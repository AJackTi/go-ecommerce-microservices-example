package defaultlogger

import (
	"os"

	"github.com/AJackTi/go-ecommerce-microservices-example/internal/pkg/constants"
	"github.com/AJackTi/go-ecommerce-microservices-example/internal/pkg/logger"
	"github.com/AJackTi/go-ecommerce-microservices-example/internal/pkg/logger/config"
	logrous "github.com/AJackTi/go-ecommerce-microservices-example/internal/pkg/logger/logrus"
	"github.com/AJackTi/go-ecommerce-microservices-example/internal/pkg/logger/models"
	"github.com/AJackTi/go-ecommerce-microservices-example/internal/pkg/logger/zap"
)

var l logger.Logger

func initLogger() {
	logType := os.Getenv("LogConfig_LogType")

	switch logType {
	case "Zap", "":
		l = zap.NewZapLogger(
			&config.LogOptions{LogType: models.Zap, CallerEnabled: false},
			constants.Dev,
		)
	case "Logrus":
		l = logrous.NewLogrusLogger(&config.LogOptions{LogType: models.Logrus, CallerEnabled: false},
			constants.Dev,
		)
	default:
	}
}

func GetLogger() logger.Logger {
	if l == nil {
		initLogger()
	}

	return l
}
