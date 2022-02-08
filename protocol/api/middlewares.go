package api

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gwuhaolin/livego/lib"
)

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := c.Cookie("token")
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": "unauthorized"})
			return
		}
		claims, err := lib.DecodeToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": "unauthorized"})
			return
		}
		if claims.ExpiresAt < time.Now().Unix() {
			c.AbortWithStatusJSON(401, gin.H{"error": "expired"})
			return
		}
		context := &RequestUser{
			ID:       claims.Id,
			Username: claims.Subject,
		}
		c.Set("user", context)
	}
}
