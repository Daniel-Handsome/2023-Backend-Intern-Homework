package article

import (
	"github.com/Daniel-Handsome/2023-Backend-intern-Homework/model"
	"github.com/Daniel-Handsome/2023-Backend-intern-Homework/pb"
	"github.com/golang/protobuf/ptypes/timestamp"
)

func ProtoToOrderColumn(v pb.OrderColumn) model.OrderColumn {
	switch v {
	case pb.OrderColumn_CreateAt:
		return model.CreateAt
	case pb.OrderColumn_UpdateAt:
		return model.UpdateAt
	default:
		return model.Id
	}
}

func ArticlesToProto(articles []model.Article) []*pb.Article {
	var result []*pb.Article
	for _, article := range articles {
		result = append(result, ArticleToProto(article))
	}
	return result
}

func ArticleToProto(article model.Article) *pb.Article {
	return &pb.Article{
		Uuid:      article.Uuid,
		Title:     article.Title,
		Content:   article.Content,
		CreatedAt: &timestamp.Timestamp{Seconds: article.CreatedAt.Unix()},
		UpdateAt:  &timestamp.Timestamp{Seconds: article.UpdatedAt.Unix()},
	}
}
