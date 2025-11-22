package tracing

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel/attribute"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
)

// Config is universal config being used to tune tracing
type Config struct {
	// Protocol sets how we send spans to jaeger - over udp or over http
	Protocol string `yaml:"protocol" validate:"required,oneof=udp http UDP HTTP otlp_http OTLP_HTTP"`

	/*
		Configuration for Jaeger exporter to send spans to a Jaeger agent over compact thrift protocol over UDP
	*/

	// Host - sets hostname of Jaeger agent, overrides environment value of OTEL_EXPORTER_JAEGER_AGENT_HOST.
	// Default value is `localhost`
	Host string `yaml:"host"`
	// Port - sets port where Jaeger agent listens, overrides environment value of OTEL_EXPORTER_JAEGER_AGENT_PORT.
	// Default value is `6831`
	Port string `yaml:"port"`

	/*
		Configuration for Jaeger exporter to use full URL to the Jaeger HTTP Thrift collector.
	*/

	// Endpoint is the URL for the Jaeger collector that spans are sent to,
	// overrides value of environment variable OTEL_EXPORTER_JAEGER_ENDPOINT.
	// Default value is `http://localhost:14268/api/traces`
	Endpoint string `yaml:"endpoint"`
	// Username used for basic authorization to access Jaeger collector. Setting value overrides environment
	// variable OTEL_EXPORTER_JAEGER_USER. Default is empty
	Username string `yaml:"username"`
	// Password  used for basic authorization to access Jaeger collector. Setting value overrides environment
	// variable OTEL_EXPORTER_JAEGER_PASSWORD. Default is empty
	Password string `yaml:"password"`

	/*
		Configuration for Jaeger exporter to use full URL to  HTTP OTLP collector.
	*/

	// OTLPEndpoint sets HTTP OTLP collector url, overrides value of environment variable of OTEL_EXPORTER_OTLP_TRACES_ENDPOINT
	// Default is  "https://localhost:4318/v1/traces". See full documentation for environment options supported:
	// https://pkg.go.dev/go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp#pkg-overview
	OTLPEndpoint string `yaml:"otlp_endpoint"`

	// Insecure allows insecure connections over HTTP protocol where connection via HTTPS are preferred
	Insecure bool `yaml:"insecure"`
	// Ratio sets percent of spans to record, where 1 - means every span is recorded, 0 - no spans recorded and .05 means only 5% of spans are recorded
	Ratio float64 `yaml:"ratio" validate:"required,lte=1,gte=0"`

	/*
	   Extra trace provider options like samplers and so on applied for all exporters
	*/
	// TraceProviderOptions allows to add extra tracer provider options like custom sampler and so on
	TraceProviderOptions []tracesdk.TracerProviderOption `yaml:"-"`
}

func (c *Config) String() string {
	switch c.Protocol {
	case "otlp_http", "OTLP_HTTP":
		return fmt.Sprintf("Sending spans via OTLP HTTP into %s with ratio %.0f%%...",
			c.OTLPEndpoint, 100*c.Ratio)
	case "udp", "UDP":
		return fmt.Sprintf("Sending spans over compact thrift protocol over udp into %s:%s with ratio %.0f%%...",
			c.Host, c.Port, 100*c.Ratio)
	case "http", "HTTP":
		return fmt.Sprintf("Sending spans over compact thrift protocol over http into %s with ratio %.0f%%...",
			c.Endpoint, 100*c.Ratio)
	default:
		return fmt.Sprintf("Unknown protocol %s", c.Protocol)
	}
}

// Start starts telemetry exporter
func Start(cfg Config, extraAttributes ...attribute.KeyValue) (err error) {
	return StartWithContext(context.Background(), cfg, extraAttributes...)
}

// StartWithContext starts telemetry exporter with context provided
func StartWithContext(ctx context.Context, cfg Config, extraAttributes ...attribute.KeyValue) (err error) {
	switch cfg.Protocol {
	case "otlp_http", "OTLP_HTTP":
		return ConfigureOTLPoverHTTP(ctx, OTLPoverHTTPConfig{
			Endpoint:             cfg.OTLPEndpoint,
			Compression:          true,
			Ratio:                cfg.Ratio,
			Insecure:             cfg.Insecure,
			TraceProviderOptions: cfg.TraceProviderOptions,
		}, extraAttributes...)
	case "udp", "UDP":
		return ConfigureUDP(UDPConfig{
			Host:                 cfg.Host,
			Port:                 cfg.Port,
			Ratio:                cfg.Ratio,
			TraceProviderOptions: cfg.TraceProviderOptions,
		}, extraAttributes...)
	case "http", "HTTP":
		return ConfigureHTTP(HTTPConfig{
			Endpoint:             cfg.Endpoint,
			Username:             cfg.Username,
			Password:             cfg.Password,
			TraceProviderOptions: cfg.TraceProviderOptions,
		}, extraAttributes...)
	default:
		return fmt.Errorf("unknowwn protocol: %s", cfg.Protocol)
	}
}
