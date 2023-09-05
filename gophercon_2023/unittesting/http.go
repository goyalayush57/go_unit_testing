package unittesting

import (
	"io/ioutil"
	"net/http"
)

type API struct {
	Client  *http.Client
	baseURL string
}

func (api *API) DoGoodStuff() (statuscode int, body []byte, err error) {
	resp, err := api.Client.Get(api.baseURL + "/some/path")
	if err != nil {
		return http.StatusInternalServerError, []byte(""), err
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	//perform operations here
	return resp.StatusCode, body, err
}

func DoBadStuff() (statuscode int, body []byte, err error) {
	resp, err := http.Get("example.com" + "/some/path") //internally calls DefaultClient
	if err != nil {
		return http.StatusInternalServerError, []byte(""), err
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	//perform operations here
	return resp.StatusCode, body, err
}
