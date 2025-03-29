package environment

import "github.com/AJackTi/go-ecommerce-microservices-example/internal/pkg/constants"

type Environment string

var (
	Development = Environment(constants.Dev)
	Test        = Environment(constants.Test)
	Production  = Environment(constants.Production)
)
