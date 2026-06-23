package system

import (
	"github.com/gin-gonic/gin"
)

// RegisterRouter 注册路由
// nokia-alto/v1/system
func RegisterRouter(r *gin.RouterGroup) {

	r.GET("/", todoSystemHandler)

	r.GET("/oltList", getOLTListHandler)
	r.PUT("/oltList", addOLTListHandler)
	r.POST("/oltConn", connectOltHandler)
	r.POST("/rpc", rpcPushHandler)
	r.GET("/nt_gw", ntGWHandler)
}
