package database

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/robertogsf/POC/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	//var DB *gorm.DB
	var err error
	p := Config("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)
	if err != nil {
		panic("failed to parse database port")
	}
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		Config("DB_HOST"), port, Config("DB_USER"), Config("DB_PASSWORD"),
		Config("DB_NAME"))
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connection Opened to Database")
	DB.AutoMigrate(&models.Product{})
	DB.AutoMigrate(&models.User{})
	fmt.Println("Database Migrated")
}

func Config(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Print("Error loading .env file")
	}
	return os.Getenv(key)
}
