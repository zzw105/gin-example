package router

import (
	"gin-example/handler"

	"github.com/gin-gonic/gin"
)

func WorkCalendarRouter(r *gin.Engine) *gin.Engine {

	api := r.Group("/work-calendar")
	{
		api.GET("/", handler.GetWorkCalendar)
		api.POST("/", handler.ModifyDayMessage)
		// api.PUT("/", handler.UpdateWorkCalendar)
	}
	// 路由设置
	// r.GET("/user", handler.GetUsers)
	// r.POST("/user", handler.CreateUser)
	// r.PUT("/user/:id", handler.UpdateUser)

	// 可以在这里添加更多路由
	// r.POST("/users", handler.CreateUser)
	// r.PUT("/users/:id", handler.UpdateUser)

	return r
}
