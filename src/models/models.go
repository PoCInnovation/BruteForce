package models

type boolflags struct {
	Verbose bool
}

type MatchCriteria struct {
	StatusCodes  []int
	Headers      map[string]string
	BodyContains string
}

type ForcingParams struct {
	Workers   int
	Url       string
	Wordlist  string
	BoolFlags boolflags
	Criteria  MatchCriteria
}
