package models

import (
	"Gin_web/conf"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 为什么要这样写 因为这样可以给其他要用到db的地方直接导入 如果InitDB有返回值，应该也可以调用InitDB初始化一个db给其他地方用

var DB *gorm.DB

func InitDB() {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		conf.DatabaseConf.User,
		conf.DatabaseConf.Password,
		conf.DatabaseConf.Host,
		conf.DatabaseConf.Name)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("init error")
	}
	// 自动迁移
	dbErr := DB.AutoMigrate(&User{})

	if dbErr != nil {
		panic(nil)
	}
}
