package tracing

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"
	"go.opentelemetry.io/otel/exporters/jaeger"
)

const defaultTimeout = 10 * time.Second

var exp *jaeger.Exporter

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
