package cli

import (
	"bruteforce/src/matching"
	"bruteforce/src/models"
	"errors"
	"flag"
	"fmt"
	"os"
)

func Parse_cli_args() (models.Forcing_params, error) {
	var params models.Forcing_params

	UrlError := errors.New("No url given")
	ThreadsError := errors.New("Wrong number of threads given")
	WordListError := errors.New("No wordlist given")

	// forkptr := flag.Bool("v", false, "Verbose program")
	statusPtr := flag.String("status-codes", "200,401,403,404,429,500", "Comma-separated list of status codes to match")
	headerPtr := flag.String("header", "", "Header to match, formatted as \"key: value\"")
	bodyPtr := flag.String("body", "", "String to match in response body")
	wordlistPtr := flag.String("wordlist", "", "Wordlist to bruteforce url with")
	flag.IntVar(&params.Workers, "threads", 1, "Number of threads to be used")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: bruteforce [options] <url>\n")
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
	}

	flag.Parse()

	if params.Workers < 1 {
		return params, ThreadsError
	}
	fmt.Print(flag.Args())
	if len(flag.Args()) < 1 {
		return params, UrlError
	}

	params.Url = flag.Args()[0]
	// params.BoolFlags.Verbose = *forkptr
	params.Criteria = matcher.MatchParser(*statusPtr, *headerPtr, *bodyPtr)
	params.Wordlist = *wordlistPtr
	if params.Wordlist == "" {
		return params, WordListError
	}

	return params, nil
}
