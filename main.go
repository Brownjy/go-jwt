package main

import (
	"fmt"
	"github.com/go-jwt/conf"
	"github.com/go-jwt/route"
)

func main() {
	//初始化config
	conf.ConfigInit()
	//加载路由route
	r := route.Router()
	//启动gin
	err := r.Run(":8080")
	if err != nil {
		fmt.Println("启动gin服务失败...,err: ", err)
	}
}
