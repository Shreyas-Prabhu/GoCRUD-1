package main

import (
	"github.com/Shreyas-Prabhu/GoCRUD-1.git/controllers"
	"github.com/Shreyas-Prabhu/GoCRUD-1.git/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}
func main() {
	r := gin.Default()
	r.GET("/posts", controllers.GetPost)
	r.GET("/post/:id", controllers.GetSinglePost)
	r.POST("/posts", controllers.PostCreateInsert)
	r.PUT("/post/:id", controllers.UpdatePosts)
	r.DELETE("/post/:id", controllers.DeletePost)
	r.Run()
}
