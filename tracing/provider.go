package tracing

import (
	"os"
	"runtime"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.27.0"
)

func makeProvider(opts []tracesdk.TracerProviderOption, ratio float64, extraAttributes ...attribute.KeyValue) error {
	hostname, err := os.Hostname()
	if err != nil {
		return err
	}
	extraAttributes = append(extraAttributes,
		semconv.HostID(hostname),
		semconv.OSName(runtime.GOOS),
		semconv.HostArchKey.String(runtime.GOARCH),
		semconv.TelemetrySDKLanguageGo,
	)
	var traceProviderOptions []tracesdk.TracerProviderOption
	traceProviderOptions = append(traceProviderOptions,
		// Always be sure to batch in production.
		tracesdk.WithBatcher(exp),
		// set sampling part of data
		tracesdk.WithSampler(tracesdk.ParentBased(tracesdk.TraceIDRatioBased(ratio))),
		// Record information about this application in a Resource.
		tracesdk.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			extraAttributes...,
		)),
	)
	traceProviderOptions = append(traceProviderOptions, opts...)
	tp := tracesdk.NewTracerProvider(traceProviderOptions...)

	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	))

	// Register our TracerProvider as the global so any imported
	// instrumentation in the future will default to using it.
	otel.SetTracerProvider(tp)
	return nil
}
