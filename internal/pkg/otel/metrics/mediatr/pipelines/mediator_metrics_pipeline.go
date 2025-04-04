package pipelines

import (
	"context"
	"fmt"
	"time"

	"github.com/mehdihadeli/go-mediatr"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"

	"github.com/AJackTi/go-ecommerce-microservices-example/internal/pkg/core/cqrs"
	"github.com/AJackTi/go-ecommerce-microservices-example/internal/pkg/core/events"
	"github.com/AJackTi/go-ecommerce-microservices-example/internal/pkg/otel/constants/telemetrytags"
	"github.com/AJackTi/go-ecommerce-microservices-example/internal/pkg/otel/metrics"
	customAttribute "github.com/AJackTi/go-ecommerce-microservices-example/internal/pkg/otel/tracing/attribute"
	"github.com/AJackTi/go-ecommerce-microservices-example/internal/pkg/reflection/typemapper"
)

type mediatorMetricsPipeline struct {
	config *config
	meter  metrics.AppMetrics
}

func NewMediatorMetricsPipeline(appMetrics metrics.AppMetrics, opts ...Option) mediatr.PipelineBehavior {
	cfg := defaultConfig
	for _, opt := range opts {
		opt.apply(cfg)
	}

	return &mediatorMetricsPipeline{
		config: cfg,
		meter:  appMetrics,
	}
}

func (m *mediatorMetricsPipeline) Handle(ctx context.Context, request interface{}, next mediatr.RequestHandlerFunc) (interface{}, error) {
	payloadSnakeTypeName := typemapper.GetSnakeTypeName(request)
	typeName := typemapper.GetTypeName(request)

	nameTag := telemetrytags.App.RequestName
	typeNameTag := telemetrytags.App.RequestType
	payloadTag := telemetrytags.App.Request
	resultSnakeTypeNameTag := telemetrytags.App.RequestResultName
	resultTag := telemetrytags.App.RequestResult

	if cqrs.IsCommand(request) {
		nameTag = telemetrytags.App.CommandName
		typeNameTag = telemetrytags.App.CommandType
		payloadTag = telemetrytags.App.Command
		resultSnakeTypeNameTag = telemetrytags.App.CommandResultName
		resultTag = telemetrytags.App.CommandResult
	} else if cqrs.IsQuery(request) {
		nameTag = telemetrytags.App.QueryName
		typeNameTag = telemetrytags.App.QueryType
		payloadTag = telemetrytags.App.Query
		resultSnakeTypeNameTag = telemetrytags.App.QueryResultName
		resultTag = telemetrytags.App.QueryResult
	} else if events.IsEvent(request) {
		nameTag = telemetrytags.App.EventName
		typeNameTag = telemetrytags.App.EventType
		payloadTag = telemetrytags.App.Event
		resultSnakeTypeNameTag = telemetrytags.App.EventResultName
		resultTag = telemetrytags.App.EventResult
	}

	successRequestsCounter, err := m.meter.Int64Counter(
		fmt.Sprintf("%s.success_total", payloadSnakeTypeName),
		metric.WithUnit("count"),
		metric.WithDescription(fmt.Sprintf("Measures the number of the success '%s' (%s)", payloadSnakeTypeName, typeName)),
	)
	if err != nil {
		return nil, err
	}

	failureRequestsCounter, err := m.meter.Int64Counter(
		fmt.Sprintf("%s.failed_total", payloadSnakeTypeName),
		metric.WithUnit("count"),
		metric.WithDescription(fmt.Sprintf("Measures the number of the failed '%s' (%s)", payloadSnakeTypeName, typeName)),
	)
	if err != nil {
		return nil, err
	}

	totalRequestsCounter, err := m.meter.Int64Counter(
		fmt.Sprintf("%s.total", payloadSnakeTypeName),
		metric.WithUnit("count"),
		metric.WithDescription(fmt.Sprintf("Measures the total number of '%s' (%s)", payloadSnakeTypeName, typeName)),
	)
	if err != nil {
		return nil, err
	}

	durationValueRecorder, err := m.meter.Int64Histogram(
		fmt.Sprintf("%s.duration", payloadSnakeTypeName),
		metric.WithUnit("ms"),
		metric.WithDescription(fmt.Sprintf("Measures the duration of '%s' (%s)", payloadSnakeTypeName, typeName)),
	)
	if err != nil {
		return nil, err
	}

	// Start recording the duration
	startTime := time.Now()

	response, err := next(ctx)

	// Calculate the duration
	duration := time.Since(startTime).Microseconds()

	// response will be nill if we have an error
	responseSnakeName := typemapper.GetSnakeTypeName(response)

	opt := metric.WithAttributes(
		attribute.String(nameTag, payloadSnakeTypeName),
		attribute.String(typeNameTag, typeName),
		customAttribute.Object(payloadTag, request),
		attribute.String(resultSnakeTypeNameTag, responseSnakeName),
		customAttribute.Object(resultTag, response),
	)

	// Record metrics
	totalRequestsCounter.Add(ctx, 1, opt)

	if err == nil {
		successRequestsCounter.Add(ctx, 1, opt)
	} else {
		failureRequestsCounter.Add(ctx, 1, opt)
	}

	durationValueRecorder.Record(ctx, duration, opt)

	return response, err
}
