package logger

import (
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func InitGlobalLogger() {
	location, err := time.LoadLocation("Europe/Moscow")
	if err != nil {
		panic(err)
	}

	writer := zerolog.NewConsoleWriter(
		TimeLocation(location),
		TimeFormatter("2006-01-02 15:04:05"),
		DisableColoredConsole(false),
	)

	writer.FormatLevel = FormatLevel("[%s]")
	writer.FormatCaller = FormatCaller("%s")

	log.Logger = zerolog.New(writer).
		With().
		Caller().
		Timestamp().
		Logger()
}
