package test

import (
	"github.com/gin-gonic/gin"
)

// RegisterRouter 注册路由
//nokia-alto/v1/test/
func RegisterRouter(r *gin.RouterGroup) {

	r.GET("/basicInfo", getBasicInfoHandler)
	r.POST("/basicInfo", setBasicInfoHandler)
	r.PUT("/basicInfo", setBasicInfoHandler)
	r.DELETE("/basicInfo", setBasicInfoHandler)
	r.GET("/someXML", someXMLHandler)
	r.GET("/testRpc", testRpcHandler)
	r.POST("/upload", uploadXmlHandler)
}
