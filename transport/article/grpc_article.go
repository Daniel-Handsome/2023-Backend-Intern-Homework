package article

import (
	"context"

	"github.com/Daniel-Handsome/2023-Backend-intern-Homework/pb"
	"github.com/Daniel-Handsome/2023-Backend-intern-Homework/service/article"
)

type GrpcServer struct {
	srv article.ArticleService
}

func NewGrpcServer(srv article.ArticleService) *GrpcServer {
	return &GrpcServer{srv: srv}
}

func (g GrpcServer) GetArticlesPage(ctx context.Context, rq *pb.GetArticlesPageReq) (*pb.GetArticlesPageRes, error) {
	articles, nextKey, err := g.srv.GetArticlesPage(ctx, rq.GetHeadKey())
	if err != nil {
		return nil, err
	}
	return &pb.GetArticlesPageRes{
		Articles:   articles,
		NexPageKey: nextKey,
	}, nil
}
