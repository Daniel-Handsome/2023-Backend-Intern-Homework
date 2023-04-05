package Setup

import (
	"fmt"

	"github.com/Daniel-Handsome/2023-Backend-intern-Homework/pb"
	"github.com/Daniel-Handsome/2023-Backend-intern-Homework/utils"
	"google.golang.org/grpc"
)

func UserClient() pb.UserServiceClient {
	cc := getGrpcConnection(int(utils.GetConfigToInt("Grpc_port")))
	return pb.NewUserServiceClient(cc)
}

func ArticleClient() pb.ArticleServiceClient {
	cc := getGrpcConnection(int(utils.GetConfigToInt("Grpc_port")))
	return pb.NewArticleServiceClient(cc)
}

func getGrpcConnection(grpcPort int) *grpc.ClientConn {
	address := fmt.Sprintf(":%d", grpcPort)
	cc, _ := grpc.Dial(address, grpc.WithInsecure())
	return cc
}
