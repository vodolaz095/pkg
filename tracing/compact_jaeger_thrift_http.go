package tracing

import (
	"github.com/rs/zerolog/log"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
)

// HTTPConfig is used to fine tune jaeger exporter to deliver spans via compact thrift protocol to collector http endpoint
type HTTPConfig struct {
	Endpoint string  `yaml:"endpoint"`
	Username string  `yaml:"username"`
	Password string  `yaml:"password"`
	Ratio    float64 `yaml:"ratio" validate:"lte=1,gte=0"`
	// TraceProviderOptions allows to add extra tracer provider options like custom sampler and so on
	TraceProviderOptions []tracesdk.TracerProviderOption `yaml:"-"`
}

// ConfigureHTTP fine tunes jaeger exporter to deliver spans via compact thrift protocol to collector http endpoint
func ConfigureHTTP(cfg HTTPConfig, extraAttributes ...attribute.KeyValue) (err error) {
	if cfg.Ratio == 0 {
		log.Debug().Msgf("Tracing disabled")
		return nil
	}
	opts := make([]jaeger.CollectorEndpointOption, 0)
	if cfg.Endpoint != "" {
		opts = append(opts, jaeger.WithEndpoint(cfg.Endpoint))
		log.Debug().Msgf("Sending traces using compact jaeger thrift protocol via http into %s...", cfg.Endpoint)
	} else {
		where := loadFromEnv("OTEL_EXPORTER_JAEGER_ENDPOINT", "http://localhost:14268/api/traces")
		log.Debug().Msgf("Sending traces using compact jaeger thrift protocol via http into %s...", where)
	}
	if cfg.Username != "" {
		opts = append(opts, jaeger.WithUsername(cfg.Username))
	}
	if cfg.Password != "" {
		opts = append(opts, jaeger.WithPassword(cfg.Password))
	}
	// export via compact thrift protocol over http
	exp, err = jaeger.New(jaeger.WithCollectorEndpoint(opts...))
	if err != nil {
		return err
	}
	return makeProvider(cfg.TraceProviderOptions, cfg.Ratio, extraAttributes...)
}
