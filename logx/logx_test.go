package logx

import "testing"

func TestGetLogger(t *testing.T) {
	logger := Get()
	logger.Info("info test")
	logger.Debug("debug test")
}
