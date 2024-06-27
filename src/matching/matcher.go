package matcher

import (
	"log"
)

func MatchParser(url string, statuses string) {
	log.Println("Matching using the following status codes:", statuses)

	matchCodes, err := parseStatusCodes(statuses)
	if err != nil {
        log.Fatal("Error parsing status codes:", err)
    }
	log.Printf("Matching status codes %d...", matchCodes)

    matched, msg := matchStatusCode(url, matchCodes)
    if matched {
        log.Println("Matched:", msg)
    } else {
        log.Println("Not matched:", msg)
    }
}
