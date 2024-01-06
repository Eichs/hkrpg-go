package Gm

import (
	proto "github.com/Eichs/hkrpg-go/protocol/gmpb"
	"github.com/gin-gonic/gin"
)

func WorldLevel(c *gin.Context) {
	cmd := stou32(c.Query("cmd"))
	uid := stou32(c.Query("uid"))
	worldLevel := stou32(c.Query("world_level"))
	if worldLevel < 0 || worldLevel > 6 || uid == 0 {
		c.JSON(404, gin.H{
			"code": -1,
		})
		return
	}

	message := &proto.GmWorldLevel{
		WorldLevel: worldLevel,
	}

	ToGate(c, uid, cmd, message)
}
