package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"url_shortner/database"
	"url_shortner/models"
	"url_shortner/utils"

	"github.com/gin-gonic/gin"
)

func ShortenURL(ctx *gin.Context) {
	var body models.Body

	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	shortCode := utils.GeneratorShortCode()
	// Verificar si la URL tiene el esquema (http:// o https://), si no, agregar http://
	if !strings.HasPrefix(body.URL, "http://") && !strings.HasPrefix(body.URL, "https://") {
		body.URL = "http://" + body.URL
	}

	url := models.URL{
		OriginalURL: body.URL,
		ShortURL:    shortCode,
	}

	fmt.Println("URL original:", body.URL)
	if err := database.DB.Create(&url).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	response := map[string]string{
		"short_url": fmt.Sprintf("http://localhost:8080/r/%s", shortCode),
	}
	log.Println("URL acortada:", response["short_url"])

	ctx.JSON(http.StatusOK, gin.H{"Content-Type": "application/json", "data": response})
}

func RedirectHandler(ctx *gin.Context) {
	// Obtén el parámetro dinámico de la URL
	shorturl := ctx.Param("shorturl")
	log.Println("ShortURL recibido:", shorturl)

	var url models.URL

	// Busca en la base de datos el short_url
	if err := database.DB.Where("short_url = ?", shorturl).First(&url).Error; err != nil {
		log.Println("Error al buscar en la base de datos:", err.Error())
		ctx.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
		return
	}

	// Redirige a la URL original
	log.Println("Redirigiendo a:", url.OriginalURL)
	ctx.Redirect(http.StatusFound, url.OriginalURL)
}
