package repository

import "gorm.io/gorm"

type BaseRepo struct {
	orm *gorm.DB
}

func NewRepo(orm *gorm.DB) *BaseRepo {
	return &BaseRepo{orm: orm}
}

func (r *BaseRepo) Transaction(execute func(tx *gorm.DB) error) error {
	return r.orm.Transaction(execute)
}

func (r *BaseRepo) Begin() *gorm.DB {
	return r.orm.Begin()
}
