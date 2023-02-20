package apix

import (
	"github.com/gin-gonic/gin"
	"github.com/qxsugar/pkg/logx"
	"net/http"
)

type HandlerFunc func(ctx *gin.Context) (interface{}, error)

type RouterGroup struct {
	gin *gin.RouterGroup
}

func NewRouterGroup(group *gin.RouterGroup) *RouterGroup {
	return &RouterGroup{
		gin: group,
	}
}

func (r *RouterGroup) GET(relativePath string, handler HandlerFunc) *RouterGroup {
	r.gin.GET(relativePath, Wrapper(handler))
	return r
}

func (r *RouterGroup) POST(relativePath string, handler HandlerFunc) *RouterGroup {
	r.gin.POST(relativePath, Wrapper(handler))
	return r
}

func (r *RouterGroup) DELETE(relativePath string, handler HandlerFunc) *RouterGroup {
	r.gin.DELETE(relativePath, Wrapper(handler))
	return r
}

func (r *RouterGroup) PATCH(relativePath string, handler HandlerFunc) *RouterGroup {
	r.gin.PATCH(relativePath, Wrapper(handler))
	return r
}

func (r *RouterGroup) PUT(relativePath string, handler HandlerFunc) *RouterGroup {
	r.gin.PUT(relativePath, Wrapper(handler))
	return r
}

func Wrapper(fun HandlerFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		logger := logx.Get()

		resp, err := fun(ctx)
		if err == nil {
			ctx.JSON(http.StatusOK, RespBody{Succeeded: true, RespData: resp})
			return
		}

		httpCode := http.StatusOK
		respBody := RespBody{}

		switch exception := err.(type) {
		case ApiException:
			httpCode = exception.HttpCode()
			respBody.Succeeded = false
			respBody.Code = exception.Code()
			respBody.Info = exception.Info()
			if gin.IsDebugging() {
				respBody.Desc = exception.Desc()
			}
		default:
			httpCode = http.StatusInternalServerError
			respBody.Succeeded = false
			respBody.Code = -1
			respBody.Info = "Unknown Error"
			if gin.IsDebugging() && err != nil {
				respBody.Desc = err.Error()
			}
		}

		// limit http code
		if httpCode < http.StatusContinue || httpCode > http.StatusNetworkAuthenticationRequired {
			httpCode = http.StatusInternalServerError
		}

		logger.Warnf("failed to handler http, httpCode: %d, code: %d, info: %s, desc: %s", httpCode, respBody.Code, respBody.Info, respBody.Desc)

		ctx.JSON(httpCode, respBody)
		return
	}
}
