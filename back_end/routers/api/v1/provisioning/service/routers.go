package service

import (
	"github.com/gin-gonic/gin"
)

// RegisterRouter 注册路由
//nokia-alto/v1/provisioning/service
func RegisterRouter(r *gin.RouterGroup) {

	r.GET("/services", getAllServicesHandler)
	r.PUT("/services", createServicesHandler)
}
