package voice

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

func TestVoiceCall(t *testing.T) {
	payload, err := json.Marshal(map[string]string{
		"callFrom": "+234" + fake.CharactersN(10),
		"callTo":   "+234" + fake.CharactersN(10),
	})

	if err != nil {
		t.Fatal(err)
	}

	Run(t, testCase{
		name: "should make voice call",
		args: payload,
		want: `{"status":"Success"}`,
		err:  nil,
		f:    VoiceCall,
	})
}

func TestFetchQueueCalls(t *testing.T) {
	payload, err := json.Marshal(map[string]string{
		"phoneNumbers": "+234" + fake.CharactersN(10),
	})

	if err != nil {
		t.Fatal(err)
	}

	Run(t, testCase{
		name: "should fetch queued calls",
		args: payload,
		want: `{"status":"Success"}`,
		err:  nil,
		f:    FetchQueueCalls,
	})
}

func TestUploadMediaFile(t *testing.T) {
	payload, err := json.Marshal(map[string]string{
		"phoneNumber": "+234" + fake.CharactersN(10),
		"url":         fake.DomainName(),
	})

	if err != nil {
		t.Fatal(err)
	}

	Run(t, testCase{
		name: "should upload media file",
		args: payload,
		want: `{"response":"The request has been fulfilled and resulted in a new resource being created."}`,
		err:  nil,
		f:    UploadMediaFile,
	})
}
