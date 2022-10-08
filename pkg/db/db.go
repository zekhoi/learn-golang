package db

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/zekhoi/learn-golang/pkg/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error when loading .env file")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))

	db, err = ConnectDB(dsn)
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&entity.Shorten{})
}

func ConnectDB(dsn string) (*gorm.DB, error) {

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Database not connected")
	}
	fmt.Printf("Database connected\n")

	return db, nil
}

func GetDB() *gorm.DB {
	return db
}
