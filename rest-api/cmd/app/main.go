package main

import (
	"rest-api/api/handlers"
	"rest-api/configs"
	_ "rest-api/docs"
	"rest-api/internal/app/services"
	"rest-api/internal/db/postgres"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title REST API
// @version 1.0
// @description REST API documentation.
// @host localhost:8080
func main() {
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	v1 := e.Group("/api/v1")

	taskRepository := postgres.NewTaskRepository()
	taskService := services.NewTaskService(taskRepository)
	taskHandlers := handlers.NewTaskHandlers(taskService)
	taskHandlers.RegisterRoutes(v1, "/tasks")

	addr := configs.Environment.Host + ":" + configs.Environment.Port
	e.Logger.Fatal(e.Start(addr))
}
