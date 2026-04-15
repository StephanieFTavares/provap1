package main

import (
	"github.com/StephanieFTavares/provap1/database"
	"github.com/StephanieFTavares/provap1/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()

	r := gin.Default()
	routes.SetupRoutes(r)

	r.Run(":8080")
}
