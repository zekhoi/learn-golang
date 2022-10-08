package handler

import (
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/zekhoi/learn-golang/pkg/request"
	"github.com/zekhoi/learn-golang/pkg/service"
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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = url.ParseRequestURI(shortenRequest.OriginalUrl)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shorten, err := h.service.CreateShorten(shortenRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"code":   http.StatusCreated,
		"status": http.StatusText(http.StatusCreated),
		"data":   shorten,
	})
}

func (h *shortenHandler) GetShortenByCode(c *gin.Context) {
	code := c.Param("code")

	shorten, err := h.service.GetShortenByCode(code)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":   http.StatusOK,
		"status": http.StatusText(http.StatusOK),
		"data":   shorten,
	})
}

func (h *shortenHandler) GetAllShorten(c *gin.Context) {
	shortens, err := h.service.GetShortens()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":   http.StatusOK,
		"status": http.StatusText(http.StatusOK),
		"data":   shortens,
	})
}
