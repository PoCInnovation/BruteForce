package matcher

import (
	"bruteforce/src/models"
	"fmt"
	"log"
	"strings"
)

func matchContents(body []byte, criteria models.MatchCriteria) error {
	bodyStr := string(body)

	if len(criteria.Body.BodyContains) > 0 {
		for _, content := range criteria.Body.BodyContains {
			if !strings.Contains(bodyStr, content) {
				return fmt.Errorf("body content mismatch: missing %s", content)
			}
		}
	}
	return nil
}

func parseBody(body string) models.BodyMatch {
	if body == "" {
		return models.BodyMatch{}
	}

	parts := strings.Split(body, ",")
	firstPart := strings.TrimSpace(parts[0])
	mode := true

	if firstPart == "all" {
		parts = parts[1:]
		log.Println("[INFO] Matching criteria for body is set to check if.")
	} else if firstPart == "one" {
		mode = !mode
		parts = parts[1:]
		log.Println("[INFO] Matching criteria for body is set to check if one true.")
	}

	var parsedBody []string
	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part != "" {
			parsedBody = append(parsedBody, part)
		}
	}
	return models.BodyMatch{BodyContains: parsedBody, MatchAllBody: mode}
}
