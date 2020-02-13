// Package bvnr implements access to nibss.BVNPlaceHolder sandbox.
package placeholder

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/enyata/innovation-sandbox-go/nibss"
	req "github.com/enyata/innovation-sandbox-go/nibss/request"
)

// Reset sends an HTTP request to nibss.BVNPlaceHolder.Reset sandbox API.
// It returns nibss.ResetCredentials and any error encountered.
func Reset(c nibss.NibssCredentials, overrideOpts ...req.Option) (nibss.ResetCredentials, error) {
	var resetCredentials nibss.ResetCredentials

	option := append([]req.Option{
		req.WithMethod("POST"),
		req.WithPath("nibss/BVNPlaceHolder/Reset"),
		req.WithHeader("Sandbox-Key", c.SandboxKey),
		req.WithHeader("OrganisationCode", nibss.Encode(c.OrganisationCode)),
		req.WithBody([]byte("")),
	}, overrideOpts...)
	req, err := req.New(option...)

	if err != nil {
		return resetCredentials, err
	}

	_, err = req.Make(func(response *http.Response) error {
		defer response.Body.Close()

		requiredCredentials := map[string]string{
			"Aes_key":      "",
			"Code":         "",
			"Email":        "",
			"Ivkey":        "",
			"Name":         "",
			"Password":     "",
			"Responsecode": "",
		}

		for paramName := range requiredCredentials {
			requiredCredentials[paramName] = response.Header.Get(paramName)
		}

		headers, err := json.Marshal(requiredCredentials)

		if err != nil {
			return err
		}

		return json.Unmarshal(headers, &resetCredentials)
	})

	return resetCredentials, err
}

// ValidateRecord sends an HTTP request to nibss.BVNPlaceHolder.ValidateRecord sandbox API.
// It returns a decrypted HTTP response body string and any error encountered.
func ValidateRecord(c nibss.NibssCredentials, cr nibss.Crypt, data []byte, overrideOpts ...req.Option) (string, error) {
	var responseString string

	encryptedData, err := cr.Encrypt(data)

	if err != nil {
		return responseString, err
	}

	authorization := nibss.Encode(fmt.Sprintf("%s:%s", cr.Code, cr.Password))
	today := time.Now().Format("20060102") /*YYYYMMDD format*/
	signature := nibss.Sha256(fmt.Sprintf("%s%s%s", cr.Code, today, cr.Password))
	option := append([]req.Option{
		req.WithMethod("POST"),
		req.WithPath("nibss/BVNPlaceHolder/ValidateRecord"),
		req.WithDefaultHeaders(),
		req.WithHeader("Sandbox-Key", c.SandboxKey),
		req.WithHeader("OrganisationCode", nibss.Encode(c.OrganisationCode)),
		req.WithHeader("Authorization", authorization),
		req.WithHeader("SIGNATURE", signature),
		req.WithBody([]byte(hex.EncodeToString(encryptedData))),
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

		decryptedData, err := cr.Decrypt(string(data))
		responseString = string(decryptedData)
		return err
	})

	return responseString, err
}

// ValidateRecords sends an HTTP request to nibss.BVNPlaceHolder.ValidateRecords sandbox API.
// It returns a decrypted HTTP response body string and any error encountered.
func ValidateRecords(c nibss.NibssCredentials, cr nibss.Crypt, data []byte, overrideOpts ...req.Option) (string, error) {
	var responseString string

	encryptedData, err := cr.Encrypt(data)

	if err != nil {
		return responseString, err
	}

	authorization := nibss.Encode(fmt.Sprintf("%s:%s", cr.Code, cr.Password))
	today := time.Now().Format("20060102") /*YYYYMMDD format*/
	signature := nibss.Sha256(fmt.Sprintf("%s%s%s", cr.Code, today, cr.Password))
	option := append([]req.Option{
		req.WithMethod("POST"),
		req.WithPath("nibss/BVNPlaceHolder/ValidateRecords"),
		req.WithDefaultHeaders(),
		req.WithHeader("Sandbox-Key", c.SandboxKey),
		req.WithHeader("OrganisationCode", nibss.Encode(c.OrganisationCode)),
		req.WithHeader("Authorization", authorization),
		req.WithHeader("SIGNATURE", signature),
		req.WithBody([]byte(hex.EncodeToString(encryptedData))),
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

		decryptedData, err := cr.Decrypt(string(data))
		responseString = string(decryptedData)
		return err
	})

	return responseString, err
}
