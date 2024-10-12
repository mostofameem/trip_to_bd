package mongodb

import (
	"context"
	"fmt"
	"log/slog"
	"post-service/config"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func connect(db *config.MongoDBConfig) *MongoDB {
	ctx, cancle := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancle()

	mongodbSource := fmt.Sprintf(
		"mongodb://%s:%s@%s:%d",
		db.User,
		db.Pass,
		db.Host,
		db.Port,
	)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongodbSource))
	if err != nil {
		slog.Info("Mongodb connection Failed")
		return nil
	}

	database := client.Database(db.Name)

	return &MongoDB{
		Client:   client,
		Database: database,
	}
}
