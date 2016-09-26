package testing

import (
	"testing"
)

func TestUnmarshalJsonGET(t *testing.T) {

	api := "http://jendela.data.kemdikbud.go.id/api/index.php/CWilayah/wilayahGET"
	request, err := NewRequest(api, "GET")

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
