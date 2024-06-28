package main

import (
	"fmt"
	"log"
	"io/ioutil"
	"net/http"
)

func queryExecute(path string, method string) {

	client := &http.Client{}
	req, err := http.NewRequest(method, "http://localhost:3333" + path, nil)
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

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
	  log.Fatal(err)
	}

	fmt.Println(string(body))
  }