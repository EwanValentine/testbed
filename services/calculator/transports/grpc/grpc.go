package grpc

import (
	"github.com/EwanValentine/testbed/services/calculator/internal/config"
	"go.uber.org/zap"
)

type Service struct {
	logger *zap.Logger
	conf   config.Config
}

func New(logger *zap.Logger, conf config.Config) *Service {
	return &Service{
		logger: logger,
		conf:   conf,
	}
}
