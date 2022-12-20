package vic

import (
	"fmt"
)

// Value type with associated metadata and attribute implementations
type valType int

const (
	UNKNOWN valType = iota
	NULL
	NUMBER
	BOOL
	CTXV
)
type Value struct {
	typ valType
	num float64
	b   bool
	ctx *Expression
	lit string
}
func NewNullValue(lit string) *Value {
	return &Value{typ: NULL, lit: lit}
}
func NewNumberValue(num float64, lit string) *Value {
	return &Value{typ: NUMBER, num: num, lit: lit}
}
func NewBoolValue(b bool, lit string) *Value {
	return &Value{typ: BOOL, b: b, lit: lit}
}
func NewCtxValue(e *Expression) (*Value, error) {
	if e.typ != CTX {
		return nil, fmt.Errorf("Expected CTX expression but got %v", e.typ)
	}
	return &Value{typ: CTXV, ctx: e, lit: e.lit}, nil
}
