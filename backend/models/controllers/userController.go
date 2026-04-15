package controllers

import (
	"context"
	"net/http"

	"condlink/database"
	"condlink/models"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var user models.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Erro nos dados"})
		return
	}

	user.Active = false
	user.Role = "morador"

	collection := database.DB.Collection("users")

	_, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao salvar"})
		return
	}

	c.JSON(http.StatusOK, user)
}
