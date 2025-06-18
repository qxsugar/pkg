package kit

// Error codes
// Reference: [Google's API Design Guide](https://www.bookstack.cn/read/API-design-guide/API-design-guide-07-%E9%94%99%E8%AF%AF.md)

const (
	OK                    = 0   // No error
	ErrInvalidArgument    = 400 // The client specified an invalid argument. Check error message and error details for more information.
	ErrFailedPrecondition = 400 // The request can not be executed in the current system state, such as deleting a non-empty directory.
	ErrOutOfRange         = 400 // The client specified an invalid range.
	ErrUnauthenticated    = 401 // The request did not pass authentication due to missing, invalid, or expired OAuth token.
	ErrPermissionDenied   = 403 // The client does not have enough permission. This could be because the OAuth token does not have the correct scope, the client does not have permission, or the client project has not enabled the API.
	ErrNotFound           = 404 // The specified resource could not be found, or the request was denied for reasons that are not disclosed (such as a whitelist).
	ErrAborted            = 409 // Concurrent conflict, such as read-modify-write conflict.
	ErrAlreadyExists      = 409 // The resource that the client tried to create already exists.
	ErrResourceExhausted  = 429 // Resource quota reached rate limit. The client should look for google.rpc.QuotaFailure error details for more information.
	ErrCancelled          = 499 // The client cancelled the request.
	ErrDataLoss           = 500 // Irrecoverable data loss or data corruption. The client should report the error to the user.
	ErrUnknown            = 500 // Unknown server error. Typically a server error.
	ErrInternal           = 500 // Internal server error. Typically a server error.
	ErrNotImplemented     = 501 // The server did not implement the API method.
	ErrUnavailable        = 503 // Service suspended. Typically the server has been shut down.
	ErrDeadlineExceeded   = 504 // The request deadline has been exceeded. If it happens repeatedly, consider reducing the complexity of the request.
)

var messages = map[int]string{
	ErrInvalidArgument:   "Invalid argument",
	ErrUnauthenticated:   "Invalid identity",
	ErrPermissionDenied:  "Insufficient permissions",
	ErrNotFound:          "Resource does not exist",
	ErrAlreadyExists:     "Resource already exists",
	ErrResourceExhausted: "System busy",
	ErrCancelled:         "Client cancelled request",
	ErrInternal:          "Internal error",
	ErrNotImplemented:    "Method not implemented",
	ErrUnavailable:       "Service suspended",
	ErrDeadlineExceeded:  "System unable to execute",
}
