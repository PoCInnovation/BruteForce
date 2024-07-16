package main

import (
	"bruteforce/src/matching"
	"flag"
	"fmt"
	"log"
)

func main() {
	usagePtr := flag.Bool("help", false, "a bool")
	statusPtr := flag.String("status-codes", "200,401,403,404,429,500", "Comma-separated list of status codes to match")
	headerPtr := flag.String("header", "", "Header to match")
	bodyPtr := flag.String("body", "", "String to match in response body")

	flag.Parse()
	if *usagePtr {
		fmt.Println("Usage for matching options:\n\nuse flag -status-codes=\"<status-options>\"")
		fmt.Println("\tstatus-options\t200,202,400,404 (Comma-separated list)")
		fmt.Println("\t\t\tall")
		fmt.Println("\t\t\t[By default: 200,401,403,404,429,500]")
		fmt.Println("\nuse flag -header=\"<key:value>\"")
		fmt.Println("use flag -body=\"<string-in-body>\"")
		return
	}

	criteria := matcher.MatchParser(*statusPtr, *headerPtr, *bodyPtr)
	result, err := matcher.MatchResponse("http://example.com", criteria)
	if !result {
		log.Fatal(err)
	}
}
