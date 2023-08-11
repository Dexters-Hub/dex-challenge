package main

import (
	"restaurant-service/middleware"
	"restaurant-service/routes"

	"github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    r.Use(middleware.CorsMiddleware())

    routes.SetupRoutes(r)

    r.Run(":8080")
}
