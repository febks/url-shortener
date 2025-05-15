package handlers

import (
	"net/http"
	"os"
	"url-shortener/models"
	"url-shortener/utils"

	"github.com/gin-gonic/gin"
)

var store = make(map[string]string)

func ShortenURL(c *gin.Context) {
	var req models.URLRequest
	errorResponse := models.ShortenResponse{
		Success: false,
		Message: "Invalid request",
		Data:    models.DataPayload{},
	}

	if err := c.BindJSON(&req); err != nil || req.URL == "" {
		c.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	code := utils.GenerateCode(6)
	store[code] = req.URL
	response := models.ShortenResponse{
		Success: true,
		Message: "URL shortened successfully",
		Data: models.DataPayload{
			OriginalURL: req.URL,
			ShortCode:   code,
			ResultURL:   os.Getenv("REQUEST_HOST") + os.Getenv("ENDPOINT") + code,
		},
	}

	c.JSON(http.StatusOK, response)
}

func ResolveURL(c *gin.Context) {
	code := c.Param("code")
	longURL, exists := store[code]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
		return
	}
	c.Redirect(http.StatusMovedPermanently, longURL)
}
