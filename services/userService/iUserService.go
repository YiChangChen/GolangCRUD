package userservice

import "crud.com/crud/model"

type IUserService interface {
	GetUserList()
	GetUser(in *model.User) (*model.User, error)
	CreateUser(in *model.User)
	UpdateUser(in *model.User)
	DeleteUser(in *model.User)
}
