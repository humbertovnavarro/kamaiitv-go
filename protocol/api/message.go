package api

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gwuhaolin/livego/mongo"
	"github.com/gwuhaolin/livego/protocol/socketio"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

type IngressMessage struct {
	ToRoom  string `json:"toRoom"`
	Content string `json:"content"`
	Private bool   `json:"private"`
}

type EgressMessage struct {
	ToRoom   string `json:"toRoom"`
	FromId   string `json:"fromId"`
	FromName string `json:"fromName"`
	Content  string `json:"content"`
}

func PostMessage(c *gin.Context) {
	iMessage := &IngressMessage{}
	if err := c.BindJSON(iMessage); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "json syntax error"})
		return
	}
	if len(iMessage.ToRoom) == 0 || len(iMessage.Content) == 0 || len(iMessage.Content) > 500 {
		c.AbortWithStatusJSON(400, gin.H{"error": "bad request"})
		return
	}
	user := c.MustGet("user").(*RequestUser)
	eMessage := &EgressMessage{
		ToRoom:   iMessage.ToRoom,
		FromId:   user.ID,
		FromName: user.Username,
		Content:  iMessage.Content,
	}
	if !iMessage.Private {
		eMessage.ToRoom = iMessage.ToRoom + ":public"
	}
	_, err := mongo.MessageCollection.InsertOne(c, &bson.M{
		"toRoom":    eMessage.ToRoom,
		"fromId":    eMessage.FromId,
		"fromName":  eMessage.FromName,
		"content":   eMessage.Content,
		"createdAt": time.Now().Unix(),
		"deleted":   false,
	})
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": "internal error"})
		log.Errorf("PostMessage insert error: %v", err)
		return
	}
	socketio.IO.BroadcastToRoom("/", eMessage.ToRoom, "message", eMessage)
	c.JSON(200, gin.H{"error": "sent to " + eMessage.ToRoom})
}
