package controllers

import (
	"net/http"

	"github.com/aliemreipek/go-task-master/pkg/config"
	"github.com/aliemreipek/go-task-master/pkg/models"
	"github.com/gin-gonic/gin"
)

// CreateTask handles the creation of a new task
// POST /tasks
func CreateTask(c *gin.Context) {
	var task models.Task

	// Bind the incoming JSON body to the Task struct
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload", "details": err.Error()})
		return
	}

	// Save the task to the database
	if err := config.DB.Create(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
		return
	}

	c.JSON(http.StatusCreated, task)
}

// GetAllTasks retrieves all tasks from the database
// GET /tasks
func GetAllTasks(c *gin.Context) {
	var tasks []models.Task

	// Fetch all tasks using GORM
	if err := config.DB.Find(&tasks).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tasks"})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

// GetTaskByID retrieves a specific task by its ID
// GET /tasks/:id
func GetTaskByID(c *gin.Context) {
	id := c.Param("id")
	var task models.Task

	// Find the task by ID
	if err := config.DB.First(&task, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	c.JSON(http.StatusOK, task)
}

// UpdateTask updates an existing task
// PUT /tasks/:id
func UpdateTask(c *gin.Context) {
	id := c.Param("id")
	var task models.Task

	// Check if the task exists
	if err := config.DB.First(&task, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	var input models.Task
	// Bind JSON to a new struct to avoid zero-value overwrites if needed,
	// but for simplicity, we bind to the same struct here or a DTO.
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Update fields
	config.DB.Model(&task).Updates(input)

	c.JSON(http.StatusOK, task)
}

// DeleteTask removes a task from the database
// DELETE /tasks/:id
func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	var task models.Task

	// Check if task exists before deleting
	if err := config.DB.First(&task, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	// Perform soft delete
	config.DB.Delete(&task)

	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}
