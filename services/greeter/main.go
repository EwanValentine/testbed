package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/EwanValentine/testbed/services/greeter/gen/go/proto"
)

type service struct {
	pb.UnimplementedGreeterServiceServer
}

func (s *service) Greet(_ context.Context, req *pb.GreetRequest) (*pb.GreetResponse, error) {
	return &pb.GreetResponse{
		Greeting: "Hello " + req.Name,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Panic(err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterGreeterServiceServer(grpcServer, &service{})
	log.Fatal(grpcServer.Serve(lis))
}
