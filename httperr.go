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
	ErrBadRequest      = httpErr(http.StatusBadRequest)
	ErrUnauthenticated = httpErr(http.StatusUnauthorized)
	ErrForbidden       = httpErr(http.StatusForbidden)
)
