package config

import (
	"fmt"
	"log"
	"os"

	"github.com/aliemreipek/go-task-master/pkg/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	// install the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/Istanbul",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Could not connect to database! Is Docker up and running?")
	}

	// Automatic migration (Creates Tables)
	// Converts the Task struct into a database table.
	err = db.AutoMigrate(&models.Task{})
	if err != nil {
		log.Fatal("Migration error:", err)
	}

	DB = db
	fmt.Println("🚀 Database connection is successful and tables created!")
}
