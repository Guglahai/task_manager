package main

import (
	"net/http"
	"slices"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Task struct {
	ID   string `json:"id"`
	Task string `json:"task"`
}

var tasks = []Task{}

func getTasks(c echo.Context) error {
	return c.JSON(http.StatusOK, tasks)
}

func postTasks(c echo.Context) error {
	var req Task
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}
	task := Task{
		ID:   uuid.NewString(),
		Task: req.Task,
	}
	tasks = append(tasks, task)
	return c.JSON(http.StatusCreated, tasks)
}

func patchTasks(c echo.Context) error {
	id := c.Param("id")

	var req Task
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Task = req.Task
			return c.JSON(http.StatusOK, tasks[i])
		}
	}
	return c.JSON(http.StatusBadRequest, map[string]string{"error": "Task not found"})
}

func deleteTasks(c echo.Context) error {
	id := c.Param("id")

	for i, task := range tasks {
		if task.ID == id {
			tasks = slices.Delete(tasks, i, i+1)
			return c.NoContent(http.StatusNoContent)
		}
	}
	return c.JSON(http.StatusBadRequest, id)
}

func main() {
	e := echo.New()

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	e.GET("/", getTasks)
	e.POST("/", postTasks)
	e.PATCH("/", patchTasks)
	e.DELETE("/:id", deleteTasks)

	e.Start("localhost:8080")
}
