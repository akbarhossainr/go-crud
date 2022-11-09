package controllers

import (
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
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"post": post,
	})
}

func GetPosts(c *gin.Context) {
	var posts []models.Post
	result := initializers.DB.Find(&posts)

	if result.Error != nil {
		c.JSON(400, gin.H{
			"message": "No data found",
		})
		return
	}

	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func GetPost(c *gin.Context) {
	id := c.Param("id")
	var post models.Post

	result := initializers.DB.First(&post, id)

	if result.Error != nil {
		c.JSON(400, gin.H{
			"message": "No data found",
		})
		return
	}

	c.JSON(200, gin.H{
		"post": post,
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
	initializers.DB.Model(&post).Updates(models.Post{
		Title: data.Title,
		Body:  data.Body,
	})

	// return the response
	c.JSON(200, gin.H{
		"post": post,
	})
}

func DeletePost(c *gin.Context) {
	id := c.Param("id")

	initializers.DB.Delete(&models.Post{}, id)

	c.JSON(200, gin.H{
		"message": "Post deleted successfully",
	})
}
