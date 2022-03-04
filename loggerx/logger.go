package loggerx

import (
	"go.uber.org/zap"
)

var logger *zap.SugaredLogger

func init() {
	_logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	zap.ReplaceGlobals(_logger)
	logger = _logger.Sugar()
}

func GetLogger() *zap.SugaredLogger {
	return logger
}
