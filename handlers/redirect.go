package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"link-shortener/storage"
)

type RedirectRequest struct {
	ShortCode string `json:"short_code" binding:"required"`
}

func RedirectURL(db *storage.SQLiteDB) gin.HandlerFunc {
	return func(c *gin.Context) {
		shortCode := c.Param("shortCode")
		longURL, err := db.GetLink(shortCode)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Link not found"})
			return
		}
		c.Redirect(http.StatusFound, longURL)
	}
}
