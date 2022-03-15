package w

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/0014526/gin-demo/dao/custom"
	"github.com/0014526/gin-demo/dao/model"
	"github.com/0014526/gin-demo/initialize"
	"github.com/sony/sonyflake"
)

func CreateUserHandle(context *gin.Context) {
	user:=model.User{}

	flake:=sonyflake.NewSonyflake(sonyflake.Settings{})
	user_id, _:=flake.NextID()
	user.UserId= int(user_id)
	err := context.ShouldBind(&user)
	if err!=nil {
		log.Panicln(err)
	}
	custom.CreateUser(initialize.DB,&user)
}
