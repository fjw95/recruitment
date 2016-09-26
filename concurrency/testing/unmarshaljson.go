package testing

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

var (
	client = &http.Client{}
)

type Wilayah struct {
	data []interface{}
}

func NewRequest(api string, method string) (*http.Request, error) {

	request, err := http.NewRequest(method, api, nil)

	if err != nil {
		return nil, err
	}

	return request, nil
}

func GetResponse(request *http.Request) (*http.Response, error) {

	response, err := client.Do(request)

	if err != nil {
		return nil, err
	}

	return response, nil
}

func ReadResponseBody(response *http.Response) ([]byte, error) {

	responseBody, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	return responseBody, nil
}

func UnmarshallJson(responseBody []byte) ([]interface{}, error) {

	var w Wilayah

	responseBody = bytes.TrimPrefix(responseBody, []byte{239, 187, 191})

	err := json.Unmarshal(responseBody, &w)

	if err != nil {
		return nil, err
	}

	return w.data, nil
}
