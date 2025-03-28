package contracts

import (
	"go.uber.org/fx"

	"github.com/AJackTi/go-ecommerce-microservices/internal/pkg/config/environment"
	"github.com/AJackTi/go-ecommerce-microservices/internal/pkg/logger"
)

type ApplicationBuilder interface {
	// ProvideModule register modules directly instead and modules should not register with `provided` function
	ProvideModule(module fx.Option)
	// Provide register functions constructors as dependency resolver
	Provide(constructors ...interface{})
	Decorate(constructors ...interface{})
	Build() Application

	GetProviders() []interface{}
	GetDecorates() []interface{}
	Options() []fx.Option
	Logger() logger.Logger
	Environment() environment.Environment
}
