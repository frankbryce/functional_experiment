// Code generated by goyacc -o parse.go vic.y. DO NOT EDIT.

//line vic.y:3
package vic

import __yyfmt__ "fmt"

//line vic.y:3

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

func mergeIds(e1, e2 *Expression) map[string]bool {
	ids := make(map[string]bool)
	if e1 != nil {
		for id, _ := range e1.ids {
			ids[id] = true
		}
	}
	if e2 != nil {
		for id, _ := range e2.ids {
			ids[id] = true
		}
	}
	return ids
}

func subtractId(e *Expression, id *Identifier) map[string]bool {
	ids := make(map[string]bool)
	for id, _ := range e.ids {
		ids[id] = true
	}
	if _, ok := ids[id.lit]; ok {
		delete(ids, id.lit)
	}
	return ids
}

//line vic.y:46
type yySymType struct {
	yys        int
	Literal    string
	Value      *Value
	Statement  *Statement
	Expression *Expression
	Identifier *Identifier
}

const TSLASH = 57346
const TDASH = 57347
const TPLUS = 57348
const TASTERISK = 57349
const TPERCENT = 57350
const TCARET = 57351
const TLPAREN = 57352
const TRPAREN = 57353
const TLBRACK = 57354
const TRBRACK = 57355
const TEQUALS = 57356
const TDOT = 57357
const TCOMMA = 57358
const ILLEGAL = 57359
const TSTRING = 57360
const TNUMBER = 57361
const TFALSE = 57362
const TTRUE = 57363
const TNULL = 57364
const NEGATE = 57365

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"TSLASH",
	"TDASH",
	"TPLUS",
	"TASTERISK",
	"TPERCENT",
	"TCARET",
	"TLPAREN",
	"TRPAREN",
	"TLBRACK",
	"TRBRACK",
	"TEQUALS",
	"TDOT",
	"TCOMMA",
	"ILLEGAL",
	"TSTRING",
	"TNUMBER",
	"TFALSE",
	"TTRUE",
	"TNULL",
	"NEGATE",
}

var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line vic.y:224

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 63

var yyAct = [...]int{
	5, 35, 6, 2, 1, 3, 38, 24, 4, 22,
	23, 37, 36, 21, 32, 25, 26, 27, 28, 29,
	30, 7, 33, 2, 17, 31, 8, 16, 18, 19,
	0, 9, 20, 18, 19, 0, 0, 20, 39, 3,
	10, 11, 12, 13, 17, 15, 14, 16, 18, 19,
	0, 34, 20, 17, 15, 14, 16, 18, 19, 19,
	0, 20, 20,
}

var yyPact = [...]int{
	-13, -1000, -6, -1000, 21, 49, 3, -1000, 21, 21,
	-8, -1000, -1000, -1000, 21, 21, 21, 21, 21, 21,
	-13, 21, 50, 40, -18, 20, 20, 25, 25, 50,
	50, -1, -5, 49, -1000, -1000, -1000, 21, -1000, 49,
}

var yyPgo = [...]int{
	0, 21, 14, 4, 0, 2,
}

var yyR1 = [...]int{
	0, 3, 3, 5, 1, 1, 1, 1, 1, 2,
	2, 2, 4, 4, 4, 4, 4, 4, 4, 4,
	4, 4, 4, 4,
}

var yyR2 = [...]int{
	0, 2, 3, 1, 1, 3, 1, 1, 1, 0,
	1, 3, 1, 1, 3, 3, 3, 3, 3, 2,
	3, 3, 4, 4,
}

var yyChk = [...]int{
	-1000, -3, -5, 18, 14, -4, -5, -1, 5, 10,
	19, 20, 21, 22, 6, 5, 7, 4, 8, 9,
	12, 10, -4, -4, 15, -4, -4, -4, -4, -4,
	-4, -3, -2, -4, 11, 19, 13, 16, 11, -4,
}

var yyDef = [...]int{
	0, -2, 0, 3, 1, 2, 12, 13, 0, 0,
	4, 6, 7, 8, 0, 0, 0, 0, 0, 0,
	0, 9, 19, 0, 0, 14, 15, 16, 17, 18,
	20, 0, 0, 10, 21, 5, 23, 0, 22, 11,
}

var yyTok1 = [...]int{
	1,
}

var yyTok2 = [...]int{
	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23,
}

var yyTok3 = [...]int{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := yyPact[state]
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && yyChk[yyAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || yyExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := yyExca[i]
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		token = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = yyTok3[i+0]
		if token == char {
			token = yyTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ??, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		yyDollar = yyS[yypt-2 : yypt+1]
//line vic.y:83
		{ // deletes identifier from runtime
			stmt = &Statement{id: yyDollar[1].Identifier, ex: nil, lit: yyDollar[1].Identifier.lit + yyDollar[2].Literal}
			yyVAL.Statement = stmt
		}
	case 2:
		yyDollar = yyS[yypt-3 : yypt+1]
//line vic.y:87
		{
			stmt = &Statement{id: yyDollar[1].Identifier, ex: yyDollar[3].Expression, lit: yyDollar[1].Identifier.lit + yyDollar[2].Literal + yyDollar[3].Expression.lit}
			yyVAL.Statement = stmt
		}
	case 3:
		yyDollar = yyS[yypt-1 : yypt+1]
//line vic.y:92
		{
			yyVAL.Identifier = &Identifier{typ: RAW, lit: yyDollar[1].Literal}
		}
	case 4:
		yyDollar = yyS[yypt-1 : yypt+1]
//line vic.y:96
		{
			f, err := strconv.ParseFloat(yyDollar[1].Literal, 64)
			if err != nil {
				panic(fmt.Errorf("Error during ParseFloat: %v", err))
			}
			yyVAL.Value = NewNumberValue(f, yyDollar[1].Literal)
		}
	case 5:
		yyDollar = yyS[yypt-3 : yypt+1]
//line vic.y:101
		{
			lit := yyDollar[1].Literal + yyDollar[2].Literal + yyDollar[3].Literal
			f, err := strconv.ParseFloat(lit, 64)
			if err != nil {
				panic(fmt.Errorf("Error during ParseFloat: %v", err))
			}
			yyVAL.Value = NewNumberValue(f, lit)
		}
	case 6:
		yyDollar = yyS[yypt-1 : yypt+1]
//line vic.y:107
		{
			yyVAL.Value = NewBoolValue(false, yyDollar[1].Literal)
		}
	case 7:
		yyDollar = yyS[yypt-1 : yypt+1]
//line vic.y:110
		{
			yyVAL.Value = NewBoolValue(true, yyDollar[1].Literal)
		}
	case 8:
		yyDollar = yyS[yypt-1 : yypt+1]
//line vic.y:113
		{
			yyVAL.Value = NewNullValue(yyDollar[1].Literal)
		}
	case 9:
		yyDollar = yyS[yypt-0 : yypt+1]
//line vic.y:117
		{
			yyVAL.Expression = &Expression{
				typ: ARGS, e: []*Expression{}, ids: make(map[string]bool), lit: "",
			}
		}
	case 10:
		yyDollar = yyS[yypt-1 : yypt+1]
//line vic.y:122
		{
			yyVAL.Expression = &Expression{
				typ: ARGS, e: []*Expression{yyDollar[1].Expression}, ids: yyDollar[1].Expression.ids, lit: yyDollar[1].Expression.lit,
			}
		}
	case 11:
		yyDollar = yyS[yypt-3 : yypt+1]
//line vic.y:127
		{
			yyVAL.Expression = &Expression{
				typ: ARGS,
				e:   append(yyDollar[1].Expression.e, yyDollar[3].Expression),
				ids: mergeIds(yyDollar[1].Expression, yyDollar[3].Expression),
				lit: yyDollar[1].Expression.lit + yyDollar[2].Literal + yyDollar[3].Expression.lit,
			}
		}
	case 12:
		yyDollar = yyS[yypt-1 : yypt+1]
//line vic.y:136
		{
			ids := make(map[string]bool)
			ids[yyDollar[1].Identifier.lit] = true
			yyVAL.Expression = &Expression{typ: ID, id: yyDollar[1].Identifier, lit: yyDollar[1].Identifier.lit, ids: ids}
		}
	case 13:
		yyDollar = yyS[yypt-1 : yypt+1]
//line vic.y:141
		{
			yyVAL.Expression = &Expression{typ: VAL, val: yyDollar[1].Value, lit: yyDollar[1].Value.lit, ids: make(map[string]bool)}
		}
	case 14:
		yyDollar = yyS[yypt-3 : yypt+1]
//line vic.y:144
		{
			yyVAL.Expression = &Expression{
				typ: PLUS, e: []*Expression{yyDollar[1].Expression, yyDollar[3].Expression}, lit: yyDollar[1].Expression.lit + yyDollar[2].Literal + yyDollar[3].Expression.lit,
				ids: mergeIds(yyDollar[1].Expression, yyDollar[3].Expression),
			}
		}
	case 15:
		yyDollar = yyS[yypt-3 : yypt+1]
//line vic.y:150
		{
			yyVAL.Expression = &Expression{
				typ: MINUS, e: []*Expression{yyDollar[1].Expression, yyDollar[3].Expression}, lit: yyDollar[1].Expression.lit + yyDollar[2].Literal + yyDollar[3].Expression.lit,
				ids: mergeIds(yyDollar[1].Expression, yyDollar[3].Expression),
			}
		}
	case 16:
		yyDollar = yyS[yypt-3 : yypt+1]
//line vic.y:156
		{
			yyVAL.Expression = &Expression{
				typ: MULT, e: []*Expression{yyDollar[1].Expression, yyDollar[3].Expression}, lit: yyDollar[1].Expression.lit + yyDollar[2].Literal + yyDollar[3].Expression.lit,
				ids: mergeIds(yyDollar[1].Expression, yyDollar[3].Expression),
			}
		}
	case 17:
		yyDollar = yyS[yypt-3 : yypt+1]
//line vic.y:162
		{
			yyVAL.Expression = &Expression{
				typ: DIV, e: []*Expression{yyDollar[1].Expression, yyDollar[3].Expression}, lit: yyDollar[1].Expression.lit + yyDollar[2].Literal + yyDollar[3].Expression.lit,
				ids: mergeIds(yyDollar[1].Expression, yyDollar[3].Expression),
			}
		}
	case 18:
		yyDollar = yyS[yypt-3 : yypt+1]
//line vic.y:168
		{
			yyVAL.Expression = &Expression{
				typ: MOD, e: []*Expression{yyDollar[1].Expression, yyDollar[3].Expression}, lit: yyDollar[1].Expression.lit + yyDollar[2].Literal + yyDollar[3].Expression.lit,
				ids: mergeIds(yyDollar[1].Expression, yyDollar[3].Expression),
			}
		}
	case 19:
		yyDollar = yyS[yypt-2 : yypt+1]
//line vic.y:174
		{
			yyVAL.Expression = &Expression{
				typ: NEG, e: []*Expression{yyDollar[2].Expression}, lit: yyDollar[1].Literal + yyDollar[2].Expression.lit,
				ids: yyDollar[2].Expression.ids,
			}
		}
	case 20:
		yyDollar = yyS[yypt-3 : yypt+1]
//line vic.y:180
		{
			yyVAL.Expression = &Expression{
				typ: POW, e: []*Expression{yyDollar[1].Expression, yyDollar[3].Expression}, lit: yyDollar[1].Expression.lit + yyDollar[2].Literal + yyDollar[3].Expression.lit,
				ids: mergeIds(yyDollar[1].Expression, yyDollar[3].Expression),
			}
		}
	case 21:
		yyDollar = yyS[yypt-3 : yypt+1]
//line vic.y:186
		{
			yyVAL.Expression = &Expression{
				typ: PAREN, e: []*Expression{yyDollar[2].Expression}, lit: yyDollar[1].Literal + yyDollar[2].Expression.lit + yyDollar[3].Literal,
				ids: yyDollar[2].Expression.ids,
			}
		}
	case 22:
		yyDollar = yyS[yypt-4 : yypt+1]
//line vic.y:192
		{
			yyVAL.Expression = &Expression{
				typ: CALL,
				id:  yyDollar[1].Identifier,
				e:   yyDollar[3].Expression.e,
				lit: yyDollar[1].Identifier.lit + yyDollar[2].Literal + yyDollar[3].Expression.lit + yyDollar[4].Literal,
				ids: yyDollar[3].Expression.ids,
			}
		}
	case 23:
		yyDollar = yyS[yypt-4 : yypt+1]
//line vic.y:201
		{
			ids := make(map[string]bool)
			if yyDollar[1].Expression != nil {
				for id, _ := range yyDollar[1].Expression.ids {
					if id == yyDollar[3].Statement.id.lit {
						continue
					}
					ids[id] = true
				}
			}
			if yyDollar[3].Statement.ex != nil {
				for id, _ := range yyDollar[3].Statement.ex.ids {
					ids[id] = true
				}
			}
			yyVAL.Expression = &Expression{
				typ: CTX,
				e:   []*Expression{yyDollar[1].Expression},
				lit: yyDollar[1].Expression.lit + yyDollar[2].Literal + yyDollar[3].Statement.lit + yyDollar[4].Literal,
				ctx: yyDollar[3].Statement,
				ids: ids,
			}
		}
	}
	goto yystack /* stack new state and value */
}
