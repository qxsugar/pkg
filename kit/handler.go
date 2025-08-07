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

// TranslateFunc 将 HandlerFunc 转换为 gin.HandlerFunc
// 并处理错误，返回统一的响应格式
// 如果发生错误，返回 RespBody 中的 Succeeded 为 false，并包含错误信息
// 如果成功，返回 RespBody 中的 Succeeded 为 true，并包含响应数据
// 如果发生业务错误，返回 RespBody 中的 Code 和 Info 字段
// 如果发生非业务错误，返回 RespBody 中的 Code 为 InternalErrorCode，并包含错误信息
// 如果在开发模式下，返回 RespBody 中的 Desc 字段包含详细错误描述
func TranslateFunc(fun HandlerFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		logger := zap.S().Named("TranslateFunc")

		resp, err := fun(ctx)
		if err != nil {
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
				respBody.Info = Messages[ErrInternal]
				if gin.IsDebugging() && err != nil {
					respBody.Desc = err.Error()
				}
			}

			logger.Debugf("failed to handler http, code: %d, info: %s, desc: %s", respBody.Code, respBody.Info, respBody.Desc)
			ctx.JSON(http.StatusOK, respBody)
			return
		}

		ctx.JSON(http.StatusOK, RespBody{Succeeded: true, RespData: resp})
		return
	}
}
