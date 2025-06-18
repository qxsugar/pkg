package kit

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

func NewProductionConfig() zap.Config {
	config := zap.NewProductionConfig()
	config.EncoderConfig.EncodeTime = func(time time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendInt64(time.Unix())
	}
	return config
}

func NewDevelopmentConfig() zap.Config {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeTime = func(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(t.Format(time.RFC3339))
	}
	return config
}

func MustProduction() *zap.Logger {
	return zap.Must(NewProductionConfig().Build())
}

func MustDevelopment() *zap.Logger {
	return zap.Must(NewDevelopmentConfig().Build())
}
