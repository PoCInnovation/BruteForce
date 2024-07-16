package matcher

import (
	"io/ioutil"
	"log"
	"net/http"
)

type MatchCriteria struct {
	StatusCodes  []int
	Headers      map[string]string
	BodyContains string
}

func matchResponse(url string, criteria MatchCriteria) (bool, string) {
	status := true
	matchError := ""

	resp, err := http.Get(url)
	if err != nil {
		return false, err.Error()
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err.Error()
	}

	status, matchError = matchStatusCode(resp, criteria.StatusCodes)
	if !status {
		return false, matchError
	}
	status, matchError = matchHeaders(resp, criteria)
	if !status {
		return false, matchError
	}
	status, matchError = matchContents(body, criteria)
	if !status {
		return false, matchError
	}

	return true, "Matched successfully for " + url
}

func MatchParser(url string, statuses string, headers string, bodyContains string) {
	matchCodes, err := parseStatusCodes(statuses)
	if err != nil {
		log.Fatal("Error parsing status codes:", err)
	}

	matchHeaders := parseHeaders(headers)
	criteria := MatchCriteria{
		StatusCodes:  matchCodes,
		Headers:      matchHeaders,
		BodyContains: bodyContains,
	}

	matched, msg := matchResponse(url, criteria)
	if matched {
		log.Println("Matched:", msg)
	} else {
		log.Println("Not matched:", msg)
	}
}
