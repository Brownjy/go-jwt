package model

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func MySqlInit(dsn string) {
	//1.连接数据库
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("连接MySQL数据库失败,err: ", err)
		return
	}
	fmt.Println("连接数据库成功！")
	DB = db
	//2.迁移表
	err = db.AutoMigrate(&User{})
	if err != nil {
		fmt.Println("迁移MySQL数据表失败,err: ", err)
		return
	}
}
