package matcher

import (
	"bruteforce/src/models"
	"io"
	"log"
	"net/http"
)

func MatchResponse(response *http.Response, criteria models.MatchCriteria) (bool, string) {
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

func MatchParser(statusPtr string, headerPtr string, bodyPtr string) models.MatchCriteria {
	matchCodes, err := parseStatusCodes(statusPtr)
	if err != nil {
		log.Fatal("Error parsing status codes:", err)
	}

	matchHeaders := parseHeaders(headerPtr)
	criteria := models.MatchCriteria{
		StatusCodes:  matchCodes,
		Headers:      matchHeaders,
		BodyContains: bodyPtr,
	}

	return criteria
}
