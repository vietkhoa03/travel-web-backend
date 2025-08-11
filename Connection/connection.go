package connectDB

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectTravelDB() *mongo.Database {
	if err := godotenv.Load(); err != nil {
		log.Println("No found .env file")
	}

	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("MONGODB_URI is not set")
	}

	// Context với timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Kết nối MongoDB
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatalf("Connect error: %v", err)
	}

	// Ping để kiểm tra kết nối
	if err := client.Ping(ctx, nil); err != nil {
		log.Fatalf("Ping error: %v", err)
	}

	// Lấy DB Travel
	db := client.Database("Travel")
	if db != nil {
		fmt.Println("connect success")
	}

	return db
}
