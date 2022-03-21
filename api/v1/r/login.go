package r

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


func Login(context *gin.Context)  {
	context.JSON(http.StatusOK,gin.H{
		"code":200,
		"message":"success",
	})
}


func LoginCallBack(context *gin.Context)  {

}

func LoginOut(context *gin.Context) {

}