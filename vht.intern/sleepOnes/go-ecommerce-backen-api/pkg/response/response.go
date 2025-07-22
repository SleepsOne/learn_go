package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseData struct {
	Code    int         `json:"code"`           //status code
	Message string      `json:"message"`        // Response message
	Data    interface{} `json:"data,omitempty"` // Use `omitempty` to omit the field if it's nil
}

//sucess response

func SuccessResponse(c *gin.Context, code int, data interface{}) {
	c.JSON(http.StatusOK, ResponseData{
		Code:    code,
		Message: msg[code],
		Data:    data,
	})

}

func ErrorResponse(c *gin.Context, code int) {
	c.JSON(http.StatusOK, ResponseData{
		Code:    code,
		Message: msg[code],
		Data:    nil,
	})
}
