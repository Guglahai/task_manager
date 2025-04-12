package main

import (
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func initDB() {
	dsn := "host=localhost user=postgres password=1103 dbname=postgres port=5432 sslmode=disable"
	var err error

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Could not connect to database: %s", err)
	}

	if err := db.AutoMigrate(&Task{}); err != nil {
		log.Fatalf("Could not migrate: %s", err)
	}
}

type Task struct {
	ID      string `gorm:"primaryKey" json:"id"`
	Task    string `json:"task"`
	Is_done bool   `json:"is_done"`
}

type TaskRequest struct {
	Task    string `json:"task"`
	Is_done bool   `json:"is_done"`
}

// Find, Create, Update, Delete

func getTasks(c echo.Context) error {
	var tasks []Task

	if err := db.Find(&tasks).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not get tasks"})
	}

	return c.JSON(http.StatusOK, tasks)
}

func postTasks(c echo.Context) error {
	var req TaskRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}
	task := Task{
		ID:      uuid.NewString(),
		Task:    req.Task,
		Is_done: req.Is_done,
	}

	if err := db.Create(&task).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not add tasks"})
	}

	return c.JSON(http.StatusCreated, task)
}

func patchTasks(c echo.Context) error {
	id := c.Param("id")

	var req Task
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	var task Task
	if err := db.First(&task, "id = ?", id).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Could not find task"})
	}

	task.Task = req.Task
	task.Is_done = req.Is_done

	if err := db.Save(&task).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not update tasks"})
	}

	return c.JSON(http.StatusOK, task)
}

func deleteTasks(c echo.Context) error {
	id := c.Param("id")

	if err := db.Delete(&Task{}, "id = ?", id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not delete tasks"})
	}

	return c.NoContent(http.StatusNoContent)
}

func main() {
	initDB()

	e := echo.New()

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	e.GET("/tasks", getTasks)
	e.POST("/tasks", postTasks)
	e.PATCH("/tasks/:id", patchTasks)
	e.DELETE("/tasks/:id", deleteTasks)

	e.Start("localhost:8080")
}
