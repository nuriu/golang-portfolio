package repositories

import (
	"task-manager/internal/db/models"
	"task-manager/internal/domain/task"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TaskSqliteRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskSqliteRepository {
	db.AutoMigrate(&models.TaskEntity{})
	return &TaskSqliteRepository{db: db}
}

func (repository *TaskSqliteRepository) Create(task *task.Task) (*task.Task, error) {
	dbTask := models.TaskFromDomainEntity(task)

	if err := repository.db.Create(dbTask).Error; err != nil {
		return nil, err
	}

	task, err := repository.Get(dbTask.ID)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (repository *TaskSqliteRepository) List(page int, pageSize int) (interface{}, error) {
	var dbTasks []models.TaskEntity
	pagination := models.PaginatedModel{
		PageSize:    pageSize,
		CurrentPage: page,
		Order:       "created_at desc",
	}

	repository.db.Scopes(models.Paginate(dbTasks, &pagination, repository.db)).Find(&dbTasks)
	pagination.Data = dbTasks

	return pagination, nil
}

func (repository *TaskSqliteRepository) Get(id uuid.UUID) (*task.Task, error) {
	var dbTask models.TaskEntity
	if err := repository.db.First(&dbTask, id).Error; err != nil {
		return nil, err
	}

	task := dbTask.ToDomainEntity()
	return task, nil
}

func (repository *TaskSqliteRepository) Delete(id uuid.UUID) error {
	return repository.db.Delete(&models.TaskEntity{}, id).Error
}
