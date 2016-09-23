package main

import (
	_ "encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type wilayah struct {
	Data []interface{} `json:"data"`
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
	// var data wilayah
	// var data1 map[string]interface{}
	client := &http.Client{}

	response, _ := client.Get(apiWilayah)
	// request, _ := http.NewRequest("GET", apiWilayah, nil)

	// response, _ := client.Do(request)

	defer response.Body.Close()

	respBody, _ := ioutil.ReadAll(response.Body)

	fmt.Println("data :", string(respBody))

	// for _, v := range data {
	// 	fmt.Printf("Kode Wilayah: %s\t Nama: %s\t Mst Kode Wilayah:	%s\n", data.kode_wilayah, data.nama, data.mst_kode_wilayah)
	// }
}
