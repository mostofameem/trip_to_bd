package mongodb

import (
	"post-service/config"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
)

type LocationTypeRepo struct {
	schema     string
	collection *mongo.Collection
}

const locationCollectionName = "locations"

var locationTypeRepo *LocationTypeRepo

var locationCntOnce = sync.Once{}

func NewLocationTypeRepo(cnf *config.MongoDBConfig) *LocationTypeRepo {
	locationCntOnce.Do(func() {
		mongodb := NewMongoDB(cnf)

		locationTypeRepo = &LocationTypeRepo{
			schema:     "posts",
			collection: mongodb.Database.Collection(locationCollectionName),
		}
	})
	return locationTypeRepo
}
