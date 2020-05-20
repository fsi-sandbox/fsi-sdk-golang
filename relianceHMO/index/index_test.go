package index

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

func TestPlans(t *testing.T) {
	payload := map[string]string{
		"type":    fake.CharactersN(30),
		"package": fake.Words(),
	}

	Run(t, testCase{
		name: "should return plans",
		want: `{"message":"OK","data":{"status":"success"}}`,
		err:  nil,
		f: func(queries map[string]string) overrideFunc {
			return func(overrideOpts req.Option) (string, error) {
				return Plans(rC, queries, overrideOpts)
			}
		}(payload),
	})
}

func TestEnrollees(t *testing.T) {
	payload := map[string]string{
		"page":  fake.DigitsN(10),
		"limit": fake.DigitsN(20),
	}

	Run(t, testCase{
		name: "should return enrollees data",
		want: `{"message":"OK","data":{"status":"success"}}`,
		err:  nil,
		f: func(queries map[string]string) overrideFunc {
			return func(overrideOpts req.Option) (string, error) {
				return Enrollees(rC, queries, overrideOpts)
			}
		}(payload),
	})
}

func TestRegister(t *testing.T) {
	payload, err := json.Marshal([]map[string]interface{}{
		{
			"enrollees": []map[string]interface{}{
				{
					"payment_frequency":    fake.Word(),
					"first_name":           fake.FirstName(),
					"last_name":            fake.LastName(),
					"email_address":        fake.EmailAddress(),
					"phone_number":         fake.Phone(),
					"plan_id":              fake.DigitsN(3),
					"can_complete_profile": "true",
					"profile": map[string]string{
						"sex":                      fake.Gender(),
						"date_of_birth":            fake.CharactersN(30),
						"first_name":               fake.FirstName(),
						"last_name":                fake.LastName(),
						"primary_phone_number":     fake.Phone(),
						"home_address":             fake.StreetAddress(),
						"has_smartphone":           "true",
						"profile_picture_filename": fake.CharactersN(20),
						"enrollee_type":            fake.CharactersN(10),
						"hmo_id":                   fake.CharactersN(30),
					},
					"dependants": []map[string]string{
						{
							"first_name":    fake.FirstName(),
							"last_name":     fake.LastName(),
							"email_address": fake.EmailAddress(),
							"phone_number":  fake.Phone(),
							"plan_id":       fake.DigitsN(3),
						},
					},
				},
			},
		},
	})

	if err != nil {
		t.Fatal(err)
	}

	Run(t, testCase{
		name: "should create enrollee",
		want: `{"message":"OK","data":{"status":"success"}}`,
		err:  nil,
		f: func(data []byte) overrideFunc {
			return func(overrideOpts req.Option) (string, error) {
				return Register(rC, data, overrideOpts)
			}
		}(payload),
	})
}

func TestWallet(t *testing.T) {
	Run(t, testCase{
		name: "should return wallet data",
		want: `{"message":"OK","data":{"status":"success"}}`,
		err:  nil,
		f: func() overrideFunc {
			return func(overrideOpts req.Option) (string, error) {
				return Wallet(rC, overrideOpts)
			}
		}(),
	})
}

func TestConsultations(t *testing.T) {
	payload := map[string]string{
		"patient_id": fake.CharactersN(30),
		"reason":     fake.Sentence(),
	}

	Run(t, testCase{
		name: "should return consultation data",
		want: `{"message":"OK","data":{"status":"success"}}`,
		err:  nil,
		f: func(queries map[string]string) overrideFunc {
			return func(overrideOpts req.Option) (string, error) {
				return Consultations(rC, queries, overrideOpts)
			}
		}(payload),
	})
}

func TestUpload(t *testing.T) {
	payload := map[string]string{
		"file_use": fake.CharactersN(10),
		"file":     fake.CharactersN(30),
	}

	Run(t, testCase{
		name: "should return uploaded",
		want: `{"message":"OK","data":{"status":"success"}}`,
		err:  nil,
		f: func(queries map[string]string) overrideFunc {
			return func(overrideOpts req.Option) (string, error) {
				return Upload(rC, queries, overrideOpts)
			}
		}(payload),
	})
}
