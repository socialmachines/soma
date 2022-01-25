package main

import (
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/socialmachines/soma"
)

var (
	ErrorUsage = errors.New("usage")
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
	fmt.Println(args)
	fmt.Println(soma.ILLEGAL)
	return nil
}
