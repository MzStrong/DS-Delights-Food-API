package database

import (
	"github.com/MzStrong/DS-Delights-Food-API/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("Database.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	DB.AutoMigrate(&models.Menu{}, &models.Disease{}, &models.MenuDisease{})
}
