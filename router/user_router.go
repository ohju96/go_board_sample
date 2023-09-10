package router

import "github.com/gin-gonic/gin"

func InitUserRouter(g *gin.Engine) {
	v1 := g.Group("/api/v1/users")
	{
		v1.POST("")    // 회원가입
		v1.GET("/:id") // 회원정보 조회
	}
}
