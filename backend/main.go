package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jhoancamilorayomejia/uptime-monitor/backend/db"
	"github.com/jhoancamilorayomejia/uptime-monitor/backend/types"
)

func main() {
	// 🔹 Conexión DB
	database, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	r := gin.Default()

	// 🔹 CORS simple
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

	// 🔹 Endpoint principal
	r.GET("/api/status", func(c *gin.Context) {

		services, err := types.GetServices(database)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		var result []types.ServiceStatus

		for _, service := range services {
			result = append(result, types.CheckService(service))
		}

		c.JSON(http.StatusOK, result)
	})

	r.Run(":8080")
}
