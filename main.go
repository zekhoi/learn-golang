package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/zekhoi/learn-golang/handler"
	"github.com/zekhoi/learn-golang/repository"
	"github.com/zekhoi/learn-golang/service"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error when loading .env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", dbHost, dbUser, dbPassword, dbName, dbPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Database not connected")
	}

	shortenRepository := repository.NewShortenRepository(db)
	shortenService := service.NewShortenService(shortenRepository)
	shortenHandler := handler.NewShortenHandler(shortenService)

	router := gin.Default()

	v1 := router.Group("/api/v1")

	v1.GET("/shorten", shortenHandler.GetAllShorten)
	v1.GET("/shorten/:code", shortenHandler.GetShortenByCode)
	v1.POST("/shorten", shortenHandler.CreateShorten)

	router.Run()
}
