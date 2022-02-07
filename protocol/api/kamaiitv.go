package api

import (
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/gwuhaolin/livego/configure"
	"github.com/gwuhaolin/livego/mongo"
	"github.com/gwuhaolin/livego/protocol/socketio"
	"github.com/gwuhaolin/livego/routes"
)

func StartKamaiiTV() {
	go listen()
}

func listen() {
	mongo.Connect(configure.Config.GetString("mongo_addr"))
	routes.CompileRegexp()
	routes.GetSecret()
	server := gin.New()
	go socketio.Start(configure.Config.GetString("redis_addr"))
	server.Use(static.Serve("/", static.LocalFile("./public", true)))
	// Public APIs
	server.POST("/api/v1/user/register", routes.RegisterUser)
	server.POST("/api/v1/user/login", routes.LoginUser)
	server.GET("/api/v1/channels/live", routes.GetLiveChannels)
	server.GET("/api/v1/channel/live/:id", routes.GetLiveChannel)
	// Authenticated Routes
	server.Use(routes.TokenAuthMiddleware())
	server.GET("/api/v1/user/streamkey", routes.GetStreamKey)
	server.DELETE("/api/v1/user/streamkey/:key", routes.DeleteStreamKey)
	server.Run(":8080")
}
