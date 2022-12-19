// Copyright 2022 Mark Stahl. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
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
		return MainUsage()
	}
	switch args[0] {
	case "tokens":
		return TokensCommand(args[1:]...)
	case "help":
		usage := CommandUsage[args[1]]
		if usage != nil {
			return usage()
		} else {
			return MainUsage()
		}
	default:
		return MainUsage()
	}
}

func MainUsage() (string, error) {
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

var CommandUsage = map[string](func() (string, error)){
	"tokens": TokensUsage,
}
