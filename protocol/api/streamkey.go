package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gwuhaolin/livego/configure"
)

func GetStreamKey(c *gin.Context) {
	user := c.MustGet("user").(*RequestUser)
	key, err := configure.RoomKeys.GetKey(user.ID)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": "internal error"})
		return
	}
	c.JSON(200, gin.H{"key": key})
}

func DeleteStreamKey(c *gin.Context) {
	key := c.Param("key")
	configure.RoomKeys.DeleteKey(key)
	c.JSON(200, gin.H{"error": "ok"})
}
