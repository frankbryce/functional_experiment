package vic

import (
	"fmt"
	"math"
	"strings"
)

// Runtime type with associated metadata and attribute implementations
type Runtime interface {
	Load(stmt string) error
	Update(Display) error
}
type runtimeImpl struct {
	// cache of Identifiers that runtime knows about
	ids map[string]*Identifier
	// cache active Identifier stmts in runtime
	stmts map[string]*Statement
	// {Identifier lit:Statement} dependency map.
	id_st_deps map[string][]*Statement
}

func NewRuntime() Runtime {
	return newRuntimeTest()
}
func newRuntimeTest() *runtimeImpl {
	return &runtimeImpl{
		ids:        make(map[string]*Identifier),
		stmts:      make(map[string]*Statement),
		id_st_deps: make(map[string][]*Statement),
	}
}
func (r *runtimeImpl) Load(stmt_str string) error {
	l := NewLexer(strings.NewReader(stmt_str))
	stmt, err := Parse(l)
	if err != nil {
		return err
	}
	return r.execute(stmt)
}
func (r *runtimeImpl) execute(s *Statement) error {
	var id *Identifier
	var ok bool
	if s.ex == nil {
		id = s.id
		if _, ok = r.stmts[id.lit]; ok {
			delete(r.ids, id.lit)
			delete(r.stmts, id.lit)
		}
	} else {
		// evaluate expression
		val, err := r.Evaluate(s.ex)
		if err != nil {
			return err
		}

		// cache evaluation & id in runtime
		if id, ok = r.ids[s.id.lit]; ok {
			id.val = val
		} else {
			id = s.id
			id.val = val
			r.ids[id.lit] = id
		}

		// cache active Statement for id
		r.stmts[id.lit] = s

		// cache id_st_deps in runtime
		for idlit, _ := range s.ex.ids {
			if _, ok := r.id_st_deps[idlit]; !ok {
				r.id_st_deps[idlit] = []*Statement{s}
			} else {
				r.id_st_deps[idlit] = append(r.id_st_deps[idlit], s)
			}
		}
	}

	// remove old statements from dep list and
	// update dependent expressioned based on id evaluation
	i := 0
	for _, dep_s := range r.id_st_deps[id.lit] {
		// only keep if it's the active statement for the id
		if r.stmts[dep_s.id.lit] == dep_s {
			r.id_st_deps[id.lit][i] = dep_s
			i += 1
			r.execute(dep_s)
		}
	}
	// ... and avoid memory leaks
	for j := i; j < len(r.id_st_deps[id.lit]); j++ {
		r.id_st_deps[id.lit][j] = nil
	}
	r.id_st_deps[id.lit] = r.id_st_deps[id.lit][:i]
	return nil
}
func (r *runtimeImpl) Evaluate(e *Expression) (*Value, error) {
	evalTwoInputs := func(op func(*Value, *Value) (*Value, error)) (*Value, error) {
		v1, err := r.Evaluate(e.e[0])
		if err != nil {
			return nil, err
		}
		v2, err := r.Evaluate(e.e[1])
		if err != nil {
			return nil, err
		}
		return op(v1, v2)
	}

	switch e.typ {
	case ID:
		if id, ok := r.ids[e.id.lit]; ok {
			return id.val, nil
		}
		return nil, fmt.Errorf("Identifier doesn't exist: %v", e.id.lit)
	case VAL:
		return e.val, nil
	case CTX:
		var restore_stmt *Statement
		var ok bool
		if restore_stmt, ok = r.stmts[e.ctx.id.lit]; !ok {
			restore_stmt = &Statement{id: e.ctx.id, lit: e.ctx.id.lit + "="}
		}
		r.execute(e.ctx)
		v, err := r.Evaluate(e.e[0])
		if err != nil {
			return nil, err
		}
		r.execute(restore_stmt)
		return v, nil
	case ARGS:
		return nil, fmt.Errorf("Unexpected evaluation of an ARGS expression: %v.", e.lit)
	case CALL:
		args := make([]*Value, len(e.e))
		for i, arg := range e.e {
			argv, err := r.Evaluate(arg)
			if err != nil {
				return nil, err
			}
			args[i] = argv
		}
		fid := strings.ToUpper(e.id.lit)
		switch fid {
		case "IF":
			return IF{}.Call(args)
		case "EQ":
			return EQ{}.Call(args)
		default:
			return PLUGIN{}.Call(args)
		}
	case NEG:
		v, err := r.Evaluate(e.e[0])
		if err != nil {
			return nil, err
		}
		if v.typ == NUMBER {
			return NewNumberValue(-v.num, e.lit), nil
		} else if v.typ == BOOL {
			return NewBoolValue(!v.b, e.lit), nil
		}
		return nil, fmt.Errorf("Invalid Value type for NEG expression: %v", v.typ)
	case PAREN:
		v, err := r.Evaluate(e.e[0])
		if err != nil {
			return nil, err
		}
		if v.typ == NUMBER {
			return NewNumberValue(v.num, e.lit), nil
		} else if v.typ == BOOL {
			return NewBoolValue(v.b, e.lit), nil
		}
		return nil, fmt.Errorf("Invalid Value type for PAREN expression: %v", v.typ)
	case MOD:
		op := func(v1, v2 *Value) (*Value, error) {
			if v1.typ == NUMBER && v2.typ == NUMBER {
				_, ans := math.Modf(v1.num / v2.num)
				return NewNumberValue(ans, e.lit), nil
			}
			return nil, fmt.Errorf("Incompatible types for MOD expression: %v and %v", v1.typ, v2.typ)
		}
		return evalTwoInputs(op)
	case POW:
		op := func(v1, v2 *Value) (*Value, error) {
			if v1.typ == NUMBER && v2.typ == NUMBER {
				return NewNumberValue(math.Pow(v1.num, v2.num), e.lit), nil
			}
			return nil, fmt.Errorf("Incompatible types for POW expression: %v and %v", v1.typ, v2.typ)
		}
		return evalTwoInputs(op)
	case MULT:
		op := func(v1, v2 *Value) (*Value, error) {
			if v1.typ == NUMBER && v2.typ == NUMBER {
				return NewNumberValue(v1.num*v2.num, e.lit), nil
			} else if v1.typ == BOOL && v2.typ == BOOL {
				return NewBoolValue(v1.b && v2.b, e.lit), nil
			}
			return nil, fmt.Errorf("Incompatible types for MULT expression: %v and %v", v1.typ, v2.typ)
		}
		return evalTwoInputs(op)
	case DIV:
		op := func(v1, v2 *Value) (*Value, error) {
			if v1.typ == NUMBER && v2.typ == NUMBER {
				return NewNumberValue(v1.num/v2.num, e.lit), nil
			}
			return nil, fmt.Errorf("Incompatible types for DIV expression: %v and %v", v1.typ, v2.typ)
		}
		return evalTwoInputs(op)
	case PLUS:
		op := func(v1, v2 *Value) (*Value, error) {
			if v1.typ == NUMBER && v2.typ == NUMBER {
				return NewNumberValue(v1.num+v2.num, e.lit), nil
			} else if v1.typ == BOOL && v2.typ == BOOL {
				return NewBoolValue(v1.b || v2.b, e.lit), nil
			}
			return nil, fmt.Errorf("Incompatible types for PLUS expression: %v and %v", v1.typ, v2.typ)
		}
		return evalTwoInputs(op)
	case MINUS:
		op := func(v1, v2 *Value) (*Value, error) {
			if v1.typ == NUMBER && v2.typ == NUMBER {
				return NewNumberValue(v1.num-v2.num, e.lit), nil
			}
			return nil, fmt.Errorf("Incompatible types for DIV expression: %v and %v", v1.typ, v2.typ)
		}
		return evalTwoInputs(op)
	default:
		return nil, fmt.Errorf("Did not recognize Expression Type %v", e.typ)
	}
}
func (r *runtimeImpl) Update(d Display) error {
	return nil
}
