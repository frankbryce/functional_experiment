package vic

import (
    "fmt"
    "strings"
    "testing"
)

func TestBasicStatement(t *testing.T) {
    l := NewLexer(strings.NewReader("A1=B1*C1"))
    e := yyParse(l)
    fmt.Println(e)
}
