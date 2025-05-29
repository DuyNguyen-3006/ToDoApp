package main

import (
	"ToDoApp/Models"
	"ToDoApp/controller"
	"ToDoApp/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariable()
	initializers.ConnectDatabase()
}

func main() {
	r := gin.Default()
	
	initializers.DB.AutoMigrate(&Models.Post{})

	r.POST("/api/v1/post", controller.PostsCreate(initializers.DB))
	r.GET("/api/v1/post", controller.GetPosts)
	r.GET("/api/v1/post/:id", controller.GetPostById)
	r.PUT("/api/v1/post/:id", controller.UpdatePostById(initializers.DB))
	r.DELETE("/api/v1/post/:id", controller.DeletePostById)
	r.Run()
}
