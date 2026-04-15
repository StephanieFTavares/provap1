package database

import (
	"log"

	"github.com/StephanieFTavares/provap1/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	db, err := gorm.Open(sqlite.Open("condlink.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Erro ao conectar banco")
	}

	db.AutoMigrate(&models.User{})

	DB = db
}
