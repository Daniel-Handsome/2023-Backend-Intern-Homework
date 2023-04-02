package repo

import (
	"context"
	"errors"

	"github.com/Daniel-Handsome/2023-Backend-intern-Homework/orm"
	"gorm.io/gorm"
)

var ErrArticleNotFound = errors.New("article not found")

type ArticleRepository interface {
    Create(ctx context.Context, Article *orm.Article) error
    Update(ctx context.Context, Article *orm.Article) error
    Delete(ctx context.Context, id int64) error
}

type articleRepository struct {
    db *gorm.DB
}

func NewArticleRepository (db *gorm.DB) ArticleRepository {
    return &articleRepository{db: db}
}

func (r *articleRepository) Create(ctx context.Context, Article *orm.Article) error {
    err := r.db.WithContext(ctx).Create(Article).Error
    return err
}

func (r *articleRepository) Update(ctx context.Context, Article *orm.Article) error {
    err := r.db.WithContext(ctx).Model(Article).Updates(orm.Article{}).Error
    return err
}

func (r *articleRepository) Delete(ctx context.Context, id int64) error {
    err := r.db.WithContext(ctx).Where("id = ?", id).Delete(&orm.Article{}).Error
	return err
}