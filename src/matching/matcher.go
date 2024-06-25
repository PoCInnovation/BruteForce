package matcher

import (
	"fmt"
	"strings"
	"log"
	"net/http"
)

func matchStatusCode(url string, matchCodes []int) (bool, string) {
    resp, err := http.Get(url)
    if err != nil {
        return false, ""
    }
    defer resp.Body.Close()

    for _, code := range matchCodes {
        if resp.StatusCode == code {
            return true, fmt.Sprintf("Status Code: %d", code)
        }
    }
    return false, fmt.Sprintf("Status Code: %d", resp.StatusCode)
}

func parseStatusCodes(statusCodeList string) ([]int, error) {
    codeStrs := strings.Split(statusCodeList, ",")
    var codes []int
    for _, codeStr := range codeStrs {
        var code int
		if _, err := fmt.Sscanf(codeStr, "%d", &code); err != nil {
            return nil, err
        }
		if code < 600 && code >= 100 {
        	codes = append(codes, code)
		} else {
			fmt.Printf("[WARN] `%d` not considered, invalid status code.", code)
		}
    }
	if len(codes) == 1 && codes[0] == 0 {
		return nil, fmt.Errorf("no valid status code given")
	}
    return codes, nil
}

func MatchParser(statuses string) {
	log.Println("Matching using the following status codes:", statuses)
	matchCodes, err := parseStatusCodes(statuses)
	if err != nil {
        log.Fatal("Error parsing status codes:", err)
    }
	log.Println("Matching status codes:", matchCodes)
}
