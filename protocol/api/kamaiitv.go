package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gwuhaolin/livego/mongo"
	"github.com/gwuhaolin/livego/routes"
)

var GinServer = &gin.Engine{}

func StartKamaiiTV() {
	mongo.Connect()
	routes.CompileRegexp()
	routes.GetSecret()
	server := gin.Default()
	defer server.Run(":8080")
	GinServer = server
	server.POST("/api/v1/user/register", routes.RegisterUser)
	server.GET("/api/v1/user/login", routes.LoginUser)
	server.GET("/api/v1/channels/live", routes.GetLiveChannels)
	server.GET("/api/v1/channel/live/:id", routes.GetLiveChannel)
	server.Use(routes.TokenAuthMiddleware())
	server.GET("/api/v1/user/streamkey", routes.GetStreamKey)
	server.DELETE("/api/v1/user/streamkey/:key", routes.DeleteStreamKey)
}
