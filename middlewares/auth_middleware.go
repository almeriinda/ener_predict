package middlewares

import (
	"net/http"
	"strings"
	"github.com/gin-gonic/gin"
	"ener_predict/services"
)

// AuthMiddleware verifica o token JWT presente no header Authorization
func AuthMiddleware(c *gin.Context) {
	// Obter o token do header Authorization
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token não fornecido"})
		c.Abort()
		return
	}

	// Extrair o token do header
	token := strings.TrimPrefix(authHeader, "Bearer ")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token mal formatado"})
		c.Abort()
		return
	}

	// Verificar o token
	_, err := services.ParseToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
		c.Abort()
		return
	}

	// Token válido, prossiga com a requisição
	c.Next()
}
