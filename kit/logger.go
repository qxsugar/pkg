package kit

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

// NewProductionConfig returns a zap configuration optimized for production use.
// It uses JSON encoding and Unix timestamp format.
func NewProductionConfig() zap.Config {
	config := zap.NewProductionConfig()
	config.EncoderConfig.EncodeTime = func(time time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendInt64(time.Unix())
	}
	return config
}

// NewDevelopmentConfig returns a zap configuration optimized for development.
// It uses console encoding and RFC3339 timestamp format.
func NewDevelopmentConfig() zap.Config {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeTime = func(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(t.Format(time.RFC3339))
	}
	return config
}

// MustProduction creates a production logger and panics if it fails.
func MustProduction() *zap.Logger {
	return zap.Must(NewProductionConfig().Build())
}

// MustDevelopment creates a development logger and panics if it fails.
func MustDevelopment() *zap.Logger {
	return zap.Must(NewDevelopmentConfig().Build())
}
