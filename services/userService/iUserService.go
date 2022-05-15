package userservice

import "crud.com/crud/model"

type IUserService interface {
	GetUserList()
	GetUser(in *model.User)
	CreateUser(in *model.User)
	UpdateUser(in *model.User)
	ModifyUser(in *model.User)
	DeleteUser(in *model.User)
}
