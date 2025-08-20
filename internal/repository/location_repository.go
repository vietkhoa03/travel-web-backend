package repository

import (
	"context"
	"fmt"
	"time"
	model "travel-web-backend/internal/entity"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type LocationRepository interface {
	GetAllLocation(ctx context.Context) ([]model.Location, error)
	CreateLocation(ctx context.Context, location model.Location) (model.Location, error)
	UpdateLocation(ctx context.Context, id string, location model.Location) (model.Location, error)
	SearchLocationsByName(ctx context.Context, query string, page, limit int) ([]model.Location, error)
	GetLocationByID(ctx context.Context, id string) (model.Location, error)
	DeleteLocation(ctx context.Context, id string) error
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

func (r *locationRepository) GetAllLocation(ctx context.Context) ([]model.Location, error) {
	collection := r.db.Collection("Locations")

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	var location []model.Location
	if err := cursor.All(ctx, &location); err != nil {
		return nil, err
	}

	return location, nil
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

func (r *locationRepository) CreateLocation(ctx context.Context, location model.Location) (model.Location, error) {
	collection := r.db.Collection("Locations")

	location.ID = primitive.NewObjectID()
	location.CreatedAt = time.Now()
	location.UpdatedAt = location.CreatedAt

	_, err := collection.InsertOne(ctx, location)
	if err != nil {
		return model.Location{}, err
	}
	return location, nil
}

func (r *locationRepository) UpdateLocation(ctx context.Context, id string, location model.Location) (model.Location, error) {
	collection := r.db.Collection("Locations")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return model.Location{}, fmt.Errorf("invalid ObjectID: %w", err)
	}
	location.UpdatedAt = time.Now()
	update := bson.M{
		"$set": bson.M{
			"name":        location.Name,
			"description": location.Description,
			"images":      location.Images,
			"notes":       location.Notes,
			"highlights":  location.Highlights,
			"updatedat":   location.UpdatedAt,
		},
	}
	_, err = collection.UpdateOne(ctx, bson.M{"_id": objID}, update)
	if err != nil {
		return model.Location{}, err
	}
	return r.GetLocationByID(ctx, id)
}

func (r *locationRepository) DeleteLocation(ctx context.Context, id string) error {
	colection := r.db.Collection("Locations")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("invalid ObjectID: %w", err)
	}

	_, err = colection.DeleteOne(ctx, bson.M{"_id": objID})
	return err
}
