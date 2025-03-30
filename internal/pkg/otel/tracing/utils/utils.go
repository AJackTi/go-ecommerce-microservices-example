package utils

import (
	"context"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"

	"github.com/AJackTi/go-ecommerce-microservices-example/internal/pkg/otel/constants/telemetrytags"
	"github.com/AJackTi/go-ecommerce-microservices-example/internal/pkg/utils/errorutils"
)

type traceContextKeyType int

const parentSpanKey traceContextKeyType = iota + 1

func ContextWithParentSpan(parent context.Context, span trace.Span) context.Context {
	return context.WithValue(parent, parentSpanKey, span)
}

func TraceStatusFromSpan(span trace.Span, err error) error {
	isError := err != nil

	var (
		code        codes.Code
		description = ""
	)

	if isError {
		code = codes.Error
		description = err.Error()
	} else {
		code = codes.Ok
	}

	span.SetStatus(code, description)

	if isError {
		stackTraceError := errorutils.ErrorsWithStack(err)

		// https://opentelemetry.io/docs/instrumentation/go/manual/#record-errors
		span.SetAttributes(
			attribute.String(telemetrytags.Exceptions.Message, err.Error()),
			attribute.String(telemetrytags.Exceptions.Stacktrace, stackTraceError),
		)
		span.RecordError(err)
	}

	return err
}
