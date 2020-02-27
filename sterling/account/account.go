package account

import (
	"io/ioutil"
	"net/http"

	req "github.com/enyata/innovation-sandbox-go/request"
	"github.com/enyata/innovation-sandbox-go/sterling"
)

func InterbankTransferReq(c sterling.SterlingCredentials, data []byte, overrideOpts ...req.Option) (string, error) {
	var responseString string

	option := append([]req.Option{
		req.WithMethod("POST"),
		req.WithPath("sterling/accountapi/api/Spay/InterbankTransferReq"),
		req.WithHeader("Sandbox-Key", c.SandboxKey),
		req.WithHeader("Ocp-Apim-Subscription-Key", c.SubscriptionKey),
		req.WithHeader("Ocp-Apim-Trace", "true"),
		req.WithHeader("Appid", c.Appid),
		req.WithHeader("Content-Type", "application/json"),
		req.WithHeader("ipval", c.Ipval),
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
