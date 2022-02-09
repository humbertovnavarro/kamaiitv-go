package api

import (
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/gwuhaolin/livego/lib"
	"github.com/gwuhaolin/livego/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func LoginUser(c *gin.Context) {
	registration := UserRegistration{}
	if err := c.BindJSON(&registration); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "json syntax error"})
		return
	}
	if registration.Email == "" {
		c.AbortWithStatusJSON(400, gin.H{"error": "valid email is required"})
		return
	}
	if registration.Password == "" || !lib.IsValidPassword.MatchString(registration.Password) {
		c.AbortWithStatusJSON(400, gin.H{"error": "valid password is required"})
		return
	}
	if registration.Username == "" || !lib.IsValidPassword.MatchString(registration.Username) {
		c.AbortWithStatusJSON(400, gin.H{"error": "valid username is required"})
		return
	}
	query := bson.M{"usernameLower": strings.ToLower(registration.Username)}
	email := bson.M{"emailLower": strings.ToLower(registration.Email)}
	var user = &mongo.User{}
	found := mongo.UserCollection.FindOne(c, bson.M{
		"$or": bson.A{query, email},
	})
	if found.Err() != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "username or email not found"})
		return
	}
	found.Decode(&user)
	if user.Status == "unverified" {
		c.AbortWithStatusJSON(401, gin.H{"error": "unverified"})
		return
	}
	userPassword := []byte(user.Password)
	dbPassword := []byte(registration.Password)
	if err := bcrypt.CompareHashAndPassword(userPassword, dbPassword); err != nil {
		c.AbortWithStatusJSON(401, gin.H{"error": "unauthorized"})
		return
	}
	tkStrct := jwt.StandardClaims{
		Id:       user.ID,
		Subject:  user.Username,
		Audience: c.Request.UserAgent(),
	}
	token, err := lib.SignToken(tkStrct)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": "internal server error"})
		return
	}
	c.SetCookie("token", token, int(time.Hour*24*14), "", "", false, true)
	c.JSON(200, gin.H{"token": token})
}
