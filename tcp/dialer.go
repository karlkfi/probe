package tcp

import (
	"net"
)

// Dialer opens and returns a connection to a network address.
// Satisfied by net.Dialer
type Dialer interface {
	Dial(network, address string) (net.Conn, error)
}

func NewDialer() *net.Dialer {
	return &net.Dialer{}
}
