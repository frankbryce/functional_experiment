// +build exclude
%{
package vic
%}

%union {
    Literal string
    Statement *Statement
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
%left NEGATE
%right TCARET
%right TDOT
%right TLBRACK TRBRACK

// grammar productions
%type <Statement> statement
%type <Expression> expression
%type <Identifier> identifier

%%

statement : identifier TEQUALS expression { $$ = &Statement{id: $1, ex: $3} }
;
identifier : TSTRING { $$ = &Identifier{ typ:RAW, lit:$1 } }
           | identifier TDOT identifier { $$ = &Identifier{ typ:DOT, root:$1, dot:$3 } }
           | identifier TLBRACK identifier TRBRACK {
               $$ = &Identifier{ typ:BRACK, root: $1, brack:$3 }
           }
;
expression : identifier { $$ = &Expression{ typ:ID, id:$1 } }
           | TNUMBER { $$ = &Expression{ typ:VAL, lit: $1 } }
           | TNUMBER TDOT TNUMBER { $$ = &Expression{ typ:VAL, lit: $1+$2+$3 } }
           | TFALSE { $$ = &Expression{ typ:VAL, lit: $1 } }
           | TTRUE { $$ = &Expression{ typ:VAL, lit: $1 } }
           | TNULL { $$ = &Expression{ typ:VAL, lit: $1 } }
           | expression TPLUS expression {
               $$ = &Expression{ typ:PLUS, e:[]*Expression{$1, $3} } }
           | expression TDASH expression {
               $$ = &Expression{ typ:MINUS, e:[]*Expression{$1, $3} } }
           | expression TASTERISK expression {
               $$ = &Expression{ typ:MULT, e:[]*Expression{$1, $3} } }
           | expression TSLASH expression {
               $$ = &Expression{ typ:DIV, e:[]*Expression{$1, $3} } }
           | TDASH expression %prec NEGATE {
               $$ = &Expression{ typ:NEG, e:[]*Expression{$2} } }
           | expression TCARET expression {
               $$ = &Expression{ typ:POW, e:[]*Expression{$1, $3} } }
           | TLPAREN expression TRPAREN {
               $$ = &Expression{ typ:PAREN, e:[]*Expression{$2} } }
;

%%
