package main

import (
	"example.com/go-rest-api-gin/database"
	"example.com/go-rest-api-gin/router"
	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()

	route := gin.Default()
	api := route.Group("/api")

	{
		router.UserRoutes(api.Group("/v1"))
		router.PhotoRoutes(api.Group("/v1"))
	}

	route.Run(":2173")
}
