package tcp

import (
	"fmt"
)

// probeError satisfies the errors.ProbeError interface.
type probeError struct {
	address string
	timeout bool
	message string
	cause   error
}

func (pe *probeError) Address() string {
	return pe.address
}

func (pe *probeError) Cause() error {
	return pe.cause
}

func (pe *probeError) Timeout() bool {
	return pe.timeout
}

func (pe *probeError) Error() string {
	msg := fmt.Sprintf(
		"%s (address=%s, timeout=%t)",
		pe.message,
		pe.address,
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
