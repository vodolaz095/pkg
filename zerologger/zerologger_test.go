package zerologger

import (
	"os"
	"testing"

	"github.com/rs/zerolog/log"
)

type testWriter struct {
	T *testing.T
}

func (tw *testWriter) Write(p []byte) (n int, err error) {
	tw.T.Log(string(p))
	return len(p), nil
}

func TestConfigure(t *testing.T) {
	Configure(Log{
		Level:      TraceLevel,
		ToJournald: false,
	},
		&testWriter{T: t},
		os.Stdout,
	)
	log.Trace().Msg("trace")
	log.Debug().Msg("debug")
	log.Info().Msg("info")
	log.Warn().Msg("warn")
	log.Error().Msg("error")
}
