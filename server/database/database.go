package database

import (
	"akramfirmansyah/fm-dinamic-api/models"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	var err error

	dsn := os.Getenv("DB_URL")

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed connect to Database!")
	}
}

func Migrate() {
	DB.AutoMigrate(&models.Player{}, &models.Attribut{})

	fmt.Println("Success migrate!")
}
