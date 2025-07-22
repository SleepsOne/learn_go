package controller

import (
	"go-ecommerce-backend-api/internal/service"
	"go-ecommerce-backend-api/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: service.NewUserService(),
	}
}

func (uc *UserController) GetUserById(c *gin.Context) {
	uidStr := c.Query("uid")
	userName := c.Param("name")
	uid, err := strconv.Atoi(uidStr)
	if err != nil {
		response.ErrorResponse(c, response.ErrCodeParamInvalid)
		return
	}

	response.SuccessResponse(c, response.ErrCodeSuccess, uc.userService.GetInfoUser(uid, userName))

}
