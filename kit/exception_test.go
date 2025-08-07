package kit

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestException_BasicFunctionality(t *testing.T) {
	t.Run("NewException", func(t *testing.T) {
		ex := NewException()
		assert.Equal(t, OK, ex.Code())
		assert.Equal(t, "", ex.Info())
		assert.Equal(t, "", ex.Desc())
	})

	t.Run("WithCode", func(t *testing.T) {
		ex := NewException().WithCode(ErrNotFound)
		assert.Equal(t, ErrNotFound, ex.Code())
	})

	t.Run("WithInfo", func(t *testing.T) {
		info := "Resource not found"
		ex := NewException().WithInfo(info)
		assert.Equal(t, info, ex.Info())
	})

	t.Run("WithErr", func(t *testing.T) {
		err := errors.New("database connection failed")
		ex := NewException().WithErr(err)
		assert.Equal(t, err.Error(), ex.Desc())
	})

	t.Run("WithErr nil", func(t *testing.T) {
		ex := NewException().WithErr(nil)
		assert.Equal(t, "", ex.Desc())
	})
}

func TestException_ErrorMethod(t *testing.T) {
	t.Run("Error returns desc when desc is not empty", func(t *testing.T) {
		desc := "detailed error description"
		ex := &Exception{desc: desc, info: "info", code: ErrInternal}
		assert.Equal(t, desc, ex.Error())
	})

	t.Run("Error returns info when desc is empty", func(t *testing.T) {
		info := "error info"
		ex := &Exception{info: info, code: ErrInternal}
		assert.Equal(t, info, ex.Error())
	})

	t.Run("Error returns unknown message when both desc and info are empty", func(t *testing.T) {
		ex := &Exception{code: ErrInternal}
		assert.Equal(t, Messages[ErrUnknown], ex.Error())
	})
}

func TestException_BusinessErrorInterface(t *testing.T) {
	ex := NewInvalidArgumentError()

	// Test that Exception implements BusinessError interface
	var businessErr BusinessError = ex
	assert.Equal(t, ErrInvalidArgument, businessErr.Code())
	assert.Equal(t, Messages[ErrInvalidArgument], businessErr.Info())
	assert.Equal(t, "", businessErr.Desc())
}

func TestException_ErrorInterface(t *testing.T) {
	ex := NewInternalError()

	// Test that Exception implements error interface
	var err error = ex
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), Messages[ErrInternal])
}

func TestPredefinedExceptions(t *testing.T) {
	testCases := []struct {
		name     string
		factory  func() *Exception
		expected int
	}{
		{"InvalidArgument", NewInvalidArgumentError, ErrInvalidArgument},
		{"FailedPrecondition", NewFailedPreconditionError, ErrFailedPrecondition},
		{"OutOfRange", NewOutOfRangeError, ErrOutOfRange},
		{"Unauthenticated", NewUnauthenticatedError, ErrUnauthenticated},
		{"PermissionDenied", NewPermissionDeniedError, ErrPermissionDenied},
		{"NotFound", NewNotFoundError, ErrNotFound},
		{"Aborted", NewAbortedError, ErrAborted},
		{"AlreadyExists", NewAlreadyExistsError, ErrAlreadyExists},
		{"ResourceExhausted", NewResourceExhaustedError, ErrResourceExhausted},
		{"Cancelled", NewCancelledError, ErrCancelled},
		{"DataLoss", NewDataLossError, ErrDataLoss},
		{"Unknown", NewUnknownError, ErrUnknown},
		{"Internal", NewInternalError, ErrInternal},
		{"NotImplemented", NewNotImplementedError, ErrNotImplemented},
		{"Unavailable", NewUnavailableError, ErrUnavailable},
		{"DeadlineExceeded", NewDeadlineExceededError, ErrDeadlineExceeded},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ex := tc.factory()
			assert.Equal(t, tc.expected, ex.Code())

			// Check if message exists in Messages map
			if msg, exists := Messages[tc.expected]; exists {
				assert.Equal(t, msg, ex.Info())
			}
		})
	}
}

func TestException_ChainedMethods(t *testing.T) {
	t.Run("Chained method calls", func(t *testing.T) {
		originalErr := errors.New("original error")
		ex := NewException().
			WithCode(ErrNotFound).
			WithInfo("Custom info message").
			WithErr(originalErr)

		assert.Equal(t, ErrNotFound, ex.Code())
		assert.Equal(t, "Custom info message", ex.Info())
		assert.Equal(t, originalErr.Error(), ex.Desc())
	})

	t.Run("Overwriting values", func(t *testing.T) {
		ex := NewInvalidArgumentError().
			WithCode(ErrNotFound).
			WithInfo("New info")

		assert.Equal(t, ErrNotFound, ex.Code())
		assert.Equal(t, "New info", ex.Info())
	})
}

func TestException_ErrorMessages(t *testing.T) {
	t.Run("All error codes have Messages", func(t *testing.T) {
		requiredMessages := []int{
			ErrInvalidArgument,
			ErrUnauthenticated,
			ErrPermissionDenied,
			ErrNotFound,
			ErrAlreadyExists,
			ErrResourceExhausted,
			ErrCancelled,
			ErrInternal,
			ErrNotImplemented,
			ErrUnavailable,
			ErrDeadlineExceeded,
		}

		for _, code := range requiredMessages {
			msg, exists := Messages[code]
			assert.True(t, exists, "Messages for error code %d should exist", code)
			assert.NotEmpty(t, msg, "Messages for error code %d should not be empty", code)
		}
	})
}
