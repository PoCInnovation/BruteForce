package query

import (
	"bruteforce/src/matching"
	"bruteforce/src/models"
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
)

func QueryExecute(params *models.ForcingParams, path string, method string) {
	client := &http.Client{}
	body_req := []byte(params.Data)

	req, err := http.NewRequest(method, params.Url+path, bytes.NewBuffer(body_req))

	if err != nil {
		log.Fatal(err)
	}
	if params.BoolFlags.Verbose {
		log.Printf("NewRequest(%s)", params.Url+path)
	}
	q := req.URL.Query()
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	if err := matcher.MatchResponse(resp, body, params.Criteria, params); err == nil {
		fmt.Println(string(body))
	} else {
		if params.BoolFlags.Verbose {
			log.Println(err)
		}
	}
}
