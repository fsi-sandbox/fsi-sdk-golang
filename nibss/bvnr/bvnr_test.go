package bvnr

import (
	"encoding/json"
	"testing"

	"github.com/enyata/innovation-sandbox-go/nibss"

	"github.com/stretchr/testify/assert"
)

var (
	key = "0ae0db703c04119b3db7a03d7f854c13"
	// key  = "49264b2cc8fd68b33326c6d5468e5290"
	code             = "11111"
	nibssCredentials = NibssCredentials{
		SandboxKey:       key,
		OrganisationCode: code,
	}
)

func setup(t *testing.T) (nibss.Crypt, ResetCredentials) {
	resetCredentials, err := Reset(nibssCredentials)

	if err != nil {
		t.Fatal(err)
	}

	crypt := nibss.Crypt{
		AESKey:   []byte(resetCredentials.AESKey),
		IVKey:    []byte(resetCredentials.IVKey),
		Code:     resetCredentials.Code,
		Password: resetCredentials.Password,
	}

	return crypt, resetCredentials
}

func TestReset(t *testing.T) {
	credentials, err := Reset(nibssCredentials)

	assert.Nil(t, err)
	assert.NotNil(t, credentials)
}

// func TestVerifySingleBVN(t *testing.T) {

// }

func TestVerifyMultipleBVN(t *testing.T) {
	crypt, _ := setup(t)
	jsonValue, err := json.Marshal(map[string]string{
		"BVNS": "12345678901, 12345678902, 12345678903",
	})

	data, err := VerifyMultipleBVN(nibssCredentials, crypt, jsonValue)

	if err != nil {
		t.Fatal(err)
	}

	t.Log(string(jsonValue), data)
	assert.True(t, false)
}

// func TestGetSingleBVN(t *testing.T) {

// }

// func TestGetMultipleBVN(t *testing.T) {

// }

// func TestIsBVNWatchlisted(t *testing.T) {

// }
