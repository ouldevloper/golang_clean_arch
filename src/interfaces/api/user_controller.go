package api

import (
	"go_clean/src/infrastructure"
	"go_clean/src/usecases"
)

type UserController struct {
	Interactor usecases.UserInteractor
}

func NewUserController(Sqlhandler *infrastructure.SqlHandler) *UserController {
	return &UserController{
		Interactor: usecases.UserInteractor{
			UserRepository: &usecases.UserRepository{
				Sqlhandler: Sqlhandler,
			},
		},
	}
}
