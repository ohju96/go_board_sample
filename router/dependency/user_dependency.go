package dependency

import (
	"ginSample/app/controller"
	"ginSample/app/repository"
	"ginSample/app/service"
	"ginSample/config/db"
)

func InitUserDependency() controller.UserController {
	userRepository := repository.NewUserRepository(*db.MySQL)
	userService := service.NewUserService(userRepository)
	return controller.NewUserController(userService)
}
