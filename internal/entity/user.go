package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserName string `bson:"username" json:"username"`
	Email string `bson:"email" json:"email"`
	Password string `bson:"password" json:"password"`
	Role string `bson:"role" json:"role"`
	CreatedAt time.Time `bson:"createdat" json:"createdat"`
	UpdatedAt time.Time `bson:"updatedat" json:"updatedat"`
}