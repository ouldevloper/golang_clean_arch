package usecases

import "go_clean/src/domain"

type UserInteractor struct {
	repository *UserRepository
}

func (obj *UserInteractor) Add(u interface{}) {
	obj.repository.Store(u)
}

func (obj *UserInteractor) GetInfo() []domain.User {
	users := []domain.User{}
	obj.repository.Select(&users)
}
