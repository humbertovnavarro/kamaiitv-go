package api

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gwuhaolin/livego/configure"
	"github.com/gwuhaolin/livego/lib"
	"github.com/gwuhaolin/livego/mongo"
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
	endNum := startNum + 50
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
	channel := c.Param("id")
	page := c.Query("page")
	if page == "" {
		page = "0"
	}
	pageNum, err := strconv.Atoi(page)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "bad request"})
		return
	}
	filter := bson.M{
		"toRoom": channel + ":public",
	}
	findOpts := options.Find()
	findOpts.SetSort(bson.M{"createdAt": 1})
	findOpts.SetLimit(50)
	findOpts.SetSkip(int64(pageNum) * 50)
	query, err := mongo.MessageCollection.Find(c, filter, findOpts)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": "internal error"})
		return
	}
	results := &[]mongo.Message{}
	err = query.All(c, results)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": "internal error"})
		return
	}
	c.JSON(200, gin.H{"messages": results})
}

func FollowChannel(c *gin.Context) {
	channel := c.Param("id")
	context := c.MustGet("user").(*RequestUser)
	if channel == "" {
		c.AbortWithStatusJSON(400, gin.H{"error": "bad request"})
		return
	}
	if !lib.IsValidObjectID.MatchString(channel) {
		c.AbortWithStatusJSON(400, gin.H{"error": "bad request"})
		return
	}
	if context.ID == channel {
		c.AbortWithStatusJSON(400, gin.H{"error": "can not follow/unfollow yourself"})
		return
	}
	filter := bson.M{
		"follower": context.ID,
		"channel":  channel,
	}
	result := mongo.FollowerCollection.FindOneAndDelete(c, filter)
	if result.Err() == nil {
		c.JSON(200, gin.H{"error": "deleted"})
		return
	}
	mongo.FollowerCollection.InsertOne(c, filter)
	c.JSON(200, gin.H{"error": "ok"})
}
