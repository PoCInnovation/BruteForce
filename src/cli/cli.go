package cli

import (
	"bruteforce/src/models"
	"errors"
	"flag"
)

func Parse_cli_args() (models.Forcing_params, error) {
	var params models.Forcing_params
	myError := errors.New("Wrong parameters")
	forkptr := flag.Bool("v", false, "Verbose program")

	flag.IntVar(&params.Threads, "threads", 1, "Number of threads to be used")
	flag.Parse()

	if len(flag.Args()) != 1 {
		return params, myError
	}
	params.Url = flag.Args()[0]
	params.BoolFlags.Verbose = *forkptr
	return params, nil
}
