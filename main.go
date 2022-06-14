package main

import (
	"net/http"

	"crud.com/crud/controllers"
	"crud.com/crud/docs"
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
	c := controllers.NewController()
	router.LoadHTMLGlob("webroot/*")

	router.GET("/ping", test)
	router.GET("/", homePage)

	router.POST("sso", c.Sso)

	v1 := router.Group("api/v1")
	{
		user := v1.Group("/user")
		user.Use(c.AuthRequired)
		{
			user.GET("/:id", c.GetUser)
			user.POST("/create", c.CreateUser)
			user.PUT("/update", c.UpdateUser)
			user.DELETE("/delete", c.DeleteUser)
		}
	}

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
