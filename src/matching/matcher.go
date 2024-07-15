package matcher

import (
	"log"
)

func MatchParser(url string, statuses string) {
	matchCodes, err := parseStatusCodes(statuses)
	if err != nil {
        log.Fatal("Error parsing status codes:", err)
    }

    matched, msg := matchStatusCode(url, matchCodes)
    if matched {
        log.Println("Matched:", msg)
    } else {
        log.Println("Not matched:", msg)
    }
}
