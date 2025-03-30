package utils

import (
	"context"

	"go.opentelemetry.io/otel/trace"
)

type traceContextKeyType int

const parentSpanKey traceContextKeyType = iota + 1

func ContextWithParentSpan(parent context.Context, span trace.Span) context.Context {
	return context.WithValue(parent, parentSpanKey, span)
}
