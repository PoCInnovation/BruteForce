package main

import (
	"bruteforce/src/cli"
	"bruteforce/src/matching"
	"bruteforce/src/query"
	"fmt"
)

func main() {

	forcing_params, err := cli.Parse_cli_args()

	if err != nil {
		panic(err)
	}
	fmt.Println(forcing_params)

	criteria := matcher.MatchParser(&forcing_params)

	query.MainRequest(&forcing_params, criteria) // maybe like this?
}
