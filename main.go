package main

import (
	"fmt"

	"github.com/0014526/gin-demo/router"
	"github.com/0014526/gin-demo/util/setting"
	"github.com/0014526/gin-demo/initialize"
)

func init() {
	setting.InitCofig()
	initialize.InitGormMysql()
}

func main() {
	run := router.InitRouter()
	run.Run(fmt.Sprintf(":%d",setting.ServerSetting.HttpPort))
}
