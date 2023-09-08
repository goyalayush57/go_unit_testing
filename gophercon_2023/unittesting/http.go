package unittesting

import (
	"io/ioutil"
	"net/http"
)

func DoBadStuff(path string) (int, []byte, error) {
	resp, err := http.Get("example.com" + path) //internally calls DefaultClient
	if err != nil {
		return http.StatusInternalServerError, []byte(""), err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	//perform operations here
	return resp.StatusCode, body, err
}

type API struct {
	Client  *http.Client
	baseURL string
}

func (api *API) DoGoodStuff(path string) (int, []byte, error) {
	resp, err := api.Client.Get(api.baseURL + path)
	if err != nil {
		return http.StatusInternalServerError, []byte(""), err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	//perform operations here
	return resp.StatusCode, body, err
}
