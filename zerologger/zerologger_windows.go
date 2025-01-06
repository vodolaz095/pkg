//go:build windows
// +build windows

package zerologger

import (
	"io"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func Configure(params Log, writers ...io.Writer) {
	var outputsEnabled []io.Writer
	outputsEnabled = append(outputsEnabled, zerolog.ConsoleWriter{
		Out:        os.Stdout, // https://12factor.net/ru/logs
		TimeFormat: "15:04:05",
	})
	outputsEnabled = append(outputsEnabled, writers...)
	zerolog.CallerMarshalFunc = callerMarshalFunc
	sink := zerolog.New(zerolog.MultiLevelWriter(outputsEnabled...)).
		With().Timestamp().Caller().
		Logger().Level(ExtractZerologLevel(params.Level))
	log.Logger = sink
	return
}
