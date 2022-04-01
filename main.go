package main

import (
	"Gin_web/models"
	"Gin_web/routers"
	"fmt"
	"gopkg.in/ini.v1"
)

func main() {
	// 从ini文件获取配置信息
	cfg, err := ini.Load("conf/config.ini")
	if err != nil {
		panic("fail to parse 'conf/app.ini'")
	}
	Port := cfg.Section("server").Key("HttpPort").String()

	//初始化DB
	models.InitDB()

	r := routers.SetupRouter()
	//3.监听端口，默认在8080
	if err := r.Run(":" + Port); err != nil {
		fmt.Printf("startup service failed, err:%v\n\n", err)
	}

}
