package pipelines

import (
	"github.com/AJackTi/go-ecommerce-microservices-example/internal/pkg/logger"
	"github.com/AJackTi/go-ecommerce-microservices-example/internal/pkg/logger/defaultlogger"
)

type config struct {
	logger      logger.Logger
	serviceName string
}

var defaultConfig = &config{
	serviceName: "app",
	logger:      defaultlogger.GetLogger(),
}

// Option specifies instrumentation configuration options.
type Option interface {
	apply(*config)
}
