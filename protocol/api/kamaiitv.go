package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gwuhaolin/livego/routes"
)

var GinServer = &gin.Engine{}

func StartKamaiiTV() {
	server := gin.Default()
	defer server.Run(":8080")
	GinServer = server
	server.POST("/api/v1/user/register", routes.RegisterUser)
	server.GET("/api/v1/user/login", routes.LoginUser)
	server.GET("/api/v1/user/key", routes.GetUserKey)
}
