package transfer

import (
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
	args map[string]string
	want string
	err  error
	f    func(c sterling.SterlingCredentials, queries map[string]string, overrideOpts ...req.Option) (string, error)
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

func TestInterbankNameEnquiry(t *testing.T) {
	payload := map[string]string{
		"Referenceid":         fake.CharactersN(5),
		"RequestType":         fake.CharactersN(5),
		"Translocation":       fake.CharactersN(5),
		"ToAccount":           fake.CharactersN(10),
		"destinationbankcode": fake.CharactersN(10),
	}

	Run(t, testCase{
		name: "should enquiry account details",
		args: payload,
		want: `{"message":"OK","data":{"message":"success","response":"success"}}`,
		err:  nil,
		f:    InterbankNameEnquiry,
	})
}
