package userService

import (
	"task_manager/internal/taskService"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"primaryKey" json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Tasks     []taskService.Task
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
