package repository

import (
	"context"
	"errors"

	"github.com/Daniel-Handsome/2023-Backend-intern-Homework/model"
	"gorm.io/gorm"
)

type PageLinkedListRepository interface {
	GetByHeadKey(ctx context.Context, headKey string) (pageLinkedList model.PageLinkedList, err error)
	GetByType(ctx context.Context, t model.PageLinkedListType) (pageLinkedLists []model.PageLinkedList, err error)
}

type pageLinkedListRepository struct {
	*BaseRepo
	model *model.PageLinkedList
}

func NewPageLinkedListRepository(orm *gorm.DB) PageLinkedListRepository {
	return &pageLinkedListRepository{
		BaseRepo: &BaseRepo{
			orm: orm,
		},
		model: &model.PageLinkedList{},
	}
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

func (repo pageLinkedListRepository) GetByType(ctx context.Context, t model.PageLinkedListType) (pageLinkedLists []model.PageLinkedList, err error) {
	err = repo.orm.WithContext(ctx).
		Model(model.PageLinkedList{}).
		Where("type = ?", t).
		Find(&pageLinkedLists).
		Error
	if err != nil {
		return pageLinkedLists, err
	}

	return
}
