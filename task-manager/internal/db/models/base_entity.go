package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseEntity struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

func (baseEntity *BaseEntity) BeforeCreate(db *gorm.DB) (err error) {
	baseEntity.ID, err = uuid.NewV7()
	return
}
