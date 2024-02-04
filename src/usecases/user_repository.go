package usecase

import "go_clean/src/domain"

type UserRepository interface {
	Store(domain.User)
	Select() []domain.User
	Delete(id string)
}
