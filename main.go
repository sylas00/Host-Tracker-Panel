package main

import (
	"Gin_web/routers"
	"fmt"
	"gopkg.in/ini.v1"
)

func main() {
	// 从ini文件获取配置信息
	cfg, err := ini.Load("conf/config.ini")
	if err != nil {
		panic("setting.Setup, fail to parse 'conf/app.ini'")
	}
	Port := cfg.Section("server").Key("HttpPort").String()
	//db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	//if err != nil {
	//	panic("init error")
	//}
	//db.AutoMigrate(&models.Model{})

	r := routers.SetupRouter()
	//3.监听端口，默认在8080
	if err := r.Run(":" + Port); err != nil {
		fmt.Printf("startup service failed, err:%v\n\n", err)
	}

}
