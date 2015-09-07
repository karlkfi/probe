package http

import (
	"fmt"
	"net/url"
)

// timeoutable matches net/http.httpError and net/http.tlsHandshakeTimeoutError
type timeoutable interface {
	error
	Timeout() bool
}

type ProbeError interface {
	error
	fmt.Stringer
	URL() *url.URL
	StatusCode() int
	Body() string
	Cause() error
	Timeout() bool
}

type probeError struct {
	url        *url.URL
	statusCode int
	body       string
	timeout    bool
	message    string
	cause      error
}

func (pe *probeError) URL() *url.URL {
	return pe.url
}

func (pe *probeError) StatusCode() int {
	return pe.statusCode
}

func (pe *probeError) Body() string {
	return pe.body
}

func (pe *probeError) Cause() error {
	return pe.cause
}

func (pe *probeError) Timeout() bool {
	return pe.timeout
}

func (pe *probeError) Error() string {
	msg := fmt.Sprintf(
		"%s (url=%s, status-code=%d, timeout=%t)",
		pe.message,
		pe.url,
		pe.statusCode,
		pe.timeout,
	)
	if pe.cause != nil {
		msg = fmt.Sprintf("%s: %v", msg, pe.cause)
	}
	return msg
}

func (pe *probeError) String() string {
	return pe.Error()
}
