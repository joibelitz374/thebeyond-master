package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func New(env, level string) (logger *zap.Logger) {
	switch env {
	case "production":
		logger = zap.New(zapcore.NewCore(
			zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
			zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&lumberjack.Logger{
				Filename:   "/var/log/thebeyond/app.log",
				MaxSize:    100, // MB
				MaxBackups: 10,
				MaxAge:     30, // days
				Compress:   true,
			})),
			zap.NewAtomicLevelAt(logLevel(level)),
		))
	default:
		cfg := zap.NewDevelopmentConfig()
		cfg.OutputPaths = []string{"stdout"}
		cfg.ErrorOutputPaths = []string{"stderr"}
		cfg.Level = zap.NewAtomicLevelAt(logLevel(level))
		log, err := cfg.Build()
		if err != nil {
			panic(err)
		}
		logger = log
	}

	return logger
}

func logLevel(level string) (logLevel zapcore.Level) {
	switch level {
	case "debug":
		logLevel = zap.DebugLevel
	case "info":
		logLevel = zap.InfoLevel
	case "warn":
		logLevel = zap.WarnLevel
	case "error":
		logLevel = zap.ErrorLevel
	default:
		logLevel = zap.InfoLevel
	}
	return logLevel
}
