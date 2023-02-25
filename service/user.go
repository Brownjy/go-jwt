package service

import (
	"fmt"
	"github.com/go-jwt/model"
	"github.com/go-jwt/utils"
)

type UseService struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

func (service *UseService) Register() {
	var user model.User
	//1.查看数据库是否存在该用户
	count, _ := user.GetUser(service.UserName)
	if count == 1 { // 已经存在该用户
		fmt.Println("该用户已存在")
		return
	}
	//2.若不存在则将用户的传来的账号与model模型绑定
	user.UserName = service.UserName
	//  在绑定密码之前需要对密码进行加密
	err := user.SetPassword(service.Password)
	if err != nil {
		fmt.Println("加密失败, err:", err)
		return
	}
	//3.创建该user
	err = user.CreateUser()
	if err != nil {
		fmt.Println("创建用户失败, err: ", err)
		return
	}
	fmt.Println("创建用户成功")
}
func (service *UseService) Login() {
	var user model.User
	// 查找User
	_, err := user.GetUser(service.UserName)
	if err != nil {
		fmt.Println("没有该用户, err:", err)
		return
	}
	// 解析该用户的密码
	if !user.CheckPassword(service.Password) {
		fmt.Println("解析密码失败, err: ", err)
		return
	}
	// 创建用户Token
	token, err := utils.GenerateToken(user)
	if err != nil {
		fmt.Println("创建Token失败, err: ", err)
		return
	}
	fmt.Println("token: ", token)
}
