package handlers

import (
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

// @Router /api/v1/tasks [get]
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
func (handler *TaskHandler) listTasksHandler(c echo.Context) error {
	page, err := strconv.Atoi(c.QueryParam("Page"))
	if err != nil {
		return c.String(http.StatusBadRequest, "Couldn't parse 'Page' query parameter.")
	}

	pageSize, err := strconv.Atoi(c.QueryParam("PageSize"))
	if err != nil {
		return c.String(http.StatusBadRequest, "Couldn't parse 'PageSize' query parameter.")
	}

	tasks, err := handler.service.ListTasks(page, pageSize)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, tasks)
}

// @Router /api/v1/tasks [post]
// @Summary Create Task
// @Description Creates new task
// @Accept json
// @Produce json
// @Param Request body models.CreateTaskRequest true "title and description for the new task"
// @Success 204
// @Failure 400
// @Failure 401
// @Failure 500
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
