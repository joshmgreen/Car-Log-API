package db

import (
	"fmt"
	"log"
	"os"

	"github.com/joshmgreen/Car-Log-API/internal/vehicles/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, dbname, port,
	)

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	DB = database

	// Auto-migrate Vehicle table
	if err := DB.AutoMigrate(&model.Vehicle{}); err != nil {
		log.Fatalf("failed to auto-migrate: %v", err)
	}

	// Optional: Seed one vehicle if table is empty
	var count int64
	DB.Model(&model.Vehicle{}).Count(&count)
	if count == 0 {
		DB.Create(&model.Vehicle{
			Year:    2025,
			Make:    "Honda",
			Model:   "Civic Si",
			Mileage: 2499,
		})
		log.Println("Seeded initial vehicle")
	}
}
