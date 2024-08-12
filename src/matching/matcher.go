package matcher

import (
	"bruteforce/src/models"
	"log"
	"net/http"
)

func MatchResponse(response *http.Response, body []byte, criteria models.MatchCriteria) error {
	if matched, err := matchStatusCode(response, criteria); !matched {
		return err
	}
	if matched, err := matchHeaders(response, criteria); !matched {
		return err
	}
	if matched, err := matchContents(body, criteria); !matched {
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
		Headers:      matchHeaders,
		BodyContains: bodyPtr,
	}

	return criteria
}
