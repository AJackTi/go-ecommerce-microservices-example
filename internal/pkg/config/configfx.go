package config

import (
	"go.uber.org/fx"

	"github.com/AJackTi/go-ecommerce-microservices-example/internal/pkg/config/environment"
)

// Module provided to fxlog
// https://uber-go.github.io/fx/modules.html
var Module = fx.Module(
	"configfx",
	fx.Provide(func() environment.Environment {
		return environment.ConfigAppEnv()
	}),
)

var ModuleFunc = func(e environment.Environment) fx.Option {
	return fx.Module(
		"configfx",
		fx.Provide(func() environment.Environment {
			return environment.ConfigAppEnv(e)
		}),
	)
}
