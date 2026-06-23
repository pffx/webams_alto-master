package olt

import (
	"github.com/gin-gonic/gin"
)

// RegisterRouter 注册路由
// nokia-alto/v1/provisioning/olt
func RegisterRouter(r *gin.RouterGroup) {

	r.POST("/backup", backupHandler)
	r.POST("/restore", restoreHandler)
	r.POST("/reset", resetHandler)
	r.POST("/reset_all", resetAllHandler)
	r.POST("/mf_gui", webguiDeployHandler)
	r.POST("/mf_gui_un", webguiUndeployHandler)
	r.GET("/mf_gui", getWebguiDeploidResultHandler)
	r.POST("/initial_file", uploadConfigXmlHandler)
	r.POST("/ping", pingHandler)
	// r.POST("/ping_gw", getPingGW)
}
