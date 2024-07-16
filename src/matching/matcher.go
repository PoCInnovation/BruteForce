package matcher

import (
	"io"
	"log"
	"net/http"
)

type MatchCriteria struct {
	StatusCodes  []int
	Headers      map[string]string
	BodyContains string
}

func matchResponse(url string, criteria MatchCriteria) (bool, string) {
	resp, err := http.Get(url)
	if err != nil {
		return false, err.Error()
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, err.Error()
	}

	if matched, err := matchStatusCode(resp, criteria.StatusCodes); !matched {
		return false, err.Error()
	}
	if matched, err := matchHeaders(resp, criteria); !matched {
		return false, err.Error()
	}
	if matched, err := matchContents(body, criteria); !matched {
		return false, err.Error()
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
