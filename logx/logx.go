package logx

import (
	"go.uber.org/zap"
)

func init() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	zap.ReplaceGlobals(logger)
}

func Get() *zap.SugaredLogger {
	return zap.S()
}
