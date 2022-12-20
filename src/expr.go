package vic

// Expression type with associated metadata and attribute implementations
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

	CTX  // inline function support
    ARGS // external API function arg list
    CALL // external API function call
)

type Expression struct {
	typ ExpressionType
	id  *Identifier
	e   []*Expression
	val *Value
	ctx *Statement
	lit string

	// identifier strings that are in this expression
	// used for dependency tracking
	ids map[string]bool
}

