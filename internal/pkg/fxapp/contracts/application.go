package contracts

import (
	"context"

	"go.uber.org/fx"

	"github.com/AJackTi/go-ecommerce-microservices-example/internal/pkg/config/environment"
	"github.com/AJackTi/go-ecommerce-microservices-example/internal/pkg/logger"
)

type Application interface {
	Container
	RegisterHook(function interface{})
	Run()
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
	Wait() <-chan fx.ShutdownSignal
	Logger() logger.Logger
	Environment() environment.Environment
}
