package pipelines

import (
	"github.com/AJackTi/go-ecommerce-microservices-example/internal/pkg/logger"
	"github.com/AJackTi/go-ecommerce-microservices-example/internal/pkg/logger/defaultlogger"
)

type config struct {
	logger      logger.Logger
	serviceName string
}

// Option specifies instrumentation configuration options.
type Option interface {
	apply(*config)
}

var defaultConfig = &config{
	serviceName: "app",
	logger:      defaultlogger.GetLogger(),
}

type optionFunc func(*config)

func (o optionFunc) apply(c *config) {
	o(c)
}

func WithLogger(l logger.Logger) Option {
	return optionFunc(func(cfg *config) {
		if cfg.logger != nil {
			cfg.logger = l
		}
	})
}
