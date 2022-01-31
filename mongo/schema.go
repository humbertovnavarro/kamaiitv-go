package mongo

import (
	"go.mongodb.org/mongo-driver/mongo"
)

var UserCollection = &mongo.Collection{}

type User struct {
	Username      string `bson:"username,omitempty"`
	UsernameLower string `bson:"usernameLower,omitempty"`
	Password      string `bson:"password,omitempty"`
	Email         string `bson:"email,omitempty"`
	EmailLower    string `bson:"emailLower,omitempty"`
	Deleted       bool   `bson:"deleted,omitempty"`
	Status        string `bson:"status,omitempty"`
}

type Message struct {
	ID       string `bson:"_id,omitempty"`
	Username string `bson:"username,omitempty"`
	Content  string `bson:"content,omitempty"`
	Time     int64  `bson:"time,omitempty"`
	Deleted  bool   `bson:"deleted,omitempty"`
}

func schema(db *mongo.Database) {
	UserCollection = db.Collection("Users")
}
