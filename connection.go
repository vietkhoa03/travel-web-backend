package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No found .env file")
	}

	uri := os.Getenv("MONGODB_URI")

	if uri == "" {
    	log.Fatal("MONGODB_URI is not set")
	}

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))

	if err != nil {
		panic(err)
	}

	defer func() {
		if err = client.Disconnect(context.Background()); err != nil {
			panic(err)
		}
	}()

	result, err := client.ListDatabaseNames(context.Background(), bson.D{})

	if err != nil {
		log.Fatal(err)
	}

	for _, db := range result {
		fmt.Println(db)
	}
}
