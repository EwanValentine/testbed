package grpc

import (
	"go.uber.org/zap"

	"github.com/EwanValentine/testbed/services/{{cookiecutter.service_name}}/internal/config"
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
