package matcher

import (
	"bruteforce/src/models"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func matchStatusCode(resp *http.Response, criteria models.MatchCriteria) (bool, error) {
	isAll := false

	if criteria.StatusCodes[0] == 0 {
		isAll = !isAll
	} else {
		log.Printf("Matching status codes %d...", criteria.StatusCodes)
	}
	for _, code := range criteria.StatusCodes {
		if resp.StatusCode == code || isAll {
			return true, nil
		}
	}
	return false, fmt.Errorf("status code is %d", resp.StatusCode)
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
			log.Printf("[WARN] `%d` not considered, invalid status code.", code)
		}
	}

	if len(codes) == 1 && codes[0] == 0 {
		return nil, fmt.Errorf("no valid status code given")
	}
	return codes, nil
}
