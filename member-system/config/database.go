package config

import (
	"member-system/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := "root:ZSYyy1012@tcp(127.0.0.1:3306)/test_db?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败: " + err.Error())
	}

	// 自动迁移（建表）
	DB.AutoMigrate(&models.User{})
}
