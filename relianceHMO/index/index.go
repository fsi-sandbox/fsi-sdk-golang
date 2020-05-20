package index

import (
	"io/ioutil"
	"net/http"

	"github.com/enyata/innovation-sandbox-go/relianceHMO"
	req "github.com/enyata/innovation-sandbox-go/request"
)

func Plans(c relianceHMO.RelianceHMOCredentials, queries map[string]string, overrideOpts ...req.Option) (string, error) {
	var responseString string

	option := append([]req.Option{
		req.WithMethod("GET"),
		req.WithPath("relianceHMO/plans"),
		req.WithHeader("Sandbox-Key", c.SandboxKey),
		req.WithHeader("Content-Type", "application/json"),
		req.WithQueries(queries),
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

func Enrollees(c relianceHMO.RelianceHMOCredentials, queries map[string]string, overrideOpts ...req.Option) (string, error) {
	var responseString string

	option := append([]req.Option{
		req.WithMethod("GET"),
		req.WithPath("relianceHMO/enrollees"),
		req.WithHeader("Sandbox-Key", c.SandboxKey),
		req.WithHeader("Content-Type", "application/json"),
		req.WithQueries(queries),
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

func Register(c relianceHMO.RelianceHMOCredentials, data []byte, overrideOpts ...req.Option) (string, error) {
	var responseString string

	option := append([]req.Option{
		req.WithMethod("POST"),
		req.WithPath("relianceHMO/enrollees"),
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

func Wallet(c relianceHMO.RelianceHMOCredentials, overrideOpts ...req.Option) (string, error) {
	var responseString string

	option := append([]req.Option{
		req.WithMethod("GET"),
		req.WithPath("relianceHMO/wallet"),
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

func Consultations(c relianceHMO.RelianceHMOCredentials, queries map[string]string, overrideOpts ...req.Option) (string, error) {
	var responseString string

	option := append([]req.Option{
		req.WithMethod("POST"),
		req.WithPath("relianceHMO/consultations"),
		req.WithHeader("Sandbox-Key", c.SandboxKey),
		req.WithHeader("Content-Type", "application/json"),
		req.WithQueries(queries),
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

func Upload(c relianceHMO.RelianceHMOCredentials, queries map[string]string, overrideOpts ...req.Option) (string, error) {
	var responseString string

	option := append([]req.Option{
		req.WithMethod("POST"),
		req.WithPath("relianceHMO/upload"),
		req.WithHeader("Sandbox-Key", c.SandboxKey),
		req.WithHeader("Content-Type", "application/json"),
		req.WithQueries(queries),
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
