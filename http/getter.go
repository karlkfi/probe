package http

import (
	"net/http"
)

type Getter interface {
	Get(url string) (resp *http.Response, err error)
}
