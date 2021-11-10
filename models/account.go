package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type ClientAccountDbRecord struct {
	Id        primitive.ObjectID `json:"id" bson:"_id"`
	Nickname  string             `json:"nickname" bson:"nickname"`
	DiscordId uint64             `json:"discord_id" bson:"discord_id"`
	Rank      uint32             `json:"rank" bson:"rank"`
}
