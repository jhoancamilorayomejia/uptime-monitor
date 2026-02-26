package types

import (
	"database/sql"
	"net/http"
	"time"
)

// 🔹 Modelo BD
type Service struct {
	ID   int
	Name string
	URL  string
}

// 🔹 Modelo respuesta
type ServiceStatus struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	URL          string `json:"url"`
	Status       string `json:"status"`
	ResponseTime int64  `json:"responseTime"`
}

// 🔹 Obtener servicios desde PostgreSQL
func GetServices(db *sql.DB) ([]Service, error) {
	rows, err := db.Query("SELECT id, name, url FROM services")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var services []Service

	for rows.Next() {
		var s Service
		if err := rows.Scan(&s.ID, &s.Name, &s.URL); err != nil {
			return nil, err
		}
		services = append(services, s)
	}

	return services, nil
}

// 🔹 Check en tiempo real
func CheckService(service Service) ServiceStatus {
	start := time.Now()

	client := http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Get(service.URL)
	elapsed := time.Since(start).Milliseconds()

	if err != nil || resp.StatusCode >= 400 {
		return ServiceStatus{
			ID:           service.ID,
			Name:         service.Name,
			URL:          service.URL,
			Status:       "DOWN",
			ResponseTime: 0,
		}
	}

	return ServiceStatus{
		ID:           service.ID,
		Name:         service.Name,
		URL:          service.URL,
		Status:       "UP",
		ResponseTime: elapsed,
	}
}
