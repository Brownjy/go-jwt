package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-jwt/service"
	"net/http"
)

// UserLogin 用户登录
func UserLogin(c *gin.Context) {
	var userRegisterService service.UseService
	err := c.ShouldBind(&userRegisterService)
	if err != nil {
		fmt.Println("参数绑定失败！,err: ", err)
		return
	}
	userRegisterService.Login()
}

// UserRegister 用户注册
func UserRegister(c *gin.Context) {
	var userRegisterService service.UseService
	err := c.ShouldBind(&userRegisterService)
	fmt.Println(userRegisterService)
	if err != nil {
		fmt.Println("参数绑定失败！,err: ", err)
		return
	}
	userRegisterService.Register()
}

// UserSayHello 用户SayHello
func UserSayHello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "hello " + c.GetString("user"),
	})
}
