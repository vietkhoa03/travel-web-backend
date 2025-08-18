package repository

import (
	"context"
	"fmt"
	"travel-web-backend/internal/entity"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type LocationRepository interface {
	SearchLocationsByName(ctx context.Context, query string, page, limit int) ([]model.Location, error)
	GetLocationByID(ctx context.Context, id string) (model.Location, error)
}

// struct implement interface
type locationRepository struct {
	db *mongo.Database
}

// constructor trả về interface
func NewLocationRepository(db *mongo.Database) LocationRepository {
	return &locationRepository{db: db}
}

// implement method
func (r *locationRepository) SearchLocationsByName(ctx context.Context, query string, page, limit int) ([]model.Location, error) {
	collection := r.db.Collection("Locations")

	skip := int64((page - 1) * limit)
	limit64 := int64(limit)
	pattern := ".*" + query + ".*"
	filter := bson.M{
		"name": bson.M{
			"$regex":   pattern,
			"$options": "i",
		},
	}

	findOptions := options.Find().
		SetSkip(skip).
		SetLimit(limit64).
		SetSort(bson.D{{"createdat", -1}})

	fmt.Println("Mongo filter:", filter)

	cursor, err := collection.Find(ctx, filter, findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var locations []model.Location
	if err := cursor.All(ctx, &locations); err != nil {
		return nil, err
	}

	return locations, nil
}

func (r *locationRepository) GetLocationByID(ctx context.Context, id string) (model.Location, error) {
    collection := r.db.Collection("Locations")

    objID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return model.Location{}, fmt.Errorf("invalid ObjectID: %w", err)
    }

    var location model.Location
    if err := collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&location); err != nil {
        return model.Location{}, err
    }

    return location, nil
}



