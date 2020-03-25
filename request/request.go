// Package request wraps http.NewRequest exposing useful functions
// (i.e. WithMethod, WithBaseURL)
package request

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

// RequestClient struct required for an HTTP Request to nibss sandbox.
type RequestClient struct {
	method      string
	resourceURL *url.URL
	base        string
	path        string
	queries     url.Values
	headers     http.Header
	body        *bytes.Buffer
}

type Option func(*RequestClient)

const (
	defaultMethod  = http.MethodGet
	defaultBaseURL = "https://sandboxapi.fsi.ng"
)

var (
	defaultHeaders = map[string]string{
		"Accept":         "application/json",
		"Content-Type":   "application/json",
		"SIGNATURE_METH": "SHA256",
	}
)

// WithBaseURL replace HTTP Request URL.
func WithBaseURL(b string) Option {
	return func(r *RequestClient) {
		r.base = b
	}
}

// WithPath replace HTTP Request URL Path.
func WithPath(path string) Option {
	return func(r *RequestClient) {
		r.path = path
	}
}

// WithMethod replace HTTP Request Method.
func WithMethod(method string) Option {
	return func(r *RequestClient) {
		r.method = method
	}
}

// WithDefaultHeaders use default HTTP Request headers defined.
// "Accept":         "application/json"
// "Content-Type":   "application/json"
// "SIGNATURE_METH": "SHA256"
func WithDefaultHeaders() Option {
	return func(r *RequestClient) {
		for key, value := range defaultHeaders {
			r.headers.Add(key, value)
		}
	}
}

// WithHeader set an HTTP Request Header.
func WithHeader(key, value string) Option {
	return func(r *RequestClient) {
		r.headers.Set(key, value)
	}
}

// WithBody set an HTTP Request Body.
func WithBody(b []byte) Option {
	return func(r *RequestClient) {
		r.headers.Add("Content-Length", strconv.Itoa(len(b)))
		r.body = bytes.NewBuffer(b)
	}
}

// WithBody set an HTTP Request Form Body.
func WithFormBody(data map[string]string) Option {
	form := url.Values{}

	for k, v := range data {
		form.Set(k, v)
	}

	return WithBody([]byte(form.Encode()))
}

// WithQuery add an HTTP Request URL Query.
func WithQuery(key, value string) Option {
	return func(r *RequestClient) {
		r.queries.Add(key, value)
	}
}

// WithQueries add multiple HTTP Request URL Query.
func WithQueries(queries map[string]string) Option {
	return func(r *RequestClient) {
		for key, value := range queries {
			r.queries.Add(key, value)
		}
	}
}

func buildRaw(base, path string, queries url.Values) string {
	q := queries.Encode()

	if q != "" {
		q = fmt.Sprintf("?%s", q)
	}

	return fmt.Sprintf("%s/%s%s", base, path, q)
}

// New generates RequestClient with overriding options.
// It returns RequestClient and any error encountered.
func New(opts ...Option) (RequestClient, error) {
	r := RequestClient{
		method:  defaultMethod,
		base:    defaultBaseURL,
		body:    bytes.NewBuffer(nil),
		queries: make(url.Values),
		headers: make(http.Header),
	}

	for _, opt := range opts {
		opt(&r)
	}

	apiURL, err := url.Parse(buildRaw(r.base, r.path, r.queries))

	if err != nil {
		return RequestClient{}, err
	}

	r.resourceURL = apiURL

	return r, nil
}

// String returns encoded HTTP Request URL.
func (r RequestClient) String() string {
	return r.resourceURL.String()
}

// Make sends an HTTP Request and returns an HTTP Response.
// It returns an HTTP Response and any error encountered.
func (r RequestClient) Make(responseFunc func(*http.Response) error) (*http.Response, error) {
	httpClient := &http.Client{}
	req, err := http.NewRequest(r.method, r.String(), r.body)

	if err != nil {
		return nil, err
	}

	if len(r.headers) > 0 {
		for k := range r.headers {
			req.Header.Add(k, r.headers.Get(k))
		}
	}

	res, err := httpClient.Do(req)

	if err != nil {
		return nil, err
	}

	if responseFunc != nil {
		err = responseFunc(res)
	}

	return res, err
}
