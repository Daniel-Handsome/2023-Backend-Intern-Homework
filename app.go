package app

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	pbHealth "google.golang.org/grpc/health/grpc_health_v1"

	"github.com/Daniel-Handsome/2023-Backend-intern-Homework/internal/Interceptor"
	"github.com/Daniel-Handsome/2023-Backend-intern-Homework/pb"
	"github.com/Daniel-Handsome/2023-Backend-intern-Homework/pkg/repote"
	"github.com/Daniel-Handsome/2023-Backend-intern-Homework/repository"
	"github.com/Daniel-Handsome/2023-Backend-intern-Homework/service/article"
	"github.com/Daniel-Handsome/2023-Backend-intern-Homework/service/user"
	trasnportArticle "github.com/Daniel-Handsome/2023-Backend-intern-Homework/transport/article"
	trasnportUser "github.com/Daniel-Handsome/2023-Backend-intern-Homework/transport/user"
	"github.com/Daniel-Handsome/2023-Backend-intern-Homework/utils"
	log "github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/reflection"
	"gorm.io/gorm"
)

type App struct {
	gorm *gorm.DB
}

func NewApp(gorm *gorm.DB) *App {
	return &App{gorm: gorm}
}

func (a *App) Start() {
	userRepo := repository.NewUserRepository(a.gorm)
	pageLinkedListRepo := repository.NewPageLinkedListRepository(a.gorm)
	articleRepo := repository.NewArticleRepository(a.gorm)
	pageNodeRepo := repository.NewPageNodeRepository(a.gorm)

	userSrv := user.NewUserService(userRepo)
	articleSrv := article.NewArticleService(articleRepo, pageLinkedListRepo, pageNodeRepo)

	// grpc server
	userGrpc := trasnportUser.NewGrpcServer(userSrv)
	articlesGrpc := trasnportArticle.NewGrpcServer(articleSrv)

	g := new(errgroup.Group)
	g.Go(func() error {
		errReporter := repote.NewLocal()
		interceptor := Interceptor.NewGrpcInterceptor().
			WithUnaryPanicInterceptor(errReporter).
			WithStreamPanicInterceptor(errReporter)

		grpcServer := grpc.NewServer(
			interceptor.ToGrpcOptions()...,
		)
		defer grpcServer.GracefulStop()

		// grpc registers
		pb.RegisterUserServiceServer(grpcServer, userGrpc)
		pb.RegisterArticleServiceServer(grpcServer, articlesGrpc)

		reflection.Register(grpcServer)

		healthServer := health.NewServer()
		//SetServingStatus not use because default status is serving
		pbHealth.RegisterHealthServer(grpcServer, healthServer)

		listener, err := net.Listen("tcp", fmt.Sprintf(":%v", utils.GetConfigToInt("Grpc_port")))
		if err != nil {
			return err
		}

		// 創建一個 context，用於接收關閉信號
		ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
		defer stop()

		// 開啟一個 goroutine 監聽關閉信號
		go func() {
			<-ctx.Done()
			log.Println("Shutting down gracefully, please wait...")
			// 調用 grpc server 的 GracefulStop 方法進行優雅關閉
			grpcServer.GracefulStop()
		}()

		// 啟動 grpc server
		err = grpcServer.Serve(listener)
		if err != nil {
			return err
		}
		return nil
	})

	// tcp healthcheck
	//g.Go(func() error {
	//	tcpHandler := func(conn net.Conn) error {
	//		return conn.Close()
	//	}
	//	healthCheck := protocol.NewTcpProtocol(int(utils.GetConfigToInt("Grpc_port")), tcpHandler)
	//	err := healthCheck.Serve()
	//	defer func() {
	//		_ = healthCheck.Close()
	//	}()
	//
	//	return err
	//})

	log.Printf(" [127.0.0.1:%v] services up \n", utils.GetConfigToInt("Grpc_port"))

	if err := g.Wait(); err != nil {
		log.Errorf("server start error: %v", err)
	}
}
