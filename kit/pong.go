package kit

import (
	"github.com/gin-gonic/gin"
)

// Pong is a simple health check handler that returns "pong".
// It can be used to verify that the service is running.
func Pong(ctx *gin.Context) (any, error) {
	return "pong", nil
}
