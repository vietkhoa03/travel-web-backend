package repository

import (
	"context"
	"fmt"
	"travel-web-backend/internal/entity"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	Login(ctx context.Context, user model.User) (model.User, error)
	SignUp(ctx context.Context, user model.User) (model.User, error)
}

type userRepository struct {
	db *mongo.Database
}


func NewUserRepository(db *mongo.Database) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Login(ctx context.Context, user model.User) (model.User, error) {
	collection := r.db.Collection("Users")

	if user.Email == "" || user.Password == "" {
		return model.User{}, fmt.Errorf("Email and paddword are required")
	}

	var foundUser model.User
	condition := bson.M{"email": user.Email, "password": user.Password}
	if err := collection.FindOne(ctx, condition).Decode(&foundUser); err != nil {
		return model.User{}, err
	}
	return foundUser, nil
}

func (r *userRepository) SignUp(ctx context.Context, user model.User) (model.User, error) {
	collection := r.db.Collection("Users")

	if user.UserName == "" || user.Password == "" {
		return model.User{}, fmt.Errorf("Username and paddword are required")
	}

	var foundExistUser model.User
	condition := bson.M{"email": user.Email}
	err := collection.FindOne(ctx, condition).Decode(&foundExistUser)
	if err == mongo.ErrNoDocuments {
		newUser := model.User{
			UserName : user.UserName,
			Email : user.Email,
			Password : user.Password,
			Role : user.Role,
			CreatedAt : user.CreatedAt,
			UpdatedAt : user.UpdatedAt,
		}
		 _, err := collection.InsertOne(ctx, newUser)
		if err != nil {
			return model.User{}, err
		}
		return newUser, nil
	} else if err != nil {
		return model.User{}, err
	}
	return model.User{},  fmt.Errorf("User already exist")
}
