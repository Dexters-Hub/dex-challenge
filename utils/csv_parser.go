package utils

import (
	"encoding/csv"
	"log"
	"os"
	"restaurant-service/models"
	"strconv"
	"strings"
)

func ParseCSV(filePath string) ([]models.Restaurant, error) {
    file, err := os.Open(filePath)
    if err != nil {
        log.Fatal(err)
        return nil, err
    }
    defer file.Close()

    reader := csv.NewReader(file)
    records, err := reader.ReadAll()
    if err != nil {
        log.Fatal(err)
        return nil, err
    }

    var restaurants []models.Restaurant
    for _, record := range records {
        restaurant := models.Restaurant{}
        restaurant.ID, _ = strconv.Atoi(record[0])
        restaurant.Name = record[1]
        restaurant.CountryCode = record[2]
        restaurant.CityCode = record[3]
        restaurant.Address = record[4]
        restaurant.Locality = record[5]
        restaurant.LocalityVerbose = record[6]
        restaurant.Longitude, _ = strconv.ParseFloat(record[7], 64)
        restaurant.Latitude, _ = strconv.ParseFloat(record[8], 64)
        restaurant.Cuisines = record[9]
        restaurant.AverageCostForTwo, _ = strconv.Atoi(record[10])
        restaurant.Currency = record[11]
        restaurant.HasTableBooking = strings.ToLower(record[12]) == "yes"
        restaurant.HasOnlineDelivery = strings.ToLower(record[13]) == "yes"
        restaurant.IsDeliveringNow = strings.ToLower(record[14]) == "yes"
        restaurant.SwitchToOrderMenu = strings.ToLower(record[15]) == "yes"
        restaurant.PriceRange, _ = strconv.Atoi(record[16])
        restaurant.AggregateRating, _ = strconv.ParseFloat(record[17], 64)
        restaurant.RatingColor = record[18]
        restaurant.RatingText = record[19]
        restaurant.Votes, _ = strconv.Atoi(record[20])

        restaurants = append(restaurants, restaurant)
    }

    return restaurants, nil
}
