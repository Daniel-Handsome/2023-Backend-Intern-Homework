package article

import (
	"context"

	"github.com/Daniel-Handsome/2023-Backend-intern-Homework/model"
	"github.com/Daniel-Handsome/2023-Backend-intern-Homework/repository"
)

type ArticleService interface {
	GetArticlesPage(ctx context.Context, headKey string) (artcle []model.Article, nextKey string, err error)
}

type articleService struct {
	repo             repository.ArticleRepository
	pageLinkListRepo repository.PageLinkedListRepository
	pageNodeRepo     repository.PageNodeRepository
}

func NewArticleService(
	repo repository.ArticleRepository,
	pageLinkedListRepo repository.PageLinkedListRepository,
	pageNodeRepo repository.PageNodeRepository,
) ArticleService {
	return &articleService{
		repo:             repo,
		pageLinkListRepo: pageLinkedListRepo,
		pageNodeRepo:     pageNodeRepo,
	}
}

func (srv articleService) GetArticlesPage(ctx context.Context, headKey string) (articles []model.Article, nextKey string, err error) {
	pageLinkList, err := srv.pageLinkListRepo.GetByHeadKey(ctx, headKey)
	if err != nil {
		return
	}

	node, err := srv.pageNodeRepo.GetByUuid(ctx, pageLinkList.Uuid)
	if err != nil {
		return
	}

	articles, err = srv.repo.GetByIds(ctx, node.ArticleIds)
	if err != nil {
		return nil, "", err
	}

	nextKey = node.Next
	return
}
