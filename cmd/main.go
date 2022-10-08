package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zekhoi/learn-golang/pkg/db"
	"github.com/zekhoi/learn-golang/pkg/handler"
	"github.com/zekhoi/learn-golang/pkg/repository"
	"github.com/zekhoi/learn-golang/pkg/service"
)

func main() {
	db := db.GetDB()

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
