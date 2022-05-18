package userservice

import (
	"fmt"
	"log"

	"crud.com/crud/model"
	sqlservice "crud.com/crud/services/sqlService"
)

type UserService struct {
}

func (svc *UserService) CreateUser(in *model.User) {
	sqlSvc := new(sqlservice.MsSql)
	_, err := sqlSvc.CreateUser(in)
	if err != nil {
		log.Fatal("Error creating UserProfile: ", err.Error())
	}
}

func (svc *UserService) UpdateUser(in *model.User) {
	fmt.Printf("update %s \n", in.Username)
}

func (svc *UserService) ModifyUser(in *model.User) {
	fmt.Printf("modify %s \n", in.Username)
}

func (svc *UserService) DeleteUser(in *model.User) {
	fmt.Printf("delete %s \n", in.Username)
}

func (svc *UserService) GetUserList() {
	fmt.Printf("get all user \n")
}

func (svc *UserService) GetUser(in *model.User) {
	fmt.Printf("get user %s \n", in.Username)
}
