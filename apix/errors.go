package apix

import (
	"github.com/qxsugar/pkg/loggerx"
)

type Error struct {
	Code   int
	Msg    string
	ErrMsg string
}

func (e Error) Error() string {
	return e.ErrMsg
}

const TAG = "[API_ERROR]"

func newError(err error, code int, msg string) Error {
	logger := loggerx.GetLogger()
	logger.Warnw(TAG, "err", err, "code", code, "msg", msg)
	if msg == "" {
		msg = messages[code]
	}
	desc := ""
	if err != nil {
		desc = err.Error()
	}
	return Error{
		Code:   code,
		Msg:    msg,
		ErrMsg: desc,
	}
}

func NewInvalidArgumentError(err error, msg string) error {
	return newError(err, InvalidArgument, msg)
}

func NewFailedPreconditionError(err error, msg string) error {
	return newError(err, FailedPrecondition, msg)
}

func NewOutOfRangeError(err error, msg string) error {
	return newError(err, OutOfRange, msg)
}

func NewUnauthenticatedError(err error, msg string) error {
	return newError(err, Unauthenticated, msg)
}

func NewPermissionDeniedError(err error, msg string) error {
	return newError(err, PermissionDenied, msg)
}

func NewNotFoundError(err error, msg string) error {
	return newError(err, NotFound, msg)
}

func NewAbortedError(err error, msg string) error {
	return newError(err, Aborted, msg)
}

func NewAlreadyExistsError(err error, msg string) error {
	return newError(err, AlreadyExists, msg)
}

func NewResourceExhaustedError(err error, msg string) error {
	return newError(err, ResourceExhausted, msg)
}

func NewCancelledError(err error, msg string) error {
	return newError(err, Cancelled, msg)
}

func NewDataLossError(err error, msg string) error {
	return newError(err, DataLoss, msg)
}

func NewUnknownError(err error, msg string) error {
	return newError(err, Unknown, msg)
}

func NewInternalError(err error, msg string) error {
	return newError(err, Internal, msg)
}

func NewNotImplementedError(err error, msg string) error {
	return newError(err, NotImplemented, msg)
}

func NewUnavailableError(err error, msg string) error {
	return newError(err, Unavailable, msg)
}

func NewDeadlineExceededError(err error, msg string) error {
	return newError(err, DeadlineExceeded, msg)
}
