package w

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sony/sonyflake"

	"github.com/0014526/gin-demo/dao/custom"
	"github.com/0014526/gin-demo/dao/model"
)

// 创建user
func CreateUserHandle(context *gin.Context) {
	user:=model.User{}
	flake:=sonyflake.NewSonyflake(sonyflake.Settings{})
	user_id, _:=flake.NextID()
	user.UserId= int(user_id)
	user.Status= 1
	err := context.ShouldBind(&user)
	if err!=nil {
		log.Panicln(err)
	}
	err = custom.CreateUser(&user)
	if err != nil{
		context.JSON(http.StatusOK, gin.H{"error": err.Error()})
	}else {
		context.JSON(http.StatusOK, gin.H{
			"code":http.StatusOK,
			"message":"success",
			"data":user,
		})
	}
}

// 删除user
func DeleteUserHandle(context *gin.Context) {
	id:=context.PostForm("id")

	// id1 := com.StrTo(context.Param("id1")).MustInt()
	// log.Println(id1)
	err := custom.DeleteUser(id)
	if err != nil {
		context.JSON(http.StatusOK,gin.H{"error":err.Error()})
	}else{
		context.JSON(http.StatusOK, gin.H{id:"deleted"})
	}
}

// 修改user
func UpdateUserById(context *gin.Context)  {
	id:=context.Query("id")
	user:=model.User{}
	user_nike:=context.PostForm("user_nike")
	err := custom.UpdateUserName(id, user_nike)
	if err != nil {
		context.JSON(http.StatusOK,gin.H{"error":err.Error()})
	}else{
		context.JSON(http.StatusOK, gin.H{
			"code":http.StatusOK,
			"message":"success",
			"data":user,
		})
	}
}
