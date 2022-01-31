package api

import (
	"context"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gwuhaolin/livego/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

type UserRegistration struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func RegisterUser(c *gin.Context) {
	registration := UserRegistration{}
	if err := c.BindJSON(&registration); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "json syntax error"})
		return
	}
	if registration.Email == "" {
		c.AbortWithStatusJSON(400, gin.H{"error": "valid email is required"})
		return
	}
	if registration.Password == "" || !IsValidPassword.MatchString(registration.Password) {
		c.AbortWithStatusJSON(400, gin.H{"error": "valid password is required"})
		return
	}
	if registration.Username == "" || !IsValidPassword.MatchString(registration.Password) {
		c.AbortWithStatusJSON(400, gin.H{"error": "valid username is required"})
		return
	}
	query := bson.M{"usernameLower": strings.ToLower(registration.Username)}
	email := bson.M{"emailLower": strings.ToLower(registration.Email)}
	var user bson.M
	mongo.UserCollection.FindOne(context.Background(), bson.M{
		"$or": bson.A{query, email},
	}).Decode(&user)
	if user != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "username or email already exists"})
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registration.Password), bcrypt.DefaultCost)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": "internal server error"})
		return
	}
	mongo.UserCollection.InsertOne(c, &mongo.User{
		Username:      registration.Username,
		UsernameLower: strings.ToLower(registration.Username),
		Password:      string(hashedPassword),
		Email:         registration.Email,
	})
	c.JSON(200, gin.H{"error": "ok"})
}
