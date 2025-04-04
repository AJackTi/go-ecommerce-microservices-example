package pipelines

import (
	"context"
	"fmt"
	"strings"

	"github.com/mehdihadeli/go-mediatr"
	"go.opentelemetry.io/otel/attribute"

	"github.com/AJackTi/go-ecommerce-microservices-example/internal/pkg/otel/constants/telemetrytags"
	"github.com/AJackTi/go-ecommerce-microservices-example/internal/pkg/otel/constants/tracing/components"
	"github.com/AJackTi/go-ecommerce-microservices-example/internal/pkg/otel/tracing"
	customAttribute "github.com/AJackTi/go-ecommerce-microservices-example/internal/pkg/otel/tracing/attribute"
	"github.com/AJackTi/go-ecommerce-microservices-example/internal/pkg/otel/tracing/utils"
	"github.com/AJackTi/go-ecommerce-microservices-example/internal/pkg/reflection/typemapper"
)

type mediatorTracingPipeline struct {
	config *config
	tracer tracing.AppTracer
}

func NewMediatorTracingPipeline(appTracer tracing.AppTracer, opts ...Option) mediatr.PipelineBehavior {
	cfg := defaultConfig
	for _, opt := range opts {
		opt.apply(cfg)
	}

	return &mediatorTracingPipeline{
		config: cfg,
		tracer: appTracer,
	}
}

func (m *mediatorTracingPipeline) Handle(
	ctx context.Context,
	request interface{},
	next mediatr.RequestHandlerFunc,
) (interface{}, error) {
	requestName := typemapper.GetSnakeTypeName(request)

	componentName := components.RequestHandler
	requestNameTag := telemetrytags.App.RequestName
	requestTag := telemetrytags.App.Request
	requestResultNameTag := telemetrytags.App.RequestResultName
	requestResultTag := telemetrytags.App.RequestResult

	switch {
	case strings.Contains(typemapper.GetPackageName(request), "command") || strings.Contains(typemapper.GetPackageName(request), "commands"):
		componentName = components.CommandHandler
		requestNameTag = telemetrytags.App.CommandName
		requestTag = telemetrytags.App.Command
		requestResultNameTag = telemetrytags.App.CommandResultName
		requestResultTag = telemetrytags.App.CommandResult
	case strings.Contains(typemapper.GetPackageName(request), "query") || strings.Contains(typemapper.GetPackageName(request), "queries"):
		componentName = components.QueryHandler
		requestNameTag = telemetrytags.App.QueryName
		requestTag = telemetrytags.App.Query
		requestResultNameTag = telemetrytags.App.QueryResultName
		requestResultTag = telemetrytags.App.QueryResult
	case strings.Contains(typemapper.GetPackageName(request), "event") || strings.Contains(typemapper.GetPackageName(request), "events"):
		componentName = components.EventHandler
		requestNameTag = telemetrytags.App.EventName
		requestTag = telemetrytags.App.Event
		requestResultNameTag = telemetrytags.App.EventResultName
		requestResultTag = telemetrytags.App.EventResult
	}

	operationName := fmt.Sprintf("%s_handler", requestName)
	spanName := fmt.Sprintf("%s.%s/%s", componentName, operationName, requestName) // by convention

	// https://golang.org/pkg/context/#Context
	newCtx, span := m.tracer.Start(ctx, spanName)
	defer span.End()

	span.SetAttributes(
		attribute.String(requestNameTag, requestName),
		customAttribute.Object(requestTag, request),
	)

	response, err := next(newCtx)

	responseName := typemapper.GetSnakeTypeName(response)
	span.SetAttributes(
		attribute.String(requestResultNameTag, responseName),
		customAttribute.Object(requestResultTag, response),
	)

	err = utils.TraceStatusFromSpan(span, err)

	return response, err
}
