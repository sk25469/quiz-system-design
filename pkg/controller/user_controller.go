package controller

import (
	"quiz-design/pkg/model"
	"quiz-design/pkg/service"
)

type UserController struct {
	userService service.UserService
}

func CreateUserController() *UserController {
	return &UserController{
		userService: *service.CreateNewUserService(),
	}
}

func (uc *UserController) RegisterUser(userType model.UserType, name, email string) {
	uc.userService.InsertUser(name, email, userType)
}
