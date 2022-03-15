package r

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/0014526/gin-demo/dao/custom"
)

func GetUserHandle(context *gin.Context) {
	id:=context.Query("id")
	user, err := custom.GetUserById(id)
	if err != nil {
		log.Panicln(err)
	}
	context.JSON(http.StatusOK,gin.H{
		"code":http.StatusOK,
		"message":"success",
		"data":user,
	})
}

func GetUserListHandle(context *gin.Context) {
	list, err := custom.GetAllUserList()
	if err != nil {
		context.JSON(http.StatusOK, gin.H{"error": err.Error()})
	}else{
		context.JSON(http.StatusOK,gin.H{
			"code":http.StatusOK,
			"message":"success",
			"data":list,
		})
	}
}
