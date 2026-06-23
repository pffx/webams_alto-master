package provisioning

import (
	"alto_server/routers/api/v1/provisioning/olt"
	"alto_server/routers/api/v1/provisioning/ont"
	"alto_server/routers/api/v1/provisioning/pon"
	"alto_server/routers/api/v1/provisioning/service"

	"github.com/gin-gonic/gin"
)

// RegisterRouter 注册路由
//nokia-alto/v1/provisioning
func RegisterRouter(r *gin.RouterGroup) {

	service.RegisterRouter(r.Group("/service"))
	ont.RegisterRouter(r.Group("/ont"))
	olt.RegisterRouter(r.Group("/olt"))
	pon.RegisterRouter(r.Group("/pon"))

}
