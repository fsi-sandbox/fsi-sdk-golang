// Package token implements access to atlabs.token sandbox.
package token

import (
	"io/ioutil"
	"net/http"

	"github.com/enyata/innovation-sandbox-go/atlabs"
	req "github.com/enyata/innovation-sandbox-go/request"
)

// CreateCheckoutToken sends an HTTP request to atlabs.token sandbox API to retrieve
// checkout token.
// It returns an HTTP response body string and any error encountered.
func CreateCheckoutToken(c atlabs.AtlabsCredentials, data []byte, overrideOpts ...req.Option) (string, error) {
	var responseString string

	option := append([]req.Option{
		req.WithMethod("POST"),
		req.WithPath("atlabs/token/checkout"),
		req.WithHeader("Sandbox-Key", c.SandboxKey),
		req.WithHeader("Content-Type", "application/json"),
		req.WithBody(data),
	}, overrideOpts...)
	req, err := req.New(option...)

	if err != nil {
		return responseString, err
	}

	_, err = req.Make(func(response *http.Response) error {
		defer response.Body.Close()

		data, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return err
		}

		responseString = string(data)
		return err
	})

	return responseString, err
}
