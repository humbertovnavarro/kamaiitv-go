package api

import (
	"github.com/gin-contrib/static"
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
	go mongo.Connect(configure.Config.GetString("mongo_addr"))
	CompileRegexp()
	lib.GetSecret()
	server := gin.New()
	go socketio.Start(configure.Config.GetString("redis_addr"))
	server.Use(static.Serve("/", static.LocalFile("./public", true)))
	// Public APIs
	server.POST("/api/v1/user/register", RegisterUser)
	server.POST("/api/v1/user/login", LoginUser)
	server.GET("/api/v1/channels/live", GetLiveChannels)
	server.GET("/api/v1/channel/:id/live", GetLiveChannel)
	server.GET("/api/v1/channel/:id/messages", GetChannelMessage)
	// Authenticated Routes
	server.Use(TokenAuthMiddleware())
	server.POST("/api/v1/user/message", PostMessage)
	server.GET("/api/v1/user/streamkey", GetStreamKey)
	server.DELETE("/api/v1/user/streamkey/:key", DeleteStreamKey)
	server.Run(":8080")
}
