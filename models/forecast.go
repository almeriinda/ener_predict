package models

import (
	"time"
	"gorm.io/gorm"
)

// Forecast representa a previsão de consumo de energia para um usuário
type Forecast struct {
	ID             	uint      `gorm:"primaryKey" json:"id"`
	UserID         	uint      `gorm:"not null" json:"user_id"`
	ForecastDate   	time.Time `json:"forecast_date"`
	PredictedAmount float64   `gorm:"not null" json:"predicted_amount"`
	CreatedAt     	time.Time `json:"created_at"`
	UpdatedAt      	time.Time `json:"updated_at"`
	User           	User      `gorm:"foreignKey:UserID"`
}

// CreateForecast cria uma nova previsão de consumo de energia
func CreateForecast(db *gorm.DB, forecast *Forecast) error {
	return db.Create(forecast).Error
}

// GetForecastByUserID busca previsões pelo ID do usuário
func GetForecastByUserID(db *gorm.DB, userID uint) ([]Forecast, error) {
	var forecasts []Forecast
	err := db.Where("user_id = ?", userID).Find(&forecasts).Error
	return forecasts, err
}
