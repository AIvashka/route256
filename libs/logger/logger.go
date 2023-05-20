package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"sync"
)

var (
	logger *zap.Logger
	once   sync.Once
)

// Init initializes the zap logger
func Init() {
	once.Do(func() {
		config := zap.Config{
			Level:       zap.NewAtomicLevelAt(zap.DebugLevel), // Log level
			Development: false,                                // Development Mode
			Sampling: &zap.SamplingConfig{ // Sampling Configuration
				Initial:    100,
				Thereafter: 100,
			},
			Encoding:         "json",                           // Encoding (json or console)
			EncoderConfig:    zap.NewProductionEncoderConfig(), // Encoder configuration
			OutputPaths:      []string{"stdout"},               // Output paths
			ErrorOutputPaths: []string{"stderr"},               // Error output paths
		}

		config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

		var err error
		logger, err = config.Build()
		if err != nil {
			panic(err)
		}
	})
}

// GetLogger returns a zap logger instance
func GetLogger() *zap.Logger {
	if logger == nil {
		Init()
	}
	return logger
}
