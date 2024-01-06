package Gm

import (
	"github.com/Eichs/hkrpg-go/internal/Net"
	"github.com/gin-gonic/gin"
)

var (
	err error
)

func State(c *gin.Context) {
	c.JSON(200, gin.H{
		"在线玩家": Net.CLIENT_CONN_NUM,
	})
}
