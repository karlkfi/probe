package http

import (
	"crypto/tls"
	"net/http"
)

// Getter sends a request to an address and returns the response.
// Satisfied by net/http.Client
type Getter interface {
	Get(address string) (resp *http.Response, err error)
}

func NewInsecureClient() *http.Client {
	tlsConfig := &tls.Config{InsecureSkipVerify: true}
	transport := &http.Transport{TLSClientConfig: tlsConfig}
	return &http.Client{Transport: transport}
}
