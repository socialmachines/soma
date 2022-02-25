// Copyright 2022 Mark Stahl. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package main

import (
	"strings"
	"testing"
)

func Test_NoCommand_ErrorUsage(t *testing.T) {
	if _, err := Run(); err != ErrorUsage {
		t.Fatal(err)
	}
}

func Test_TokensCommand(t *testing.T) {
	output, err := Run("tokens")
	if err != nil {
		t.Fatal(err)
	}
	expected, _ := TokensUsage()
	if output != expected {
		t.Fatalf("unexpected output:\n\n%s", output)
	}
}

func Test_TokensCommand_WithArgs(t *testing.T) {
	tokens := []string{
		"UPPER_IDENT (True)",
		":= (:=)",
		"UPPER_IDENT (Object)",
		"LOWER_IDENT (new)",
		". (.)",
	}
	expected := strings.Join(tokens, "\n")

	output, err := Run("tokens", "True := Object new.")
	if err != nil {
		t.Fatal(err)
	}
	if output != expected {
		t.Fatalf("unexpected output:\n\n%s", output)
	}
}
