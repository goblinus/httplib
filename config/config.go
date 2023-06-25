package config

import (
	"fmt"
	"net/url"
	"reflect"
	"strings"

	"github.com/mitchellh/mapstructure"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

type (
	appConfig interface {
		AppName() string
		SetDefaults(v *viper.Viper)
	}
)

var (
	v           *viper.Viper = viper.NewWithOptions(viper.KeyDelimiter("_"), viper.EnvKeyReplacer(strings.NewReplacer(".", "_")))
	DecodeHooks              = viper.DecodeHook(mapstructure.ComposeDecodeHookFunc(
		mapstructure.StringToSliceHookFunc(","),
		mapstructure.StringToTimeDurationHookFunc(),
		StringToURLHook,
	))
)

// Load загружает конфигурации сервиса в переданную структуру
func Load(cfg appConfig) error {
	v.AutomaticEnv()
	cfg.SetDefaults(v)
	setDefaultPaths(v, cfg.AppName())
	if err := v.ReadInConfig(); err != nil {
		log.Debug().Msg("cannot read config from file, trying read environment's variables")
		if err := v.Unmarshal(cfg); err != nil {
			return err
		}
	}

	return nil
}

// Viper возвращает инстанс viper
func Viper() *viper.Viper {
	return v
}

func StringToURLHook(f, t reflect.Type, data interface{}) (interface{}, error) {
	if f.Kind() != reflect.String {
		return data, nil
	}

	if t != reflect.TypeOf(url.URL{}) {
		return data, nil
	}

	return url.Parse(data.(string))
}

// setDefaultPaths устанавливает пути по которым может располагаться файл конфигурации
func setDefaultPaths(v *viper.Viper, appName string) {
	v.AddConfigPath(fmt.Sprintf("/etc/%s/", appName))
	v.AddConfigPath(fmt.Sprintf("$HOME/config/%s/", appName))
	v.SetConfigName("configs")
	v.SetConfigType("yaml")
}
