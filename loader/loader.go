package loader

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"

	"github.com/vcnovaes/golex/typ"
)

func GetInputTokens(sourcefile string) []typ.Token {
	file := loadInputFile(sourcefile)
	return decodeInputTokens(file)
}

func loadInputFile(sourcefile string) *os.File {
	file, err := os.Open(sourcefile)
	if err != nil {
		fmt.Printf("Error, not possible to open file %s", sourcefile)
		fmt.Println(err)
		os.Exit(-1)
	}
	return file
}

func decodeInputTokens(file *os.File) []typ.Token {
	var tokens []typ.TokenRaw
	fileDecoder := json.NewDecoder(file)
	err := fileDecoder.Decode(&tokens)
	var finalToken []typ.Token
	for _, tokenRaw := range tokens {
		rgx, err := regexp.Compile(tokenRaw.Regex)
		if err != nil {
			fmt.Println("Error compiling regex", err)
			os.Exit(-1)
		}
		token := typ.Token{
			Id:    tokenRaw.Id,
			Regex: rgx,
		}
		finalToken = append(finalToken, token)

	}
	if err != nil {
		print("Bad decode", err)
		os.Exit(-1)
	}
	defer file.Close()
	return finalToken
}
