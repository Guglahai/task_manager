package taskService

type Task struct {
	ID      string `gorm:"primaryKey" json:"id"`
	Task    string `json:"task"`
	Is_done bool   `json:"is_done"`
}

type TaskRequest struct {
	Task    string `json:"task"`
	Is_done bool   `json:"is_done"`
}
