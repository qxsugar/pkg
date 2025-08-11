package kit

// Error codes
// Reference: [Google's API Design Guide](https://www.bookstack.cn/read/API-design-guide/API-design-guide-07-%E9%94%99%E8%AF%AF.md)

const (
	OK                    = 0     // No error
	ErrInvalidArgument    = 40001 // The client specified an invalid argument. Check error message and error details for more information.
	ErrFailedPrecondition = 40002 // The request can not be executed in the current system state, such as deleting a non-empty directory.
	ErrOutOfRange         = 40003 // The client specified an invalid range.
	ErrUnauthenticated    = 40100 // The request did not pass authentication due to missing, invalid, or expired OAuth token.
	ErrPermissionDenied   = 40300 // The client does not have enough permission. This could be because the OAuth token does not have the correct scope, the client does not have permission, or the client project has not enabled the API.
	ErrNotFound           = 40400 // The specified resource could not be found, or the request was denied for reasons that are not disclosed (such as a whitelist).
	ErrAborted            = 40901 // Concurrent conflict, such as read-modify-write conflict.
	ErrAlreadyExists      = 40902 // The resource that the client tried to create already exists.
	ErrResourceExhausted  = 42900 // Resource quota reached rate limit. The client should look for google.rpc.QuotaFailure error details for more information.
	ErrCancelled          = 49900 // The client cancelled the request.
	ErrDataLoss           = 50001 // Irrecoverable data loss or data corruption. The client should report the error to the user.
	ErrUnknown            = 50002 // Unknown server error. Typically a server error.
	ErrInternal           = 50003 // Internal server error. Typically a server error.
	ErrNotImplemented     = 50100 // The server did not implement the API method.
	ErrUnavailable        = 50300 // Service suspended. Typically the server has been shut down.
	ErrDeadlineExceeded   = 50400 // The request deadline has been exceeded. If it happens repeatedly, consider reducing the complexity of the request.
)

var Messages = map[int]string{
	OK:                   "Success",
	ErrInvalidArgument:   "Invalid argument",
	ErrFailedPrecondition: "Failed precondition",
	ErrOutOfRange:        "Out of range",
	ErrUnauthenticated:   "Invalid identity",
	ErrPermissionDenied:  "Insufficient permissions",
	ErrNotFound:          "Resource does not exist",
	ErrAborted:           "Operation aborted",
	ErrAlreadyExists:     "Resource already exists",
	ErrResourceExhausted: "System busy",
	ErrCancelled:         "Client cancelled request",
	ErrDataLoss:          "Data loss occurred",
	ErrUnknown:           "Unknown error",
	ErrInternal:          "Internal error",
	ErrNotImplemented:    "Method not implemented",
	ErrUnavailable:       "Service suspended",
	ErrDeadlineExceeded:  "System unable to execute",
}

// HTTPStatusCodes maps business error codes to HTTP status codes
var HTTPStatusCodes = map[int]int{
	OK:                   200, // OK
	ErrInvalidArgument:   400, // Bad Request
	ErrFailedPrecondition: 400, // Bad Request
	ErrOutOfRange:        400, // Bad Request
	ErrUnauthenticated:   401, // Unauthorized
	ErrPermissionDenied:  403, // Forbidden
	ErrNotFound:          404, // Not Found
	ErrAborted:           409, // Conflict
	ErrAlreadyExists:     409, // Conflict
	ErrResourceExhausted: 429, // Too Many Requests
	ErrCancelled:         499, // Client Closed Request
	ErrDataLoss:          500, // Internal Server Error
	ErrUnknown:           500, // Internal Server Error
	ErrInternal:          500, // Internal Server Error
	ErrNotImplemented:    501, // Not Implemented
	ErrUnavailable:       503, // Service Unavailable
	ErrDeadlineExceeded:  504, // Gateway Timeout
}
