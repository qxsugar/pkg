package apix

import (
	"github.com/gin-gonic/gin"
)

func Pong(ctx *gin.Context) (interface{}, error) {
	return "pong", nil
}
