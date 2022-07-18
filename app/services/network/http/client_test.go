package http

import (
	"context"
	"net"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testRequest struct {
	description string
	url         string
	method      string
	headers     map[string]string

	expectedStatusCode int
	expectedData       map[string]string
	expectedError      bool
}

func mockServer() *httptest.Server {
	netListen, _ := net.Listen("tcp", "127.0.0.1:8080")
	ts := httptest.NewUnstartedServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var fakeResponse []byte
		header := r.Header
		if header["Header_key_mock"] != nil {
			fakeResponse = []byte(`{"header_key_mock": "header_value_mock"}`)
		} else {
			fakeResponse = []byte(`{"body_key_mock": "body_value_mock"}`)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(fakeResponse)
	}))
	ts.Listener = netListen
	return ts
}

func TestRequest(t *testing.T) {
	body := map[string]string{"body_key_mock": "body_value_mock"}
	headers := map[string]string{"header_key_mock": "header_value_mock"}
	tests := []testRequest{
		{
			description:        "should return 200 status code, mapping data and nil error when request is made without headers",
			url:                "http://localhost:8080",
			method:             http.MethodGet,
			expectedStatusCode: http.StatusOK,
			expectedData:       body,
			expectedError:      false,
		},
		{
			description:        "should return 200 status code, mapping data and nil error when request is made with headers",
			url:                "http://localhost:8080",
			method:             http.MethodGet,
			expectedStatusCode: http.StatusOK,
			headers:            headers,
			expectedData:       headers,
			expectedError:      false,
		},
	}
	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			var actualData map[string]string
			request := Request{
				URL:     test.url,
				Method:  test.method,
				Headers: test.headers,
			}
			server := mockServer()
			server.Start()
			actualResponse, actualErr := request.Do(context.Background(), &actualData)
			assert.Equal(t, test.expectedStatusCode, actualResponse.StatusCode)
			assert.Equal(t, test.expectedError, actualErr != nil)
			assert.Equal(t, test.expectedData, actualData)
			defer server.Close()
		})
	}
}
