package usecases

import (
	"go_clean/src/domain"
	"go_clean/src/interfaces/database"
)

type UserRepository struct {
	database.SqlHandler
}

func (obj *UserRepository) Store(u interface{}) interface{} {
	obj.Create(u)
	return u
}

func (obj *UserRepository) Select() []domain.User {
	users := []domain.User{}
	obj.Find(&users)
	return users
}

func (obj *UserRepository) DeleteUser(id int) []domain.User {
	user := []domain.User{}
	obj.Delete(&user, id)
	return user
}
