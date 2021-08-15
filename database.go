package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	dbCtx          context.Context
	dbClient       *mongo.Client
	dbVrc          *mongo.Database
	dbVrcUsers     *mongo.Collection
	dbVrcAvatars   *mongo.Collection
	dbVrcWorlds    *mongo.Collection
	dbVrcUploaders *mongo.Collection
)

func dbInit(clusterUrl string, username string, password string) bool {
	dbUri := fmt.Sprintf("mongodb+srv://%s:%s@%s?retryWrites=true&w=majority", username, password, clusterUrl)

	var err error
	dbClient, err = mongo.NewClient(options.Client().ApplyURI(dbUri))

	if err != nil {
		return false
	}

	dbCtx = context.Background()

	err = dbClient.Connect(dbCtx)
	if err != nil {
		return false
	}

	err = dbClient.Ping(dbCtx, readpref.Primary())
	if err != nil {
		return false
	}

	dbVrc = dbClient.Database("VRChat")
	dbVrcUsers = dbVrc.Collection("users")
	dbVrcAvatars = dbVrc.Collection("avatars")
	dbVrcWorlds = dbVrc.Collection("worlds")
	dbVrcUploaders = dbVrc.Collection("uploaders")

	return true
}

func dbClose() {
	dbClient.Disconnect(dbCtx)
}
