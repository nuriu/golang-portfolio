package sqlite

import (
	"task-manager/internal/db/models"
	"task-manager/internal/domain/user"

	"gorm.io/gorm"
)

type UserSqliteRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserSqliteRepository {
	db.AutoMigrate(&models.UserEntity{})
	return &UserSqliteRepository{db: db}
}

func (repository *UserSqliteRepository) Create(user *user.User) (*user.User, error) {
	dbUser := models.UserFromDomainEntity(user)

	if err := repository.db.Create(dbUser).Error; err != nil {
		return nil, err
	}

	user, err := repository.Get(dbUser.Email)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (repository *UserSqliteRepository) Get(email string) (*user.User, error) {
	var dbUser models.UserEntity
	if err := repository.db.First(&dbUser, email).Error; err != nill {
		return nil, err
	}

	user := dbUser.ToDomainEntity()
	return user, nil
}
