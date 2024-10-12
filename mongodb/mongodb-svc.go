package mongodb

import (
	"log/slog"
	"os"
	"post-service/config"

	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDB struct {
	Client   *mongo.Client
	Database *mongo.Database
}

func NewMongoDB(mongoCnf *config.MongoDBConfig) *MongoDB {
	mongoDb := Connet(mongoCnf)

	return &MongoDB{
		Client:   mongoDb.Client,
		Database: mongoDb.Database,
	}
}

func Connet(mongoCnf *config.MongoDBConfig) *MongoDB {

	mongoDB := connect(mongoCnf)
	if mongoDB == nil {
		slog.Error("MongoDb Connection is nil")
		os.Exit(1)
	}
	slog.Info("Connected to Mongo Database")
	return mongoDB
}
