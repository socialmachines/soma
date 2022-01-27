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
	"fmt"
	"go/token"
	"path/filepath"
	"unicode"
	"unicode/utf8"
)

// An ErrorHandler may be provided to Scanner.Init. If a syntax error is
// encountered and a handler was installed, the handler is called with a
// position and an error message. The position points to the beginning of
// the offending token.
//
type ErrorHandler func(pos token.Position, msg string)

// A Scanner holds the scanner's internal state while processing
// a given text. It can be allocated as part of another data
// structure but must be initialized via Init before use.
//
type Scanner struct {
	// immutable state
	file *token.File  // source file handle
	dir  string       // directory portion of file.Name()
	src  []byte       // source
	err  ErrorHandler // error reporting; or nil

	// scanning state
	ch         rune // current character
	offset     int  // character offset
	rdOffset   int  // reading offset (position after current character)
	lineOffset int  // current line offset

	// public state - ok to modify
	ErrorCount int // number of errors encountered
}

const (
	bom = 0xFEFF // byte order mark, only permitted as very first character
	eof = -1     // end of file
)

// Init prepares the scanner s to tokenize the text src by setting the
// scanner at the beginning of src. The scanner uses the file set file
// for position information and it adds line information for each line.
// It is ok to re-use the same file when re-scanning the same file as
// line information which is already present is ignored. Init causes a
// panic if the file size does not match the src size.
//
// Calls to Scan will invoke the error handler err if they encounter a
// syntax error and err is not nil. Also, for each error encountered,
// the Scanner field ErrorCount is incremented by one. The mode parameter
// determines how comments are handled.
//
// Note that Init may call err if there is an error in the first character
// of the file.
//
func (s *Scanner) Init(file *token.File, src []byte, err ErrorHandler) {
	if file.Size() != len(src) {
		panic(fmt.Sprintf("file size (%d) does not match src len (%d)", file.Size(), len(src)))
	}

	s.file = file
	s.dir, _ = filepath.Split(file.Name())
	s.src = src
	s.err = err

	s.ch = ' '
	s.offset = 0
	s.rdOffset = 0
	s.lineOffset = 0
	s.ErrorCount = 0

	s.next()
	if s.ch == bom {
		s.next() // ignore BOM at file beginning
	}
}

func (s *Scanner) error(offs int, msg string) {
	if s.err != nil {
		s.err(s.file.Position(s.file.Pos(offs)), msg)
	}
	s.ErrorCount++
}

func (s *Scanner) errorf(offs int, format string, args ...interface{}) {
	s.error(offs, fmt.Sprintf(format, args...))
}

func (s *Scanner) Scan() (pos token.Pos, tok Token, lit string) {
	s.skipWhiteSpace()

	pos = s.file.Pos(s.offset)

	switch ch := s.ch; {
	case unicode.IsUpper(ch):
		lit = s.scanIdentifier()
		if s.ch == ':' {
			tok, lit = UPPER_KEYWORD, lit+":"
		} else {
			tok = UPPER_IDENT
		}
	case unicode.IsLower(ch):
		lit = s.scanIdentifier()
		if s.ch == ':' {
			tok, lit = LOWER_KEYWORD, lit+":"
		} else {
			tok = LOWER_IDENT
		}
	case isBinary(ch):
		tok, lit = BINARY, s.scanBinary()
	default:
		s.next()
		switch ch {
		case -1:
			tok, lit = EOF, "EOF"
		case '\'':
			tok, lit = COMMENT, s.scanComment()
		case '"':
			tok, lit = STRING, s.scanString()
		case ':':
			if s.ch == '=' {
				s.next()
				tok, lit = DECLARE, ":="
			}
		case '[':
			tok, lit = LEFT_BRACK, "["
		case ']':
			tok, lit = RIGHT_BRACK, "]"
		case ';':
			tok, lit = FLUENT, ";"
		case '.':
			tok, lit = PERIOD, "."
		default:
			s.errorf(s.file.Offset(pos), "illegal character %#U", ch)
			tok, lit = ILLEGAL, string(ch)
		}
	}
	return
}

func (s *Scanner) scanBinary() string {
	offs := s.offset
	for isBinary(s.ch) {
		s.next()
	}
	return string(s.src[offs:s.offset])
}

func (s *Scanner) scanComment() string {
	offs := s.offset - 1
	for s.ch != '\'' && s.ch != -1 {
		s.next()
	}
	if s.ch != '\'' {
		s.error(s.offset, "expecting single-quote (') to end the comment")
	}
	s.next()
	return string(s.src[offs:s.offset])
}

// scanIdentifier reads the string of valid identifier characters at s.offset.
// It must only be called when s.ch is known to be a valid letter.
//
// Be careful when making changes to this function: it is optimized and affects
// scanning performance significantly.
func (s *Scanner) scanIdentifier() string {
	offs := s.offset

	// Optimize for the common case of an ASCII identifier.
	//
	// Ranging over s.src[s.rdOffset:] lets us avoid some bounds checks, and
	// avoids conversions to runes.
	//
	// In case we encounter a non-ASCII character, fall back on the slower path
	// of calling into s.next().
	for rdOffset, b := range s.src[s.rdOffset:] {
		if 'a' <= b && b <= 'z' || 'A' <= b && b <= 'Z' || b == '_' || '0' <= b && b <= '9' {
			// Avoid assigning a rune for the common case of an ascii character.
			continue
		}
		s.rdOffset += rdOffset
		if 0 < b && b < utf8.RuneSelf {
			// Optimization: we've encountered an ASCII character that's not a letter
			// or number. Avoid the call into s.next() and corresponding set up.
			//
			// Note that s.next() does some line accounting if s.ch is '\n', so this
			// shortcut is only possible because we know that the preceding character
			// is not '\n'.
			s.ch = rune(b)
			s.offset = s.rdOffset
			s.rdOffset++
			goto exit
		}
		// We know that the preceding character is valid for an identifier because
		// scanIdentifier is only called when s.ch is a letter, so calling s.next()
		// at s.rdOffset resets the scanner state.
		s.next()
		for isLetter(s.ch) || isDigit(s.ch) {
			s.next()
		}
		goto exit
	}
	s.offset = len(s.src)
	s.rdOffset = len(s.src)
	s.ch = eof

exit:
	return string(s.src[offs:s.offset])
}

func (s *Scanner) scanString() string {
	offs := s.offset - 1
	for s.ch != '"' && s.ch != -1 {
		s.next()
	}
	if s.ch != '"' {
		s.error(s.offset, "expecting double-quote (\") to end the string")
	}
	s.next()
	return string(s.src[offs:s.offset])
}

// Read the next Unicode char into s.ch. s.ch < 0 means end-of-file.
//
func (s *Scanner) next() {
	if s.rdOffset < len(s.src) {
		s.offset = s.rdOffset
		if s.ch == '\n' {
			s.lineOffset = s.offset
			s.file.AddLine(s.offset)
		}

		r, w := rune(s.src[s.rdOffset]), 1
		switch {
		case r == 0:
			s.error(s.offset, "illegal character NUL")
		case r >= utf8.RuneSelf:
			// not ASCII
			r, w = utf8.DecodeRune(s.src[s.rdOffset:])
			if r == utf8.RuneError && w == 1 {
				s.error(s.offset, "illegal UTF-8 encoding")
			}
		}
		s.rdOffset += w
		s.ch = r
	} else {
		s.offset = len(s.src)
		if s.ch == '\n' {
			s.lineOffset = s.offset
			s.file.AddLine(s.offset)
		}
		s.ch = eof
	}
}

func isBinary(ch rune) bool {
	switch ch {
	case '!', '*', '/', '+', '|', '&', '-', '>', '<', '=', '?', '\\', '~':
		return true
	}
	return false
}

func isDigit(ch rune) bool {
	isDecimal := func(ch rune) bool {
		return '0' <= ch && ch <= '9'
	}
	return isDecimal(ch) || ch >= utf8.RuneSelf && unicode.IsDigit(ch)
}

func isLetter(ch rune) bool {
	lower := func(ch rune) rune {
		return ('a' - 'A') | ch
	}
	return 'a' <= lower(ch) && lower(ch) <= 'z' || ch == '_' || ch >= utf8.RuneSelf && unicode.IsLetter(ch)
}

func (s *Scanner) skipWhiteSpace() {
	for unicode.IsSpace(s.ch) {
		s.next()
	}
}
