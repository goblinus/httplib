package config

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setup() (*HTTPConfig, map[string]bool) {
	conf := &HTTPConfig{
		Addr: "localhost",
		Port: 8081,
		ExcludedPath: []string{
			"/ready",
			"/health",
			"/metrics",
		},
	}

	excludedPaths := map[string]bool{
		"/ready":   true,
		"/health":  true,
		"/metrics": true,
	}

	return conf, excludedPaths
}

func TestConfig(t *testing.T) {
	conf, excludedPaths := setup()

	assert.Equal(t, "localhost", conf.Addr, fmt.Sprintf("address should be localhost, got %s", conf.Addr))
	assert.Equal(t, 8081, conf.Port, "port should be 8081, got %d", conf.Port)

	assert.Equal(t, 3, len(conf.ExcludedPath), fmt.Sprintf("should be 3 excluded paths, got %d", len(conf.ExcludedPath)))
	for _, path := range conf.ExcludedPath {
		if _, ok := excludedPaths[path]; !ok {
			t.Errorf("%s is not in paths list should be", path)
		}
	}
}
