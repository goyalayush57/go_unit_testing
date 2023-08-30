package unit_testing

import (
	"io/ioutil"
	"net/http"
)

type API struct {
	Client  *http.Client
	baseURL string
}

func (api *API) DoStuff() (statuscode int, body []byte, err error) {
	resp, err := api.Client.Get(api.baseURL + "/some/path")
	if err != nil {
		return 500, nil, err
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	return resp.StatusCode, body, err
}
