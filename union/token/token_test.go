package token

import (
	"net/http"
	"net/http/httptest"
	"testing"

	req "github.com/enyata/innovation-sandbox-go/request"
	"github.com/enyata/innovation-sandbox-go/union"

	"github.com/icrowley/fake"
	"github.com/stretchr/testify/assert"
)

var (
	uC = union.UnionCredentials{
		SandboxKey: fake.CharactersN(30),
	}
)

type testCase struct {
	name string
	args map[string]string
	want string
	err  error
	f    func(c union.UnionCredentials, body map[string]string, overrideOpts ...req.Option) (string, error)
}

func Run(t *testing.T, tt testCase) {
	t.Run(tt.name, func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(tt.want))
		}))

		defer mockServer.Close()

		got, err := tt.f(uC, tt.args, req.WithBaseURL(mockServer.URL))
		assert.Equal(t, err, tt.err)
		assert.Equal(t, got, tt.want)
	})
}

func TestToken(t *testing.T) {
	payload := map[string]string{
		"client_secret": fake.CharactersN(16),
		"client_id":     fake.CharactersN(16),
		"grant_type":    fake.CharactersN(7),
		"username":      fake.CharactersN(10),
		"password":      fake.CharactersN(20),
	}

	Run(t, testCase{
		name: "should recieve oauth payload",
		args: payload,
		want: `{"message":"OK","data":{"message":"success","scope":"read"}}`,
		err:  nil,
		f:    Token,
	})
}
