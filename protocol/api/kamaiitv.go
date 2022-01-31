package api

import (
	"github.com/gin-gonic/gin"
)

var GinServer = &gin.Engine{}

func StartKamaiiTV() {
	compileRegex()
	server := gin.Default()
	defer server.Run(":8080")
	GinServer = server
	server.POST("/api/v1/user/register", RegisterUser)
}
