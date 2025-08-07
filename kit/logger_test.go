package kit

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest/observer"
)

func TestNewProductionConfig(t *testing.T) {
	config := NewProductionConfig()

	t.Run("should have production defaults", func(t *testing.T) {
		assert.Equal(t, zapcore.InfoLevel, config.Level.Level())
		assert.Equal(t, "json", config.Encoding)
		assert.Equal(t, []string{"stderr"}, config.ErrorOutputPaths)
		assert.Equal(t, []string{"stderr"}, config.OutputPaths)
	})

	t.Run("should have custom time encoder", func(t *testing.T) {
		// Build logger with memory output to test time encoding
		core, recorded := observer.New(zapcore.InfoLevel)
		logger := zap.New(core)

		testTime := time.Unix(1640995200, 0) // 2022-01-01 00:00:00 UTC

		// Test that our time encoder works
		encoder := zapcore.NewJSONEncoder(config.EncoderConfig)
		buffer, err := encoder.EncodeEntry(zapcore.Entry{
			Time:  testTime,
			Level: zapcore.InfoLevel,
		}, nil)
		assert.NoError(t, err)

		var logEntry map[string]interface{}
		err = json.Unmarshal(buffer.Bytes(), &logEntry)
		assert.NoError(t, err)

		// Should contain unix timestamp
		assert.Contains(t, logEntry, "ts")
		assert.Equal(t, float64(1640995200), logEntry["ts"])

		logger.Info("test") // Prevent unused variable warning
		assert.Equal(t, 1, recorded.Len())
	})
}

func TestNewDevelopmentConfig(t *testing.T) {
	config := NewDevelopmentConfig()

	t.Run("should have development defaults", func(t *testing.T) {
		assert.Equal(t, zapcore.DebugLevel, config.Level.Level())
		assert.Equal(t, "console", config.Encoding)
		assert.True(t, config.Development)
	})

	t.Run("should have custom time encoder", func(t *testing.T) {
		testTime := time.Date(2022, 1, 1, 12, 0, 0, 0, time.UTC)

		encoder := zapcore.NewConsoleEncoder(config.EncoderConfig)
		buffer, err := encoder.EncodeEntry(zapcore.Entry{
			Time:    testTime,
			Level:   zapcore.InfoLevel,
			Message: "test message",
		}, nil)
		assert.NoError(t, err)

		logOutput := buffer.String()

		// Should contain RFC3339 formatted time
		assert.Contains(t, logOutput, "2022-01-01T12:00:00Z")
		assert.Contains(t, logOutput, "test message")
	})
}

func TestMustProduction(t *testing.T) {
	t.Run("should create production logger without panic", func(t *testing.T) {
		assert.NotPanics(t, func() {
			logger := MustProduction()
			assert.NotNil(t, logger)

			// Test that logger works
			logger.Info("test production logger")

			// Clean up
			_ = logger.Sync()
		})
	})

	t.Run("production logger should have correct configuration", func(t *testing.T) {
		logger := MustProduction()

		// Test that it's indeed a production logger by checking core type
		core := logger.Core()
		assert.NotNil(t, core)

		// Production logger should not log debug Messages
		assert.False(t, logger.Core().Enabled(zapcore.DebugLevel))
		assert.True(t, logger.Core().Enabled(zapcore.InfoLevel))

		_ = logger.Sync()
	})
}

func TestMustDevelopment(t *testing.T) {
	t.Run("should create development logger without panic", func(t *testing.T) {
		assert.NotPanics(t, func() {
			logger := MustDevelopment()
			assert.NotNil(t, logger)

			// Test that logger works
			logger.Debug("test development logger")

			// Clean up
			_ = logger.Sync()
		})
	})

	t.Run("development logger should have correct configuration", func(t *testing.T) {
		logger := MustDevelopment()

		// Development logger should log debug Messages
		assert.True(t, logger.Core().Enabled(zapcore.DebugLevel))
		assert.True(t, logger.Core().Enabled(zapcore.InfoLevel))

		_ = logger.Sync()
	})
}

func TestTimeEncoders(t *testing.T) {
	t.Run("production time encoder should output unix timestamp", func(t *testing.T) {
		config := NewProductionConfig()
		testTime := time.Unix(1640995200, 123456789) // With nanoseconds

		encoder := zapcore.NewJSONEncoder(config.EncoderConfig)
		buffer, err := encoder.EncodeEntry(zapcore.Entry{
			Time:  testTime,
			Level: zapcore.InfoLevel,
		}, nil)
		assert.NoError(t, err)

		logOutput := buffer.String()

		// Should contain unix timestamp (seconds only, not nanoseconds)
		assert.Contains(t, logOutput, "1640995200")
		assert.NotContains(t, logOutput, "123456789") // Nanoseconds should be ignored
	})

	t.Run("development time encoder should output RFC3339", func(t *testing.T) {
		config := NewDevelopmentConfig()
		testTime := time.Date(2022, 1, 1, 12, 30, 45, 0, time.UTC)

		encoder := zapcore.NewConsoleEncoder(config.EncoderConfig)
		buffer, err := encoder.EncodeEntry(zapcore.Entry{
			Time:    testTime,
			Level:   zapcore.InfoLevel,
			Message: "test",
		}, nil)
		assert.NoError(t, err)

		logOutput := buffer.String()
		assert.Contains(t, logOutput, "2022-01-01T12:30:45Z")
	})
}

func TestLoggerIntegration(t *testing.T) {
	t.Run("production logger integration test", func(t *testing.T) {
		logger := MustProduction()
		defer logger.Sync()

		// Test various log levels
		logger.Info("info message", zap.String("key", "value"))
		logger.Warn("warning message", zap.Int("number", 42))
		logger.Error("error message", zap.Error(assert.AnError))

		// Debug should not be logged in production
		logger.Debug("debug message - should not appear")
	})

	t.Run("development logger integration test", func(t *testing.T) {
		logger := MustDevelopment()
		defer logger.Sync()

		// Test various log levels
		logger.Debug("debug message", zap.String("key", "value"))
		logger.Info("info message", zap.Int("number", 42))
		logger.Warn("warning message")
		logger.Error("error message", zap.Error(assert.AnError))
	})
}

func TestLoggerWithFields(t *testing.T) {
	t.Run("production logger with fields", func(t *testing.T) {
		logger := MustProduction().With(
			zap.String("service", "test-service"),
			zap.String("version", "1.0.0"),
		)

		logger.Info("service started", zap.Int("port", 8080))
		logger.Error("service error", zap.Error(assert.AnError))

		_ = logger.Sync()
	})

	t.Run("development logger with fields", func(t *testing.T) {
		logger := MustDevelopment().With(
			zap.String("component", "test-component"),
			zap.String("trace_id", "abc123"),
		)

		logger.Debug("component initialized")
		logger.Info("processing request", zap.String("path", "/api/test"))

		_ = logger.Sync()
	})
}

func TestConfigComparison(t *testing.T) {
	t.Run("production and development configs should differ", func(t *testing.T) {
		prodConfig := NewProductionConfig()
		devConfig := NewDevelopmentConfig()

		// Level should be different
		assert.NotEqual(t, prodConfig.Level.Level(), devConfig.Level.Level())

		// Encoding should be different
		assert.NotEqual(t, prodConfig.Encoding, devConfig.Encoding)

		// Development flag should be different
		assert.NotEqual(t, prodConfig.Development, devConfig.Development)
	})
}
