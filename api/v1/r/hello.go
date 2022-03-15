package r

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetHelloHandle(context *gin.Context) {
	context.JSON(http.StatusOK,gin.H{
		"code":http.StatusOK,
		"message":"success",
	})
}