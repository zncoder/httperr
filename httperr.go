package httperr

import (
	"errors"
	"io"
	"net/http"
	"os"

	"github.com/zncoder/assert"
)

type httpErr int

func (he httpErr) Error() string {
	return http.StatusText(int(he))
}

func Code(err error) (int, bool) {
	he, ok := err.(httpErr)
	return int(he), ok
}

func WithStatus(code int) error {
	if code == 200 {
		return nil
	}
	return httpErr(code)
}

func Error(w http.ResponseWriter, err error, format string, args ...interface{}) {
	assert.OK(err != nil)
	he, ok := err.(httpErr)
	if !ok {
		he = InternalServerError
	}
	http.Error(w, err.Error(), int(he))
}

func Is(err error, wants ...error) bool {
	for _, e := range wants {
		if errors.Is(err, e) {
			return true
		}
	}
	return false
}

func IsEOF(err error) bool {
	return Is(err, io.EOF)
}

func IsNotFound(err error) bool {
	return Is(err, NotFound, os.ErrNotExist)
}

const (
	BadRequest = httpErr(http.StatusBadRequest)
	// StatusUnauthorized is actually Unauthenticated
	Unauthenticated               = httpErr(http.StatusUnauthorized)
	Unauthorized                  = Unauthenticated // alias
	PaymentRequired               = httpErr(http.StatusPaymentRequired)
	Forbidden                     = httpErr(http.StatusForbidden)
	NotFound                      = httpErr(http.StatusNotFound)
	MethodNotAllowed              = httpErr(http.StatusMethodNotAllowed)
	NotAcceptable                 = httpErr(http.StatusNotAcceptable)
	ProxyAuthRequired             = httpErr(http.StatusProxyAuthRequired)
	RequestTimeout                = httpErr(http.StatusRequestTimeout)
	Conflict                      = httpErr(http.StatusConflict)
	Gone                          = httpErr(http.StatusGone)
	LengthRequired                = httpErr(http.StatusLengthRequired)
	PreconditionFailed            = httpErr(http.StatusPreconditionFailed)
	RequestEntityTooLarge         = httpErr(http.StatusRequestEntityTooLarge)
	RequestURITooLong             = httpErr(http.StatusRequestURITooLong)
	UnsupportedMediaType          = httpErr(http.StatusUnsupportedMediaType)
	RequestedRangeNotSatisfiable  = httpErr(http.StatusRequestedRangeNotSatisfiable)
	ExpectationFailed             = httpErr(http.StatusExpectationFailed)
	Teapot                        = httpErr(http.StatusTeapot)
	MisdirectedRequest            = httpErr(http.StatusMisdirectedRequest)
	UnprocessableEntity           = httpErr(http.StatusUnprocessableEntity)
	Locked                        = httpErr(http.StatusLocked)
	FailedDependency              = httpErr(http.StatusFailedDependency)
	TooEarly                      = httpErr(http.StatusTooEarly)
	UpgradeRequired               = httpErr(http.StatusUpgradeRequired)
	PreconditionRequired          = httpErr(http.StatusPreconditionRequired)
	TooManyRequests               = httpErr(http.StatusTooManyRequests)
	RequestHeaderFieldsTooLarge   = httpErr(http.StatusRequestHeaderFieldsTooLarge)
	UnavailableForLegalReasons    = httpErr(http.StatusUnavailableForLegalReasons)
	InternalServerError           = httpErr(http.StatusInternalServerError)
	NotImplemented                = httpErr(http.StatusNotImplemented)
	BadGateway                    = httpErr(http.StatusBadGateway)
	ServiceUnavailable            = httpErr(http.StatusServiceUnavailable)
	GatewayTimeout                = httpErr(http.StatusGatewayTimeout)
	HTTPVersionNotSupported       = httpErr(http.StatusHTTPVersionNotSupported)
	VariantAlsoNegotiates         = httpErr(http.StatusVariantAlsoNegotiates)
	InsufficientStorage           = httpErr(http.StatusInsufficientStorage)
	LoopDetected                  = httpErr(http.StatusLoopDetected)
	NotExtended                   = httpErr(http.StatusNotExtended)
	NetworkAuthenticationRequired = httpErr(http.StatusNetworkAuthenticationRequired)
)
