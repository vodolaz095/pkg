package stopper

import (
	"context"
	"os/signal"
	"syscall"
)

// NewWithContext makes child context which is terminated when process receives stop signals
func NewWithContext(parent context.Context) (ctx context.Context, cancel context.CancelFunc) {
	return signal.NotifyContext(parent,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
		syscall.SIGABRT)
}

// New makes context which is terminated when process receives stop signals
func New() (ctx context.Context, cancel context.CancelFunc) {
	return NewWithContext(context.Background())
}
