package repository

import (
	"github.com/Daniel-Handsome/2023-Backend-intern-Homework/model"
	"gorm.io/gorm"
)

type BaseRepo struct {
	orm *gorm.DB
}

func NewRepo(orm *gorm.DB, model model.BaseModel) *BaseRepo {
	return &BaseRepo{orm: orm}
}
