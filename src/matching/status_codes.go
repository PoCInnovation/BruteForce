package matcher

import (
	"fmt"
	"strings"
	"net/http"
)


func matchStatusCode(url string, matchCodes []int) (bool, string) {
    resp, err := http.Get(url)
    if err != nil {
        return false, err.Error()
    }
    defer resp.Body.Close()

    for _, code := range matchCodes {
        if resp.StatusCode == code {
            return true, fmt.Sprintf("status code is %d", code)
        }
    }
    return false, fmt.Sprintf("status code is %d", resp.StatusCode)
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
