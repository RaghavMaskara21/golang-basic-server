package mongodb

import (
	"context"
	"hayday/server/config"
	"hayday/server/internal/logger"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func Connect() *mongo.Client {
	log := logger.Log.WithFields(map[string]interface{}{
		"EVENT": "MONGO_CONNECTION",
	})
	log.Infof("Initialized the mongoDb connection")
	serverAPI := options.ServerAPI(options.ServerAPIVersion1).SetStrict(true).SetDeprecationErrors(true)
	opts := options.Client().ApplyURI(config.EnvValues.MONGO_DB_URL).SetServerAPIOptions(serverAPI).SetMinPoolSize(2).SetMaxPoolSize(100)
	client, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		log.Fatalf("failed to establish the mongo connection : ERROR : %s", err)
	}

	IsMongoDbHealthy(log, client)
	log.Infof("Connected to mongoDb successfully")
	return client
}

func IsMongoDbHealthy(log *logger.LoggerEvent, client *mongo.Client) bool {
	err := client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatalf("failed to establish the mongo connection : ERROR : %s", err)
	}
	return true
}
