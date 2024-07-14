package handlers

import (
	"net/http"
	"rest-api/api/models"
	"rest-api/internal/domain"
	"rest-api/internal/domain/task"

	"github.com/labstack/echo/v4"
)

type TaskHandler struct {
	service task.TaskService
}

func NewTaskHandlers(service task.TaskService) *TaskHandler {
	return &TaskHandler{service: service}
}

func (handler *TaskHandler) RegisterRoutes(group *echo.Group, routePrefix string) {
	group.POST(routePrefix, handler.createTaskHandler)
}

// @Summary Create Task
// @Description Create new task
// @Accept json
// @Produce json
// @Param Request body models.CreateTaskRequest true "CreateTaskRequest"
// @Success 200 {object} task.Task
// @Router /api/v1/tasks [post]
func (handler *TaskHandler) createTaskHandler(c echo.Context) error {
	model := new(models.CreateTaskRequest)

	if err := c.Bind(model); err != nil {
		return c.String(http.StatusBadRequest, "Bad Request")
	}

	task, err := handler.service.CreateTask(model.Title, model.Description)
	if err != nil {
		if domainErr, ok := err.(*domain.DomainError); ok {
			return c.String(domainErr.Code, domainErr.Message)
		}
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, task)
}
