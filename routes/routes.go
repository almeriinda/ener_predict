package routes

import (
	"ener_predict/controllers"
	"ener_predict/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	// Rotas públicas
	router.POST("/login", controllers.LoginUser)
	router.POST("/register", controllers.RegisterUser)

	// Rotas protegidas
	protected := router.Group("/")
	protected.Use(middlewares.AuthMiddleware)

	protected.GET("/forecast/:userID", controllers.GetForecast)
	//testada
	protected.GET("/consumption/:userID", controllers.GetConsumption)
	protected.GET("/users", controllers.GetAllUsers)
	protected.POST("/consumption", controllers.AddConsumption)
	protected.GET("/user", controllers.GetUserInfo)
}
