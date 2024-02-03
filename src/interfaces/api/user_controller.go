package api

import (
	"go_clean/src/domain"
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
				SqlHandler: *Sqlhandler,
			},
		},
	}
}

func (obj *UserController) GetUser() []domain.User {
	res := obj.Interactor.GetInfo()
	return res
}
