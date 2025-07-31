package middlewares

import (
	"fmt"
	"go-ecommerce-backend-api/pkg/response"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token != "valid-token" {
			response.ErrorResponse(c, response.ErrInvalidToken)
			c.Abort() // Stop the request from proceeding further
			fmt.Println("Invalid token provided")
			return
		}
		fmt.Println("Valid token provided")
		c.Next()
	}

}
