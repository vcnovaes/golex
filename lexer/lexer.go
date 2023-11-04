package lexer

import (
	"sync"

	"github.com/vcnovaes/golex/loader"
	"github.com/vcnovaes/golex/typ"
)

type Lexer struct {
	tokens []typ.Token
}

func (lex *Lexer) Init(sourcefile string) {
	lex.tokens = loader.GetInputTokens(sourcefile)
}

func (lex *Lexer) GetToken(word string) (typ.Token, []int) {
	var waitGroup sync.WaitGroup
	type LocationToken struct {
		token    typ.Token
		location []int
	}
	channel := make(chan LocationToken)

	for _, tokenStruct := range lex.tokens {
		waitGroup.Add(1)
		go func(token typ.Token) {
			defer waitGroup.Done()
			matches := token.Regex.FindStringIndex(word)
			if len(matches) > 0 {
				channel <- LocationToken{
					token:    token,
					location: matches,
				}
			}
		}(tokenStruct)
	}

	go func() {
		waitGroup.Wait()
		close(channel)
	}()

	select {
	case locationToken, ok := <-channel:
		if ok {
			return locationToken.token, locationToken.location
		}
	}
	return typ.Token{}, []int{}

}
