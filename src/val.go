package vic

// Value type with associated metadata and attribute implementations
type valType int

const (
	UNKNOWN valType = iota
	NULL
	NUMBER
	BOOL
)

type Value struct {
	typ valType
	num float64
	b   bool
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
