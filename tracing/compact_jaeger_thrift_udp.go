package tracing

import (
	"github.com/rs/zerolog/log"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
)

type UDPConfig struct {
	// Host - sets hostname of Jaeger agent, overrides environment value of OTEL_EXPORTER_JAEGER_AGENT_HOST.
	// Default value is `localhost`
	Host string `yaml:"host" validate:"hostname"`
	// Port - sets port where Jaeger agent listens, overrides environment value of OTEL_EXPORTER_JAEGER_AGENT_PORT.
	// Default value is `6831`
	Port string `yaml:"port" validate:"gte=0,lte=65535"`
	// Ratio sets percent of spans to record, where 1 - means every span is recorded, 0 - no spans recorded and .05 means only 5% of spans are recorded
	Ratio float64 `yaml:"ratio" validate:"required,lte=1,gte=0"`
}

func ConfigureUDP(cfg UDPConfig, extraAttributes ...attribute.KeyValue) (err error) {
	if cfg.Ratio == 0 {
		log.Debug().Msgf("Tracing disabled")
		return nil
	}
	var host, port string
	var opts []jaeger.AgentEndpointOption
	if cfg.Host != "" {
		opts = append(opts, jaeger.WithAgentHost(cfg.Host))
		host = cfg.Host
	} else {
		host = loadFromEnv("OTEL_EXPORTER_JAEGER_AGENT_HOST", "localhost")
	}
	if cfg.Port != "" {
		opts = append(opts, jaeger.WithAgentPort(cfg.Port))
		port = cfg.Port
	} else {
		port = loadFromEnv("OTEL_EXPORTER_JAEGER_AGENT_PORT", "6831")
	}
	log.Debug().Msgf("Sending traces using compact jaeger thrift protocol via udp into %s:%s.", host, port)
	exp, err = jaeger.New(jaeger.WithAgentEndpoint(opts...))
	if err != nil {
		return err
	}
	return makeProvider(cfg.Ratio, extraAttributes...)
}
