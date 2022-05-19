package main

import (
	"net/http"

	"crud.com/crud/model"
	userservice "crud.com/crud/services/userService"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	server.LoadHTMLGlob("webroot/*")

	server.GET("/ping", test)
	server.GET("/", homePage)

	server.GET("/:id", GetUser)
	server.POST("/create", CreateUser)
	server.PUT("/update", UpdateUser)
	server.DELETE("/delete", DeleteUser)

	server.Run(":8080")
}

func test(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func homePage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

func GetUser(c *gin.Context) {
	var data = new(model.User)
	if c.ShouldBindJSON(&data) != nil {
		c.JSON(http.StatusOK, gin.H{"error": 9999})
		return
	}

	svc := new(userservice.UserService)
	result, err := svc.GetUser(data)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": 9999})
		return
	}
	c.JSON(http.StatusOK, result)
}

func CreateUser(c *gin.Context) {
	var data = new(model.User)

	if c.ShouldBindJSON(&data) != nil {
		c.JSON(http.StatusOK, gin.H{"error": 9999})
		return
	}

	svc := new(userservice.UserService)
	svc.CreateUser(data)
}

func UpdateUser(c *gin.Context) {
	var data = new(model.User)

	if c.ShouldBindJSON(&data) != nil {
		c.JSON(http.StatusOK, gin.H{"error": 9999})
		return
	}

	svc := new(userservice.UserService)
	svc.UpdateUser(data)
}

func DeleteUser(c *gin.Context) {
	var data = new(model.User)

	if c.ShouldBindJSON(&data) != nil {
		c.JSON(http.StatusOK, gin.H{"error": 9999})
		return
	}

	svc := new(userservice.UserService)
	svc.DeleteUser(data)
}
