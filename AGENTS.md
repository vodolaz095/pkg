# AGENTS.md

## Project Overview

This repository contains various functions that are commonly copy-pasted between different projects. It includes packages for cryptographically secure random string generation, date and time helpers, healthcheck utilities, mathematical functions, application context management, tracing, and logging.

## Directory Structure

The repository is organized into several directories, each containing related functionality:

- `example`: Contains an example usage of the package.
- `math`: Includes generic mathematical functions.
- `cryptorand`: Provides functions for generating cryptographically secure random strings.
- `date`: Contains date and time helper functions.
- `healthcheck`: Includes utilities for systemd-compatible healthchecks.
- `stopper`: Provides a way to manage global application context that can be terminated by signals.
- `tracing`: Contains tools for configuring OpenTelemetry with Jaeger.
- `zerologger`: Provides an opinionated way to configure Zerolog with sane defaults.

## Build, Test, and Lint Commands

The repository uses Makefile for various commands:

- `make deps`: Downloads and verifies Go module dependencies.
- `make tools`: Checks the availability of required tools like `golint` and `govulncheck`.
- `make fix`: Runs `go fix ./...` to fix code style packages to make it more modern.
- `make lint`: Runs `gofmt`, `golint`, and `go vet` to ensure code quality.
- `make test`: Runs all tests in the repository.
- `make cover`: Runs tests with coverage.
- `make vuln`: Checks for vulnerabilities using `govulncheck`.
- `make run`: Runs the example program.

## Code Conventions

- **Naming**: Functions and variables follow the Go convention of using camelCase.
- **Style**: The code adheres to the Go standard style, as enforced by `gofmt`.
- **Testing**: Tests are written using the `testing` package and follow a similar structure to the production code.

## Gotchas and Non-Obvious Patterns

- The `tracing` package provides an opinionated way to configure OpenTelemetry with Jaeger, which can be started using Docker Compose.
- The `zerologger` package provides a sane default configuration for Zerolog.

## Skills

- `run-tests`: Skill to run unit tests using `make test`, with additional options for coverage (`make cover`) and linting before tests (`make check`). 
  Requires tools and dependencies to be set up using `make tools` and `make deps` respectively.
- `vuln-scan`: Skill to scan for vulnerabilities using `make vuln` and fix them by updating dependencies with `go get` according to semver.
  Requires `govulncheck` to be installed and available in PATH.
- `lint`: Skill to run linting on the Go codebase and automatically fix any issues found.
  Requires tools (go, gopls, staticcheck) to be installed and available in PATH.
  Usage:
  1. Verify required tools are available
  2. Run `make lint` to identify issues
  3. Analyze lint output and apply fixes
  4. Verify fixes resolved all issues
