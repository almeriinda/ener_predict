package services

import (
	"ener_predict/config"
	"ener_predict/models"
	"errors"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	UserID uint   `json:"user_id"`
	Email  string `json:"email"`
	jwt.StandardClaims
}

func GenerateToken(user models.User) (string, error) {
	claims := Claims{
		UserID: user.ID,
		Email:  user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			Issuer:    "ener_predict",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(config.JWTSecret))
}

func ParseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.JWTSecret), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, errors.New("token inválido")
	}

	return claims, nil
}

func AuthenticateUser(email, password string) (string, error) {
	db := config.GetDB()

	user, err := models.GetUserByEmail(db, email)
	if err != nil {
		log.Println("erro aqui", err)
		return "", err
	}

	if err := user.CheckPassword(password); err != nil {
		return "", errors.New("senha incorreta")
	}

	token, err := GenerateToken(user)
	if err != nil {
		return "", err
	}

	return token, nil
}

func CreateUser(user models.User) error {
	if err := config.DB.Create(&user).Error; err != nil {
		log.Println("Erro ao criar usuário no banco de dados:", err)
		return err
	}
	return nil
}
