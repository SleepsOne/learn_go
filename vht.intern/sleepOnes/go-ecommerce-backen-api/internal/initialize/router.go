package initialize

import (
	"go-ecommerce-backend-api/internal/controller"
	"go-ecommerce-backend-api/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.Use(middlewares.AuthMiddleware())

	v1 := router.Group("/v1")
	{
		// UserController
		userController := controller.NewUserController()
		v1.GET("/ping/:name", userController.GetUserById)
	}
	return router
}
