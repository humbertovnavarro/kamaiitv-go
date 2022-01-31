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

var MessageCollection = &mongo.Collection{}

type Message struct {
	ID       string `bson:"_id,omitempty"`
	Username string `bson:"username,omitempty"`
	Content  string `bson:"content,omitempty"`
	Time     int64  `bson:"time,omitempty"`
	Deleted  bool   `bson:"deleted,omitempty"`
}

var NodeCollection = &mongo.Collection{}

type Node struct {
	ID           string `bson:"_id,omitempty"`
	Hostname     string `bson:"hostname,omitempty"`
	LoadBalancer string `bson:"loadBalancer,omitempty"`
}

func schema(db *mongo.Database) {
	MessageCollection = db.Collection("Messages")
	UserCollection = db.Collection("Users")
	NodeCollection = db.Collection("Nodes")
}
