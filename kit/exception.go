package kit

type Exception struct {
	code int    // business code
	info string // business information, to user
	desc string // business description, to developer
}

var _ BusinessError = &Exception{}
var _ error = &Exception{}

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
	return messages[ErrUnknown]
}

// WithErr set desc = err.Error() when error is not nil
func (e *Exception) WithErr(err error) *Exception {
	if err != nil {
		e.desc = err.Error()
	}
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

func newException(code int, info string) *Exception {
	return (&Exception{}).WithCode(code).WithInfo(info)
}

// NewException httpCode: 200, code: 0, info: "", desc: ""
func NewException() *Exception {
	return newException(OK, "")
}

func NewInvalidArgumentError() *Exception {
	return newException(ErrInvalidArgument, messages[ErrInvalidArgument])
}

func NewFailedPreconditionError() *Exception {
	return newException(ErrFailedPrecondition, messages[ErrFailedPrecondition])
}

func NewOutOfRangeError() *Exception {
	return newException(ErrOutOfRange, messages[ErrOutOfRange])
}

func NewUnauthenticatedError() *Exception {
	return newException(ErrUnauthenticated, messages[ErrUnauthenticated])
}

func NewPermissionDeniedError() *Exception {
	return newException(ErrPermissionDenied, messages[ErrPermissionDenied])
}

func NewNotFoundError() *Exception {
	return newException(ErrNotFound, messages[ErrNotFound])
}

func NewAbortedError() *Exception {
	return newException(ErrAborted, messages[ErrAborted])
}

func NewAlreadyExistsError() *Exception {
	return newException(ErrAlreadyExists, messages[ErrAlreadyExists])
}

func NewResourceExhaustedError() *Exception {
	return newException(ErrResourceExhausted, messages[ErrResourceExhausted])
}

func NewCancelledError() *Exception {
	return newException(ErrCancelled, messages[ErrCancelled])
}

func NewDataLossError() *Exception {
	return newException(ErrDataLoss, messages[ErrDataLoss])
}

func NewUnknownError() *Exception {
	return newException(ErrUnknown, messages[ErrUnknown])
}

func NewInternalError() *Exception {
	return newException(ErrInternal, messages[ErrInternal])
}

func NewNotImplementedError() *Exception {
	return newException(ErrNotImplemented, messages[ErrNotImplemented])
}

func NewUnavailableError() *Exception {
	return newException(ErrUnavailable, messages[ErrUnavailable])
}

func NewDeadlineExceededError() *Exception {
	return newException(ErrDeadlineExceeded, messages[ErrDeadlineExceeded])
}
