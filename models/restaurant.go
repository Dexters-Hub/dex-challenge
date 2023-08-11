package models

type Restaurant struct {
    ID                 int
    Name               string
    CountryCode        string
    CityCode           string
    Address            string
    Locality           string
    LocalityVerbose    string
    Longitude          float64
    Latitude           float64
    Cuisines           string
    AverageCostForTwo  int
    Currency           string
    HasTableBooking    bool
    HasOnlineDelivery  bool
    IsDeliveringNow    bool
    SwitchToOrderMenu  bool
    PriceRange         int
    AggregateRating    float64
    RatingColor        string
    RatingText         string
    Votes              int
}
