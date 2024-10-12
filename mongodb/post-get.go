package mongodb

import (
	"context"
	"log/slog"

	"go.mongodb.org/mongo-driver/bson"
)

func (repo *LocationTypeRepo) GetLocations(ctx context.Context) (*[]Location, error) {
	data, err := repo.collection.Find(ctx, bson.D{})
	if err != nil {
		slog.Error("Failed to get locations")
		return nil, err
	}
	defer data.Close(ctx)

	var locations []Location
	if err = data.All(ctx, &locations); err != nil {
		slog.Error("Failed to struct locations")
		return nil, err
	}

	slog.Info("Data retrieved successfully")
	return &locations, nil
}
