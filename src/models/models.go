package models

type boolflags struct {
	Verbose bool
}

type Forcing_params struct {
	Workers   int
	Url       string
	Wordlist  string
	BoolFlags boolflags
	Status    string
	Header    string
	Body      string
}
