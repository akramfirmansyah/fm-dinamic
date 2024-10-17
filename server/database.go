package main

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

	dsn := "host=" + os.Getenv("DB_HOST") + " user=" + os.Getenv("BD_USER") + " password=" + os.Getenv("DB_PASSWORD") + " dbname=" + os.Getenv("DB_NAME") + " port=" + os.Getenv("DB_PORT") + " sslmode=disable TimeZone=Asia/Shanghai"

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed connect to Database!")
	}
}

func Migrate() {
	DB.AutoMigrate(&models.Player{}, &models.Attribut{})

	fmt.Println("Success migrate!")
}
