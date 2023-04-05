package Interceptor

import (
	"fmt"

	"github.com/Daniel-Handsome/2023-Backend-intern-Homework/pkg/repote"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GrpcInterceptor struct {
	unaryInterceptor  []grpc.UnaryServerInterceptor
	streamInterceptor []grpc.StreamServerInterceptor
}

func NewGrpcInterceptor() *GrpcInterceptor {
	return &GrpcInterceptor{}
}

func (g *GrpcInterceptor) WithUnaryInterceptor(interceptor ...grpc.UnaryServerInterceptor) *GrpcInterceptor {
	g.unaryInterceptor = append(g.unaryInterceptor, interceptor...)
	return g
}

func (g *GrpcInterceptor) WithStreamInterceptor(interceptor ...grpc.StreamServerInterceptor) *GrpcInterceptor {
	g.streamInterceptor = append(g.streamInterceptor, interceptor...)
	return g
}

func (g *GrpcInterceptor) WithUnaryPanicInterceptor(panicReporter repote.Reporter) *GrpcInterceptor {
	recoveryOpts := []grpc_recovery.Option{
		grpc_recovery.WithRecoveryHandler(func(p interface{}) (err error) {
			msg := p.(string)
			if err = panicReporter.Send(msg); err != nil {
				fmt.Println(err)
			}
			return status.New(codes.Unavailable, msg).Err()
		}),
	}
	return g.WithUnaryInterceptor(grpc_recovery.UnaryServerInterceptor(recoveryOpts...))
}

func (g *GrpcInterceptor) WithStreamPanicInterceptor(panicReporter repote.Reporter) *GrpcInterceptor {
	recoveryOpts := []grpc_recovery.Option{
		grpc_recovery.WithRecoveryHandler(func(p interface{}) (err error) {
			msg := p.(string)
			return status.New(codes.Unavailable, msg).Err()
		}),
	}
	return g.WithStreamInterceptor(grpc_recovery.StreamServerInterceptor(recoveryOpts...))
}

func (g *GrpcInterceptor) ToGrpcOptions(options ...grpc.ServerOption) []grpc.ServerOption {
	grpcOptions := []grpc.ServerOption{
		grpc.MaxRecvMsgSize(10 * 1024 * 1024),
		grpc.MaxSendMsgSize(10 * 1024 * 1024),
		grpc_middleware.WithUnaryServerChain(
			g.unaryInterceptor...,
		),
		grpc_middleware.WithStreamServerChain(
			g.streamInterceptor...,
		),
	}
	return append(grpcOptions, options...)
}
