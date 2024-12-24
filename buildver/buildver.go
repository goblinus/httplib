// Для корректного использования данной библиотеки можно определить в файле Makefile
// команду для сборки приложения с указанием флага -ldflags. Например:
// - для версии сборки: "-X 'github.com/goblinus/httplib/buildver.Version=${buildVersion}'"
// - для времени сборки: "-X 'github.com/goblinus/collector/httplib/buildver.BuildTime=${buildTime}'"
// - для имени пользователя-автора сборки: "-X 'github.com/goblinus/httplib/buildver.Builder=${builderName}'"
//
// Переменные `buildVersion`, `buildTime`, `builderName` должны быть определены выше команды сборки
package buildver

import (
	"time"
)

var (
	Version   string
	BuildTime string
	Builder   string
)

func NewApplicationMeta() *ApplicationMeta {
	dtValue, err := time.Parse("2006-01-02 15:04:05", BuildTime)
	if err != nil {
		panic(err)
	}

	return &ApplicationMeta{
		Version,
		Builder,
		dtValue,
	}
}

type ApplicationMeta struct {
	version   string
	builder   string
	buildTime time.Time
}

func (am *ApplicationMeta) GetVersion() string {
	return am.version
}

func (am *ApplicationMeta) GetBuilder() string {
	return am.builder
}

func (am *ApplicationMeta) GetBuildTime() string {
	return am.buildTime.Format("2006-01-02 15:04:05")
}
