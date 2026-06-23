package pm

import (
	"github.com/gin-gonic/gin"
)

// RegisterRouter 注册路由
//nokia-alto/v1/pm
func RegisterRouter(r *gin.RouterGroup) {

	r.GET("/", todoPmHandler)
}
