package request

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

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
	defaultBaseURL = ""
)

var (
	defaultHeaders = map[string]string{
		"Accept":         "application/json",
		"Content-Type":   "application/json",
		"SIGNATURE_METH": "SHA256",
	}
)

func WithBaseURL(b string) Option {
	return func(r *RequestClient) {
		r.base = b
	}
}

func WithPath(path string) Option {
	return func(r *RequestClient) {
		r.path = path
	}
}

func WithMethod(method string) Option {
	return func(r *RequestClient) {
		r.method = method
	}
}

func WithDefaultHeaders() Option {
	return func(r *RequestClient) {
		for key, value := range defaultHeaders {
			r.headers.Add(key, value)
		}
	}
}

func WithHeader(key, value string) Option {
	return func(r *RequestClient) {
		r.headers.Set(key, value)
	}
}

func WithBody(b []byte) Option {
	return func(r *RequestClient) {
		r.headers.Add("Content-Length", strconv.Itoa(len(b)))
		r.body = bytes.NewBuffer(b)
	}
}

func WithQuery(key, value string) Option {
	return func(r *RequestClient) {
		r.queries.Add(key, value)
	}
}

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
		q = fmt.Sprintf("?q=%s", q)
	}

	return fmt.Sprintf("%s/%s%s", base, path, q)
}

func New(opts ...Option) (RequestClient, error) {
	r := RequestClient{
		method:  defaultMethod,
		base:    defaultBaseURL,
		body:    nil,
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

func (r RequestClient) String() string {
	return r.resourceURL.String()
}

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

	if responseFunc != nil {
		err = responseFunc(res)
	}

	return res, err
}
