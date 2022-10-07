package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zekhoi/learn-golang/request"
	"github.com/zekhoi/learn-golang/service"
)

type shortenHandler struct {
	service service.ShortenService
}

func NewShortenHandler(service service.ShortenService) *shortenHandler {
	return &shortenHandler{service}
}

func (h *shortenHandler) CreateShorten(c *gin.Context) {
	var shortenRequest request.CreateShortenRequest

	err := c.ShouldBindJSON(&shortenRequest)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return

	}

	shorten, err := h.service.CreateShorten(shortenRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	c.JSON(http.StatusCreated, gin.H{
		"code":   http.StatusCreated,
		"status": "Created",
		"data":   shorten,
	})
}

func (h *shortenHandler) GetShortenByCode(c *gin.Context) {
	var request request.GetShortenRequest

	err := c.ShouldBindJSON(&request)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	shorten, err := h.service.GetShortenByCode(request)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	c.JSON(http.StatusOK, gin.H{
		"code":   http.StatusOK,
		"status": "OK",
		"data":   shorten,
	})
}

func (h *shortenHandler) GetAllShorten(c *gin.Context) {
	shortens, err := h.service.GetShortens()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	c.JSON(http.StatusOK, gin.H{
		"code":   http.StatusOK,
		"status": "OK",
		"data":   shortens,
	})
}
