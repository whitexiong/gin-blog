package dao

import (
	"blog/app/setting"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

// 初始化 mysql 连接
func InitMySQL(cfg *setting.MySQLConfig) (err error) {
	dsn := "root:123456@tcp(mysql.test:3306)/blog?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return
	}
	return err
}

// 关闭 mysql 连接
func Close() {
}
