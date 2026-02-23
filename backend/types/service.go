package types

import (
	"net/http"
	"time"
)

type ServiceStatus struct {
	Name         string `json:"name"`
	URL          string `json:"url"`
	Status       string `json:"status"`
	ResponseTime int64  `json:"responseTime"`
}

func CheckService(name, url string) ServiceStatus {
	start := time.Now()

	client := http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Get(url)
	elapsed := time.Since(start).Milliseconds()

	if err != nil || resp.StatusCode >= 400 {
		return ServiceStatus{
			Name:         name,
			URL:          url,
			Status:       "DOWN",
			ResponseTime: 0,
		}
	}

	return ServiceStatus{
		Name:         name,
		URL:          url,
		Status:       "UP",
		ResponseTime: elapsed,
	}
}
