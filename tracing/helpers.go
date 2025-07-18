package tracing

import (
	"fmt"
	"runtime"
	"time"

	"go.opentelemetry.io/otel/attribute"
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

// AddEventWithCodeLocation creates event in span with current code location
func AddEventWithCodeLocation(span trace.Span, message string) {
	_, file, line, ok := runtime.Caller(1)
	if ok {
		span.AddEvent(message, trace.WithAttributes(
			semconv.CodeFilepath(file),
			semconv.CodeLineNumber(line),
		))
	}
}

// AddAttributeToSpan attach any supportable value to span
func AddAttributeToSpan(span trace.Span, name string, val any) {
	switch v := val.(type) {
	case bool:
		span.SetAttributes(attribute.Bool(name, v))
	case []bool:
		span.SetAttributes(attribute.BoolSlice(name, v))
	case int:
		span.SetAttributes(attribute.Int(name, v))
	case []int:
		span.SetAttributes(attribute.IntSlice(name, v))
	case int64:
		span.SetAttributes(attribute.Int64(name, v))
	case []int64:
		span.SetAttributes(attribute.Int64Slice(name, v))
	case float64:
		span.SetAttributes(attribute.Float64(name, v))
	case []float64:
		span.SetAttributes(attribute.Float64Slice(name, v))
	case string:
		span.SetAttributes(attribute.String(name, v))
	case []string:
		span.SetAttributes(attribute.StringSlice(name, v))
	case fmt.Stringer:
		span.SetAttributes(attribute.Stringer(name, v))
	case []fmt.Stringer:
		strings := make([]string, len(v))
		for i := range v {
			strings[i] = v[i].String()
		}
		span.SetAttributes(attribute.StringSlice(name, strings))
	case map[string]fmt.Stringer:
		for k := range v {
			span.SetAttributes(attribute.Stringer(name+"."+k, v[k]))
		}
	case time.Duration:
		span.SetAttributes(attribute.String(name, v.String()))
	case time.Time:
		span.SetAttributes(attribute.String(name, v.Format(time.RFC3339Nano)))
	default:
		span.SetAttributes(attribute.String(name, fmt.Sprint(v)))
	}
}
