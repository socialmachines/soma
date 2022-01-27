// Copyright 2009 The Go Authors. All rights reserved.
// Copyright 2022 Mark Stahl. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package token defines constants representing the lexical tokens of the
// Social Machines programming language and basic operations on tokens
// (printing, predicates).
//
package soma

import (
	"strconv"
)

// Token is the set of lexical tokens of the Social Machines programming
// language.
type Token int

const (
	// Special tokens
	ILLEGAL Token = iota
	EOF

	// Lexical types
	COMMENT // 'This is a comment'
	STRING  // "This is a string"

	// Identifiers
	LOWER_IDENT   // firstName
	UPPER_IDENT   // Person
	LOWER_KEYWORD // ifTrue:
	UPPER_KEYWORD // Else:
	BINARY        // '!', '*', '/', '+', '|', '&', '-', '>', '<', '=', '?', '\', '~':

	// Grouping
	LEFT_BRACE  // {
	RIGHT_BRACE // }
	LEFT_BRACK  // [
	RIGHT_BRACK // ]
	LEFT_PAREN  // (
	RIGHT_PAREN // )

	// Assignment
	DECLARE // :=

	// Puncuation
	FLUENT // ;
	COMMA  // ,
	PERIOD // .
)

var tokens = []string{
	ILLEGAL: "ILLEGAL",
	EOF:     "EOF",

	COMMENT: "COMMENT",
	STRING:  "STRING",

	LOWER_IDENT:   "LOWER_IDENT",
	UPPER_IDENT:   "UPPER_IDENT",
	LOWER_KEYWORD: "LOWER_KEYWRD",
	UPPER_KEYWORD: "UPPER_KEYWRD",
	BINARY:        "BINARY",

	LEFT_BRACE:  "{",
	RIGHT_BRACE: "}",

	DECLARE: ":=",

	FLUENT: ";",
	PERIOD: ".",
}

// String returns the string corresponding to the token tok.
// For punctution, assignment, and groupings the string is the actual
// token character sequence (e.g., for the token COMMA, the string is
// ","). For all other tokens the string corresponds to the token
// constant name (e.g. for the token BINARY, the string is "BINARY").
//
func (tok Token) String() string {
	s := ""
	if 0 <= tok && tok < Token(len(tokens)) {
		s = tokens[tok]
	}
	if s == "" {
		s = "token(" + strconv.Itoa(int(tok)) + ")"
	}
	return s
}
