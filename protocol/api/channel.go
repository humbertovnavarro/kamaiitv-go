package api

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gwuhaolin/livego/configure"
	"github.com/gwuhaolin/livego/mongo"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	viewers, err := configure.RoomKeys.GetLiveChannel(channel)
	if err != nil {
		c.AbortWithStatusJSON(404, gin.H{"error": "not found"})
		return
	}
	c.JSON(200, gin.H{"viewers": viewers})
}

func GetChannelMessage(c *gin.Context) {
	oldest := time.Now().Unix() - int64(time.Hour)*24
	channel := c.Param("id")
	filter := bson.M{
		"createdAt": bson.M{
			"$gt": oldest,
		},
		"toRoom": channel + ":public",
	}
	findOpts := options.Find()
	findOpts.SetSort(bson.M{"createdAt": 1})
	query, err := mongo.MessageCollection.Find(c, filter, findOpts)
	if err != nil {
		log.Errorf("GetChannelMessage find error: %v", err)
		c.AbortWithStatusJSON(500, gin.H{"error": "internal error"})
		return
	}
	results := &[]mongo.Message{}
	err = query.All(c, results)
	if err != nil {
		log.Errorf("GetChannelMessage find error: %v", err)
		c.AbortWithStatusJSON(500, gin.H{"error": "internal error"})
		return
	}
	c.JSON(200, gin.H{"messages": results})
}
