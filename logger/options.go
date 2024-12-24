package logger

import (
	"fmt"
	"strings"
	"time"

	"github.com/rs/zerolog"
)

const (
	TermColorsRedFormat       = "\u001b[38;5;196m%s\u001b[0m"
	TermColorsBlueFormat      = "\u001b[38;5;21m%s\u001b[0m"
	TermColorsGreenFormat     = "\u001b[38;5;40m%s\u001b[0m"
	TermColorsYellowFormat    = "\u001b[38;5;192m%s\u001b[0m"
	TermColorsCyanFormat      = "\u001b[38;5;63m%s\u001b[0m"
	TermColorsLightGreyFormat = "\u001b[38;5;240m%s\u001b[0m"
)

func TimeFormatter(format string) func(*zerolog.ConsoleWriter) {
	return func(cw *zerolog.ConsoleWriter) {
		cw.TimeFormat = format
	}
}

func TimeLocation(location *time.Location) func(*zerolog.ConsoleWriter) {
	return func(cw *zerolog.ConsoleWriter) {
		cw.TimeLocation = location
	}
}

func DisableColoredConsole(noColor bool) func(*zerolog.ConsoleWriter) {
	return func(cw *zerolog.ConsoleWriter) {
		cw.NoColor = noColor
	}
}

func FormatLevel(format string) func(i interface{}) string {
	return func(i interface{}) string {
		levelName := strings.ToUpper(i.(string))

		switch levelName {
		case "INFO":
			levelName = fmt.Sprintf(TermColorsGreenFormat, levelName)
		case "WARN":
			levelName = fmt.Sprintf(TermColorsYellowFormat, levelName)
		case "DEBUG", "TRACE":
			levelName = fmt.Sprintf(TermColorsCyanFormat, levelName)

			levelName = fmt.Sprintf(TermColorsCyanFormat, levelName)
		default:
			levelName = fmt.Sprintf(TermColorsRedFormat, levelName)
		}

		return fmt.Sprintf(format, levelName)
	}
}

func FormatCaller(format string) func(i interface{}) string {
	return func(i interface{}) string {
		levelName := i.(string)
		levelName = fmt.Sprintf(TermColorsLightGreyFormat, levelName)
		levelName = "[" + levelName + "]"
		return fmt.Sprintf(format, levelName)
	}
}
