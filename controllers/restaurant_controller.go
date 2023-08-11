package controllers

import (
	"math"
	"net/http"
	"restaurant-service/models"
	"restaurant-service/utils"
	"sort"
	"strconv"
	"strings"

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

func ListRestaurantsSortedByRating(c *gin.Context) {
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
    sortRestaurantsByRating(cityRestaurants)

    c.JSON(http.StatusOK, cityRestaurants)
}

func sortRestaurantsByRating(restaurants []models.Restaurant) {
    sort.Slice(restaurants, func(i, j int) bool {
        return restaurants[i].AggregateRating > restaurants[j].AggregateRating
    })
}

func FilterRestaurantsByTableBooking(c *gin.Context) {
    cityCode := c.Query("city_code")
    hasTableBooking := c.Query("has_table_booking") == "true"

    restaurants, err := utils.ParseCSV("data/restaurants.csv")
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load restaurant data"})
        return
    }

    cityRestaurants := filterRestaurantsByCity(restaurants, cityCode)
    filteredRestaurants := filterRestaurantsByTableBooking(cityRestaurants, hasTableBooking)

    c.JSON(http.StatusOK, filteredRestaurants)
}

func filterRestaurantsByTableBooking(restaurants []models.Restaurant, hasTableBooking bool) []models.Restaurant {
    var filteredRestaurants []models.Restaurant
    for _, r := range restaurants {
        if r.HasTableBooking == hasTableBooking {
            filteredRestaurants = append(filteredRestaurants, r)
        }
    }
    return filteredRestaurants
}


func FilterRestaurantsByOnlineDelivery(c *gin.Context) {
    cityCode := c.Query("city_code")
    hasOnlineDelivery := c.Query("has_online_delivery") == "yes"

    restaurants, err := utils.ParseCSV("data/restaurants.csv")
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load restaurant data"})
        return
    }

    cityRestaurants := filterRestaurantsByCity(restaurants, cityCode)
    filteredRestaurants := filterRestaurantsByOnlineDelivery(cityRestaurants, hasOnlineDelivery)

    c.JSON(http.StatusOK, filteredRestaurants)
}

func filterRestaurantsByOnlineDelivery(restaurants []models.Restaurant, hasOnlineDelivery bool) []models.Restaurant {
    var filteredRestaurants []models.Restaurant
    for _, r := range restaurants {
        if r.HasOnlineDelivery == hasOnlineDelivery {
            filteredRestaurants = append(filteredRestaurants, r)
        }
    }
    return filteredRestaurants
}

func FilterRestaurantsByCuisines(c *gin.Context) {
    cityCode := c.Query("city_code")
    cuisines := c.Query("cuisines")

    restaurants, err := utils.ParseCSV("data/restaurants.csv")
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load restaurant data"})
        return
    }

    cityRestaurants := filterRestaurantsByCity(restaurants, cityCode)
    filteredRestaurants := filterRestaurantsByCuisines(cityRestaurants, cuisines)

    c.JSON(http.StatusOK, filteredRestaurants)
}

func filterRestaurantsByCuisines(restaurants []models.Restaurant, cuisines string) []models.Restaurant {
    var filteredRestaurants []models.Restaurant
    for _, r := range restaurants {
        if strings.Contains(strings.ToLower(r.Cuisines), strings.ToLower(cuisines)) {
            filteredRestaurants = append(filteredRestaurants, r)
        }
    }
    return filteredRestaurants
}