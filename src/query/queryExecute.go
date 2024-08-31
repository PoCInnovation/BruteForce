package query

import (
	"bruteforce/src/matching"
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
	log.Printf("NewRequest(%s)", params.Url+path)

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

	if err := matcher.MatchResponse(resp, body, params.Criteria); err == nil {
		fmt.Println(string(body))
	} else {
		log.Println(err)
	}
}
