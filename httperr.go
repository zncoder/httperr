package httperr

import (
	"net/http"
)

type httpErr int

func (he httpErr) Error() string {
	return http.StatusText(int(he))
}

func Code(err error) (int, bool) {
	he, ok := err.(httpErr)
	return int(he), ok
}

const (
	ErrBadRequest = httpErr(http.StatusBadRequest)
	// StatusUnauthorized is actually Unauthenticated
	ErrUnauthorized                  = httpErr(http.StatusUnauthorized)
	ErrUnauthenticated               = httpErr(http.StatusUnauthorized)
	ErrPaymentRequired               = httpErr(http.StatusPaymentRequired)
	ErrForbidden                     = httpErr(http.StatusForbidden)
	ErrNotFound                      = httpErr(http.StatusNotFound)
	ErrMethodNotAllowed              = httpErr(http.StatusMethodNotAllowed)
	ErrNotAcceptable                 = httpErr(http.StatusNotAcceptable)
	ErrProxyAuthRequired             = httpErr(http.StatusProxyAuthRequired)
	ErrRequestTimeout                = httpErr(http.StatusRequestTimeout)
	ErrConflict                      = httpErr(http.StatusConflict)
	ErrGone                          = httpErr(http.StatusGone)
	ErrLengthRequired                = httpErr(http.StatusLengthRequired)
	ErrPreconditionFailed            = httpErr(http.StatusPreconditionFailed)
	ErrRequestEntityTooLarge         = httpErr(http.StatusRequestEntityTooLarge)
	ErrRequestURITooLong             = httpErr(http.StatusRequestURITooLong)
	ErrUnsupportedMediaType          = httpErr(http.StatusUnsupportedMediaType)
	ErrRequestedRangeNotSatisfiable  = httpErr(http.StatusRequestedRangeNotSatisfiable)
	ErrExpectationFailed             = httpErr(http.StatusExpectationFailed)
	ErrTeapot                        = httpErr(http.StatusTeapot)
	ErrMisdirectedRequest            = httpErr(http.StatusMisdirectedRequest)
	ErrUnprocessableEntity           = httpErr(http.StatusUnprocessableEntity)
	ErrLocked                        = httpErr(http.StatusLocked)
	ErrFailedDependency              = httpErr(http.StatusFailedDependency)
	ErrTooEarly                      = httpErr(http.StatusTooEarly)
	ErrUpgradeRequired               = httpErr(http.StatusUpgradeRequired)
	ErrPreconditionRequired          = httpErr(http.StatusPreconditionRequired)
	ErrTooManyRequests               = httpErr(http.StatusTooManyRequests)
	ErrRequestHeaderFieldsTooLarge   = httpErr(http.StatusRequestHeaderFieldsTooLarge)
	ErrUnavailableForLegalReasons    = httpErr(http.StatusUnavailableForLegalReasons)
	ErrInternalServerError           = httpErr(http.StatusInternalServerError)
	ErrNotImplemented                = httpErr(http.StatusNotImplemented)
	ErrBadGateway                    = httpErr(http.StatusBadGateway)
	ErrServiceUnavailable            = httpErr(http.StatusServiceUnavailable)
	ErrGatewayTimeout                = httpErr(http.StatusGatewayTimeout)
	ErrHTTPVersionNotSupported       = httpErr(http.StatusHTTPVersionNotSupported)
	ErrVariantAlsoNegotiates         = httpErr(http.StatusVariantAlsoNegotiates)
	ErrInsufficientStorage           = httpErr(http.StatusInsufficientStorage)
	ErrLoopDetected                  = httpErr(http.StatusLoopDetected)
	ErrNotExtended                   = httpErr(http.StatusNotExtended)
	ErrNetworkAuthenticationRequired = httpErr(http.StatusNetworkAuthenticationRequired)
)
