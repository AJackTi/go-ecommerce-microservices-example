package config

import (
	"github.com/iancoleman/strcase"

	"github.com/AJackTi/go-ecommerce-microservices-example/internal/pkg/config"
	"github.com/AJackTi/go-ecommerce-microservices-example/internal/pkg/config/environment"
	"github.com/AJackTi/go-ecommerce-microservices-example/internal/pkg/logger/models"
	"github.com/AJackTi/go-ecommerce-microservices-example/internal/pkg/reflection/typemapper"
)

var optionName = strcase.ToLowerCamel(typemapper.GetGenericTypeNameByT[LogOptions]())

type LogOptions struct {
	LogLevel      string         `mapstructure:"level"`
	LogType       models.LogType `mapstructure:"logType"`
	CallerEnabled bool           `mapstructure:"callerEnabled"`
	EnableTracing bool           `mapstructure:"enableTracing" default:"true"`
}

func ProvideLogConfig(env environment.Environment) (*LogOptions, error) {
	return config.BindConfigKey[*LogOptions](optionName, env)
}
