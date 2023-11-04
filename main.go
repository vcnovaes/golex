package main

import (
	"fmt"

	"github.com/vcnovaes/golex/lexer"
)

func main() {

	var lexer lexer.Lexer

	lexer.Init("template.json")

	token, has := lexer.GetToken("whileflfl")
	fmt.Print(token.Id, has)
}
