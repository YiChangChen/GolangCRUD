package main

import (
	"net/http"
	"strconv"

	"crud.com/crud/docs"
	"crud.com/crud/models"
	userservice "crud.com/crud/services/userService"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func main() {
	docs.SwaggerInfo.Title = "Swagger Example API"
	docs.SwaggerInfo.Description = "This is a sample server Petstore server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	router := gin.Default()
	router.LoadHTMLGlob("webroot/*")

	router.GET("/ping", test)
	router.GET("/", homePage)

	router.Group("/api/v1/user").
		GET("/:id", GetUser).
		POST("/create", CreateUser).
		PUT("/update", UpdateUser).
		DELETE("/delete", DeleteUser)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(":8080")
}

func test(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func homePage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

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
func GetUser(c *gin.Context) {
	id := c.Param("id")
	_, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": 9999})
		return
	}

	svc := new(userservice.UserService)
	result, err := svc.GetUser(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": 9999})
		return
	}
	c.JSON(http.StatusOK, result)
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
func CreateUser(c *gin.Context) {
	var data = new(models.User)

	if c.ShouldBindJSON(&data) != nil {
		c.JSON(http.StatusOK, gin.H{"error": 9999})
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
func UpdateUser(c *gin.Context) {
	var data = new(models.User)

	if c.ShouldBindJSON(&data) != nil {
		c.JSON(http.StatusOK, gin.H{"error": 9999})
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
func DeleteUser(c *gin.Context) {
	var data = new(models.User)

	if c.ShouldBindJSON(&data) != nil {
		c.JSON(http.StatusOK, gin.H{"error": 9999})
		return
	}

	svc := new(userservice.UserService)
	svc.DeleteUser(data)
}
