package Game

import (
	"encoding/json"
	"time"

	"github.com/Eichs/hkrpg-go/internal/DataBase"
	"github.com/Eichs/hkrpg-go/pkg/alg"
	"github.com/Eichs/hkrpg-go/pkg/kcp"
	"github.com/Eichs/hkrpg-go/pkg/logger"
	"github.com/Eichs/hkrpg-go/protocol/cmd"
	"google.golang.org/protobuf/encoding/protojson"
	pb "google.golang.org/protobuf/proto"
)

type Game struct {
	Uid         uint32
	Seed        uint64
	NetMsgInput chan *NetMsg
	KcpConn     *kcp.UDPSession
	Db          *DataBase.Store
	Snowflake   *alg.SnowflakeWorker // 雪花唯一id生成器
	// 玩家数据
	Player *PlayerData
	// 密钥
	XorKey []byte
}

type NetMsg struct {
	G         *Game
	CmdId     uint16
	PlayerMsg pb.Message
	Type      string
}

func (g *Game) send(cmdid uint16, playerMsg pb.Message) {
	data := protojson.Format(playerMsg)
	logger.Debug("[UID:%v] S --> C : CmdId: %v KcpMsg: \n%s\n", g.Uid, cmdid, data)
	netMsg := new(NetMsg)
	netMsg.G = g
	netMsg.CmdId = cmdid
	netMsg.PlayerMsg = playerMsg
	netMsg.Type = "KcpMsg"
	g.NetMsgInput <- netMsg
}

func (g *Game) decodePayloadToProto(cmdId uint16, msg []byte) (protoObj pb.Message) {
	protoObj = cmd.GetSharedCmdProtoMap().GetProtoObjCacheByCmdId(cmdId)
	if protoObj == nil {
		logger.Error("get new proto object is nil")
		return nil
	}
	err := pb.Unmarshal(msg, protoObj)
	if err != nil {
		logger.Error("unmarshal proto data err: %v", err)
		return nil
	}
	data := protojson.Format(protoObj)
	logger.Debug("[UID:%v] C --> S : NAME: %s KcpMsg: \n%s\n", g.Uid, cmd.GetSharedCmdProtoMap().GetCmdNameByCmdId(cmdId), data)
	return protoObj
}

func (g *Game) UpDataPlayer() {
	var err error
	if g.KcpConn == nil {
		return
	}
	if g.Uid == 0 {
		return
	}
	dbDate := new(DataBase.Player)
	dbDate.AccountUid = g.Uid
	dbDate.PlayerData, err = json.Marshal(g.Player)
	if err != nil {
		logger.Error("json to bin error:%s", err)
		return
	}
	if err = g.Db.UpdatePlayer(dbDate); err != nil {
		logger.Error("Update Player error")
		return
	}
	logger.Info("数据库账号:%v 数据更新", g.Uid)
}

func (g *Game) AutoUpDataPlayer() {
	ticker := time.NewTicker(time.Second * 60)
	for {
		<-ticker.C
		if g.KcpConn == nil {
			return
		}
		if g.Db == nil {
			return
		}
		if g.Uid == 0 {
			continue
		}
		logger.Info("[UID:%v] || 定时保存在线玩家数据", g.Uid)
		g.UpDataPlayer()
	}
}

func (g *Game) exitGame() {
	g.UpDataPlayer()
	logger.Info("[UID:%v] || 玩家已离线", g.Uid)
	netMsg := new(NetMsg)
	netMsg.G = g
	netMsg.Type = "Close"
	g.NetMsgInput <- netMsg
	g.Db = nil
	return
}
