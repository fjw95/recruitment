package testing

import (
	"testing"
)

var (
	api = "http://jendela.data.kemdikbud.go.id/api/index.php/CWilayah/wilayahGET"
)

func TestUnmarshalJsonGET(t *testing.T) {

	request, err := NewRequest(api, "POST")

	if err != nil {
		t.Error(err)
	}

	response, err := GetResponse(request)

	if err != nil {
		t.Error(err)
	}

	responseBody, err := ReadResponseBody(response)

	if err != nil {
		t.Error(err)
	}

	wilayah, err := UnmarshallJson(responseBody)

	if err != nil {
		t.Error(err)
	}

	if wilayah == nil {
		t.Errorf("Error : Wilayah cannot nil json data failed to unmarshal")
	}

}
