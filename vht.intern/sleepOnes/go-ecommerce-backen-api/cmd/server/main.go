package main

import (
	"go-ecommerce-backend-api/internal/controller"
	router "go-ecommerce-backend-api/internal/routers"
)

func main() {
	router := router.NewRouter()

	v1 := router.Group("/v1")
	{
		// UserController
		userController := controller.NewUserController()
		v1.GET("/ping/:name", userController.GetUserById)
	}
	router.Run(":8001") // listen and serve on 0.0.0.0:8080
}
