package server

import (
	"net/http"
	"task-manager/internal/app/services"
	"task-manager/internal/db/sqlite"
	"task-manager/internal/http/handlers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"gorm.io/gorm"
)

func Run(address string, db *gorm.DB) {
	e := echo.New()

	e.HTTPErrorHandler = handlers.HandleError

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	v1 := e.Group("/api/v1")
	// v1.Use((echojwt.WithConfig(echojwt.Config{
	// 	SigningKey:   []byte(configs.Environment.JWTSecret),
	// 	ErrorHandler: handlers.HandleJWTError,
	// })))

	taskRepository := sqlite.NewTaskRepository(db)
	taskService := services.NewTaskService(taskRepository)
	taskHandlers := handlers.NewTaskHandlers(taskService)
	taskHandlers.RegisterRoutes(v1, "/tasks")

	e.Logger.Fatal(e.Start(address))
}
