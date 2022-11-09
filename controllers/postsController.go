package controllers

import (
	"net/http"

	"github.com/akbarhossainr/go-crud/initializers"
	"github.com/akbarhossainr/go-crud/models"
	"github.com/gin-gonic/gin"
)

func CreatePosts(c *gin.Context) {
	var req struct {
		Title string
		Body  string
	}

	c.Bind(&req)

	post := models.Post{Title: req.Title, Body: req.Body}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "Something went wrong!",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"error":   false,
		"message": "Post created successfully!",
	})
}

func GetPosts(c *gin.Context) {
	var posts []models.Post
	result := initializers.DB.Find(&posts)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "No record found!",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error": false,
		"data":  posts,
	})
}

func GetPost(c *gin.Context) {
	id := c.Param("id")
	var post models.Post

	result := initializers.DB.First(&post, id)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "No record found!",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error": false,
		"data":  post,
	})
}

func UpdatePost(c *gin.Context) {
	// get request body
	id := c.Param("id")

	var data struct {
		Title string
		Body  string
	}

	c.Bind(&data)

	// find post by id
	var post models.Post
	initializers.DB.First(&post, id)

	// update the post
	result := initializers.DB.Model(&post).Updates(models.Post{
		Title: data.Title,
		Body:  data.Body,
	})

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "Something went wrong!",
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"error":   false,
		"message": "Post updated successfully!",
	})
}

func DeletePost(c *gin.Context) {
	id := c.Param("id")

	result := initializers.DB.Delete(&models.Post{}, id)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "Something went wrong!",
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"error":   false,
		"message": "Post deleted successfully!",
	})
}
