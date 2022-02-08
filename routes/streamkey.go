package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gwuhaolin/livego/configure"
)

func GetStreamKey(c *gin.Context) {
	id := c.GetString("id")
	if id == "" {
		c.AbortWithStatusJSON(500, gin.H{"error": "internal error"})
		return
	}
	key, err := configure.RoomKeys.GetKey(id)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": "internal error"})
		return
	}
	c.JSON(200, gin.H{"key": key, "error": "ok"})
}

func DeleteStreamKey(c *gin.Context) {
	key := c.Param("key")
	configure.RoomKeys.DeleteKey(key)
	c.JSON(200, gin.H{"error": "ok"})
}
