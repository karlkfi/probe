package errors

// TimeoutableError is satisfied by net/http.httpError, net/http.tlsHandshakeTimeoutError, and net.errTimeout
type TimeoutableError interface {
	error
	Timeout() bool
}
