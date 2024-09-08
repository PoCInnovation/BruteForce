package matcher

import (
	"bruteforce/src/models"
	"log"
	"net/http"
	"strings"
)

func MatchResponse(response *http.Response, body []byte, criteria models.MatchCriteria) error {
	if err := matchStatusCode(response, criteria); err != nil {
		return err
	}
	if err := matchHeaders(response, criteria); err != nil {
		return err
	}
	if err := matchContents(body, criteria); err != nil {
		return err
	}

	return nil
}

func MatchParser(statusPtr string, headerPtr string, bodyPtr string) models.MatchCriteria {
	matchCodes, err := parseStatusCodes(statusPtr)
	if err != nil {
		log.Fatal("Error parsing status codes:", err)
	}

	matchHeaders := parseHeaders(headerPtr)
	criteria := models.MatchCriteria{
		StatusCodes:  matchCodes,
		Headers:      models.HeaderMatch {Headers: matchHeaders, MatchAllHeader: false},
		BodyContains: models.BodyMatch {BodyContains: bodyPtr, MatchAllBody: false},
	}

	return criteria
}
