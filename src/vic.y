// +build exclude
%{
package vic

import (
    "fmt"
    "strconv"
)

var stmt *Statement
func Parse(l *Lexer) (*Statement, error) {
    e := yyParse(l)
    if e != 0 {
        return nil, fmt.Errorf("Error code %v while Parsing %v", e, l)
    }
    return stmt, nil
}

func mergeIds(e1,e2 *Expression) map[string]bool {
    ids := make(map[string]bool)
    if e1 != nil {
        for id, _ := range(e1.ids) {
            ids[id] = true
        }
    }
    if e2 != nil {
        for id, _ := range(e2.ids) {
            ids[id] = true
        }
    }
    return ids
}

func subtractId(e *Expression, id *Identifier) map[string]bool {
    ids := make(map[string]bool)
    for id, _ := range(e.ids) {
        ids[id] = true
    }
    if _, ok := ids[id.lit]; ok {
        delete(ids, id.lit)
    }
    return ids
}
%}

%union {
    Literal    string
    Value      *Value
    Statement  *Statement
    Expression *Expression
    Identifier *Identifier
}

// math symbols
%token <Literal> TSLASH TDASH TPLUS TASTERISK TPERCENT TCARET
// other symbols
%token <Literal> TLPAREN TRPAREN TLBRACK TRBRACK TEQUALS TDOT
// other lex tokens for lexer
%token <Literal> ILLEGAL

// literals
%token <Literal> TSTRING TNUMBER
%token <Literal> TFALSE TTRUE TNULL

%left TPLUS TDASH
%left TASTERISK TSLASH
%left TAMPERSAND
%left NEGATE
%right TCARET
%right TDOT
%right TLBRACK

// grammar productions
%type <Value> value
%type <Statement> statement
%type <Expression> expression
%type <Identifier> identifier

%%

statement   : identifier TEQUALS { // deletes identifier from runtime
                stmt = &Statement{id:$1, ex:nil, lit:$1.lit+$2 }
                $$ = stmt
            }
            | identifier TEQUALS expression {
                stmt = &Statement{id:$1, ex:$3, lit:$1.lit+$2+$3.lit }
                $$ = stmt
            }
;
identifier  : TSTRING { 
                $$ = &Identifier{ typ:RAW, lit:$1 }
            }
;
value       : TNUMBER {
                f, err := strconv.ParseFloat($1, 64)
                if err != nil { panic(fmt.Errorf("Error during ParseFloat: %v", err)) }
                $$ = NewNumberValue(f, $1)
            }
            | TNUMBER TDOT TNUMBER {
                lit := $1+$2+$3
                f, err := strconv.ParseFloat(lit, 64)
                if err != nil { panic(fmt.Errorf("Error during ParseFloat: %v", err)) }
                $$ = NewNumberValue(f, lit)
            }
            | TFALSE {
                $$ = NewBoolValue(false, $1)
            }
            | TTRUE {
                $$ = NewBoolValue(true, $1)
            }
            | TNULL {
                $$ = NewNullValue($1)
            }
;
expression  : identifier {
                ids := make(map[string]bool)
                ids[$1.lit] = true
                $$ = &Expression{ typ:ID, id:$1, lit:$1.lit, ids:ids }
            }
            | value {
                $$ = &Expression{ typ:VAL, val:$1, lit:$1.lit, ids:make(map[string]bool) }
            }
            | expression TPLUS expression {
                $$ = &Expression{
                    typ:PLUS, e:[]*Expression{$1, $3}, lit:$1.lit+$2+$3.lit,
                    ids:mergeIds($1,$3),
                }
            }
            | expression TDASH expression {
                $$ = &Expression{
                    typ:MINUS, e:[]*Expression{$1, $3}, lit:$1.lit+$2+$3.lit,
                    ids:mergeIds($1,$3),
                }
            }
            | expression TASTERISK expression {
                $$ = &Expression{
                    typ:MULT, e:[]*Expression{$1, $3}, lit:$1.lit+$2+$3.lit,
                    ids:mergeIds($1,$3),
                }
            }
            | expression TSLASH expression {
                $$ = &Expression{
                    typ:DIV, e:[]*Expression{$1, $3}, lit:$1.lit+$2+$3.lit,
                    ids:mergeIds($1,$3),
                }
            }
            | TDASH expression %prec NEGATE {
                $$ = &Expression{
                    typ:NEG, e:[]*Expression{$2}, lit:$1+$2.lit,
                    ids:$2.ids,
                }
            }
            | expression TCARET expression {
                $$ = &Expression{
                    typ:POW, e:[]*Expression{$1, $3}, lit:$1.lit+$2+$3.lit,
                    ids:mergeIds($1,$3),
                }
            }
            | TLPAREN expression TRPAREN {
                $$ = &Expression{
                    typ:PAREN, e:[]*Expression{$2}, lit:$1+$2.lit+$3,
                    ids:$2.ids,
                }
            }
            | expression TLBRACK statement TRBRACK {
                ids := make(map[string]bool)
                if $1 != nil {
                    for id, _ := range($1.ids) {
                        if id == $3.id.lit { continue }
                        ids[id] = true
                    }
                }
                if $3.ex != nil {
                    for id, _ := range($3.ex.ids) {
                        ids[id] = true
                    }
                }
                $$ = &Expression {
                    typ:CTX,
                    e:[]*Expression{$1},
                    lit:$1.lit+$2+$3.lit+$4,
                    ctx:$3,
                    ids:ids,
                }
            }
;

%%
