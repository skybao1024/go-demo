package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func JSON(c *gin.Context, httpStatus int, code int, message string, data interface{}) {
	c.JSON(httpStatus, Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}

func Success(c *gin.Context, data interface{}) {
	JSON(c, http.StatusOK, 200, "success", data)
}

func SuccessWithoutData(c *gin.Context) {
	c.Status(http.StatusNoContent)
}

func Fail(c *gin.Context, code int, message string) {
	JSON(c, http.StatusOK, code, message, nil)
}

func BadRequest(c *gin.Context, message string) {
	Fail(c, http.StatusBadRequest, message)
}

func Unauthorized(c *gin.Context, message string) {
	Fail(c, http.StatusUnauthorized, message)
}

func InternalServerError(c *gin.Context, message string) {
	Fail(c, http.StatusInternalServerError, message)
}

// 可以根据需要添加更多的辅助函数
