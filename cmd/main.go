package main

import (
	"os"

	"github.com/henriquevschroeder/code-pix/application/grpc"
	"github.com/henriquevschroeder/code-pix/infrastructure/db"
	"github.com/jinzhu/gorm"
)

var database *gorm.DB

func main() {
	database = db.ConnectDB(os.Getenv("env"))
	grpc.StartGrpcServer(database, 50051)
}
