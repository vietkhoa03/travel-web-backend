package model
import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)
type Location struct {
    ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
    Name        string             `bson:"name" json:"name"`
    Description string             `bson:"description" json:"description"`
    Image       string             `bson:"images" json:"image"`
    Notes       string             `bson:"notes" json:"notes"`
    Highlights  string             `bson:"highlights" json:"highlights"`
    CreatedAt   time.Time          `bson:"createdat" json:"createdat"`
    UpdatedAt   time.Time          `bson:"updatedat" json:"updatedat"`
}