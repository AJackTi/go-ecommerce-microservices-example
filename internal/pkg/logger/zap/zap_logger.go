package zap

import (
	"os"
	"time"

	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/AJackTi/go-ecommerce-microservices-example/internal/pkg/config/environment"
	"github.com/AJackTi/go-ecommerce-microservices-example/internal/pkg/constants"
	"github.com/AJackTi/go-ecommerce-microservices-example/internal/pkg/logger"
	"github.com/AJackTi/go-ecommerce-microservices-example/internal/pkg/logger/config"
	"github.com/AJackTi/go-ecommerce-microservices-example/internal/pkg/logger/models"
)

type zapLogger struct {
	level       string
	sugarLogger *zap.SugaredLogger
	logger      *zap.Logger
	logOptions  *config.LogOptions
}

type ZapLogger interface {
	logger.Logger
	InternalLogger() *zap.Logger
	DPanic(args ...interface{})
	DPanicf(template string, args ...interface{})
	Sync() error
}

var loggerLevelMap = map[string]zapcore.Level{
	"debug": zap.DebugLevel,
	"info":  zap.InfoLevel,
	"warn":  zap.WarnLevel,
	"error": zap.ErrorLevel,
	"panic": zap.PanicLevel,
	"fatal": zap.FatalLevel,
}

// NewZapLogger create new zap logger
func NewZapLogger(cfg *config.LogOptions, env environment.Environment) ZapLogger {
	zapLogger := &zapLogger{level: cfg.LogLevel, logOptions: cfg}
	zapLogger.initLogger(env)

	return zapLogger
}

func (l *zapLogger) DPanic(args ...interface{}) {
	l.sugarLogger.DPanic(args...)
}

func (l *zapLogger) DPanicf(template string, args ...interface{}) {
	l.sugarLogger.DPanicf(template, args...)
}

func (l *zapLogger) InternalLogger() *zap.Logger {
	return l.logger
}

func (l *zapLogger) Sync() error {
	go func() {
		err := l.logger.Sync()
		if err != nil {
			l.logger.Error("error while syncing", zap.Error(err))
		}
	}() // nolint: errcheck

	return l.sugarLogger.Sync()
}

func (l *zapLogger) getLoggerLevel() zapcore.Level {
	level, exist := loggerLevelMap[l.level]
	if !exist {
		return zapcore.DebugLevel
	}

	return level
}

func (l *zapLogger) initLogger(env environment.Environment) {
	logLevel := l.getLoggerLevel()

	logWriter := zapcore.AddSync(os.Stdout)

	var encoderCfg zapcore.EncoderConfig
	var encoder zapcore.Encoder

	if env.IsProduction() {
		encoderCfg = zap.NewProductionEncoderConfig()
		encoderCfg.NameKey = "[SERVICE]"
		encoderCfg.TimeKey = "[TIME]"
		encoderCfg.LevelKey = "[LEVEL]"
		encoderCfg.FunctionKey = "[CALLER]"
		encoderCfg.CallerKey = "[LINE]"
		encoderCfg.MessageKey = "[MESSAGE]"
		encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
		encoderCfg.EncodeLevel = zapcore.CapitalLevelEncoder
		encoderCfg.EncodeCaller = zapcore.ShortCallerEncoder
		encoderCfg.EncodeName = zapcore.FullNameEncoder
		encoderCfg.EncodeDuration = zapcore.StringDurationEncoder
		encoder = zapcore.NewJSONEncoder(encoderCfg)
	} else {
		encoderCfg = zap.NewDevelopmentEncoderConfig()
		encoderCfg.NameKey = "[SERVICE]"
		encoderCfg.TimeKey = "[TIME]"
		encoderCfg.LevelKey = "[LEVEL]"
		encoderCfg.FunctionKey = "[CALLER]"
		encoderCfg.CallerKey = "[LINE]"
		encoderCfg.MessageKey = "[MESSAGE]"
		encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
		encoderCfg.EncodeName = zapcore.FullNameEncoder
		encoderCfg.EncodeDuration = zapcore.StringDurationEncoder
		encoderCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder
		encoderCfg.EncodeCaller = zapcore.FullCallerEncoder
		encoderCfg.ConsoleSeparator = " | "
		encoder = zapcore.NewConsoleEncoder(encoderCfg)
	}

	core := zapcore.NewCore(encoder, logWriter, zap.NewAtomicLevelAt(logLevel))

	var options []zap.Option

	if l.logOptions.CallerEnabled {
		options = append(options, zap.AddCaller())
		options = append(options, zap.AddCallerSkip(1))
	}

	logger := zap.New(core, options...)

	if l.logOptions.EnableTracing {
		// add logs as events to tracing
		logger = otelzap.New(logger).Logger
	}

	l.logger = logger
	l.sugarLogger = logger.Sugar()
}

func (z *zapLogger) Configure(cfg func(internalLog interface{})) {
	cfg(z.logger)
}

func (z *zapLogger) Debug(args ...interface{}) {
	z.sugarLogger.Debug(args...)
}

func (z *zapLogger) Debugf(template string, args ...interface{}) {
	z.sugarLogger.Debugf(template, args...)
}

func (z *zapLogger) Debugw(msg string, fields logger.Fields) {
	zapFields := mapToZapFields(fields)
	z.logger.Debug(msg, zapFields...)
}

func (z *zapLogger) Err(msg string, err error) {
	z.logger.Error(msg, zap.Error(err))
}

func (z *zapLogger) Error(args ...interface{}) {
	z.sugarLogger.Error(args...)
}

func (z *zapLogger) Errorf(template string, args ...interface{}) {
	z.sugarLogger.Errorf(template, args...)
}

func (z *zapLogger) Errorw(msg string, fields logger.Fields) {
	zapFields := mapToZapFields(fields)
	z.logger.Error(msg, zapFields...)
}

func (z *zapLogger) Fatal(args ...interface{}) {
	z.sugarLogger.Fatal(args...)
}

func (z *zapLogger) Fatalf(template string, args ...interface{}) {
	z.sugarLogger.Fatalf(template, args...)
}

func (z *zapLogger) GrpcClientInterceptorLogger(method string, req interface{}, reply interface{}, time time.Duration, metaData map[string][]string, err error) {
	z.Info(
		constants.GRPC,
		zap.String(constants.Method, method),
		zap.Any(constants.Request, req),
		zap.Any(constants.Reply, reply),
		zap.Duration(constants.Time, time),
		zap.Any(constants.Metadata, metaData),
		zap.Error(err))
}

func (z *zapLogger) GrpcMiddlewareAccessLogger(method string, time time.Duration, metaData map[string][]string, err error) {
	z.Info(constants.GRPC,
		zap.String(constants.Method, method),
		zap.Duration(constants.Time, time),
		zap.Any(constants.Metadata, metaData),
		zap.Error(err))
}

func (z *zapLogger) Info(args ...interface{}) {
	z.sugarLogger.Info(args...)
}

func (z *zapLogger) Infof(template string, args ...interface{}) {
	z.sugarLogger.Infof(template, args...)
}

func (z *zapLogger) Infow(msg string, fields logger.Fields) {
	zapFields := mapToZapFields(fields)
	z.logger.Info(msg, zapFields...)
}

func (z *zapLogger) LogType() models.LogType {
	return models.Zap
}

func (z *zapLogger) Printf(template string, args ...interface{}) {
	z.sugarLogger.Infof(template, args...)
}

func (z *zapLogger) Warn(args ...interface{}) {
	z.sugarLogger.Warn(args...)
}

func (z *zapLogger) WarnMsg(msg string, err error) {
	z.logger.Warn(msg, zap.String("error", err.Error()))
}

func (z *zapLogger) Warnf(template string, args ...interface{}) {
	z.sugarLogger.Warnf(template, args...)
}

// WithName and logger microservice name
func (z *zapLogger) WithName(name string) {
	z.logger = z.logger.Named(name)
	z.sugarLogger = z.sugarLogger.Named(name)
}

func mapToZapFields(data map[string]interface{}) []zap.Field {
	fields := make([]zap.Field, 0, len(data))

	for key, value := range data {
		field := zap.Field{
			Key:       key,
			Type:      getFieldType(value),
			Interface: value,
		}

		fields = append(fields, field)
	}

	return fields
}

func getFieldType(value interface{}) zapcore.FieldType {
	switch value.(type) {
	case string:
		return zapcore.StringType
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return zapcore.Int64Type
	case bool:
		return zapcore.BoolType
	case float32, float64:
		return zapcore.Float64Type
	case error:
		return zapcore.ErrorType
	default:
		// uses reflection to serialize arbitrary objects, so it can be slow  and allocation-heavy.
		return zapcore.ReflectType
	}
}
