package task

type TaskService interface {
	CreateTask(title string, description string) (*Task, error)
}
