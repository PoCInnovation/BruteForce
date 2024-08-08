package matcher

import (
	"bruteforce/src/models"
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

func MatchParser(params *models.Forcing_params) MatchCriteria {
	matchCodes, err := parseStatusCodes(params.Status)
	if err != nil {
		log.Fatal("Error parsing status codes:", err)
	}

	matchHeaders := parseHeaders(params.Header)
	criteria := MatchCriteria{
		StatusCodes:  matchCodes,
		Headers:      matchHeaders,
		BodyContains: params.Body,
	}

	return criteria
}
