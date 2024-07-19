package handlers

import (
	"fmt"
	"net/http"
	"rest-api/api/models"
	"rest-api/internal/domain/task"
	"strconv"

	"github.com/labstack/echo/v4"
)

type TaskHandler struct {
	service task.TaskService
}

func NewTaskHandlers(service task.TaskService) *TaskHandler {
	return &TaskHandler{service: service}
}

func (handler *TaskHandler) RegisterRoutes(group *echo.Group, routePrefix string) {
	group.GET(routePrefix, handler.listTasksHandler)
	group.POST(routePrefix, handler.createTaskHandler)
}

// @Summary List Tasks
// @Description Returns list of the created tasks
// @Accept json
// @Produce json
// @Param Page query int true "Page to retrieve"
// @Param PageSize query int true "Items per page"
// @Success 200 {array} task.Task
// @Failure 400
// @Failure 401
// @Failure 500
// @Router /api/v1/tasks [get]
func (handler *TaskHandler) listTasksHandler(c echo.Context) error {
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	pageSize, err := strconv.Atoi(c.QueryParam("pageSize"))
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	res := fmt.Sprintf("%d - %d", page, pageSize)

	return c.String(http.StatusOK, res)
}

// @Summary Create Task
// @Description Creates new task
// @Accept json
// @Produce json
// @Param Request body models.CreateTaskRequest true
// @Success 204
// @Failure 400 {string}
// @Failure 401
// @Failure 500
// @Router /api/v1/tasks [post]
func (handler *TaskHandler) createTaskHandler(c echo.Context) error {
	model := new(models.CreateTaskRequest)

	if err := c.Bind(model); err != nil {
		return c.String(http.StatusBadRequest, "Bad Request")
	}

	err := handler.service.CreateTask(model.Title, model.Description)
	if err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}
