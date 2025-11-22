package main

import (
	"time"

	"github.com/rs/zerolog/log"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.27.0"
	"go.opentelemetry.io/otel/trace"
	"golang.org/x/sync/errgroup"

	"github.com/vodolaz095/pkg/stopper"
	"github.com/vodolaz095/pkg/tracing"
	"github.com/vodolaz095/pkg/zerologger"
)

func main() {
	initialCtx, cancel := stopper.New()
	defer cancel()

	// configure logging
	zerologger.Configure(zerologger.Log{Level: zerologger.TraceLevel})

	// configure tracing
	tracingCfg := tracing.Config{
		Protocol:     "otlp_http",
		OTLPEndpoint: "http://localhost:4318/v1/traces",
		Ratio:        1,
		Insecure:     true,
		TraceProviderOptions: []tracesdk.TracerProviderOption{
			tracesdk.WithSampler(tracesdk.AlwaysSample()),
		},
	}
	log.Info().Msg(tracingCfg.String())

	err := tracing.StartWithContext(initialCtx, tracingCfg,
		semconv.ServiceName("pkg_example"),
		attribute.String("environment", "example"),
	)
	if err != nil {
		log.Error().Err(err).Msgf("error setting tracing")
		return
	}

	eg, ctx := errgroup.WithContext(initialCtx)
	eg.Go(func() error {
		tt := time.NewTicker(time.Second)
		for {
			select {
			case <-ctx.Done():
				tt.Stop()
				log.Info().Msg("ticker 1 is stopping...")
				return nil
			case <-tt.C:
				log.Warn().Msg("Hello from ticker 1!")
			}
		}
	})

	eg.Go(func() error {
		time.Sleep(500 * time.Millisecond)
		tt := time.NewTicker(time.Second)
		for {
			select {
			case <-ctx.Done():
				tt.Stop()
				log.Info().Msg("ticker 2 is stopping...")
				return nil
			case <-tt.C:
				log.Info().Msg("Hello from ticker 2!")
			}
		}
	})

	eg.Go(func() error {
		tt := time.NewTicker(3 * time.Second)
		var n = 0
		for {
			select {
			case <-ctx.Done():
				tt.Stop()
				log.Info().Msg("ticker 3 is stopping...")
				return nil
			case now := <-tt.C:
				n++
				_, span := otel.Tracer("ticker").Start(ctx, "loop",
					trace.WithAttributes(attribute.Int("iteration", n)))

				tracing.AddAttributeToSpan(span, "timestamp", now)
				span.AddEvent("Hello from ticker 3!", trace.WithAttributes(attribute.Int("iteration", n)))
				log.Info().
					Int("iteration", n).
					Str("trace_id", span.SpanContext().TraceID().String()).
					Msg("Hello from ticker 3!")
				span.End()
			}
		}
	})

	eg.Go(func() error {
		return tracing.Wait(ctx)
	})

	err = eg.Wait()
	if err != nil {
		log.Error().Err(err).Msgf("error starting application")
		return
	}
	log.Info().Msg("Application is stopped")
}
