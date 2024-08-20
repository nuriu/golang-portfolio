package repositories

import (
	"task-manager/internal/db/models"
	"task-manager/internal/domain/user"

	"gorm.io/gorm"
)

type UserGormRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserGormRepository {
	db.AutoMigrate(&models.UserEntity{})
	return &UserGormRepository{db: db}
}

func (repository *UserGormRepository) Create(user *user.User) (*user.User, error) {
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

func (repository *UserGormRepository) Get(email string) (*user.User, error) {
	var dbUser models.UserEntity
	if err := repository.db.Where(&user.User{Email: email}).First(&dbUser).Error; err != nil {
		return nil, err
	}

	user := dbUser.ToDomainEntity()
	return user, nil
}
