package main

import (
	"log"
	"task_manager/internal/db"
	"task_manager/internal/handlers"
	"task_manager/internal/taskService"
	"task_manager/internal/userService"
	"task_manager/internal/web/tasks"
	"task_manager/internal/web/users"

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

	userRepo := userService.NewUserRepository(database)
	userServices := userService.NewUserService(userRepo)
	userHandlers := handlers.NewUserHandler(userServices)

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	strictHandler := tasks.NewStrictHandler(taskHandlers, nil)
	tasks.RegisterHandlers(e, strictHandler)

	usersStrictHandler := users.NewStrictHandler(userHandlers, nil)
	users.RegisterHandlers(e, usersStrictHandler)

	if err := e.Start("localhost:8080"); err != nil {
		log.Fatalf("failed to start with err: %v", err)
	}
}
