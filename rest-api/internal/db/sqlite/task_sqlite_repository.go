package sqlite

import (
	"rest-api/internal/db/models"
	"rest-api/internal/domain/task"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TaskSqliteRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskSqliteRepository {
	db.AutoMigrate(&models.Task{})
	return &TaskSqliteRepository{db: db}
}

func (repository *TaskSqliteRepository) Create(task *task.Task) error {
	dbTask := models.MapToDBTask(task)

	if err := repository.db.Create(dbTask).Error; err != nil {
		return err
	}

	return nil
}

// TODO: paginate
func (repository *TaskSqliteRepository) List(page int, pageSize int) ([]*task.Task, error) {
	var dbTasks []models.Task

	if err := repository.db.Find(&dbTasks).Error; err != nil {
		return nil, err
	}

	tasks := make([]*task.Task, len(dbTasks))
	for i, dbTask := range dbTasks {
		tasks[i] = models.MapFromDBTask(&dbTask)
	}

	return tasks, nil
}

func (repository *TaskSqliteRepository) Get(id uuid.UUID) (*task.Task, error) {
	var dbTask models.Task
	if err := repository.db.First(&dbTask, id).Error; err != nil {
		return nil, err
	}

	task := models.MapFromDBTask(&dbTask)
	return task, nil
}

func (repository *TaskSqliteRepository) Delete(id uuid.UUID) error {
	return repository.db.Delete(&models.Task{}, id).Error
}
