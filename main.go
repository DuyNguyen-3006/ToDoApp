package main

import (
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

	r.POST("/api/v1/post", controller.PostsCreate)
	r.GET("/api/v1/post", controller.GetPosts)
	r.GET("/api/v1/post/:id", controller.GetPostById)
	r.PUT("/api/v1/post/:id", controller.UpdatePostById)
	r.DELETE("/api/v1/post/:id", controller.DeletePostById)
	r.Run()
}
