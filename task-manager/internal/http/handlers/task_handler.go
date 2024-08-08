package handlers

import (
	"net/http"
	"strconv"
	"task-manager/internal/domain/task"
	"task-manager/internal/http/models"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type TaskHandler struct {
	service task.TaskService
}

func NewTaskHandler(service task.TaskService) *TaskHandler {
	return &TaskHandler{service}
}

func (handler *TaskHandler) RegisterRoutes(group *echo.Group, routePrefix string) {
	group.GET(routePrefix, handler.listTasksHandler)
	group.POST(routePrefix, handler.createTaskHandler)
	group.GET(routePrefix+"/:id", handler.getTaskHandler)
	group.DELETE(routePrefix+"/:id", handler.deleteTaskHandler)
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
// @Success 201 {object} task.Task
// @Failure 400
// @Failure 401
// @Failure 500
func (handler *TaskHandler) createTaskHandler(c echo.Context) error {
	model := new(models.CreateTaskRequest)

	if err := c.Bind(model); err != nil {
		return c.String(http.StatusBadRequest, "Bad Request")
	}

	task, err := handler.service.CreateTask(model.Title, model.Description)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, task)
}

// @Router /api/v1/tasks/{id} [get]
// @Summary Get Task
// @Description Returns the task with given id
// @Accept json
// @Produce json
// @Param id path uuid.UUID true "Task ID"
// @Success 200 {object} task.Task
// @Failure 400
// @Failure 401
// @Failure 500
func (handler *TaskHandler) getTaskHandler(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	task, err := handler.service.GetTask(id)
	if err != nil {
		return c.String(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, task)
}

// @Router /api/v1/tasks/{id} [delete]
// @Summary Delete Task
// @Description Deletes the task with given id
// @Accept json
// @Produce json
// @Param id path uuid.UUID true "Task ID"
// @Success 204
// @Failure 400
// @Failure 401
// @Failure 500
func (handler *TaskHandler) deleteTaskHandler(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	err = handler.service.DeleteTask(id)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}
