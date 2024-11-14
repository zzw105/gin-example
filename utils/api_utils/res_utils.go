package api_utils

import (
	"gin-example/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SuccessResponse 返回成功的统一格式
func SuccessResponse(c *gin.Context, data interface{}, options ...*types.Options) (int, types.BaseResponse) {
	// 默认值
	resp := types.BaseResponse{
		Code:    200, // 默认 code 200
		Data:    data,
		Message: "Success", // 默认 message "Success"
	}

	if len(options) > 0 {
		resp.Code = options[0].Code
		resp.Message = options[0].Message
	}

	c.JSON(http.StatusOK, resp)

	return http.StatusOK, resp
}

// ErrorResponse 返回错误的统一格式
func ErrorResponse(c *gin.Context, code int, message string) {
	c.JSON(http.StatusOK, types.BaseResponse{
		Code:    code,
		Data:    nil,
		Message: message,
	})
}
