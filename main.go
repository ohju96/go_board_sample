package main

import (
	"ginSample/config"
	"ginSample/config/db"
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
	toml := config.InitToml("./config/config.toml")

	// db
	db.InitMySQL(&toml)
}
