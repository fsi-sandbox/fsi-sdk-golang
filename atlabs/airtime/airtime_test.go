package airtime

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/enyata/innovation-sandbox-go/atlabs"
	req "github.com/enyata/innovation-sandbox-go/request"
	"github.com/icrowley/fake"
	"github.com/stretchr/testify/assert"
)

var (
	aC = atlabs.AtlabsCredentials{
		SandboxKey: fake.CharactersN(30),
	}
)

type testCase struct {
	name string
	args []byte
	want string
	err  error
	f    func(c atlabs.AtlabsCredentials, data []byte, overrideOpts ...req.Option) (string, error)
}

func Run(t *testing.T, tt testCase) {
	t.Run(tt.name, func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(tt.want))
		}))

		defer mockServer.Close()

		got, err := tt.f(aC, tt.args, req.WithBaseURL(mockServer.URL))
		assert.Equal(t, err, tt.err)
		assert.Equal(t, got, tt.want)
	})
}

func TestSendAirtime(t *testing.T) {
	payload, err := json.Marshal(map[string][]map[string]string{
		"recipients": []map[string]string{
			{
				"phoneNumber":  "+234" + fake.CharactersN(10),
				"amount":       fake.CharactersN(1),
				"currencyCode": "NGN",
			},
		},
	})

	if err != nil {
		t.Fatal(err)
	}

	Run(t, testCase{
		name: "should send Airtime amount to phoneNumber",
		args: payload,
		want: `{"errorMessage":"None","numSent":"1"`,
		err:  nil,
		f:    SendAirtime,
	})
}
