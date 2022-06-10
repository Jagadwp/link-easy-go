package db

import (
	"fmt"
	"log"
	"os"

	model "github.com/Jagadwp/link-easy-go/internal/entity"
	"github.com/Jagadwp/link-easy-go/internal/shared/config"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var database *gorm.DB
var e error

func DatabaseInit() {

	err := godotenv.Load(config.EnvPath)
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=require TimeZone=Asia/Jakarta", host, user, password, dbName, port)
	database, e = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if e != nil {
		panic(e)
	}

	initMigrate()
}

func DB() *gorm.DB {
	return database
}

func initMigrate() {
	database.AutoMigrate(&model.User{})
	// database.AutoMigrate(&_model.UrlData{})
}