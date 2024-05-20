package gohttpclient

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"time"
)

type httpClient struct {
	config HttpClientConfig
}

type HttpClientConfig struct {
	RequestTimeout time.Duration
}

func NewHttpClient(config HttpClientConfig) *httpClient {
	return &httpClient{
		config: config,
	}
}

func (c *httpClient) Post(ctx context.Context, url string, data any, headers map[string]string) (*http.Response, error) {
	body, err := encodeBody(data, headers)
	if err != nil {
		return nil, err
	}
	return c.Request(ctx, http.MethodPost, url, body, headers)
}

func (c *httpClient) Get(ctx context.Context, url string, params map[string]any, headers map[string]string) (*http.Response, error) {
	if params != nil {
		queryParamsString, err := urlEncode(params)
		if err != nil {
			return nil, err
		}
		url = fmt.Sprintf("%s?%s", url, queryParamsString)
	}

	return c.Request(ctx, http.MethodGet, url, nil, headers)
}

func (c *httpClient) Request(context context.Context, method string, url string, body []byte, headers map[string]string) (*http.Response, error) {
	bodyReader := bytes.NewReader(body)

	req, err := http.NewRequest(method, url, bodyReader)

	if err != nil {
		return nil, err
	}

	for key, val := range headers {
		req.Header.Set(key, val)
	}

	client := http.Client{
		Timeout: c.config.RequestTimeout,
	}

	response, err := client.Do(req)

	return response, err

}
