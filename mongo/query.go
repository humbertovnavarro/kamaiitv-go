package mongo

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Common queries
func UpsertNodeInfo(hostname string) (err error) {
	ctx, c := context.WithTimeout(context.Background(), 5*time.Second)
	defer c()
	search := bson.M{
		"hostname": hostname,
	}
	update := bson.M{
		"$set": bson.M{
			"hostname": hostname,
		},
	}
	_, err = NodeCollection.UpdateOne(ctx, search, update, options.Update().SetUpsert(true))
	if err != nil {
		log.Fatal(err)
	}
	return
}
