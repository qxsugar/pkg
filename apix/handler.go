package apix

import (
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
	r.gin.GET(relativePath, W(handler))
	return r
}

func (r *RouterGroup) POST(relativePath string, handler HandlerFunc) *RouterGroup {
	r.gin.POST(relativePath, W(handler))
	return r
}

func (r *RouterGroup) DELETE(relativePath string, handler HandlerFunc) *RouterGroup {
	r.gin.DELETE(relativePath, W(handler))
	return r
}

func (r *RouterGroup) PATCH(relativePath string, handler HandlerFunc) *RouterGroup {
	r.gin.PATCH(relativePath, W(handler))
	return r
}

func (r *RouterGroup) PUT(relativePath string, handler HandlerFunc) *RouterGroup {
	r.gin.PUT(relativePath, W(handler))
	return r
}

func W(fun HandlerFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		resp, err := fun(ctx)
		if err != nil {
			if e, ok := err.(RespBody); ok {
				ctx.JSON(e.Code, e)
			} else {
				ctx.JSON(Internal, newError(err, Internal, "服务内部出错"))
			}
			return
		}

		ctx.JSON(http.StatusOK, RespBody{
			Succeeded: true,
			RespData:  resp,
		})
		return
	}
}
