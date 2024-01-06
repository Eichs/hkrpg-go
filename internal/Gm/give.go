package Gm

import (
	proto "github.com/Eichs/hkrpg-go/protocol/gmpb"
	"github.com/gin-gonic/gin"
)

func Give(c *gin.Context) {
	cmd := stou32(c.Query("cmd"))
	uid := stou32(c.Query("uid"))
	all := stou32(c.Query("all"))
	itemId := c.Query("item_id")
	itemCount := c.Query("item_count") // 数量

	message := &proto.GmGive{
		ItemId:    stou32(itemId),
		ItemCount: stou32(itemCount),
	}
	if all == 1 {
		message.GiveAll = true
	}

	ToGate(c, uid, cmd, message)
}
