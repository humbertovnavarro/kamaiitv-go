package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func StartStreamApi(addr string) {
	fmt.Println("Starting KamaiiTV stream api")
	router := gin.New()
	defer router.Run(addr)
	router.POST("/publish", publish)
	router.POST("/done", done)
}

func StartPublicApi(addr string) {
	fmt.Println("Starting KamaiiTV public api")
	router := gin.New()
	defer router.Run(addr)
}
