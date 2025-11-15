package router

import (
	"task_manager/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	taskController := controllers.NewTaskController()
	r := gin.Default()

	tasks := r.Group("/tasks")
	{
		tasks.POST("", taskController.CreateTask)
		tasks.GET("", taskController.GetAllTasks)
		tasks.GET("/:id", taskController.GetTask)
		tasks.PUT("/:id", taskController.UpdateTask)
		tasks.DELETE("/:id", taskController.DeleteTask)
	}

	return r
}
