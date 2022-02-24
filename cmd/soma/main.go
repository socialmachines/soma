// Copyright 2013 Ben Johnson. All rights reserved.
// Copyright 2022 Mark Stahl. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the MIT-LICENSE file.
package main

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/socialmachines/soma"
)

var (
	VERSION = "0.1.0"

	// Returned when the command usage is printed and
	// the process should exit with code 2
	ErrorUsage = errors.New("usage")
)

func main() {
	output, err := Run(os.Args[1:]...)
	if err == ErrorUsage {
		fmt.Println(output)
		os.Exit(2)
	}
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(output)
}

func Run(args ...string) (string, error) {
	if len(args) == 0 || strings.HasPrefix(args[0], "-") {
		return Usage()
	}
	switch args[0] {
	case "tokens":
		return TokensCommand(args[1:]...)
	default:
		return Usage()
	}
}

func Usage() (string, error) {
	usage := strings.TrimLeft(fmt.Sprintf(`
Social Machines v%s

Usage:
  soma [command [arguments*]]

Command:
  help <command>      Prints the help for any command
  tokens <expression> Prints the list of tokens for a provided string
`, VERSION), "\n")
	return usage, ErrorUsage
}

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
Prints the tokens and literals for expression provided.

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
