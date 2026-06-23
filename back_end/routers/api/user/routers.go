package user

import (
	"alto_server/common/middlewares"

	"github.com/gin-gonic/gin"
)

// RegisterRouter 注册路由
//nokia-alto/user
func RegisterRouter(router *gin.RouterGroup) {

	user := router.Group("/user")
	{
		//使用Token管理中间件, 保存用户信息
		user.Use(middlewares.JWTAuthz())
		// 用户路由
		user.POST("/register", registerHandler)
		user.POST("/role", roleHandler)
	}
}
