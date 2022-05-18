package sqlservice

import "crud.com/crud/model"

type ISqlService interface {
	GetUserList()
	GetUser(in *model.User)
	CreateUser(in *model.User) (int64, error)
	UpdateUser(in *model.User)
	ModifyUser(in *model.User)
	DeleteUser(in *model.User)
}
