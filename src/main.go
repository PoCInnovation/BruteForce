package main

import (
	"flag"
	"log"
	"fmt"
	"bruteforce/src/matching"
)

func main() {
	matchPtr := flag.String("match", "", "a string")
	usagePtr := flag.Bool("help", false, "a bool")

	flag.Parse()
	if *usagePtr {
		fmt.Println("Usage: use flag match=<option>")
		fmt.Println("	option	status[100-599|all]")
		return;
	}
	if *matchPtr != "" {
		matcher.MatchParser(*matchPtr)
	} else {
		log.Fatal("No match indicated.")
	}
}
