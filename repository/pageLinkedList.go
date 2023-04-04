package repository

import (
	"context"
	"errors"

	"github.com/Daniel-Handsome/2023-Backend-intern-Homework/model"
	"gorm.io/gorm"
)

type PageLinkedListRepository interface {
	GetByHeadKey(ctx context.Context, headKey string) (pageLinkedList model.PageLinkedList, err error)
}

type pageLinkedListRepository struct {
	*BaseRepo
}

func NewPageLinkedListRepository(orm *gorm.DB) PageLinkedListRepository {
	return &pageLinkedListRepository{NewRepo(orm.Model(model.PageLinkedList{}))}
}

func (repo pageLinkedListRepository) GetByHeadKey(ctx context.Context, headKey string) (pageLinkedList model.PageLinkedList, err error) {
	query := repo.orm.WithContext(ctx).
		Where("head = ?", headKey)

	if err := query.First(&pageLinkedList).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return pageLinkedList, ErrUserNotFound
		}
	}

	return
}