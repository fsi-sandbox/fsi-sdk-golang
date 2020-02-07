package bvnr

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

type NibssCredentials struct {
	SandboxKey       string
	OrganisationCode string
}

type ResetCredentials struct {
	AESKey   string `json:"Aes_key"`
	Code     string `json:"Code"`
	Email    string `json:"Email"`
	IVKey    string `json:"Ivkey"`
	Name     string `json:"Name"`
	Password string `json:"Password"`
}

func Reset(c NibssCredentials) (ResetCredentials, error) {
	var resetCredentials ResetCredentials

	req, err := req.New(
		req.WithMethod("POST"),
		req.WithPath("nibss/bvnr/Reset"),
		req.WithHeader("Sandbox-Key", c.SandboxKey),
		req.WithHeader("OrganisationCode", nibss.Encode(c.OrganisationCode)),
		req.WithBody([]byte("")),
	)

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

func VerifySingleBVN(c NibssCredentials, cr nibss.Crypt, data []byte) (string, error) {
	var responseString string

	encryptedData, err := cr.Encrypt(data)

	if err != nil {
		return responseString, err
	}

	authorization := nibss.Encode(fmt.Sprintf("%s:%s", cr.Code, cr.Password))
	today := time.Now().Format("20060102") /*YYYYMMDD format*/
	signature := nibss.Sha256(fmt.Sprintf("%s%s%s", cr.Code, today, cr.Password))

	req, err := req.New(
		req.WithMethod("POST"),
		req.WithPath("nibss/bvnr/VerifySingleBVN"),
		req.WithDefaultHeaders(),
		req.WithHeader("Sandbox-Key", c.SandboxKey),
		req.WithHeader("OrganisationCode", nibss.Encode(c.OrganisationCode)),
		req.WithHeader("Authorization", authorization),
		req.WithHeader("SIGNATURE", signature),
		req.WithBody([]byte(hex.EncodeToString(encryptedData))),
	)

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

func VerifyMultipleBVN(c NibssCredentials, cr nibss.Crypt, data []byte) (string, error) {
	var responseString string

	encryptedData, err := cr.Encrypt(data)

	if err != nil {
		return responseString, err
	}

	authorization := nibss.Encode(fmt.Sprintf("%s:%s", cr.Code, cr.Password))
	today := time.Now().Format("20060102") /*YYYYMMDD format*/
	signature := nibss.Sha256(fmt.Sprintf("%s%s%s", cr.Code, today, cr.Password))

	req, err := req.New(
		req.WithMethod("POST"),
		req.WithPath("nibss/bvnr/VerifyMultipleBVN"),
		req.WithDefaultHeaders(),
		req.WithHeader("Sandbox-Key", c.SandboxKey),
		req.WithHeader("OrganisationCode", nibss.Encode(c.OrganisationCode)),
		req.WithHeader("Authorization", authorization),
		req.WithHeader("SIGNATURE", signature),
		req.WithBody([]byte(hex.EncodeToString(encryptedData))),
	)

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

func GetSingleBVN(c NibssCredentials, cr nibss.Crypt, data []byte) (string, error) {
	var responseString string

	encryptedData, err := cr.Encrypt(data)

	if err != nil {
		return responseString, err
	}

	authorization := nibss.Encode(fmt.Sprintf("%s:%s", cr.Code, cr.Password))
	today := time.Now().Format("20060102") /*YYYYMMDD format*/
	signature := nibss.Sha256(fmt.Sprintf("%s%s%s", cr.Code, today, cr.Password))

	req, err := req.New(
		req.WithMethod("POST"),
		req.WithPath("nibss/bvnr/GetSingleBVN"),
		req.WithDefaultHeaders(),
		req.WithHeader("Sandbox-Key", c.SandboxKey),
		req.WithHeader("OrganisationCode", nibss.Encode(c.OrganisationCode)),
		req.WithHeader("Authorization", authorization),
		req.WithHeader("SIGNATURE", signature),
		req.WithBody([]byte(hex.EncodeToString(encryptedData))),
	)

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

func GetMultipleBVN(c NibssCredentials, cr nibss.Crypt, data []byte) (string, error) {
	var responseString string

	encryptedData, err := cr.Encrypt(data)

	if err != nil {
		return responseString, err
	}

	authorization := nibss.Encode(fmt.Sprintf("%s:%s", cr.Code, cr.Password))
	today := time.Now().Format("20060102") /*YYYYMMDD format*/
	signature := nibss.Sha256(fmt.Sprintf("%s%s%s", cr.Code, today, cr.Password))

	req, err := req.New(
		req.WithMethod("POST"),
		req.WithPath("nibss/bvnr/GetMultipleBVN"),
		req.WithDefaultHeaders(),
		req.WithHeader("Sandbox-Key", c.SandboxKey),
		req.WithHeader("OrganisationCode", nibss.Encode(c.OrganisationCode)),
		req.WithHeader("Authorization", authorization),
		req.WithHeader("SIGNATURE", signature),
		req.WithBody([]byte(hex.EncodeToString(encryptedData))),
	)

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

func IsBVNWatchlisted(c NibssCredentials, cr nibss.Crypt, data []byte) (string, error) {
	var responseString string

	encryptedData, err := cr.Encrypt(data)

	if err != nil {
		return responseString, err
	}

	authorization := nibss.Encode(fmt.Sprintf("%s:%s", cr.Code, cr.Password))
	today := time.Now().Format("20060102") /*YYYYMMDD format*/
	signature := nibss.Sha256(fmt.Sprintf("%s%s%s", cr.Code, today, cr.Password))

	req, err := req.New(
		req.WithMethod("POST"),
		req.WithPath("nibss/bvnr/IsBVNWatchlisted"),
		req.WithDefaultHeaders(),
		req.WithHeader("Sandbox-Key", c.SandboxKey),
		req.WithHeader("OrganisationCode", nibss.Encode(c.OrganisationCode)),
		req.WithHeader("Authorization", authorization),
		req.WithHeader("SIGNATURE", signature),
		req.WithBody([]byte(hex.EncodeToString(encryptedData))),
	)

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