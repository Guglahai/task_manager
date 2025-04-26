package main

import (
	"log"
	"task_manager/internal/db"
	"task_manager/internal/handlers"
	"task_manager/internal/taskService"
	"task_manager/internal/web/tasks"

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

	strictHandler := tasks.NewStrictHandler(taskHandlers, nil)
	tasks.RegisterHandlers(e, strictHandler)

	if err := e.Start("localhost:8080"); err != nil {
		log.Fatalf("failed to start with err: %v", err)
	}
}
