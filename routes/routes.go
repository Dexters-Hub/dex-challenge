package routes

import (
	"restaurant-service/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
    r.GET("/restaurants-in-city", controllers.ListRestaurantsInCity)
    r.GET("/food-options-near", controllers.ListFoodOptionsNearLocation)
    r.GET("/restaurants-sorted-by-rating", controllers.ListRestaurantsSortedByRating)
    r.GET("/filter-by-table-booking", controllers.FilterRestaurantsByTableBooking)
    r.GET("/filter-by-online-delivery", controllers.FilterRestaurantsByOnlineDelivery)
    r.GET("/filter-by-cuisines", controllers.FilterRestaurantsByCuisines)
}