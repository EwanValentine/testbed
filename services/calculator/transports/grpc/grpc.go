package grpc

import (
	"context"

	"go.uber.org/zap"

	pb "github.com/EwanValentine/testbed/services/calculator/gen/go/proto"
	"github.com/EwanValentine/testbed/services/calculator/internal/config"
)

type Service struct {
	logger *zap.Logger
	conf   config.Config
	pb.UnimplementedCalculatorServiceServer
}

func New(logger *zap.Logger, conf config.Config) *Service {
	return &Service{
		logger: logger,
		conf:   conf,
	}
}

func (s *Service) Add(ctx context.Context, req *pb.AddRequest) (*pb.AddResponse, error) {
	result := req.Values.GetA() + req.Values.GetB()
	return &pb.AddResponse{
		Result: result,
	}, nil
}

func (s *Service) Subtract(ctx context.Context, req *pb.SubtractRequest) (*pb.SubtractResponse, error) {
	result := req.Values.GetA() - req.Values.GetB()
	return &pb.SubtractResponse{
		Result: result,
	}, nil
}

func (s *Service) Multiply(ctx context.Context, req *pb.MultiplyRequest) (*pb.MultiplyResponse, error) {
	result := req.Values.GetA() * req.Values.GetB()
	return &pb.MultiplyResponse{
		Result: result,
	}, nil
}
