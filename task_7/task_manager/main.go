package main

import (
	"log"
	"task_manager/router"
)

func main() {
	r := router.SetupRouter()

	log.Println("Server starting on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

