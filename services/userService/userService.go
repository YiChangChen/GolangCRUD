package userservice

import (
	"fmt"

	"crud.com/crud/model"
)

type UserService struct {
}

func (svc *UserService) CreateUser(in *model.User) {
	fmt.Printf("add %s \n", in.Username)
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
