// Package wallet implements access to relianceHMO.wallet sandbox.
package wallet

import (
	"io/ioutil"
	"net/http"

	"github.com/enyata/innovation-sandbox-go/relianceHMO"
	req "github.com/enyata/innovation-sandbox-go/request"
)

// Fund sends an HTTP request to relianceHMO.wallet.fund sandbox API.
// It returns an HTTP response body string and any error encountered.
func Fund(c relianceHMO.RelianceHMOCredentials, data []byte, overrideOpts ...req.Option) (string, error) {
	var responseString string

	option := append([]req.Option{
		req.WithMethod("POST"),
		req.WithPath("relianceHMO/wallet/fund"),
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

// Transactions sends an HTTP request to relianceHMO.wallet.transactions sandbox API.
// It returns an HTTP response body string and any error encountered.
func Transactions(c relianceHMO.RelianceHMOCredentials, overrideOpts ...req.Option) (string, error) {
	var responseString string

	option := append([]req.Option{
		req.WithMethod("GET"),
		req.WithPath("relianceHMO/wallet/transactions"),
		req.WithHeader("Sandbox-Key", c.SandboxKey),
		req.WithHeader("Content-Type", "application/json"),
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
