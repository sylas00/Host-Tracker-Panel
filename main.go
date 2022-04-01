package main

import (
	"Gin_web/conf"
	"Gin_web/models"
	"Gin_web/routers"
	"fmt"
)

func main() {

	//初始化配置项
	conf.InitConfig()
	//初始化DB
	models.InitDB()

	r := routers.SetupRouter()
	//3.从配置文件取监听端口，默认是8080
	Port := fmt.Sprintf(":%d", conf.ServerConf.HttpPort)
	if err := r.Run(Port); err != nil {
		fmt.Printf("startup service failed, err:%v\n\n", err)
	}

}
