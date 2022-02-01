package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/gwuhaolin/livego/configure"
)

var secret []byte

type AuthToken struct {
	Id        string `json:"id"`
	Ip        string `json:"ip"`
	UserAgent string `json:"userAgent"`
	jwt.StandardClaims
}

func getSecret() {
	configure.Config.GetString("secret")
}

func signToken(id string, ip string, agent string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &AuthToken{
		Id:        id,
		Ip:        ip,
		UserAgent: agent,
	})
	tokenString, err := token.SignedString(secret)
	return tokenString, err
}

func tokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}
