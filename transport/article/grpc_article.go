package article

import (
	"context"

	"github.com/Daniel-Handsome/2023-Backend-intern-Homework/internal/Error"
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
		return nil, Error.ErrServerError.Error(err)
	}
	return &pb.GetArticlesPageRes{
		Articles:   ArticlesToProto(articles),
		NexPageKey: nextKey,
	}, nil
}

func (g GrpcServer) UpdateArticlesPage(ctx context.Context, req *pb.UpdateArticlesPageReq) (*pb.UpdateArticlesPageRes, error) {
	err := g.srv.UpdateArticlesPage(ctx,
		ProtoToOrderColumn(req.GetOrderColumns()),
	)

	if err != nil {
		return &pb.UpdateArticlesPageRes{}, Error.ErrServerError.Error(err)
	}
	return &pb.UpdateArticlesPageRes{}, nil
}
