pkg
=====================================
[![Go](https://github.com/vodolaz095/pkg/actions/workflows/go.yml/badge.svg)](https://github.com/vodolaz095/pkg/actions/workflows/go.yml)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/vodolaz095/pkg)](https://pkg.go.dev/github.com/vodolaz095/pkg?tab=doc)
[![Go Report Card](https://goreportcard.com/badge/github.com/vodolaz095/pkg)](https://goreportcard.com/report/github.com/vodolaz095/pkg)

Various functions copy-pasted between different projects.

See example: [example.go](example%2Fexample.go)

cryptorand
=======================================
Generate cryptographically secure random strings with alphabet provided

Documentation: https://pkg.go.dev/github.com/vodolaz095/pkg/cryptorand

date
=======================================
Date and time helpers - see [moments_test.go](date%2Fmoments_test.go)

Documentation: https://pkg.go.dev/github.com/vodolaz095/pkg/date


healthcheck
=======================================
Systemd compatible healthcheck.

Documentation: https://www.freedesktop.org/software/systemd/man/latest/sd_notify.html

Usage example: 
https://github.com/vodolaz095/stocks_broadcaster/blob/a03cf70efc1e333e959f58bd295aa2701cca37c8/main.go#L131-L160

math
=======================================
Various generic mathematical functions copy-pasted between different projects.

Documentation: https://pkg.go.dev/github.com/vodolaz095/pkg/math

stopper
=======================================
Make global application context which can be terminated by signals

Documentation: https://pkg.go.dev/github.com/vodolaz095/pkg/stopper

tracing
=======================================
Opinionated way to configure OpenTelemetry with `jaegertracing/all-in-one` started with docker compose like this

```yaml

version: "3.11"

volumes:
  jaeger_temp_old:
  jaeger_temp:

services:
  jaeger_old:
    container_name: jaeger
    image: docker.io/jaegertracing/all-in-one:1.57.0
    volumes:
      - jaeger_temp_old:/tmp
    ports:
      - "16686:16686/tcp" # webui is listening
      - "14268:14268/tcp" # accepting spans in compact jaeger thrift format over http
      - "6831:6831/udp" # accepting spans in compact jaeger thrift format over udp

  jaeger:
    container_name: jaeger
    image: docker.io/jaegertracing/all-in-one:1.67.0
    volumes:
      - jaeger_temp:/tmp
    ports:
      - "16686:16686/tcp" # webui is listening
      - "14268:14268/tcp" # accepting spans in compact jaeger thrift format over http
      - "4318:4318/tcp" # accepting spans in OTLP format over http


```
See example: [example.go](example%2Fexample.go)

zerologger
=======================================
Opinionated way to configure zerolog with sane defaults

Documentation: https://pkg.go.dev/github.com/vodolaz095/pkg/zerologger
Usage example: [zerologger_test.go](zerologger%2Fzerologger_test.go)
