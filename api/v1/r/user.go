package r

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/0014526/gin-demo/dao/custom"
	"github.com/0014526/gin-demo/initialize"
)

func GetUserHandle(context *gin.Context) {
	user, err := custom.GetUserById(initialize.DB, 11)
	if err != nil {
		log.Panicln(err)
	}
	context.JSON(http.StatusOK,user)
}

func GetUserListHandle(context *gin.Context) {

}
