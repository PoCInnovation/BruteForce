package cli

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"os"

	"bruteforce/src/matching"
	"bruteforce/src/models"
)

func errorHandling(params models.ForcingParams) (models.ForcingParams, error) {
	ThreadsError := errors.New("wrong number of threads given")
	WordListError := errors.New("no wordlist given")
	DataError := errors.New("Invalid JSON data")
	methodError := errors.New("Invalid HTTP method")
	var method_list = [7]string{"POST", "GET", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}

	if params.Data != "" && !json.Valid([]byte(params.Data)) {
		return params, DataError
	}
	if params.Workers < 1 {
		return params, ThreadsError
	}
	if params.Wordlist == "" && params.BoolFlags.Generate == false {
		return params, WordListError
	}
	for i := 0; i < 7; i++ {
		if params.Method == method_list[i] {
			return params, nil
		}
	}
	return params, methodError
}

func ParseCliArgs() (models.ForcingParams, error) {
	var params models.ForcingParams
	UrlError := errors.New("no url given")

	forkptr := flag.Bool("v", false, "Verbose program")
	printbodyptr := flag.Bool("p", false, "Prints body when matched")
	statusPtr := flag.String("status-codes", "200,401,403,404,429,500", "Comma-separated list of status codes to match")
	headerPtr := flag.String("header", "", "Header to match, formatted as \"key: value\"")
	bodyPtr := flag.String("body", "", "String to match in response body")
	wordlistPtr := flag.String("wordlist", "", "Wordlist to bruteforce url with")
	methodPtr := flag.String("method", "GET", "Method to bruteforce with")
	postDataptr := flag.String("data", "", "JSON Data to inlude in body when bruteforcing")
	scrapPtr := flag.String("scrape", "", "Scrape the original url for keywords to generate the wordlist: ALL/TECH/SITE")
	genptr := flag.String("generate", "", "JSON to generate a custom wordlist using chat gpt based on your api key: {\"sitewords\": [\"?\", ...], \"techwords\": [\"?\", ...], \"len\": ?}")
	flag.IntVar(&params.Workers, "threads", 1, "Number of threads to be used")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: bruteforce [options] <url>\n")
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
	}

	flag.Parse()

	if len(flag.Args()) < 1 {
		return params, UrlError
	}

	params.Url = flag.Args()[0]
	params.BoolFlags.Verbose = *forkptr
	params.Wordlist = *wordlistPtr
	params.Criteria = matcher.MatchParser(*statusPtr, *headerPtr, *bodyPtr)
	params.Data = *postDataptr
	params.Method = *methodPtr
	params.BoolFlags.BodyToFile = *printbodyptr
	params.ScrapeLevel = *scrapPtr

	err := json.Unmarshal([]byte(*genptr), &params.PromptInfo)
	params.BoolFlags.Generate = !(err != nil)

	return errorHandling(params)
}
