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
	assertNum := func(id string, typ valType, num float64) assert {
		return assert{id: id, val: &Value{typ: typ, num: num}}
	}
	assertEmpty := func(id string) assert {
		return assert{id: id}
	}
	testCases := []testCase{
		{
			lines: []line{
				line{stmt: "c=1",
					asserts: []assert{assertNum("c", NUMBER, 1)}},
				line{stmt: "d=c+1",
					asserts: []assert{assertNum("c", NUMBER, 1),
						assertNum("d", NUMBER, 2)}},
				line{stmt: "a=3",
					asserts: []assert{assertNum("a", NUMBER, 3),
						assertNum("c", NUMBER, 1),
						assertNum("d", NUMBER, 2)}},
				line{stmt: "b=5",
					asserts: []assert{assertNum("a", NUMBER, 3),
						assertNum("b", NUMBER, 5),
						assertNum("c", NUMBER, 1),
						assertNum("d", NUMBER, 2)}},
				line{stmt: "c=a*b",
					asserts: []assert{assertNum("a", NUMBER, 3),
						assertNum("b", NUMBER, 5),
						assertNum("c", NUMBER, 15),
						assertNum("d", NUMBER, 16)}},
				line{stmt: "a=b*2",
					asserts: []assert{assertNum("a", NUMBER, 10),
						assertNum("b", NUMBER, 5),
						assertNum("c", NUMBER, 50),
						assertNum("d", NUMBER, 51)}},
			},
		}, {
			lines: []line{
				line{stmt: "a=(b+1)[b=1]",
					asserts: []assert{
						assertNum("a", NUMBER, 2),
						assertEmpty("b"),
					}},
				line{stmt: "a=b[b=1]+1",
					asserts: []assert{
						assertNum("a", NUMBER, 2),
						assertEmpty("b"),
					}},
				line{stmt: "square=(n*n)[n=11]",
					asserts: []assert{
						assertNum("square", NUMBER, 121),
						assertEmpty("n"),
					}},
				line{stmt: "mult=(n*m)[n=10][m=16]-1",
					asserts: []assert{
						assertNum("mult", NUMBER, 159),
						assertEmpty("n"),
						assertEmpty("m"),
					}},
				line{stmt: "n=1[unrelated=11]",
					asserts: []assert{
						assertNum("n", NUMBER, 1),
						assertEmpty("unrelated"),
					}},
			},
		}, {
		    lines: []line{
		        line{stmt: "n=5", asserts: []assert{}},
		        line{stmt: "nn=n*n", asserts: []assert{}},
		        line{stmt: "m=1+nn[n=4]",
		            asserts: []assert{
		                assertNum("n", NUMBER, 5),
		                assertNum("nn", NUMBER, 25),
		                assertNum("m", NUMBER, 17),
		            }},
                line{stmt: "m=",
                    asserts: []assert{
                        assertEmpty("m"),
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
				if assert.val.typ == NUMBER {
					if id.val.num != assert.val.num {
						t.Errorf("test case %v, line %v: expected %v to have num %v, got num %v",
							i, j, id.lit, assert.val.num, id.val.num)
					}
				} else {
					t.Errorf("test case %v, line %v: unhandled type %v", i, j, assert.val.typ)
				}
			}
		}
	}
}
