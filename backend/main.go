package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// Puerto por variable de entorno
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Router Gin
	router := gin.Default()

	// Endpoint básico de health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"message": "Uptime Monitor API running",
		})
	})

	// Levantar servidor
	router.Run(":" + port)
}
