package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HandlerFunc func(ctx *gin.Context) (interface{}, error)

func W(fun HandlerFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		resp, err := fun(ctx)
		if err != nil {
			if e, ok := err.(Error); ok {
				ctx.JSON(e.Code, RespBody{
					Succeeded: false,
					Code:      e.Code,
					Msg:       e.Msg,
					Desc:      e.ErrMsg,
				})
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
