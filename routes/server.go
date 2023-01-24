package routes

import (
	"github.com/gin-gonic/gin"
	"go-identity-service/controllers"
	"go-identity-service/middlewares"
)

func SetupServer() *gin.Engine {
	router := initRouter()
	return router
}

func initRouter() *gin.Engine {
	router := gin.Default()
	router.POST("/singup", controllers.SingUp)
	router.POST("/signin", controllers.SingIn)
	api := router.Group("/api").Use(middlewares.Auth())
	{
		api.GET("/ping", controllers.Ping)
	}
	return router
}
