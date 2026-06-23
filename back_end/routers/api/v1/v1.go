package v1

import (
	"alto_server/routers/api/v1/alarm"
	"alto_server/routers/api/v1/pm"
	"alto_server/routers/api/v1/provisioning"
	"alto_server/routers/api/v1/software"
	"alto_server/routers/api/v1/system"
	"alto_server/routers/api/v1/test"

	"github.com/gin-gonic/gin"
)

func setUpConfig(router *gin.RouterGroup) {
	//使用Token管理中间件, 保存用户信息
	// router.Use(middlewares.JWTAuthz())
	// 使用权限管理中间件
	//copy all the files to the same level with Alto.exe
	// e := casbin.NewEnforcer("conf/authz/model.conf", "conf/authz/policy.csv")
	// router.Use(middlewares.NewAuthorizer(e))

}

// RegisterRouter 注册路由
// nokia-alto/v1
func RegisterRouter(router *gin.RouterGroup) {

	v1 := router.Group("/v1")
	{
		setUpConfig(v1)
		// 用户路由
		test.RegisterRouter(v1.Group("/test"))
		provisioning.RegisterRouter(v1.Group("/provisioning"))
		alarm.RegisterRouter(v1.Group("/alarm"))
		pm.RegisterRouter(v1.Group("/pm"))
		system.RegisterRouter(v1.Group("/system"))
		software.RegisterRouter(v1.Group("/software"))
	}
}
