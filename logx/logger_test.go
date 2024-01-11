package logx

import (
	"testing"
)

func TestNewLogger(t *testing.T) {
	logger := NewLogger().WithOptions()
	logger.Info("test log message")
}

func TestNewSugaredLogger(t *testing.T) {
	logger := NewSugaredLogger()
	logger.Infow("test log message")
}

func TestNewDevelopmentLogger(t *testing.T) {
	logger := NewDevelopmentLogger()
	logger.Info("test log message")
}

func TestNewProductionLogger(t *testing.T) {
	logger := NewProductionLogger()
	logger.Info("test log message")
}
