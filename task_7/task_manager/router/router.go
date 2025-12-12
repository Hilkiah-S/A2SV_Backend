package router

import (
	"task_manager/controllers"
	"task_manager/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	taskController := controllers.NewTaskController()
	authController := controllers.NewAuthController()
	r := gin.Default()

	auth := r.Group("/auth")
	{
		auth.POST("/register", authController.Register)
		auth.POST("/login", authController.Login)
	}

	admin := r.Group("/admin")
	admin.Use(middleware.AuthMiddleware())
	admin.Use(middleware.AdminMiddleware())
	{
		admin.POST("/promote", authController.Promote)
	}

	tasks := r.Group("/tasks")
	tasks.Use(middleware.AuthMiddleware())
	{
		tasks.GET("", taskController.GetAllTasks)
		tasks.GET("/:id", taskController.GetTask)
	}

	adminTasks := r.Group("/tasks")
	adminTasks.Use(middleware.AuthMiddleware())
	adminTasks.Use(middleware.AdminMiddleware())
	{
		adminTasks.POST("", taskController.CreateTask)
		adminTasks.PUT("/:id", taskController.UpdateTask)
		adminTasks.DELETE("/:id", taskController.DeleteTask)
	}

	return r
}

