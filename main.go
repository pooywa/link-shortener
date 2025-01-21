package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/shorten")
	r.GET("/:shortCode")

	log.Println("Server is running on localhost:8080")
	r.Run(":8080")

}
