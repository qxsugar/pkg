package apix

import (
	"go.uber.org/zap"
	"net/http"
)

type Exception struct {
	HttpCode int
	Code     int
	Msg      string
	Desc     string
}

func (e *Exception) GetHttpCode() int {
	return e.HttpCode
}

func (e *Exception) GetCode() int {
	return e.Code
}

func (e *Exception) GetMsg() string {
	return e.Msg
}

func (e *Exception) GetDesc() string {
	return e.Desc
}

func (e Exception) Error() string {
	return e.Msg
}

var _ ApiException = (*Exception)(nil)

func newException(httpCode, code int, msg string, err error) error {
	logger := zap.S()
	logger.Infof("%s: httpCode:%d, code: %d, msg: %s, err: %v", "[API_ERROR]", httpCode, code, msg, err)
	if msg == "" {
		msg = messages[code]
	}

	desc := ""
	if err != nil {
		desc = err.Error()
	}

	return &Exception{
		HttpCode: httpCode,
		Code:     code,
		Msg:      msg,
		Desc:     desc,
	}
}

// NewCustomError 自定义消息
func NewCustomError(code int, msg string) error {
	return newException(http.StatusOK, code, msg, nil)
}

func NewInvalidArgumentError(err error, msg string) error {
	return newException(InvalidArgument, InvalidArgument, msg, err)
}

func NewFailedPreconditionError(err error, msg string) error {
	return newException(FailedPrecondition, FailedPrecondition, msg, err)
}

func NewOutOfRangeError(err error, msg string) error {
	return newException(OutOfRange, OutOfRange, msg, err)
}

func NewUnauthenticatedError(err error, msg string) error {
	return newException(Unauthenticated, Unauthenticated, msg, err)
}

func NewPermissionDeniedError(err error, msg string) error {
	return newException(PermissionDenied, PermissionDenied, msg, err)
}

func NewNotFoundError(err error, msg string) error {
	return newException(NotFound, NotFound, msg, err)
}

func NewAbortedError(err error, msg string) error {
	return newException(Aborted, Aborted, msg, err)
}

func NewAlreadyExistsError(err error, msg string) error {
	return newException(AlreadyExists, AlreadyExists, msg, err)
}

func NewResourceExhaustedError(err error, msg string) error {
	return newException(ResourceExhausted, ResourceExhausted, msg, err)
}

func NewCancelledError(err error, msg string) error {
	return newException(Cancelled, Cancelled, msg, err)
}

func NewDataLossError(err error, msg string) error {
	return newException(DataLoss, DataLoss, msg, err)
}

func NewUnknownError(err error, msg string) error {
	return newException(Unknown, Unknown, msg, err)
}

func NewInternalError(err error, msg string) error {
	return newException(Internal, Internal, msg, err)
}

func NewNotImplementedError(err error, msg string) error {
	return newException(NotImplemented, NotImplemented, msg, err)
}

func NewUnavailableError(err error, msg string) error {
	return newException(Unavailable, Unavailable, msg, err)
}

func NewDeadlineExceededError(err error, msg string) error {
	return newException(DeadlineExceeded, DeadlineExceeded, msg, err)
}
