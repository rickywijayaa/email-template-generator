package database

import (
	"email-template-generator/entity"
	"email-template-generator/log"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Connection struct {
	DB *gorm.DB
}

type Model struct {
	Model interface{}
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

	migrate(db)
	seed(db)

	log.Info("Success Connect Database")

	return &Connection{
		DB: db,
	}
}

func modelList() []Model {
	return []Model{
		{Model: entity.User{}},
		{Model: entity.Email{}},
		{Model: entity.SystemConfig{}},
	}
}

func migrate(db *gorm.DB) {
	for _, item := range modelList() {
		err := db.Debug().AutoMigrate(item.Model)
		if err != nil {
			log.Fatal(fmt.Sprintf("Error Migrate - %s", err.Error()))
			return
		}
	}

	log.Info("Migrate Successfully")
}

func seed(db *gorm.DB) {
	var user entity.User
	_ = db.Where("email = ?", "admin@gmail.com").Find(&user).Error

	if user.ID == 0 {
		password, _ := bcrypt.GenerateFromPassword([]byte("12345Qwe!"), bcrypt.MinCost)
		db.Create(&entity.User{
			Name:     "Admin",
			Email:    "admin@gmail.com",
			Password: string(password),
			Token:    "",
		})
	}

	var systemConfig entity.SystemConfig
	_ = db.Where("code = ? AND value = ?", "Language", "Indonesia").Find(&systemConfig)
	if systemConfig.ID == 0 {
		db.Create(&entity.SystemConfig{
			Code:  "Language",
			Value: "Indonesia",
		})
	}

	_ = db.Where("code = ? AND value = ?", "Language", "English").Find(&systemConfig)
	if systemConfig.ID == 0 {
		db.Create(&entity.SystemConfig{
			Code:  "Language",
			Value: "English",
		})
	}

	_ = db.Where("code = ? AND value = ?", "TemplateEmail", "Job Application").Find(&systemConfig)
	if systemConfig.ID == 0 {
		db.Create(&entity.SystemConfig{
			Code:  "TemplateEmail",
			Value: "Job Application",
		})
	}

	log.Info("Success seeding")
}
