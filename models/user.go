package models

import (
	"gorm.io/gorm"
	"time"
)

// default:(-)是无默认值

type User struct {
	ID       int64          `gorm:"primaryKey;column:id"`
	Username string         `gorm:"column:username;type:varchar(255);default:(-)"`
	Password string         `gorm:"column:password;type:varchar(255);default:(-)"`
	CreateAt time.Time      `gorm:"column:create_time;type:timestamp;default:(-)"`
	UpdateAt time.Time      `gorm:"column:update_time;type:timestamp;default:(-)"`
	DeleteAt gorm.DeletedAt `gorm:"column:delete_time;type:timestamp;default:(-)"`
}
