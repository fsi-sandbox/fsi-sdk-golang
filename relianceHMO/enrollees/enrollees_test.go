package enrollees

import (
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

func TestEnrollee(t *testing.T) {
	APIPath := fake.CharactersN(30)
	Run(t, testCase{
		name: "should return enrollees",
		want: `{"message":"OK","data":{"status":"success"}}`,
		err:  nil,
		f: func(path string) overrideFunc {
			return func(overrideOpts req.Option) (string, error) {
				return Enrollee(rC, path, overrideOpts)
			}
		}(APIPath),
	})
}

func TestProfile(t *testing.T) {
	payload := map[string]string{
		"sex":                      fake.Gender(),
		"date_of_birth":            fake.CharactersN(30),
		"home_address":             fake.StreetAddress(),
		"has_smartphone":           "true",
		"profile_picture_filename": fake.CharactersN(20),
		"hash":                     fake.CharactersN(30),
	}

	Run(t, testCase{
		name: "should return enrollees profile data",
		want: `{"message":"OK","data":{"status":"success"}}`,
		err:  nil,
		f: func(queries map[string]string) overrideFunc {
			return func(overrideOpts req.Option) (string, error) {
				return Profile(rC, queries, overrideOpts)
			}
		}(payload),
	})
}

func TestValidate(t *testing.T) {
	payload := map[string]string{
		"hmo_id": fake.CharactersN(30),
	}

	Run(t, testCase{
		name: "should validate enrollee",
		want: `{"message":"OK","data":{"status":"success"}}`,
		err:  nil,
		f: func(queries map[string]string) overrideFunc {
			return func(overrideOpts req.Option) (string, error) {
				return Validate(rC, queries, overrideOpts)
			}
		}(payload),
	})
}

func TestIDCard(t *testing.T) {
	payload := map[string]string{
		"hmo_id": fake.CharactersN(30),
	}

	Run(t, testCase{
		name: "should return enrollee ID Card",
		want: `{"message":"OK","data":{"status":"success"}}`,
		err:  nil,
		f: func(queries map[string]string) overrideFunc {
			return func(overrideOpts req.Option) (string, error) {
				return IDCard(rC, queries, overrideOpts)
			}
		}(payload),
	})
}
