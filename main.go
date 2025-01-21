package main

import (
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"link-shortener/handlers"
	"link-shortener/storage"
)

func main() {
	db, err := storage.InitDB("links.db")
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	r := gin.Default()

	r.POST("/shorten", handlers.ShortenURL(db))
	r.GET("/:shortCode", handlers.RedirectURL(db))

	log.Println("Server is running on localhost:8080")
	r.Run(":8080")

}
