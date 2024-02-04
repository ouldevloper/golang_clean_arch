package api

import (
	"go_clean/src/domain"
	"go_clean/src/interfaces/database"
	"go_clean/src/usecases"
)

type UserController struct {
	Interactor usecases.UserInteractor
}

func NewUserController(Sqlhandler database.SqlHandler) *UserController {
	return &UserController{
		Interactor: usecases.UserInteractor{
			UserRepository: &usecases.UserRepository{
				SqlHandler: Sqlhandler,
			},
		},
	}
}

func (obj *UserController) GetUser() []domain.User {
	res := obj.Interactor.
	return res
}

func (obj *UserController) Create(u *domain.User) {
	obj.Create(u)
}

func (obj *UserController) DeleteUser() {
	obj.Interactor.UserRepository.Delete()
}
