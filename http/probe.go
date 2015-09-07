package http

import (
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/golang/glog"
)

// Probe performs a GET request to the provided url.
// If the HTTP response code is successful (i.e. 400 > code >= 200), Probe returns without error.
// Otherwise an error is returned.
func Probe(url *url.URL, client Getter) ProbeError {
	urlStr := url.String()
	glog.V(2).Infof("Request URL: %s", urlStr)

	res, err := client.Get(urlStr)
	if err != nil {
		pErr := &probeError{
			url:     url,
			message: "probe/http: get error",
			cause:   err,
		}
		if res != nil {
			pErr.statusCode = res.StatusCode
		}
		// detect timeouts
		if tErr, ok := err.(timeoutable); ok && tErr.Timeout() {
			pErr.timeout = true
		}
		return pErr
	}

	glog.V(2).Infof("Response Status: %s", res.Status)

	body := ""
	if res.Body != nil {
		defer res.Body.Close()
		b, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return &probeError{
				url:        url,
				statusCode: res.StatusCode,
				body:       body,
				message:    "probe/http: failed to read response body",
				cause:      err,
			}
		}
		body = string(b)
		glog.V(3).Infof("Response Body:\n%s", body)
	}

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		return &probeError{
			url:        url,
			statusCode: res.StatusCode,
			body:       body,
			message:    "probe/http: invalid status code",
		}
	}

	return nil
}
