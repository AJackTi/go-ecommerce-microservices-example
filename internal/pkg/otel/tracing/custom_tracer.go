package tracing

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"

	"github.com/AJackTi/go-ecommerce-microservices-example/internal/pkg/otel/tracing/utils"
)

type AppTracer interface {
	trace.Tracer
}

type appTracer struct {
	trace.Tracer
}

func NewAppTracer(name string, options ...trace.TracerOption) AppTracer {
	// without registering `NewOtelTracing` is uses global empty (NoopTracer) TraceProvider but after using `NewOtelTracing`, global TraceProvider will be replaced
	tracer := otel.Tracer(name, options...)
	return &appTracer{
		Tracer: tracer,
	}
}

func (a *appTracer) Start(ctx context.Context, spanName string, opts ...trace.SpanStartOption) (context.Context, trace.Span) {
	parentSpan := trace.SpanFromContext(ctx)
	if parentSpan != nil {
		utils.ContextWithParentSpan(ctx, parentSpan)
	}

	return a.Tracer.Start(ctx, spanName, opts...)
}
