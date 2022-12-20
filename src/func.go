package vic

import (
	"fmt"
)

type Function interface {
	Call([]*Value) (*Value, error)
}

// built in functions
type IF struct{}

func (_ IF) Call(args []*Value) (*Value, error) {
	if len(args) != 3 {
		return nil, fmt.Errorf("Expected 3 args passed to IF(), got %v", args)
	}

	pred, tru, fal := args[0], args[1], args[2]
	if pred == nil {
		fmt.Errorf("Expected IF() predicate to be non-nil, got nil")
	}
	if tru == nil {
		fmt.Errorf("Expected IF() true branch expression to be non-nil, got nil")
	}
	if fal == nil {
		fmt.Errorf("Expected IF() false branch expression to be non-nil, got nil")
	}

	if pred.typ != BOOL {
		fmt.Errorf("Expected predicate to IF() to have type BOOL, got type %v", pred.typ)
	}
	if pred.b {
		return tru, nil
	} else {
		return fal, nil
	}
}

type EQ struct {
	v1 *Expression
	v2 *Expression
}

func (_ EQ) Call(args []*Value) (*Value, error) {
	if len(args) != 2 {
		return nil, fmt.Errorf("Expected 2 args passed to EQ(), got %v", args)
	}
	if args[0] == nil {
		fmt.Errorf("Expected first arg to EQ() to be non-nil, got nil")
	}
	if args[1] == nil {
		fmt.Errorf("Expected second arg to EQ() to be non-nil, got nil")
	}
	TRUE := &Value{typ: BOOL, b: true, lit: "TRUE"}
	FALSE := &Value{typ: BOOL, b: false, lit: "FALSE"}
	if args[0].typ != args[1].typ {
		return FALSE, nil
	}
	switch args[0].typ {
	case NUMBER:
		if args[0].num == args[1].num {
			return TRUE, nil
		}
		return FALSE, nil
	case BOOL:
		if args[0].b == args[1].b {
			return TRUE, nil
		}
		return FALSE, nil
	default:
		return nil, fmt.Errorf("Unexpected type evaluated in EQ(): %v", args[0].typ)
	}
}

// TODO: handle this. For now everything returns error.
type PLUGIN struct{}

func (_ PLUGIN) Call(_ []*Value) (*Value, error) {
	return nil, fmt.Errorf("Unimplemented: called a PLUGIN function")
}
