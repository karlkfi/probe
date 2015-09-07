package tcp

import (
	"net"

	"github.com/golang/glog"

	"github.com/karlkfi/probe/errors"
)

// Probe opens and closes a TCP connection to the given URL.
func Probe(address string, dialer Dialer) errors.ProbeError {
	glog.V(2).Infof("Connection Address: %s", address)

	conn, err := dialer.Dial("tcp", address)
	if err != nil {
		pErr := &probeError{
			address: address,
			message: "probe/tcp: dial error",
			cause:   err,
		}
		glog.V(2).Infof("Connection Errored (address=%s): %#v", address, err)

		//unwrap *net.OpError
		if oErr, ok := err.(*net.OpError); ok && oErr.Err != nil {
			glog.V(1).Infof("Net %q Error (address=%s): %#v", oErr.Op, address, oErr.Err)
			err = oErr.Err
		}

		// detect timeouts
		if tErr, ok := err.(errors.TimeoutableError); ok && tErr.Timeout() {
			glog.V(1).Infof("Connection Timed Out (address=%s): %#v", address, err)
			pErr.timeout = true
		}
		return pErr
	}

	glog.V(2).Infof("Connection Openned (address=%s)", address)

	err = conn.Close()
	if err != nil {
		return &probeError{
			address: address,
			message: "probe/tcp: failed to close connection",
			cause:   err,
		}
	}

	glog.V(2).Infof("Connection Closed (address=%s)", address)

	return nil
}
