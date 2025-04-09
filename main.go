package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type requestBody struct {
	Task string `json:"task"`
}

var task string

func getTasks(c echo.Context) error {
	return c.JSON(http.StatusOK, "Hello "+task)
}

func postTasks(c echo.Context) error {
	var req requestBody
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}
	task = req.Task
	return c.JSON(http.StatusCreated, req)
}

func main() {
	e := echo.New()

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	e.GET("/", getTasks)
	e.POST("/", postTasks)

	e.Start("localhost:8080")
}
