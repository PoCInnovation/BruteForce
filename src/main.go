package main

import (
	"bruteforce/src/cli"
	"fmt"
)

func main() {
	forcing_params, err := cli.Parse_cli_args()

	if err != nil {
		panic(err)
	}
	fmt.Println(forcing_params)
}
