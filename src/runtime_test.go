package vic

import (
	"testing"
)

func TestRuntimeLoadsEvaluatesCaches(t *testing.T) {
	r := newRuntimeTest()

	assertValNotNil := func(id *Identifier) {
		if id.val == nil {
			t.Fatalf("expected %v to have non-nil value, got nil", id.lit)
		}
	}
	assertType := func(id *Identifier, typ valType) {
		if id.val.typ != typ {
			t.Errorf("expected %v to have type %v, got type %v", id.lit, typ, id.val.typ)
		}
	}
	assertNum := func(id *Identifier, num float64) {
		if id.val.num != num {
			t.Errorf("expected %v to have num %v, got num %v", id.lit, num, id.val.num)
		}
	}

	err := r.Load("c=1")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	c, ok := r.ids["c"]
	if !ok {
		t.Fatal("expected runtime to create id c, got nil")
	}
	assertValNotNil(c)
	assertType(c, NUMBER)
	assertNum(c, 1)

	err = r.Load("d=c+1")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	d, ok := r.ids["d"]
	if !ok {
		t.Fatal("expected runtime to create id d, got nil")
	}
	assertValNotNil(d)
	assertType(d, NUMBER)
	assertNum(d, 2)

	err = r.Load("a=5")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	a, ok := r.ids["a"]
	if !ok {
		t.Fatal("expected runtime to create id a, got nil")
	}
	assertValNotNil(a)
	assertType(a, NUMBER)
	assertNum(a, 5)

	err = r.Load("b=3")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	b, ok := r.ids["b"]
	if !ok {
		t.Fatal("expected runtime to create id b, got nil")
	}
	assertValNotNil(b)
	assertType(b, NUMBER)
	assertNum(b, 3)

	r.Load("c=a*b")
	// silently assert that same ID object is used in runtime
	assertValNotNil(c)
	assertType(c, NUMBER)
	assertNum(c, 15)
	assertNum(d, 16)

	r.Load("a=b*2")
	// silently assert that same ID object is used in runtime
	assertValNotNil(a)
	assertType(a, NUMBER)
	assertNum(a, 6)
	assertNum(c, 18)
	assertNum(d, 19)
}
