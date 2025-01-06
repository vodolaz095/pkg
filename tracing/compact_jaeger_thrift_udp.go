package tracing

import (
	"fmt"
	"strings"

	"github.com/rs/zerolog/log"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
)

type UDPConfig struct {
	Endpoint string  `yaml:"endpoint" validate:"hostname_port"`
	Ratio    float64 `yaml:"ratio" validate:"lte=1,gte=0"`
}

func ConfigureUDP(cfg UDPConfig, extraAttributes ...attribute.KeyValue) (err error) {
	if cfg.Ratio == 0 {
		log.Debug().Msgf("Tracing disabled")
		return nil
	}
	if cfg.Endpoint == "" {
		host := loadFromEnv("OTEL_EXPORTER_JAEGER_AGENT_HOST", "localhost")
		port := loadFromEnv("OTEL_EXPORTER_JAEGER_AGENT_PORT", "6831")
		log.Debug().Msgf("Sending traces using compact jaeger thrift protocol via udp into %s:%s.", host, port)
		exp, err = jaeger.New(jaeger.WithAgentEndpoint())
	} else {
		parts := strings.Split(cfg.Endpoint, ":")
		if len(parts) != 2 {
			return fmt.Errorf("malformed endpoint: %s", cfg.Endpoint)
		}
		log.Debug().Msgf("Sending traces using compact jaeger thrift protocol via udp into %s...", cfg.Endpoint)
		// export via compact thrift protocol over upd - important
		exp, err = jaeger.New(jaeger.WithAgentEndpoint(
			jaeger.WithAgentHost(parts[0]),
			jaeger.WithAgentPort(parts[1]),
		))
	}
	if err != nil {
		return err
	}
	return makeProvider(cfg.Ratio, extraAttributes...)
}
