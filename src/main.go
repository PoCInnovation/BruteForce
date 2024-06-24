package main

import (
	"flag"
	"log"
	"fmt"
	// "BruteForce/src/matcher"
)

var (
    WarningLogger *log.Logger
    InfoLogger    *log.Logger
    ErrorLogger   *log.Logger
)

func main() {
	matchPtr := flag.String("match", "", "a string")
	usagePtr := flag.Bool("help", false, "a bool")

	flag.Parse()
	if *usagePtr {
		fmt.Println("Usage: use flag match=<option>")
		return;
	}
	if *matchPtr != "" {
		log.Println("match:", *matchPtr)
		return;
	}
	log.Println("No match indicated.")
}
