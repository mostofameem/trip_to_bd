package mongodb

import (
	"context"
	"log/slog"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type Comment struct {
	Userid     int
	UserName   string
	Content    string
	Created_at time.Time
	Updated_at time.Time
}

func (repo *LocationTypeRepo) AddComment(ctx context.Context, locationId int, cmnt Comment) error {
	postKey := GetKey(locationId)

	filter := bson.M{"_id": postKey}

	update := bson.M{
		"$push": bson.M{"comments": cmnt},
	}

	_, err := repo.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		slog.Error("Error adding comment")
		return err
	}

	slog.Info("Comment added successfully")
	return nil
}
