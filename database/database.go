package database

import (
	"email-template-generator/log"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Connection struct {
	DB *gorm.DB
}

var (
	DB_HOST     string
	DB_USER     string
	DB_URL      string
	DB_NAME     string
	DB_PASSWORD string
	DB_PORT     string
	DSN         string
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	DB_HOST = os.Getenv("DB_HOST")
	DB_USER = os.Getenv("DB_USER")
	DB_URL = os.Getenv("DB_URL")
	DB_NAME = os.Getenv("DB_NAME")
	DB_PASSWORD = os.Getenv("DB_PASSWORD")
	DB_PORT = os.Getenv("DB_PORT")
	DSN = fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s sslmode=disable TimeZone=Asia/Shanghai", DB_HOST, DB_USER, DB_PASSWORD, DB_PORT, DB_NAME)
}

func New() *Connection {
	db, err := gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if err != nil {
		panic("Failed Connect Database")
	}

	log.Info("Success Connect Database")
	return &Connection{
		DB: db,
	}
}
