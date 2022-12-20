package vic

// Identifier type with associated metadata and attribute implementations
type IdentifierType int

const (
	RAW IdentifierType = iota
	FUNC
)

type Identifier struct {
	typ IdentifierType
	lit string
    // TODO: fun *Function
	val *Value
}

