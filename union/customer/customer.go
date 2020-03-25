// Package customer implements access to union.customer sandbox.
package customer

import (
	"io/ioutil"
	"net/http"

	req "github.com/enyata/innovation-sandbox-go/request"
	"github.com/enyata/innovation-sandbox-go/union"
)

// CustomerEnquiry sends an HTTP request to union.cbacustomerenquiry sandbox API.
// It returns an HTTP response body string and any error encountered.
func CustomerEnquiry(c union.UnionCredentials, queries map[string]string, data []byte, overrideOpts ...req.Option) (string, error) {
	var responseString string

	option := append([]req.Option{
		req.WithMethod("POST"),
		req.WithPath("union/secured/cbacustomerenquiry"),
		req.WithHeader("Sandbox-Key", c.SandboxKey),
		req.WithHeader("Content-Type", "application/json"),
		req.WithQueries(queries),
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

// CustomerAccountEnquiry sends an HTTP request to union.cbacustomeraccountenquiry sandbox API.
// It returns an HTTP response body string and any error encountered.
func CustomerAccountEnquiry(c union.UnionCredentials, queries map[string]string, data []byte, overrideOpts ...req.Option) (string, error) {
	var responseString string

	option := append([]req.Option{
		req.WithMethod("POST"),
		req.WithPath("union/secured/cbacustomeraccountenquiry"),
		req.WithHeader("Sandbox-Key", c.SandboxKey),
		req.WithHeader("Content-Type", "application/json"),
		req.WithQueries(queries),
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

// AccountEnquiry sends an HTTP request to union.cbaaccountenquiry sandbox API.
// It returns an HTTP response body string and any error encountered.
func AccountEnquiry(c union.UnionCredentials, queries map[string]string, data []byte, overrideOpts ...req.Option) (string, error) {
	var responseString string

	option := append([]req.Option{
		req.WithMethod("POST"),
		req.WithPath("union/secured/cbaaccountenquiry"),
		req.WithHeader("Sandbox-Key", c.SandboxKey),
		req.WithHeader("Content-Type", "application/json"),
		req.WithQueries(queries),
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

// ChangeUserCredentials sends an HTTP request to union.changeusercredentials sandbox API.
// It returns an HTTP response body string and any error encountered.
func ChangeUserCredentials(c union.UnionCredentials, queries map[string]string, data []byte, overrideOpts ...req.Option) (string, error) {
	var responseString string

	option := append([]req.Option{
		req.WithMethod("POST"),
		req.WithPath("union/secured/changeusercredentials"),
		req.WithHeader("Sandbox-Key", c.SandboxKey),
		req.WithHeader("Content-Type", "application/json"),
		req.WithQueries(queries),
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
