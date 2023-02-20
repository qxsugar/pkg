package apix

import (
	"net/http"
)

type ApiException interface {
	HttpCode() int
	Code() int
	Info() string
	Desc() string
}

type Exception struct {
	httpCode int    // http code
	code     int    // business code
	info     string // information about
	desc     string // description of exception
}

var _ ApiException = &Exception{}

func (e *Exception) HttpCode() int {
	return e.httpCode
}

func (e *Exception) Code() int {
	return e.code
}

func (e *Exception) Info() string {
	return e.info
}

func (e *Exception) Desc() string {
	return e.desc
}

func (e *Exception) Error() string {
	if e.desc != "" {
		return e.desc
	}
	return "unknown error"
}

func (e *Exception) WithErr(err error) *Exception {
	if err != nil {
		e.desc = e.Error()
	}
	return e
}

func (e *Exception) WithHttpCode(code int) *Exception {
	e.httpCode = code
	return e
}

func (e *Exception) WithCode(code int) *Exception {
	e.code = code
	return e
}

func (e *Exception) WithInfo(info string) *Exception {
	e.info = info
	return e
}

// NewException httpCode: 200, code: 0, info: "", desc: ""
func NewException() *Exception {
	return &Exception{
		httpCode: http.StatusOK,
	}
}

func NewExceptionWithStatusOk(code int, info string) *Exception {
	return NewException().WithHttpCode(http.StatusOK).WithCode(code).WithInfo(info)
}

func NewExceptionWithStatusOkAutoMsg(code int) *Exception {
	return NewException().WithHttpCode(http.StatusOK).WithCode(code).WithInfo("")
}

func NewExceptionAutoMsg(httpCode, code int) *Exception {
	return NewException().WithHttpCode(httpCode).WithCode(code).WithInfo("")
}

func NewInvalidArgumentError() *Exception {
	return NewException().WithHttpCode(InvalidArgument).WithCode(InvalidArgument).WithInfo(messages[InvalidArgument])
}

func NewFailedPreconditionError() *Exception {
	return NewException().WithHttpCode(FailedPrecondition).WithCode(FailedPrecondition).WithInfo(messages[FailedPrecondition])
}

func NewOutOfRangeError() *Exception {
	return NewException().WithHttpCode(OutOfRange).WithCode(OutOfRange).WithInfo(messages[OutOfRange])
}

func NewUnauthenticatedError() *Exception {
	return NewException().WithHttpCode(Unauthenticated).WithCode(Unauthenticated).WithInfo(messages[Unauthenticated])
}

func NewPermissionDeniedError() *Exception {
	return NewException().WithHttpCode(PermissionDenied).WithCode(PermissionDenied).WithInfo(messages[PermissionDenied])
}

func NewNotFoundError() *Exception {
	return NewException().WithHttpCode(NotFound).WithCode(NotFound).WithInfo(messages[NotFound])
}

func NewAbortedError() *Exception {
	return NewException().WithHttpCode(Aborted).WithCode(Aborted).WithInfo(messages[Aborted])
}

func NewAlreadyExistsError() *Exception {
	return NewException().WithHttpCode(AlreadyExists).WithCode(AlreadyExists).WithInfo(messages[AlreadyExists])
}

func NewResourceExhaustedError() *Exception {
	return NewException().WithHttpCode(ResourceExhausted).WithCode(ResourceExhausted).WithInfo(messages[ResourceExhausted])
}

func NewCancelledError() *Exception {
	return NewException().WithHttpCode(Cancelled).WithCode(Cancelled).WithInfo(messages[Cancelled])
}

func NewDataLossError() *Exception {
	return NewException().WithHttpCode(DataLoss).WithCode(DataLoss).WithInfo(messages[DataLoss])
}

func NewUnknownError() *Exception {
	return NewException().WithHttpCode(Unknown).WithCode(Unknown).WithInfo(messages[Unknown])
}

func NewInternalError() *Exception {
	return NewException().WithHttpCode(Internal).WithCode(Internal).WithInfo(messages[Internal])
}

func NewNotImplementedError() *Exception {
	return NewException().WithHttpCode(NotImplemented).WithCode(NotImplemented).WithInfo(messages[NotImplemented])
}

func NewUnavailableError() *Exception {
	return NewException().WithHttpCode(Unavailable).WithCode(Unavailable).WithInfo(messages[Unavailable])
}

func NewDeadlineExceededError() *Exception {
	return NewException().WithHttpCode(DeadlineExceeded).WithCode(DeadlineExceeded).WithInfo(messages[DeadlineExceeded])
}
