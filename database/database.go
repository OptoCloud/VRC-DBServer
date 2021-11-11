package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	internalCtx         context.Context
	internalClient      *mongo.Client
	internalDatabase    *mongo.Database
	CollectionConfig    *mongo.Collection
	CollectionUploaders *mongo.Collection
	CollectionUsers     *mongo.Collection
	CollectionAvatars   *mongo.Collection
	CollectionWorlds    *mongo.Collection
)

func Open(clusterUrl string, username string, password string) error {
	dbUri := fmt.Sprintf("mongodb+srv://%s:%s@%s?retryWrites=true&w=majority", username, password, clusterUrl)

	var err error

	log.Println("Creating database client")
	internalClient, err = mongo.NewClient(options.Client().ApplyURI(dbUri))
	if err != nil {
		return err
	}

	internalCtx = context.Background()

	log.Println("Connecting to MongoDB cluster...")
	err = internalClient.Connect(internalCtx)
	if err != nil {
		return err
	}

	log.Println("Verifying connection is up...")
	err = internalClient.Ping(internalCtx, readpref.Primary())
	if err != nil {
		return err
	}

	log.Println("Getting MongoDB collections...")
	internalDatabase = internalClient.Database("VRChat")
	CollectionConfig = internalDatabase.Collection("config")
	CollectionUploaders = internalDatabase.Collection("uploaders")
	CollectionUsers = internalDatabase.Collection("users")
	CollectionAvatars = internalDatabase.Collection("avatars")
	CollectionWorlds = internalDatabase.Collection("worlds")

	err = initConfigWatcher()
	if err != nil {
		return err
	}

	log.Println("Database setup done!")

	return nil
}
func Close() {
	internalClient.Disconnect(internalCtx)
}
