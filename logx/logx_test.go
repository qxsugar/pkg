package logx

import "testing"

func TestGetLogger(t *testing.T) {
	logger := GetLogger()
	logger.Info("info test")
	logger.Debug("debug test")
}
