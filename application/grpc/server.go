package grpc

import (
	"fmt"
	"log"
	"net"

	"github.com/jinzhu/gorm"
	"google.golang.org/grpc"
)

func StartGrpcServer(databse *gorm.DB, port int) {
	grpcServer := grpc.NewServer()

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
