package logger

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func InitGlobalLogger() {
	log.Logger = zerolog.New(zerolog.ConsoleWriter{
		Out:        os.Stderr,
		TimeFormat: time.RFC3339,
		NoColor:    false,
		FormatLevel: func(i interface{}) string {
			return strings.ToUpper(fmt.Sprintf("[%-6s]", i))
		},
		FormatFieldName: func(i interface{}) string {
			return fmt.Sprintf("%s:", i)
		},
		FormatFieldValue: func(i interface{}) string {
			return strings.ToUpper(fmt.Sprintf("[%s]", i))
		},
	}).With().Caller().Timestamp().Logger()
}
