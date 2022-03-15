package model

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/0014526/gin-demo/util/setting"
)
var (
	db *gorm.DB
)


func Setup() {
	var err error
	db, err = gorm.Open(setting.MysqlSetting.Type, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.MysqlSetting.User,
		setting.MysqlSetting.Password,
		setting.MysqlSetting.Host,
		setting.MysqlSetting.DB))
	log.Println(db)
	if err != nil {
		log.Fatalf("model.Setup err: %v", err)
	}
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}

func CloseDB() {
	defer db.Close()
}