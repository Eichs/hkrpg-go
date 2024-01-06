package Game

import (
	"time"

	"github.com/Eichs/hkrpg-go/gdconf"
	"github.com/Eichs/hkrpg-go/protocol/cmd"
	"github.com/Eichs/hkrpg-go/protocol/proto"
)

func (g *Game) StaminaInfoScNotify() {
	notify := &proto.StaminaInfoScNotify{
		NextRecoverTime: 0,
		Stamina:         g.Player.DbItem.MaterialMap[11].Num,
		ReserveStamina:  g.Player.DbItem.MaterialMap[12].Num,
	}
	g.Send(cmd.StaminaInfoScNotify, notify)
}

func (g *Game) HandleGetBasicInfoCsReq() {
	rsp := new(proto.GetBasicInfoScRsp)
	rsp.CurDay = 1
	rsp.NextRecoverTime = 1698768000
	rsp.GameplayBirthday = g.Player.Birthday
	rsp.PlayerSettingInfo = &proto.PlayerSettingInfo{}

	g.Send(cmd.GetBasicInfoScRsp, rsp)
}

func (g *Game) HandleGetArchiveDataCsReq(payloadMsg []byte) {
	rsp := new(proto.GetArchiveDataScRsp)
	var archiveAvatarIdList []uint32
	for _, a := range g.Player.DbAvatar.Avatar {
		archiveAvatarIdList = append(archiveAvatarIdList, a.AvatarId)
	}
	archiveData := &proto.ArchiveData{
		ArchiveAvatarIdList:    make([]uint32, 0),
		ArchiveEquipmentIdList: make([]uint32, 0),
		ArchiveMonsterIdList:   make([]*proto.MonsterArchive, 0),
		RelicList:              make([]*proto.RelicArchive, 0),
	}
	for _, monsterList := range gdconf.GetMonsterConfigMap() {
		archiveMonsterIdList := &proto.MonsterArchive{
			Num:       1,
			MonsterId: monsterList.MonsterID,
		}
		archiveData.ArchiveMonsterIdList = append(archiveData.ArchiveMonsterIdList, archiveMonsterIdList)
	}
	for _, relicList := range gdconf.GetRelicMap() {
		archiveRelicList := &proto.RelicArchive{
			RelicId: relicList.ID,
			Slot:    relicList.Type,
		}
		archiveData.RelicList = append(archiveData.RelicList, archiveRelicList)
	}

	rsp.ArchiveData = archiveData

	g.Send(cmd.GetArchiveDataScRsp, rsp)
}

func (g *Game) HandleGetPlayerBoardDataCsReq(payloadMsg []byte) {
	rsp := &proto.GetPlayerBoardDataScRsp{
		CurrentHeadIconId:    g.Player.HeadImage,
		UnlockedHeadIconList: make([]*proto.HeadIcon, 0),
		Signature:            g.Player.Signature,
		//	Unk1:                 "",
	}

	for _, avatar := range g.Player.DbItem.HeadIcon {
		headIcon := &proto.HeadIcon{
			Id: avatar,
		}
		rsp.UnlockedHeadIconList = append(rsp.UnlockedHeadIconList, headIcon)
	}

	g.Send(cmd.GetPlayerBoardDataScRsp, rsp)
}

func (g *Game) SetHeadIconCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.SetHeadIconCsReq, payloadMsg)
	req := msg.(*proto.SetHeadIconCsReq)

	g.Player.HeadImage = req.Id

	rsp := &proto.SetHeadIconScRsp{
		CurrentHeadIconId: req.Id,
	}

	g.Send(cmd.SetHeadIconScRsp, rsp)

}

func (g *Game) SetHeroBasicTypeCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.SetHeroBasicTypeCsReq, payloadMsg)
	req := msg.(*proto.SetHeroBasicTypeCsReq)

	g.Player.DbAvatar.MainAvatar = req.BasicType

	rsp := &proto.SetHeroBasicTypeScRsp{
		BasicType: req.BasicType,
	}

	g.Send(cmd.SetHeroBasicTypeScRsp, rsp)

}

func (g *Game) HandleGetFriendLoginInfoCsReq(payloadMsg []byte) {
	rsp := new(proto.GetFriendLoginInfoScRsp)
	rsp.FriendUidList = []uint32{99}

	g.Send(cmd.GetFriendLoginInfoScRsp, rsp)
}

func (g *Game) HandleGetRogueHandbookDataCsReq(payloadMsg []byte) {
	rsp := new(proto.GetRogueHandbookDataScRsp)
	handbookInfo := &proto.RogueHandbookData{
		RogueCurrentVersion: 1,
		IsMiracleUnlock:     true,
	}
	rsp.HandbookInfo = handbookInfo

	g.Send(cmd.GetRogueHandbookDataScRsp, rsp)
}

func (g *Game) HandleGetChallengeCsReq(payloadMsg []byte) {
	rsp := new(proto.GetChallengeScRsp)
	rsp.ChallengeList = make([]*proto.Challenge, 0)
	for _, challengeList := range gdconf.GetChallengeMazeConfigMap() {
		challenge := &proto.Challenge{
			ChallengeId: challengeList.ID,
		}
		rsp.ChallengeList = append(rsp.ChallengeList, challenge)
	}
	g.Send(cmd.GetChallengeScRsp, rsp)
}

func (g *Game) HandleGetChatEmojiListCsReq(payloadMsg []byte) {
	rsp := new(proto.GetChallengeScRsp)
	// TODO 是的，没错，还是同样的原因
	g.Send(cmd.GetChatEmojiListScRsp, rsp)
}

func (g *Game) HandleGetAssistHistoryCsReq() {
	rsp := new(proto.GetChallengeScRsp)
	// TODO 是的，没错，还是同样的原因
	g.Send(cmd.GetAssistHistoryScRsp, rsp)
}

func (g *Game) GetMailCsReq() {
	rsp := new(proto.GetMailScRsp)
	rsp.TotalNum = 1
	rsp.IsEnd = true
	mailList := &proto.ClientMail{
		Sender:  "Eichs",
		Content: "欢迎来到 hkrpg-go server",
		Title:   "欢迎来到 hkrpg-go server",
	}
	rsp.MailList = append(rsp.MailList, mailList)

	g.Send(cmd.GetMailScRsp, rsp)
}

func (g *Game) SetClientPausedCsReq() {
	rsp := new(proto.SetClientPausedScRsp)
	g.Player.IsPaused = !g.Player.IsPaused
	rsp.Paused = g.Player.IsPaused

	g.Send(cmd.SetClientPausedScRsp, rsp)
}

func (g *Game) GetFirstTalkNpcCsReq() {
	rsp := new(proto.GetChallengeScRsp)
	// TODO 是的，没错，还是同样的原因
	g.Send(cmd.GetFirstTalkNpcScRsp, rsp)
}

func (g *Game) HandleGetJukeboxDataCsReq(payloadMsg []byte) {
	rsp := new(proto.GetJukeboxDataScRsp)
	rsp.PlayingId = 210000
	rsp.MusicList = make([]*proto.UnlockedMusic, 0)
	for _, backMusicList := range gdconf.GetBackGroundMusicMap() {
		musicList := &proto.UnlockedMusic{
			GroupId: backMusicList.GroupID,
			Unkbool: true,
			Id:      backMusicList.ID,
		}
		rsp.MusicList = append(rsp.MusicList, musicList)
	}
	g.Send(cmd.GetJukeboxDataScRsp, rsp)
}

func (g *Game) HandleGetPhoneDataCsReq(payloadMsg []byte) {
	rsp := new(proto.GetPhoneDataScRsp)
	rsp.CurChatBubble = 220000
	rsp.CurPhoneTheme = 221000
	rsp.OwnedChatBubbles = []uint32{220002, 220000, 220001}
	rsp.OwnedPhoneThemes = []uint32{221000, 221001, 221002, 221003}

	g.Send(cmd.GetPhoneDataScRsp, rsp)
}

func (g *Game) SetNicknameCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.SetNicknameCsReq, payloadMsg)
	req := msg.(*proto.SetNicknameCsReq)

	if g.Player.IsNickName {
		g.Player.NickName = req.Nickname
	}

	g.Player.IsNickName = !g.Player.IsNickName

	rsp := new(proto.GetChallengeScRsp)
	// TODO 是的，没错，还是同样的原因

	g.PlayerPlayerSyncScNotify()
	g.Send(cmd.SetNicknameScRsp, rsp)
}

func (g *Game) SetGameplayBirthdayCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.SetGameplayBirthdayCsReq, payloadMsg)
	req := msg.(*proto.SetGameplayBirthdayCsReq)

	g.Player.Birthday = req.Birthday

	rsp := &proto.SetGameplayBirthdayScRsp{Birthday: req.Birthday}

	g.Send(cmd.SetGameplayBirthdayScRsp, rsp)
}

func (g *Game) SetSignatureCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.SetSignatureCsReq, payloadMsg)
	req := msg.(*proto.SetSignatureCsReq)

	g.Player.Signature = req.Signature

	rsp := &proto.SetSignatureScRsp{Signature: req.Signature}

	g.Send(cmd.SetSignatureScRsp, rsp)
}

func (g *Game) HandlePlayerHeartBeatCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.PlayerHeartBeatCsReq, payloadMsg)
	req := msg.(*proto.PlayerHeartbeatCsReq)

	rsp := new(proto.PlayerHeartbeatScRsp)
	rsp.ServerTimeMs = uint64(time.Now().UnixNano() / 1e6)
	rsp.ClientTimeMs = req.ClientTimeMs

	g.Send(cmd.PlayerHeartBeatScRsp, rsp)
}

func (g *Game) InteractPropCsReq() {
	rsp := new(proto.InteractPropScRsp)

	g.Send(cmd.InteractPropScRsp, rsp)
}

func (g *Game) TextJoinQueryCsReq() {
	rsp := new(proto.TextJoinQueryScRsp)
	for _, textJoin := range gdconf.GetTextJoinConfigMap() {
		textJoinList := &proto.TextJoinInfo{
			TextItemId:       textJoin.TextJoinID,
			TextItemConfigId: textJoin.TextJoinItemList[len(textJoin.TextJoinItemList)-1],
		}
		rsp.TextJoinList = append(rsp.TextJoinList, textJoinList)
	}

	g.Send(cmd.TextJoinQueryScRsp, rsp)
}
