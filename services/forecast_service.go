package services

import (
	"bytes"
	"encoding/json"
	"os/exec"
	"time"
)

// ForecastRequest define a estrutura para a solicitação de previsão
type ForecastRequest struct {
	// Adapte os campos conforme necessário
	UserID   uint      `json:"user_id"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
}

// ForecastResponse define a estrutura para a resposta de previsão
type ForecastResponse struct {
	Date   time.Time `json:"date"`
	Value  float64   `json:"value"`
}

// RunForecast executa o script de previsão e retorna os resultados
func RunForecast(request ForecastRequest) ([]ForecastResponse, error) {
	// Converta a solicitação para JSON
	requestData, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	// Execute o script de previsão Python
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

	// Parse a resposta do script Python
	var forecastResponse []ForecastResponse
	err = json.Unmarshal(out.Bytes(), &forecastResponse)
	if err != nil {
		return nil, err
	}

	return forecastResponse, nil
}
