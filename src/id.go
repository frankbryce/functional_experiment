package vic

// Identifier type with associated metadata and attribute implementations
type IdentifierType int

const (
	RAW IdentifierType = iota
)

type Identifier struct {
	typ IdentifierType
	lit string
	val *Value
}
