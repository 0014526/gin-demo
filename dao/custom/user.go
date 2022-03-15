package custom

import (
	"github.com/0014526/gin-demo/dao/model"
	"github.com/0014526/gin-demo/initialize"
)


func GetUserById(id string) (model.User,error) {
	user:=model.User{}
	initialize.DB.Model("tbl_user").Find(&user,"id=?",id)
	return user,nil
}

func GetAllUserList() ([]model.User,error)  {
	users:=[]model.User{}
	initialize.DB.Model("tbl_user").Find(&users,"1=1")
	return users,nil
}

func CreateUser(user *model.User) error{
	err:=initialize.DB.Model("tbl_user").Create(&user).Error
	return err
}

func UpdateUserName(id string, user_nike string) error {
	user:=model.User{}
	initialize.DB.Model(&user).Where("id=?",id).Update("user_nike",user_nike)
	return nil
}

func DeleteUser(id string) error{
	initialize.DB.Model("tbl_user").Where("id=?",id).Update("status",0)
	return nil
}