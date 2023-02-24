package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/go-jwt/utils"
	"net/http"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.Request.Header.Get("Authorization")
		if len(auth) == 0 {
			c.Abort()
			c.String(http.StatusOK, "未登录无权限")
			return
		}
		// 校验token，只要出错直接拒绝请求
		claims, err := utils.ParseToken(auth)
		// 设置当前用户
		c.Set("user", claims.User.UserName)
		if err != nil {
			c.Abort()
			message := err.Error()
			c.JSON(http.StatusOK, message)
			return
		} else {
			println("token 正确")
		}
		c.Next()
	}
}
