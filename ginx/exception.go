package ginx

import (
	"net/http"
)

const (
	defaultHttpCode = http.StatusOK
	maxHttpCode     = 600
	minHttpCode     = 100
)

type Exception struct {
	httpCode int // http code

	code int    // business code
	info string // business information
	desc string // business description
}

var _ BusinessError = &Exception{}
var _ HttpError = &Exception{}
var _ error = &Exception{}

func (e *Exception) HttpCode() int {
	if e.httpCode < minHttpCode || e.httpCode > maxHttpCode {
		return defaultHttpCode
	}
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
	if e.info != "" {
		return e.info
	}
	return "unknown error"
}

// WithErr set desc = err.Error() when error is not nil
func (e *Exception) WithErr(err error) *Exception {
	if err != nil {
		e.desc = err.Error()
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
	return newException(http.StatusOK, OK, "")
}

func newException(httpCode, code int, message string) *Exception {
	return NewException().
		WithHttpCode(httpCode).
		WithCode(code).
		WithInfo(message)
}

func NewExceptionWithStatusOk(code int, info string) *Exception {
	return newException(http.StatusOK, code, info)
}

func NewInvalidArgumentError() *Exception {
	return newException(ErrInvalidArgument, ErrInvalidArgument, messages[ErrInvalidArgument])
}

func NewFailedPreconditionError() *Exception {
	return newException(ErrFailedPrecondition, ErrFailedPrecondition, messages[ErrFailedPrecondition])
}

func NewOutOfRangeError() *Exception {
	return newException(ErrOutOfRange, ErrOutOfRange, messages[ErrOutOfRange])
}

func NewUnauthenticatedError() *Exception {
	return newException(ErrUnauthenticated, ErrUnauthenticated, messages[ErrUnauthenticated])
}

func NewPermissionDeniedError() *Exception {
	return newException(ErrPermissionDenied, ErrPermissionDenied, messages[ErrPermissionDenied])
}

func NewNotFoundError() *Exception {
	return newException(ErrNotFound, ErrNotFound, messages[ErrNotFound])
}

func NewAbortedError() *Exception {
	return newException(ErrAborted, ErrAborted, messages[ErrAborted])
}

func NewAlreadyExistsError() *Exception {
	return newException(ErrAlreadyExists, ErrAlreadyExists, messages[ErrAlreadyExists])
}

func NewResourceExhaustedError() *Exception {
	return newException(ErrResourceExhausted, ErrResourceExhausted, messages[ErrResourceExhausted])
}

func NewCancelledError() *Exception {
	return newException(ErrCancelled, ErrCancelled, messages[ErrCancelled])
}

func NewDataLossError() *Exception {
	return newException(ErrDataLoss, ErrDataLoss, messages[ErrDataLoss])
}

func NewUnknownError() *Exception {
	return newException(ErrUnknown, ErrUnknown, messages[ErrUnknown])
}

func NewInternalError() *Exception {
	return newException(ErrInternal, ErrInternal, messages[ErrInternal])
}

func NewNotImplementedError() *Exception {
	return newException(ErrNotImplemented, ErrNotImplemented, messages[ErrNotImplemented])
}

func NewUnavailableError() *Exception {
	return newException(ErrUnavailable, ErrUnavailable, messages[ErrUnavailable])
}

func NewDeadlineExceededError() *Exception {
	return newException(ErrDeadlineExceeded, ErrDeadlineExceeded, messages[ErrDeadlineExceeded])
}
