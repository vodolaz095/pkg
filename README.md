pkg
=====================================
[![Go](https://github.com/vodolaz095/pkg/actions/workflows/go.yml/badge.svg)](https://github.com/vodolaz095/pkg/actions/workflows/go.yml)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/vodolaz095/pkg)](https://pkg.go.dev/github.com/vodolaz095/pkg?tab=doc)
[![Go Report Card](https://goreportcard.com/badge/github.com/vodolaz095/pkg)](https://goreportcard.com/report/github.com/vodolaz095/pkg)

Various functions copy-pasted between different projects.

cryptorand
=======================================
Generate cryptographically secure random strings with alphabet provided

date
=======================================
Date and time helpers - see [moments_test.go](date%2Fmoments_test.go)

math
=======================================
Various generic mathematical functions copy-pasted between different projects.

stopper
=======================================
Make global application context which can be terminated by signals

tracing
=======================================
Opinionated way to configure OpenTelemetry with `jaegertracing/all-in-one` started like with docker compose like this

```yaml

version: "3.11"

services:
  jaeger:
    container_name: jaeger
    image: docker.io/jaegertracing/all-in-one:1.37
    ports:
      - "16686:16686/tcp" # webui is listening
      - "14268:14268/tcp" # accepting spans in compact jaeger thrift format over http
      - "6831:6831/udp" # accepting spans in compact jaeger thrift format over udp

```

zerologger
=======================================
Opinionated way to configure zerolog with sane defaults
