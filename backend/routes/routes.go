package routes

import (
	"github.com/StephanieFTavares/provap1/controllers"
	"github.com/StephanieFTavares/provap1/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api")

	api.POST("/register", controllers.Register)
	api.POST("/login", controllers.Login)

	protected := api.Group("/")
	protected.Use(middleware.AuthMiddleware())

	protected.GET("/users", controllers.GetUsers)
	protected.PUT("/approve/:id", controllers.ApproveUser)
}
