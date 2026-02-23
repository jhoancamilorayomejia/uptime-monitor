package main

import (
	"net/http"

	"github.com/jhoancamilorayomejia/uptime-monitor/backend/types"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	r.GET("/api/status", func(c *gin.Context) {

		services := []types.ServiceStatus{
			types.CheckService("Google", "https://www.google.com"),
			types.CheckService("GitHub", "https://api.github.com"),
			types.CheckService("Servicio Falso", "https://no-existe-123.com"),
		}

		c.JSON(http.StatusOK, services)
	})

	r.Run(":8080")
}
