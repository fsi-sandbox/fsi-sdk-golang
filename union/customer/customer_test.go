package customer

import (
	"encoding/json"
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
	args []interface{}
	want string
	err  error
	f    func(c union.UnionCredentials, queries map[string]string, data []byte, overrideOpts ...req.Option) (string, error)
}

func Run(t *testing.T, tt testCase) {
	t.Run(tt.name, func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(tt.want))
		}))

		defer mockServer.Close()

		got, err := tt.f(uC, tt.args[0].(map[string]string), tt.args[1].([]byte), req.WithBaseURL(mockServer.URL))
		assert.Equal(t, err, tt.err)
		assert.Equal(t, got, tt.want)
	})
}

func TestCustomerEnquiry(t *testing.T) {
	args := make([]interface{}, 2)
	args[0] = map[string]string{
		"access_token": fake.CharactersN(40),
	}

	payload, err := json.Marshal(map[string]string{
		"accountNumber": fake.CharactersN(10),
		"accountType":   fake.CharactersN(4),
	})

	args[1] = payload
	if err != nil {
		t.Fatal(err)
	}

	Run(t, testCase{
		name: "should enquiry customer details",
		args: args,
		want: `{"message":"OK","data":{"message":"Customer Enquiry Successful","response":"success"}}`,
		err:  nil,
		f:    CustomerEnquiry,
	})
}

func TestCustomerAccountEnquiry(t *testing.T) {
	args := make([]interface{}, 2)
	args[0] = map[string]string{
		"access_token": fake.CharactersN(40),
	}

	payload, err := json.Marshal(map[string]string{
		"accountNumber": fake.CharactersN(10),
		"accountType":   fake.CharactersN(4),
	})

	args[1] = payload
	if err != nil {
		t.Fatal(err)
	}

	Run(t, testCase{
		name: "should enquiry customer account details",
		args: args,
		want: `{"message":"OK","data":{"message":"Account Enquiry Successful","response":"success"}}`,
		err:  nil,
		f:    CustomerAccountEnquiry,
	})
}

func TestAccountEnquiry(t *testing.T) {
	args := make([]interface{}, 2)
	args[0] = map[string]string{
		"access_token": fake.CharactersN(40),
	}

	payload, err := json.Marshal(map[string]string{
		"accountNumber": fake.CharactersN(10),
		"accountType":   fake.CharactersN(4),
	})

	args[1] = payload
	if err != nil {
		t.Fatal(err)
	}

	Run(t, testCase{
		name: "should enquiry account details",
		args: args,
		want: `{"message":"OK","data":{"message":"Enquiry Successful","response":"success"}}`,
		err:  nil,
		f:    AccountEnquiry,
	})
}

func TestChangeUserCredentials(t *testing.T) {
	args := make([]interface{}, 2)
	args[0] = map[string]string{
		"access_token": fake.CharactersN(40),
	}

	payload, err := json.Marshal(map[string]string{
		"username":     fake.CharactersN(6),
		"oldPassword":  fake.CharactersN(14),
		"password":     fake.CharactersN(20),
		"moduleId":     fake.CharactersN(10),
		"clientSecret": fake.CharactersN(5),
	})

	args[1] = payload
	if err != nil {
		t.Fatal(err)
	}

	Run(t, testCase{
		name: "should change user password",
		args: args,
		want: `{"message":"OK","data":{"message":"Password changes successfully"}}`,
		err:  nil,
		f:    ChangeUserCredentials,
	})
}
