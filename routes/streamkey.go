package routes

import (
	"github.com/gin-gonic/gin"
)

func GetStreamKey(c *gin.Context) {
	token, err := extractToken(c)
	if err != nil {
		c.JSON(401, err.Error())
		c.Abort()
		return
	}
	token.Claims["id"]
}
