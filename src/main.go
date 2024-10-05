package main

import (
	"bruteforce/src/cli"
	"bruteforce/src/query"
	"bruteforce/src/wordlistgen"

	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	forcingParams, err := cli.ParseCliArgs()

	if err != nil {
		panic(err)
	}
	wordlistgen.OverwriteWordlist(&forcingParams)
	query.MainRequest(&forcingParams)
}
