package vic

import (
    "fmt"
    "strings"
)

type Runtime interface {
    LoadStatement(stmt string) error
    Update(Display) error
}

type runtimeImpl struct {
}

func (r *runtimeImpl) LoadStatement(stmt_str string) error {
    l := NewLexer(strings.NewReader(stmt_str))
    stmt, e := Parse(l)
    if e != 0 {
        return fmt.Errorf("Parsing failed with error code %v", e)
    }
    fmt.Printf("Parsed statement %v", stmt.lit)
    return nil
}

func (r *runtimeImpl) Update(d Display) error {
    return nil
}

// Idea: GUID to not re-update values since last update?
type Display interface {
    SetStringValue(name string, value string) error
    SetIntegerValue(name string, value int64) error
    SetDecimalValue(name string, value float64) error
}

type defaultDisplay struct {}

func (d *defaultDisplay) SetStringValue(name string, value string) error {
    return nil
}
func (d *defaultDisplay) SetIntegerValue(name string, value int64) error {
    return nil
}
func (d *defaultDisplay) SetDecimalValue(name string, value float64) error {
    return nil
}
