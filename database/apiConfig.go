package database

import (
	"log"
	"sync"
	"vrcdb/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	apiConfigMtx sync.RWMutex
	apiConfig    models.ApiConfig
)

func initConfigWatcher() error {
	log.Println("Initializing Config listener...")

	stream, err := CollectionConfig.Watch(internalCtx, mongo.Pipeline{
		bson.D{{
			Key: "$match",
			Value: bson.D{{
				Key: "$or",
				Value: []bson.D{
					{{Key: "operationType", Value: "insert"}},
					{{Key: "operationType", Value: "replace"}},
					{{Key: "operationType", Value: "update"}},
				},
			}},
		}},
	}, options.ChangeStream().SetFullDocument(options.UpdateLookup))
	if err != nil {
		return err
	}

	result := CollectionConfig.FindOne(internalCtx, bson.M{})
	err = result.Err()
	if err != nil {
		return err
	}

	err = result.Decode(&apiConfig)
	if err != nil {
		return err
	}

	go func() {
		defer stream.Close(internalCtx)

		var event struct {
			FullDocument *models.ApiConfig `bson:"fullDocument"`
		}
		event.FullDocument = &apiConfig

		for stream.Next(internalCtx) {
			apiConfigMtx.Lock()
			err := stream.Decode(&event)
			if err != nil {
				apiConfigMtx.Unlock()
				log.Println(err.Error())
				return
			}
			apiConfigMtx.Unlock()
		}
	}()

	return nil
}

func GetApiConfig() models.ApiConfig {
	apiConfigMtx.RLock()
	defer apiConfigMtx.RUnlock()

	return apiConfig
}
