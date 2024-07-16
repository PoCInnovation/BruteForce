package matcher

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func matchStatusCode(resp *http.Response, matchCodes []int) (bool, string) {
	isAll := false

	if matchCodes[0] == 0 {
		isAll = !isAll
	} else {
		log.Printf("Matching status codes %d...", matchCodes)
	}
	for _, code := range matchCodes {
		if resp.StatusCode == code || isAll {
			return true, fmt.Sprintf("status code is %d", resp.StatusCode)
		}
	}
	return false, fmt.Sprintf("status code is %d", resp.StatusCode)
}

func parseStatusCodes(statusCodeList string) ([]int, error) {
	codeStrs := strings.Split(statusCodeList, ",")
	if statusCodeList == "all" {
		log.Println("Matching all status codes")
		return []int{0}, nil
	}

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
