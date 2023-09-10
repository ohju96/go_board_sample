package router

import (
	"ginSample/router/dependency"
	"github.com/gin-gonic/gin"
)

func InitUserRouter(g *gin.Engine) {
	v1 := g.Group("/api/v1/users")
	{
		userController := dependency.InitUserDependency()

		v1.POST("", userController.CreateUser) // 회원가입
		v1.GET("/:id")                         // 회원정보 조회
	}
}
