package models

import (
	"time"

	"gorm.io/gorm"
)

type Consumption struct {
	ID              uint      `gorm:"primaryKey" json:"id"`
	UserID          uint      `gorm:"not null" json:"user_id"`
	ConsumptionDate time.Time `json:"consumption_date"`
	Amount          float64   `gorm:"not null" json:"amount"`
	Month           int       `json:"month"`
	Day             int       `json:"day"`
	Temperature     float64   `json:"temperature"`
	UsageHours      float64   `json:"usage_hours"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	User            User      `gorm:"foreignKey:UserID"`
}

func GetConsumptionByUserID(db *gorm.DB, userID uint) ([]Consumption, error) {
	var consumptions []Consumption
	err := db.Where("user_id = ?", userID).Find(&consumptions).Error
	return consumptions, err
}
