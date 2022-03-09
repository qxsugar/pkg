package apix

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Pong(ctx *gin.Context) {
	ctx.String(http.StatusOK, "pong")
}
