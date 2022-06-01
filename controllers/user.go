package controllers

import (
	"net/http"
	"strconv"

	"crud.com/crud/models"
	userservice "crud.com/crud/services/userService"
	"github.com/gin-gonic/gin"
)

// @Summary 查詢user
// @Description search user by id
// @Tags     user
// @Accept  json
// @Produce  json
// @Param   id   path      string  true  "Search User"
// @Success 200 {string} string	"ok"
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Router /user/{id} [get]
func (c *Controller) GetUser(ctx *gin.Context) {
	id := ctx.Param("id")
	_, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": 9999})
		return
	}

	svc := new(userservice.UserService)
	result, err := svc.GetUser(id)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": 9999})
		return
	}
	ctx.JSON(http.StatusOK, result)
}

// @Summary 新增user
// @Description create user
// @Tags     user
// @Accept  json
// @Produce  json
// @Param   user  body      models.User  true  "Create User"
// @Success 200 {string} string	"ok"
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Router /user/create [post]
func (c *Controller) CreateUser(ctx *gin.Context) {
	var data = new(models.User)

	if ctx.ShouldBindJSON(&data) != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": 9999})
		return
	}

	svc := new(userservice.UserService)
	svc.CreateUser(data)
}

// @Summary 更新user
// @Description update user
// @Tags    user
// @Accept  json
// @Produce  json
// @Param   user  body      models.User  true  "update User"
// @Success 200 {string} string	"ok"
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Router /user/update [put]
func (c *Controller) UpdateUser(ctx *gin.Context) {
	var data = new(models.User)

	if ctx.ShouldBindJSON(&data) != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": 9999})
		return
	}

	svc := new(userservice.UserService)
	svc.UpdateUser(data)
}

// @Summary 刪除user
// @Description delete user
// @Tags    user
// @Accept  json
// @Produce  json
// @Param   user  body      models.User  true  "delete User"
// @Success 200 {string} string	"ok"
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Router /user/delete [delete]
func (c *Controller) DeleteUser(ctx *gin.Context) {
	var data = new(models.User)

	if ctx.ShouldBindJSON(&data) != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": 9999})
		return
	}

	svc := new(userservice.UserService)
	svc.DeleteUser(data)
}
