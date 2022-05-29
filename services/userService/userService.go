package userservice

import (
	"fmt"

	"crud.com/crud/models"
	sqlservice "crud.com/crud/services/sqlService"
)

type UserService struct {
}

func (svc *UserService) CreateUser(in *models.User) {
	sqlSvc := new(sqlservice.MsSql)
	_, err := sqlSvc.CreateUser(in)
	if err != nil {
		fmt.Printf("Error creating UserProfile: %s \n", err.Error())
	}
	fmt.Printf("Createing: %s \n", in.Username)
}

func (svc *UserService) UpdateUser(in *models.User) {
	sqlSvc := new(sqlservice.MsSql)
	_, err := sqlSvc.UpdateUser(in)
	if err != nil {
		fmt.Printf("Error updating UserProfile: %s \n", err.Error())
	}
	fmt.Printf("update %s \n", in.Username)
}

func (svc *UserService) DeleteUser(in *models.User) {
	sqlSvc := new(sqlservice.MsSql)
	_, err := sqlSvc.DeleteUser(in)
	if err != nil {
		fmt.Printf("Error delete UserProfile: %s \n", err.Error())
	}
	fmt.Printf("delete %s \n", in.Username)
}

func (svc *UserService) GetUserList() {
	fmt.Printf("get all user \n")
}

func (svc *UserService) GetUser(in string) (*models.User, error) {
	sqlSvc := new(sqlservice.MsSql)
	user, err := sqlSvc.GetUser(in)
	if err != nil {
		fmt.Printf("Error get UserProfile: %s \n", err.Error())
	}
	fmt.Printf("get user %s \n", user.Username)
	return user, err
}
