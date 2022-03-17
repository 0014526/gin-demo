package r

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/0014526/gin-demo/util/util"
)

func GetHelloHandle(context *gin.Context) {
	context.JSON(http.StatusOK,gin.H{
		"code":http.StatusOK,
		"message":"success",
	})
}

func TestBcrypt(context *gin.Context) {
	password:=context.Query("password")
	generatePassword, err := util.GeneratePassword(password)
	if err != nil {
		log.Println("加密出错了")
	}
	passwordstring := string(generatePassword)
	log.Println(passwordstring)
	log.Println(generatePassword)
	mysql_password := "$2a$10$I8WaWXgiBw8j/IBejb3t/.s5NoOYLvoQzL6mIM2g3TEu4z0HenzqK"
	isOk, _ := util.ValidatePassword(passwordstring, mysql_password)
	if !isOk {
		log.Println("密码错误")
		return
	}
	log.Println(isOk)
}

func TestReflect() {

}