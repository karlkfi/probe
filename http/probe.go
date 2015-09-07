package http

import (
	"io/ioutil"
	"net"
	"net/http"
	"net/url"

	"github.com/golang/glog"

	"github.com/karlkfi/probe/errors"
)

// Probe sends an HTTP GET request to the provided address
// and validates that the response status code is between 200 (inclusive) and 400 (exclusive).
func Probe(address string, client Getter) errors.HTTPProbeError {
	glog.V(2).Infof("Request Address: %s", address)

	res, err := client.Get(address)
	if err != nil {
		pErr := &probeError{
			address: address,
			message: "probe/http: get error",
			cause:   err,
		}
		if res != nil {
			pErr.statusCode = res.StatusCode
		}
		glog.V(1).Infof("Request Errored (address=%s): %#v", address, err)

		// unwrap *url.Error
		if uErr, ok := err.(*url.Error); ok && uErr.Err != nil {
			glog.V(1).Infof("URL %q Error (address=%s): %#v", uErr.Op, address, uErr.Err)
			err = uErr.Err
		}

		//unwrap *net.OpError
		if oErr, ok := err.(*net.OpError); ok && oErr.Err != nil {
			glog.V(1).Infof("Net %q Error (address=%s): %#v", oErr.Op, address, oErr.Err)
			err = oErr.Err
		}

		// detect timeouts
		if tErr, ok := err.(errors.TimeoutableError); ok && tErr.Timeout() {
			glog.V(1).Infof("Request Timed Out (address=%s): %#v", address, tErr)
			pErr.timeout = true
		}
		return pErr
	}

	glog.V(2).Infof("Response Status (address=%s): %s", address, res.Status)

	body := ""
	if res.Body != nil {
		defer res.Body.Close()
		b, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return &probeError{
				address:    address,
				statusCode: res.StatusCode,
				body:       body,
				message:    "probe/http: failed to read response body",
				cause:      err,
			}
		}
		body = string(b)
		glog.V(3).Infof("Response Body (address=%s):\n%s", address, body)
	}

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		return &probeError{
			address:    address,
			statusCode: res.StatusCode,
			body:       body,
			message:    "probe/http: invalid status code",
		}
	}

	return nil
}
