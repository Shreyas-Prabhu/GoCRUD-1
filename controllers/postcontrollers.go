package controllers

import (
	models "github.com/Shreyas-Prabhu/GoCRUD-1.git/Models"
	"github.com/Shreyas-Prabhu/GoCRUD-1.git/initializers"
	"github.com/gin-gonic/gin"
)

func GetPost(c *gin.Context) {

	var posts []models.Post
	initializers.DB.Find(&posts)

	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func GetSinglePost(c *gin.Context) {

	var post models.Post
	initializers.DB.First(&post, c.Param("id"))

	c.JSON(200, gin.H{
		"posts": post,
	})
}

func PostCreateInsert(c *gin.Context) {

	//Get data of request body

	var body struct {
		Body  string
		Title string
	}
	c.Bind(&body)
	posts := models.Post{Title: body.Title, Body: body.Body}
	result := initializers.DB.Create(&posts)
	if result.Error != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"Post": posts,
	})
}

func UpdatePosts(c *gin.Context) {

	//Get data of request body
	id := c.Param("id")

	var body struct {
		Body  string
		Title string
	}
	c.Bind(&body)

	var post models.Post
	initializers.DB.First(&post, id)

	initializers.DB.Model(&post).Updates(models.Post{Title: body.Title, Body: body.Body})

	c.JSON(200, gin.H{
		"Post": post,
	})
}

func DeletePost(c *gin.Context) {
	id := c.Param("id")

	initializers.DB.Delete(&models.Post{}, id)
	c.JSON(200, gin.H{
		"Post": "Deleted",
	})
}
