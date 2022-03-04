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
	TOK_ILLEGAL Token = iota
	TOK_EOF

	// Lexical types
	TOK_COMMENT // 'This is a comment'
	TOK_STRING  // "This is a string"
	TOK_INT     // 122334

	// Identifiers
	TOK_LOWER_IDENT   // firstName
	TOK_UPPER_IDENT   // Person
	TOK_LOWER_KEYWORD // ifTrue:
	TOK_UPPER_KEYWORD // Else:
	TOK_BINARY        // '!', '*', '/', '+', '|', '&', '-', '>', '<', '=', '?', '\', '~', '^', '%'
	TOK_ATTR_GET      // @name
	TOK_ATTR_SET      // @name:

	// Grouping
	TOK_LEFT_BRACE  // { blocks
	TOK_RIGHT_BRACE // }
	TOK_LEFT_BRACK  // [ maps, arrays
	TOK_RIGHT_BRACK // ]
	TOK_LEFT_PAREN  // ( grouping
	TOK_RIGHT_PAREN // )

	// Assignment
	TOK_ASSIGN // :=
	TOK_DEFINE // ->

	// Puncuation
	TOK_COMMA      // ,
	TOK_SEMI_COLON // ;
	TOK_PERIOD     // .
)

var tokens = []string{
	TOK_ILLEGAL: "ILLEGAL",
	TOK_EOF:     "EOF",

	TOK_COMMENT: "COMMENT",
	TOK_STRING:  "STRING",
	TOK_INT:     "INT",

	TOK_LOWER_IDENT:   "IDENT",
	TOK_UPPER_IDENT:   "IDENT",
	TOK_LOWER_KEYWORD: "KEYWORD",
	TOK_UPPER_KEYWORD: "KEYWORD",
	TOK_BINARY:        "BINARY",

	TOK_LEFT_BRACE:  "LEFT_BRACE",
	TOK_RIGHT_BRACE: "RIGHT_BRACE",
	TOK_LEFT_BRACK:  "LEFT_BRACK",
	TOK_RIGHT_BRACK: "RIGHT_BRACK",
	TOK_LEFT_PAREN:  "LEFT_PAREN",
	TOK_RIGHT_PAREN: "RIGHT_PAREN",

	TOK_ASSIGN: "ASSIGN",
	TOK_DEFINE: "DEFINE",

	TOK_COMMA:      "COMMA",
	TOK_SEMI_COLON: "SEMI-COLON",
	TOK_PERIOD:     "PERIOD",
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
