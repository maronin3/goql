package config

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Client
var Database []string

func DBConnection() {
	for {
		err := MongoDB()
		if err != nil {
			log.Printf("database connection failed, will retry again in 10 Second : %v\n", err.Error())
			timer := time.NewTimer(10 * time.Second)
			<-timer.C
		} else {
			break
		}
	}
}

func MongoDB() error {
	uri := os.Getenv("MONGO_URI")
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		return err
	}

	databalist, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		return err
	}

	DB = client
	Database = databalist

	return nil
}
