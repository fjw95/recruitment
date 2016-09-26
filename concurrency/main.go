package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Object struct
type Wilayah struct {
	data []interface{}
}

/*type dataWilayah struct {
	kode_wilayah     string
	nama             string
	mst_kode_wilayah string
}*/

var (
	apiMuseum  = "http://jendela.data.kemdikbud.go.id/api/index.php/CcariMuseum/searchGET"
	apiWilayah = "http://jendela.data.kemdikbud.go.id/api/index.php/CWilayah/wilayahGET"
)

func main() {

	// declare struct wilayah
	var datWil Wilayah

	// declare var client
	client := &http.Client{}

	// declare var new request
	req, err := http.NewRequest("GET", apiWilayah, nil)

	// check if error
	if err != nil {
		fmt.Println("error:", err)
	}

	// get response from request
	response, err := client.Do(req)
	if err != nil {
		fmt.Println("error:", err)
	}

	// close response body
	defer response.Body.Close()

	// get bytes code response body
	respBody, err := ioutil.ReadAll(response.Body)

	// check if null
	if err != nil {
		fmt.Println("error:", err)
	}

	/*
		remove bytes code sesuai refernsi
		http://stackoverflow.com/questions/31398044/got-error-invalid-character-%C3%AF-looking-for-beginning-of-value-from-json-unmar
		jika tidak remove akan --> error: invalid character 'ï' looking for beginning of value
	*/
	respBody = bytes.TrimPrefix(respBody, []byte{239, 187, 191})

	err = json.Unmarshal(respBody, &datWil)

	// check if error
	if err != nil {
		fmt.Println("error:", err)
	}

	// print response body
	fmt.Println("response Body : ", string(respBody))
	// print data hasil unmarshal json
	fmt.Println("\ndata :", datWil.data)

	// var data1 map[string]interface{}
	// request, _ := http.NewRequest("GET", apiWilayah, nil)

	// response, _ := client.Do(request)
	// for _, v := range data {
	// 	fmt.Printf("Kode Wilayah: %s\t Nama: %s\t Mst Kode Wilayah:	%s\n", data.kode_wilayah, data.nama, data.mst_kode_wilayah)
	// }
}
