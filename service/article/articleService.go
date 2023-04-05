package article

import (
	"context"
	"fmt"

	"github.com/Daniel-Handsome/2023-Backend-intern-Homework/model"
	"github.com/Daniel-Handsome/2023-Backend-intern-Homework/repository"
	"gorm.io/gorm"
)

type ArticleService interface {
	GetArticlesPage(ctx context.Context, headKey string) (articles []model.Article, nextKey string, err error)
	UpdateArticlesPage(ctx context.Context, column model.OrderColumn) error
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

	node, err := srv.pageNodeRepo.GetByUuid(ctx, pageLinkList.Head)
	if err != nil {
		return
	}

	articles, err = srv.repo.GetPage(ctx, node.ArticleIds)
	if err != nil {
		return nil, "", err
	}

	nextKey = node.Next
	return
}

func (srv articleService) UpdateArticlesPage(ctx context.Context, column model.OrderColumn) error {
	err := srv.repo.Transaction(func(tx *gorm.DB) error {
		pageLinkLists, err := srv.pageLinkListRepo.GetByType(ctx, model.RECOMMEND)
		for _, pageLinkList := range pageLinkLists {
			if err != nil {
				return err
			}

			var articleIds []int64
			//head

			headNode, err := srv.pageNodeRepo.GetByUuid(ctx, pageLinkList.Head)
			articleIds = append(articleIds, headNode.ArticleIds...)

			nextUUID := headNode.Next
			for {
				node, err := srv.pageNodeRepo.GetByUuid(ctx, nextUUID)
				if err != nil {
					break
				}
				articleIds = append(articleIds, node.ArticleIds...)

				if node.Next == "" {
					break
				}

				nextUUID = node.Next
			}

			if articleIds == nil {
				return fmt.Errorf("no articles")
			}

			newArticleIds, err := srv.repo.GetIdsSortByOmitId(ctx, articleIds, column)
			if err != nil {
				return err
			}

			//update
			err = srv.updatePageNodes(ctx, pageLinkList.Head, newArticleIds)
			if err != nil {
				fmt.Println(4)
				return err
			}
		}

		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (srv *articleService) updatePageNodes(
	ctx context.Context,
	headUUID string,
	newArticleIDs []int64,
) error {
	// 先取得 head 對應的 PageNode 記錄
	head, err := srv.pageNodeRepo.UpdateArticleIds(ctx, headUUID, newArticleIDs[:model.PAGE_LIMIT])
	if err != nil {
		return err
	}
	// 取得下一個 PageNode 的 uuid，開始循環更新
	nextUUID := head.Next
	currentIndex := model.PAGE_LIMIT
	for nextUUID != "" {
		// 更新 article_ids
		newIds := newArticleIDs[(currentIndex):(currentIndex + 2)]
		node, err := srv.pageNodeRepo.UpdateArticleIds(ctx, nextUUID, newIds)
		if err != nil {
			return err
		}
		currentIndex = currentIndex + model.PAGE_LIMIT
		// 更新下一個 PageNode
		nextUUID = node.Next
	}

	return nil
}
