package router

import "github.com/gin-gonic/gin"

func MainRouter(g *gin.Engine) {
	InitUserRouter(g) // USER
}
