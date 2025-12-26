package models

import (
	"gorm.io/gorm"
	"time"
)

// Task struct comes across the 'tasks' table in the database.
type Task struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Title       string         `gorm:"size:255;not null" json:"title"`
	Description string         `gorm:"type:text" json:"description"`
	Status      string         `gorm:"default:'pending'" json:"status"` // pending, in-progress, completed
	DueDate     time.Time      `json:"due_date"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"` // Soft delete özelliği
}
