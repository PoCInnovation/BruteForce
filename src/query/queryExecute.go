package query

import (
	"bruteforce/src/models"
	"fmt"
	"io"
	"log"
	"net/http"
)

func QueryExecute(params *models.Forcing_params, path string, method string) {
	client := &http.Client{}
	req, err := http.NewRequest(method, params.Url+path, nil)
	if err != nil {
		log.Fatal(err)
	}

	q := req.URL.Query()
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))
}
