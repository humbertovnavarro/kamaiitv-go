package routes

import (
	"fmt"
	"net/http"

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
func verifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secret, nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func extractToken(c *gin.Context) (*jwt.Token, error) {
	tokenString := c.GetHeader("Authorization")
	token, err := verifyToken(tokenString)
	token.Claims["id"]
	return token, err
}

func tokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		_, err := verifyToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}
		c.Next()
	}
}
