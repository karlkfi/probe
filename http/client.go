package http

import (
	"crypto/tls"
	"net/http"
)

func NewInsecureClient() *http.Client {
	tlsConfig := &tls.Config{InsecureSkipVerify: true}
	transport := &http.Transport{TLSClientConfig: tlsConfig}
	return &http.Client{Transport: transport}
}
