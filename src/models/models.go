package models

type boolflags struct {
	Verbose bool
}

type Forcing_params struct {
	Threads   int
	Url       string
	BoolFlags boolflags
}
