package main

import (
	"fmt"
	connectDB "travel-web-backend/database"
	controller "travel-web-backend/internal/controller/view"
	repository "travel-web-backend/internal/repository"
	"travel-web-backend/internal/service"
	route "travel-web-backend/router"

	"github.com/gin-gonic/gin"
)

func main() {
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
	r = route.SetupRouter(locationController, userController)
	// Chạy server
	r.Run(":8080")

	// Ví dụ: in ra tên DB để kiểm tra
	fmt.Println("Database name:", travelDB.Name())
}
