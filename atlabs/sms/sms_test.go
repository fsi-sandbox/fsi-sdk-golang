package sms

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

func TestMessage(t *testing.T) {
	payload, err := json.Marshal(map[string]string{
		"to":      "+234" + fake.CharactersN(10),
		"from":    fake.CharactersN(10),
		"message": fake.Sentence(),
	})

	if err != nil {
		t.Fatal(err)
	}

	Run(t, testCase{
		name: "should send SMS message",
		args: payload,
		want: `{"SMSMessageData":{Message:"Sent to 1/1 Total Cost: NGN 0"}}`,
		err:  nil,
		f:    Message,
	})
}

func TestPremiumMessage(t *testing.T) {
	payload, err := json.Marshal(map[string]string{
		"to":                   "+234" + fake.CharactersN(10),
		"from":                 fake.CharactersN(10),
		"message":              fake.Sentence(),
		"keyword":              fake.CharactersN(15),
		"linkId":               fake.CharactersN(30),
		"retryDurationInHours": "1",
	})

	if err != nil {
		t.Fatal(err)
	}

	Run(t, testCase{
		name: "should send SMS premium message",
		args: payload,
		want: `{"SMSMessageData":{Message:"Sent to 1/1"}}`,
		err:  nil,
		f:    PremiumMessage,
	})
}

func TestCreatePremiumSubscription(t *testing.T) {
	payload, err := json.Marshal(map[string]string{
		"shortCode":     fake.CharactersN(10),
		"keyword":       fake.CharactersN(15),
		"phoneNumber":   "+234" + fake.CharactersN(10),
		"checkoutToken": fake.CharactersN(100),
	})

	if err != nil {
		t.Fatal(err)
	}

	Run(t, testCase{
		name: "should create premium subscription",
		args: payload,
		want: `{"status":"Success","description":"Waiting for user input"}`,
		err:  nil,
		f:    CreatePremiumSubscription,
	})
}

func TestDeletePremiumSubscription(t *testing.T) {
	payload, err := json.Marshal(map[string]string{
		"shortCode":   fake.CharactersN(10),
		"keyword":     fake.CharactersN(15),
		"phoneNumber": "+234" + fake.CharactersN(10),
	})

	if err != nil {
		t.Fatal(err)
	}

	Run(t, testCase{
		name: "should delete premium subscription",
		args: payload,
		want: `{"status":"Success","description": "Succeeded"}`,
		err:  nil,
		f:    DeletePremiumSubscription,
	})
}

func TestFetchPremiumSubscription(t *testing.T) {
	payload, err := json.Marshal(map[string]string{
		"shortCode":      fake.CharactersN(10),
		"keyword":        fake.CharactersN(15),
		"lastReceivedId": fake.CharactersN(5),
	})

	if err != nil {
		t.Fatal(err)
	}

	Run(t, testCase{
		name: "should fetch premium subscription",
		args: payload,
		want: `{"responses":[]}`,
		err:  nil,
		f:    FetchPremiumSubscription,
	})
}

func TestFetchMessage(t *testing.T) {
	payload, err := json.Marshal(map[string]string{
		"lastReceivedId": fake.CharactersN(5),
	})

	if err != nil {
		t.Fatal(err)
	}

	Run(t, testCase{
		name: "should fetch messages",
		args: payload,
		want: `{"SMSMessageData":{Messages:[]}}`,
		err:  nil,
		f:    FetchMessage,
	})
}
