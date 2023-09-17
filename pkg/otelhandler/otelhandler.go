package otelhandler

import (
	"context"
	"log/slog"

	"go.opentelemetry.io/otel/trace"
)

type OTELHandler struct {
	slog.Handler
}

func (h OTELHandler) Handle(ctx context.Context, rec slog.Record) error {
	span := trace.SpanFromContext(ctx)
	spanCtx := span.SpanContext()
	attrs := []slog.Attr{}
	if spanCtx.HasTraceID() {
		traceID := span.SpanContext().TraceID()
		attrs = append(attrs, slog.String("traceID", traceID.String()))
	}
	if spanCtx.HasSpanID() {
		spanID := span.SpanContext().SpanID()
		attrs = append(attrs, slog.String("spanID", spanID.String()))
	}
	rec.AddAttrs(attrs...)
	return h.Handler.Handle(ctx, rec)
}
