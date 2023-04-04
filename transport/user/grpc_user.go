package user

import (
	"context"

	"github.com/Daniel-Handsome/2023-Backend-intern-Homework/pb"
	"github.com/Daniel-Handsome/2023-Backend-intern-Homework/service/user"
)

type GrpcServer struct {
	srv user.UserService
}

func NewGrpcServer(srv user.UserService) *GrpcServer {
	return &GrpcServer{srv: srv}
}

func (g GrpcServer) GetUserArticlesHeadKey(ctx context.Context, req *pb.GetUserArticlesHeadKeyReq) (*pb.GetUserArticlesHeadKeyRes, error) {
	userArticlesPageKey, err := g.srv.GetUserArticlesHeadKey(ctx, req.GetUserId())
	if err != nil {
		return nil, err
	}

	return &pb.GetUserArticlesHeadKeyRes{ArticlePageHeadKey: userArticlesPageKey}, nil
}
