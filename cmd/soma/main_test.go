package main

import (
	"testing"
)

func TestExe(t *testing.T) {
	m := NewMain()
	if err := m.Run("hello world"); err != nil {
		t.Fatal(err)
	}
}
