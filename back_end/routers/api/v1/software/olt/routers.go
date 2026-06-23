package olt

import (
	"github.com/gin-gonic/gin"
)

// RegisterRouter 注册路由
// nokia-alto/v1/software/olt/
func RegisterRouter(r *gin.RouterGroup) {

	r.GET("/server_software", getAvailableSoftwareList)
	r.GET("/software", getOltSoftwareList)
	r.PUT("/software", handleOltSoftware)
	r.PUT("/software/migration", handleOltConfigMigration)
	r.PUT("/software/migration_upload", handleOltConfigMigrationUploading)
	//Add For Test
	r.POST("/downloadOltSoftware", downloadOltSoftware)
	r.POST("/activeOltSoftware", activeOltSoftware)
	r.POST("/commitOltSoftware", commitOltSoftware)
	//Add For Test
}
