package router

import (
	"github.com/gin-gonic/gin"
)

// SetupRouter 用于设置所有路由
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// 路由设置
	// r.GET("/user", handler.GetUsers)
	r = UserRouter(r)

	// 可以在这里添加更多路由
	// r.POST("/users", handler.CreateUser)
	// r.PUT("/users/:id", handler.UpdateUser)

	return r
}
