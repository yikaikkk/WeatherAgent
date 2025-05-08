package httpService

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

type HttpClient struct {
	*http.Client
	host    string
	timeout time.Duration
}

func NewHttpClient(host string, timeout time.Duration) *HttpClient {
	return &HttpClient{
		Client:  &http.Client{},
		host:    host,
		timeout: timeout,
	}
}

func (c *HttpClient) Get(path string) ([]byte, error) {
	req, err := http.NewRequest("GET", c.host+path, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "*")
	c.Timeout = c.timeout
	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	rsp, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return rsp, nil
}

func (c *HttpClient) Post(path string, body []byte) ([]byte, error) {
	var buf bytes.Buffer
	if body != nil {
		if err := json.NewEncoder(&buf).Encode(body); err != nil {
			return nil, err
		}
	}
	req, err := http.NewRequest("POST", c.host+path, &buf)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	c.Timeout = c.timeout
	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	rsp, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return rsp, nil
}
