package tracing

import (
	"runtime"

	semconv "go.opentelemetry.io/otel/semconv/v1.27.0"
	"go.opentelemetry.io/otel/trace"
)

// AttachCodeLocationToSpan attach current code location to span
func AttachCodeLocationToSpan(span trace.Span) {
	_, file, line, ok := runtime.Caller(1)
	if ok {
		span.SetAttributes(semconv.CodeFilepath(file), semconv.CodeLineNumber(line))
	}
}
