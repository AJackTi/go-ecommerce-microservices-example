package fxapp

import (
	"context"

	"go.uber.org/fx"

	"github.com/AJackTi/go-ecommerce-microservices-example/internal/pkg/config/environment"
	"github.com/AJackTi/go-ecommerce-microservices-example/internal/pkg/fxapp/contracts"
	"github.com/AJackTi/go-ecommerce-microservices-example/internal/pkg/logger"
)

type application struct {
	provides    []interface{}
	decorates   []interface{}
	invokes     []interface{}
	options     []fx.Option
	logger      logger.Logger
	fxApp       *fx.App
	environment environment.Environment
}

func NewApplication(
	providers []interface{},
	decorates []interface{},
	options []fx.Option,
	logger logger.Logger,
	env environment.Environment) contracts.Application {
	return &application{
		provides:    providers,
		decorates:   decorates,
		options:     options,
		logger:      logger,
		environment: env,
	}
}

func (a *application) Environment() environment.Environment {
	return a.environment
}

func (a *application) Logger() logger.Logger {
	return a.logger
}

func (a *application) RegisterHook(function interface{}) {
	a.invokes = append(a.invokes, function)
}

func (a *application) ResolveFunc(function interface{}) {
	a.invokes = append(a.invokes, function)
}

func (a *application) ResolveFuncWithParamTag(function interface{}, paramTagName string) {
	a.invokes = append(a.invokes, fx.Annotate(function, fx.ParamTags(paramTagName)))
}

func (a *application) Run() {
	// build phase of container will do in this stage, containing provides and invokes but app not started yet and will be started in the future with `fxApp.Register`
	fxApp := CreateFxApp(a)

	a.fxApp = fxApp

	// running phase will do in this stage and all register event hooks like OnStart and OnStop
	// instead of run for handling start and stop and create a ctx and cancel we can handle them manually with appconfigfx.start and appconfigfx.stop
	// https://github.com/uber-go/fx/blob/v1.20.0/app.go#L573
	fxApp.Run()
}

func (a *application) Start(ctx context.Context) error {
	// build phase of container will do in this stage, containing provides and invokes but app not started yet and will be started in the future with `fxApp.Register`
	fxApp := CreateFxApp(a)
	a.fxApp = fxApp

	return fxApp.Start(ctx)
}

func (a *application) Stop(ctx context.Context) error {
	if a.fxApp == nil {
		a.logger.Fatal("Failed to stop because application not started.")
	}

	return a.fxApp.Stop(ctx)
}

func (a *application) Wait() <-chan fx.ShutdownSignal {
	if a.fxApp == nil {
		a.logger.Fatal("Failed to wait because application not started.")
	}

	return a.fxApp.Wait()
}
