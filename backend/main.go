package main

import (
	"condlink/database"
	"condlink/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectMongo()

	r := gin.Default()

	r.Use(cors.Default())

	routes.SetupRoutes(r)

	r.Run(":8080")
}
