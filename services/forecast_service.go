package services

import (
	"bytes"
	"encoding/json"
	"os/exec"
	"time"
)

type ForecastRequest struct {
	UserID    uint      `json:"user_id"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
}

type ForecastResponse struct {
	Date  time.Time `json:"date"`
	Value float64   `json:"value"`
}

func RunForecast(request ForecastRequest) ([]ForecastResponse, error) {
	requestData, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	cmd := exec.Command("python3", "scripts/run_forecast.py")
	cmd.Stdin = bytes.NewReader(requestData)

	var out bytes.Buffer
	cmd.Stdout = &out
	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	err = cmd.Run()
	if err != nil {
		return nil, err
	}

	var forecastResponse []ForecastResponse
	err = json.Unmarshal(out.Bytes(), &forecastResponse)
	if err != nil {
		return nil, err
	}

	return forecastResponse, nil
}
