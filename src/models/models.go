package models

type PromptBuilder struct {
	SiteKeywords []string `json:"sitewords"`
	TechKeywords []string `json:"techwords"`
	WordlistLen  int      `json:"len"`
}

type boolflags struct {
	Verbose    bool
	BodyToFile bool
	Generate   bool
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
	Workers     int
	Url         string
	Wordlist    string
	BoolFlags   boolflags
	Criteria    MatchCriteria
	Data        string
	Method      string
	PromptInfo  PromptBuilder
	ScrapeLevel string
}
