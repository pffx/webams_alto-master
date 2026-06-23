package routers

import (
	//"fmt"

	"alto_server/common/utils"
	auth "alto_server/routers/api/auth"
	user "alto_server/routers/api/user"
	v1 "alto_server/routers/api/v1"
	command "alto_server/ws/command"
	"net/http"

	// docs "alto_server/docs"

	"github.com/gin-contrib/static"
	// swaggerFiles "github.com/swaggo/files"
	// ginSwagger "github.com/swaggo/gin-swagger"

	"path/filepath"

	"github.com/gin-gonic/gin"
)

// InitRouter 初始化路由
func InitRouter() *gin.Engine {

	router := gin.Default()

	setUpConfig(router)
	setUpRouter(router)
	setUpWebSocket(router)

	return router
}

// 处理跨域请求,支持options访问
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin") //请求头部
		if origin != "" {
			// 当Access-Control-Allow-Credentials为true时，将*替换为指定的域名
			//c.Header("Access-Control-Allow-Origin", "http://localhost:5600")
			c.Header("Access-Control-Allow-Origin", origin)
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, X-Extra-Header, Content-Type, Accept, Authorization, token")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type, newtoken")
			c.Header("Access-Control-Allow-Credentials", "true")
			//c.Header("Cache-Control", "no-cache")
			//c.Header("Access-Control-Max-Age", "86400") // 可选
			//c.Header("Content-Type", "application/json")// 可选
		}

		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}

		c.Next()
	}
}

// 初始化应用设置
func setUpConfig(router *gin.Engine) {
	// 设置静态文件处理
	router.Use(func(c *gin.Context) {
		if c.Request.URL.Path == "/" {
			// 设置缓存头，告诉浏览器不要缓存该文件
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Pragma, Expires, Content-Language, Content-Type, newtoken")
			c.Header("Cache-Control", "no-cache, no-store, must-revalidate")
			c.Header("Pragma", "no-cache")
			c.Header("Expires", "0")
		}
		c.Next()
	})
	//static html files
	router.Use(static.Serve("/", static.LocalFile(utils.GetReactAppFilePath(), true)))
	router.NoRoute(func(c *gin.Context) {
		// 设置缓存头，告诉浏览器不要缓存该文件
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Pragma, Expires, Content-Language, Content-Type, newtoken")
		c.Header("Cache-Control", "no-cache, no-store, must-revalidate")
		c.Header("Pragma", "no-cache")
		c.Header("Expires", "0")
		c.File(filepath.Join(utils.GetReactAppFilePath(), "index.html"))
	})
	//OLT software stored in this folder, used for downloading
	router.Use(static.Serve("/software", static.LocalFile(utils.GetSoftwarePath(), true)))

	// 使用swagger自动生成接口文档
	// docs.SwaggerInfo.BasePath = "/nokia-alto"
	// router.StaticFile("swagger.json", "./docs/swagger.json")
	// url := ginSwagger.URL("http://localhost:5600/swagger.json")
	// router.GET("/help/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	router.Use(Cors())
}

func setUpWebSocket(router *gin.Engine) {
	wsGroup := router.Group("/ws")
	{
		wsGroup.GET("/cmd", command.ServeWs)
	}

}

// 设置路由
// nokia-alto
func setUpRouter(router *gin.Engine) {
	alto := router.Group("/nokia-alto")
	{
		v1.RegisterRouter(alto)
		auth.RegisterRouter(alto)
		user.RegisterRouter(alto)
	}
}
