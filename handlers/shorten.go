package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShortenURL(db *models.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req shortenRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

	}

	c.JSON(http.StatusOK, gin.H{"shortCode": shortCode})
}
