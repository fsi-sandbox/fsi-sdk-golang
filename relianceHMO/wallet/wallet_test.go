package wallet

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/enyata/innovation-sandbox-go/relianceHMO"
	req "github.com/enyata/innovation-sandbox-go/request"

	"github.com/icrowley/fake"
	"github.com/stretchr/testify/assert"
)

var (
	rC = relianceHMO.RelianceHMOCredentials{
		SandboxKey: fake.CharactersN(30),
	}
)

type overrideFunc func(overrideOpts req.Option) (string, error)

type testCase struct {
	name string
	want string
	err  error
	f    overrideFunc
}

func Run(t *testing.T, tt testCase) {
	t.Run(tt.name, func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(tt.want))
		}))

		defer mockServer.Close()

		got, err := tt.f(req.WithBaseURL(mockServer.URL))
		assert.Equal(t, err, tt.err)
		assert.Equal(t, got, tt.want)
	})
}

func TestFund(t *testing.T) {
	payload, err := json.Marshal(map[string]string{
		"amount": fake.CharactersN(5),
	})

	if err != nil {
		t.Fatal(err)
	}

	Run(t, testCase{
		name: "should fund wallet",
		want: `{"message":"OK","data":{"status":"success"}}`,
		err:  nil,
		f: func(data []byte) overrideFunc {
			return func(overrideOpts req.Option) (string, error) {
				return Fund(rC, data, overrideOpts)
			}
		}(payload),
	})
}

func TestTransactions(t *testing.T) {
	Run(t, testCase{
		name: "should retrieve wallet transactions",
		want: `{"message":"OK","data":{"status":"success"}}`,
		err:  nil,
		f: func() overrideFunc {
			return func(overrideOpts req.Option) (string, error) {
				return Transactions(rC, overrideOpts)
			}
		}(),
	})
}
