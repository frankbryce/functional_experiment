package main
import (
    "bufio"
    "bytes"
    "fmt"
    "io"
    "os"
    "strings"
)

type Token int

const (
    // Special tokens
    ILLEGAL Token = iota
    EOF
    WS

    // Symbols
    SLASH
    DASH
    PLUS
    ASTERISK
    LPAREN
    RPAREN
    PERCENT
    EQUALS
    CARET

    // Literals
    IDENT  // identifiers

    // Keywords
    // TODO
)
func (t Token) String() string {
    switch t {
    case EOF:
        return "EOF"
    case WS:
        return "WS"
    case SLASH:
        return "SLASH"
    case DASH:
        return "DASH"
    case PLUS:
        return "PLUS"
    case ASTERISK:
        return "ASTERISK"
    case LPAREN:
        return "LPAREN"
    case RPAREN:
        return "RPAREN"
    case PERCENT:
        return "PERCENT"
    case EQUALS:
        return "EQUALS"
    case CARET:
        return "CARET"
    case IDENT:
        return "IDENT"
    }
    return "ILLEGAL"
}

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

// Scanner
type Scanner struct {
    r *bufio.Reader
}
func NewScanner(r io.Reader) *Scanner {
    return &Scanner{r: bufio.NewReader(r)}
}
func (s *Scanner) read() rune {
    ch, _, err := s.r.ReadRune()
    if err != nil {
        return eof
    }
    return ch
}
func (s *Scanner) unread() { _ = s.r.UnreadRune() }
func (s *Scanner) Scan() (tok Token, lit string) {
    ch := s.read()
    if isWhitespace(ch) {
        s.unread()
        return s.scanWhitespace()
    } else if isLetter(ch) {
        s.unread()
        return s.scanIdent()
    }

    switch ch {
    case eof:
        return EOF, string(ch)
    case '/':
        return SLASH, string(ch)
    case '-':
        return DASH, string(ch)
    case '+':
        return PLUS, string(ch)
    case '*':
        return ASTERISK, string(ch)
    case '(':
        return LPAREN, string(ch)
    case ')':
        return RPAREN, string(ch)
    case '%':
        return PERCENT, string(ch)
    case '=':
        return EQUALS, string(ch)
    case '^':
        return CARET, string(ch)
    }

    return ILLEGAL, string(ch)
}
func (s *Scanner) scanWhitespace() (tok Token, lit string) {
    var buf bytes.Buffer
    buf.WriteRune(s.read())
    for {
        if ch := s.read(); ch == eof {
            break;
        } else if !isWhitespace(ch) {
            s.unread()
            break
        } else {
            buf.WriteRune(ch)
        }
    }
    return WS, buf.String()
}
func (s *Scanner) scanIdent() (tok Token, lit string) {
    var buf bytes.Buffer
    buf.WriteRune(s.read())
    for {
        if ch := s.read(); ch == eof {
            break
        } else if !isLetter(ch) && !isDigit(ch) && ch != '_' {
            s.unread()
            break
        } else {
            _, _ = buf.WriteRune(ch)
        }
    }

    // TODO: handle keywords

    return IDENT, buf.String()
}
func (s *Scanner) scanFullIgnoreWS() (toks []Token, lits []string) {
    toks = []Token{}
    lits = []string{}
    for tok, lit := s.Scan(); tok != EOF; tok, lit = s.Scan() {
        if tok == WS { continue }
        toks = append(toks, tok)
        lits = append(lits, lit)
    }
    return
}

// Parser structs
// This Parser will only take statements of the form
//   lhs = rhs
// where rhs is an expression of (possibly nested) binary or unary operations on IDENT Tokens
// Atom Expression
//   IDENT
// valid Recursive Expressions
//   LPAREN Expression RPAREN
//   DASH Expression
//   Expression PLUS Expression
//   Expression ASTERISK Expression
//   Expression DASH Expression
//   Expression SLASH Expression
//   Expression PERCENT Expression
//   Expression CARET Expression
type ExpressionType int
const (
    // Expression Types
    NEGATE ExpressionType = iota
    PAREN
    MOD
    POWER
    MULTIPLY
    DIVIDE
    ADD
    SUBTRACT

    // Just an IDENT Token
    ATOM
)
type Expression struct {
    typ ExpressionType
    ex1 *Expression
    ex2 *Expression
    lit string
}
type Statement struct {
    lhs string  // must be IDENT Token
    // must be separated by EQUALS Token
    rhs Expression
}
func NewExpression(toks []Token, lits []string) (*Expression, error) {
    if len(toks) != len(lits) {
        return nil, fmt.Errorf("expected same length for tokens and literals")
    }
    if len(toks) == 0 { return nil, fmt.Errorf("no tokens given to NewExpression") }
    if len(toks) == 1 {
        if toks[0] == IDENT {
            return &Expression {
                typ: ATOM,
                lit: lits[0],
            }, nil
        }
        return nil, fmt.Errorf("Expected expression with one token to be an IDENT token")
    }
    if len(toks) == 2 {
        // TODO
    }
    return nil, nil
}

func main() {
    s := NewScanner(strings.NewReader(os.Args[1]))
    fmt.Println(s.scanFullIgnoreWS())
}
