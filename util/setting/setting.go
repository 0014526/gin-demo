package setting

import (
	"log"
	"time"

	"gopkg.in/ini.v1"
)

var (
	cfg *ini.File
)

type ServerConfig struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	MaxMultipartMemory int64
}
var ServerSetting = &ServerConfig{}

type MySqlConfig struct {
	Type     string
	User     string
	Password string
	DB       string
	Host     string
	Port     int
}
var MysqlSetting = &MySqlConfig{}

type RedisConfig struct {
	Host        string
	Password    string
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
}
var RedisSetting = &RedisConfig{}

// 加载ini文件绑定到struct中
func InitCofig() {
	var err error
	cfg, err = ini.Load("config/config.ini")
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse 'config/config.ini': %v", err)
	}

	mapTo("server", ServerSetting)
	mapTo("mysql", MysqlSetting)
	mapTo("redis", RedisSetting)

	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second
	RedisSetting.IdleTimeout = RedisSetting.IdleTimeout * time.Second
}

// 绑定ini文件
func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}
