package main

import (
	"Gin_web/routers"
	"fmt"
)

func main() {
	//db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	//if err != nil {
	//	panic("init error")
	//}
	//db.AutoMigrate(&models.Model{})

	r := routers.SetupRouter()
	// 3.监听端口，默认在8080
	// Run("里面不指定端口号默认为8080")
	if err := r.Run(":8000"); err != nil {
		fmt.Println("startup service failed, err:%v\n", err)
	}

}
