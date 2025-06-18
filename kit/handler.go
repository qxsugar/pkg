package kit

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

const (
	InternalErrorCode = -1
)

type HandlerFunc func(ctx *gin.Context) (any, error)

type RouterGroup struct {
	gin *gin.RouterGroup
}

func NewRouterGroup(group *gin.RouterGroup) *RouterGroup {
	return &RouterGroup{
		gin: group,
	}
}

func (r *RouterGroup) GET(relativePath string, handler HandlerFunc) *RouterGroup {
	r.gin.GET(relativePath, TranslateFunc(handler))
	return r
}

func (r *RouterGroup) POST(relativePath string, handler HandlerFunc) *RouterGroup {
	r.gin.POST(relativePath, TranslateFunc(handler))
	return r
}

func (r *RouterGroup) DELETE(relativePath string, handler HandlerFunc) *RouterGroup {
	r.gin.DELETE(relativePath, TranslateFunc(handler))
	return r
}

func (r *RouterGroup) PATCH(relativePath string, handler HandlerFunc) *RouterGroup {
	r.gin.PATCH(relativePath, TranslateFunc(handler))
	return r
}

func (r *RouterGroup) PUT(relativePath string, handler HandlerFunc) *RouterGroup {
	r.gin.PUT(relativePath, TranslateFunc(handler))
	return r
}

func TranslateFunc(fun HandlerFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var (
			logger = zap.S().Named("TranslateFunc")
		)

		resp, err := fun(ctx)
		if err == nil {
			ctx.JSON(http.StatusOK, RespBody{Succeeded: true, RespData: resp})
			return
		}

		respBody := RespBody{
			Succeeded: false,
		}

		switch ex := err.(type) {
		case BusinessError:
			respBody.Code = ex.Code()
			respBody.Info = ex.Info()
			if gin.IsDebugging() {
				respBody.Desc = ex.Desc()
			}
		default:
			respBody.Code = InternalErrorCode
			respBody.Info = messages[ErrInternal]
			if gin.IsDebugging() && err != nil {
				respBody.Desc = err.Error()
			}
		}

		logger.Infof("failed to handler http, code: %d, info: %s, desc: %s", respBody.Code, respBody.Info, respBody.Desc)

		ctx.JSON(http.StatusOK, respBody)
		return
	}
}
