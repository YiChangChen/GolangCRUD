package sqlservice

import "crud.com/crud/models"

type ISqlService interface {
	GetUserList()
	GetUser(in string) (*models.User, error)
	CreateUser(in *models.User) (int64, error)
	UpdateUser(in *models.User) (int64, error)
	DeleteUser(in *models.User) (int64, error)
}
