package controller

import (
	"ginSample/app/service"
	"github.com/gin-gonic/gin"
)

type UserController interface {
	CreateUser(ctx *gin.Context)
}

type userController struct {
	userService service.UserService
}

func (u userController) CreateUser(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func NewUserController(userService service.UserService) UserController {
	return &userController{
		userService: userService,
	}
}
