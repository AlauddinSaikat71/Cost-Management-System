package db

import (
	"costmanagement/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open("sqlite3", "test.db")

	if err != nil {
		panic("Failed to connect to database")
	}
	database.AutoMigrate(&models.User{})
	database.AutoMigrate(&models.Attachment{})
	database.AutoMigrate(&models.Cost{})
	database.AutoMigrate(&models.CostAttachment{})
	database.AutoMigrate(&models.Payment{})

	DB = database
}
