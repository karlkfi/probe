package http

import (
	"net/http"
	"net/url"
	"time"
)

type Prober interface {
	Probe(url *url.URL) error
}

type prober struct {
	client *http.Client
}

func NewProber(client *http.Client) Prober {
	return &prober{
		client: client,
	}
}

func (pr *prober) Probe(url *url.URL) error {
	return Probe(url, pr.client)
}
