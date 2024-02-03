package database

import "go_clean/src/domain"

type UserRepository interface {
	Store(*domain.User)
	Select() []domain.User
	Delete(int)
}
