package controllers

import (
	"net/http"
	"strings"

	"ener_predict/config"
	"ener_predict/models"
	"ener_predict/services"

	"github.com/gin-gonic/gin"
)

func GetUserInfo(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token não fornecido"})
		return
	}

	token := strings.TrimPrefix(authHeader, "Bearer ")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token mal formatado"})
		return
	}

	claims, err := services.ParseToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
		return
	}

	db := config.GetDB()
	user, err := models.GetUserByID(db, claims.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar usuário"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user_id": user.ID,
		"email":   user.Email,
	})
}
