package models

type MonitorDomain struct {
	ID       int64  `gorm:"column:id"`
	username string `gorm:"column:username"`
	password string `gorm:"column:password"`
}
