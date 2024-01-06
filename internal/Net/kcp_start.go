/*
好用的kcp
爱 来自 hk4e-go
*/
package Net

import (
	"encoding/binary"
	"strconv"
	"sync/atomic"
	"time"

	"github.com/Eichs/hkrpg-go/internal/DataBase"
	"github.com/Eichs/hkrpg-go/internal/Game"
	"github.com/Eichs/hkrpg-go/pkg/alg"
	"github.com/Eichs/hkrpg-go/pkg/config"
	"github.com/Eichs/hkrpg-go/pkg/kcp"
	"github.com/Eichs/hkrpg-go/pkg/logger"
	"github.com/Eichs/hkrpg-go/pkg/random"
	"github.com/Eichs/hkrpg-go/protocol/cmd"
	pb "google.golang.org/protobuf/proto"
)

const (
	PacketMaxLen            = 343 * 1024 // 最大应用层包长度
	KcpConnAddrChangeNotify = "KcpConnAddrChangeNotify"
)

var CLIENT_CONN_NUM int32 = 0 // 当前客户端连接数
var KCPCONNMANAGER *KcpConnManager

type KcpConnManager struct {
	kcpListener *kcp.Listener
	// 会话
	sessionIdCounter uint32
	sessionMap       map[uint32]*Game.Game
	// 事件
	kcpEventChan chan *KcpEvent
}

type KcpEvent struct {
	SessionId    uint32
	EventId      string
	EventMessage any
}

type GmMsg struct {
	CmdId     uint16
	ProtoData []byte
}

func Run() error {
	k := KcpConnManager{}
	k.sessionMap = make(map[uint32]*Game.Game)

	addr := "0.0.0.0:" + strconv.Itoa(int(config.GetConfig().Game.Port))
	kcpListener, err := kcp.ListenWithOptions(addr)
	if err != nil {
		logger.Error("listen kcp err: %v", err)
		return err
	}
	logger.Info("kcp服务在 %s 上启动", addr)
	k.kcpListener = kcpListener
	go k.kcpNetInfo()
	go k.kcpEnetHandle(kcpListener)

	for {
		kcpConn, err := kcpListener.AcceptKCP()
		if err != nil {
			logger.Error("accept kcp err: %v", err)
			continue
		}

		kcpConn.SetACKNoDelay(true)
		kcpConn.SetWriteDelay(false)
		kcpConn.SetWindowSize(256, 256)
		kcpConn.SetMtu(1200)
		// 读取密钥相关文件
		g := NewGame(kcpConn)
		g.XorKey = config.GetConfig().Ec2b.XorKey()
		g.Db = DataBase.DBASE
		g.Snowflake = alg.NewSnowflakeWorker(int64(1))
		CLIENT_CONN_NUM++
		go g.AutoUpDataPlayer()
		go k.recvHandle(g)
		go k.sendNet(g)
	}
}

func NewGame(kcpConn *kcp.UDPSession) *Game.Game {
	g := new(Game.Game)
	g.KcpConn = kcpConn
	g.NetMsgInput = make(chan *Game.NetMsg, 1000)
	return g
}

func (k *KcpConnManager) recvHandle(g *Game.Game) {
	payload := make([]byte, PacketMaxLen)

	for {
		var bin []byte = nil
		recvLen, err := g.KcpConn.Read(payload)
		if err != nil {
			logger.Debug("exit recv loop, conn read err: %v", err)
			return
		}
		bin = payload[:recvLen]
		kcpMsgList := make([]*KcpMsg, 0)
		DecodeBinToPayload(bin, &kcpMsgList, g.XorKey)
		for _, v := range kcpMsgList {
			// name := g.ServerCmdProtoMap.GetCmdNameByCmdId(v.CmdId)
			// logger.Error("C --> S: %v", v.CmdId)
			// payloadMsg := DecodePayloadToProto(g, v) TODO 由于 req 大部分缺失，所以不预处理数据
			g.RegisterMessage(v.CmdId, v.ProtoData)
		}
	}
}

// kcp连接事件处理函数
func (k *KcpConnManager) kcpEnetHandle(listener *kcp.Listener) {
	KCPCONNMANAGER = k
	logger.Info("kcp enet handle start")
	for {
		enetNotify := <-listener.GetEnetNotifyChan()
		logger.Info("[Kcp Enet] addr: %v, conv: %v, sessionId: %v, connType: %v, enetType: %v",
			enetNotify.Addr, enetNotify.Conv, enetNotify.SessionId, enetNotify.ConnType, enetNotify.EnetType)
		switch enetNotify.ConnType {
		case kcp.ConnEnetSyn:
			if enetNotify.EnetType != kcp.EnetClientConnectKey {
				logger.Error("enet type not match, sessionId: %v", enetNotify.SessionId)
				continue
			}
			sessionId := atomic.AddUint32(&k.sessionIdCounter, 1)
			listener.SendEnetNotifyToPeer(&kcp.Enet{
				Addr:      enetNotify.Addr,
				SessionId: sessionId,
				Conv:      binary.BigEndian.Uint32(random.GetRandomByte(4)),
				ConnType:  kcp.ConnEnetEst,
				EnetType:  enetNotify.EnetType,
			})
		case kcp.ConnEnetAddrChange:
			// 连接地址改变通知
			k.kcpEventChan <- &KcpEvent{
				SessionId:    enetNotify.SessionId,
				EventId:      KcpConnAddrChangeNotify,
				EventMessage: enetNotify.Addr,
			}
		default:
		}
	}
}

func (k *KcpConnManager) sendNet(g *Game.Game) {
	for {
		netMsg := <-g.NetMsgInput
		switch netMsg.Type {
		case "KcpMsg":
			k.SendHandle(netMsg.G, netMsg.CmdId, netMsg.PlayerMsg)
		case "Close":
			atomic.AddInt32(&CLIENT_CONN_NUM, -1)
			g.KcpConn.Close()
			delete(k.sessionMap, g.Uid)
			return
		}
	}
}

// 发送事件处理
func (k *KcpConnManager) SendHandle(g *Game.Game, cmdid uint16, playerMsg pb.Message) {
	rspMsg := new(ProtoMsg)
	rspMsg.CmdId = cmdid
	rspMsg.PayloadMessage = playerMsg
	kcpMsg := EncodeProtoToPayload(rspMsg)
	if kcpMsg.CmdId == 0 {
		logger.Error("cmdid error")
	}
	// logger.Debug("S --> C: %v", kcpMsg.CmdId)
	binMsg := EncodePayloadToBin(kcpMsg, g.XorKey)
	_, err := g.KcpConn.Write(binMsg)
	if err != nil {
		logger.Debug("exit send loop, conn write err: %v", err)
		return
	}
	// 密钥交换
	if cmdid == cmd.PlayerGetTokenScRsp {
		if g.Seed == 0 {
			return
		}
		g.XorKey = createXorPad(g.Seed)
		logger.Info("uid:%v,seed:%v,密钥交换成功", g.Uid, g.Seed)
		if k.sessionMap[g.Uid] == nil {
			k.sessionMap[g.Uid] = g
		}
		// 如果不为空则是断线重连
	}
}

func createXorPad(seed uint64) []byte {
	keyBlock := random.NewKeyBlock(seed, false)
	xorKey := keyBlock.XorKey()
	key := make([]byte, 4096)
	copy(key, xorKey[:])
	return key
}

func GmToGs(uid uint32, gmMsg *GmMsg) bool {
	if KCPCONNMANAGER.sessionMap[uid] == nil {
		return false
	}
	game := KCPCONNMANAGER.sessionMap[uid]
	payloadMsg := DecodeGmPayloadToProto(game, gmMsg)
	go game.GMRegisterMessage(gmMsg.CmdId, payloadMsg)
	return true
}

func (k *KcpConnManager) kcpNetInfo() {
	ticker := time.NewTicker(time.Second * 60)
	kcpErrorCount := uint64(0)
	for {
		<-ticker.C
		snmp := kcp.DefaultSnmp.Copy()
		kcpErrorCount += snmp.KCPInErrors
		logger.Info("kcp send: %v B/s, kcp recv: %v B/s", snmp.BytesSent/60, snmp.BytesReceived/60)
		logger.Info("udp send: %v B/s, udp recv: %v B/s", snmp.OutBytes/60, snmp.InBytes/60)
		logger.Info("udp send: %v pps, udp recv: %v pps", snmp.OutPkts/60, snmp.InPkts/60)
		clientConnNum := atomic.LoadInt32(&CLIENT_CONN_NUM)
		logger.Info("conn num: %v, new conn num: %v, kcp error num: %v", clientConnNum, snmp.CurrEstab, kcpErrorCount)
		kcp.DefaultSnmp.Reset()
	}
}
