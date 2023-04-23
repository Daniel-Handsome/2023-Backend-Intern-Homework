package repository

import (
	"context"
	"errors"

	"github.com/Daniel-Handsome/2023-Backend-intern-Homework/model"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

var ErrArticleNotFound = errors.New("article not found")

type ArticleRepository interface {
	GetIdsSortByOmitId(ctx context.Context, ids []int64, column model.OrderColumn) (newIds []int64, err error)
	GetPage(ctx context.Context, ids []int64) (articles []model.Article, err error)
	UpdateArticlesPage(ctx context.Context) error
	Transaction(execute func(tx *gorm.DB) error) error
}

type articleRepository struct {
	*BaseRepo
	model *model.Article
}

func NewArticleRepository(orm *gorm.DB) ArticleRepository {
	return &articleRepository{
		BaseRepo: &BaseRepo{
			orm: orm,
		},
		model: &model.Article{},
	}
}

func (repo *articleRepository) GetPage(ctx context.Context, ids []int64) (articles []model.Article, err error) {
	err = repo.orm.WithContext(ctx).
		Where("id = ANY(?)", pq.Array(ids)).
		Order("sort").
		Find(&articles).Error
	if err != nil {
		return
	}
	return
}

func (repo articleRepository) GetIdsSortByOmitId(ctx context.Context, ids []int64, column model.OrderColumn) (newIds []int64, err error) {
	columName, err := column.GetName()
	if err != nil {
		return
	}
	err = repo.orm.WithContext(ctx).
		Select("id").
		Where("id = ANY(?)", pq.Array(ids)).
		Order(columName).
		Pluck("id", &newIds).
		Error
	return
}

func (repo articleRepository) UpdateArticlesPage(ctx context.Context) error {

	return nil
}
