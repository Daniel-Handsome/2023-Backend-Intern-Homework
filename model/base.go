package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseModel struct {
	gorm.Model
	Uuid string
}

func (b *BaseModel) BeforeCreate(tx *gorm.DB) error {
	b.Uuid = uuid.New().String()
	return nil
}
