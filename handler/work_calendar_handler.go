package handler

import (
	"fmt"
	"gin-example/db"
	"gin-example/model"
	"gin-example/types"
	"gin-example/utils/api_utils"

	"github.com/gin-gonic/gin"
)

// 获取用户列表
func GetWorkCalendar(c *gin.Context) {
	workCalendars, err := model.GetAllWorkCalendar()
	fmt.Println(workCalendars) // log:[]
	if err != nil {
		api_utils.ErrorResponse(c, 500, err.Error())
		return
	}
	api_utils.SuccessResponse(c, workCalendars)
}

// 假设数据库中已经存在表a，其中b是我们需要检查的字段
func ModifyDayMessage(c *gin.Context) {
	var workCalendar types.WorkCalendar

	if err := c.ShouldBindJSON(&workCalendar); err != nil {
		// 如果绑定失败，返回错误信息
		api_utils.ErrorResponse(c, 500, err.Error())
		return
	}
	// // 使用 = 进行精确匹配
	// query := "SELECT COUNT(*) FROM a WHERE b = ?"
	// // 使用精确匹配查询 b 字段
	// rows, err := db.Query(query, searchString)

	rows, err := db.DB.Query("SELECT day FROM work_calendar WHERE day = ?", workCalendar.Day)

	if err != nil {
		CreateWorkCalendar(c, workCalendar)
		return
	}
	defer rows.Close()

	count := 0
	// 使用 rows.Next() 来遍历所有行，并计数
	for rows.Next() {
		count++
	}

	// 如果 count 大于 0，表示找到了匹配的记录
	if count > 0 {
		UpdateWorkCalendar(c, workCalendar)
		return
	}

	// 如果没有找到匹配的记录
	CreateWorkCalendar(c, workCalendar)
}

// 创建用户
func CreateWorkCalendar(c *gin.Context, workCalendar types.WorkCalendar) {

	model.CreateWorkCalendar(workCalendar)
	// 如果绑定成功，返回成功的 JSON 数据
	api_utils.SuccessResponse(c, workCalendar)

}

// 更新用户
func UpdateWorkCalendar(c *gin.Context, workCalendar types.WorkCalendar) {
	// var workCalendar types.WorkCalendar

	// if err := c.ShouldBindJSON(&workCalendar); err != nil {
	// 	// 如果绑定失败，返回错误信息
	// 	api_utils.ErrorResponse(c, 500, err.Error())
	// 	return
	// }

	if _, err := model.UpdateWorkCalendar(workCalendar); err != nil {
		api_utils.ErrorResponse(c, 500, err.Error())
		return
	}

	// 如果绑定成功，返回成功的 JSON 数据
	api_utils.SuccessResponse(c, nil)

}
