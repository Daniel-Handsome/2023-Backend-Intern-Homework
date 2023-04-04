package article

import (
	"github.com/Daniel-Handsome/2023-Backend-intern-Homework/model"
	"github.com/Daniel-Handsome/2023-Backend-intern-Homework/pb"
	"github.com/golang/protobuf/ptypes/timestamp"
)

func ArticleToProto(infos []model.Article) []*pb.Article {
	var result []*pb.Article
	for _, val := range infos {
		result = append(result, PostInfoToProto(val))
	}
	return result
}

func ArticleToProto(article model.Article) *pb.Article {
	return &pb.Article{
		Uuid:        article.Uuid,
		Title:       article.Title,
		Content:     article.Content,
		PostStartAt: &timestamp.Timestamp{Seconds: data.PostStartAt.Unix()},
		PostEndAt:   &timestamp.Timestamp{Seconds: data.PostEndAt.Unix()},
	}
}
