package mongo

import (
	"go.mongodb.org/mongo-driver/mongo"
)

var UserCollection = &mongo.Collection{}

type User struct {
	ID            string `bson:"_id"`
	Username      string `bson:"username"`
	UsernameLower string `bson:"usernameLower"`
	Password      string `bson:"password"`
	Email         string `bson:"email"`
	EmailLower    string `bson:"emailLower"`
	Deleted       bool   `bson:"deleted"`
	Status        string `bson:"status"`
}

var MessageCollection = &mongo.Collection{}

type Message struct {
	ToRoom    string `bson:"toRoom"`
	FromId    string `bson:"fromId"`
	FromName  string `bson:"fromName"`
	Content   string `bson:"content"`
	CreatedAt int64  `bson:"createdAt"`
	Deleted   bool   `bson:"deleted"`
}

var NodeCollection = &mongo.Collection{}

type Node struct {
	ID           string `bson:"_id"`
	Hostname     string `bson:"hostname"`
	LoadBalancer string `bson:"loadBalancer"`
}

var FollowerCollection = &mongo.Collection{}

type Follower struct {
	ID       string `bson:"_id,omitempty"`
	Follower string `bson:"follower,omitempty"`
	Channel  string `bson:"channel,omitempty"`
}

func schema(db *mongo.Database) {
	MessageCollection = db.Collection("Messages")
	UserCollection = db.Collection("Users")
	NodeCollection = db.Collection("Nodes")
	FollowerCollection = db.Collection("Followers")
}
