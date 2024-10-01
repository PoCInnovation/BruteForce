package main

import (
	"bruteforce/src/cli"
	"bruteforce/src/query"
)

func main() {

	forcingParams, err := cli.ParseCliArgs()

	if err != nil {
		panic(err)
	}

	query.MainRequest(&forcingParams)
}
