package vic

// Statement type with associated metadata and attribute implementations
type Statement struct {
	id  *Identifier
	ex  *Expression
	lit string
}
