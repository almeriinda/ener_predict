package controllers

import (
	"ener_predict/config"
	"ener_predict/models"
	"ener_predict/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddConsumption(c *gin.Context) {
	var consumption models.Consumption
	if err := c.ShouldBindJSON(&consumption); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	db := config.GetDB()

	if _, err := services.CreateConsumption(db, consumption); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao adicionar consumo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Consumo adicionado com sucesso"})

}

func GetConsumption(c *gin.Context) {
	userID := c.Param("userID")
	db := config.GetDB()

	consumptionData, err := services.GetConsumptionByUserID(db, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar consumo"})
		return
	}

	c.JSON(http.StatusOK, consumptionData)
}
