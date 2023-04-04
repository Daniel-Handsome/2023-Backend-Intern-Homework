package repository

import (
	"context"

	"github.com/Daniel-Handsome/2023-Backend-intern-Homework/model"
	"gorm.io/gorm"
)

type PageNodeRepository interface {
	GetByUuid(ctx context.Context, uuid string) (node model.PageNode, err error)
}

type pageNodeRepository struct {
	*BaseRepo
}

func NewPageNodeRepository(orm *gorm.DB) PageNodeRepository {
	return &pageNodeRepository{NewRepo(orm.Model(model.PageNode{}))}
}

func (repo pageNodeRepository) GetByUuid(ctx context.Context, uuid string) (node model.PageNode, err error) {
	query := repo.orm.WithContext(ctx).Where("uuid = ?", uuid)

	if err := query.First(&node).Error; err != nil {
		return node, err
	}

	return
}
