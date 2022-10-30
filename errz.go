package faceitgo

import "fmt"

type (
	HTTPErr struct {
		statusCode int
		message    string
		cause      error
	}
)

func (e *HTTPErr) Error() string {
	return fmt.Sprintf("HTTP error %d: %s", e.statusCode, e.message)
}

func (e *HTTPErr) Unwrap() error {
	return e.cause
}

func (e *HTTPErr) Wrap(cause error) *HTTPErr {
	e.cause = cause
	return e
}

func (e *HTTPErr) Is(target error) bool {
	t, ok := target.(*HTTPErr)
	if !ok {
		return false
	}

	return e.statusCode == t.statusCode
}

func (e *HTTPErr) StatusCode() int {
	return e.statusCode
}

func (e *HTTPErr) Msg(msg string) *HTTPErr {
	e.message = msg
	return e
}

func (e *HTTPErr) Msgf(format string, a ...interface{}) *HTTPErr {
	e.message = fmt.Sprintf(format, a...)
	return e
}

var (
	errNotFound            = HTTPErr{statusCode: 404, message: "not found"}             // 404 Not Found
	errBadRequest          = HTTPErr{statusCode: 400, message: "bad request"}           // 400 Bad Request
	errUnauthorized        = HTTPErr{statusCode: 401, message: "unauthorized"}          // 401 Unauthorized
	errForbidden           = HTTPErr{statusCode: 403, message: "forbidden"}             // 403 Forbidden
	errTooManyRequests     = HTTPErr{statusCode: 429, message: "too many requests"}     // 429 Too Many Requests
	errInternalServerError = HTTPErr{statusCode: 500, message: "internal server error"} // 500 Internal Server Error
	errBadGateway          = HTTPErr{statusCode: 502, message: "bad gateway"}           // 502 Bad Gateway
	errServiceUnavailable  = HTTPErr{statusCode: 503, message: "service unavailable"}   // 503 Service Unavailable
	errGatewayTimeout      = HTTPErr{statusCode: 504, message: "gateway timeout"}       // 504 Gateway Timeout
)

func ErrNotFound() *HTTPErr {
	err := errNotFound
	return &err
}

func ErrBadRequest() *HTTPErr {
	err := errBadRequest
	return &err
}

func ErrUnauthorized() *HTTPErr {
	err := errUnauthorized
	return &err
}

func ErrForbidden() *HTTPErr {
	err := errForbidden
	return &err
}

func ErrTooManyRequests() *HTTPErr {
	err := errTooManyRequests
	return &err
}

func ErrInternalServerError() *HTTPErr {
	err := errInternalServerError
	return &err
}

func ErrBadGateway() *HTTPErr {
	err := errBadGateway
	return &err
}

func ErrServiceUnavailable() *HTTPErr {
	err := errServiceUnavailable
	return &err
}

func ErrGatewayTimeout() *HTTPErr {
	err := errGatewayTimeout
	return &err
}

func IsNotFound(err error) bool {
	return errNotFound.Is(err)
}

func IsBadRequest(err error) bool {
	return errBadRequest.Is(err)
}

func IsUnauthorized(err error) bool {
	return errUnauthorized.Is(err)
}

func IsForbidden(err error) bool {
	return errForbidden.Is(err)
}

func IsTooManyRequests(err error) bool {
	return errTooManyRequests.Is(err)
}

func IsInternalServerError(err error) bool {
	return errInternalServerError.Is(err)
}

func IsBadGateway(err error) bool {
	return errBadGateway.Is(err)
}

func IsServiceUnavailable(err error) bool {
	return errServiceUnavailable.Is(err)
}

func IsGatewayTimeout(err error) bool {
	return errGatewayTimeout.Is(err)
}
