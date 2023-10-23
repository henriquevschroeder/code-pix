package grpc

import (
	"fmt"
	"log"
	"net"

	"github.com/henriquevschroeder/code-pix/application/grpc/pb"
	"github.com/henriquevschroeder/code-pix/application/usecase"
	"github.com/henriquevschroeder/code-pix/infrastructure/repository"
	"github.com/jinzhu/gorm"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func StartGrpcServer(database *gorm.DB, port int) {
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	pixRepository := repository.PixKeyRepositoryDb{Db: database}
	pixUseCase := usecase.PixUseCase{PixKeyRepository: pixRepository}
	pixGrpcService := NewPixGrpcService(pixUseCase)
	pb.RegisterPixServiceServer(grpcServer, pixGrpcService)

	address := fmt.Sprintf("0.0.0.0:%d", port)

	listener, err := net.Listen("tcp", address)

	if err != nil {
		log.Fatal("could not start gRPC server", err)
	}

	log.Printf("gRPC server running on port %d", port)

	err = grpcServer.Serve(listener)

	if err != nil {
		log.Fatal("could not start gRPC server", err)
	}
}
