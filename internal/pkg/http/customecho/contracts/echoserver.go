package contracts

import (
	"context"

	"github.com/labstack/echo/v4"

	"github.com/AJackTi/go-ecommerce-microservices-example/internal/pkg/http/customecho/config"
	"github.com/AJackTi/go-ecommerce-microservices-example/internal/pkg/logger"
)

type EchoHttpServer interface {
	RunHttpServer(configEcho ...func(echo *echo.Echo)) error
	GracefulShutdown(ctx context.Context) error
	ApplyVersioningFromHeader()
	GetEchoInstance() *echo.Echo
	Logger() logger.Logger
	Cfg() *config.EchoHttpOptions
	SetupDefaultMiddlewares()
	RouteBuilder() *RouteBuilder
	AuthMiddlewares(middlewares ...echo.MiddlewareFunc)
	ConfigGroup(groupName string, groupFunc func(group *echo.Group))
}
