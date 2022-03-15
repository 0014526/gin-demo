package initialize

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/0014526/gin-demo/util/setting"
)
var	DB *gorm.DB


func InitGormMysql() {
	var err error
	DB, err = gorm.Open(setting.MysqlSetting.Type, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.MysqlSetting.User,
		setting.MysqlSetting.Password,
		setting.MysqlSetting.Host,
		setting.MysqlSetting.DB))
	if err != nil {
		log.Fatalf("model.Setup err: %v", err)
	}
	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(100)
}

func CloseDB() {
	defer DB.Close()
}

func GetDB() *gorm.DB{
	return DB
}
