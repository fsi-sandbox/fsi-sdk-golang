package clients

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

func TestSignup(t *testing.T) {
	payload, err := json.Marshal([]map[string]interface{}{
		{
			"transfer_code":     fake.CharactersN(30),
			"company_name":      fake.Company(),
			"company_address":   fake.StreetAddress(),
			"state_code":        fake.StateAbbrev(),
			"payment_frequency": fake.Word(),
			"company_admin": map[string]string{
				"first_name":    fake.FirstName(),
				"last_name":     fake.LastName(),
				"email_address": fake.EmailAddress(),
				"phone_number":  fake.Phone(),
			},
			"enrollees": []map[string]interface{}{
				{
					"plan_id":       fake.DigitsN(3),
					"first_name":    fake.FirstName(),
					"last_name":     fake.LastName(),
					"email_address": fake.EmailAddress(),
					"phone_number":  fake.Phone(),
				},
				{
					"plan_id":       fake.DigitsN(3),
					"first_name":    fake.FirstName(),
					"last_name":     fake.LastName(),
					"email_address": fake.EmailAddress(),
					"phone_number":  fake.Phone(),
				},
			},
		},
	})

	if err != nil {
		t.Fatal(err)
	}

	Run(t, testCase{
		name: "should add clients",
		want: `{"message":"OK","data":{"status":"success"}}`,
		err:  nil,
		f: func(data []byte) overrideFunc {
			return func(overrideOpts req.Option) (string, error) {
				return Signup(rC, data, overrideOpts)
			}
		}(payload),
	})
}

func TestRenew(t *testing.T) {
	APIPath := fake.CharactersN(30)
	payload, err := json.Marshal([]map[string]interface{}{
		{
			"transfer_code": fake.CharactersN(30),
			"add": []map[string]interface{}{
				{
					"plan_id":       fake.DigitsN(3),
					"first_name":    fake.FirstName(),
					"last_name":     fake.LastName(),
					"email_address": fake.EmailAddress(),
					"phone_number":  fake.Phone(),
				},
			},
			"remove": []string{fake.CharactersN(30), fake.CharactersN(30)},
			"update": []map[string]interface{}{
				{
					"plan_id":    fake.DigitsN(3),
					"user_token": fake.CharactersN(30),
				},
			},
		},
	})

	if err != nil {
		t.Fatal(err)
	}

	Run(t, testCase{
		name: "should retrieve renewed data",
		want: `{"message":"OK","data":{"status":"success"}}`,
		err:  nil,
		f: func(path string, data []byte) overrideFunc {
			return func(overrideOpts req.Option) (string, error) {
				return Renew(rC, path, data, overrideOpts)
			}
		}(APIPath, payload),
	})
}
