package grpc

import (
	"context"
	"fmt"
	"github.com/dhiemaz/AccountService/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func RunServer() {
	log.Println("starting gRPC server...")

	cfg := config.GetConfig()
	tls := cfg.Server.TLS

	opts := []grpc.ServerOption{}

	if tls {
		serverCert := "server.crt"
		serverKey := "server.key"
		creds, err := credentials.NewServerTLSFromFile(serverCert, serverKey)
		if err != nil {
			log.Fatalln("failed to loading certificates, err : ", err)
			os.Exit(1)
		}

		opts = append(opts, grpc.Creds(creds))
	}

	opts = append(opts, grpc.UnaryInterceptor(serverInterceptor))

	// ctn := container.NewContainer()
	grpcServer := grpc.NewServer(opts...)

	//Apply(grpcServer, ctn)

	svcHost := cfg.Server.Host
	svcPort := cfg.Server.Port

	go func() {
		lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", svcHost, svcPort))
		if err != nil {
			log.Fatalf("failed listen, err : %v", err)
		}

		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed start gRPC server, err : %v", err)
		}
	}()

	fmt.Printf("BL gRPC server is running at %s:%d\n", svcHost, svcPort)

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	signal := <-c
	log.Fatalf("process killed with signal: %v\n", signal.String())
}

func serverInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	//m, err := getMetadata(ctx)
	//logUUID, _ := uuid.NewRandom()
	//logger.WriteLog(logger.Info, req, m, "request", "", info.FullMethod, logUUID.String(), codes.OK)

	//if err != nil {
	//	return container.MetaData{}, status.Errorf(codes.InvalidArgument, err.Error())
	//}

	//log.Println("Metadata ", m)
	//if _, err := validateMetadata(info, m); err != nil {
	//	return nil, err
	//}

	//if v, ok := req.(validator); ok {
	//	if err := v.Validate(); err != nil {
	//		logger.WriteLog(logger.Error, req, m, "request", err.Error(), info.FullMethod, logUUID.String(), codes.InvalidArgument)
	//		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	//	}
	//}

	// Calls the handler
	res, err := handler(ctx, req)
	if err != nil {
		if _, ok := status.FromError(err); ok {
			//logger.WriteLog(logger.Error, nil, m, "response", e.Message(), info.FullMethod, logUUID.String(), e.Code())
		}
	} else {
		//logger.WriteLog(logger.Info, res, m, "response", codes.OK.String(), info.FullMethod, logUUID.String(), codes.OK)
	}
	return res, err
}
