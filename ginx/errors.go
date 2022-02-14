package ginx

import (
	"github.com/qxsugar/pkg/zapx"
)

const TAG = "ExceptionX"

func newFailedRespBody(err error, code int, msg string) FailedRespBody {
	logger := zapx.GetLogger()
	logger.Warnw(TAG, "err", err, "code", code, "msg", msg)
	if msg == "" {
		msg = messages[code]
	}
	desc := ""
	if err != nil {
		desc = err.Error()
	}
	return FailedRespBody{
		Succeeded: false,
		RespData:  nil,
		Code:      code,
		Msg:       msg,
		Desc:      desc,
	}
}

func NewInvalidArgumentError(err error, msg string) FailedRespBody {
	return newFailedRespBody(err, InvalidArgument, msg)
}

func NewFailedPreconditionError(err error, msg string) FailedRespBody {
	return newFailedRespBody(err, FailedPrecondition, msg)
}

func NewOutOfRangeError(err error, msg string) FailedRespBody {
	return newFailedRespBody(err, OutOfRange, msg)
}

func NewUnauthenticatedError(err error, msg string) FailedRespBody {
	return newFailedRespBody(err, Unauthenticated, msg)
}

func NewPermissionDeniedError(err error, msg string) FailedRespBody {
	return newFailedRespBody(err, PermissionDenied, msg)
}

func NewNotFoundError(err error, msg string) FailedRespBody {
	return newFailedRespBody(err, NotFound, msg)
}

func NewAbortedError(err error, msg string) FailedRespBody {
	return newFailedRespBody(err, Aborted, msg)
}

func NewAlreadyExistsError(err error, msg string) FailedRespBody {
	return newFailedRespBody(err, AlreadyExists, msg)
}

func NewResourceExhaustedError(err error, msg string) FailedRespBody {
	return newFailedRespBody(err, ResourceExhausted, msg)
}

func NewCancelledError(err error, msg string) FailedRespBody {
	return newFailedRespBody(err, Cancelled, msg)
}

func NewDataLossError(err error, msg string) FailedRespBody {
	return newFailedRespBody(err, DataLoss, msg)
}

func NewUnknownError(err error, msg string) FailedRespBody {
	return newFailedRespBody(err, Unknown, msg)
}

func NewInternalError(err error, msg string) FailedRespBody {
	return newFailedRespBody(err, Internal, msg)
}

func NewNotImplementedError(err error, msg string) FailedRespBody {
	return newFailedRespBody(err, NotImplemented, msg)
}

func NewUnavailableError(err error, msg string) FailedRespBody {
	return newFailedRespBody(err, Unavailable, msg)
}

func NewDeadlineExceededError(err error, msg string) FailedRespBody {
	return newFailedRespBody(err, DeadlineExceeded, msg)
}
