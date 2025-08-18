package route

import (
	controller "travel-web-backend/internal/controller/view"

	"github.com/gin-gonic/gin"
)

func SetupRouter(
	locationController *controller.LocationController,
	userController *controller.UserController,
) *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		locationRoutes := v1.Group("/locations")
		{
			locationRoutes.GET("/search", locationController.SearchLocationsByName)
			locationRoutes.GET("/:id", locationController.GetLocationByID)
		}

		userRoutes := v1.Group("/users")
		{
			userRoutes.POST("/signup", userController.SignUp)
			userRoutes.POST("/login", userController.Login)
		}
	}

	return r
}
