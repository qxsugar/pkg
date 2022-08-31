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
		return e.Error()
	}
	return "unknown error"
}

func (e *Exception) WithErr(err error) *Exception {
	e.err = err
	return e
}

var _ ApiException = (*Exception)(nil)

// NewError 根据状态码、错误码、错误描述创建一个Error
func NewError(httpCode, businessCode int, msg string) *Exception {
	return &Exception{
		httpCode:     httpCode,
		businessCode: businessCode,
		msg:          msg,
	}
}

// NewErrorWithStatusOk 状态码默认200，根据错误码、错误描述创建一个Error
func NewErrorWithStatusOk(businessCode int, msg string) *Exception {
	return &Exception{
		httpCode:     http.StatusOK,
		businessCode: businessCode,
		msg:          msg,
	}
}

// NewErrorWithStatusOkAutoMsg 状态码默认200，根据错误码创建一个Error（错误描述从 错误码表 中获取）
func NewErrorWithStatusOkAutoMsg(businessCode int) error {
	return &Exception{
		httpCode:     http.StatusOK,
		businessCode: businessCode,
		msg:          "",
	}
}

// NewErrorAutoMsg 根据状态码、错误码创建一个Error
func NewErrorAutoMsg(httpCode, businessCode int) error {
	return &Exception{
		httpCode:     httpCode,
		businessCode: businessCode,
		msg:          "",
	}
}

func newError(httpCode, businessCode int, msg string) *Exception {
	if msg == "" {
		msg = messages[businessCode]
	}

	return &Exception{
		httpCode:     httpCode,
		businessCode: businessCode,
		msg:          msg,
	}
}

// NewInvalidArgumentError 参数错误
func NewInvalidArgumentError(msg string, err error) error {
	return newError(InvalidArgument, InvalidArgument, msg).WithErr(err)
}

// NewFailedPreconditionError 请求不能在当前系统状态下执行，例如删除非空目录
func NewFailedPreconditionError(msg string, err error) error {
	return newError(FailedPrecondition, FailedPrecondition, msg).WithErr(err)
}

// NewOutOfRangeError 客户端指定了无效的范围。
func NewOutOfRangeError(msg string, err error) error {
	return newError(OutOfRange, OutOfRange, msg).WithErr(err)
}

// NewUnauthenticatedError 由于遗失，无效或过期的OAuth令牌而导致请求未通过身份验证。
func NewUnauthenticatedError(msg string, err error) error {
	return newError(Unauthenticated, Unauthenticated, msg).WithErr(err)
}

// NewPermissionDeniedError 客户端没有足够的权限。这可能是因为OAuth令牌没有正确的范围，客户端没有权限，或者客户端项目尚未启用API。
func NewPermissionDeniedError(msg string, err error) error {
	return newError(PermissionDenied, PermissionDenied, msg).WithErr(err)
}

// NewNotFoundError 找不到指定的资源，或者该请求被未公开的原因（例如白名单）拒绝。
func NewNotFoundError(msg string, err error) error {
	return newError(NotFound, NotFound, msg).WithErr(err)
}

// NewAbortedError 并发冲突，例如读-修改-写冲突。
func NewAbortedError(msg string, err error) error {
	return newError(Aborted, Aborted, msg).WithErr(err)
}

// NewAlreadyExistsError 客户端尝试创建的资源已存在。
func NewAlreadyExistsError(msg string, err error) error {
	return newError(AlreadyExists, AlreadyExists, msg).WithErr(err)
}

// NewResourceExhaustedError 资源配额达到速率限制。 客户端应该查找google.rpc.QuotaFailure错误详细信息以获取更多信息。
func NewResourceExhaustedError(msg string, err error) error {
	return newError(ResourceExhausted, ResourceExhausted, msg).WithErr(err)
}

// NewCancelledError 客户端取消请求
func NewCancelledError(msg string, err error) error {
	return newError(Cancelled, Cancelled, msg).WithErr(err)
}

// NewDataLossError 不可恢复的数据丢失或数据损坏。 客户端应该向用户报告错误。
func NewDataLossError(msg string, err error) error {
	return newError(DataLoss, DataLoss, msg).WithErr(err)
}

// NewUnknownError 未知的服务器错误。 通常是服务器错误。
func NewUnknownError(msg string, err error) error {
	return newError(Unknown, Unknown, msg).WithErr(err)
}

// NewInternalError 内部服务错误。 通常是服务器错误。
func NewInternalError(msg string, err error) error {
	return newError(Internal, Internal, msg).WithErr(err)
}

// NewNotImplementedError 服务器未实现该API方法。
func NewNotImplementedError(msg string, err error) error {
	return newError(NotImplemented, NotImplemented, msg).WithErr(err)
}

// NewUnavailableError 暂停服务。通常是服务器已经关闭。
func NewUnavailableError(msg string, err error) error {
	return newError(Unavailable, Unavailable, msg).WithErr(err)
}

// NewDeadlineExceededError 已超过请求期限。如果重复发生，请考虑降低请求的复杂性。
func NewDeadlineExceededError(msg string, err error) error {
	return newError(DeadlineExceeded, DeadlineExceeded, msg).WithErr(err)
}
