package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gwuhaolin/livego/configure"
	"github.com/gwuhaolin/livego/lib"
	"github.com/gwuhaolin/livego/mongo"
	"github.com/gwuhaolin/livego/protocol/socketio"
)

func StartKamaiiTV() {
	go listen()
}

type RequestUser struct {
	ID       string
	Username string
}

func listen() {
	lib.CompileRegexp()
	lib.GetSecret()
	go mongo.Connect(configure.Config.GetString("mongo_addr"))
	go socketio.Start(configure.Config.GetString("redis_addr"))
	server := gin.New()
	defer server.Run(":8080")
	// HTTP Proxy NextJS
	// Authenticated Routes
	authenticated := TokenAuthMiddleware()
	server.POST("/api/v1/channel/:id/follow", FollowChannel, authenticated)
	server.POST("/api/v1/user/message", PostMessage, authenticated)
	server.GET("/api/v1/user/streamkey", GetStreamKey, authenticated)
	server.GET("/api/v1/user/following", GetFollowing, authenticated)
	server.DELETE("/api/v1/user/streamkey/:key", DeleteStreamKey, authenticated)
	// Public APIs
	server.POST("/api/v1/user/register", RegisterUser)
	server.POST("/api/v1/user/login", LoginUser)
	server.GET("/api/v1/channels/live", GetLiveChannels)
	server.GET("/api/v1/channel/:id/live", GetLiveChannel)
	server.GET("/api/v1/channel/:id/messages", GetChannelMessage)
	server.Use(ReverseProxy("http://localhost:3000"))
}
