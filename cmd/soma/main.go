// Copyright 2013 Ben Johnson. All rights reserved.
// Copyright 2022 Mark Stahl. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the MIT-LICENSE file.
package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	// "github.com/socialmachines/soma"
)

var (
	VERSION = "0.1.0"

	// Returned when the command usage is printed and
	// the process should exit with code 2
	ErrorUsage = errors.New("usage")

	// The CLI receives a command that is not known
	ErrorUnknownCommand = errors.New("unknown command")
)

func main() {
	m := NewMain()
	if err := m.Run(os.Args[1:]...); err == ErrorUsage {
		os.Exit(2)
	} else if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

type Main struct {
	Stdin  io.Reader
	Stdout io.Writer
	Stderr io.Writer
}

func NewMain() *Main {
	return &Main{
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}
}

func (m *Main) Run(args ...string) error {
	if len(args) == 0 || strings.HasPrefix(args[0], "-") {
		fmt.Fprintln(m.Stderr, m.Usage())
		return ErrorUsage
	}
	switch args[0] {
	default:
		return ErrorUnknownCommand
	}
}

func (m *Main) Usage() string {
	return strings.TrimLeft(fmt.Sprintf(`
Social Machines v%s

Usage:
    soma [command [arguments*]]

Command:
    tokens    Prints the list of tokens for a provided string

Use "soma <command> -h" for more information about the command
`, VERSION), "\n")
}
