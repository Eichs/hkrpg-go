package Game

import (
	"time"

	"github.com/Eichs/hkrpg-go/pkg/logger"
	"github.com/Eichs/hkrpg-go/protocol/cmd"
	pb "google.golang.org/protobuf/proto"
)

func (g *Game) RegisterMessage(cmdId uint16, payloadMsg []byte /*payloadMsg pb.Message*/) {
	g.LastActiveTime = time.Now().Unix()
	switch cmdId {
	case cmd.GetBasicInfoCsReq:
		g.HandleGetBasicInfoCsReq()
	case cmd.GetPlayerBoardDataCsReq:
		g.HandleGetPlayerBoardDataCsReq(payloadMsg)
	case cmd.GetCurChallengeCsReq:
		g.GetCurChallengeCsReq(payloadMsg)
	case cmd.GetEnteredSceneCsReq:
		g.HandleGetEnteredSceneCsReq(payloadMsg)
	case cmd.QueryProductInfoCsReq:
		g.HandleQueryProductInfoCsReq(payloadMsg)
	case cmd.GetFriendLoginInfoCsReq:
		g.HandleGetFriendLoginInfoCsReq(payloadMsg) // 获取好友登录信息
	case cmd.GetRogueHandbookDataCsReq:
		g.HandleGetRogueHandbookDataCsReq(payloadMsg) // 获取帮助手册
	case cmd.GetChatEmojiListCsReq:
		g.HandleGetChatEmojiListCsReq(payloadMsg) // 获取聊天表情
	case cmd.GetJukeboxDataCsReq:
		g.HandleGetJukeboxDataCsReq(payloadMsg) // 点歌？
	case cmd.GetPhoneDataCsReq:
		g.HandleGetPhoneDataCsReq(payloadMsg) // 获取手机信息?
	case cmd.TextJoinQueryCsReq:
		g.TextJoinQueryCsReq() //
	// 登录
	case cmd.PlayerLoginCsReq:
		g.HandlePlayerLoginCsReq(payloadMsg) // 玩家登录请求 第二个登录包
	case cmd.PlayerLoginFinishCsReq:
		g.HandlePlayerLoginFinishCsReq(payloadMsg) // 登录完成包
	case cmd.PlayerLogoutCsReq:
		g.PlayerLogoutCsReq() // 客户端退出游戏通知
	case cmd.GetDailyActiveInfoCsReq:
		// TODO 每日任务
	// 队伍
	case cmd.GetAllLineupDataCsReq:
		g.HandleGetAllLineupDataCsReq(payloadMsg) // 获取队伍信息请求
	case cmd.GetCurLineupDataCsReq:
		g.HandleGetCurLineupDataCsReq(payloadMsg) // 获取当前上场队伍请求
	case cmd.JoinLineupCsReq:
		g.HandleJoinLineupCsReq(payloadMsg) // 更新队伍请求
	case cmd.SwitchLineupIndexCsReq:
		g.HandleSwitchLineupIndexCsReq(payloadMsg) // 上场队伍更新请求
	case cmd.SwapLineupCsReq:
		g.HandleSwapLineupCsReq(payloadMsg) // 队伍角色交换请求
	case cmd.SetLineupNameCsReq:
		g.SetLineupNameCsReq(payloadMsg) // 修改队伍名称
	case cmd.ReplaceLineupCsReq:
		g.ReplaceLineupCsReq(payloadMsg) // 快速入队
	case cmd.ChangeLineupLeaderCsReq:
		g.ChangeLineupLeaderCsReq(payloadMsg) // 切换角色
	case cmd.QuitLineupCsReq:
		g.QuitLineupCsReq(payloadMsg) // 角色离队
	// 角色管理
	case cmd.GetHeroBasicTypeInfoCsReq:
		g.HandleGetHeroBasicTypeInfoCsReq(payloadMsg) // 请求主角基本信息
	case cmd.GetAvatarDataCsReq:
		g.HandleGetAvatarDataCsReq(payloadMsg) // 请求全部角色信息
	case cmd.RankUpAvatarCsReq:
		g.RankUpAvatarCsReq(payloadMsg) // 提高角色命座
	case cmd.AvatarExpUpCsReq:
		g.AvatarExpUpCsReq(payloadMsg) // 角色升级
	case cmd.PromoteAvatarCsReq:
		g.PromoteAvatarCsReq(payloadMsg) // 角色突破
	case cmd.UnlockSkilltreeCsReq:
		g.UnlockSkilltreeCsReq(payloadMsg) // 行迹升级
	case cmd.TakePromotionRewardCsReq:
		g.TakePromotionRewardCsReq(payloadMsg) // 领取角色突破奖励
	// 光锥
	case cmd.DressAvatarCsReq:
		g.DressAvatarCsReq(payloadMsg) // 角色光锥装备
	case cmd.ExpUpEquipmentCsReq:
		g.ExpUpEquipmentCsReq(payloadMsg) // 光锥升级
	case cmd.RankUpEquipmentCsReq:
		g.RankUpEquipmentCsReq(payloadMsg) // 光锥叠影
	case cmd.PromoteEquipmentCsReq:
		g.PromoteEquipmentCsReq(payloadMsg) // 光锥突破
	// 场景
	case cmd.GetSceneMapInfoCsReq:
		g.HanldeGetSceneMapInfoCsReq(payloadMsg) // 获取地图信息
	case cmd.GetCurSceneInfoCsReq:
		g.HandleGetCurSceneInfoCsReq(payloadMsg) // 获取场景信息(关键包)
	case cmd.SceneEntityMoveCsReq:
		g.SceneEntityMoveCsReq(payloadMsg) // 场景实体移动
	case cmd.GetRogueScoreRewardInfoCsReq:
		g.GetRogueScoreRewardInfoCsReq()
	case cmd.EnterSceneCsReq:
		g.EnterSceneCsReq(payloadMsg) // 场景传送
	// 战斗
	case cmd.GetChallengeCsReq:
		g.HandleGetChallengeCsReq(payloadMsg) // 获取挑战id列表
	case cmd.SceneCastSkillCsReq:
		g.SceneCastSkillCsReq(payloadMsg) // 场景开启战斗
	case cmd.PVEBattleResultCsReq:
		g.PVEBattleResultCsReq(payloadMsg) // PVE战斗结算
	case cmd.GetRogueInfoCsReq:
		g.GetRogueInfoCsReq(payloadMsg) // 获取模拟宇宙
	case cmd.StartRogueCsReq:
		g.StartRogueCsReq(payloadMsg) // 模拟宇宙,启动!
	case cmd.LeaveRogueCsReq:
		g.LeaveRogueCsReq(payloadMsg) // 模拟宇宙撤离请求
	case cmd.QuitRogueCsReq:
		g.QuitRogueCsReq(payloadMsg) // 模拟宇宙结算请求
	case cmd.GetRogueTalentInfoCsReq:
		g.GetRogueTalentInfoCsReq() // 获取天赋信息
	case cmd.StartCocoonStageCsReq:
		g.StartCocoonStageCsReq(payloadMsg) // 副本/周本等

	// 背包
	case cmd.GetBagCsReq:
		g.HandleGetBagCsReq(payloadMsg) // 获取背包物品
	// 交易
	case cmd.GetShopListCsReq:
		g.GetShopListCsReq() // 获取商店物品列表
	case cmd.ExchangeHcoinCsReq:
		g.ExchangeHcoinCsReq(payloadMsg) // 梦华兑换
	// 社交
	case cmd.GetMailCsReq:
		g.GetMailCsReq() // 获取邮件
	// 卡池
	case cmd.GetGachaInfoCsReq:
		g.HandleGetGachaInfoCsReq(payloadMsg) // 获取卡池信息
	case cmd.DoGachaCsReq:
		g.DoGachaCsReq(payloadMsg) // 抽卡请求
	case cmd.GetGachaCeilingCsReq:
		g.HandleGetGachaCeilingCsReq(payloadMsg) // 基础卡池保底达到进度请求
	// 任务
	case cmd.GetQuestDataCsReq:
		g.GetQuestDataCsReq(payloadMsg) // 获取任务信息
	case cmd.GetMissionStatusCsReq:
		g.HandleGetMissionStatusCsReq(payloadMsg)
	// 活动
	case cmd.GetActivityScheduleConfigCsReq:
		g.HandleGetActivityScheduleConfigCsReq(payloadMsg) // 活动配置请求
	// 基础
	case cmd.SetClientPausedCsReq:
		g.SetClientPausedCsReq() // 客户端暂停请求
	case cmd.PlayerHeartBeatCsReq:
		g.HandlePlayerHeartBeatCsReq(payloadMsg) // 心跳包
	case cmd.SyncClientResVersionCsReq:
		g.SyncClientResVersionCsReq(payloadMsg) // 版本同步
	case cmd.GetAssistHistoryCsReq:
		g.HandleGetAssistHistoryCsReq() // 漫游签证
	case cmd.SetHeadIconCsReq:
		g.SetHeadIconCsReq(payloadMsg) // 切换头像
	case cmd.SetHeroBasicTypeCsReq:
		g.SetHeroBasicTypeCsReq(payloadMsg) // 切换主角类型
	case cmd.SetNicknameCsReq:
		g.SetNicknameCsReq(payloadMsg) // 修改昵称请求
	case cmd.SetGameplayBirthdayCsReq:
		g.SetGameplayBirthdayCsReq(payloadMsg) // 修改生日请求
	case cmd.SetSignatureCsReq:
		g.SetSignatureCsReq(payloadMsg) // 简介修改请求
	// 成就
	case cmd.GetArchiveDataCsReq:
		g.HandleGetArchiveDataCsReq(payloadMsg) // 获取成就
	// 乱七八糟
	case cmd.InteractPropCsReq:
		g.InteractPropCsReq()
	case cmd.GetFirstTalkNpcCsReq:
		g.GetFirstTalkNpcCsReq()
	default:
		logger.Debug("C --> S error router: %v", cmdId)
	}
	return
}

func (g *Game) GMRegisterMessage(cmdId uint16, payloadMsg pb.Message) {
	switch cmdId {
	case cmd.GmGive:
		g.GmGive(payloadMsg) // 获取物品
	case cmd.GmWorldLevel:
		g.GmWorldLevel(payloadMsg) // 设置世界等级
	}
}
