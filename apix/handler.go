package apix

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HandlerFunc func(ctx *gin.Context) (interface{}, error)

func W(fun HandlerFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		resp, err := fun(ctx)
		if err != nil {
			if e, ok := err.(RespBody); ok {
				ctx.JSON(e.Code, e)
				return
			} else {
				ctx.JSON(Internal, newError(err, Internal, "服务内部出错"))
				return
			}
		}

		ctx.JSON(http.StatusOK, RespBody{
			Succeeded: true,
			RespData:  resp,
		})
		return
	}
}
