package mongodb

import (
	"post-service/config"

	"go.mongodb.org/mongo-driver/mongo"
)

type LocationTypeRepo struct {
	schema     string
	collection *mongo.Collection
}

const locationCollectionName = "locations"

func NewLocationTypeRepo(cnf *config.MongoDBConfig) *LocationTypeRepo {
	mongodb := NewMongoDB(cnf)
	return &LocationTypeRepo{
		schema:     "posts",
		collection: mongodb.Database.Collection(locationCollectionName),
	}
}
