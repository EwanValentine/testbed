package main

import (
	"log"
	"net"

	"go.uber.org/zap"
	"google.golang.org/grpc"

	pb "github.com/EwanValentine/testbed/services/calculator/gen/go/proto"
	"github.com/EwanValentine/testbed/services/calculator/internal/config"
	grpcService "github.com/EwanValentine/testbed/services/calculator/transports/grpc"
)

func main() {
	logger, _ := zap.NewProduction()
	conf := config.Load()

	lis, err := net.Listen("tcp", conf.Addr)
	if err != nil {
		log.Panic(err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterCalculatorServiceServer(grpcServer, grpcService.New(logger, conf))
	log.Fatal(grpcServer.Serve(lis))
}
