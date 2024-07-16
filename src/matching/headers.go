package matcher

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func matchHeaders(resp *http.Response, criteria MatchCriteria) (bool, error) {
	for key, value := range criteria.Headers {
		if resp.Header.Get(key) != value {
			return false, fmt.Errorf("header mismatch: %s=%s\nheaders: %s", key, value, resp.Header)
		}
	}
	return true, nil
}

func parseHeaders(headersList string) map[string]string {
	if headersList == "" {
		return nil
	}

	headers := make(map[string]string)
	headerPairs := strings.Split(headersList, ",")

	for _, pair := range headerPairs {
		parts := strings.SplitN(pair, ":", 2)
		if len(parts) == 2 {
			headers[strings.TrimSpace(parts[0])] = strings.TrimSpace(parts[1])
		} else {
			log.Printf("[WARN] Invalid header format: %s", pair)
		}
	}
	return headers
}
