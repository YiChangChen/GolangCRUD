package userservice

import "crud.com/crud/models"

type IUserService interface {
	GetUserList()
	GetUser(in string) (*models.User, error)
	CreateUser(in *models.User)
	UpdateUser(in *models.User)
	DeleteUser(in *models.User)
}
