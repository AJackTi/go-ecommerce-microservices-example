package fxapp

import (
	"go.uber.org/fx"

	"github.com/AJackTi/go-ecommerce-microservices-example/internal/pkg/config/environment"
	"github.com/AJackTi/go-ecommerce-microservices-example/internal/pkg/fxapp/contracts"
	"github.com/AJackTi/go-ecommerce-microservices-example/internal/pkg/logger"
	"github.com/AJackTi/go-ecommerce-microservices-example/internal/pkg/logger/config"
	logrous "github.com/AJackTi/go-ecommerce-microservices-example/internal/pkg/logger/logrus"
	"github.com/AJackTi/go-ecommerce-microservices-example/internal/pkg/logger/models"
	"github.com/AJackTi/go-ecommerce-microservices-example/internal/pkg/logger/zap"
)

type applicationBuilder struct {
	provides    []interface{}
	decorates   []interface{}
	options     []fx.Option
	logger      logger.Logger
	environment environment.Environment
}

func NewApplicationBuilder(environments ...environment.Environment) contracts.ApplicationBuilder {
	env := environment.ConfigAppEnv(environments...)

	var logger logger.Logger
	logOption, err := config.ProvideLogConfig(env)
	if err != nil || logOption == nil {
		logger = zap.NewZapLogger(logOption, env)
	} else if logOption.LogType == models.Logrus {
		logger = logrous.NewLogrusLogger(logOption, env)
	} else {
		logger = zap.NewZapLogger(logOption, env)
	}

	return &applicationBuilder{
		logger:      logger,
		environment: env,
	}
}

func (a *applicationBuilder) Build() contracts.Application {
	app := NewApplication(a.provides, a.decorates, a.options, a.logger, a.environment)

	return app
}

func (a *applicationBuilder) Decorate(constructors ...interface{}) {
	a.decorates = append(a.decorates, constructors...)
}

func (a *applicationBuilder) Environment() environment.Environment {
	return a.environment
}

func (a *applicationBuilder) GetDecorates() []interface{} {
	return a.decorates
}

func (a *applicationBuilder) GetProviders() []interface{} {
	return a.provides
}

func (a *applicationBuilder) Logger() logger.Logger {
	return a.logger
}

func (a *applicationBuilder) Options() []fx.Option {
	return a.options
}

func (a *applicationBuilder) Provide(constructors ...interface{}) {
	a.provides = append(a.provides, constructors...)
}

func (a *applicationBuilder) ProvideModule(module fx.Option) {
	a.options = append(a.options, module)
}
