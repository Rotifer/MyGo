package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// Generated by curl-to-Go: https://mholt.github.io/curl-to-go

	// curl 'https://rest.ensembl.org/lookup/id/ENSG00000157764?expand=1' -H 'Content-type:application/json'

	req, err := http.NewRequest("GET", "https://rest.ensembl.org/lookup/id/ENSG00000157764?expand=1", nil)
	if err != nil {
		// handle err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body) // response body is []byte
	fmt.Println(string(body))
}
