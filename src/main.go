package main

import (
	"bruteforce/src/cli"
	"bruteforce/src/query"
	"fmt"
)

func main() {

	forcingParams, err := cli.ParseCliArgs()

	if err != nil {
		panic(err)
	}
	fmt.Println(forcingParams)

	query.MainRequest(&forcingParams)
}
