package main

import (
	"ToDoApp/Models"
	"ToDoApp/initializers"
)

func init() {
	initializers.LoadEnvVariable()
	initializers.ConnectDatabase()
}
func main() {
	initializers.DB.AutoMigrate(&Models.Post{})
}
