package httpService

import (
	"net/http"
	"time"
)

type HttpClient struct {
	*http.Client
	host    string
	name    string
	timeout time.Duration
}

func NewHttpClient(host, name string, timeout time.Duration) *HttpClient {
	return &HttpClient{
		Client:  &http.Client{},
		host:    host,
		name:    name,
		timeout: timeout,
	}
}
