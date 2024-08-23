package models

import (
	"task-manager/internal/domain/user"
	"time"

	"gorm.io/gorm"
)

type UserEntity struct {
	BaseEntity
	Email    string
	Password string
}

func (UserEntity) TableName() string {
	return "users"
}

func UserFromDomainEntity(u *user.User) *UserEntity {
	deleteInfo := gorm.DeletedAt{Valid: false}
	if u.DeletedAt != nil {
		deleteInfo.Time = *u.DeletedAt
		deleteInfo.Valid = true
	}

	return &UserEntity{
		BaseEntity: BaseEntity{
			ID:        u.ID,
			CreatedAt: u.CreatedAt,
			DeletedAt: deleteInfo,
		},
		Email:    u.Email,
		Password: u.Password,
	}
}

func (ue *UserEntity) ToDomainEntity() *user.User {
	var deletedAt *time.Time = nil
	if ue.DeletedAt.Valid {
		deletedAt = &ue.DeletedAt.Time
	}

	return &user.User{
		ID:        ue.ID,
		Email:     ue.Email,
		Password:  ue.Password,
		CreatedAt: ue.CreatedAt,
		UpdatedAt: ue.UpdatedAt,
		DeletedAt: deletedAt,
	}
}
