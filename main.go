package main

import (
	"ginSample/config"
	"ginSample/router"
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	Init(app)

	app.Run()
}

func Init(app *gin.Engine) {
	// router
	router.MainRouter(app)

	// toml
	config.InitToml("./config/config.toml")

}
