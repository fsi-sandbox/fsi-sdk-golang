package account

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	req "github.com/enyata/innovation-sandbox-go/request"
	"github.com/enyata/innovation-sandbox-go/sterling"

	"github.com/icrowley/fake"
	"github.com/stretchr/testify/assert"
)

var (
	sC = sterling.SterlingCredentials{
		SandboxKey:      fake.CharactersN(30),
		SubscriptionKey: fake.CharactersN(10),
		Appid:           fake.CharactersN(3),
		Ipval:           fake.CharactersN(2),
	}
)

type testCase struct {
	name string
	args []byte
	want string
	err  error
	f    func(c sterling.SterlingCredentials, data []byte, overrideOpts ...req.Option) (string, error)
}

func Run(t *testing.T, tt testCase) {
	t.Run(tt.name, func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(tt.want))
		}))

		defer mockServer.Close()

		got, err := tt.f(sC, tt.args, req.WithBaseURL(mockServer.URL))
		assert.Equal(t, err, tt.err)
		assert.Equal(t, got, tt.want)
	})
}

func TestInterbankTransferReq(t *testing.T) {
	payload, err := json.Marshal(map[string]string{
		"Referenceid":           fake.CharactersN(10),
		"RequestType":           fake.CharactersN(10),
		"Translocation":         fake.CharactersN(10),
		"SessionID":             fake.CharactersN(5),
		"FromAccount":           fake.CharactersN(5),
		"ToAccount":             fake.CharactersN(5),
		"Amount":                fake.CharactersN(5),
		"DestinationBankCode":   fake.CharactersN(5),
		"NEResponse":            fake.CharactersN(5),
		"BenefiName":            fake.CharactersN(5),
		"PaymentReference":      fake.CharactersN(5),
		"OriginatorAccountName": fake.CharactersN(5),
		"translocation":         fake.CharactersN(5),
	})

	if err != nil {
		t.Fatal(err)
	}

	Run(t, testCase{
		name: "should submit transaction for processing",
		args: payload,
		want: `{"message":"OK","data":{"message":"success","response":"success"}}`,
		err:  nil,
		f:    InterbankTransferReq,
	})
}
