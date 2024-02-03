package usecases

import (
	"go_clean/src/domain"
	"go_clean/src/infrastructure"
)

type UserRepository struct {
	infrastructure.SqlHandler
}

func (obj *UserRepository) Store(u *domain.User) {
	obj.Create(u)
}


func (obj *UserRepository)Select()[]domain.User{
	users:=[]domain.User
	obj.FindAll(&users)
	return users
}


func (obj*UserRepository) Delete(id int){
	userv:= []domain.User{}
	obj.Delete(&user,id)
}