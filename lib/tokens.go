package lib

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/gwuhaolin/livego/configure"
)

var secret []byte

func GetSecret() {
	secretString := configure.Config.GetString("secret")
	if len(secretString) < 16 {
		panic("secret is empty or too short")
	}
	secret = []byte(secretString)
}

func SignToken(token jwt.StandardClaims) (string, error) {
	token.ExpiresAt = time.Now().Add(time.Hour * 24 * 7).Unix()
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, token)
	return claims.SignedString(secret)
}

func DecodeToken(tokenString string) (*jwt.StandardClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	claims := token.Claims.(*jwt.StandardClaims)
	return claims, nil
}

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := c.Cookie("token")
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": "unauthorized"})
			return
		}
		claims, err := DecodeToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": "unauthorized"})
			return
		}
		if claims.ExpiresAt < time.Now().Unix() {
			c.AbortWithStatusJSON(401, gin.H{"error": "expired"})
			return
		}
		c.Set("id", claims.Id)
	}
}
