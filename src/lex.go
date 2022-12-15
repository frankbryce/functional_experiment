package vic

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"

	"unicode/utf8"
)

// Some helpers
func isWhitespace(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n'
}
func isLetter(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}
func isDigit(ch rune) bool {
	return ch >= '0' && ch <= '9'
}

var eof = rune(0)

// Lexer
type Lexer struct {
	r *bufio.Reader
}

func NewLexer(r io.Reader) *Lexer {
	return &Lexer{r: bufio.NewReader(r)}
}
func (l *Lexer) read() rune {
	ch, _, err := l.r.ReadRune()
	if err != nil {
		return eof
	}
	return ch
}
func (l *Lexer) Error(s string) {
	fmt.Printf("Error: %s", s)
}
func (l *Lexer) unread() { _ = l.r.UnreadRune() }
func (l *Lexer) Lex(out *yySymType) int {
	l.skipWhitespace()
    ch := l.read()
	out.Literal = string(ch)
    if isLetter(ch) {
		l.unread()
		return l.scanIdent(out)
	}

	switch ch {
	case eof:
		return -1  // EOF
	case '/':
		return TSLASH
	case '-':
		return TDASH
	case '+':
		return TPLUS
	case '*':
		return TASTERISK
	case '(':
		return TLPAREN
	case ')':
		return TRPAREN
	case '%':
		return TPERCENT
	case '=':
		return TEQUALS
	case '^':
		return TCARET
	case '.':
		return TDOT
	}

	return ILLEGAL
}
func (l *Lexer) skipWhitespace() {
	for {
		if ch := l.read(); ch == eof {
			break
		} else if !isWhitespace(ch) {
			l.unread()
			break
		}
	}
}
func (l *Lexer) scanIdent(out *yySymType) int {
	var buf bytes.Buffer
	buf.WriteRune(l.read())
	for {
		if ch := l.read(); ch == eof {
			break
		} else if !isLetter(ch) && !isDigit(ch) && ch != '_' {
			l.unread()
			break
		} else {
			_, _ = buf.WriteRune(ch)
		}
	}

	out.Literal = buf.String()
	if len(buf.String()) == 0 {
		return TSTRING
	}
	str := strings.ToLower(buf.String())
	r, _ := utf8.DecodeRuneInString(str)
	switch {
	case str == "true":
		return TTRUE
	case str == "false":
		return TFALSE
	case str == "null":
		return TNULL
	case isLetter(r) || r == '_':
		return TSTRING
	case isDigit(r):
		return TNUMBER
	}

	return ILLEGAL
}

type IdentifierType int

const (
	RAW IdentifierType = iota
	DOT
	BRACK
)

type Identifier struct {
    typ   IdentifierType
	lit   string
	root  *Identifier
	dot   *Identifier
	brack *Identifier
}
type ExpressionType int

const (
	// Expression Types
	NEG ExpressionType = iota
	PAREN
	MOD
	POW
	MULT
	DIV
	PLUS
	MINUS

	ID
	VAL
)

type Expression struct {
	typ ExpressionType
	id  *Identifier
	e   []*Expression
	lit string
}
type Statement struct {
	id *Identifier
	ex *Expression
}

