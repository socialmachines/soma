package soma

import (
	"go/token"
	"testing"
)

func TestAssignment(t *testing.T) {
	received := `
		True := Object new.
	`
	expected := []Token{
		UPPER_IDENT,
		DECLARE,
		UPPER_IDENT,
		LOWER_IDENT,
		PERIOD,
	}
	testTokens(t, received, expected)
}

func TestUnary(t *testing.T) {
	received := `
		'True not'
		True externalBehavior: "not" Does: {
			|t|
			False
		}.
	`
	expected := []Token{
		COMMENT,
		UPPER_IDENT,
		LOWER_KEYWORD,
		STRING,
		UPPER_KEYWORD,
		LEFT_BRACE,
		BINARY,
		LOWER_IDENT,
		BINARY,
		UPPER_IDENT,
		RIGHT_BRACE,
		PERIOD,
	}
	testTokens(t, received, expected)
}

func TestKeyword(t *testing.T) {
	received := `
		'True ifTrue: { "do something" } Else: { "do something else" }'
		True externalBehavior: "ifTrue:Else:" Does: {
			|trueBlock. elseBlock. t|
			(trueBlock value)
		}.
	`
	expected := []Token{
		COMMENT,
		UPPER_IDENT,
		LOWER_KEYWORD,
		STRING,
		UPPER_KEYWORD,
		LEFT_BRACE,
		BINARY,
		LOWER_IDENT,
		PERIOD,
		LOWER_IDENT,
		PERIOD,
		LOWER_IDENT,
		BINARY,
		LEFT_PAREN,
		LOWER_IDENT,
		LOWER_IDENT,
		RIGHT_PAREN,
		RIGHT_BRACE,
		PERIOD,
	}
	testTokens(t, received, expected)
}

func initScanner(expr string) Scanner {
	src := []byte(expr)

	fset := token.NewFileSet()
	file := fset.AddFile("", fset.Base(), len(src))

	var s Scanner
	s.Init(file, src, nil)

	return s
}

func testTokens(t *testing.T, expr string, tokens []Token) {
	s := initScanner(expr)
	msg := "Expected (%s) -- Received (%s)\n"

	for _, token := range tokens {
		_, tok, _ := s.Scan()
		if tok != token {
			t.Fatalf(msg, token, tok)
		}
	}
}
