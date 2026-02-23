package main

import (
	"net/http"

	"github.com/jhoancamilorayomejia/uptime-monitor/types"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

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
