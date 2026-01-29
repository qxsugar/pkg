// Package kit provides utilities and middleware for building web applications with the Gin framework.
// It includes standardized error handling, business error management, logging utilities,
// database helpers, and synchronization utilities.
package kit

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

const (
	InternalErrorCode = -1
)

// HandlerFunc defines a custom handler function that returns a response data and an error.
// This allows for standardized error handling through the TranslateFunc middleware.
type HandlerFunc func(ctx *gin.Context) (any, error)

// RouterGroup wraps gin.RouterGroup and provides methods that accept HandlerFunc
// instead of gin.HandlerFunc, enabling automatic error handling.
type RouterGroup struct {
	gin *gin.RouterGroup
}

// NewRouterGroup creates a new RouterGroup wrapper around the given gin.RouterGroup.
func NewRouterGroup(group *gin.RouterGroup) *RouterGroup {
	return &RouterGroup{
		gin: group,
	}
}

// GET registers a GET route with the given path and handler.
func (r *RouterGroup) GET(relativePath string, handler HandlerFunc) *RouterGroup {
	r.gin.GET(relativePath, TranslateFunc(handler))
	return r
}

// POST registers a POST route with the given path and handler.
func (r *RouterGroup) POST(relativePath string, handler HandlerFunc) *RouterGroup {
	r.gin.POST(relativePath, TranslateFunc(handler))
	return r
}

// DELETE registers a DELETE route with the given path and handler.
func (r *RouterGroup) DELETE(relativePath string, handler HandlerFunc) *RouterGroup {
	r.gin.DELETE(relativePath, TranslateFunc(handler))
	return r
}

// PATCH registers a PATCH route with the given path and handler.
func (r *RouterGroup) PATCH(relativePath string, handler HandlerFunc) *RouterGroup {
	r.gin.PATCH(relativePath, TranslateFunc(handler))
	return r
}

// PUT registers a PUT route with the given path and handler.
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
		if ctx.IsAborted() {
			// 如果内部中断了。就不执行转换
			return
		}

		if err != nil {
			respBody := RespBody{
				Succeeded: false,
			}

			var httpStatus int
			switch ex := err.(type) {
			case BusinessError:
				respBody.Code = ex.Code()
				respBody.Info = ex.Info()
				if gin.IsDebugging() {
					respBody.Desc = ex.Desc()
				}
				// Get appropriate HTTP status code, default to 500 if not found
				if status, exists := HTTPStatusCodes[ex.Code()]; exists {
					httpStatus = status
				} else {
					httpStatus = http.StatusInternalServerError
				}
			default:
				respBody.Code = InternalErrorCode
				respBody.Info = Messages[ErrInternal]
				httpStatus = http.StatusInternalServerError
				if gin.IsDebugging() && err != nil {
					respBody.Desc = err.Error()
				}
			}

			logger.Debugf("failed to handler http, code: %d, info: %s, desc: %s, http_status: %d",
				respBody.Code, respBody.Info, respBody.Desc, httpStatus)
			ctx.JSON(httpStatus, respBody)
			return
		}

		ctx.JSON(http.StatusOK, RespBody{Succeeded: true, RespData: resp})
	}
}
