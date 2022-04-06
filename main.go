package main

import (
	"Gin_web/configs"
	"Gin_web/internal/dao"
	"Gin_web/internal/routers"
	"fmt"
)

func main() {

	//初始化配置项
	configs.InitConfig()
	//初始化DB
	dao.InitDB()

	r := routers.SetupRouter()
	//3.从配置文件取监听端口，默认是8080
	Port := fmt.Sprintf(":%d", configs.ServerConf.HttpPort)
	if err := r.Run(Port); err != nil {
		fmt.Printf("startup service failed, err:%v\n\n", err)
	}

}
