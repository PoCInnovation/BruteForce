package cli

import (
	"bruteforce/src/models"
	"errors"
	"flag"
)

func Parse_cli_args() (models.Forcing_params, error) {
	var params models.Forcing_params
	UrlError := errors.New("No url given")
	ThreadsError := errors.New("Wrong number of threads given")

	forkptr := flag.Bool("v", false, "Verbose program")

	flag.IntVar(&params.Threads, "threads", 1, "Number of threads to be used")
	flag.Parse()

	if params.Threads < 1 {
		return params, ThreadsError
	}
	if len(flag.Args()) != 1 {
		return params, UrlError
	}
	params.Url = flag.Args()[0]
	params.BoolFlags.Verbose = *forkptr
	return params, nil
}
