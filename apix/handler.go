package apix

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

var Debug = true

func IsDebug() bool {
	return Debug
}

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

type RespBody struct {
	Succeeded bool        `json:"succeeded"`
	RespData  interface{} `json:"resp_data"`
	HttpCode  int         `json:"-"`
	Code      int         `json:"code,omitempty"`
	Msg       string      `json:"msg,omitempty"`
	Desc      string      `json:"desc,omitempty"`
}

type ApiException interface {
	GetHttpCode() int
	GetBusinessCode() int
	GetMsg() string
	GetDesc() string
	Error() string
}

func Wrapper(fun HandlerFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		logger := zap.S()
		resp, err := fun(ctx)
		if err == nil {
			ctx.JSON(http.StatusOK, RespBody{
				Succeeded: true,
				RespData:  resp,
			})
			return
		}

		apiException, ok := err.(ApiException)

		respBody := RespBody{}
		httpCode := http.StatusOK
		if !ok {
			logger.Warnf("failed to handler http, unkonwn error: %v", err)
			httpCode = http.StatusInternalServerError
			respBody.Succeeded = false
			respBody.Code = -1
			respBody.Msg = "Unknown Error"
			if IsDebug() && err != nil {
				respBody.Desc = err.Error()
			}
		} else {
			if apiException.GetHttpCode() < 100 || apiException.GetHttpCode() > 505 {
				httpCode = http.StatusInternalServerError
			} else {
				httpCode = apiException.GetHttpCode()
			}
			respBody.Succeeded = false
			respBody.Code = apiException.GetBusinessCode()
			respBody.Msg = apiException.GetMsg()
			if IsDebug() {
				respBody.Desc = err.Error()
			}
		}

		ctx.JSON(httpCode, respBody)
		return
	}
}
