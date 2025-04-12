package main

import (
	"log"
	"task_manager/internal/db"
	"task_manager/internal/handlers"
	"task_manager/internal/taskService"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	database, err := db.InitDB()
	if err != nil {
		log.Fatalf("Could not connect to database: %s", err)
	}

	e := echo.New()

	taskRepo := taskService.NewTaskRepository(database)
	taskServices := taskService.NewTaskService(taskRepo)
	taskHandlers := handlers.NewTaskHandler(taskServices)

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	e.GET("/tasks", taskHandlers.GetTasks)
	e.POST("/tasks", taskHandlers.PostTasks)
	e.PATCH("/tasks/:id", taskHandlers.PatchTasks)
	e.DELETE("/tasks/:id", taskHandlers.DeleteTasks)

	e.Start("localhost:8080")
}
