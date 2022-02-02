package routes

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gwuhaolin/livego/configure"
)

func GetLiveChannels(c *gin.Context) {
	startNum := uint64(0)
	cursor := c.Query("cursor")
	if cursor != "" {
		newStartNum, err := strconv.ParseUint(cursor, 10, 64)
		if err != nil {
			c.AbortWithStatusJSON(400, gin.H{"error": "bad request"})
			return
		}
		startNum = newStartNum
	}
	endNum := startNum + 100
	keys, newCursor, err := configure.RoomKeys.GetAllLiveChannels(startNum, endNum)
	if newCursor <= startNum && startNum != 0 {
		c.AbortWithStatus(204)
		return
	}
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": "internal error"})
		return
	}
	res := gin.H{
		"cursor": newCursor,
		"keys":   keys,
		"error":  "ok",
	}
	c.JSON(200, res)
}

func GetLiveChannel(c *gin.Context) {
	channel := c.Param("id")
	channel, err := configure.RoomKeys.GetLiveChannel(channel)
	if err != nil {
		c.AbortWithStatusJSON(404, gin.H{"error": "not found"})
		return
	}
	c.JSON(200, gin.H{"channel": channel, "error": "ok"})
}
