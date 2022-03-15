package router

import (
	"github.com/gin-gonic/gin"
	"github.com/0014526/gin-demo/api/v1/r"
)
func InitRouter() *gin.Engine {
	router:=gin.New()
	v1 := router.Group("/api/v1")
	{
		// 后期封装到中间件中，批量处理
		v1.GET("/hello", r.GetHelloHandle)
	}

	return router
}