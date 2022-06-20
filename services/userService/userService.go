package userservice

import (
	"fmt"

	"crud.com/crud/models"
	sqlservice "crud.com/crud/services/sqlService"
)

type UserService struct {
	sqlservice *sqlservice.Dto
}

func NewUserService() *UserService {
	return &UserService{}
}

func (svc *UserService) CreateUser(in *models.User) (int64, error) {
	res, err := svc.sqlservice.CreateUser(in)
	if err != nil {
		return -1, err
	}
	return res, nil
}

func (svc *UserService) UpdateUser(in *models.User) (int64, error) {
	res, err := svc.sqlservice.UpdateUser(in)
	if err != nil {
		return -1, err
	}
	return res, nil
}

func (svc *UserService) DeleteUser(in *models.User) (int64, error) {
	res, err := svc.sqlservice.DeleteUser(in)
	if err != nil {
		return -1, err
	}
	return res, nil
}

func (svc *UserService) GetUserList() {
	fmt.Printf("get all user \n")
}

func (svc *UserService) GetUser(in string) (*models.User, error) {
	user, err := svc.sqlservice.GetUser(in)
	if err != nil {
		fmt.Printf("Error get UserProfile: %s \n", err.Error())
	}
	fmt.Printf("get user %s \n", user.Username)
	return user, err
}
