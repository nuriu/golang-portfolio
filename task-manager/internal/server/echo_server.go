package server

import (
	"net/http"
	"task-manager/configs"
	"task-manager/internal/app/services"
	"task-manager/internal/db/sqlite"
	"task-manager/internal/http/handlers"

	echojwt "github.com/labstack/echo-jwt/v4"
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

	v1User := v1.Group("/users")
	userRepository := sqlite.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	userHandlers := handlers.NewUserHandler(userService)
	userHandlers.RegisterRoutes(v1User, "")

	v1Task := v1.Group("/tasks")
	v1Task.Use((echojwt.WithConfig(echojwt.Config{
		SigningKey:   []byte(configs.Environment.JWTSecret),
		ErrorHandler: handlers.HandleJWTError,
	})))
	taskRepository := sqlite.NewTaskRepository(db)
	taskService := services.NewTaskService(taskRepository)
	taskHandlers := handlers.NewTaskHandler(taskService)
	taskHandlers.RegisterRoutes(v1Task, "")

	e.Logger.Fatal(e.Start(address))
}
