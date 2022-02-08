package routes

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
	if registration.Username == "" || !IsValidPassword.MatchString(registration.Username) {
		c.AbortWithStatusJSON(400, gin.H{"error": "valid username is required"})
		return
	}
	query := bson.M{"usernameLower": strings.ToLower(registration.Username)}
	var user = &mongo.User{}
	mongo.UserCollection.FindOne(context.Background(), query).Decode(&user)
	if user.Username != "" {
		c.AbortWithStatusJSON(400, gin.H{"error": "username already exists"})
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registration.Password), bcrypt.DefaultCost)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": "internal server error"})
		return
	}
	mongo.UserCollection.InsertOne(c, bson.M{
		"username":      registration.Username,
		"usernameLower": strings.ToLower(registration.Username),
		"password":      string(hashedPassword),
		"email":         registration.Email,
		"emailLower":    strings.ToLower(registration.Email),
		"status":        "unverified",
	})
	c.JSON(200, gin.H{"error": "ok"})
}
