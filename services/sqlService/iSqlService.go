package sqlservice

import "crud.com/crud/model"

type ISqlService interface {
	GetUserList()
	GetUser(in *model.User) (*model.User, error)
	CreateUser(in *model.User) (int64, error)
	UpdateUser(in *model.User) (int64, error)
	DeleteUser(in *model.User) (int64, error)
}
