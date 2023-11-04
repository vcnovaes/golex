package typ

import "regexp"

type TokenRaw struct {
	Id    string `json:"token"`
	Regex string `json:"regex"`
}

type Token struct {
	Id    string
	Regex *regexp.Regexp
}

type Lexemma struct {
	location []int
	token    string
}
