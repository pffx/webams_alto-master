package auth

import (
	"alto_server/common/middlewares"

	"github.com/gin-gonic/gin"
)

// RegisterRouter 注册路由
//nokia-alto/auth
func RegisterRouter(router *gin.RouterGroup) {
	auth := router.Group("/auth")
	{
		// 用户路由
		auth.POST("/login", LoginHandler)
		auth.POST("/logout", LogoutHandler)
		auth.GET("/token", middlewares.JWTAuthz(), UpdateTokenHandler)
	}
}
