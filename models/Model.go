package models

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	fmt.Println("Db init")
	dsn := "root:123456@tcp(127.0.0.1:3306)/days7?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db, err := DB.DB()
	if err != nil {
		panic(err)
	}
	db.SetMaxIdleConns(50)
	// 打开链接
	db.SetMaxOpenConns(200)
}
