package logx

import (
	"go.uber.org/zap"
)

func NewLogger() *zap.Logger {
	return zap.L()
}

func NewSugaredLogger() *zap.SugaredLogger {
	return zap.S()
}

func NewDevelopmentLogger() *zap.Logger {
	return zap.Must(zap.NewDevelopment())
}

func NewProductionLogger() *zap.Logger {
	return zap.Must(zap.NewProduction())
}

func init() {
	zap.ReplaceGlobals(zap.Must(zap.NewDevelopment()))
}
