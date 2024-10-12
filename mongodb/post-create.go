package mongodb

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/exp/slog"
)

type Author struct {
	UserID   int    `json:"userid" bson:"userid"`
	Username string `json:"username" bson:"username"`
}

type Location struct {
	ID           int       `json:"id" bson:"id"`
	Title        string    `json:"title" bson:"title"`
	Descriptions string    `json:"content" bson:"content"`
	BestTime     string    `json:"best_time" bson:"best_time"`
	PictureUrl   string    `json:"picture_url" bson:"picture_url"`
	Rating       float32   `json:"rating" bson:"rating"`
	Voted        int       `json:"voted" bson:"voted"`
	Author       Author    `json:"author" bson:"author"`
	ReviewIDs    []int     `json:"review_ids" bson:"review_ids"`
	CreatedAt    time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" bson:"updated_at"`
}

func (repo *LocationTypeRepo) AddLocation(location *Location) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	location.ReviewIDs = append(location.ReviewIDs, location.ID)
	mongoDoc, err := BsonM(location)
	if err != nil {
		slog.Error("Error converting struct to BSON", err)
		return err
	}

	key := GetKey(location.ID)
	mongoDoc["_id"] = key

	_, err = repo.collection.InsertOne(ctx, mongoDoc)
	if err != nil {
		slog.Error("Error inserting document", err)
		return err
	}

	return nil
}

func GetKey(id int) string {
	return fmt.Sprintf("posts:%d", id)
}

func BsonM(data interface{}) (bson.M, error) {
	result := bson.M{}

	val := reflect.Indirect(reflect.ValueOf(data))
	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldName := typ.Field(i).Tag.Get("bson")

		if fieldName == "" {
			fieldName = typ.Field(i).Name
		}

		if val.Field(i).CanInterface() {
			result[fieldName] = field.Interface()
		}
	}

	return result, nil
}
