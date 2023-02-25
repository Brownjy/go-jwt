package route

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-jwt/controller"
	"github.com/go-jwt/middleware"
	"net/http"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.POST("/login", controller.UserLogin)
	r.POST("/Register", controller.UserRegister)
	authorized := r.Group("/", middleware.JWTAuth())
	{
		authorized.GET("/sayHello", controller.UserSayHello)
		authorized.GET("/say", func(c *gin.Context) {
			value, exists := c.Get("user")
			if !exists {
				fmt.Println("不存在user")
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"msg": value,
			})
		})
	}
	return r
}
