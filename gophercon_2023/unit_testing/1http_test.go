package unit_testing

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

//Testing using httptest.NewServer
func TestAPI_DoStuff_NewServer(t *testing.T) {
	var mockedServer *httptest.Server
	mockClient := &http.Client{}
	type fields struct {
		Client  *http.Client
		baseURL string
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
		setup   func() //not auto generated //Test Fixtures
	}{
		{
			name: "Testing HTTP Client",
			fields: fields{
				Client: mockClient,
			},
			setup: func() {
				// Start a local HTTP server
				mockedServer = httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
					// Test request parameters
					stringCompare(t, req.URL.String(), "/some/path")
					// Send response to be tested
					rw.Write([]byte(`OK`)) //If the method returned body then this can also be tested
				}))
				mockClient = mockedServer.Client()
			},
			want: []byte(`OK`),
		},
		{
			name: "Testing wrong path",
			fields: fields{
				Client: mockClient,
			},
			setup: func() {
				// Start a local HTTP server
				mockedServer = httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
					// Test request parameters
					stringCompare(t, req.URL.String(), "/some/path/wrong")
					// Send response to be tested
					rw.Write([]byte(`OK`)) //If the method returned body then this can also be tested
				}))
				// Close the server when test finishes
				defer mockedServer.Close()
				mockClient = mockedServer.Client()
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt.setup() //not auto generated
		// Close the server when test finishes
		defer mockedServer.Close()
		t.Run(tt.name, func(t *testing.T) {
			api := &API{
				Client:  tt.fields.Client,
				baseURL: mockedServer.URL, //not auto generated
			}
			_, got, err := api.DoStuff()
			if (err != nil) != tt.wantErr {
				t.Errorf("API.DoStuff() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("API.DoStuff() = %v, want %v", string(got), string(tt.want))
			}
		})
	}

}

//Testing by Replacing http.Transport

/*
Transport specifies the mechanism by which individual HTTP requests are made.
Instead of using the default http.Transport, we’ll replace it with our own implementation.
To implement a transport, we’ll have to implement http.RoundTripper interface.
*/
func TestAPI_DoStuff_Transport(t *testing.T) {
	mockClient := &http.Client{}
	type fields struct {
		Client  *http.Client
		baseURL string
	}
	tests := []struct {
		name    string
		fields  fields
		want    int
		want1   []byte
		wantErr bool
		setup   func() //Test Fixtures //not Auto Generated
	}{
		{
			name: "Ok Response with Body",
			fields: fields{
				Client:  mockClient,
				baseURL: "http://example.com",
			},
			setup: func() {
				mockClient.Transport = RoundTripFuncMock(func(req *http.Request) *http.Response {
					// Test request parameters
					stringCompare(t, req.URL.String(), "http://example.com/some/path")
					return &http.Response{
						StatusCode: 200,
						// Send response to be tested
						Body: ioutil.NopCloser(bytes.NewBufferString(`OK`)),
						// Must be set to non-nil value or it panics
						Header: make(http.Header),
					}
				})
			},
			want:  200,
			want1: []byte("OK"),
		},
		{
			name: "400 Status Code with Body",
			fields: fields{
				Client:  mockClient,
				baseURL: "http://example.com",
			},
			setup: func() {
				mockClient.Transport = RoundTripFuncMock(func(req *http.Request) *http.Response {
					// Compare request parameters
					stringCompare(t, req.URL.String(), "http://example.com/some/path")
					return &http.Response{
						StatusCode: 400,
						// Send response to be tested
						Body: ioutil.NopCloser(bytes.NewBufferString(`Bad Request Received`)),
						// Must be set to non-nil value or it panics
						Header: make(http.Header),
					}
				})
			},
			want:  400,
			want1: []byte("Bad Request Received"),
		},
	}
	for _, tt := range tests {
		tt.setup()
		t.Run(tt.name, func(t *testing.T) {
			api := &API{
				Client:  tt.fields.Client,
				baseURL: tt.fields.baseURL,
			}
			got, got1, err := api.DoStuff()
			if (err != nil) != tt.wantErr {
				t.Errorf("API.DoStuff() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("API.DoStuff() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("API.DoStuff() = %v, want %v", string(got1), tt.want)
			}
		})
	}
}

/****************************** Mocks  *****************************************/

// RoundTripFunc .
type RoundTripFuncMock func(req *http.Request) *http.Response

// RoundTrip .
func (f RoundTripFuncMock) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

func stringCompare(t *testing.T, got string, expected string) {
	if got != expected {
		t.Errorf("API.DoStuff() got = %v, want %v", got, expected)
	}

}

/****************************** Mocks End *****************************************/
