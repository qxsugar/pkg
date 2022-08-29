package apix

import (
	"fmt"
	"github.com/gin-gonic/gin"
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
	GetCode() int
	GetMsg() string
	GetDesc() string
}

func Wrapper(fun HandlerFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		resp, err := fun(ctx)
		if err == nil {
			ctx.JSON(http.StatusOK, RespBody{
				Succeeded: true,
				RespData:  resp,
			})
			return
		}

		apiException, ok := err.(ApiException)
		if !ok {
			ctx.JSON(http.StatusInternalServerError, RespBody{
				Succeeded: false,
				RespData:  nil,
				Code:      http.StatusInternalServerError,
				Msg:       "Unknown Error",
				Desc:      err.Error(),
			})
			return
		}

		if apiException.GetHttpCode() < 100 || apiException.GetHttpCode() > 505 {
			ctx.JSON(http.StatusInternalServerError, RespBody{
				Succeeded: false,
				RespData:  nil,
				Code:      apiException.GetCode(),
				Msg:       fmt.Sprintf("invalid http code: %d", apiException.GetHttpCode()),
				Desc:      apiException.GetDesc(),
			})
			return
		}

		ctx.JSON(apiException.GetHttpCode(), RespBody{
			Succeeded: false,
			RespData:  nil,
			Code:      apiException.GetCode(),
			Msg:       apiException.GetMsg(),
			Desc:      apiException.GetDesc(),
		})
		return
	}
}
