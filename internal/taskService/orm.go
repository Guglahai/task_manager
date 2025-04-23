package taskService

import "github.com/google/uuid"

type Task struct {
	ID      uuid.UUID `gorm:"primaryKey" json:"id"`
	Task    string    `json:"task"`
	Is_done bool      `json:"is_done"`
}

type TaskRequest struct {
	Task    string `json:"task"`
	Is_done bool   `json:"is_done"`
}
