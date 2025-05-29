package controller

import (
	"ToDoApp/DTOs/Post"
	"ToDoApp/Models"
	"ToDoApp/initializers"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"net/http"
	"time"
)

var validate = validator.New()

func PostsCreate(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req Post.PostRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Error Data: " + err.Error()})
			return
		}
		if err := validate.Struct(req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "You must fill all fields"})
			return
		}

		if !req.Status.IsValid() {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Status must be: ToDo, InProgress, Done"})
			return
		}

		post := Models.Post{
			Title:  req.Title,
			Body:   req.Body,
			Name:   req.Name,
			Status: req.Status,
			Time:   time.Now(),
		}

		if err := db.Create(&post).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Không thể tạo bài viết: " + err.Error()})
			return
		}

		response := Post.PostResponse{
			ID:        post.ID,
			Title:     post.Title,
			Body:      post.Body,
			Name:      post.Name,
			Status:    post.Status,
			Time:      post.Time,
			CreatedAt: post.CreatedAt,
		}

		c.JSON(http.StatusCreated, response)
	}
}

func GetPosts(c *gin.Context) {
	var posts []Models.Post
	if err := initializers.DB.Find(&posts).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "post not found"})
		return
	}
	response := make([]Post.PostResponse, len(posts))
	for i, post := range posts {
		response[i] = Post.PostResponse{
			ID:        post.ID,
			Title:     post.Title,
			Body:      post.Body,
			Name:      post.Name,
			Status:    post.Status,
			Time:      post.Time,
			CreatedAt: post.CreatedAt,
		}
	}
	c.JSON(http.StatusOK, response)
}
func GetPostById(c *gin.Context) {
	var post Models.Post
	id := c.Param("id")
	result := initializers.DB.First(&post, id)
	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Post Not Found"})
		return
	}
	var response Post.PostResponse
	response.ID = post.ID
	response.Title = post.Title
	response.Body = post.Body
	response.Name = post.Name
	response.Status = post.Status
	response.Time = post.Time
	response.CreatedAt = post.CreatedAt
	c.JSON(http.StatusOK, response)
}

func UpdatePostById(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var req Post.PostRequest
		err := c.Bind(&req)
		if err != nil {
			return
		}
		if !req.Status.IsValid() {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Trạng thái không hợp lệ: " + string(req.Status)})
			return
		}
		var post Models.Post
		result := initializers.DB.First(&post, id)
		if result.Error != nil {
			c.JSON(404, gin.H{"error": "Post Not Found"})
			return
		}
		updatedPost := Models.Post{
			Title:  post.Title,
			Body:   post.Body,
			Name:   post.Name,
			Status: post.Status,
		}
		if err := db.Model(&post).Updates(updatedPost).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Fail to update: " + err.Error()})
			return
		}
		response := Post.UpdatePostResponse{
			ID:        post.ID,
			Title:     post.Title,
			Body:      post.Body,
			Name:      post.Name,
			Status:    post.Status,
			UpdatedAt: post.UpdatedAt,
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "update success",
			"data":    response,
		})

	}

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
