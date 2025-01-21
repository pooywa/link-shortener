package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"link-shortener/storage"
	"link-shortener/utils"
)

type ShortenRequest struct {
	LongURL string `json:"long_url" binding:"required,url"`
}

func ShortenURL(db *storage.SQLiteDB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req ShortenRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		shortCode := utils.GenerateShortCode(req.LongURL)
		err := db.SaveLink(shortCode, req.LongURL, time.Now())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save link"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"shortCode": shortCode})
	}
}
