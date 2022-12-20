package vic

import (
	"fmt"
)

// Display type with associated metadata and attribute implementations
type Display interface {
	SetValue(val *Value) error
}
type defaultDisplay struct{}

func (d *defaultDisplay) SetValue(v Value) error {
	switch v.typ {
	case NUMBER:
		return d.setNumberValue(v.lit, v.num)
	case BOOL:
		return d.setBoolValue(v.lit, v.b)
	}
	return fmt.Errorf("Did no recognize Value type for %v: %v", v.lit, v.typ)
}
func (d *defaultDisplay) setNumberValue(name string, val float64) error {
	return nil
}
func (d *defaultDisplay) setBoolValue(name string, val bool) error {
	return nil
}

