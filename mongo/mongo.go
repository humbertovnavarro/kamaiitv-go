package mongo

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/gwuhaolin/livego/configure"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var Ctx context.Context
var DB *mongo.Database = nil

func Connect() {
	uri := configure.Config.GetString("mongo_uri")
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
	schema(DB)
}
func Disconnect() {
	log.Info("disconnected from mongodb")
	DB.Client().Disconnect(context.Background())
}
