package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShortenURL(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "URL shortened"})
}
