package matcher

import (
	"bruteforce/src/models"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func matchHeaders(resp *http.Response, criteria models.MatchCriteria) error {
	for key, value := range criteria.Header.Headers {
		if resp.Header.Get(key) != value {
			return fmt.Errorf("header mismatch: %s=%s\nheaders: %s", key, value, resp.Header)
		}
	}
	return nil
}

func parseHeaders(headersList string) models.HeaderMatch {
	if headersList == "" {
		return models.HeaderMatch{}
	}

	headers := make(map[string]string)
	headerPairs := strings.Split(headersList, ",")
	mode := true
	firstPair := strings.TrimSpace(headerPairs[0])

	if strings.HasPrefix(firstPair, "all") {
		headerPairs = headerPairs[1:]
		log.Println("[INFO] Matching criteria for header is set to check if.")
	} else if strings.HasPrefix(firstPair, "one") {
		mode = !mode
		headerPairs = headerPairs[1:]
		log.Println("[INFO] Matching criteria for header is set to check if one true.")
	}

	for _, pair := range headerPairs {
		parts := strings.SplitN(pair, ":", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])

			log.Printf("[INFO] Parsed header: %s=%s", key, value)
			headers[key] = value
		} else {
			log.Printf("[WARN] Invalid header format: %s", pair)
		}
	}
	return models.HeaderMatch{Headers: headers, MatchAllHeader: mode}
}
