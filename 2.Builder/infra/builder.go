package infra

import (
	"fmt"
	"net/http"
	"strings"
)

type HTTPRequestBuilder struct {
	Method string
	URL    string
	Header http.Header
	Body   string
	err    error
}

func NewHTTPRequestBuilder() *HTTPRequestBuilder {
	return &HTTPRequestBuilder{
		Header: make(http.Header),
	}
}

func (b *HTTPRequestBuilder) WithMethod(method string) *HTTPRequestBuilder {
	if strings.TrimSpace(method) == "" {
		b.err = fmt.Errorf("method cannot be empty")
	}

	b.Method = strings.ToUpper(method)
	return b
}

func (b *HTTPRequestBuilder) WithURL(url string) *HTTPRequestBuilder {
	if strings.TrimSpace(url) == "" {
		b.err = fmt.Errorf("url cannot be empty")
	}

	b.URL = url
	return b
}

func (b *HTTPRequestBuilder) WithHeader(key, value string) *HTTPRequestBuilder {
	b.Header.Add(key, value)
	return b
}

func (b *HTTPRequestBuilder) WithBody(body string) *HTTPRequestBuilder {
	b.Body = body
	return b
}

func (b *HTTPRequestBuilder) Build() (*http.Request, error) {
	if b.err != nil {
		// error occurred in one of the builder's steps
		return nil, b.err
	}

	req, err := http.NewRequest(b.Method, b.URL, strings.NewReader(b.Body))
	if err != nil {
		return nil, err
	}

	req.Header = b.Header

	return req, nil
}
