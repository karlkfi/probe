package errors

import (
	"fmt"
)

type ProbeError interface {
	error
	fmt.Stringer
	Address() string
	Cause() error
	Timeout() bool
}

type HTTPProbeError interface {
	ProbeError
	StatusCode() int
	Body() string
}
