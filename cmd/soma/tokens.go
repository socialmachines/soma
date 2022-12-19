package main

import (
	"fmt"
	"strings"

	"github.com/socialmachines/soma"
)

func TokensCommand(args ...string) (string, error) {
	if len(args) != 1 {
		return TokensUsage()
	}
	var s soma.Scanner
	s.FromString(args[0])

	results := []string{}
	_, tok, lit := s.Scan()
	for tok != soma.TOK_EOF {
		results = append(results, fmt.Sprintf("%s (%s)", tok, lit))
		_, tok, lit = s.Scan()
	}
	return strings.Join(results, "\n"), nil
}

func TokensUsage() (string, error) {
	usage := strings.TrimLeft(`
Prints the tokens and literals of the expression.

Usage:
  soma tokens <expression>

Example:
  $ soma tokens "True := Object new."
  UPPER_IDENT (True)
  := (:=)
  UPPER_IDENT (Object)
  LOWER_IDENT (new)
  . (.)
`, "\n")
	return usage, nil
}
