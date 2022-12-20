package vic

import (
	"testing"
)

func TestRuntimeLoadsEvaluatesCaches(t *testing.T) {
	type assert struct {
		id  string
		val *Value
	}
	type line struct {
		stmt    string
		asserts []assert
	}
	type testCase struct {
		lines []line
	}
	assertNum := func(id string, num float64) assert {
		return assert{id: id, val: &Value{typ: NUMBER, num: num}}
	}
	assertBool := func(id string, b bool) assert {
		return assert{id: id, val: &Value{typ: BOOL, b: b}}
	}
	assertEmpty := func(id string) assert {
		return assert{id: id}
	}
	testCases := []testCase{
		{ // basic statement and expression tests
			lines: []line{
				line{stmt: "c=1",
					asserts: []assert{
						assertNum("c", 1)}},
				line{stmt: "d=c+1",
					asserts: []assert{
						assertNum("c", 1),
						assertNum("d", 2)}},
				line{stmt: "a=3",
					asserts: []assert{
						assertNum("a", 3),
						assertNum("c", 1),
						assertNum("d", 2)}},
				line{stmt: "b=5",
					asserts: []assert{
						assertNum("a", 3),
						assertNum("b", 5),
						assertNum("c", 1),
						assertNum("d", 2)}},
				line{stmt: "c=a*b",
					asserts: []assert{
						assertNum("a", 3),
						assertNum("b", 5),
						assertNum("c", 15),
						assertNum("d", 16)}},
				line{stmt: "a=b*2",
					asserts: []assert{
						assertNum("a", 10),
						assertNum("b", 5),
						assertNum("c", 50),
						assertNum("d", 51)}},
			},
		}, { // basic contextual evaluation tests
			lines: []line{
				line{stmt: "a=(b+1)[b=1]",
					asserts: []assert{
						assertNum("a", 2),
						assertEmpty("b"),
					}},
				line{stmt: "a=b[b=1]+1",
					asserts: []assert{
						assertNum("a", 2),
						assertEmpty("b"),
					}},
				line{stmt: "square=(n*n)[n=11]",
					asserts: []assert{
						assertNum("square", 121),
						assertEmpty("n"),
					}},
				line{stmt: "mult=(n*m)[n=10][m=16]-1",
					asserts: []assert{
						assertNum("mult", 159),
						assertEmpty("n"),
						assertEmpty("m"),
					}},
				line{stmt: "n=1[unrelated=11]",
					asserts: []assert{
						assertNum("n", 1),
						assertEmpty("unrelated"),
					}},
				line{stmt: "p=5", asserts: []assert{}},
				line{stmt: "pp=p*p", asserts: []assert{}},
				line{stmt: "q=1+pp[p=4]",
					asserts: []assert{
						assertNum("p", 5),
						assertNum("pp", 25),
						assertNum("q", 17),
					}},
				line{stmt: "q=",
					asserts: []assert{
						assertEmpty("q"),
					}},
			},
		}, { // basic builtin function tests
			lines: []line{
				line{stmt: "a=IF(TRUE,5,10)",
					asserts: []assert{
						assertNum("a", 5),
					}},
				line{stmt: "even11=IF( EQ(n%2,0), TRUE, FALSE)[n=11]",
					asserts: []assert{
						assertBool("even11", false),
					}},
			},
		},
	}

	for i, testCase := range testCases {
		r := newRuntimeTest()
		for j, line := range testCase.lines {
			err := r.Load(line.stmt)
			if err != nil {
				t.Fatalf("test case %v, line %v: expected no error, got %v", i, j, err)
			}
			for _, assert := range line.asserts {
				id, ok := r.ids[assert.id]
				if !ok && assert.val != nil {
					t.Fatalf("test case %v, line %v: expected runtime to create id %v, got nil",
						i, j, assert.id)
				}
				if ok && assert.val == nil {
					t.Fatalf("test case %v, line %v: expected runtime to have no id %v, got %v",
						i, j, assert.id, id)
				}
				if ok && id.val == nil {
					t.Fatalf("test case %v, line %v: expected %v to have non-nil value, got nil",
						i, j, id.lit)
				}
				if !ok {
					continue
				}
				if id.val.typ != assert.val.typ {
					t.Errorf("test case %v, line %v: expected %v to have type %v, got type %v",
						i, j, id.lit, assert.val.typ, id.val.typ)
				}
				switch assert.val.typ {
				case NUMBER:
					if id.val.num != assert.val.num {
						t.Errorf("test case %v, line %v: expected %v to have num %v, got num %v",
							i, j, id.lit, assert.val.num, id.val.num)
					}
				case BOOL:
					if id.val.b != assert.val.b {
						t.Errorf("test case %v, line %v: expected %v to have bool %v, got bool %v",
							i, j, id.lit, assert.val.b, id.val.b)
					}
				default:
					t.Errorf("test case %v, line %v: unhandled type %v", i, j, assert.val.typ)
				}
			}
		}
	}
}
