package controller

import (
	"ToDoApp/Models"
	"ToDoApp/initializers"
	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {

	var body struct {
		Body  string
		Title string
	}
	err := c.Bind(&body)
	if err != nil {
		return
	}
	post := Models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error})
	} else {
		c.JSON(200, gin.H{"data": post})
	}
}

func GetPosts(c *gin.Context) {
	var posts []Models.Post
	initializers.DB.Find(&posts)
	if posts == nil {
		c.JSON(404, gin.H{"error": "Not Found"})
		return
	}
	c.JSON(200, gin.H{"data": posts})
}
func GetPostById(c *gin.Context) {
	var post Models.Post
	id := c.Param("id")
	result := initializers.DB.First(&post, id)
	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Post Not Found"})
		return
	}
	c.JSON(200, gin.H{"data": post})
}

func UpdatePostById(c *gin.Context) {
	//Get Id from url
	id := c.Param("id")
	var body struct {
		Body  string
		Title string
	}
	err := c.Bind(&body)
	if err != nil {
		return
	}
	var post Models.Post
	result := initializers.DB.First(&post, id)
	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Post Not Found"})
		return
	}
	//Update Post
	initializers.DB.Model(&post).Updates(Models.Post{Title: body.Title, Body: body.Body})
	c.JSON(200, gin.H{"Post updated: ": post})
}

func DeletePostById(c *gin.Context) {
	id := c.Param("id")
	var post Models.Post
	result := initializers.DB.First(&post, id)
	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Post Not Found"})
		return
	}
	initializers.DB.Delete(&post)
	c.JSON(200, gin.H{"Post deleted: ": post})
}
