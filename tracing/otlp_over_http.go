package tracing

import (
	"context"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
)

// OTLPoverHTTPConfig is used to fine http otlp exporter to deliver spans via OTLP over HTTP protocol
type OTLPoverHTTPConfig struct {
	Endpoint    string                 `yaml:"endpoint"`
	Compression bool                   `yaml:"compression"`
	Ratio       float64                `yaml:"ratio" validate:"lte=1,gte=0"`
	Headers     map[string]string      `yaml:"headers"`
	Opts        []otlptracehttp.Option `yaml:"-"`
}

// ConfigureOTLPoverHTTP fine tunes OTLP HTTP exporter to deliver spans via OTLP over HTTP protocol to collector http endpoint
func ConfigureOTLPoverHTTP(ctx context.Context, cfg OTLPoverHTTPConfig, extraAttributes ...attribute.KeyValue) (err error) {
	var opts []otlptracehttp.Option
	if cfg.Endpoint != "" {
		opts = append(opts, otlptracehttp.WithEndpointURL(cfg.Endpoint))
	}
	if cfg.Compression {
		opts = append(opts, otlptracehttp.WithCompression(otlptracehttp.GzipCompression))
	}
	if len(cfg.Headers) > 0 {
		opts = append(opts, otlptracehttp.WithHeaders(cfg.Headers))
	}
	opts = append(opts, cfg.Opts...)
	exp, err = otlptracehttp.New(ctx, opts...)
	if err != nil {
		return err
	}
	return makeProvider(cfg.Ratio, extraAttributes...)
}
