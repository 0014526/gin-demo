package router

import (
	"github.com/gin-gonic/gin"
	"github.com/0014526/gin-demo/api/v1/r"
	"github.com/0014526/gin-demo/api/v1/w"
)
func InitRouter() *gin.Engine {
	router:=gin.New()
	v1 := router.Group("/api/v1")
	{
		// 后期封装到中间件中，批量处理
		v1.GET("/hello", r.GetHelloHandle)
		r1:=v1.Group("/r")
		{
			r1.GET("/getUserById",r.GetUserHandle)
			r1.GET("/getAllUser",r.GetUserListHandle)
		}
		w1:=v1.Group("/w")
		{
			w1.POST("/CreateUser",w.CreateUserHandle)
		}
	}

	return router
}