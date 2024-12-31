// Для корректного использования данного пакета определить в файле Makefile
package buildmeta

import (
	"time"
)

const dateTimeFormat = "2006-01-02 15:04:05"

func NewBuildMeta(version, release, builder, buildDateTime string) *BuildMeta {
	dateTime, err := time.Parse(dateTimeFormat, buildDateTime)
	if err != nil {
		panic(err)
	}

	return &BuildMeta{
		version:       version,
		release:       release,
		builder:       builder,
		buildDateTime: dateTime,
	}
}

type BuildMeta struct {
	version       string
	release       string
	builder       string
	buildDateTime time.Time
}

func (m *BuildMeta) GetVersion() string {
	return m.version
}

func (m *BuildMeta) GetRelease() string {
	return m.release
}

func (m *BuildMeta) GetBuilder() string {
	return m.builder
}

func (m *BuildMeta) GetBuildTime() string {
	return m.buildDateTime.Format(dateTimeFormat)
}
