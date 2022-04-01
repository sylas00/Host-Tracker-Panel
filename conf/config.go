package conf

import (
	"gopkg.in/ini.v1"
	"log"
)

// Database 数据库配置
type Database struct {
	Type     string
	User     string
	Password string
	Host     string
	Name     string
}

var DatabaseConf = &Database{}

// Server 服务器的配置
type Server struct {
	//RunMode  string
	HttpPort int
}

var ServerConf = &Server{}

var cfg *ini.File

func InitConfig() {
	// 从ini文件获取配置信息
	var err error
	cfg, err = ini.Load("conf/config.ini")
	if err != nil {
		panic("fail to parse 'conf/app.ini'")
	}
	//取单个配置项
	//Port := cfg.Section("server").Key("HttpPort").String()

	//直接映射
	mapTo("database", DatabaseConf)
	mapTo("server", ServerConf)
}

// mapTo map section
func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}
