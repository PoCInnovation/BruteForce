package main

import (
	"flag"
	"fmt"
	"bruteforce/src/matching"
)

func main() {
	usagePtr := flag.Bool("help", false, "a bool")
	statusPtr := flag.String("status-codes", "200,401,403,404,429,500", "Comma-separated list of status codes to match")

	flag.Parse()
	if *usagePtr {
		fmt.Println("Usage for matching options:\n\nuse flag -status-codes=\"<status-options>\"")
		fmt.Println("\tstatus-options\t200,202,400,404 (Comma-separated list)")
		fmt.Println("\t\t\tall")
		fmt.Println("\t\t\t[By default: 200,401,403,404,429,500]")
		return;
	}
	matcher.MatchParser("http://example.com", *statusPtr)
}
