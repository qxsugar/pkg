package apix

import (
	"net/http"
)

type Exception struct {
	httpCode     int
	businessCode int
	msg          string
	err          error
}

func (e *Exception) GetHttpCode() int {
	return e.httpCode
}

func (e *Exception) GetBusinessCode() int {
	return e.businessCode
}

func (e *Exception) GetMsg() string {
	return e.msg
}

func (e Exception) Error() string {
	if e.err != nil {
		return e.err.Error()
	}
	return "unknown error"
}

func (e *Exception) WithErr(err error) *Exception {
	e.err = err
	return e
}

func (e *Exception) WithHttpCode(code int) *Exception {
	e.httpCode = code
	return e
}

func (e *Exception) WithBusinessCode(code int) *Exception {
	e.businessCode = code
	return e
}

func (e *Exception) WithMsg(msg string) *Exception {
	e.msg = msg
	return e
}

var _ ApiException = (*Exception)(nil)

// NewError httpCode: 200, code: 0, msg: "", desc: ""
func NewError() *Exception {
	return &Exception{
		httpCode: http.StatusOK,
	}
}

// NewErrorWithStatusOk httpCode: 200, code: businessCode, msg: msg, desc: ""
func NewErrorWithStatusOk(businessCode int, msg string) *Exception {
	return NewError().WithHttpCode(http.StatusOK).WithBusinessCode(businessCode).WithMsg(msg)
}

// NewErrorWithStatusOkAutoMsg httpCode: 200, code: businessCode, msg: "", desc: ""
func NewErrorWithStatusOkAutoMsg(businessCode int) *Exception {
	return NewError().WithHttpCode(http.StatusOK).WithBusinessCode(businessCode).WithMsg("")
}

// NewErrorAutoMsg httpCode: httpCode, code: 0, msg: "", desc: ""
func NewErrorAutoMsg(httpCode, businessCode int) *Exception {
	return NewError().WithHttpCode(httpCode).WithBusinessCode(businessCode).WithMsg("")
}

// NewInvalidArgumentError 参数错误
func NewInvalidArgumentError() *Exception {
	return NewError().WithHttpCode(InvalidArgument).WithBusinessCode(InvalidArgument).WithMsg(messages[InvalidArgument])
}

// NewFailedPreconditionError 请求不能在当前系统状态下执行，例如删除非空目录
func NewFailedPreconditionError() *Exception {
	return NewError().WithHttpCode(FailedPrecondition).WithBusinessCode(FailedPrecondition).WithMsg(messages[FailedPrecondition])
}

// NewOutOfRangeError 客户端指定了无效的范围。
func NewOutOfRangeError() *Exception {
	return NewError().WithHttpCode(OutOfRange).WithBusinessCode(OutOfRange).WithMsg(messages[OutOfRange])
}

// NewUnauthenticatedError 由于遗失，无效或过期的OAuth令牌而导致请求未通过身份验证。
func NewUnauthenticatedError() *Exception {
	return NewError().WithHttpCode(Unauthenticated).WithBusinessCode(Unauthenticated).WithMsg(messages[Unauthenticated])
}

// NewPermissionDeniedError 客户端没有足够的权限。这可能是因为OAuth令牌没有正确的范围，客户端没有权限，或者客户端项目尚未启用API。
func NewPermissionDeniedError() *Exception {
	return NewError().WithHttpCode(PermissionDenied).WithBusinessCode(PermissionDenied).WithMsg(messages[PermissionDenied])
}

// NewNotFoundError 找不到指定的资源，或者该请求被未公开的原因（例如白名单）拒绝。
func NewNotFoundError() *Exception {
	return NewError().WithHttpCode(NotFound).WithBusinessCode(NotFound).WithMsg(messages[NotFound])
}

// NewAbortedError 并发冲突，例如读-修改-写冲突。
func NewAbortedError() *Exception {
	return NewError().WithHttpCode(Aborted).WithBusinessCode(Aborted).WithMsg(messages[Aborted])
}

// NewAlreadyExistsError 客户端尝试创建的资源已存在。
func NewAlreadyExistsError() *Exception {
	return NewError().WithHttpCode(AlreadyExists).WithBusinessCode(AlreadyExists).WithMsg(messages[AlreadyExists])
}

// NewResourceExhaustedError 资源配额达到速率限制。 客户端应该查找google.rpc.QuotaFailure错误详细信息以获取更多信息。
func NewResourceExhaustedError() *Exception {
	return NewError().WithHttpCode(ResourceExhausted).WithBusinessCode(ResourceExhausted).WithMsg(messages[ResourceExhausted])
}

// NewCancelledError 客户端取消请求
func NewCancelledError() *Exception {
	return NewError().WithHttpCode(Cancelled).WithBusinessCode(Cancelled).WithMsg(messages[Cancelled])
}

// NewDataLossError 不可恢复的数据丢失或数据损坏。 客户端应该向用户报告错误。
func NewDataLossError() *Exception {
	return NewError().WithHttpCode(DataLoss).WithBusinessCode(DataLoss).WithMsg(messages[DataLoss])
}

// NewUnknownError 未知的服务器错误。 通常是服务器错误。
func NewUnknownError() *Exception {
	return NewError().WithHttpCode(Unknown).WithBusinessCode(Unknown).WithMsg(messages[Unknown])
}

// NewInternalError 内部服务错误。 通常是服务器错误。
func NewInternalError() *Exception {
	return NewError().WithHttpCode(Internal).WithBusinessCode(Internal).WithMsg(messages[Internal])
}

// NewNotImplementedError 服务器未实现该API方法。
func NewNotImplementedError() *Exception {
	return NewError().WithHttpCode(NotImplemented).WithBusinessCode(NotImplemented).WithMsg(messages[NotImplemented])
}

// NewUnavailableError 暂停服务。通常是服务器已经关闭。
func NewUnavailableError() *Exception {
	return NewError().WithHttpCode(Unavailable).WithBusinessCode(Unavailable).WithMsg(messages[Unavailable])
}

// NewDeadlineExceededError 已超过请求期限。如果重复发生，请考虑降低请求的复杂性。
func NewDeadlineExceededError() *Exception {
	return NewError().WithHttpCode(DeadlineExceeded).WithBusinessCode(DeadlineExceeded).WithMsg(messages[DeadlineExceeded])
}
