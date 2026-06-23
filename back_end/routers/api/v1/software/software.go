package software

import (
	"alto_server/routers/api/v1/software/olt"
	"alto_server/routers/api/v1/software/ont"

	"github.com/gin-gonic/gin"
)

// RegisterRouter 注册路由
//nokia-alto/v1/software
func RegisterRouter(r *gin.RouterGroup) {

	ont.RegisterRouter(r.Group("/ont"))
	olt.RegisterRouter(r.Group("/olt"))
}
