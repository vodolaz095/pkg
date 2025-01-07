//go:build !windows
// +build !windows

package zerologger

import (
	"io"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/journald"
	"github.com/rs/zerolog/log"
)

// Configure tunes global logger of zerolog
func Configure(params Log, writers ...io.Writer) {
	var outputsEnabled []io.Writer
	if params.ToJournald {
		outputsEnabled = append(outputsEnabled, journald.NewJournalDWriter())
	} else {
		outputsEnabled = append(outputsEnabled, zerolog.ConsoleWriter{
			Out:        os.Stdout, // https://12factor.net/ru/logs
			TimeFormat: "15:04:05",
		})
	}
	outputsEnabled = append(outputsEnabled, writers...)
	zerolog.CallerMarshalFunc = callerMarshalFunc
	sink := zerolog.New(zerolog.MultiLevelWriter(outputsEnabled...)).
		With().Timestamp().Caller().
		Logger().Level(ExtractZerologLevel(params.Level))
	log.Logger = sink
	return
}
