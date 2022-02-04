package mongo

import (
	"context"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/gwuhaolin/livego/configure"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var Ctx context.Context
var DB *mongo.Database = nil

func Connect() {
	uri := configure.Config.GetString("mongo_uri")
	fmt.Print("connecting to mongodb: ", uri, "\n")
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	Ctx = ctx
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	log.Info("connected to mongodb")
	DB = client.Database("kamaiitv")
	db, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	log.Info("mongodb databases: ", db)
	schema(DB)
}
func Disconnect() {
	log.Info("disconnected from mongodb")
	DB.Client().Disconnect(context.Background())
}
