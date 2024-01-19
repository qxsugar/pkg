package ginx

import (
	"github.com/gin-gonic/gin"
)

func Pong(ctx *gin.Context) (any, error) {
	return "pong", nil
}
