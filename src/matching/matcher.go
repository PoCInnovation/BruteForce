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

func MatchResponse(response *http.Response, criteria MatchCriteria) (bool, string) {
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return false, err.Error()
	}

	if matched, err := matchStatusCode(response, criteria.StatusCodes); !matched {
		return false, err.Error()
	}
	if matched, err := matchHeaders(response, criteria); !matched {
		return false, err.Error()
	}
	if matched, err := matchContents(body, criteria); !matched {
		return false, err.Error()
	}

	return true, "matched successfully"
}

func MatchParser(statuses string, headers string, bodyContains string) MatchCriteria {
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

	return criteria
}
