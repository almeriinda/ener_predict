package services

import (
	"ener_predict/models"
	"time"

	"gorm.io/gorm"
)

func CreateConsumption(db *gorm.DB, consumption models.Consumption) (models.Consumption, error) {
	consumption.CreatedAt = time.Now()

	result := db.Create(&consumption)
	if result.Error != nil {
		return models.Consumption{}, result.Error
	}
	return consumption, nil
}

func GetConsumptions(db *gorm.DB, userID uint) ([]models.Consumption, error) {
	var consumptions []models.Consumption
	result := db.Where("user_id = ?", userID).Find(&consumptions)
	if result.Error != nil {
		return nil, result.Error
	}
	return consumptions, nil
}

func GetConsumptionByID(db *gorm.DB, id uint) (models.Consumption, error) {
	var consumption models.Consumption
	result := db.First(&consumption, id)
	if result.Error != nil {
		return models.Consumption{}, result.Error
	}
	return consumption, nil
}

func UpdateConsumption(db *gorm.DB, id uint, updatedConsumption models.Consumption) (models.Consumption, error) {
	var consumption models.Consumption
	result := db.First(&consumption, id)
	if result.Error != nil {
		return models.Consumption{}, result.Error
	}

	consumption.Amount = updatedConsumption.Amount
	consumption.ConsumptionDate = updatedConsumption.ConsumptionDate

	result = db.Save(&consumption)
	if result.Error != nil {
		return models.Consumption{}, result.Error
	}
	return consumption, nil
}

func DeleteConsumption(db *gorm.DB, id uint) error {
	var consumption models.Consumption
	result := db.Delete(&consumption, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetConsumptionByUserID(db *gorm.DB, userID string) ([]models.Consumption, error) {
	var consumption []models.Consumption
	if err := db.Where("user_id = ?", userID).Find(&consumption).Error; err != nil {
		return nil, err
	}
	return consumption, nil
}
