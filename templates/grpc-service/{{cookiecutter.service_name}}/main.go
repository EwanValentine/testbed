package main

import (
	"log"
	"net"

	"go.uber.org/zap"
	"google.golang.org/grpc"

	"github.com/EwanValentine/testbed/services/{{cookiecutter.service_name}}/internal/config"
	pb "github.com/EwanValentine/testbed/services/{{cookiecutter.service_name}}/gen/go/proto"
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
	pb.Register{{cookiecutter.service_name.capitalize() }}ServiceServer(grpcServer, grpc.New(logger, conf))
	log.Fatal(grpcServer.Serve(lis))
}
