package user

import "github.com/gin-gonic/gin"

type IUserController interface {
	GetUserList(c *gin.Context)

	GetUser(c *gin.Context)
	CreateUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
}
