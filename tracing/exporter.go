package tracing

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"
	"go.opentelemetry.io/otel/sdk/trace"
)

const defaultTimeout = 10 * time.Second

var exp trace.SpanExporter

// Wait allows span exporter to inhibit application shutdown until it sends all traces.
func Wait(ctx context.Context) (err error) {
	if exp != nil {
		<-ctx.Done()
		shutdownContext, cancel := context.WithTimeout(context.Background(), defaultTimeout)
		defer cancel()
		err = exp.Shutdown(shutdownContext)
		if err != nil {
			log.Error().Err(err).Msgf("Error saving telemetry spans: %s", err)
			return err
		}
		log.Debug().Msgf("All spans are saved")
		return nil
	}
	return nil
}

// Shutdown forces exporter to save all spans and turn off.
func Shutdown(ctx context.Context) (err error) {
	if exp != nil {
		shutdownContext, cancel := context.WithTimeout(ctx, defaultTimeout)
		defer cancel()
		err = exp.Shutdown(shutdownContext)
		if err != nil {
			log.Error().Err(err).Msgf("Error saving telemetry spans: %s", err)
			return err
		}
		log.Debug().Msgf("All spans are saved")
		return nil
	}
	return nil
}
