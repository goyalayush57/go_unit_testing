package unittesting

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

/*

######################## Testing HTTP calls ########################

1. httptest.NewServer
2. Mock the methods of Client struct
3. Mock Transport layer Roundtripper

*/

//Testing using httptest.NewServer
func TestAPI_DoStuff_NewServer(t *testing.T) {
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
		//Test Fixtures //not auto generated
		setupHandler http.HandlerFunc
	}{
		{
			name: "Success 200",
			setupHandler: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte("OK"))
			},
			want:  200,
			want1: []byte(`OK`),
		},
		{
			name: "Bad Request",
			setupHandler: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(""))
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a test server with the provided handler
			ts := httptest.NewServer(http.HandlerFunc(tt.setupHandler))
			defer ts.Close()
			api := &API{
				Client:  ts.Client(),
				baseURL: ts.URL, //not auto generated
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

//Testing by Replacing http.Transport
/*
Transport specifies the mechanism by which individual HTTP requests are made.
Instead of using the default http.Transport, we’ll replace it with our own implementation.
To implement a transport, we’ll have to implement http.RoundTripper interface.
*/
func TestAPI_DoStuff_Transport(t *testing.T) {
	mockClient := &http.Client{}
	type fields struct {
		Client  *http.Client //added mock when running tc
		baseURL string
	}
	tests := []struct {
		name    string
		fields  fields
		want    int
		want1   []byte
		wantErr bool
		setup   func() *http.Client //Test Fixtures //not Auto Generated
	}{
		{
			name: "Ok Response with Body",
			fields: fields{

				baseURL: "http://example.com",
			},
			setup: func() *http.Client {
				mockClient.Transport = RoundTripFuncMock(func(req *http.Request) (*http.Response, error) {
					// Test request parameters
					stringCompare(t, req.URL.String(), "http://example.com/some/path")
					return &http.Response{
						StatusCode: 200,
						// Send response to be tested
						Body: ioutil.NopCloser(bytes.NewBufferString(`OK`)),
						// Must be set to non-nil value or it panics
						Header: make(http.Header),
					}, nil
				})
				return mockClient
			},
			want:  200,
			want1: []byte("OK"),
		},
		{
			name: "400 Status Code with Body",
			fields: fields{
				baseURL: "http://example.com",
			},
			setup: func() *http.Client {
				mockClient.Transport = RoundTripFuncMock(func(req *http.Request) (*http.Response, error) {
					// Compare request parameters
					stringCompare(t, req.URL.String(), "http://example.com/some/path")
					return &http.Response{
						StatusCode: 400,
						// Send response to be tested
						Body: ioutil.NopCloser(bytes.NewBufferString(`Bad Request Received`)),
						// Must be set to non-nil value or it panics
						Header: make(http.Header),
					}, nil
				})
				return mockClient
			},
			want:  400,
			want1: []byte("Bad Request Received"),
		},
		{
			name: "Get call error",
			fields: fields{
				baseURL: "http://example.com",
			},
			setup: func() *http.Client {
				mockClient.Transport = RoundTripFuncMock(func(req *http.Request) (*http.Response, error) {
					// Compare request parameters
					stringCompare(t, req.URL.String(), "http://example.com/some/path")
					return &http.Response{}, fmt.Errorf("Error while GET redirect")
				})
				return mockClient
			},
			want:    http.StatusInternalServerError,
			want1:   []byte(""),
			wantErr: true,
		},
		{
			name: "GET call error via Dial func",
			fields: fields{
				baseURL: "http://example.com",
			},
			setup: func() *http.Client {
				mockClient.Transport = &http.Transport{
					// Customize the Transport to return an error
					// when making the HTTP request
					Proxy:       nil,
					DialContext: nil,
					Dial: func(network, addr string) (net.Conn, error) {
						return nil, fmt.Errorf("Error while GET redirect")
					},
				}
				return mockClient
			},
			want:    http.StatusInternalServerError,
			want1:   []byte(""),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		mockClient = tt.setup() //not autogenerated
		t.Run(tt.name, func(t *testing.T) {
			api := &API{
				Client:  mockClient,
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
				t.Errorf("API.DoStuff() = %v, want %v", string(got1), tt.want1)
			}
		})
	}
}

/****************************** Mocks  *****************************************/

// RoundTripFunc.
type RoundTripFuncMock func(req *http.Request) (*http.Response, error)

// RoundTrip .
func (f RoundTripFuncMock) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}

func stringCompare(t *testing.T, got string, expected string) {
	if got != expected {
		t.Errorf("API.DoStuff() got = %v, want %v", got, expected)
	}

}

/****************************** Mocks End *****************************************/