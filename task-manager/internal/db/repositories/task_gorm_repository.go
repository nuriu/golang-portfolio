package repositories

import (
	"task-manager/internal/db/models"
	"task-manager/internal/domain/task"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TaskGormRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskGormRepository {
	db.AutoMigrate(&models.TaskEntity{})
	return &TaskGormRepository{db: db}
}

func (repository *TaskGormRepository) Create(task *task.Task) (*task.Task, error) {
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

func (repository *TaskGormRepository) List(page int, pageSize int) (interface{}, error) {
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

func (repository *TaskGormRepository) Get(id uuid.UUID) (*task.Task, error) {
	var dbTask models.TaskEntity
	if err := repository.db.First(&dbTask, id).Error; err != nil {
		return nil, err
	}

	task := dbTask.ToDomainEntity()
	return task, nil
}

func (repository *TaskGormRepository) Delete(id uuid.UUID) error {
	return repository.db.Delete(&models.TaskEntity{}, id).Error
}

func (repository *TaskGormRepository) Update(id uuid.UUID, task *task.Task) {
	dbModel := models.TaskFromDomainEntity(task)
	repository.db.Save(&dbModel)
}
