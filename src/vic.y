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
    for id, _ := range(e1.ids) {
        ids[id] = true
    }
    for id, _ := range(e2.ids) {
        ids[id] = true
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
%token <Literal> TLPAREN TRPAREN TEQUALS TDOT
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

// grammar productions
%type <Value> constant
%type <Statement> statement
%type <Expression> expression
%type <Identifier> identifier

%%

statement   : identifier TEQUALS expression {
                stmt = &Statement{id:$1, ex:$3, lit:$1.lit+$2+$3.lit}
            }
;
identifier  : TSTRING { 
                $$ = &Identifier{ typ:RAW, lit:$1 }
            }
            | identifier TDOT identifier {
                $$ = &Identifier{ typ:DOT, root:$1, dot:$3, lit:$1.lit+$2+$3.lit }
            }
;
constant    : TNUMBER {
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
expression  : identifier {
                ids := make(map[string]bool)
                ids[$1.lit] = true
                $$ = &Expression{ typ:ID, id:$1, lit:$1.lit, ids:ids }
            }
            | constant {
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
;

%%
