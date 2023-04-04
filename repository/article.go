package repository

import (
	"context"
	"errors"

	"github.com/Daniel-Handsome/2023-Backend-intern-Homework/model"
	"gorm.io/gorm"
)

var ErrArticleNotFound = errors.New("article not found")

type ArticleRepository interface {
	Create(ctx context.Context, Article *model.Article) error
	Update(ctx context.Context, Article *model.Article) error
	Delete(ctx context.Context, id int64) error
	GetByIds(ctx context.Context, ids []int64) (articles []model.Article, err error)
}

type articleRepository struct {
	*BaseRepo
}

func NewArticleRepository(orm *gorm.DB) ArticleRepository {
	return &articleRepository{NewRepo(orm.Model(model.Article{}))}
}

func (repo *articleRepository) Create(ctx context.Context, Article *model.Article) error {
	err := repo.orm.WithContext(ctx).Create(Article).Error
	return err
}

func (repo *articleRepository) Update(ctx context.Context, Article *model.Article) error {
	err := repo.orm.WithContext(ctx).Model(Article).Updates(model.Article{}).Error
	return err
}

func (repo *articleRepository) Delete(ctx context.Context, id int64) error {
	err := repo.orm.WithContext(ctx).Where("id = ?", id).Delete(&model.Article{}).Error
	return err
}

func (repo *articleRepository) GetByIds(ctx context.Context, ids []int64) (articles []model.Article, err error) {
	err = repo.orm.WithContext(ctx).Where("id = ANY(?)", ids).Find(&articles).Error
	if err != nil {
		return
	}
	return
}
