// Package airtime implements access to atlabs.airtime sandbox.
package airtime

import (
	"io/ioutil"
	"net/http"

	"github.com/enyata/innovation-sandbox-go/atlabs"
	req "github.com/enyata/innovation-sandbox-go/request"
)

// SendAirtime sends an HTTP request to atlabs.airtime sandbox API.
// It returns an HTTP response body string and any error encountered.
func SendAirtime(c atlabs.AtlabsCredentials, data []byte, overrideOpts ...req.Option) (string, error) {
	var responseString string

	option := append([]req.Option{
		req.WithMethod("POST"),
		req.WithPath("atlabs/airtime/send"),
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
