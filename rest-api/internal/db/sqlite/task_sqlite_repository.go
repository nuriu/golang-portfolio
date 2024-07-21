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

func (repository *TaskSqliteRepository) List(page int, pageSize int) (interface{}, error) {
	var dbTasks []models.Task
	pagination := models.PaginatedModel{
		PageSize:    pageSize,
		CurrentPage: page,
		Order:       "CreatedAt desc",
	}

	repository.db.Scopes(models.Paginate(dbTasks, &pagination, repository.db)).Find(&dbTasks)
	pagination.Data = dbTasks

	return pagination, nil
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
