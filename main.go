package main

import (
	"fmt"
	"log"
	"os"
	connectDB "travel-web-backend/database"
	controller "travel-web-backend/internal/controller/view"
	repository "travel-web-backend/internal/repository"
	"travel-web-backend/internal/service"
	route "travel-web-backend/router"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	frontendURL := os.Getenv("FE_URL")
	port := os.Getenv("PORT")

	travelDB := connectDB.ConnectTravelDB()
	locationRepo := repository.NewLocationRepository(travelDB)
	userRepo := repository.NewUserRepository(travelDB)

	// Init services
	locationService := service.NewLocationService(locationRepo)
	userService := service.NewUserService(userRepo)

	// Init controllers
	locationController := controller.NewLocationController(locationService)
	userController := controller.NewUserController(userService)

	// Router
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", frontendURL)
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	r = route.SetupRouter(r, locationController, userController)
	// Chạy server
	r.Run(":" + port)

	fmt.Println("Database name:", travelDB.Name())
}
