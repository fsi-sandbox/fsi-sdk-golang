// Package token implements access to union.oauth sandbox.
package token

import (
	"io/ioutil"
	"net/http"

	req "github.com/enyata/innovation-sandbox-go/request"
	"github.com/enyata/innovation-sandbox-go/union"
)

// Token sends an HTTP request to union.oauth sandbox API to retrieve
// oauth credetials.
// It returns an HTTP response body string and any error encountered.
func Token(c union.UnionCredentials, data map[string]string, overrideOpts ...req.Option) (string, error) {
	var responseString string

	option := append([]req.Option{
		req.WithMethod("POST"),
		req.WithPath("union/oauth/token"),
		req.WithHeader("Sandbox-Key", c.SandboxKey),
		req.WithHeader("Content-Type", "application/x-www-form-urlencoded"),
		req.WithFormBody(data),
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
