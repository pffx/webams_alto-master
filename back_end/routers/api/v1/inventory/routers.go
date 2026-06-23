package inventory

import (
	"github.com/gin-gonic/gin"
)

// RegisterRouter 注册路由
//nokia-alto/v1/inventory
func RegisterRouter(r *gin.RouterGroup) {

	r.GET("/", todoInventoryHandler)
}
