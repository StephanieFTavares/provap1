package main

import (
	"./database"
	"./routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()

	r := gin.Default()

	r.Use(cors.Default())

	routes.SetupRoutes(r)

	r.Run(":8080")
}
