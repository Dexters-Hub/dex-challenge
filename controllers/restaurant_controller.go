package controllers

import (
	"math"
	"net/http"
	"restaurant-service/models"
	"restaurant-service/utils"
	"strconv"
	"github.com/gin-gonic/gin"
)

func ListRestaurantsInCity(c *gin.Context) {
    cityCode := c.Query("city_code")
    if cityCode == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "City code is required"})
        return
    }

    restaurants, err := utils.ParseCSV("data/restaurants.csv")
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load restaurant data"})
        return
    }

    cityRestaurants := filterRestaurantsByCity(restaurants, cityCode)

    c.JSON(http.StatusOK, cityRestaurants)
}

func filterRestaurantsByCity(restaurants []models.Restaurant, cityCode string) []models.Restaurant {
    var cityRestaurants []models.Restaurant
    for _, r := range restaurants {
        if r.CityCode == cityCode {
            cityRestaurants = append(cityRestaurants, r)
        }
    }
    return cityRestaurants
}

func ListFoodOptionsNearLocation(c *gin.Context) {
    latitudeStr := c.Query("latitude")
    longitudeStr := c.Query("longitude")

    if latitudeStr == "" || longitudeStr == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Latitude and longitude are required"})
        return
    }

    latitude, err := strconv.ParseFloat(latitudeStr, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid latitude format"})
        return
    }

    longitude, err := strconv.ParseFloat(longitudeStr, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid longitude format"})
        return
    }

    restaurants, err := utils.ParseCSV("data/restaurants.csv")
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load restaurant data"})
        return
    }

    var nearbyRestaurants []models.Restaurant
    for _, r := range restaurants {
        dist := distance(latitude, longitude, r.Latitude, r.Longitude)
        if dist <= 5.0 {
            nearbyRestaurants = append(nearbyRestaurants, r)
        }
    }

    c.JSON(http.StatusOK, nearbyRestaurants)
}

func distance(lat1, lon1, lat2, lon2 float64) float64 {

	return math.Abs(lat1-lat2) + math.Abs(lon1-lon2)
}
