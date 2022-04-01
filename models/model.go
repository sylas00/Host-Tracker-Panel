package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// 为什么要这样写 因为这样可以给其他要用到db的地方直接导入 如果InitDB有返回值，应该也可以调用InitDB初始化一个db给其他地方用

var db *gorm.DB

func InitDB() {
	var err error
	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("init error")
	}
	// 自动迁移
	dbErr := db.AutoMigrate(&User{})

	if dbErr != nil {
		panic(nil)
	}
}
