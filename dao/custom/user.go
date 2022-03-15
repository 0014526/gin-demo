package custom

import (
	"github.com/jinzhu/gorm"

	"github.com/0014526/gin-demo/dao/model"
)

func GetUserById(db *gorm.DB,id int) (model.User,error) {
	user:=model.User{}
	db.Model("tbl_user").Find(&user,"id=?",id)
	return user,nil
}

func GetAllUserList(db *gorm.DB) ([]model.User,error)  {
	users:=[]model.User{}
	db.Model("tbl_user").Find(&users,"1=1")
	return users,nil
}

func CreateUser(db *gorm.DB,user *model.User) bool{
	result:=db.Model("tbl_user").Create(&user)
	if result.RowsAffected>=1 {
		return true
	}else{
		return false
	}
}

func UpdateUserName(db *gorm.DB, id int, user_name string) int64 {
	result:=db.Model("tbl_user").Where("id=?",id).Update("user_name",user_name)
	return result.RowsAffected
}