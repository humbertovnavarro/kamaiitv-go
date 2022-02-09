package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gwuhaolin/livego/mongo"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

func GetFollowing(c *gin.Context) {
	context := c.MustGet("user").(*RequestUser)
	filter := bson.M{
		"follower": context.ID,
	}

	query, err := mongo.FollowerCollection.Find(c, filter)
	if err != nil {
		log.Info("GetFollowing find error: %v", err)
		c.AbortWithStatusJSON(500, gin.H{"error": "internal error"})
		return
	}
	following := []string{}
	for query.Next(c) {
		var follower mongo.Follower
		if err := query.Decode(&follower); err != nil {
			c.AbortWithStatusJSON(500, gin.H{"error": "internal error"})
			return
		}
		following = append(following, follower.Channel)
	}
	c.JSON(200, gin.H{"following": following})
}
