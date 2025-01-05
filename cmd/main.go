package main

import (
	"log"
	"url_shortner/database"
	"url_shortner/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	database.DBConnect()

	database.Migrate()

	r.POST("/shorten", handlers.ShortenURL)
	r.GET("/r/:shorturl", handlers.RedirectHandler)

	log.Println("NAZI -- Server running on port 8080 -- NAZI")
	r.Run(":8080")
}
