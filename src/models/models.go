package models

type boolflags struct {
	Verbose bool
}

type HeaderMatch struct {
	Headers        map[string]string
	MatchAllHeader bool
}

type BodyMatch struct {
	BodyContains []string
	MatchAllBody bool
}

type MatchCriteria struct {
	StatusCodes []int
	Header      HeaderMatch
	Body        BodyMatch
}

type ForcingParams struct {
	Workers   int
	Url       string
	Wordlist  string
	BoolFlags boolflags
	Criteria  MatchCriteria
}
