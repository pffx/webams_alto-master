package log

import (
	"github.com/gin-gonic/gin"
)

// RegisterRouter 注册路由
//nokia-alto/v1/log
func RegisterRouter(r *gin.RouterGroup) {

	r.GET("/", todoLogHandler)
}
