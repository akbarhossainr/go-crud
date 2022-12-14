package main

import (
	"github.com/akbarhossainr/go-crud/controllers"
	"github.com/akbarhossainr/go-crud/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/posts", controllers.GetPosts)
	r.POST("/posts", controllers.CreatePosts)
	r.GET("/posts/:id", controllers.GetPost)
	r.PUT("/posts/:id", controllers.UpdatePost)
	r.DELETE("/posts/:id", controllers.DeletePost)

	r.Run() // listen and serve on 0.0.0.0:8080
}
