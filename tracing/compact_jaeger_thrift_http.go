package tracing

import (
	"github.com/rs/zerolog/log"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
)

type HTTPConfig struct {
	Endpoint string  `yaml:"endpoint" validate:"url"`
	Username string  `yaml:"username"`
	Password string  `yaml:"password"`
	Ratio    float64 `yaml:"ratio" validate:"lte=1,gte=0"`
}

func ConfigureHTTP(cfg HTTPConfig, extraAttributes ...attribute.KeyValue) (err error) {
	if cfg.Ratio == 0 {
		log.Debug().Msgf("Tracing disabled")
		return nil
	}
	opts := make([]jaeger.CollectorEndpointOption, 0)
	opts = append(opts, jaeger.WithEndpoint(cfg.Endpoint))
	if cfg.Username != "" {
		opts = append(opts, jaeger.WithUsername(cfg.Username))
	}
	if cfg.Password != "" {
		opts = append(opts, jaeger.WithPassword(cfg.Password))
	}
	log.Debug().Msgf("Sending traces using compact jaeger thrift protocol via http into %s...", cfg.Endpoint)
	// export via compact thrift protocol over http
	exp, err = jaeger.New(jaeger.WithCollectorEndpoint(opts...))
	if err != nil {
		return err
	}
	return makeProvider(cfg.Ratio, extraAttributes...)
}
