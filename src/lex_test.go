package vic

import (
	"strings"
	"testing"
)

func TestBasicStatementsParseCorrectly(t *testing.T) {
	ins := []string{
		"A1=B1*C1",
		"test=(p+q)*(r+s)",
		"what.if.I.use.dots=(some.obj*(nested_param)+other.thing)",
		"a=-b--c+-d*-p--q",
		"c=1",
	}
	for _, inp := range ins {
		l := NewLexer(strings.NewReader(inp))
		stmt, e := Parse(l)
		if e != nil {
			t.Errorf("expected err to be nil, got %v", e)
		} else if stmt.lit != inp {
			t.Errorf("expected parsed lit to be %v, got %v", inp, stmt.lit)
		}
	}
}

func TestIncorrectStatementsReturnErrors(t *testing.T) {
	ins := []string{
		"-negated_identifier=something_else+a_thing",
		"expression_cannot*go_here=identifier.with.dot",
		"Things_must_have_equals_signs",
	}
	for _, inp := range ins {
		l := NewLexer(strings.NewReader(inp))
		_, e := Parse(l)
		if e == nil {
			t.Error("expected an error, but got nil")
		}
	}
}
