package utilities

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

func TestProviders(t *testing.T) {
	payload := map[string]string{
		"state":   fake.State(),
		"plan_id": fake.CharactersN(30),
		"tiers":   fake.Word(),
		"page":    fake.DigitsN(3),
		"limit":   fake.DigitsN(3),
	}

	Run(t, testCase{
		name: "should retrieve providers",
		want: `{"message":"OK","data":{"status":"success"}}`,
		err:  nil,
		f: func(queries map[string]string) overrideFunc {
			return func(overrideOpts req.Option) (string, error) {
				return Providers(rC, queries, overrideOpts)
			}
		}(payload),
	})
}

func TestStates(t *testing.T) {
	Run(t, testCase{
		name: "should retrieve states",
		want: `{"message":"OK","data":{"status":"success"}}`,
		err:  nil,
		f: func() overrideFunc {
			return func(overrideOpts req.Option) (string, error) {
				return States(rC, overrideOpts)
			}
		}(),
	})
}

func TestBenefits(t *testing.T) {
	payload := map[string]string{
		"plan": fake.Word(),
	}

	Run(t, testCase{
		name: "should retrieve benefits",
		want: `{"message":"OK","data":{"status":"success"}}`,
		err:  nil,
		f: func(queries map[string]string) overrideFunc {
			return func(overrideOpts req.Option) (string, error) {
				return Benefits(rC, queries, overrideOpts)
			}
		}(payload),
	})
}

func TestTitles(t *testing.T) {
	Run(t, testCase{
		name: "should retrieve titles",
		want: `{"message":"OK","data":{"status":"success"}}`,
		err:  nil,
		f: func() overrideFunc {
			return func(overrideOpts req.Option) (string, error) {
				return Titles(rC, overrideOpts)
			}
		}(),
	})
}

func TestOccupations(t *testing.T) {
	Run(t, testCase{
		name: "should retrieve occupations",
		want: `{"message":"OK","data":{"status":"success"}}`,
		err:  nil,
		f: func() overrideFunc {
			return func(overrideOpts req.Option) (string, error) {
				return Occupations(rC, overrideOpts)
			}
		}(),
	})
}

func TestMaritalStatuses(t *testing.T) {
	Run(t, testCase{
		name: "should retrieve marital statuses",
		want: `{"message":"OK","data":{"status":"success"}}`,
		err:  nil,
		f: func() overrideFunc {
			return func(overrideOpts req.Option) (string, error) {
				return MaritalStatuses(rC, overrideOpts)
			}
		}(),
	})
}
