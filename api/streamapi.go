package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func StartStreamApi(addr string) {
	fmt.Print("Starting KamaiiTV")
	router := gin.New()
	defer router.Run(addr)
	router.POST("/publish", publish)
	router.POST("/done", done)
	router.POST("/viewer/connect", viewerConnect)
	router.POST("/viewer/disconnect", viewerDisconnect)
}
