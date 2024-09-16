package controllers

import (
	"net/http"
	"os/exec"

	"github.com/gin-gonic/gin"
)

func GetForecast(c *gin.Context) {
	month := c.Query("month")
	day := c.Query("day")
	temperature := c.Query("temperature")
	usageHours := c.Query("usage_hours")

	if month == "" || day == "" || temperature == "" || usageHours == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parâmetros inválidos"})
		return
	}

	cmd := exec.Command("python3", "scripts/run_forecast.py", month, day, temperature, usageHours)

	output, err := cmd.CombinedOutput()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao gerar previsão", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"forecast": string(output)})
}
