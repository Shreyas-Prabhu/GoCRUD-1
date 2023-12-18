package main

import (
	models "github.com/Shreyas-Prabhu/GoCRUD-1.git/Models"
	"github.com/Shreyas-Prabhu/GoCRUD-1.git/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
