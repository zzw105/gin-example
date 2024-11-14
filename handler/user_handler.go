package handler

import (
	"gin-example/model"
	"gin-example/types"
	"gin-example/utils/api_utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 获取用户列表
func GetUsers(c *gin.Context) {
	users, err := model.GetAllUsers()
	if err != nil {
		api_utils.ErrorResponse(c, 500, err.Error())
		return
	}
	api_utils.SuccessResponse(c, users)
}

// 创建用户
func CreateUser(c *gin.Context) {
	var user types.User

	if err := c.ShouldBindJSON(&user); err != nil {
		// 如果绑定失败，返回错误信息
		api_utils.ErrorResponse(c, 500, err.Error())
		return
	}
	model.CreateUser(user)
	// 如果绑定成功，返回成功的 JSON 数据
	api_utils.SuccessResponse(c, user)

}

// 更新用户
func UpdateUser(c *gin.Context) {
	idParam := c.Param("id")
	var user types.User

	// 将 id 从字符串转换为整数
	id, err := strconv.Atoi(idParam)
	if err != nil {
		// 如果转换失败，返回 400 错误，告诉用户 id 必须是一个有效的数字
		api_utils.ErrorResponse(c, 500, "Invalid user ID, must be a number")
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		// 如果绑定失败，返回错误信息
		api_utils.ErrorResponse(c, 500, err.Error())
		return
	}

	if _, err := model.UpdateUser(id, user); err != nil {
		api_utils.ErrorResponse(c, 500, err.Error())
		return
	}

	// 如果绑定成功，返回成功的 JSON 数据
	api_utils.SuccessResponse(c, nil)

}
