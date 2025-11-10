package main

import (
	"fmt"

	"library_management/controllers"
	"library_management/services"
)

func main() {
	lib := services.NewLibrary()
	ctrl := controllers.NewController(lib)
	fmt.Println("Welcome to the Library Management System!")
	fmt.Println("Preloaded members: 1) Alice, 2) Bob")
	ctrl.Run()
}


