package bvnr

import (
	"encoding/hex"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/enyata/innovation-sandbox-go/nibss"
	req "github.com/enyata/innovation-sandbox-go/request"

	"github.com/icrowley/fake"
	"github.com/stretchr/testify/assert"
)

var (
	key  = fake.CharactersN(30)
	code = fake.CharactersN(5)
	nC   = nibss.NibssCredentials{
		SandboxKey:       key,
		OrganisationCode: code,
	}
	rC = nibss.ResetCredentials{
		AESKey:   fake.CharactersN(16),
		Code:     fake.CharactersN(5),
		Email:    fake.EmailAddress(),
		IVKey:    fake.CharactersN(16),
		Name:     fake.FullName(),
		Password: fake.SimplePassword(),
	}
)

type testCase struct {
	name string
	args []byte
	want string
	f    func(c nibss.NibssCredentials, cr nibss.Crypt, data []byte, overrideOpts ...req.Option) (string, error)
}

func Run(t *testing.T, tt testCase) {
	crypt, _ := setup(t, rC)

	t.Run(tt.name, func(t *testing.T) {
		enc, _ := crypt.Encrypt([]byte(tt.want))

		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(hex.EncodeToString(enc)))
		}))

		defer mockServer.Close()

		got, err := tt.f(nC, crypt, tt.args, req.WithBaseURL(mockServer.URL))
		assert.Nil(t, err)
		assert.Equal(t, got, tt.want)
	})
}

func setup(t *testing.T, want nibss.ResetCredentials) (nibss.Crypt, nibss.ResetCredentials) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Aes_key", want.AESKey)
		w.Header().Set("Code", want.Code)
		w.Header().Set("Email", want.Email)
		w.Header().Set("Ivkey", want.IVKey)
		w.Header().Set("Name", want.Name)
		w.Header().Set("Password", want.Password)
		w.Header().Set("Responsecode", "000")
		w.Write([]byte(""))
	}))

	defer mockServer.Close()

	resetCredentials, err := Reset(nC, req.WithBaseURL(mockServer.URL))

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
	want := nibss.ResetCredentials{
		AESKey:   fake.CharactersN(16),
		Code:     fake.CharactersN(5),
		Email:    fake.EmailAddress(),
		IVKey:    fake.CharactersN(16),
		Name:     fake.FullName(),
		Password: fake.SimplePassword(),
	}

	_, got := setup(t, want)
	assert.Equal(t, got, want)
}

func TestVerifySingleBVN(t *testing.T) {
	payload, err := json.Marshal(map[string]string{
		"BVN": fake.CharactersN(10),
	})

	if err != nil {
		t.Fatal(err)
	}

	tt := testCase{
		name: "should verify single BVN",
		args: payload,
		want: `{"message":"OK","data":{"ResponseCode":"00"}}`,
		f:    VerifySingleBVN,
	}

	Run(t, tt)
}

func TestVerifyMultipleBVN(t *testing.T) {
	payload, err := json.Marshal(map[string]string{
		"BVNS": fake.CharactersN(40),
	})

	if err != nil {
		t.Fatal(err)
	}

	tt := testCase{
		name: "should verify multiple BVN",
		args: payload,
		want: `{"message":"OK","data":{"ResponseCode":"00"}}`,
		f:    VerifyMultipleBVN,
	}

	Run(t, tt)
}

func TestGetSingleBVN(t *testing.T) {
	payload, err := json.Marshal(map[string]string{
		"BVN": fake.CharactersN(10),
	})

	if err != nil {
		t.Fatal(err)
	}

	tt := testCase{
		name: "should get single BVN",
		args: payload,
		want: `{"message":"OK","data":{"ResponseCode":"00"}}`,
		f:    GetSingleBVN,
	}

	Run(t, tt)
}

func TestGetMultipleBVN(t *testing.T) {
	payload, err := json.Marshal(map[string]string{
		"BVNS": fake.CharactersN(40),
	})

	if err != nil {
		t.Fatal(err)
	}

	tt := testCase{
		name: "should get multiple BVN",
		args: payload,
		want: `{"message":"OK","data":{"ResponseCode":"00"}}`,
		f:    GetMultipleBVN,
	}

	Run(t, tt)
}

func TestIsBVNWatchlisted(t *testing.T) {
	payload, err := json.Marshal(map[string]string{
		"BVN": fake.CharactersN(10),
	})

	if err != nil {
		t.Fatal(err)
	}

	tt := testCase{
		name: "should check if BVN is WatchListed",
		args: payload,
		want: `{"message":"OK","data":{"ResponseCode":"00"}}`,
		f:    IsBVNWatchlisted,
	}

	Run(t, tt)
}
