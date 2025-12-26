package routes

import (
	"github.com/aliemreipek/go-task-master/pkg/controllers"
	"github.com/gin-gonic/gin"
)

// SetupRoutes configures the API endpoints
func SetupRoutes(r *gin.Engine) {
	// Group routes under /api/v1 for versioning
	api := r.Group("/api/v1")
	{
		api.POST("/tasks", controllers.CreateTask)       // Create
		api.GET("/tasks", controllers.GetAllTasks)       // Read All
		api.GET("/tasks/:id", controllers.GetTaskByID)   // Read One
		api.PUT("/tasks/:id", controllers.UpdateTask)    // Update
		api.DELETE("/tasks/:id", controllers.DeleteTask) // Delete
	}
}
