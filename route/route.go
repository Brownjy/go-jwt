package route

import (
	"github.com/gin-gonic/gin"
	"github.com/go-jwt/controller"
	"github.com/go-jwt/middleware"
)

func Router() *gin.Engine {
	r := gin.Default()
	//r.POST("/login", controller.UserLogin)
	r.POST("/Register", controller.UserRegister)
	authorized := r.Group("/", middleware.JWTAuth())
	{
		authorized.GET("/sayHello", controller.UserSayHello)
	}
	return r
}
