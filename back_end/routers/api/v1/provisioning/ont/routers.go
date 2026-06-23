package ont

import (
	"github.com/gin-gonic/gin"
)

// RegisterRouter 注册路由
//nokia-alto/v1/provisioning/ont
func RegisterRouter(r *gin.RouterGroup) {

	r.GET("/todo", todoHandler)
}
