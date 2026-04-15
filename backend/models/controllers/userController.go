package controllers

import (
	"net/http"

	"condlink/database"
	"condlink/models"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)

	user.Active = false
	user.Role = "morador"

	database.DB.Create(&user)

	c.JSON(http.StatusOK, user)
}

func Login(c *gin.Context) {
	var input models.User
	var user models.User

	c.BindJSON(&input)

	database.DB.Where("email = ? AND password = ?", input.Email, input.Password).First(&user)

	if user.ID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciais inválidas"})
		return
	}

	if !user.Active {
		c.JSON(http.StatusForbidden, gin.H{"error": "Usuário não aprovado"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func GetUsers(c *gin.Context) {
	var users []models.User
	database.DB.Find(&users)
	c.JSON(http.StatusOK, users)
}

func ApproveUser(c *gin.Context) {
	id := c.Param("id")

	var user models.User
	database.DB.First(&user, id)

	user.Active = true
	database.DB.Save(&user)

	c.JSON(http.StatusOK, user)
}
