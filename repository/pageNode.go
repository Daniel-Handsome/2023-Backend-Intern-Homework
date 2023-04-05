package repository

import (
	"context"
	"errors"

	"github.com/Daniel-Handsome/2023-Backend-intern-Homework/model"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

var ErrPageNodeNotFound = errors.New("page node not found")

type PageNodeRepository interface {
	GetByUuid(ctx context.Context, uuid string) (node model.PageNode, err error)
	GetArticlesIdsAndNextByUuid(ctx context.Context, uuid string) (node model.PageNode, err error)
	UpdateArticleIds(ctx context.Context, uuid string, articleIds []int64) (node model.PageNode, err error)
}

type pageNodeRepository struct {
	*BaseRepo
}

func NewPageNodeRepository(orm *gorm.DB) PageNodeRepository {
	return &pageNodeRepository{NewRepo(orm.Model(model.PageNode{}))}
}

func (repo pageNodeRepository) GetByUuid(ctx context.Context, uuid string) (model.PageNode, error) {
	var node model.PageNode
	query := repo.orm.WithContext(ctx).Where("uuid = ?", uuid)

	if err := query.First(&node).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return node, ErrPageNodeNotFound
		}
		return node, err
	}

	return node, nil
}

func (repo pageNodeRepository) GetArticlesIdsAndNextByUuid(ctx context.Context, uuid string) (node model.PageNode, err error) {
	err = repo.orm.WithContext(ctx).
		Select("next", "article_ids").
		Where("uuid = ?", uuid).
		First(&node).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return node, ErrUserNotFound
		}
		return node, err
	}

	return
}

func (repo pageNodeRepository) UpdateArticleIds(ctx context.Context, uuid string, articleIds []int64) (model.PageNode, error) {
	var pageNode model.PageNode
	err := repo.orm.Model(&model.PageNode{}).
		WithContext(ctx).
		Where("uuid = ?", uuid).Updates(
		map[string]interface{}{
			"article_ids": pq.Array(articleIds),
		}).First(&pageNode).Error
	if err != nil {
		return pageNode, err
	}
	return pageNode, nil
}
