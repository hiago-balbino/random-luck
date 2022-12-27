package logger

import (
	"sync"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	// Logger is an exportable variable to be used in log output.
	Logger *zap.Logger
	once   sync.Once
)

// GetLogger is a function to initialize the zap logger once and return.
func GetLogger() *zap.Logger {
	once.Do(func() {
		config := zap.Config{
			Level:            zap.NewAtomicLevelAt(getLogLevel()),
			Encoding:         "json",
			OutputPaths:      []string{"stdout"},
			ErrorOutputPaths: []string{"stderr"},
			EncoderConfig: zapcore.EncoderConfig{
				MessageKey:  "message",
				LevelKey:    "level",
				TimeKey:     "time",
				EncodeTime:  zapcore.ISO8601TimeEncoder,
				EncodeLevel: zapcore.CapitalLevelEncoder,
			},
		}

		Logger = zap.Must(config.Build())
	})

	return Logger
}

func getLogLevel() zapcore.Level {
	logLevel := viper.GetString("LOG_LEVEL")

	switch logLevel {
	case "DEBUG":
		return zap.DebugLevel
	case "INFO":
		return zap.InfoLevel
	case "WARN":
		return zap.WarnLevel
	case "ERROR":
		return zap.ErrorLevel
	default:
		return zap.ErrorLevel
	}
}
