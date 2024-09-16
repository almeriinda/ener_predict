package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"ener_predict/config"
	"ener_predict/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Erro ao carregar .env, usando variáveis do ambiente")
	}

	config.ConnectDB()

	router := gin.Default()

	routes.SetupRoutes(router)

	server := &http.Server{
		Handler:      router,
		Addr:         ":" + os.Getenv("PORT"),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	log.Printf("Servidor iniciado em http://localhost:%s", os.Getenv("PORT"))
	log.Fatal(server.ListenAndServe())
}
