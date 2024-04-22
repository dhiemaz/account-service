package grpc

import (
	"context"
	"fmt"
	"github.com/dhiemaz/AccountService/config"
	"github.com/dhiemaz/AccountService/infrastructures/logger"
	"github.com/dhiemaz/AccountService/infrastructures/server/grpc/proto/access_svc"
	"github.com/dhiemaz/AccountService/infrastructures/server/grpc/proto/account_svc"
	"github.com/dhiemaz/AccountService/infrastructures/server/grpc/proto/office_svc"
	"github.com/dhiemaz/AccountService/infrastructures/server/grpc/proto/role_svc"
	"github.com/dhiemaz/AccountService/internal/container"
	accessHandler "github.com/dhiemaz/AccountService/internal/domain/access/handler"
	accessUsecase "github.com/dhiemaz/AccountService/internal/domain/access/usecase"
	accountHandler "github.com/dhiemaz/AccountService/internal/domain/account/handler"
	accountUsecase "github.com/dhiemaz/AccountService/internal/domain/account/usecase"
	officeHandler "github.com/dhiemaz/AccountService/internal/domain/office/handler"
	officeUsecase "github.com/dhiemaz/AccountService/internal/domain/office/usecase"
	roleHandler "github.com/dhiemaz/AccountService/internal/domain/role/handler"
	roleUsecase "github.com/dhiemaz/AccountService/internal/domain/role/usecase"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func RunServer(cfg *config.Config) {
	logger.WithFields(logger.Fields{"component": "cmd", "action": "RunServer"}).
		Infof("starting gRPC server...")

	tls := cfg.GrpcTLS
	opts := []grpc.ServerOption{}

	if tls {
		serverCert := "server.crt"
		serverKey := "server.key"
		creds, err := credentials.NewServerTLSFromFile(serverCert, serverKey)
		if err != nil {
			logger.WithFields(logger.Fields{"component": "cmd", "action": "RunServer"}).
				Errorf("failed to loading certificates, error : %v", err)

			os.Exit(1)
		}

		opts = append(opts, grpc.Creds(creds))
	}

	opts = append(opts, grpc.UnaryInterceptor(serverInterceptor))
	server := grpc.NewServer(opts...)

	ctn := container.NewContainer()
	Apply(server, ctn)

	svcHost := cfg.GrpcHost
	svcPort := cfg.GrpcPort

	go func() {
		lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", svcHost, svcPort))
		if err != nil {
			logger.WithFields(logger.Fields{"component": "cmd", "action": "RunServer"}).
				Errorf("failed listen, error : %v", err)

			os.Exit(1)
		}

		if err := server.Serve(lis); err != nil {
			logger.WithFields(logger.Fields{"component": "cmd", "action": "RunServer"}).
				Errorf("failed start gRPC server, error : %v", err)
		}
	}()

	logger.WithFields(logger.Fields{"component": "cmd", "action": "RunServer"}).
		Infof("gRPC server is running at %s:%d", svcHost, svcPort)

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	signal := <-c

	logger.WithFields(logger.Fields{"component": "cmd", "action": "RunServer"}).
		Infof("process killed with signal: %v", signal.String())

}

func serverInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	// Calls the handler
	res, err := handler(ctx, req)
	if err != nil {
		if e, ok := status.FromError(err); ok {
			logger.WithFields(logger.Fields{"component": "grpc", "action": info.FullMethod}).
				Errorf("code : %v, response : %v, ", e.Code(), e.Message())
		}
	} else {
		logger.WithFields(logger.Fields{"component": "grpc", "action": info.FullMethod}).
			Infof("code : %v, response : %v, ", codes.OK, codes.OK.String())
	}
	return res, err
}

func Apply(server *grpc.Server, ctn *container.Container) {
	access_svc.RegisterAccessServiceServer(server, accessHandler.NewAccessHandler(ctn.Resolve("accessSvc").(accessUsecase.AccessUsecase)))
	account_svc.RegisterEmployeeServiceServer(server, accountHandler.NewAccountHandler(ctn.Resolve("accountSvc").(accountUsecase.AccountUsecase)))
	office_svc.RegisterOfficeServiceServer(server, officeHandler.NewOfficeHandler(ctn.Resolve("officeSvc").(officeUsecase.OfficeUsecase)))
	role_svc.RegisterRoleServiceServer(server, roleHandler.NewRoleHandler(ctn.Resolve("roleSvc").(roleUsecase.RoleUsecase)))
}
