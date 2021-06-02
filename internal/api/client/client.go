// Package client implement client requests to plantbook-server
package client

import (
	"crypto/tls"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/pkg/errors"
)

const (
	maxIdleConns int = 10
)

type API struct {
	baseURL *url.URL
	handler *http.Client
}

// New builder for API
func New(uri string, timeout time.Duration) (*API, error) {
	u, err := url.Parse(uri)
	if err != nil {
		return nil, errors.WithMessage(err, "pass wrong uri")
	}
	api := &API{
		baseURL: u,
	}
	// nolint:gosec
	cfg := &tls.Config{
		InsecureSkipVerify: true,
	}
	// set timeouts and border for opened connections
	api.handler = &http.Client{Timeout: timeout, Transport: &http.Transport{
		MaxIdleConns:    maxIdleConns,
		IdleConnTimeout: timeout,
		TLSClientConfig: cfg,
	}}
	return api, nil
}

// helpers

func (a *API) urlWithQuery(path, query string) string {
	res := a.url() + "/" + strings.TrimLeft(path, "/")
	if query != "" {
		res += "?" + query
	}
	return res
}

func (a *API) url() string {
	return a.baseURL.Scheme + "://" + a.baseURL.Host
}
