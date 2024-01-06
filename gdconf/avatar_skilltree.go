package gdconf

import (
	"fmt"
	"os"

	"github.com/Eichs/hkrpg-go/pkg/logger"
	"github.com/hjson/hjson-go/v4"
)

type AvatarSkilltree struct {
	PointID  uint32 `json:"PointID"`
	Level    uint32 `json:"Level"`
	MaxLevel uint32 `json:"MaxLevel"`
	AvatarID uint32 `json:"AvatarID"`
}

func (g *GameDataConfig) loadAvatarSkilltree() {
	g.AvatarSkilltreeMap = make(map[string]map[string]*AvatarSkilltree)
	playerElementsFilePath := g.excelPrefix + "AvatarSkillTreeConfig.json"
	playerElementsFile, err := os.ReadFile(playerElementsFilePath)
	if err != nil {
		info := fmt.Sprintf("open file error: %v", err)
		panic(info)
	}

	err = hjson.Unmarshal(playerElementsFile, &g.AvatarSkilltreeMap)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}
	logger.Info("load %v AvatarSkilltree", len(g.AvatarSkilltreeMap))
}

func GetAvatarSkilltreeById(avatarId, level string) *AvatarSkilltree {
	return CONF.AvatarSkilltreeMap[avatarId][level]
}

func GetAvatarSkilltreeMap() map[string]map[string]*AvatarSkilltree {
	return CONF.AvatarSkilltreeMap
}
