package routes

import (
	"restaurant-service/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
    r.GET("/restaurants-in-city", controllers.ListRestaurantsInCity)
    r.GET("/food-options-near", controllers.ListFoodOptionsNearLocation)
}
