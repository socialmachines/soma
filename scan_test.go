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
		TOK_UPPER_IDENT,
		TOK_DECLARE,
		TOK_UPPER_IDENT,
		TOK_LOWER_IDENT,
		TOK_PERIOD,
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
		TOK_COMMENT,
		TOK_UPPER_IDENT,
		TOK_LOWER_KEYWORD,
		TOK_STRING,
		TOK_UPPER_KEYWORD,
		TOK_LEFT_BRACE,
		TOK_BINARY,
		TOK_LOWER_IDENT,
		TOK_BINARY,
		TOK_UPPER_IDENT,
		TOK_RIGHT_BRACE,
		TOK_PERIOD,
	}
	testTokens(t, received, expected)
}

func TestUnaryDefine(t *testing.T) {
	received := `
		+ (t True) not -> False.
	`
	expected := []Token{
		TOK_BINARY,
		TOK_LEFT_PAREN,
		TOK_LOWER_IDENT,
		TOK_UPPER_IDENT,
		TOK_RIGHT_PAREN,
		TOK_LOWER_IDENT,
		TOK_DEFINE,
		TOK_UPPER_IDENT,
		TOK_PERIOD,
	}
	testTokens(t, received, expected)
}

func TestKeyword(t *testing.T) {
	received := `
		'True ifTrue: { "do something" } Else: { "do something else" }'
		True externalBehavior: "ifTrue:Else:" Does: {
			|trueBlock. elseBlock. t|
			trueBlock value
		}.
	`
	expected := []Token{
		TOK_COMMENT,
		TOK_UPPER_IDENT,
		TOK_LOWER_KEYWORD,
		TOK_STRING,
		TOK_UPPER_KEYWORD,
		TOK_LEFT_BRACE,
		TOK_BINARY,
		TOK_LOWER_IDENT,
		TOK_PERIOD,
		TOK_LOWER_IDENT,
		TOK_PERIOD,
		TOK_LOWER_IDENT,
		TOK_BINARY,
		TOK_LOWER_IDENT,
		TOK_LOWER_IDENT,
		TOK_RIGHT_BRACE,
		TOK_PERIOD,
	}
	testTokens(t, received, expected)
}

func TestKeywordDefine(t *testing.T) {
	received := `
		+ (t True) ifTrue: tBlock Else: fBlock -> tBlock value.
	`
	expected := []Token{
		TOK_BINARY,
		TOK_LEFT_PAREN,
		TOK_LOWER_IDENT,
		TOK_UPPER_IDENT,
		TOK_RIGHT_PAREN,
		TOK_LOWER_KEYWORD,
		TOK_LOWER_IDENT,
		TOK_UPPER_KEYWORD,
		TOK_LOWER_IDENT,
		TOK_DEFINE,
		TOK_LOWER_IDENT,
		TOK_LOWER_IDENT,
		TOK_PERIOD,
	}
	testTokens(t, received, expected)
}

func TestAttributeGetter(t *testing.T) {
	received := `
		+ (f Foo) getHello -> f @hello.
	`
	expected := []Token{
		TOK_BINARY,
		TOK_LEFT_PAREN,
		TOK_LOWER_IDENT,
		TOK_UPPER_IDENT,
		TOK_RIGHT_PAREN,
		TOK_LOWER_IDENT,
		TOK_DEFINE,
		TOK_LOWER_IDENT,
		TOK_ATTR_GET,
		TOK_PERIOD,
	}
	testTokens(t, received, expected)
}

func TestAttributeSetter(t *testing.T) {
	received := `
	+ (f Foo) setHello: hello -> f @hello: hello.
`
	expected := []Token{
		TOK_BINARY,
		TOK_LEFT_PAREN,
		TOK_LOWER_IDENT,
		TOK_UPPER_IDENT,
		TOK_RIGHT_PAREN,
		TOK_LOWER_KEYWORD,
		TOK_LOWER_IDENT,
		TOK_DEFINE,
		TOK_LOWER_IDENT,
		TOK_ATTR_SET,
		TOK_LOWER_IDENT,
		TOK_PERIOD,
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
