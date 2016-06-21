// CAUTION: Generated file - DO NOT EDIT.

// Copyright 2016 The GC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This is a derived work, the original is published at
//
//        https://github.com/golang/go/blob/release-branch.go1.4/src/cmd/gc/go.y
//
// The original work is
//
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the GO-LICENSE file.

package gc

import __yyfmt__ "fmt"

import (
	"math/big"
	"strconv"
	"strings"
	"unicode"

	"github.com/cznic/xc"
)

type yySymType struct {
	yys   int
	node  Node
	Token xc.Token
}

type yyXError struct {
	state, xsym int
}

const (
	yyDefault     = 57421
	yyEofCode     = 57344
	ADD_ASSIGN    = 57346
	ANDAND        = 57347
	ANDNOT        = 57348
	ANDNOT_ASSIGN = 57349
	AND_ASSIGN    = 57350
	BAD_FLOAT_LIT = 57351
	BODY          = 57352
	BREAK         = 57353
	CASE          = 57354
	CHAN          = 57355
	CHAR_LIT      = 57356
	COLAS         = 57357
	COMM          = 57358
	CONST         = 57359
	CONTINUE      = 57360
	DDD           = 57361
	DEC           = 57362
	DEFAULT       = 57363
	DEFER         = 57364
	DIV_ASSIGN    = 57365
	ELSE          = 57366
	EQ            = 57367
	ERRCHECK      = 57368
	FALLTHROUGH   = 57369
	FLOAT_LIT     = 57370
	FOR           = 57371
	FUNC          = 57372
	GEQ           = 57373
	GO            = 57374
	GOTO          = 57375
	GTGT          = 57376
	IDENTIFIER    = 57377
	IF            = 57378
	IMAG_LIT      = 57379
	IMPORT        = 57380
	INC           = 57381
	INTERFACE     = 57382
	INT_LIT       = 57383
	LEQ           = 57384
	LSH           = 57385
	LSH_ASSIGN    = 57386
	LTLT          = 57387
	MAP           = 57388
	MOD_ASSIGN    = 57389
	MUL_ASSIGN    = 57390
	NEQ           = 57391
	NO_RESULT     = 57392
	OROR          = 57393
	OR_ASSIGN     = 57394
	PACKAGE       = 57395
	PARAMS        = 57396
	RANGE         = 57397
	RETURN        = 57398
	RSH           = 57399
	RSH_ASSIGN    = 57400
	RXCHAN        = 57401
	SELECT        = 57402
	STRING_LIT    = 57403
	STRUCT        = 57404
	SUB_ASSIGN    = 57405
	SWITCH        = 57406
	TXCHAN        = 57407
	TYPE          = 57408
	VAR           = 57409
	XOR_ASSIGN    = 57410
	yyErrCode     = 57345

	yyMaxDepth = 200
	yyTabOfs   = -295
)

var (
	yyXLAT = map[int]int{
		59:    0,   // ';' (253x)
		125:   1,   // '}' (243x)
		40:    2,   // '(' (235x)
		42:    3,   // '*' (225x)
		38:    4,   // '&' (186x)
		43:    5,   // '+' (186x)
		45:    6,   // '-' (186x)
		94:    7,   // '^' (186x)
		57358: 8,   // COMM (186x)
		57354: 9,   // CASE (183x)
		57363: 10,  // DEFAULT (183x)
		91:    11,  // '[' (174x)
		41:    12,  // ')' (173x)
		57377: 13,  // IDENTIFIER (171x)
		44:    14,  // ',' (167x)
		57352: 15,  // BODY (152x)
		57403: 16,  // STRING_LIT (150x)
		57355: 17,  // CHAN (139x)
		57372: 18,  // FUNC (138x)
		57382: 19,  // INTERFACE (138x)
		57388: 20,  // MAP (138x)
		57401: 21,  // RXCHAN (138x)
		57404: 22,  // STRUCT (138x)
		58:    23,  // ':' (117x)
		61:    24,  // '=' (115x)
		57424: 25,  // ArrayType (115x)
		57430: 26,  // ChanType (115x)
		57451: 27,  // FuncType (115x)
		57464: 28,  // InterfaceType (115x)
		57469: 29,  // MapType (115x)
		57486: 30,  // SliceType (115x)
		57492: 31,  // StructType (115x)
		57356: 32,  // CHAR_LIT (110x)
		57357: 33,  // COLAS (110x)
		57370: 34,  // FLOAT_LIT (110x)
		57379: 35,  // IMAG_LIT (110x)
		57383: 36,  // INT_LIT (110x)
		57361: 37,  // DDD (109x)
		33:    38,  // '!' (103x)
		123:   39,  // '{' (103x)
		93:    40,  // ']' (101x)
		57426: 41,  // BasicLiteral (90x)
		57505: 42,  // TypeLiteral (89x)
		57434: 43,  // CompLitType (88x)
		57470: 44,  // Operand (88x)
		57475: 45,  // PrimaryExpression (88x)
		57508: 46,  // UnaryExpression (87x)
		37:    47,  // '%' (83x)
		47:    48,  // '/' (83x)
		60:    49,  // '<' (83x)
		62:    50,  // '>' (83x)
		124:   51,  // '|' (83x)
		57347: 52,  // ANDAND (83x)
		57348: 53,  // ANDNOT (83x)
		57367: 54,  // EQ (83x)
		57373: 55,  // GEQ (83x)
		57384: 56,  // LEQ (83x)
		57385: 57,  // LSH (83x)
		57391: 58,  // NEQ (83x)
		57393: 59,  // OROR (83x)
		57399: 60,  // RSH (83x)
		57442: 61,  // Expression (80x)
		57346: 62,  // ADD_ASSIGN (64x)
		57350: 63,  // AND_ASSIGN (64x)
		57349: 64,  // ANDNOT_ASSIGN (64x)
		57365: 65,  // DIV_ASSIGN (64x)
		57386: 66,  // LSH_ASSIGN (64x)
		57389: 67,  // MOD_ASSIGN (64x)
		57390: 68,  // MUL_ASSIGN (64x)
		57394: 69,  // OR_ASSIGN (64x)
		57400: 70,  // RSH_ASSIGN (64x)
		57405: 71,  // SUB_ASSIGN (64x)
		57410: 72,  // XOR_ASSIGN (64x)
		57362: 73,  // DEC (61x)
		57381: 74,  // INC (61x)
		46:    75,  // '.' (42x)
		57376: 76,  // GTGT (40x)
		57443: 77,  // ExpressionList (34x)
		57477: 78,  // QualifiedIdent (34x)
		57345: 79,  // error (32x)
		57502: 80,  // Typ (26x)
		57408: 81,  // TYPE (25x)
		57378: 82,  // IF (24x)
		57353: 83,  // BREAK (23x)
		57359: 84,  // CONST (23x)
		57360: 85,  // CONTINUE (23x)
		57364: 86,  // DEFER (23x)
		57369: 87,  // FALLTHROUGH (23x)
		57371: 88,  // FOR (23x)
		57374: 89,  // GO (23x)
		57375: 90,  // GOTO (23x)
		57398: 91,  // RETURN (23x)
		57402: 92,  // SELECT (23x)
		57406: 93,  // SWITCH (23x)
		57409: 94,  // VAR (23x)
		57425: 95,  // Assignment (14x)
		57484: 96,  // SimpleStatement (14x)
		57387: 97,  // LTLT (11x)
		57454: 98,  // IdentifierList (10x)
		57344: 99,  // $end (8x)
		57427: 100, // Block (8x)
		57436: 101, // ConstDecl (7x)
		57448: 102, // ForStatement (7x)
		57457: 103, // IfStatement (7x)
		57465: 104, // LBrace (7x)
		57474: 105, // Parameters (7x)
		57481: 106, // SelectStatement (7x)
		57485: 107, // SimpleStatementOpt (7x)
		57489: 108, // StatementNonDecl (7x)
		57498: 109, // SwitchStatement (7x)
		57503: 110, // TypeDecl (7x)
		57509: 111, // VarDecl (7x)
		57431: 112, // CommaOpt (6x)
		57435: 113, // CompLitValue (6x)
		57482: 114, // SemicolonOpt (6x)
		57487: 115, // Statement (6x)
		57499: 116, // TagOpt (6x)
		57366: 117, // ELSE (5x)
		57455: 118, // IdentifierOpt (5x)
		57483: 119, // Signature (5x)
		57422: 120, // Argument (4x)
		57452: 121, // GenericArgumentsOpt (4x)
		57380: 122, // IMPORT (4x)
		57468: 123, // LBraceCompLitValue (4x)
		57488: 124, // StatementList (4x)
		57428: 125, // Body (3x)
		57437: 126, // ConstSpec (3x)
		57445: 127, // ExpressionOpt (3x)
		57460: 128, // ImportSpec (3x)
		57472: 129, // ParameterDecl (3x)
		57397: 130, // RANGE (3x)
		57506: 131, // TypeSpec (3x)
		57510: 132, // VarSpec (3x)
		57423: 133, // ArgumentList (2x)
		57429: 134, // Call (2x)
		57432: 135, // CompLitItem (2x)
		57453: 136, // GenericParametersOpt (2x)
		57456: 137, // IfHeader (2x)
		57462: 138, // InterfaceMethodDecl (2x)
		57466: 139, // LBraceCompLitItem (2x)
		57473: 140, // ParameterDeclList (2x)
		57480: 141, // ResultOpt (2x)
		57490: 142, // StructFieldDecl (2x)
		57493: 143, // SwitchBody (2x)
		57494: 144, // SwitchCase (2x)
		57495: 145, // SwitchCaseBlock (2x)
		57411: 146, // $@1 (1x)
		57412: 147, // $@10 (1x)
		57413: 148, // $@2 (1x)
		57414: 149, // $@3 (1x)
		57415: 150, // $@4 (1x)
		57416: 151, // $@5 (1x)
		57417: 152, // $@6 (1x)
		57418: 153, // $@7 (1x)
		57419: 154, // $@8 (1x)
		57420: 155, // $@9 (1x)
		57433: 156, // CompLitItemList (1x)
		57438: 157, // ConstSpecList (1x)
		57439: 158, // Elif (1x)
		57440: 159, // ElifList (1x)
		57441: 160, // ElseOpt (1x)
		57444: 161, // ExpressionListOpt (1x)
		57446: 162, // File (1x)
		57447: 163, // ForHeader (1x)
		57449: 164, // FuncBodyOpt (1x)
		57450: 165, // FuncDecl (1x)
		57458: 166, // ImportDecl (1x)
		57459: 167, // ImportList (1x)
		57461: 168, // ImportSpecList (1x)
		57463: 169, // InterfaceMethodDeclList (1x)
		57467: 170, // LBraceCompLitItemList (1x)
		57395: 171, // PACKAGE (1x)
		57471: 172, // PackageClause (1x)
		57476: 173, // Prologue (1x)
		57478: 174, // Range (1x)
		57479: 175, // ReceiverOpt (1x)
		57491: 176, // StructFieldDeclList (1x)
		57496: 177, // SwitchCaseList (1x)
		57497: 178, // SwitchHeader (1x)
		57500: 179, // TopLevelDecl (1x)
		57501: 180, // TopLevelDeclList (1x)
		57407: 181, // TXCHAN (1x)
		57504: 182, // TypeList (1x)
		57507: 183, // TypeSpecList (1x)
		57511: 184, // VarSpecList (1x)
		57421: 185, // $default (0x)
		57351: 186, // BAD_FLOAT_LIT (0x)
		57368: 187, // ERRCHECK (0x)
		57392: 188, // NO_RESULT (0x)
		57396: 189, // PARAMS (0x)
	}

	yySymNames = []string{
		"';'",
		"'}'",
		"'('",
		"'*'",
		"'&'",
		"'+'",
		"'-'",
		"'^'",
		"COMM",
		"CASE",
		"DEFAULT",
		"'['",
		"')'",
		"IDENTIFIER",
		"','",
		"BODY",
		"STRING_LIT",
		"CHAN",
		"FUNC",
		"INTERFACE",
		"MAP",
		"RXCHAN",
		"STRUCT",
		"':'",
		"'='",
		"ArrayType",
		"ChanType",
		"FuncType",
		"InterfaceType",
		"MapType",
		"SliceType",
		"StructType",
		"CHAR_LIT",
		"COLAS",
		"FLOAT_LIT",
		"IMAG_LIT",
		"INT_LIT",
		"DDD",
		"'!'",
		"'{'",
		"']'",
		"BasicLiteral",
		"TypeLiteral",
		"CompLitType",
		"Operand",
		"PrimaryExpression",
		"UnaryExpression",
		"'%'",
		"'/'",
		"'<'",
		"'>'",
		"'|'",
		"ANDAND",
		"ANDNOT",
		"EQ",
		"GEQ",
		"LEQ",
		"LSH",
		"NEQ",
		"OROR",
		"RSH",
		"Expression",
		"ADD_ASSIGN",
		"AND_ASSIGN",
		"ANDNOT_ASSIGN",
		"DIV_ASSIGN",
		"LSH_ASSIGN",
		"MOD_ASSIGN",
		"MUL_ASSIGN",
		"OR_ASSIGN",
		"RSH_ASSIGN",
		"SUB_ASSIGN",
		"XOR_ASSIGN",
		"DEC",
		"INC",
		"'.'",
		"GTGT",
		"ExpressionList",
		"QualifiedIdent",
		"error",
		"Typ",
		"TYPE",
		"IF",
		"BREAK",
		"CONST",
		"CONTINUE",
		"DEFER",
		"FALLTHROUGH",
		"FOR",
		"GO",
		"GOTO",
		"RETURN",
		"SELECT",
		"SWITCH",
		"VAR",
		"Assignment",
		"SimpleStatement",
		"LTLT",
		"IdentifierList",
		"$end",
		"Block",
		"ConstDecl",
		"ForStatement",
		"IfStatement",
		"LBrace",
		"Parameters",
		"SelectStatement",
		"SimpleStatementOpt",
		"StatementNonDecl",
		"SwitchStatement",
		"TypeDecl",
		"VarDecl",
		"CommaOpt",
		"CompLitValue",
		"SemicolonOpt",
		"Statement",
		"TagOpt",
		"ELSE",
		"IdentifierOpt",
		"Signature",
		"Argument",
		"GenericArgumentsOpt",
		"IMPORT",
		"LBraceCompLitValue",
		"StatementList",
		"Body",
		"ConstSpec",
		"ExpressionOpt",
		"ImportSpec",
		"ParameterDecl",
		"RANGE",
		"TypeSpec",
		"VarSpec",
		"ArgumentList",
		"Call",
		"CompLitItem",
		"GenericParametersOpt",
		"IfHeader",
		"InterfaceMethodDecl",
		"LBraceCompLitItem",
		"ParameterDeclList",
		"ResultOpt",
		"StructFieldDecl",
		"SwitchBody",
		"SwitchCase",
		"SwitchCaseBlock",
		"$@1",
		"$@10",
		"$@2",
		"$@3",
		"$@4",
		"$@5",
		"$@6",
		"$@7",
		"$@8",
		"$@9",
		"CompLitItemList",
		"ConstSpecList",
		"Elif",
		"ElifList",
		"ElseOpt",
		"ExpressionListOpt",
		"File",
		"ForHeader",
		"FuncBodyOpt",
		"FuncDecl",
		"ImportDecl",
		"ImportList",
		"ImportSpecList",
		"InterfaceMethodDeclList",
		"LBraceCompLitItemList",
		"PACKAGE",
		"PackageClause",
		"Prologue",
		"Range",
		"ReceiverOpt",
		"StructFieldDeclList",
		"SwitchCaseList",
		"SwitchHeader",
		"TopLevelDecl",
		"TopLevelDeclList",
		"TXCHAN",
		"TypeList",
		"TypeSpecList",
		"VarSpecList",
		"$default",
		"BAD_FLOAT_LIT",
		"ERRCHECK",
		"NO_RESULT",
		"PARAMS",
	}

	yyTokenLiteralStrings = map[int]string{
		57358: "<-",
		57354: "case",
		57363: "default",
		57377: "identifier",
		57352: "{",
		57403: "literal",
		57355: "chan",
		57372: "func",
		57382: "interface",
		57388: "map",
		57401: "<-",
		57404: "struct",
		57356: "literal",
		57357: ":=",
		57370: "literal",
		57379: "literal",
		57383: "literal",
		57361: "...",
		57347: "&&",
		57348: "&^",
		57367: "==",
		57373: ">=",
		57384: "<=",
		57385: "<<",
		57391: "!=",
		57393: "||",
		57399: ">>",
		57346: "+=",
		57350: "&=",
		57349: "&^=",
		57365: "/=",
		57386: "<<=",
		57389: "%=",
		57390: "*=",
		57394: "|=",
		57400: ">>=",
		57405: "-=",
		57410: "^=",
		57362: "--",
		57381: "++",
		57376: "»",
		57408: "type",
		57378: "if",
		57353: "break",
		57359: "const",
		57360: "continue",
		57364: "defer",
		57369: "fallthrough",
		57371: "for",
		57374: "go",
		57375: "goto",
		57398: "return",
		57402: "select",
		57406: "switch",
		57409: "var",
		57387: "«",
		57366: "else",
		57380: "import",
		57397: "range",
		57395: "package",
		57407: "<-",
		57351: "0x1p10",
		57368: "// ERROR",
	}

	yyReductions = map[int]struct{ xsym, components int }{
		0:   {0, 1},
		1:   {162, 2},
		2:   {120, 1},
		3:   {120, 1},
		4:   {133, 1},
		5:   {133, 3},
		6:   {25, 4},
		7:   {25, 4},
		8:   {95, 3},
		9:   {95, 3},
		10:  {95, 3},
		11:  {95, 3},
		12:  {95, 3},
		13:  {95, 3},
		14:  {95, 3},
		15:  {95, 3},
		16:  {95, 3},
		17:  {95, 3},
		18:  {95, 3},
		19:  {95, 3},
		20:  {41, 1},
		21:  {41, 1},
		22:  {41, 1},
		23:  {41, 1},
		24:  {41, 1},
		25:  {146, 0},
		26:  {100, 4},
		27:  {148, 0},
		28:  {125, 4},
		29:  {134, 2},
		30:  {134, 4},
		31:  {134, 5},
		32:  {26, 2},
		33:  {26, 3},
		34:  {26, 3},
		35:  {112, 0},
		36:  {112, 1},
		37:  {135, 1},
		38:  {135, 3},
		39:  {135, 3},
		40:  {135, 1},
		41:  {135, 3},
		42:  {135, 3},
		43:  {156, 1},
		44:  {156, 3},
		45:  {43, 1},
		46:  {43, 1},
		47:  {43, 1},
		48:  {43, 1},
		49:  {113, 2},
		50:  {113, 4},
		51:  {101, 3},
		52:  {101, 5},
		53:  {101, 2},
		54:  {126, 1},
		55:  {126, 3},
		56:  {126, 4},
		57:  {157, 1},
		58:  {157, 3},
		59:  {158, 4},
		60:  {159, 0},
		61:  {159, 2},
		62:  {160, 0},
		63:  {160, 2},
		64:  {61, 1},
		65:  {61, 3},
		66:  {61, 3},
		67:  {61, 3},
		68:  {61, 3},
		69:  {61, 3},
		70:  {61, 3},
		71:  {61, 3},
		72:  {61, 3},
		73:  {61, 3},
		74:  {61, 3},
		75:  {61, 3},
		76:  {61, 3},
		77:  {61, 3},
		78:  {61, 3},
		79:  {61, 3},
		80:  {61, 3},
		81:  {61, 3},
		82:  {61, 3},
		83:  {61, 3},
		84:  {61, 3},
		85:  {127, 0},
		86:  {127, 1},
		87:  {77, 1},
		88:  {77, 3},
		89:  {161, 0},
		90:  {161, 1},
		91:  {163, 1},
		92:  {163, 5},
		93:  {163, 1},
		94:  {102, 3},
		95:  {164, 0},
		96:  {164, 1},
		97:  {149, 0},
		98:  {165, 7},
		99:  {27, 2},
		100: {27, 5},
		101: {121, 0},
		102: {121, 3},
		103: {136, 0},
		104: {136, 3},
		105: {118, 0},
		106: {118, 1},
		107: {98, 1},
		108: {98, 3},
		109: {137, 1},
		110: {137, 3},
		111: {103, 5},
		112: {166, 3},
		113: {166, 5},
		114: {166, 2},
		115: {128, 2},
		116: {128, 2},
		117: {128, 3},
		118: {128, 3},
		119: {168, 1},
		120: {168, 3},
		121: {167, 0},
		122: {167, 3},
		123: {28, 3},
		124: {150, 0},
		125: {28, 6},
		126: {151, 0},
		127: {138, 3},
		128: {138, 1},
		129: {169, 1},
		130: {169, 3},
		131: {104, 1},
		132: {104, 1},
		133: {139, 1},
		134: {139, 3},
		135: {139, 3},
		136: {139, 1},
		137: {170, 1},
		138: {170, 3},
		139: {123, 2},
		140: {123, 4},
		141: {29, 5},
		142: {44, 3},
		143: {44, 3},
		144: {44, 1},
		145: {152, 0},
		146: {44, 5},
		147: {44, 2},
		148: {172, 3},
		149: {129, 2},
		150: {129, 3},
		151: {129, 2},
		152: {129, 1},
		153: {140, 1},
		154: {140, 3},
		155: {105, 2},
		156: {105, 4},
		157: {45, 1},
		158: {45, 2},
		159: {45, 5},
		160: {45, 5},
		161: {45, 3},
		162: {45, 4},
		163: {45, 8},
		164: {45, 6},
		165: {45, 2},
		166: {45, 2},
		167: {45, 5},
		168: {173, 2},
		169: {78, 1},
		170: {78, 3},
		171: {174, 4},
		172: {174, 4},
		173: {174, 2},
		174: {175, 0},
		175: {175, 1},
		176: {141, 0},
		177: {141, 1},
		178: {141, 1},
		179: {153, 0},
		180: {106, 3},
		181: {114, 0},
		182: {114, 1},
		183: {119, 2},
		184: {96, 1},
		185: {96, 1},
		186: {96, 2},
		187: {96, 2},
		188: {96, 3},
		189: {107, 0},
		190: {107, 1},
		191: {30, 3},
		192: {115, 0},
		193: {115, 1},
		194: {115, 1},
		195: {115, 1},
		196: {115, 1},
		197: {115, 1},
		198: {115, 1},
		199: {124, 1},
		200: {124, 3},
		201: {108, 2},
		202: {108, 2},
		203: {108, 2},
		204: {108, 1},
		205: {108, 1},
		206: {108, 2},
		207: {108, 2},
		208: {108, 3},
		209: {108, 1},
		210: {108, 2},
		211: {108, 1},
		212: {108, 1},
		213: {108, 1},
		214: {142, 3},
		215: {142, 3},
		216: {142, 2},
		217: {142, 4},
		218: {142, 5},
		219: {142, 5},
		220: {176, 1},
		221: {176, 3},
		222: {31, 3},
		223: {154, 0},
		224: {31, 6},
		225: {143, 2},
		226: {155, 0},
		227: {143, 4},
		228: {144, 3},
		229: {144, 5},
		230: {144, 5},
		231: {144, 2},
		232: {144, 2},
		233: {144, 2},
		234: {145, 2},
		235: {177, 1},
		236: {177, 2},
		237: {178, 1},
		238: {178, 2},
		239: {178, 3},
		240: {178, 9},
		241: {147, 0},
		242: {109, 4},
		243: {116, 0},
		244: {116, 1},
		245: {179, 1},
		246: {179, 1},
		247: {179, 1},
		248: {179, 1},
		249: {179, 1},
		250: {179, 1},
		251: {180, 0},
		252: {180, 3},
		253: {80, 3},
		254: {80, 2},
		255: {80, 1},
		256: {80, 1},
		257: {80, 1},
		258: {80, 1},
		259: {80, 1},
		260: {80, 2},
		261: {80, 1},
		262: {80, 1},
		263: {110, 3},
		264: {110, 5},
		265: {110, 2},
		266: {182, 1},
		267: {182, 3},
		268: {42, 2},
		269: {42, 1},
		270: {42, 1},
		271: {42, 1},
		272: {42, 1},
		273: {42, 1},
		274: {42, 1},
		275: {42, 1},
		276: {131, 3},
		277: {183, 1},
		278: {183, 3},
		279: {46, 2},
		280: {46, 2},
		281: {46, 2},
		282: {46, 2},
		283: {46, 2},
		284: {46, 2},
		285: {46, 2},
		286: {46, 1},
		287: {111, 3},
		288: {111, 5},
		289: {111, 2},
		290: {132, 3},
		291: {132, 2},
		292: {132, 4},
		293: {184, 1},
		294: {184, 3},
	}

	yyXErrors = map[yyXError]string{
		yyXError{0, 99}:   "invalid empty input",
		yyXError{1, -1}:   "expected $end",
		yyXError{56, -1}:  "expected '('",
		yyXError{284, -1}: "expected '('",
		yyXError{285, -1}: "expected '('",
		yyXError{286, -1}: "expected '('",
		yyXError{287, -1}: "expected '('",
		yyXError{288, -1}: "expected '('",
		yyXError{289, -1}: "expected '('",
		yyXError{26, -1}:  "expected ')'",
		yyXError{99, -1}:  "expected ')'",
		yyXError{173, -1}: "expected ')'",
		yyXError{193, -1}: "expected ')'",
		yyXError{214, -1}: "expected ')'",
		yyXError{215, -1}: "expected ')'",
		yyXError{229, -1}: "expected ')'",
		yyXError{308, -1}: "expected ')'",
		yyXError{309, -1}: "expected ')'",
		yyXError{330, -1}: "expected ')'",
		yyXError{332, -1}: "expected ')'",
		yyXError{351, -1}: "expected ')'",
		yyXError{353, -1}: "expected ')'",
		yyXError{362, -1}: "expected ')'",
		yyXError{379, -1}: "expected ')'",
		yyXError{501, -1}: "expected ')'",
		yyXError{297, -1}: "expected ':'",
		yyXError{7, -1}:   "expected ';'",
		yyXError{11, -1}:  "expected ';'",
		yyXError{23, -1}:  "expected ';'",
		yyXError{29, -1}:  "expected ';'",
		yyXError{30, -1}:  "expected ';'",
		yyXError{73, -1}:  "expected ';'",
		yyXError{74, -1}:  "expected ';'",
		yyXError{75, -1}:  "expected ';'",
		yyXError{76, -1}:  "expected ';'",
		yyXError{77, -1}:  "expected ';'",
		yyXError{78, -1}:  "expected ';'",
		yyXError{79, -1}:  "expected ';'",
		yyXError{445, -1}: "expected ';'",
		yyXError{446, -1}: "expected ';'",
		yyXError{455, -1}: "expected ';'",
		yyXError{494, -1}: "expected '='",
		yyXError{48, -1}:  "expected '['",
		yyXError{303, -1}: "expected ']'",
		yyXError{406, -1}: "expected ']'",
		yyXError{512, -1}: "expected ']'",
		yyXError{318, -1}: "expected '}'",
		yyXError{348, -1}: "expected '}'",
		yyXError{389, -1}: "expected '}'",
		yyXError{418, -1}: "expected '}'",
		yyXError{290, -1}: "expected argument list or one of ['!', '&', '(', ')', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{46, -1}:  "expected block statement or if statement header or one of ['!', '&', '(', '*', '+', '-', ';', '[', '^', <-, chan, func, identifier, interface, literal, map, struct, {]",
		yyXError{430, -1}: "expected block statement or if statement header or one of ['!', '&', '(', '*', '+', '-', ';', '[', '^', <-, chan, func, identifier, interface, literal, map, struct, {]",
		yyXError{427, -1}: "expected block statement or one of ['{', if]",
		yyXError{423, -1}: "expected block statement or {",
		yyXError{432, -1}: "expected block statement or {",
		yyXError{450, -1}: "expected block statement or {",
		yyXError{72, -1}:  "expected body of the switch statement or switch statement header or one of ['!', '&', '(', '*', '+', '-', ';', '[', '^', <-, chan, func, identifier, interface, literal, map, struct, {]",
		yyXError{57, -1}:  "expected body of the switch statement or {",
		yyXError{235, -1}: "expected body of the switch statement or {",
		yyXError{236, -1}: "expected body of the switch statement or {",
		yyXError{375, -1}: "expected body of the switch statement or {",
		yyXError{55, -1}:  "expected call or composite literal value or one of [!=, %=, &&, &=, &^, &^=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '[', ']', '^', '{', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{282, -1}: "expected call or composite literal value or one of ['(', '.', '[', '{']",
		yyXError{36, -1}:  "expected chan",
		yyXError{381, -1}: "expected composite literal item list or one of ['!', '&', '(', '*', '+', '-', '[', '^', '{', '}', <-, chan, func, identifier, interface, literal, map, struct, {]",
		yyXError{291, -1}: "expected composite literal item list or one of ['!', '&', '(', '*', '+', '-', '[', '^', '{', '}', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{388, -1}: "expected composite literal item or one of ['!', '&', '(', '*', '+', '-', '[', '^', '{', '}', <-, chan, func, identifier, interface, literal, map, struct, {]",
		yyXError{317, -1}: "expected composite literal item or one of ['!', '&', '(', '*', '+', '-', '[', '^', '{', '}', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{392, -1}: "expected composite literal value or expression or one of ['!', '&', '(', '*', '+', '-', '[', '^', '{', <-, chan, func, identifier, interface, literal, map, struct, {]",
		yyXError{321, -1}: "expected composite literal value or expression or one of ['!', '&', '(', '*', '+', '-', '[', '^', '{', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{324, -1}: "expected composite literal value or expression or one of ['!', '&', '(', '*', '+', '-', '[', '^', '{', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{54, -1}:  "expected composite literal value or one of ['{', {]",
		yyXError{490, -1}: "expected constant specification list or one of [')', identifier]",
		yyXError{41, -1}:  "expected constant specification or one of ['(', identifier]",
		yyXError{502, -1}: "expected constant specification or one of [')', identifier]",
		yyXError{426, -1}: "expected else if clause or optional else clause or one of [';', '}', case, default, else]",
		yyXError{425, -1}: "expected else if list clause or optional else clause or one of [';', '}', case, default, else]",
		yyXError{458, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, range, struct]",
		yyXError{470, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, range, struct]",
		yyXError{115, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{117, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{459, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{460, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{461, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{462, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{463, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{464, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{465, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{466, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{467, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{468, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{469, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{493, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{495, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{510, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{511, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{33, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '[', ']', '^', ..., <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{278, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct, {]",
		yyXError{61, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{64, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{128, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{130, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{131, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{132, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{133, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{134, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{135, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{136, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{137, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{138, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{139, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{140, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{141, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{142, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{143, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{144, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{145, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{146, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{147, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{148, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{149, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{212, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{269, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{270, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{451, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{472, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{486, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{293, -1}: "expected expression or optional expression or one of ['!', '&', '(', '*', '+', '-', ':', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{49, -1}:  "expected expression or type literal or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{329, -1}: "expected expression/type literal or one of ['!', '&', '(', ')', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{267, -1}: "expected expression/type literal or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{44, -1}:  "expected for statement header or one of ['!', '&', '(', '*', '+', '-', ';', '[', '^', <-, chan, func, identifier, interface, literal, map, range, struct, {]",
		yyXError{202, -1}: "expected function/method signature or '('",
		yyXError{416, -1}: "expected function/method signature or '('",
		yyXError{442, -1}: "expected function/method signature or '('",
		yyXError{413, -1}: "expected function/method signature or one of ['(', '.', ';', '}']",
		yyXError{102, -1}: "expected function/method signature or one of ['(', «]",
		yyXError{441, -1}: "expected function/method signature or optional generic parameters or one of ['(', «]",
		yyXError{45, -1}:  "expected function/method signature or optional receiver or one of ['(', identifier, «]",
		yyXError{3, -1}:   "expected identifier",
		yyXError{65, -1}:  "expected identifier",
		yyXError{103, -1}: "expected identifier",
		yyXError{175, -1}: "expected identifier",
		yyXError{439, -1}: "expected identifier",
		yyXError{179, -1}: "expected identifier list or identifier",
		yyXError{220, -1}: "expected identifier list or identifier",
		yyXError{10, -1}:  "expected import specification list or one of [')', '.', identifier, literal]",
		yyXError{6, -1}:   "expected import specification or one of ['(', '.', identifier, literal]",
		yyXError{27, -1}:  "expected import specification or one of [')', '.', identifier, literal]",
		yyXError{411, -1}: "expected interface method declaration list or identifier",
		yyXError{409, -1}: "expected interface method declaration list or one of ['}', identifier]",
		yyXError{419, -1}: "expected interface method declaration or one of ['}', identifier]",
		yyXError{51, -1}:  "expected left brace or one of ['(', ')', ',', ':', '=', '{', ..., :=, {]",
		yyXError{47, -1}:  "expected left brace or one of ['{', {]",
		yyXError{71, -1}:  "expected left brace or one of ['{', {]",
		yyXError{397, -1}: "expected left brace or one of ['{', {]",
		yyXError{12, -1}:  "expected literal or ",
		yyXError{13, -1}:  "expected literal or ",
		yyXError{125, -1}: "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '[', ']', '^', '{', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, literal, {, |=, ||, »]",
		yyXError{50, -1}:  "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '[', ']', '^', '{', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{53, -1}:  "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '[', ']', '^', '{', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{122, -1}: "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '[', ']', '^', '{', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{216, -1}: "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '[', ']', '^', '{', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{294, -1}: "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '[', ']', '^', '{', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{295, -1}: "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '[', ']', '^', '{', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{302, -1}: "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '[', ']', '^', '{', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{304, -1}: "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '[', ']', '^', '{', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{305, -1}: "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '[', ']', '^', '{', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{307, -1}: "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '[', ']', '^', '{', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{310, -1}: "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '[', ']', '^', '{', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{316, -1}: "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '[', ']', '^', '{', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{319, -1}: "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '[', ']', '^', '{', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{328, -1}: "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '[', ']', '^', '{', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{333, -1}: "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '[', ']', '^', '{', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{334, -1}: "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '[', ']', '^', '{', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{380, -1}: "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '[', ']', '^', '{', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{382, -1}: "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '[', ']', '^', '{', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{387, -1}: "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '[', ']', '^', '{', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{390, -1}: "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '[', ']', '^', '{', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{400, -1}: "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '[', ']', '^', '{', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{403, -1}: "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '[', ']', '^', '{', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{404, -1}: "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '[', ']', '^', '{', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{42, -1}:  "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', ')', '*', '+', ',', '-', '/', ':', ';', '<', '=', '>', ']', '^', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{150, -1}: "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', ')', '*', '+', ',', '-', '/', ':', ';', '<', '=', '>', ']', '^', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{151, -1}: "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', ')', '*', '+', ',', '-', '/', ':', ';', '<', '=', '>', ']', '^', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{152, -1}: "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', ')', '*', '+', ',', '-', '/', ':', ';', '<', '=', '>', ']', '^', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{153, -1}: "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', ')', '*', '+', ',', '-', '/', ':', ';', '<', '=', '>', ']', '^', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{154, -1}: "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', ')', '*', '+', ',', '-', '/', ':', ';', '<', '=', '>', ']', '^', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{155, -1}: "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', ')', '*', '+', ',', '-', '/', ':', ';', '<', '=', '>', ']', '^', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{156, -1}: "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', ')', '*', '+', ',', '-', '/', ':', ';', '<', '=', '>', ']', '^', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{157, -1}: "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', ')', '*', '+', ',', '-', '/', ':', ';', '<', '=', '>', ']', '^', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{158, -1}: "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', ')', '*', '+', ',', '-', '/', ':', ';', '<', '=', '>', ']', '^', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{159, -1}: "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', ')', '*', '+', ',', '-', '/', ':', ';', '<', '=', '>', ']', '^', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{160, -1}: "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', ')', '*', '+', ',', '-', '/', ':', ';', '<', '=', '>', ']', '^', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{161, -1}: "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', ')', '*', '+', ',', '-', '/', ':', ';', '<', '=', '>', ']', '^', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{162, -1}: "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', ')', '*', '+', ',', '-', '/', ':', ';', '<', '=', '>', ']', '^', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{163, -1}: "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', ')', '*', '+', ',', '-', '/', ':', ';', '<', '=', '>', ']', '^', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{164, -1}: "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', ')', '*', '+', ',', '-', '/', ':', ';', '<', '=', '>', ']', '^', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{165, -1}: "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', ')', '*', '+', ',', '-', '/', ':', ';', '<', '=', '>', ']', '^', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{166, -1}: "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', ')', '*', '+', ',', '-', '/', ':', ';', '<', '=', '>', ']', '^', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{167, -1}: "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', ')', '*', '+', ',', '-', '/', ':', ';', '<', '=', '>', ']', '^', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{168, -1}: "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', ')', '*', '+', ',', '-', '/', ':', ';', '<', '=', '>', ']', '^', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{169, -1}: "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', ')', '*', '+', ',', '-', '/', ':', ';', '<', '=', '>', ']', '^', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{204, -1}: "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', ')', '*', '+', ',', '-', '/', ':', ';', '<', '=', '>', ']', '^', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{205, -1}: "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', ')', '*', '+', ',', '-', '/', ':', ';', '<', '=', '>', ']', '^', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{206, -1}: "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', ')', '*', '+', ',', '-', '/', ':', ';', '<', '=', '>', ']', '^', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{207, -1}: "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', ')', '*', '+', ',', '-', '/', ':', ';', '<', '=', '>', ']', '^', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{208, -1}: "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', ')', '*', '+', ',', '-', '/', ':', ';', '<', '=', '>', ']', '^', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{209, -1}: "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', ')', '*', '+', ',', '-', '/', ':', ';', '<', '=', '>', ']', '^', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{211, -1}: "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', ')', '*', '+', ',', '-', '/', ':', ';', '<', '=', '>', ']', '^', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{129, -1}: "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', ')', '*', '+', ',', '-', '/', ';', '<', '=', '>', '^', '|', '}', *=, +=, -=, /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{43, -1}:  "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', '*', '+', ',', '-', '/', ';', '<', '=', '>', '^', '|', '}', *=, ++, +=, --, -=, /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{262, -1}: "expected one of [!=, &&, &^, '%', '&', ')', '*', '+', ',', '-', '/', ':', '<', '=', '>', '^', '|', ..., :=, <-, <<, <=, ==, >=, >>, ||]",
		yyXError{118, -1}: "expected one of [!=, &&, &^, '%', '&', ')', '*', '+', ',', '-', '/', ';', '<', '>', '^', '|', '}', <-, <<, <=, ==, >=, >>, case, default, {, ||]",
		yyXError{401, -1}: "expected one of [!=, &&, &^, '%', '&', ')', '*', '+', '-', '/', '<', '>', '^', '|', <-, <<, <=, ==, >=, >>, ||]",
		yyXError{313, -1}: "expected one of [!=, &&, &^, '%', '&', '*', '+', ',', '-', '/', ':', '<', '>', '^', '|', '}', <-, <<, <=, ==, >=, >>, ||]",
		yyXError{383, -1}: "expected one of [!=, &&, &^, '%', '&', '*', '+', ',', '-', '/', ':', '<', '>', '^', '|', '}', <-, <<, <=, ==, >=, >>, ||]",
		yyXError{323, -1}: "expected one of [!=, &&, &^, '%', '&', '*', '+', ',', '-', '/', '<', '>', '^', '|', '}', <-, <<, <=, ==, >=, >>, ||]",
		yyXError{326, -1}: "expected one of [!=, &&, &^, '%', '&', '*', '+', ',', '-', '/', '<', '>', '^', '|', '}', <-, <<, <=, ==, >=, >>, ||]",
		yyXError{393, -1}: "expected one of [!=, &&, &^, '%', '&', '*', '+', ',', '-', '/', '<', '>', '^', '|', '}', <-, <<, <=, ==, >=, >>, ||]",
		yyXError{296, -1}: "expected one of [!=, &&, &^, '%', '&', '*', '+', '-', '/', ':', '<', '>', ']', '^', '|', <-, <<, <=, ==, >=, >>, ||]",
		yyXError{299, -1}: "expected one of [!=, &&, &^, '%', '&', '*', '+', '-', '/', ':', '<', '>', ']', '^', '|', <-, <<, <=, ==, >=, >>, ||]",
		yyXError{271, -1}: "expected one of [!=, &&, &^, '%', '&', '*', '+', '-', '/', ':', '<', '>', '^', '|', <-, <<, <=, ==, >=, >>, ||]",
		yyXError{273, -1}: "expected one of [!=, &&, &^, '%', '&', '*', '+', '-', '/', ':', '<', '>', '^', '|', <-, <<, <=, ==, >=, >>, ||]",
		yyXError{371, -1}: "expected one of [!=, &&, &^, '%', '&', '*', '+', '-', '/', ';', '<', '>', '^', '|', '}', <-, <<, <=, ==, >=, >>, case, default, ||]",
		yyXError{372, -1}: "expected one of [!=, &&, &^, '%', '&', '*', '+', '-', '/', ';', '<', '>', '^', '|', '}', <-, <<, <=, ==, >=, >>, case, default, ||]",
		yyXError{513, -1}: "expected one of [!=, &&, &^, '%', '&', '*', '+', '-', '/', '<', '>', ']', '^', '|', <-, <<, <=, ==, >=, >>, ||]",
		yyXError{279, -1}: "expected one of [!=, &&, &^, '%', '&', '*', '+', '-', '/', '<', '>', '^', '|', <-, <<, <=, ==, >=, >>, {, ||]",
		yyXError{452, -1}: "expected one of [!=, &&, &^, '%', '&', '*', '+', '-', '/', '<', '>', '^', '|', <-, <<, <=, ==, >=, >>, {, ||]",
		yyXError{473, -1}: "expected one of [!=, &&, &^, '%', '&', '*', '+', '-', '/', '<', '>', '^', '|', <-, <<, <=, ==, >=, >>, {, ||]",
		yyXError{487, -1}: "expected one of [!=, &&, &^, '%', '&', '*', '+', '-', '/', '<', '>', '^', '|', <-, <<, <=, ==, >=, >>, {, ||]",
		yyXError{34, -1}:  "expected one of [%=, &=, &^=, ',', '=', *=, +=, -=, /=, :=, <<=, >>=, ^=, |=]",
		yyXError{447, -1}: "expected one of [%=, &=, &^=, ',', '=', *=, +=, -=, /=, :=, <<=, >>=, ^=, |=]",
		yyXError{191, -1}: "expected one of ['(', ')', '*', ',', ':', ';', '=', '[', ']', '{', '}', ..., :=, <-, case, chan, default, func, identifier, interface, literal, map, struct, {, »]",
		yyXError{194, -1}: "expected one of ['(', ')', '*', ',', ':', ';', '=', '[', ']', '{', '}', ..., :=, <-, case, chan, default, func, identifier, interface, literal, map, struct, {, »]",
		yyXError{91, -1}:  "expected one of ['(', ')', '*', ',', ';', '=', '[', '}', <-, case, chan, default, func, identifier, interface, map, struct, »]",
		yyXError{177, -1}: "expected one of ['(', ')', '*', ',', ';', '=', '[', '}', <-, case, chan, default, func, identifier, interface, map, struct, »]",
		yyXError{104, -1}: "expected one of ['(', ')', ',', '.', ':', ';', '=', ']', '{', '}', ..., :=, case, default, literal, {, «, »]",
		yyXError{176, -1}: "expected one of ['(', ')', ',', ':', ';', '=', ']', '{', '}', ..., :=, case, default, literal, {, «, »]",
		yyXError{107, -1}: "expected one of ['(', ')', ',', ':', ';', '=', ']', '{', '}', ..., :=, case, default, literal, {, »]",
		yyXError{108, -1}: "expected one of ['(', ')', ',', ':', ';', '=', ']', '{', '}', ..., :=, case, default, literal, {, »]",
		yyXError{109, -1}: "expected one of ['(', ')', ',', ':', ';', '=', ']', '{', '}', ..., :=, case, default, literal, {, »]",
		yyXError{110, -1}: "expected one of ['(', ')', ',', ':', ';', '=', ']', '{', '}', ..., :=, case, default, literal, {, »]",
		yyXError{111, -1}: "expected one of ['(', ')', ',', ':', ';', '=', ']', '{', '}', ..., :=, case, default, literal, {, »]",
		yyXError{113, -1}: "expected one of ['(', ')', ',', ':', ';', '=', ']', '{', '}', ..., :=, case, default, literal, {, »]",
		yyXError{114, -1}: "expected one of ['(', ')', ',', ':', ';', '=', ']', '{', '}', ..., :=, case, default, literal, {, »]",
		yyXError{171, -1}: "expected one of ['(', ')', ',', ':', ';', '=', ']', '{', '}', ..., :=, case, default, literal, {, »]",
		yyXError{172, -1}: "expected one of ['(', ')', ',', ':', ';', '=', ']', '{', '}', ..., :=, case, default, literal, {, »]",
		yyXError{174, -1}: "expected one of ['(', ')', ',', ':', ';', '=', ']', '{', '}', ..., :=, case, default, literal, {, »]",
		yyXError{178, -1}: "expected one of ['(', ')', ',', ':', ';', '=', ']', '{', '}', ..., :=, case, default, literal, {, »]",
		yyXError{183, -1}: "expected one of ['(', ')', ',', ':', ';', '=', ']', '{', '}', ..., :=, case, default, literal, {, »]",
		yyXError{184, -1}: "expected one of ['(', ')', ',', ':', ';', '=', ']', '{', '}', ..., :=, case, default, literal, {, »]",
		yyXError{185, -1}: "expected one of ['(', ')', ',', ':', ';', '=', ']', '{', '}', ..., :=, case, default, literal, {, »]",
		yyXError{203, -1}: "expected one of ['(', ')', ',', ':', ';', '=', ']', '{', '}', ..., :=, case, default, literal, {, »]",
		yyXError{338, -1}: "expected one of ['(', ')', ',', ':', ';', '=', ']', '{', '}', ..., :=, case, default, literal, {, »]",
		yyXError{349, -1}: "expected one of ['(', ')', ',', ':', ';', '=', ']', '{', '}', ..., :=, case, default, literal, {, »]",
		yyXError{408, -1}: "expected one of ['(', ')', ',', ':', ';', '=', ']', '{', '}', ..., :=, case, default, literal, {, »]",
		yyXError{410, -1}: "expected one of ['(', ')', ',', ':', ';', '=', ']', '{', '}', ..., :=, case, default, literal, {, »]",
		yyXError{421, -1}: "expected one of ['(', ')', ',', ':', ';', '=', ']', '{', '}', ..., :=, case, default, literal, {, »]",
		yyXError{506, -1}: "expected one of ['(', ')', ',', ':', ';', '=', ']', '{', '}', ..., :=, case, default, literal, {, »]",
		yyXError{507, -1}: "expected one of ['(', ')', ',', ':', ';', '=', ']', '{', '}', ..., :=, case, default, literal, {, »]",
		yyXError{509, -1}: "expected one of ['(', ')', ',', ':', ';', '=', ']', '{', '}', ..., :=, case, default, literal, {, »]",
		yyXError{515, -1}: "expected one of ['(', ')', ',', ':', ';', '=', ']', '{', '}', ..., :=, case, default, literal, {, »]",
		yyXError{517, -1}: "expected one of ['(', ')', ',', ':', ';', '=', ']', '{', '}', ..., :=, case, default, literal, {, »]",
		yyXError{519, -1}: "expected one of ['(', ')', ',', ':', ';', '=', ']', '{', '}', ..., :=, case, default, literal, {, »]",
		yyXError{37, -1}:  "expected one of ['(', ')', ',', ':', '=', '{', ..., :=, {]",
		yyXError{38, -1}:  "expected one of ['(', ')', ',', ':', '=', '{', ..., :=, {]",
		yyXError{39, -1}:  "expected one of ['(', ')', ',', ':', '=', '{', ..., :=, {]",
		yyXError{40, -1}:  "expected one of ['(', ')', ',', ':', '=', '{', ..., :=, {]",
		yyXError{82, -1}:  "expected one of ['(', ')', ',', ':', '=', ..., :=]",
		yyXError{83, -1}:  "expected one of ['(', ')', ',', ':', '=', ..., :=]",
		yyXError{210, -1}: "expected one of ['(', ')', ',', ':', '=', ..., :=]",
		yyXError{263, -1}: "expected one of ['(', ')', ',', ':', '=', ..., :=]",
		yyXError{402, -1}: "expected one of ['(', ')']",
		yyXError{340, -1}: "expected one of ['(', '*', ',', '.', ';', '[', '}', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{224, -1}: "expected one of ['(', '*', '[', <-, chan, func, identifier, interface, map, struct]",
		yyXError{311, -1}: "expected one of ['(', '.', '[', '{', {]",
		yyXError{292, -1}: "expected one of ['(', identifier]",
		yyXError{377, -1}: "expected one of ['(', identifier]",
		yyXError{264, -1}: "expected one of [')', ',', ':', '=', ..., :=]",
		yyXError{275, -1}: "expected one of [')', ',', ':', '=', ..., :=]",
		yyXError{119, -1}: "expected one of [')', ',', ';', '}', case, default]",
		yyXError{170, -1}: "expected one of [')', ',', ';', '}', case, default]",
		yyXError{496, -1}: "expected one of [')', ',', ';', '}', case, default]",
		yyXError{497, -1}: "expected one of [')', ',', ';', '}', case, default]",
		yyXError{188, -1}: "expected one of [')', ',']",
		yyXError{189, -1}: "expected one of [')', ',']",
		yyXError{195, -1}: "expected one of [')', ',']",
		yyXError{196, -1}: "expected one of [')', ',']",
		yyXError{198, -1}: "expected one of [')', ',']",
		yyXError{199, -1}: "expected one of [')', ',']",
		yyXError{200, -1}: "expected one of [')', ',']",
		yyXError{116, -1}: "expected one of [')', ';', '=', '}', case, default]",
		yyXError{222, -1}: "expected one of [')', ';', '}', case, default]",
		yyXError{20, -1}:  "expected one of [')', ';']",
		yyXError{22, -1}:  "expected one of [')', ';']",
		yyXError{25, -1}:  "expected one of [')', ';']",
		yyXError{28, -1}:  "expected one of [')', ';']",
		yyXError{97, -1}:  "expected one of [')', ';']",
		yyXError{101, -1}: "expected one of [')', ';']",
		yyXError{227, -1}: "expected one of [')', ';']",
		yyXError{231, -1}: "expected one of [')', ';']",
		yyXError{500, -1}: "expected one of [')', ';']",
		yyXError{503, -1}: "expected one of [')', ';']",
		yyXError{265, -1}: "expected one of [',', ':', '=', :=]",
		yyXError{312, -1}: "expected one of [',', ':', '}']",
		yyXError{471, -1}: "expected one of [',', ';', '}', case, default, {]",
		yyXError{474, -1}: "expected one of [',', ';', '}', case, default, {]",
		yyXError{475, -1}: "expected one of [',', ';', '}', case, default, {]",
		yyXError{476, -1}: "expected one of [',', ';', '}', case, default, {]",
		yyXError{477, -1}: "expected one of [',', ';', '}', case, default, {]",
		yyXError{478, -1}: "expected one of [',', ';', '}', case, default, {]",
		yyXError{479, -1}: "expected one of [',', ';', '}', case, default, {]",
		yyXError{480, -1}: "expected one of [',', ';', '}', case, default, {]",
		yyXError{481, -1}: "expected one of [',', ';', '}', case, default, {]",
		yyXError{482, -1}: "expected one of [',', ';', '}', case, default, {]",
		yyXError{483, -1}: "expected one of [',', ';', '}', case, default, {]",
		yyXError{484, -1}: "expected one of [',', ';', '}', case, default, {]",
		yyXError{485, -1}: "expected one of [',', ';', '}', case, default, {]",
		yyXError{368, -1}: "expected one of [',', ';', '}', case, default]",
		yyXError{314, -1}: "expected one of [',', '}']",
		yyXError{320, -1}: "expected one of [',', '}']",
		yyXError{322, -1}: "expected one of [',', '}']",
		yyXError{325, -1}: "expected one of [',', '}']",
		yyXError{384, -1}: "expected one of [',', '}']",
		yyXError{385, -1}: "expected one of [',', '}']",
		yyXError{391, -1}: "expected one of [',', '}']",
		yyXError{394, -1}: "expected one of [',', '}']",
		yyXError{123, -1}: "expected one of [',', »]",
		yyXError{124, -1}: "expected one of [',', »]",
		yyXError{127, -1}: "expected one of [',', »]",
		yyXError{201, -1}: "expected one of [',', »]",
		yyXError{223, -1}: "expected one of [',', »]",
		yyXError{300, -1}: "expected one of [':', ']']",
		yyXError{428, -1}: "expected one of [';', '}', case, default, else]",
		yyXError{433, -1}: "expected one of [';', '}', case, default, else]",
		yyXError{436, -1}: "expected one of [';', '}', case, default, else]",
		yyXError{9, -1}:   "expected one of [';', '}', case, default, literal]",
		yyXError{58, -1}:  "expected one of [';', '}', case, default, {]",
		yyXError{488, -1}: "expected one of [';', '}', case, default, {]",
		yyXError{489, -1}: "expected one of [';', '}', case, default, {]",
		yyXError{62, -1}:  "expected one of [';', '}', case, default]",
		yyXError{63, -1}:  "expected one of [';', '}', case, default]",
		yyXError{66, -1}:  "expected one of [';', '}', case, default]",
		yyXError{68, -1}:  "expected one of [';', '}', case, default]",
		yyXError{69, -1}:  "expected one of [';', '}', case, default]",
		yyXError{70, -1}:  "expected one of [';', '}', case, default]",
		yyXError{94, -1}:  "expected one of [';', '}', case, default]",
		yyXError{95, -1}:  "expected one of [';', '}', case, default]",
		yyXError{100, -1}: "expected one of [';', '}', case, default]",
		yyXError{218, -1}: "expected one of [';', '}', case, default]",
		yyXError{225, -1}: "expected one of [';', '}', case, default]",
		yyXError{230, -1}: "expected one of [';', '}', case, default]",
		yyXError{238, -1}: "expected one of [';', '}', case, default]",
		yyXError{239, -1}: "expected one of [';', '}', case, default]",
		yyXError{247, -1}: "expected one of [';', '}', case, default]",
		yyXError{248, -1}: "expected one of [';', '}', case, default]",
		yyXError{249, -1}: "expected one of [';', '}', case, default]",
		yyXError{250, -1}: "expected one of [';', '}', case, default]",
		yyXError{251, -1}: "expected one of [';', '}', case, default]",
		yyXError{252, -1}: "expected one of [';', '}', case, default]",
		yyXError{253, -1}: "expected one of [';', '}', case, default]",
		yyXError{254, -1}: "expected one of [';', '}', case, default]",
		yyXError{256, -1}: "expected one of [';', '}', case, default]",
		yyXError{259, -1}: "expected one of [';', '}', case, default]",
		yyXError{276, -1}: "expected one of [';', '}', case, default]",
		yyXError{369, -1}: "expected one of [';', '}', case, default]",
		yyXError{370, -1}: "expected one of [';', '}', case, default]",
		yyXError{373, -1}: "expected one of [';', '}', case, default]",
		yyXError{374, -1}: "expected one of [';', '}', case, default]",
		yyXError{376, -1}: "expected one of [';', '}', case, default]",
		yyXError{396, -1}: "expected one of [';', '}', case, default]",
		yyXError{429, -1}: "expected one of [';', '}', case, default]",
		yyXError{431, -1}: "expected one of [';', '}', case, default]",
		yyXError{453, -1}: "expected one of [';', '}', case, default]",
		yyXError{491, -1}: "expected one of [';', '}', case, default]",
		yyXError{498, -1}: "expected one of [';', '}', case, default]",
		yyXError{504, -1}: "expected one of [';', '}', case, default]",
		yyXError{258, -1}: "expected one of [';', '}']",
		yyXError{345, -1}: "expected one of [';', '}']",
		yyXError{350, -1}: "expected one of [';', '}']",
		yyXError{355, -1}: "expected one of [';', '}']",
		yyXError{356, -1}: "expected one of [';', '}']",
		yyXError{358, -1}: "expected one of [';', '}']",
		yyXError{359, -1}: "expected one of [';', '}']",
		yyXError{364, -1}: "expected one of [';', '}']",
		yyXError{365, -1}: "expected one of [';', '}']",
		yyXError{367, -1}: "expected one of [';', '}']",
		yyXError{399, -1}: "expected one of [';', '}']",
		yyXError{414, -1}: "expected one of [';', '}']",
		yyXError{415, -1}: "expected one of [';', '}']",
		yyXError{417, -1}: "expected one of [';', '}']",
		yyXError{420, -1}: "expected one of [';', '}']",
		yyXError{435, -1}: "expected one of [';', '}']",
		yyXError{233, -1}: "expected one of [';', {]",
		yyXError{234, -1}: "expected one of [';', {]",
		yyXError{422, -1}: "expected one of [';', {]",
		yyXError{449, -1}: "expected one of [';', {]",
		yyXError{245, -1}: "expected one of ['}', case, default]",
		yyXError{277, -1}: "expected one of ['}', case, default]",
		yyXError{213, -1}: "expected optional comma or one of [!=, &&, &^, '%', '&', ')', '*', '+', ',', '-', '/', '<', '>', '^', '|', <-, <<, <=, ==, >=, >>, ||]",
		yyXError{327, -1}: "expected optional comma or one of [')', ',', ...]",
		yyXError{190, -1}: "expected optional comma or one of [')', ',']",
		yyXError{331, -1}: "expected optional comma or one of [')', ',']",
		yyXError{315, -1}: "expected optional comma or one of [',', '}']",
		yyXError{386, -1}: "expected optional comma or one of [',', '}']",
		yyXError{67, -1}:  "expected optional expression list or one of ['!', '&', '(', '*', '+', '-', ';', '[', '^', '}', <-, case, chan, default, func, identifier, interface, literal, map, struct]",
		yyXError{298, -1}: "expected optional expression or one of ['!', '&', '(', '*', '+', '-', ':', '[', ']', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{301, -1}: "expected optional expression or one of ['!', '&', '(', '*', '+', '-', '[', ']', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{443, -1}: "expected optional function body or one of [';', '{']",
		yyXError{444, -1}: "expected optional function body or one of [';', '{']",
		yyXError{120, -1}: "expected optional generic arguments or one of [!=, %=, &&, &=, &^, &^=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '[', ']', '^', '{', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||, «]",
		yyXError{52, -1}:  "expected optional generic arguments or one of [!=, %=, &&, &=, &^, &^=, '%', '&', '(', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '[', '^', '{', '|', '}', *=, ++, +=, --, -=, /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, |=, ||, «]",
		yyXError{280, -1}: "expected optional generic arguments or one of [!=, &&, &^, '%', '&', '(', '*', '+', '-', '.', '/', '<', '>', '[', '^', '{', '|', :=, <-, <<, <=, ==, >=, >>, {, ||, «]",
		yyXError{112, -1}: "expected optional generic arguments or one of ['(', ')', ',', ':', ';', '=', ']', '{', '}', ..., :=, case, default, literal, {, «, »]",
		yyXError{219, -1}: "expected optional generic parameters or type or one of ['(', '*', '[', <-, chan, func, identifier, interface, map, struct, «]",
		yyXError{59, -1}:  "expected optional identifier or one of [';', '}', case, default, identifier]",
		yyXError{60, -1}:  "expected optional identifier or one of [';', '}', case, default, identifier]",
		yyXError{181, -1}: "expected optional result or one of ['(', ')', '*', ',', ':', ';', '=', '[', ']', '{', '}', ..., :=, <-, case, chan, default, func, identifier, interface, literal, map, struct, {, »]",
		yyXError{440, -1}: "expected optional result or one of ['(', '*', '[', '{', <-, chan, func, identifier, interface, map, struct, {]",
		yyXError{24, -1}:  "expected optional semicolon or one of [')', ';']",
		yyXError{96, -1}:  "expected optional semicolon or one of [')', ';']",
		yyXError{226, -1}: "expected optional semicolon or one of [')', ';']",
		yyXError{499, -1}: "expected optional semicolon or one of [')', ';']",
		yyXError{346, -1}: "expected optional semicolon or one of [';', '}']",
		yyXError{412, -1}: "expected optional semicolon or one of [';', '}']",
		yyXError{454, -1}: "expected optional simple statement or one of ['!', '&', '(', '*', '+', '-', ';', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{437, -1}: "expected optional simple statement or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct, {]",
		yyXError{456, -1}: "expected optional simple statement or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct, {]",
		yyXError{343, -1}: "expected optional tag or one of [';', '}', literal]",
		yyXError{354, -1}: "expected optional tag or one of [';', '}', literal]",
		yyXError{357, -1}: "expected optional tag or one of [';', '}', literal]",
		yyXError{360, -1}: "expected optional tag or one of [';', '}', literal]",
		yyXError{363, -1}: "expected optional tag or one of [';', '}', literal]",
		yyXError{366, -1}: "expected optional tag or one of [';', '}', literal]",
		yyXError{180, -1}: "expected parameter declaration list or one of ['(', ')', '*', '[', ..., <-, chan, func, identifier, interface, map, struct]",
		yyXError{182, -1}: "expected parameter declaration list or type or one of ['(', ')', '*', '[', ..., <-, chan, func, identifier, interface, map, struct]",
		yyXError{192, -1}: "expected parameter declaration or one of ['(', ')', '*', '[', ..., <-, chan, func, identifier, interface, map, struct]",
		yyXError{281, -1}: "expected primary expression or one of ['(', '*', '[', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{352, -1}: "expected qualified identifier or identifier",
		yyXError{361, -1}: "expected qualified identifier or identifier",
		yyXError{342, -1}: "expected qualified identifier or one of ['(', identifier]",
		yyXError{344, -1}: "expected qualified identifier or one of ['*', identifier]",
		yyXError{0, -1}:   "expected source file or package",
		yyXError{337, -1}: "expected struct field declaration list or one of ['(', '*', '}', identifier]",
		yyXError{339, -1}: "expected struct field declaration list or one of ['(', '*', identifier]",
		yyXError{347, -1}: "expected struct field declaration or one of ['(', '*', '}', identifier]",
		yyXError{237, -1}: "expected switch case/default clause list or one of ['}', case, default]",
		yyXError{240, -1}: "expected switch case/default clause list or one of [case, default]",
		yyXError{241, -1}: "expected switch case/default clause statement block or one of ['}', case, default]",
		yyXError{121, -1}: "expected type list or one of ['(', '*', '[', <-, chan, func, identifier, interface, map, struct]",
		yyXError{283, -1}: "expected type literal or one of ['*', '[', <-, chan, func, interface, map, struct]",
		yyXError{81, -1}:  "expected type literal or unary expression or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{187, -1}: "expected type or one of ['(', ')', '*', ',', '.', '[', ..., <-, chan, func, identifier, interface, map, struct, «]",
		yyXError{492, -1}: "expected type or one of ['(', ')', '*', ',', ';', '=', '[', '}', <-, case, chan, default, func, identifier, interface, map, struct]",
		yyXError{92, -1}:  "expected type or one of ['(', '*', ',', '=', '[', <-, chan, func, identifier, interface, map, struct]",
		yyXError{341, -1}: "expected type or one of ['(', '*', ',', '[', <-, chan, func, identifier, interface, map, struct]",
		yyXError{35, -1}:  "expected type or one of ['(', '*', '[', <-, chan, func, identifier, interface, map, struct]",
		yyXError{105, -1}: "expected type or one of ['(', '*', '[', <-, chan, func, identifier, interface, map, struct]",
		yyXError{106, -1}: "expected type or one of ['(', '*', '[', <-, chan, func, identifier, interface, map, struct]",
		yyXError{126, -1}: "expected type or one of ['(', '*', '[', <-, chan, func, identifier, interface, map, struct]",
		yyXError{186, -1}: "expected type or one of ['(', '*', '[', <-, chan, func, identifier, interface, map, struct]",
		yyXError{197, -1}: "expected type or one of ['(', '*', '[', <-, chan, func, identifier, interface, map, struct]",
		yyXError{221, -1}: "expected type or one of ['(', '*', '[', <-, chan, func, identifier, interface, map, struct]",
		yyXError{306, -1}: "expected type or one of ['(', '*', '[', <-, chan, func, identifier, interface, map, struct]",
		yyXError{378, -1}: "expected type or one of ['(', '*', '[', <-, chan, func, identifier, interface, map, struct]",
		yyXError{405, -1}: "expected type or one of ['(', '*', '[', <-, chan, func, identifier, interface, map, struct]",
		yyXError{407, -1}: "expected type or one of ['(', '*', '[', <-, chan, func, identifier, interface, map, struct]",
		yyXError{505, -1}: "expected type or one of ['(', '*', '[', <-, chan, func, identifier, interface, map, struct]",
		yyXError{508, -1}: "expected type or one of ['(', '*', '[', <-, chan, func, identifier, interface, map, struct]",
		yyXError{514, -1}: "expected type or one of ['(', '*', '[', <-, chan, func, identifier, interface, map, struct]",
		yyXError{516, -1}: "expected type or one of ['(', '*', '[', <-, chan, func, identifier, interface, map, struct]",
		yyXError{518, -1}: "expected type or one of ['(', '*', '[', <-, chan, func, identifier, interface, map, struct]",
		yyXError{217, -1}: "expected type specification list or one of [')', identifier]",
		yyXError{80, -1}:  "expected type specification or one of ['(', identifier]",
		yyXError{228, -1}: "expected type specification or one of [')', identifier]",
		yyXError{84, -1}:  "expected unary expression or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{85, -1}:  "expected unary expression or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{86, -1}:  "expected unary expression or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{87, -1}:  "expected unary expression or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{88, -1}:  "expected unary expression or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{89, -1}:  "expected unary expression or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{93, -1}:  "expected variable specification list or one of [')', identifier]",
		yyXError{90, -1}:  "expected variable specification or one of ['(', identifier]",
		yyXError{98, -1}:  "expected variable specification or one of [')', identifier]",
		yyXError{438, -1}: "expected {",
		yyXError{448, -1}: "expected {",
		yyXError{457, -1}: "expected {",
	}

	yyParseTab = [520][]uint16{
		// 0
		{162: 296, 171: 298, 299, 297},
		{99: 295},
		{2: 44, 44, 44, 44, 44, 44, 44, 11: 44, 13: 44, 16: 44, 44, 44, 44, 44, 44, 44, 32: 44, 34: 44, 44, 44, 38: 44, 79: 44, 81: 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 99: 44, 180: 327},
		{13: 325},
		{2: 174, 174, 174, 174, 174, 174, 174, 11: 174, 13: 174, 16: 174, 174, 174, 174, 174, 174, 174, 32: 174, 34: 174, 174, 174, 38: 174, 79: 174, 81: 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 174, 99: 174, 122: 174, 167: 300},
		// 5
		{2: 127, 127, 127, 127, 127, 127, 127, 11: 127, 13: 127, 16: 127, 127, 127, 127, 127, 127, 127, 32: 127, 34: 127, 127, 127, 38: 127, 79: 127, 81: 127, 127, 127, 127, 127, 127, 127, 127, 127, 127, 127, 127, 127, 127, 99: 127, 122: 301, 166: 302},
		{2: 305, 13: 304, 16: 190, 32: 190, 34: 190, 190, 190, 75: 307, 118: 308, 128: 306},
		{303},
		{2: 173, 173, 173, 173, 173, 173, 173, 11: 173, 13: 173, 16: 173, 173, 173, 173, 173, 173, 173, 32: 173, 34: 173, 173, 173, 38: 173, 79: 173, 81: 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 173, 99: 173, 122: 173},
		{189, 189, 9: 189, 189, 16: 189, 32: 189, 34: 189, 189, 189},
		// 10
		{12: 318, 304, 16: 190, 32: 190, 34: 190, 190, 190, 75: 307, 118: 308, 128: 320, 168: 319},
		{181},
		{16: 313, 32: 309, 34: 310, 311, 312, 41: 316},
		{16: 313, 32: 309, 34: 310, 311, 312, 41: 314},
		{275, 275, 275, 275, 275, 275, 275, 275, 275, 275, 275, 275, 275, 14: 275, 275, 23: 275, 275, 33: 275, 37: 275, 39: 275, 275, 47: 275, 275, 275, 275, 275, 275, 275, 275, 275, 275, 275, 275, 275, 275, 62: 275, 275, 275, 275, 275, 275, 275, 275, 275, 275, 275, 275, 275, 275, 79: 275},
		// 15
		{274, 274, 274, 274, 274, 274, 274, 274, 274, 274, 274, 274, 274, 14: 274, 274, 23: 274, 274, 33: 274, 37: 274, 39: 274, 274, 47: 274, 274, 274, 274, 274, 274, 274, 274, 274, 274, 274, 274, 274, 274, 62: 274, 274, 274, 274, 274, 274, 274, 274, 274, 274, 274, 274, 274, 274, 79: 274},
		{273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 14: 273, 273, 23: 273, 273, 33: 273, 37: 273, 39: 273, 273, 47: 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 62: 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 79: 273},
		{272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 14: 272, 272, 23: 272, 272, 33: 272, 37: 272, 39: 272, 272, 47: 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 62: 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 79: 272},
		{271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 14: 271, 271, 23: 271, 271, 33: 271, 37: 271, 39: 271, 271, 47: 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 62: 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 79: 271},
		{179, 12: 179, 79: 315},
		// 20
		{177, 12: 177},
		{180, 12: 180, 79: 317},
		{178, 12: 178},
		{183},
		{322, 12: 114, 114: 321},
		// 25
		{176, 12: 176},
		{12: 324},
		{12: 113, 304, 16: 190, 32: 190, 34: 190, 190, 190, 75: 307, 118: 308, 128: 323},
		{175, 12: 175},
		{182},
		// 30
		{326},
		{2: 147, 147, 147, 147, 147, 147, 147, 11: 147, 13: 147, 16: 147, 147, 147, 147, 147, 147, 147, 32: 147, 34: 147, 147, 147, 38: 147, 79: 147, 81: 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 99: 147, 122: 147},
		{2: 344, 376, 380, 381, 382, 383, 384, 11: 328, 13: 347, 16: 313, 330, 340, 342, 343, 331, 366, 25: 332, 377, 346, 378, 333, 334, 335, 309, 34: 310, 311, 312, 38: 379, 41: 345, 351, 349, 348, 350, 337, 61: 338, 77: 329, 79: 373, 81: 375, 341, 354, 336, 355, 356, 357, 339, 359, 360, 362, 352, 367, 385, 353, 364, 99: 294, 101: 368, 358, 361, 106: 363, 108: 372, 365, 370, 371, 165: 369, 179: 374},
		{2: 344, 376, 380, 381, 382, 383, 384, 11: 328, 13: 415, 16: 313, 330, 397, 342, 343, 331, 366, 25: 332, 377, 346, 378, 333, 334, 335, 309, 34: 310, 311, 312, 807, 379, 40: 809, 345, 351, 349, 348, 350, 337, 61: 808},
		{14: 423, 24: 805, 33: 806, 62: 754, 756, 755, 757, 758, 759, 760, 761, 762, 763, 764},
		// 35
		{2: 400, 401, 11: 328, 13: 399, 17: 330, 397, 342, 343, 331, 366, 25: 402, 403, 404, 405, 406, 408, 409, 78: 407, 80: 802, 181: 803},
		{17: 800},
		{2: 26, 12: 26, 14: 26, 250, 23: 26, 26, 33: 26, 37: 26, 39: 250},
		{2: 22, 12: 22, 14: 22, 249, 23: 22, 22, 33: 22, 37: 22, 39: 249},
		{2: 21, 12: 21, 14: 21, 248, 23: 21, 21, 33: 21, 37: 21, 39: 248},
		// 40
		{2: 20, 12: 20, 14: 20, 247, 23: 20, 20, 33: 20, 37: 20, 39: 247},
		{2: 785, 13: 386, 98: 787, 126: 786},
		{231, 231, 3: 231, 231, 231, 231, 231, 231, 231, 231, 12: 231, 14: 231, 231, 23: 231, 231, 33: 231, 37: 231, 40: 231, 47: 231, 231, 231, 231, 231, 231, 231, 231, 231, 231, 231, 231, 231, 231, 62: 231, 231, 231, 231, 231, 231, 231, 231, 231, 231, 231, 231, 231},
		{110, 110, 3: 427, 426, 428, 429, 433, 444, 110, 110, 14: 208, 110, 24: 208, 33: 208, 47: 425, 430, 431, 432, 434, 435, 436, 437, 438, 439, 440, 441, 442, 443, 62: 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 783, 784},
		{106, 2: 344, 376, 380, 381, 382, 383, 384, 11: 328, 13: 415, 15: 106, 313, 330, 397, 342, 343, 331, 366, 25: 332, 377, 346, 378, 333, 334, 335, 309, 34: 310, 311, 312, 38: 379, 41: 345, 351, 349, 348, 350, 337, 61: 338, 77: 742, 95: 353, 528, 107: 744, 130: 746, 163: 745, 174: 743},
		// 45
		{2: 475, 13: 121, 97: 474, 105: 735, 119: 473, 175: 734},
		{106, 2: 344, 376, 380, 381, 382, 383, 384, 11: 328, 13: 415, 15: 106, 313, 330, 397, 342, 343, 331, 366, 25: 332, 377, 346, 378, 333, 334, 335, 309, 34: 310, 311, 312, 38: 379, 41: 345, 351, 349, 348, 350, 337, 61: 338, 77: 329, 95: 353, 528, 107: 717, 137: 718},
		{15: 630, 39: 631, 104: 704},
		{11: 700},
		{2: 344, 376, 380, 381, 382, 383, 384, 11: 328, 13: 415, 16: 313, 330, 397, 342, 343, 331, 366, 25: 332, 377, 346, 378, 333, 334, 335, 309, 34: 310, 311, 312, 38: 379, 41: 345, 697, 349, 348, 350, 337, 61: 696},
		// 50
		{151, 151, 151, 151, 151, 151, 151, 151, 151, 151, 151, 151, 151, 14: 151, 151, 23: 151, 151, 33: 151, 37: 151, 39: 151, 151, 47: 151, 151, 151, 151, 151, 151, 151, 151, 151, 151, 151, 151, 151, 151, 62: 151, 151, 151, 151, 151, 151, 151, 151, 151, 151, 151, 151, 151, 151},
		{2: 24, 12: 24, 14: 24, 150, 23: 24, 24, 33: 24, 37: 24, 39: 150, 152: 692},
		{194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 14: 194, 23: 690, 194, 33: 194, 39: 194, 47: 194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 62: 194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 97: 416, 121: 417},
		{138, 138, 138, 138, 138, 138, 138, 138, 138, 138, 138, 138, 138, 14: 138, 138, 23: 138, 138, 33: 138, 37: 138, 39: 138, 138, 47: 138, 138, 138, 138, 138, 138, 138, 138, 138, 138, 138, 138, 138, 138, 62: 138, 138, 138, 138, 138, 138, 138, 138, 138, 138, 138, 138, 138, 138},
		{15: 630, 39: 631, 104: 676, 123: 677},
		// 55
		{9, 9, 585, 9, 9, 9, 9, 9, 9, 9, 9, 588, 9, 14: 9, 9, 23: 9, 9, 33: 9, 37: 9, 39: 586, 9, 47: 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 62: 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 672, 113: 590, 134: 589},
		{2: 507},
		{15: 116, 153: 670},
		{111, 111, 9: 111, 111, 15: 111},
		{190, 190, 9: 190, 190, 13: 304, 118: 669},
		// 60
		{190, 190, 9: 190, 190, 13: 304, 118: 668},
		{2: 344, 376, 380, 381, 382, 383, 384, 11: 328, 13: 415, 16: 313, 330, 397, 342, 343, 331, 366, 25: 332, 377, 346, 378, 333, 334, 335, 309, 34: 310, 311, 312, 38: 379, 41: 345, 351, 349, 348, 350, 337, 61: 667},
		{91, 91, 9: 91, 91},
		{90, 90, 9: 90, 90},
		{2: 344, 376, 380, 381, 382, 383, 384, 11: 328, 13: 415, 16: 313, 330, 397, 342, 343, 331, 366, 25: 332, 377, 346, 378, 333, 334, 335, 309, 34: 310, 311, 312, 38: 379, 41: 345, 351, 349, 348, 350, 337, 61: 666},
		// 65
		{13: 665},
		{86, 86, 9: 86, 86},
		{206, 206, 344, 376, 380, 381, 382, 383, 384, 206, 206, 328, 13: 415, 16: 313, 330, 397, 342, 343, 331, 366, 25: 332, 377, 346, 378, 333, 334, 335, 309, 34: 310, 311, 312, 38: 379, 41: 345, 351, 349, 348, 350, 337, 61: 413, 77: 663, 161: 664},
		{84, 84, 9: 84, 84},
		{83, 83, 9: 83, 83},
		// 70
		{82, 82, 9: 82, 82},
		{15: 630, 39: 631, 104: 632},
		{106, 2: 344, 376, 380, 381, 382, 383, 384, 11: 328, 13: 415, 15: 106, 313, 330, 397, 342, 343, 331, 366, 25: 332, 377, 346, 378, 333, 334, 335, 309, 34: 310, 311, 312, 38: 379, 41: 345, 351, 349, 348, 350, 337, 61: 338, 77: 329, 95: 353, 528, 107: 529, 178: 530},
		{50},
		{49},
		// 75
		{48},
		{47},
		{46},
		{45},
		{527},
		// 80
		{2: 512, 13: 514, 131: 513},
		{2: 344, 376, 380, 381, 382, 383, 384, 11: 328, 13: 415, 16: 313, 330, 397, 342, 343, 331, 366, 25: 332, 377, 346, 378, 333, 334, 335, 309, 34: 310, 311, 312, 38: 379, 41: 345, 505, 349, 348, 350, 506},
		{2: 25, 12: 25, 14: 25, 23: 25, 25, 33: 25, 37: 25},
		{2: 23, 12: 23, 14: 23, 23: 23, 23, 33: 23, 37: 23},
		{2: 344, 376, 380, 381, 382, 383, 384, 11: 328, 13: 415, 16: 313, 330, 397, 342, 343, 331, 366, 25: 332, 377, 346, 378, 333, 334, 335, 309, 34: 310, 311, 312, 38: 379, 41: 345, 351, 349, 348, 350, 504},
		// 85
		{2: 344, 376, 380, 381, 382, 383, 384, 11: 328, 13: 415, 16: 313, 330, 397, 342, 343, 331, 366, 25: 332, 377, 346, 378, 333, 334, 335, 309, 34: 310, 311, 312, 38: 379, 41: 345, 351, 349, 348, 350, 503},
		{2: 344, 376, 380, 381, 382, 383, 384, 11: 328, 13: 415, 16: 313, 330, 397, 342, 343, 331, 366, 25: 332, 377, 346, 378, 333, 334, 335, 309, 34: 310, 311, 312, 38: 379, 41: 345, 351, 349, 348, 350, 502},
		{2: 344, 376, 380, 381, 382, 383, 384, 11: 328, 13: 415, 16: 313, 330, 397, 342, 343, 331, 366, 25: 332, 377, 346, 378, 333, 334, 335, 309, 34: 310, 311, 312, 38: 379, 41: 345, 351, 349, 348, 350, 501},
		{2: 344, 376, 380, 381, 382, 383, 384, 11: 328, 13: 415, 16: 313, 330, 397, 342, 343, 331, 366, 25: 332, 377, 346, 378, 333, 334, 335, 309, 34: 310, 311, 312, 38: 379, 41: 345, 351, 349, 348, 350, 500},
		{2: 344, 376, 380, 381, 382, 383, 384, 11: 328, 13: 415, 16: 313, 330, 397, 342, 343, 331, 366, 25: 332, 377, 346, 378, 333, 334, 335, 309, 34: 310, 311, 312, 38: 379, 41: 345, 351, 349, 348, 350, 499},
		// 90
		{2: 388, 13: 386, 98: 387, 132: 389},
		{188, 188, 188, 188, 9: 188, 188, 188, 188, 188, 188, 17: 188, 188, 188, 188, 188, 188, 24: 188, 76: 188},
		{2: 400, 401, 11: 328, 13: 399, 398, 17: 330, 397, 342, 343, 331, 366, 24: 410, 402, 403, 404, 405, 406, 408, 409, 78: 407, 80: 411},
		{12: 390, 386, 98: 387, 132: 392, 184: 391},
		{6, 6, 9: 6, 6},
		// 95
		{8, 8, 9: 8, 8},
		{393, 12: 114, 114: 394},
		{2, 12: 2},
		{12: 113, 386, 98: 387, 132: 396},
		{12: 395},
		// 100
		{7, 7, 9: 7, 7},
		{1, 12: 1},
		{2: 475, 97: 474, 105: 476, 119: 473},
		{13: 472},
		{126, 126, 126, 9: 126, 126, 12: 126, 14: 126, 126, 126, 23: 126, 126, 33: 126, 37: 126, 39: 126, 126, 75: 470, 126, 97: 126},
		// 105
		{2: 400, 401, 11: 328, 13: 399, 17: 330, 397, 342, 343, 331, 366, 25: 402, 403, 404, 405, 406, 408, 409, 78: 407, 80: 468},
		{2: 400, 401, 11: 328, 13: 399, 17: 330, 397, 342, 343, 331, 366, 25: 402, 403, 404, 405, 406, 408, 409, 78: 407, 80: 467},
		{40, 40, 40, 9: 40, 40, 12: 40, 14: 40, 40, 40, 23: 40, 40, 33: 40, 37: 40, 39: 40, 40, 76: 40},
		{39, 39, 39, 9: 39, 39, 12: 39, 14: 39, 39, 39, 23: 39, 39, 33: 39, 37: 39, 39: 39, 39, 76: 39},
		{38, 38, 38, 9: 38, 38, 12: 38, 14: 38, 38, 38, 23: 38, 38, 33: 38, 37: 38, 39: 38, 38, 76: 38},
		// 110
		{37, 37, 37, 9: 37, 37, 12: 37, 14: 37, 37, 37, 23: 37, 37, 33: 37, 37: 37, 39: 37, 37, 76: 37},
		{36, 36, 36, 9: 36, 36, 12: 36, 14: 36, 36, 36, 23: 36, 36, 33: 36, 37: 36, 39: 36, 36, 76: 36},
		{194, 194, 194, 9: 194, 194, 12: 194, 14: 194, 194, 194, 23: 194, 194, 33: 194, 37: 194, 39: 194, 194, 76: 194, 97: 416, 121: 466},
		{34, 34, 34, 9: 34, 34, 12: 34, 14: 34, 34, 34, 23: 34, 34, 33: 34, 37: 34, 39: 34, 34, 76: 34},
		{33, 33, 33, 9: 33, 33, 12: 33, 14: 33, 33, 33, 23: 33, 33, 33: 33, 37: 33, 39: 33, 33, 76: 33},
		// 115
		{2: 344, 376, 380, 381, 382, 383, 384, 11: 328, 13: 415, 16: 313, 330, 397, 342, 343, 331, 366, 25: 332, 377, 346, 378, 333, 334, 335, 309, 34: 310, 311, 312, 38: 379, 41: 345, 351, 349, 348, 350, 337, 61: 413, 77: 465},
		{4, 4, 9: 4, 4, 12: 4, 24: 412},
		{2: 344, 376, 380, 381, 382, 383, 384, 11: 328, 13: 415, 16: 313, 330, 397, 342, 343, 331, 366, 25: 332, 377, 346, 378, 333, 334, 335, 309, 34: 310, 311, 312, 38: 379, 41: 345, 351, 349, 348, 350, 337, 61: 413, 77: 414},
		{208, 208, 3: 427, 426, 428, 429, 433, 444, 208, 208, 12: 208, 14: 208, 208, 47: 425, 430, 431, 432, 434, 435, 436, 437, 438, 439, 440, 441, 442, 443},
		{3, 3, 9: 3, 3, 12: 3, 14: 423},
		// 120
		{194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 14: 194, 194, 23: 194, 194, 33: 194, 37: 194, 39: 194, 194, 47: 194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 62: 194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 97: 416, 121: 417},
		{2: 400, 401, 11: 328, 13: 399, 17: 330, 397, 342, 343, 331, 366, 25: 402, 403, 404, 405, 406, 408, 409, 78: 407, 80: 419, 182: 418},
		{148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 14: 148, 148, 23: 148, 148, 33: 148, 37: 148, 39: 148, 148, 47: 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 62: 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148},
		{14: 421, 76: 420},
		{14: 29, 76: 29},
		// 125
		{193, 193, 193, 193, 193, 193, 193, 193, 193, 193, 193, 193, 193, 14: 193, 193, 193, 23: 193, 193, 33: 193, 37: 193, 39: 193, 193, 47: 193, 193, 193, 193, 193, 193, 193, 193, 193, 193, 193, 193, 193, 193, 62: 193, 193, 193, 193, 193, 193, 193, 193, 193, 193, 193, 193, 193, 193, 193},
		{2: 400, 401, 11: 328, 13: 399, 17: 330, 397, 342, 343, 331, 366, 25: 402, 403, 404, 405, 406, 408, 409, 78: 407, 80: 422},
		{14: 28, 76: 28},
		{2: 344, 376, 380, 381, 382, 383, 384, 11: 328, 13: 415, 16: 313, 330, 397, 342, 343, 331, 366, 25: 332, 377, 346, 378, 333, 334, 335, 309, 34: 310, 311, 312, 38: 379, 41: 345, 351, 349, 348, 350, 337, 61: 424},
		{207, 207, 3: 427, 426, 428, 429, 433, 444, 207, 207, 12: 207, 14: 207, 207, 24: 207, 33: 207, 47: 425, 430, 431, 432, 434, 435, 436, 437, 438, 439, 440, 441, 442, 443, 62: 207, 207, 207, 207, 207, 207, 207, 207, 207, 207, 207},
		// 130
		{2: 344, 376, 380, 381, 382, 383, 384, 11: 328, 13: 415, 16: 313, 330, 397, 342, 343, 331, 366, 25: 332, 377, 346, 378, 333, 334, 335, 309, 34: 310, 311, 312, 38: 379, 41: 345, 351, 349, 348, 350, 337, 61: 464},
		{2: 344, 376, 380, 381, 382, 383, 384, 11: 328, 13: 415, 16: 313, 330, 397, 342, 343, 331, 366, 25: 332, 377, 346, 378, 333, 334, 335, 309, 34: 310, 311, 312, 38: 379, 41: 345, 351, 349, 348, 350, 337, 61: 463},
		{2: 344, 376, 380, 381, 382, 383, 384, 11: 328, 13: 415, 16: 313, 330, 397, 342, 343, 331, 366, 25: 332, 377, 346, 378, 333, 334, 335, 309, 34: 310, 311, 312, 38: 379, 41: 345, 351, 349, 348, 350, 337, 61: 462},
		{2: 344, 376, 380, 381, 382, 383, 384, 11: 328, 13: 415, 16: 313, 330, 397, 342, 343, 331, 366, 25: 332, 377, 346, 378, 333, 334, 335, 309, 34: 310, 311, 312, 38: 379, 41: 345, 351, 349, 348, 350, 337, 61: 461},
		{2: 344, 376, 380, 381, 382, 383, 384, 11: 328, 13: 415, 16: 313, 330, 397, 342, 343, 331, 366, 25: 332, 377, 346, 378, 333, 334, 335, 309, 34: 310, 311, 312, 38: 379, 41: 345, 351, 349, 348, 350, 337, 61: 460},
		// 135
		{2: 344, 376, 380, 381, 382, 383, 384, 11: 328, 13: 415, 16: 313, 330, 397, 342, 343, 331, 366, 25: 332, 377, 346, 378, 333, 334, 335, 309, 34: 310, 311, 312, 38: 379, 41: 345, 351, 349, 348, 350, 337, 61: 459},
		{2: 344, 376, 380, 381, 382, 383, 384, 11: 328, 13: 415, 16: 313, 330, 397, 342, 343, 331, 366, 25: 332, 377, 346, 378, 333, 334, 335, 309, 34: 310, 311, 312, 38: 379, 41: 345, 351, 349, 348, 350, 337, 61: 458},
		{2: 344, 376, 380, 381, 382, 383, 384, 11: 328, 13: 415, 16: 313, 330, 397, 342, 343, 331, 366, 25: 332, 377, 346, 378, 333, 334, 335, 309, 34: 310, 311, 312, 38: 379, 41: 345, 351, 349, 348, 350, 337, 61: 457},
		{2: 344, 376, 380, 381, 382, 383, 384, 11: 328, 13: 415, 16: 313, 330, 397, 342, 343, 331, 366, 25: 332, 377, 346, 378, 333, 334, 335, 309, 34: 310, 311, 312, 38: 379, 41: 345, 351, 349, 348, 350, 337, 61: 456},
		{2: 344, 376, 380, 381, 382, 383, 384, 11: 328, 13: 415, 16: 313, 330, 397, 342, 343, 331, 366, 25: 332, 377, 346, 378, 333, 334, 335, 309, 34: 310, 311, 312, 38: 379, 41: 345, 351, 349, 348, 350, 337, 61: 455},
		// 140
		{2: 344, 376, 380, 381, 382, 383, 384, 11: 328, 13: 415, 16: 313, 330, 397, 342, 343, 331, 366, 25: 332, 377, 346, 378, 333, 334, 335, 309, 34: 310, 311, 312, 38: 379, 41: 345, 351, 349, 348, 350, 337, 61: 454},
		{2: 344, 376, 380, 381, 382, 383, 384, 11: 328, 13: 415, 16: 313, 330, 397, 342, 343, 331, 366, 25: 332, 377, 346, 378, 333, 334, 335, 309, 34: 310, 311, 312, 38: 379, 41: 345, 351, 349, 348, 350, 337, 61: 453},
		{2: 344, 376, 380, 381, 382, 383, 384, 11: 328, 13: 415, 16: 313, 330, 397, 342, 343, 331, 366, 25: 332, 377, 346, 378, 333, 334, 335, 309, 34: 310, 311, 312, 38: 379, 41: 345, 351, 349, 348, 350, 337, 61: 452},
		{2: 344, 376, 380, 381, 382, 383, 384, 11: 328, 13: 415, 16: 313, 330, 397, 342, 343, 331, 366, 25: 332, 377, 346, 378, 333, 334, 335, 309, 34: 310, 311, 312, 38: 379, 41: 345, 351, 349, 348, 350, 337, 61: 451},
		{2: 344, 376, 380, 381, 382, 383, 384, 11: 328, 13: 415, 16: 313, 330, 397, 342, 343, 331, 366, 25: 332, 377, 346, 378, 333, 334, 335, 309, 34: 310, 311, 312, 38: 379, 41: 345, 351, 349, 348, 350, 337, 61: 450},
		// 145
		{2: 344, 376, 380, 381, 382, 383, 384, 11: 328, 13: 415, 16: 313, 330, 397, 342, 343, 331, 366, 25: 332, 377, 346, 378, 333, 334, 335, 309, 34: 310, 311, 312, 38: 379, 41: 345, 351, 349, 348, 350, 337, 61: 449},
		{2: 344, 376, 380, 381, 382, 383, 384, 11: 328, 13: 415, 16: 313, 330, 397, 342, 343, 331, 366, 25: 332, 377, 346, 378, 333, 334, 335, 309, 34: 310, 311, 312, 38: 379, 41: 345, 351, 349, 348, 350, 337, 61: 448},
		{2: 344, 376, 380, 381, 382, 383, 384, 11: 328, 13: 415, 16: 313, 330, 397, 342, 343, 331, 366, 25: 332, 377, 346, 378, 333, 334, 335, 309, 34: 310, 311, 312, 38: 379, 41: 345, 351, 349, 348, 350, 337, 61: 447},
		{2: 344, 376, 380, 381, 382, 383, 384, 11: 328, 13: 415, 16: 313, 330, 397, 342, 343, 331, 366, 25: 332, 377, 346, 378, 333, 334, 335, 309, 34: 310, 311, 312, 38: 379, 41: 345, 351, 349, 348, 350, 337, 61: 446},
		{2: 344, 376, 380, 381, 382, 383, 384, 11: 328, 13: 415, 16: 313, 330, 397, 342, 343, 331, 366, 25: 332, 377, 346, 378, 333, 334, 335, 309, 34: 310, 311, 312, 38: 379, 41: 345, 351, 349, 348, 350, 337, 61: 445},
		// 150
		{211, 211, 3: 427, 426, 428, 429, 433, 211, 211, 211, 12: 211, 14: 211, 211, 23: 211, 211, 33: 211, 37: 211, 40: 211, 47: 425, 430, 431, 432, 434, 435, 436, 437, 438, 439, 440, 441, 442, 443, 62: 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211},
		{212, 212, 3: 212, 212, 212, 212, 212, 212, 212, 212, 12: 212, 14: 212, 212, 23: 212, 212, 33: 212, 37: 212, 40: 212, 47: 212, 212, 212, 212, 212, 212, 212, 212, 212, 212, 212, 212, 212, 212, 62: 212, 212, 212, 212, 212, 212, 212, 212, 212, 212, 212, 212, 212},
		{213, 213, 3: 427, 426, 428, 429, 433, 213, 213, 213, 12: 213, 14: 213, 213, 23: 213, 213, 33: 213, 37: 213, 40: 213, 47: 425, 430, 431, 432, 434, 435, 436, 437, 438, 439, 440, 441, 213, 443, 62: 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213},
		{214, 214, 3: 427, 426, 428, 429, 433, 214, 214, 214, 12: 214, 14: 214, 214, 23: 214, 214, 33: 214, 37: 214, 40: 214, 47: 425, 430, 214, 214, 434, 214, 436, 214, 214, 214, 440, 214, 214, 443, 62: 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214},
		{215, 215, 3: 215, 215, 215, 215, 215, 215, 215, 215, 12: 215, 14: 215, 215, 23: 215, 215, 33: 215, 37: 215, 40: 215, 47: 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 62: 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215},
		// 155
		{216, 216, 3: 427, 426, 428, 429, 433, 216, 216, 216, 12: 216, 14: 216, 216, 23: 216, 216, 33: 216, 37: 216, 40: 216, 47: 425, 430, 216, 216, 434, 216, 436, 216, 216, 216, 440, 216, 216, 443, 62: 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216},
		{217, 217, 3: 427, 426, 428, 429, 433, 217, 217, 217, 12: 217, 14: 217, 217, 23: 217, 217, 33: 217, 37: 217, 40: 217, 47: 425, 430, 217, 217, 434, 217, 436, 217, 217, 217, 440, 217, 217, 443, 62: 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 217},
		{218, 218, 3: 427, 426, 428, 429, 433, 218, 218, 218, 12: 218, 14: 218, 218, 23: 218, 218, 33: 218, 37: 218, 40: 218, 47: 425, 430, 218, 218, 434, 218, 436, 218, 218, 218, 440, 218, 218, 443, 62: 218, 218, 218, 218, 218, 218, 218, 218, 218, 218, 218, 218, 218},
		{219, 219, 3: 219, 219, 219, 219, 219, 219, 219, 219, 12: 219, 14: 219, 219, 23: 219, 219, 33: 219, 37: 219, 40: 219, 47: 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 62: 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219},
		{220, 220, 3: 427, 426, 428, 429, 433, 220, 220, 220, 12: 220, 14: 220, 220, 23: 220, 220, 33: 220, 37: 220, 40: 220, 47: 425, 430, 431, 432, 434, 220, 436, 437, 438, 439, 440, 441, 220, 443, 62: 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 220},
		// 160
		{221, 221, 3: 427, 426, 221, 221, 221, 221, 221, 221, 12: 221, 14: 221, 221, 23: 221, 221, 33: 221, 37: 221, 40: 221, 47: 425, 430, 221, 221, 221, 221, 436, 221, 221, 221, 440, 221, 221, 443, 62: 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221},
		{222, 222, 3: 427, 426, 222, 222, 222, 222, 222, 222, 12: 222, 14: 222, 222, 23: 222, 222, 33: 222, 37: 222, 40: 222, 47: 425, 430, 222, 222, 222, 222, 436, 222, 222, 222, 440, 222, 222, 443, 62: 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222},
		{223, 223, 3: 427, 426, 428, 429, 433, 223, 223, 223, 12: 223, 14: 223, 223, 23: 223, 223, 33: 223, 37: 223, 40: 223, 47: 425, 430, 223, 223, 434, 223, 436, 223, 223, 223, 440, 223, 223, 443, 62: 223, 223, 223, 223, 223, 223, 223, 223, 223, 223, 223, 223, 223},
		{224, 224, 3: 427, 426, 428, 429, 433, 224, 224, 224, 12: 224, 14: 224, 224, 23: 224, 224, 33: 224, 37: 224, 40: 224, 47: 425, 430, 224, 224, 434, 224, 436, 224, 224, 224, 440, 224, 224, 443, 62: 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224},
		{225, 225, 3: 225, 225, 225, 225, 225, 225, 225, 225, 12: 225, 14: 225, 225, 23: 225, 225, 33: 225, 37: 225, 40: 225, 47: 225, 225, 225, 225, 225, 225, 225, 225, 225, 225, 225, 225, 225, 225, 62: 225, 225, 225, 225, 225, 225, 225, 225, 225, 225, 225, 225, 225},
		// 165
		{226, 226, 3: 427, 426, 226, 226, 226, 226, 226, 226, 12: 226, 14: 226, 226, 23: 226, 226, 33: 226, 37: 226, 40: 226, 47: 425, 430, 226, 226, 226, 226, 436, 226, 226, 226, 440, 226, 226, 443, 62: 226, 226, 226, 226, 226, 226, 226, 226, 226, 226, 226, 226, 226},
		{227, 227, 3: 427, 426, 227, 227, 227, 227, 227, 227, 12: 227, 14: 227, 227, 23: 227, 227, 33: 227, 37: 227, 40: 227, 47: 425, 430, 227, 227, 227, 227, 436, 227, 227, 227, 440, 227, 227, 443, 62: 227, 227, 227, 227, 227, 227, 227, 227, 227, 227, 227, 227, 227},
		{228, 228, 3: 228, 228, 228, 228, 228, 228, 228, 228, 12: 228, 14: 228, 228, 23: 228, 228, 33: 228, 37: 228, 40: 228, 47: 228, 228, 228, 228, 228, 228, 228, 228, 228, 228, 228, 228, 228, 228, 62: 228, 228, 228, 228, 228, 228, 228, 228, 228, 228, 228, 228, 228},
		{229, 229, 3: 229, 229, 229, 229, 229, 229, 229, 229, 12: 229, 14: 229, 229, 23: 229, 229, 33: 229, 37: 229, 40: 229, 47: 229, 229, 229, 229, 229, 229, 229, 229, 229, 229, 229, 229, 229, 229, 62: 229, 229, 229, 229, 229, 229, 229, 229, 229, 229, 229, 229, 229},
		{230, 230, 3: 230, 230, 230, 230, 230, 230, 230, 230, 12: 230, 14: 230, 230, 23: 230, 230, 33: 230, 37: 230, 40: 230, 47: 230, 230, 230, 230, 230, 230, 230, 230, 230, 230, 230, 230, 230, 230, 62: 230, 230, 230, 230, 230, 230, 230, 230, 230, 230, 230, 230, 230},
		// 170
		{5, 5, 9: 5, 5, 12: 5, 14: 423},
		{35, 35, 35, 9: 35, 35, 12: 35, 14: 35, 35, 35, 23: 35, 35, 33: 35, 37: 35, 39: 35, 35, 76: 35},
		{41, 41, 41, 9: 41, 41, 12: 41, 14: 41, 41, 41, 23: 41, 41, 33: 41, 37: 41, 39: 41, 41, 76: 41},
		{12: 469},
		{42, 42, 42, 9: 42, 42, 12: 42, 14: 42, 42, 42, 23: 42, 42, 33: 42, 37: 42, 39: 42, 42, 76: 42},
		// 175
		{13: 471},
		{125, 125, 125, 9: 125, 125, 12: 125, 14: 125, 125, 125, 23: 125, 125, 33: 125, 37: 125, 39: 125, 125, 76: 125, 97: 125},
		{187, 187, 187, 187, 9: 187, 187, 187, 187, 187, 187, 17: 187, 187, 187, 187, 187, 187, 24: 187, 76: 187},
		{196, 196, 196, 9: 196, 196, 12: 196, 14: 196, 196, 196, 23: 196, 196, 33: 196, 37: 196, 39: 196, 196, 76: 196},
		{13: 386, 98: 496},
		// 180
		{2: 400, 401, 11: 328, 486, 482, 17: 330, 397, 342, 343, 331, 366, 25: 402, 403, 404, 405, 406, 408, 409, 37: 481, 78: 407, 80: 490, 129: 484, 140: 485},
		{119, 119, 477, 401, 9: 119, 119, 328, 119, 399, 119, 119, 119, 330, 397, 342, 343, 331, 366, 119, 119, 402, 403, 404, 405, 406, 408, 409, 33: 119, 37: 119, 39: 119, 119, 76: 119, 78: 407, 80: 479, 105: 478, 141: 480},
		{2: 400, 401, 11: 328, 486, 482, 17: 330, 397, 342, 343, 331, 366, 25: 402, 403, 404, 405, 406, 408, 409, 37: 481, 78: 407, 80: 483, 129: 484, 140: 485},
		{118, 118, 118, 9: 118, 118, 12: 118, 14: 118, 118, 118, 23: 118, 118, 33: 118, 37: 118, 39: 118, 118, 76: 118},
		{117, 117, 117, 9: 117, 117, 12: 117, 14: 117, 117, 117, 23: 117, 117, 33: 117, 37: 117, 39: 117, 117, 76: 117},
		// 185
		{112, 112, 112, 9: 112, 112, 12: 112, 14: 112, 112, 112, 23: 112, 112, 33: 112, 37: 112, 39: 112, 112, 76: 112},
		{2: 400, 401, 11: 328, 13: 399, 17: 330, 397, 342, 343, 331, 366, 25: 402, 403, 404, 405, 406, 408, 409, 78: 407, 80: 495},
		{2: 400, 401, 11: 328, 126, 399, 126, 17: 330, 397, 342, 343, 331, 366, 25: 402, 403, 404, 405, 406, 408, 409, 37: 492, 75: 470, 78: 407, 80: 493, 97: 126},
		{12: 469, 14: 143},
		{12: 142, 14: 142},
		// 190
		{12: 260, 14: 487, 112: 488},
		{140, 140, 140, 140, 9: 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 33: 140, 37: 140, 39: 140, 140, 76: 140},
		{2: 400, 401, 11: 328, 259, 482, 17: 330, 397, 342, 343, 331, 366, 25: 402, 403, 404, 405, 406, 408, 409, 37: 481, 78: 407, 80: 490, 129: 491},
		{12: 489},
		{139, 139, 139, 139, 9: 139, 139, 139, 139, 139, 139, 139, 139, 139, 139, 139, 139, 139, 139, 139, 139, 33: 139, 37: 139, 39: 139, 139, 76: 139},
		// 195
		{12: 143, 14: 143},
		{12: 141, 14: 141},
		{2: 400, 401, 11: 328, 13: 399, 17: 330, 397, 342, 343, 331, 366, 25: 402, 403, 404, 405, 406, 408, 409, 78: 407, 80: 494},
		{12: 144, 14: 144},
		{12: 145, 14: 145},
		// 200
		{12: 146, 14: 146},
		{14: 398, 76: 497},
		{2: 475, 105: 476, 119: 498},
		{195, 195, 195, 9: 195, 195, 12: 195, 14: 195, 195, 195, 23: 195, 195, 33: 195, 37: 195, 39: 195, 195, 76: 195},
		{10, 10, 3: 10, 10, 10, 10, 10, 10, 10, 10, 12: 10, 14: 10, 10, 23: 10, 10, 33: 10, 37: 10, 40: 10, 47: 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 62: 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10},
		// 205
		{11, 11, 3: 11, 11, 11, 11, 11, 11, 11, 11, 12: 11, 14: 11, 11, 23: 11, 11, 33: 11, 37: 11, 40: 11, 47: 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 62: 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11},
		{12, 12, 3: 12, 12, 12, 12, 12, 12, 12, 12, 12: 12, 14: 12, 12, 23: 12, 12, 33: 12, 37: 12, 40: 12, 47: 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 62: 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12},
		{13, 13, 3: 13, 13, 13, 13, 13, 13, 13, 13, 12: 13, 14: 13, 13, 23: 13, 13, 33: 13, 37: 13, 40: 13, 47: 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 62: 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13},
		{15, 15, 3: 15, 15, 15, 15, 15, 15, 15, 15, 12: 15, 14: 15, 15, 23: 15, 15, 33: 15, 37: 15, 40: 15, 47: 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 62: 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15},
		{16, 16, 3: 16, 16, 16, 16, 16, 16, 16, 16, 12: 16, 14: 16, 16, 23: 16, 16, 33: 16, 37: 16, 40: 16, 47: 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 62: 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16},
		// 210
		{2: 507, 12: 27, 14: 27, 23: 27, 27, 33: 27, 37: 27},
		{14, 14, 3: 14, 14, 14, 14, 14, 14, 14, 14, 12: 14, 14: 14, 14, 23: 14, 14, 33: 14, 37: 14, 40: 14, 47: 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 62: 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14},
		{2: 344, 376, 380, 381, 382, 383, 384, 11: 328, 13: 415, 16: 313, 330, 397, 342, 343, 331, 366, 25: 332, 377, 346, 378, 333, 334, 335, 309, 34: 310, 311, 312, 38: 379, 41: 345, 351, 349, 348, 350, 337, 61: 508},
		{3: 427, 426, 428, 429, 433, 444, 12: 260, 14: 509, 47: 425, 430, 431, 432, 434, 435, 436, 437, 438, 439, 440, 441, 442, 443, 112: 510},
		{12: 259},
		// 215
		{12: 511},
		{128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 14: 128, 128, 23: 128, 128, 33: 128, 37: 128, 39: 128, 128, 47: 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 62: 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128},
		{12: 520, 514, 131: 522, 183: 521},
		{30, 30, 9: 30, 30},
		{2: 192, 192, 11: 192, 13: 192, 17: 192, 192, 192, 192, 192, 192, 97: 515, 136: 516},
		// 220
		{13: 386, 98: 518},
		{2: 400, 401, 11: 328, 13: 399, 17: 330, 397, 342, 343, 331, 366, 25: 402, 403, 404, 405, 406, 408, 409, 78: 407, 80: 517},
		{19, 19, 9: 19, 19, 12: 19},
		{14: 398, 76: 519},
		{2: 191, 191, 11: 191, 13: 191, 17: 191, 191, 191, 191, 191, 191},
		// 225
		{32, 32, 9: 32, 32},
		{523, 12: 114, 114: 524},
		{18, 12: 18},
		{12: 113, 514, 131: 526},
		{12: 525},
		// 230
		{31, 31, 9: 31, 31},
		{17, 12: 17},
		{2: 43, 43, 43, 43, 43, 43, 43, 11: 43, 13: 43, 16: 43, 43, 43, 43, 43, 43, 43, 32: 43, 34: 43, 43, 43, 38: 43, 79: 43, 81: 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 99: 43},
		{105, 15: 105},
		{573, 15: 58},
		// 235
		{15: 54, 147: 531},
		{15: 532, 143: 533},
		{1: 534, 9: 69, 69, 155: 535},
		{53, 53, 9: 53, 53},
		{70, 70, 9: 70, 70},
		// 240
		{9: 537, 538, 144: 539, 540, 177: 536},
		{1: 571, 9: 537, 538, 144: 539, 572},
		{2: 344, 376, 380, 381, 382, 383, 384, 11: 328, 13: 415, 16: 313, 330, 397, 342, 343, 331, 366, 25: 332, 377, 346, 378, 333, 334, 335, 309, 34: 310, 311, 312, 38: 379, 41: 345, 558, 349, 348, 350, 337, 61: 557, 79: 561, 120: 559, 133: 560},
		{23: 555, 79: 556},
		{103, 103, 344, 376, 380, 381, 382, 383, 384, 103, 103, 328, 13: 347, 16: 313, 330, 397, 342, 343, 331, 366, 25: 332, 377, 346, 378, 333, 334, 335, 309, 34: 310, 311, 312, 38: 379, 541, 41: 345, 351, 349, 348, 350, 337, 61: 338, 77: 329, 79: 547, 81: 375, 341, 354, 336, 355, 356, 357, 339, 359, 360, 362, 352, 367, 385, 353, 364, 100: 542, 543, 358, 361, 106: 363, 108: 546, 365, 544, 545, 115: 548, 124: 549},
		// 245
		{1: 60, 9: 60, 60},
		{270, 270, 270, 270, 270, 270, 270, 270, 270, 11: 270, 13: 270, 16: 270, 270, 270, 270, 270, 270, 270, 32: 270, 34: 270, 270, 270, 38: 270, 270, 79: 270, 81: 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 146: 552},
		{102, 102, 9: 102, 102},
		{101, 101, 9: 101, 101},
		{100, 100, 9: 100, 100},
		// 250
		{99, 99, 9: 99, 99},
		{98, 98, 9: 98, 98},
		{97, 97, 9: 97, 97},
		{96, 96, 9: 96, 96},
		{550, 61, 9: 61, 61},
		// 255
		{103, 103, 344, 376, 380, 381, 382, 383, 384, 103, 103, 328, 13: 347, 16: 313, 330, 397, 342, 343, 331, 366, 25: 332, 377, 346, 378, 333, 334, 335, 309, 34: 310, 311, 312, 38: 379, 541, 41: 345, 351, 349, 348, 350, 337, 61: 338, 77: 329, 79: 547, 81: 375, 341, 354, 336, 355, 356, 357, 339, 359, 360, 362, 352, 367, 385, 353, 364, 100: 542, 543, 358, 361, 106: 363, 108: 546, 365, 544, 545, 115: 551},
		{95, 95, 9: 95, 95},
		{103, 103, 344, 376, 380, 381, 382, 383, 384, 11: 328, 13: 347, 16: 313, 330, 397, 342, 343, 331, 366, 25: 332, 377, 346, 378, 333, 334, 335, 309, 34: 310, 311, 312, 38: 379, 541, 41: 345, 351, 349, 348, 350, 337, 61: 338, 77: 329, 79: 547, 81: 375, 341, 354, 336, 355, 356, 357, 339, 359, 360, 362, 352, 367, 385, 353, 364, 100: 542, 543, 358, 361, 106: 363, 108: 546, 365, 544, 545, 115: 548, 124: 553},
		{550, 554},
		{269, 269, 9: 269, 269},
		// 260
		{64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 13: 64, 16: 64, 64, 64, 64, 64, 64, 64, 32: 64, 34: 64, 64, 64, 38: 64, 64, 79: 64, 81: 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64},
		{62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 13: 62, 16: 62, 62, 62, 62, 62, 62, 62, 32: 62, 34: 62, 62, 62, 38: 62, 62, 79: 62, 81: 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62},
		{3: 427, 426, 428, 429, 433, 444, 12: 293, 14: 293, 23: 293, 293, 33: 293, 37: 293, 47: 425, 430, 431, 432, 434, 435, 436, 437, 438, 439, 440, 441, 442, 443},
		{2: 507, 12: 292, 14: 292, 23: 292, 292, 33: 292, 37: 292},
		{12: 291, 14: 291, 23: 291, 291, 33: 291, 37: 291},
		// 265
		{14: 562, 23: 563, 564, 33: 565},
		{63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 13: 63, 16: 63, 63, 63, 63, 63, 63, 63, 32: 63, 34: 63, 63, 63, 38: 63, 63, 79: 63, 81: 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63},
		{2: 344, 376, 380, 381, 382, 383, 384, 11: 328, 13: 415, 16: 313, 330, 397, 342, 343, 331, 366, 25: 332, 377, 346, 378, 333, 334, 335, 309, 34: 310, 311, 312, 38: 379, 41: 345, 558, 349, 348, 350, 337, 61: 557, 120: 570},
		{67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 13: 67, 16: 67, 67, 67, 67, 67, 67, 67, 32: 67, 34: 67, 67, 67, 38: 67, 67, 79: 67, 81: 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67},
		{2: 344, 376, 380, 381, 382, 383, 384, 11: 328, 13: 415, 16: 313, 330, 397, 342, 343, 331, 366, 25: 332, 377, 346, 378, 333, 334, 335, 309, 34: 310, 311, 312, 38: 379, 41: 345, 351, 349, 348, 350, 337, 61: 568},
		// 270
		{2: 344, 376, 380, 381, 382, 383, 384, 11: 328, 13: 415, 16: 313, 330, 397, 342, 343, 331, 366, 25: 332, 377, 346, 378, 333, 334, 335, 309, 34: 310, 311, 312, 38: 379, 41: 345, 351, 349, 348, 350, 337, 61: 566},
		{3: 427, 426, 428, 429, 433, 444, 23: 567, 47: 425, 430, 431, 432, 434, 435, 436, 437, 438, 439, 440, 441, 442, 443},
		{65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 13: 65, 16: 65, 65, 65, 65, 65, 65, 65, 32: 65, 34: 65, 65, 65, 38: 65, 65, 79: 65, 81: 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65},
		{3: 427, 426, 428, 429, 433, 444, 23: 569, 47: 425, 430, 431, 432, 434, 435, 436, 437, 438, 439, 440, 441, 442, 443},
		{66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 13: 66, 16: 66, 66, 66, 66, 66, 66, 66, 32: 66, 34: 66, 66, 66, 38: 66, 66, 79: 66, 81: 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66},
		// 275
		{12: 290, 14: 290, 23: 290, 290, 33: 290, 37: 290},
		{68, 68, 9: 68, 68},
		{1: 59, 9: 59, 59},
		{2: 344, 376, 380, 381, 382, 383, 384, 11: 328, 13: 575, 15: 57, 313, 330, 397, 342, 343, 331, 366, 25: 332, 377, 346, 378, 333, 334, 335, 309, 34: 310, 311, 312, 38: 379, 41: 345, 351, 349, 348, 350, 337, 61: 574},
		{3: 427, 426, 428, 429, 433, 444, 15: 56, 47: 425, 430, 431, 432, 434, 435, 436, 437, 438, 439, 440, 441, 442, 443},
		// 280
		{2: 194, 194, 194, 194, 194, 194, 194, 11: 194, 15: 194, 33: 576, 39: 194, 47: 194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 194, 75: 194, 97: 416, 121: 417},
		{2: 344, 578, 11: 328, 13: 415, 16: 313, 330, 397, 342, 343, 331, 366, 25: 332, 377, 346, 378, 333, 334, 335, 309, 34: 310, 311, 312, 41: 345, 351, 349, 348, 577},
		{2: 585, 11: 588, 39: 586, 75: 587, 113: 590, 134: 589},
		{3: 578, 11: 328, 17: 330, 397, 342, 343, 331, 366, 25: 580, 377, 581, 378, 582, 583, 584, 42: 579},
		{2: 27},
		// 285
		{2: 26},
		{2: 24},
		{2: 22},
		{2: 21},
		{2: 20},
		// 290
		{2: 344, 376, 380, 381, 382, 383, 384, 11: 328, 623, 415, 16: 313, 330, 397, 342, 343, 331, 366, 25: 332, 377, 346, 378, 333, 334, 335, 309, 34: 310, 311, 312, 38: 379, 41: 345, 558, 349, 348, 350, 337, 61: 557, 120: 559, 133: 622},
		{1: 611, 344, 376, 380, 381, 382, 383, 384, 11: 328, 13: 415, 16: 313, 330, 397, 342, 343, 331, 366, 25: 332, 377, 346, 378, 333, 334, 335, 309, 34: 310, 311, 312, 38: 379, 586, 41: 345, 351, 349, 348, 350, 337, 61: 608, 113: 607, 135: 609, 156: 610},
		{2: 601, 13: 602},
		{2: 344, 376, 380, 381, 382, 383, 384, 11: 328, 13: 415, 16: 313, 330, 397, 342, 343, 331, 366, 210, 25: 332, 377, 346, 378, 333, 334, 335, 309, 34: 310, 311, 312, 38: 379, 41: 345, 351, 349, 348, 350, 337, 61: 591, 127: 592},
		{130, 130, 130, 130, 130, 130, 130, 130, 130, 130, 130, 130, 130, 14: 130, 130, 23: 130, 130, 33: 130, 37: 130, 39: 130, 130, 47: 130, 130, 130, 130, 130, 130, 130, 130, 130, 130, 130, 130, 130, 130, 62: 130, 130, 130, 130, 130, 130, 130, 130, 130, 130, 130, 130, 130, 130},
		// 295
		{129, 129, 129, 129, 129, 129, 129, 129, 129, 129, 129, 129, 129, 14: 129, 129, 23: 129, 129, 33: 129, 37: 129, 39: 129, 129, 47: 129, 129, 129, 129, 129, 129, 129, 129, 129, 129, 129, 129, 129, 129, 62: 129, 129, 129, 129, 129, 129, 129, 129, 129, 129, 129, 129, 129, 129},
		{3: 427, 426, 428, 429, 433, 444, 23: 209, 40: 600, 47: 425, 430, 431, 432, 434, 435, 436, 437, 438, 439, 440, 441, 442, 443},
		{23: 593},
		{2: 344, 376, 380, 381, 382, 383, 384, 11: 328, 13: 415, 16: 313, 330, 397, 342, 343, 331, 366, 210, 25: 332, 377, 346, 378, 333, 334, 335, 309, 34: 310, 311, 312, 38: 379, 40: 210, 345, 351, 349, 348, 350, 337, 61: 594, 127: 595},
		{3: 427, 426, 428, 429, 433, 444, 23: 209, 40: 209, 47: 425, 430, 431, 432, 434, 435, 436, 437, 438, 439, 440, 441, 442, 443},
		// 300
		{23: 596, 40: 597},
		{2: 344, 376, 380, 381, 382, 383, 384, 11: 328, 13: 415, 16: 313, 330, 397, 342, 343, 331, 366, 25: 332, 377, 346, 378, 333, 334, 335, 309, 34: 310, 311, 312, 38: 379, 40: 210, 345, 351, 349, 348, 350, 337, 61: 594, 127: 598},
		{131, 131, 131, 131, 131, 131, 131, 131, 131, 131, 131, 131, 131, 14: 131, 131, 23: 131, 131, 33: 131, 37: 131, 39: 131, 131, 47: 131, 131, 131, 131, 131, 131, 131, 131, 131, 131, 131, 131, 131, 131, 62: 131, 131, 131, 131, 131, 131, 131, 131, 131, 131, 131, 131, 131, 131},
		{40: 599},
		{132, 132, 132, 132, 132, 132, 132, 132, 132, 132, 132, 132, 132, 14: 132, 132, 23: 132, 132, 33: 132, 37: 132, 39: 132, 132, 47: 132, 132, 132, 132, 132, 132, 132, 132, 132, 132, 132, 132, 132, 132, 62: 132, 132, 132, 132, 132, 132, 132, 132, 132, 132, 132, 132, 132, 132},
		// 305
		{133, 133, 133, 133, 133, 133, 133, 133, 133, 133, 133, 133, 133, 14: 133, 133, 23: 133, 133, 33: 133, 37: 133, 39: 133, 133, 47: 133, 133, 133, 133, 133, 133, 133, 133, 133, 133, 133, 133, 133, 133, 62: 133, 133, 133, 133, 133, 133, 133, 133, 133, 133, 133, 133, 133, 133},
		{2: 400, 401, 11: 328, 13: 399, 17: 330, 397, 342, 343, 331, 366, 25: 402, 403, 404, 405, 406, 408, 409, 78: 407, 80: 604, 603},
		{134, 134, 134, 134, 134, 134, 134, 134, 134, 134, 134, 134, 134, 14: 134, 134, 23: 134, 134, 33: 134, 37: 134, 39: 134, 134, 47: 134, 134, 134, 134, 134, 134, 134, 134, 134, 134, 134, 134, 134, 134, 62: 134, 134, 134, 134, 134, 134, 134, 134, 134, 134, 134, 134, 134, 134},
		{12: 606},
		{12: 605},
		// 310
		{135, 135, 135, 135, 135, 135, 135, 135, 135, 135, 135, 135, 135, 14: 135, 135, 23: 135, 135, 33: 135, 37: 135, 39: 135, 135, 47: 135, 135, 135, 135, 135, 135, 135, 135, 135, 135, 135, 135, 135, 135, 62: 135, 135, 135, 135, 135, 135, 135, 135, 135, 135, 135, 135, 135, 135},
		{2: 136, 11: 136, 15: 55, 39: 136, 75: 136},
		{1: 258, 14: 258, 23: 619},
		{1: 255, 3: 427, 426, 428, 429, 433, 444, 14: 255, 23: 616, 47: 425, 430, 431, 432, 434, 435, 436, 437, 438, 439, 440, 441, 442, 443},
		{1: 252, 14: 252},
		// 315
		{1: 260, 14: 612, 112: 613},
		{246, 246, 246, 246, 246, 246, 246, 246, 246, 246, 246, 246, 246, 14: 246, 246, 23: 246, 246, 33: 246, 37: 246, 39: 246, 246, 47: 246, 246, 246, 246, 246, 246, 246, 246, 246, 246, 246, 246, 246, 246, 62: 246, 246, 246, 246, 246, 246, 246, 246, 246, 246, 246, 246, 246, 246},
		{1: 259, 344, 376, 380, 381, 382, 383, 384, 11: 328, 13: 415, 16: 313, 330, 397, 342, 343, 331, 366, 25: 332, 377, 346, 378, 333, 334, 335, 309, 34: 310, 311, 312, 38: 379, 586, 41: 345, 351, 349, 348, 350, 337, 61: 608, 113: 607, 135: 615},
		{1: 614},
		{245, 245, 245, 245, 245, 245, 245, 245, 245, 245, 245, 245, 245, 14: 245, 245, 23: 245, 245, 33: 245, 37: 245, 39: 245, 245, 47: 245, 245, 245, 245, 245, 245, 245, 245, 245, 245, 245, 245, 245, 245, 62: 245, 245, 245, 245, 245, 245, 245, 245, 245, 245, 245, 245, 245, 245},
		// 320
		{1: 251, 14: 251},
		{2: 344, 376, 380, 381, 382, 383, 384, 11: 328, 13: 415, 16: 313, 330, 397, 342, 343, 331, 366, 25: 332, 377, 346, 378, 333, 334, 335, 309, 34: 310, 311, 312, 38: 379, 586, 41: 345, 351, 349, 348, 350, 337, 61: 618, 113: 617},
		{1: 254, 14: 254},
		{1: 253, 3: 427, 426, 428, 429, 433, 444, 14: 253, 47: 425, 430, 431, 432, 434, 435, 436, 437, 438, 439, 440, 441, 442, 443},
		{2: 344, 376, 380, 381, 382, 383, 384, 11: 328, 13: 415, 16: 313, 330, 397, 342, 343, 331, 366, 25: 332, 377, 346, 378, 333, 334, 335, 309, 34: 310, 311, 312, 38: 379, 586, 41: 345, 351, 349, 348, 350, 337, 61: 621, 113: 620},
		// 325
		{1: 257, 14: 257},
		{1: 256, 3: 427, 426, 428, 429, 433, 444, 14: 256, 47: 425, 430, 431, 432, 434, 435, 436, 437, 438, 439, 440, 441, 442, 443},
		{12: 260, 14: 624, 37: 626, 112: 625},
		{266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 14: 266, 266, 23: 266, 266, 33: 266, 37: 266, 39: 266, 266, 47: 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 62: 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266},
		{2: 344, 376, 380, 381, 382, 383, 384, 11: 328, 259, 415, 16: 313, 330, 397, 342, 343, 331, 366, 25: 332, 377, 346, 378, 333, 334, 335, 309, 34: 310, 311, 312, 38: 379, 41: 345, 558, 349, 348, 350, 337, 61: 557, 120: 570},
		// 330
		{12: 629},
		{12: 260, 14: 509, 112: 627},
		{12: 628},
		{264, 264, 264, 264, 264, 264, 264, 264, 264, 264, 264, 264, 264, 14: 264, 264, 23: 264, 264, 33: 264, 37: 264, 39: 264, 264, 47: 264, 264, 264, 264, 264, 264, 264, 264, 264, 264, 264, 264, 264, 264, 62: 264, 264, 264, 264, 264, 264, 264, 264, 264, 264, 264, 264, 264, 264},
		{265, 265, 265, 265, 265, 265, 265, 265, 265, 265, 265, 265, 265, 14: 265, 265, 23: 265, 265, 33: 265, 37: 265, 39: 265, 265, 47: 265, 265, 265, 265, 265, 265, 265, 265, 265, 265, 265, 265, 265, 265, 62: 265, 265, 265, 265, 265, 265, 265, 265, 265, 265, 265, 265, 265, 265},
		// 335
		{164, 164, 164, 164, 164, 164, 164, 164, 164, 11: 164, 13: 164, 15: 164, 164, 164, 164, 164, 164, 164, 164, 32: 164, 34: 164, 164, 164, 38: 164, 164, 79: 164, 81: 164, 164, 164, 164, 164, 164, 164, 164, 164, 164, 164, 164, 164, 164},
		{163, 163, 163, 163, 163, 163, 163, 163, 163, 11: 163, 13: 163, 15: 163, 163, 163, 163, 163, 163, 163, 163, 32: 163, 34: 163, 163, 163, 38: 163, 163, 79: 163, 81: 163, 163, 163, 163, 163, 163, 163, 163, 163, 163, 163, 163, 163, 163},
		{1: 633, 72, 72, 13: 72, 154: 634},
		{73, 73, 73, 9: 73, 73, 12: 73, 14: 73, 73, 73, 23: 73, 73, 33: 73, 37: 73, 39: 73, 73, 76: 73},
		{2: 639, 637, 13: 635, 78: 638, 98: 636, 142: 640, 176: 641},
		// 340
		{126, 126, 188, 188, 11: 188, 13: 188, 188, 16: 126, 188, 188, 188, 188, 188, 188, 75: 470},
		{2: 400, 401, 11: 328, 13: 399, 398, 17: 330, 397, 342, 343, 331, 366, 25: 402, 403, 404, 405, 406, 408, 409, 78: 407, 80: 661},
		{2: 656, 13: 399, 78: 655},
		{52, 52, 16: 651, 116: 654},
		{3: 647, 13: 399, 78: 646},
		// 345
		{75, 75},
		{642, 114, 114: 643},
		{1: 113, 639, 637, 13: 635, 78: 638, 98: 636, 142: 645},
		{1: 644},
		{71, 71, 71, 9: 71, 71, 12: 71, 14: 71, 71, 71, 23: 71, 71, 33: 71, 37: 71, 39: 71, 71, 76: 71},
		// 350
		{74, 74},
		{12: 652},
		{13: 399, 78: 648},
		{12: 649},
		{52, 52, 16: 651, 116: 650},
		// 355
		{77, 77},
		{51, 51},
		{52, 52, 16: 651, 116: 653},
		{78, 78},
		{79, 79},
		// 360
		{52, 52, 16: 651, 116: 660},
		{13: 399, 78: 657},
		{12: 658},
		{52, 52, 16: 651, 116: 659},
		{76, 76},
		// 365
		{81, 81},
		{52, 52, 16: 651, 116: 662},
		{80, 80},
		{205, 205, 9: 205, 205, 14: 423},
		{85, 85, 9: 85, 85},
		// 370
		{88, 88, 9: 88, 88},
		{89, 89, 3: 427, 426, 428, 429, 433, 444, 89, 89, 47: 425, 430, 431, 432, 434, 435, 436, 437, 438, 439, 440, 441, 442, 443},
		{92, 92, 3: 427, 426, 428, 429, 433, 444, 92, 92, 47: 425, 430, 431, 432, 434, 435, 436, 437, 438, 439, 440, 441, 442, 443},
		{93, 93, 9: 93, 93},
		{94, 94, 9: 94, 94},
		// 375
		{15: 532, 143: 671},
		{115, 115, 9: 115, 115},
		{2: 673, 13: 602},
		{2: 400, 401, 11: 328, 13: 399, 17: 330, 397, 342, 343, 331, 366, 25: 402, 403, 404, 405, 406, 408, 409, 78: 407, 80: 604, 674},
		{12: 675},
		// 380
		{136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 14: 136, 136, 23: 136, 136, 33: 136, 37: 136, 39: 136, 136, 47: 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 62: 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136},
		{1: 682, 344, 376, 380, 381, 382, 383, 384, 11: 328, 13: 415, 15: 630, 313, 330, 397, 342, 343, 331, 366, 25: 332, 377, 346, 378, 333, 334, 335, 309, 34: 310, 311, 312, 38: 379, 631, 41: 345, 351, 349, 348, 350, 337, 61: 678, 104: 676, 123: 679, 139: 680, 170: 681},
		{137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 14: 137, 137, 23: 137, 137, 33: 137, 37: 137, 39: 137, 137, 47: 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 62: 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137},
		{1: 162, 3: 427, 426, 428, 429, 433, 444, 14: 162, 23: 687, 47: 425, 430, 431, 432, 434, 435, 436, 437, 438, 439, 440, 441, 442, 443},
		{1: 159, 14: 159},
		// 385
		{1: 158, 14: 158},
		{1: 260, 14: 683, 112: 684},
		{156, 156, 156, 156, 156, 156, 156, 156, 156, 156, 156, 156, 156, 14: 156, 156, 23: 156, 156, 33: 156, 37: 156, 39: 156, 156, 47: 156, 156, 156, 156, 156, 156, 156, 156, 156, 156, 156, 156, 156, 156, 62: 156, 156, 156, 156, 156, 156, 156, 156, 156, 156, 156, 156, 156, 156},
		{1: 259, 344, 376, 380, 381, 382, 383, 384, 11: 328, 13: 415, 15: 630, 313, 330, 397, 342, 343, 331, 366, 25: 332, 377, 346, 378, 333, 334, 335, 309, 34: 310, 311, 312, 38: 379, 631, 41: 345, 351, 349, 348, 350, 337, 61: 678, 104: 676, 123: 679, 139: 686},
		{1: 685},
		// 390
		{155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 14: 155, 155, 23: 155, 155, 33: 155, 37: 155, 39: 155, 155, 47: 155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 62: 155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 155},
		{1: 157, 14: 157},
		{2: 344, 376, 380, 381, 382, 383, 384, 11: 328, 13: 415, 15: 630, 313, 330, 397, 342, 343, 331, 366, 25: 332, 377, 346, 378, 333, 334, 335, 309, 34: 310, 311, 312, 38: 379, 631, 41: 345, 351, 349, 348, 350, 337, 61: 688, 104: 676, 123: 689},
		{1: 161, 3: 427, 426, 428, 429, 433, 444, 14: 161, 47: 425, 430, 431, 432, 434, 435, 436, 437, 438, 439, 440, 441, 442, 443},
		{1: 160, 14: 160},
		// 395
		{103, 103, 344, 376, 380, 381, 382, 383, 384, 103, 103, 328, 13: 347, 16: 313, 330, 397, 342, 343, 331, 366, 25: 332, 377, 346, 378, 333, 334, 335, 309, 34: 310, 311, 312, 38: 379, 541, 41: 345, 351, 349, 348, 350, 337, 61: 338, 77: 329, 79: 547, 81: 375, 341, 354, 336, 355, 356, 357, 339, 359, 360, 362, 352, 367, 385, 353, 364, 100: 542, 543, 358, 361, 106: 363, 108: 546, 365, 544, 545, 115: 691},
		{87, 87, 9: 87, 87},
		{15: 630, 39: 631, 104: 693},
		{103, 103, 344, 376, 380, 381, 382, 383, 384, 11: 328, 13: 347, 16: 313, 330, 397, 342, 343, 331, 366, 25: 332, 377, 346, 378, 333, 334, 335, 309, 34: 310, 311, 312, 38: 379, 541, 41: 345, 351, 349, 348, 350, 337, 61: 338, 77: 329, 79: 547, 81: 375, 341, 354, 336, 355, 356, 357, 339, 359, 360, 362, 352, 367, 385, 353, 364, 100: 542, 543, 358, 361, 106: 363, 108: 546, 365, 544, 545, 115: 548, 124: 694},
		{550, 695},
		// 400
		{149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 14: 149, 149, 23: 149, 149, 33: 149, 37: 149, 39: 149, 149, 47: 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 62: 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149},
		{3: 427, 426, 428, 429, 433, 444, 12: 699, 47: 425, 430, 431, 432, 434, 435, 436, 437, 438, 439, 440, 441, 442, 443},
		{2: 507, 12: 698},
		{152, 152, 152, 152, 152, 152, 152, 152, 152, 152, 152, 152, 152, 14: 152, 152, 23: 152, 152, 33: 152, 37: 152, 39: 152, 152, 47: 152, 152, 152, 152, 152, 152, 152, 152, 152, 152, 152, 152, 152, 152, 62: 152, 152, 152, 152, 152, 152, 152, 152, 152, 152, 152, 152, 152, 152},
		{153, 153, 153, 153, 153, 153, 153, 153, 153, 153, 153, 153, 153, 14: 153, 153, 23: 153, 153, 33: 153, 37: 153, 39: 153, 153, 47: 153, 153, 153, 153, 153, 153, 153, 153, 153, 153, 153, 153, 153, 153, 62: 153, 153, 153, 153, 153, 153, 153, 153, 153, 153, 153, 153, 153, 153},
		// 405
		{2: 400, 401, 11: 328, 13: 399, 17: 330, 397, 342, 343, 331, 366, 25: 402, 403, 404, 405, 406, 408, 409, 78: 407, 80: 701},
		{40: 702},
		{2: 400, 401, 11: 328, 13: 399, 17: 330, 397, 342, 343, 331, 366, 25: 402, 403, 404, 405, 406, 408, 409, 78: 407, 80: 703},
		{154, 154, 154, 9: 154, 154, 12: 154, 14: 154, 154, 154, 23: 154, 154, 33: 154, 37: 154, 39: 154, 154, 76: 154},
		{1: 705, 13: 171, 150: 706},
		// 410
		{172, 172, 172, 9: 172, 172, 12: 172, 14: 172, 172, 172, 23: 172, 172, 33: 172, 37: 172, 39: 172, 172, 76: 172},
		{13: 708, 78: 709, 138: 710, 169: 707},
		{714, 114, 114: 713},
		{126, 126, 169, 75: 470, 151: 711},
		{167, 167},
		// 415
		{166, 166},
		{2: 475, 105: 476, 119: 712},
		{168, 168},
		{1: 716},
		{1: 113, 13: 708, 78: 709, 138: 715},
		// 420
		{165, 165},
		{170, 170, 170, 9: 170, 170, 12: 170, 14: 170, 170, 170, 23: 170, 170, 33: 170, 37: 170, 39: 170, 170, 76: 170},
		{732, 15: 186},
		{15: 719, 125: 720},
		{268, 268, 268, 268, 268, 268, 268, 268, 268, 11: 268, 13: 268, 16: 268, 268, 268, 268, 268, 268, 268, 32: 268, 34: 268, 268, 268, 38: 268, 268, 79: 268, 81: 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 148: 729},
		// 425
		{235, 235, 9: 235, 235, 117: 235, 159: 721},
		{233, 233, 9: 233, 233, 117: 722, 158: 723, 160: 724},
		{39: 541, 82: 725, 100: 726},
		{234, 234, 9: 234, 234, 117: 234},
		{184, 184, 9: 184, 184},
		// 430
		{106, 2: 344, 376, 380, 381, 382, 383, 384, 11: 328, 13: 415, 15: 106, 313, 330, 397, 342, 343, 331, 366, 25: 332, 377, 346, 378, 333, 334, 335, 309, 34: 310, 311, 312, 38: 379, 41: 345, 351, 349, 348, 350, 337, 61: 338, 77: 329, 95: 353, 528, 107: 717, 137: 727},
		{232, 232, 9: 232, 232},
		{15: 719, 125: 728},
		{236, 236, 9: 236, 236, 117: 236},
		{103, 103, 344, 376, 380, 381, 382, 383, 384, 11: 328, 13: 347, 16: 313, 330, 397, 342, 343, 331, 366, 25: 332, 377, 346, 378, 333, 334, 335, 309, 34: 310, 311, 312, 38: 379, 541, 41: 345, 351, 349, 348, 350, 337, 61: 338, 77: 329, 79: 547, 81: 375, 341, 354, 336, 355, 356, 357, 339, 359, 360, 362, 352, 367, 385, 353, 364, 100: 542, 543, 358, 361, 106: 363, 108: 546, 365, 544, 545, 115: 548, 124: 730},
		// 435
		{550, 731},
		{267, 267, 9: 267, 267, 117: 267},
		{2: 344, 376, 380, 381, 382, 383, 384, 11: 328, 13: 415, 15: 106, 313, 330, 397, 342, 343, 331, 366, 25: 332, 377, 346, 378, 333, 334, 335, 309, 34: 310, 311, 312, 38: 379, 41: 345, 351, 349, 348, 350, 337, 61: 338, 77: 329, 95: 353, 528, 107: 733},
		{15: 185},
		{13: 736},
		// 440
		{2: 477, 401, 11: 328, 13: 120, 15: 119, 17: 330, 397, 342, 343, 331, 366, 25: 402, 403, 404, 405, 406, 408, 409, 39: 119, 78: 407, 80: 479, 105: 478, 141: 480},
		{2: 192, 97: 515, 136: 737},
		{2: 475, 105: 476, 119: 738},
		{198, 39: 198, 149: 739},
		{200, 39: 541, 100: 740, 164: 741},
		// 445
		{199},
		{197},
		{14: 423, 24: 753, 33: 765, 62: 754, 756, 755, 757, 758, 759, 760, 761, 762, 763, 764},
		{15: 204},
		{749, 15: 202},
		// 450
		{15: 719, 125: 748},
		{2: 344, 376, 380, 381, 382, 383, 384, 11: 328, 13: 415, 16: 313, 330, 397, 342, 343, 331, 366, 25: 332, 377, 346, 378, 333, 334, 335, 309, 34: 310, 311, 312, 38: 379, 41: 345, 351, 349, 348, 350, 337, 61: 747},
		{3: 427, 426, 428, 429, 433, 444, 15: 122, 47: 425, 430, 431, 432, 434, 435, 436, 437, 438, 439, 440, 441, 442, 443},
		{201, 201, 9: 201, 201},
		{106, 2: 344, 376, 380, 381, 382, 383, 384, 11: 328, 13: 415, 16: 313, 330, 397, 342, 343, 331, 366, 25: 332, 377, 346, 378, 333, 334, 335, 309, 34: 310, 311, 312, 38: 379, 41: 345, 351, 349, 348, 350, 337, 61: 338, 77: 329, 95: 353, 528, 107: 750},
		// 455
		{751},
		{2: 344, 376, 380, 381, 382, 383, 384, 11: 328, 13: 415, 15: 106, 313, 330, 397, 342, 343, 331, 366, 25: 332, 377, 346, 378, 333, 334, 335, 309, 34: 310, 311, 312, 38: 379, 41: 345, 351, 349, 348, 350, 337, 61: 338, 77: 329, 95: 353, 528, 107: 752},
		{15: 203},
		{2: 344, 376, 380, 381, 382, 383, 384, 11: 328, 13: 415, 16: 313, 330, 397, 342, 343, 331, 366, 25: 332, 377, 346, 378, 333, 334, 335, 309, 34: 310, 311, 312, 38: 379, 41: 345, 351, 349, 348, 350, 337, 61: 413, 77: 780, 130: 781},
		{2: 344, 376, 380, 381, 382, 383, 384, 11: 328, 13: 415, 16: 313, 330, 397, 342, 343, 331, 366, 25: 332, 377, 346, 378, 333, 334, 335, 309, 34: 310, 311, 312, 38: 379, 41: 345, 351, 349, 348, 350, 337, 61: 413, 77: 779},
		// 460
		{2: 344, 376, 380, 381, 382, 383, 384, 11: 328, 13: 415, 16: 313, 330, 397, 342, 343, 331, 366, 25: 332, 377, 346, 378, 333, 334, 335, 309, 34: 310, 311, 312, 38: 379, 41: 345, 351, 349, 348, 350, 337, 61: 413, 77: 778},
		{2: 344, 376, 380, 381, 382, 383, 384, 11: 328, 13: 415, 16: 313, 330, 397, 342, 343, 331, 366, 25: 332, 377, 346, 378, 333, 334, 335, 309, 34: 310, 311, 312, 38: 379, 41: 345, 351, 349, 348, 350, 337, 61: 413, 77: 777},
		{2: 344, 376, 380, 381, 382, 383, 384, 11: 328, 13: 415, 16: 313, 330, 397, 342, 343, 331, 366, 25: 332, 377, 346, 378, 333, 334, 335, 309, 34: 310, 311, 312, 38: 379, 41: 345, 351, 349, 348, 350, 337, 61: 413, 77: 776},
		{2: 344, 376, 380, 381, 382, 383, 384, 11: 328, 13: 415, 16: 313, 330, 397, 342, 343, 331, 366, 25: 332, 377, 346, 378, 333, 334, 335, 309, 34: 310, 311, 312, 38: 379, 41: 345, 351, 349, 348, 350, 337, 61: 413, 77: 775},
		{2: 344, 376, 380, 381, 382, 383, 384, 11: 328, 13: 415, 16: 313, 330, 397, 342, 343, 331, 366, 25: 332, 377, 346, 378, 333, 334, 335, 309, 34: 310, 311, 312, 38: 379, 41: 345, 351, 349, 348, 350, 337, 61: 413, 77: 774},
		// 465
		{2: 344, 376, 380, 381, 382, 383, 384, 11: 328, 13: 415, 16: 313, 330, 397, 342, 343, 331, 366, 25: 332, 377, 346, 378, 333, 334, 335, 309, 34: 310, 311, 312, 38: 379, 41: 345, 351, 349, 348, 350, 337, 61: 413, 77: 773},
		{2: 344, 376, 380, 381, 382, 383, 384, 11: 328, 13: 415, 16: 313, 330, 397, 342, 343, 331, 366, 25: 332, 377, 346, 378, 333, 334, 335, 309, 34: 310, 311, 312, 38: 379, 41: 345, 351, 349, 348, 350, 337, 61: 413, 77: 772},
		{2: 344, 376, 380, 381, 382, 383, 384, 11: 328, 13: 415, 16: 313, 330, 397, 342, 343, 331, 366, 25: 332, 377, 346, 378, 333, 334, 335, 309, 34: 310, 311, 312, 38: 379, 41: 345, 351, 349, 348, 350, 337, 61: 413, 77: 771},
		{2: 344, 376, 380, 381, 382, 383, 384, 11: 328, 13: 415, 16: 313, 330, 397, 342, 343, 331, 366, 25: 332, 377, 346, 378, 333, 334, 335, 309, 34: 310, 311, 312, 38: 379, 41: 345, 351, 349, 348, 350, 337, 61: 413, 77: 770},
		{2: 344, 376, 380, 381, 382, 383, 384, 11: 328, 13: 415, 16: 313, 330, 397, 342, 343, 331, 366, 25: 332, 377, 346, 378, 333, 334, 335, 309, 34: 310, 311, 312, 38: 379, 41: 345, 351, 349, 348, 350, 337, 61: 413, 77: 769},
		// 470
		{2: 344, 376, 380, 381, 382, 383, 384, 11: 328, 13: 415, 16: 313, 330, 397, 342, 343, 331, 366, 25: 332, 377, 346, 378, 333, 334, 335, 309, 34: 310, 311, 312, 38: 379, 41: 345, 351, 349, 348, 350, 337, 61: 413, 77: 766, 130: 767},
		{107, 107, 9: 107, 107, 14: 423, 107},
		{2: 344, 376, 380, 381, 382, 383, 384, 11: 328, 13: 415, 16: 313, 330, 397, 342, 343, 331, 366, 25: 332, 377, 346, 378, 333, 334, 335, 309, 34: 310, 311, 312, 38: 379, 41: 345, 351, 349, 348, 350, 337, 61: 768},
		{3: 427, 426, 428, 429, 433, 444, 15: 123, 47: 425, 430, 431, 432, 434, 435, 436, 437, 438, 439, 440, 441, 442, 443},
		{276, 276, 9: 276, 276, 14: 423, 276},
		// 475
		{277, 277, 9: 277, 277, 14: 423, 277},
		{278, 278, 9: 278, 278, 14: 423, 278},
		{279, 279, 9: 279, 279, 14: 423, 279},
		{280, 280, 9: 280, 280, 14: 423, 280},
		{281, 281, 9: 281, 281, 14: 423, 281},
		// 480
		{282, 282, 9: 282, 282, 14: 423, 282},
		{283, 283, 9: 283, 283, 14: 423, 283},
		{284, 284, 9: 284, 284, 14: 423, 284},
		{285, 285, 9: 285, 285, 14: 423, 285},
		{286, 286, 9: 286, 286, 14: 423, 286},
		// 485
		{287, 287, 9: 287, 287, 14: 423, 287},
		{2: 344, 376, 380, 381, 382, 383, 384, 11: 328, 13: 415, 16: 313, 330, 397, 342, 343, 331, 366, 25: 332, 377, 346, 378, 333, 334, 335, 309, 34: 310, 311, 312, 38: 379, 41: 345, 351, 349, 348, 350, 337, 61: 782},
		{3: 427, 426, 428, 429, 433, 444, 15: 124, 47: 425, 430, 431, 432, 434, 435, 436, 437, 438, 439, 440, 441, 442, 443},
		{109, 109, 9: 109, 109, 15: 109},
		{108, 108, 9: 108, 108, 15: 108},
		// 490
		{12: 793, 386, 98: 787, 126: 795, 157: 794},
		{242, 242, 9: 242, 242},
		{241, 241, 400, 401, 9: 241, 241, 328, 241, 399, 398, 17: 330, 397, 342, 343, 331, 366, 24: 788, 402, 403, 404, 405, 406, 408, 409, 78: 407, 80: 789},
		{2: 344, 376, 380, 381, 382, 383, 384, 11: 328, 13: 415, 16: 313, 330, 397, 342, 343, 331, 366, 25: 332, 377, 346, 378, 333, 334, 335, 309, 34: 310, 311, 312, 38: 379, 41: 345, 351, 349, 348, 350, 337, 61: 413, 77: 792},
		{24: 790},
		// 495
		{2: 344, 376, 380, 381, 382, 383, 384, 11: 328, 13: 415, 16: 313, 330, 397, 342, 343, 331, 366, 25: 332, 377, 346, 378, 333, 334, 335, 309, 34: 310, 311, 312, 38: 379, 41: 345, 351, 349, 348, 350, 337, 61: 413, 77: 791},
		{239, 239, 9: 239, 239, 12: 239, 14: 423},
		{240, 240, 9: 240, 240, 12: 240, 14: 423},
		{244, 244, 9: 244, 244},
		{797, 12: 114, 114: 796},
		// 500
		{238, 12: 238},
		{12: 799},
		{12: 113, 386, 98: 787, 126: 798},
		{237, 12: 237},
		{243, 243, 9: 243, 243},
		// 505
		{2: 400, 401, 11: 328, 13: 399, 17: 330, 397, 342, 343, 331, 366, 25: 402, 403, 404, 405, 406, 408, 409, 78: 407, 80: 801},
		{261, 261, 261, 9: 261, 261, 12: 261, 14: 261, 261, 261, 23: 261, 261, 33: 261, 37: 261, 39: 261, 261, 76: 261},
		{263, 263, 263, 9: 263, 263, 12: 263, 14: 263, 263, 263, 23: 263, 263, 33: 263, 37: 263, 39: 263, 263, 76: 263},
		{2: 400, 401, 11: 328, 13: 399, 17: 330, 397, 342, 343, 331, 366, 25: 402, 403, 404, 405, 406, 408, 409, 78: 407, 80: 804},
		{262, 262, 262, 9: 262, 262, 12: 262, 14: 262, 262, 262, 23: 262, 262, 33: 262, 37: 262, 39: 262, 262, 76: 262},
		// 510
		{2: 344, 376, 380, 381, 382, 383, 384, 11: 328, 13: 415, 16: 313, 330, 397, 342, 343, 331, 366, 25: 332, 377, 346, 378, 333, 334, 335, 309, 34: 310, 311, 312, 38: 379, 41: 345, 351, 349, 348, 350, 337, 61: 413, 77: 780},
		{2: 344, 376, 380, 381, 382, 383, 384, 11: 328, 13: 415, 16: 313, 330, 397, 342, 343, 331, 366, 25: 332, 377, 346, 378, 333, 334, 335, 309, 34: 310, 311, 312, 38: 379, 41: 345, 351, 349, 348, 350, 337, 61: 413, 77: 766},
		{40: 813},
		{3: 427, 426, 428, 429, 433, 444, 40: 811, 47: 425, 430, 431, 432, 434, 435, 436, 437, 438, 439, 440, 441, 442, 443},
		{2: 400, 401, 11: 328, 13: 399, 17: 330, 397, 342, 343, 331, 366, 25: 402, 403, 404, 405, 406, 408, 409, 78: 407, 80: 810},
		// 515
		{104, 104, 104, 9: 104, 104, 12: 104, 14: 104, 104, 104, 23: 104, 104, 33: 104, 37: 104, 39: 104, 104, 76: 104},
		{2: 400, 401, 11: 328, 13: 399, 17: 330, 397, 342, 343, 331, 366, 25: 402, 403, 404, 405, 406, 408, 409, 78: 407, 80: 812},
		{288, 288, 288, 9: 288, 288, 12: 288, 14: 288, 288, 288, 23: 288, 288, 33: 288, 37: 288, 39: 288, 288, 76: 288},
		{2: 400, 401, 11: 328, 13: 399, 17: 330, 397, 342, 343, 331, 366, 25: 402, 403, 404, 405, 406, 408, 409, 78: 407, 80: 814},
		{289, 289, 289, 9: 289, 289, 12: 289, 14: 289, 289, 289, 23: 289, 289, 33: 289, 37: 289, 39: 289, 289, 76: 289},
	}
)

var yyDebug = 0

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyLexerEx interface {
	yyLexer
	Reduced(rule, state int, lval *yySymType) bool
}

func yySymName(c int) (s string) {
	x, ok := yyXLAT[c]
	if ok {
		return yySymNames[x]
	}

	if c < 0x7f {
		return __yyfmt__.Sprintf("'%c'", c)
	}

	return __yyfmt__.Sprintf("%d", c)
}

func yylex1(yylex yyLexer, lval *yySymType) (n int) {
	n = yylex.Lex(lval)
	if n <= 0 {
		n = yyEofCode
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("\nlex %s(%#x %d), PrettyString(lval.Token): %v\n", yySymName(n), n, n, PrettyString(lval.Token))
	}
	return n
}

func yyParse(yylex yyLexer) int {
	const yyError = 79

	yyEx, _ := yylex.(yyLexerEx)
	var yyn int
	var yylval yySymType
	var yyVAL yySymType
	yyS := make([]yySymType, 200)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yyerrok := func() {
		if yyDebug >= 2 {
			__yyfmt__.Printf("yyerrok()\n")
		}
		Errflag = 0
	}
	_ = yyerrok
	yystate := 0
	yychar := -1
	var yyxchar int
	var yyshift int
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	if yychar < 0 {
		yychar = yylex1(yylex, &yylval)
		var ok bool
		if yyxchar, ok = yyXLAT[yychar]; !ok {
			yyxchar = len(yySymNames) // > tab width
		}
	}
	if yyDebug >= 4 {
		var a []int
		for _, v := range yyS[:yyp+1] {
			a = append(a, v.yys)
		}
		__yyfmt__.Printf("state stack %v\n", a)
	}
	row := yyParseTab[yystate]
	yyn = 0
	if yyxchar < len(row) {
		if yyn = int(row[yyxchar]); yyn != 0 {
			yyn += yyTabOfs
		}
	}
	switch {
	case yyn > 0: // shift
		yychar = -1
		yyVAL = yylval
		yystate = yyn
		yyshift = yyn
		if yyDebug >= 2 {
			__yyfmt__.Printf("shift, and goto state %d\n", yystate)
		}
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	case yyn < 0: // reduce
	case yystate == 1: // accept
		if yyDebug >= 2 {
			__yyfmt__.Println("accept")
		}
		goto ret0
	}

	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			if yyDebug >= 1 {
				__yyfmt__.Printf("no action for %s in state %d\n", yySymName(yychar), yystate)
			}
			msg, ok := yyXErrors[yyXError{yystate, yyxchar}]
			if !ok {
				msg, ok = yyXErrors[yyXError{yystate, -1}]
			}
			if !ok && yyshift != 0 {
				msg, ok = yyXErrors[yyXError{yyshift, yyxchar}]
			}
			if !ok {
				msg, ok = yyXErrors[yyXError{yyshift, -1}]
			}
			if yychar > 0 {
				ls := yyTokenLiteralStrings[yychar]
				if ls == "" {
					ls = yySymName(yychar)
				}
				if ls != "" {
					switch {
					case msg == "":
						msg = __yyfmt__.Sprintf("unexpected %s", ls)
					default:
						msg = __yyfmt__.Sprintf("unexpected %s, %s", ls, msg)
					}
				}
			}
			if msg == "" {
				msg = "syntax error"
			}
			yylex.Error(msg)
			Nerrs++
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				row := yyParseTab[yyS[yyp].yys]
				if yyError < len(row) {
					yyn = int(row[yyError]) + yyTabOfs
					if yyn > 0 { // hit
						if yyDebug >= 2 {
							__yyfmt__.Printf("error recovery found error shift in state %d\n", yyS[yyp].yys)
						}
						yystate = yyn /* simulate a shift of "error" */
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
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery failed\n")
			}
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yySymName(yychar))
			}
			if yychar == yyEofCode {
				goto ret1
			}

			yychar = -1
			goto yynewstate /* try again in the same state */
		}
	}

	r := -yyn
	x0 := yyReductions[r]
	x, n := x0.xsym, x0.components
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= n
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	exState := yystate
	yystate = int(yyParseTab[yyS[yyp].yys][x]) + yyTabOfs
	/* reduction by production r */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce using rule %v (%s), and goto state %d\n", r, yySymNames[x], yystate)
	}

	switch r {
	case 1:
		{
			lx := yylex.(*lexer)
			lhs := &File{
				Prologue:         yyS[yypt-1].node.(*Prologue),
				TopLevelDeclList: yyS[yypt-0].node.(*TopLevelDeclList).reverse(),
			}
			yyVAL.node = lhs
			lhs.DotImports = lx.dotImports
			lhs.Path = lx.name
			lhs.Scope = lx.fileScope
			lhs.UnboundImports = lx.unboundImports
			lx.pkg.Files = append(lx.pkg.Files, lhs)
		}
	case 2:
		{
			yyVAL.node = &Argument{
				Expression: yyS[yypt-0].node.(*Expression),
			}
		}
	case 3:
		{
			yyVAL.node = &Argument{
				Case:        1,
				TypeLiteral: yyS[yypt-0].node.(*TypeLiteral),
			}
		}
	case 4:
		{
			yyVAL.node = &ArgumentList{
				Argument: yyS[yypt-0].node.(*Argument),
			}
		}
	case 5:
		{
			yyVAL.node = &ArgumentList{
				Case:         1,
				ArgumentList: yyS[yypt-2].node.(*ArgumentList),
				Token:        yyS[yypt-1].Token,
				Argument:     yyS[yypt-0].node.(*Argument),
			}
		}
	case 6:
		{
			yyVAL.node = &ArrayType{
				Token:  yyS[yypt-3].Token,
				Token2: yyS[yypt-2].Token,
				Token3: yyS[yypt-1].Token,
				Typ:    yyS[yypt-0].node.(*Typ),
			}
		}
	case 7:
		{
			yyVAL.node = &ArrayType{
				Case:       1,
				Token:      yyS[yypt-3].Token,
				Expression: yyS[yypt-2].node.(*Expression),
				Token2:     yyS[yypt-1].Token,
				Typ:        yyS[yypt-0].node.(*Typ),
			}
		}
	case 8:
		{
			yyVAL.node = &Assignment{
				ExpressionList:  yyS[yypt-2].node.(*ExpressionList).reverse(),
				Token:           yyS[yypt-1].Token,
				ExpressionList2: yyS[yypt-0].node.(*ExpressionList).reverse(),
			}
		}
	case 9:
		{
			yyVAL.node = &Assignment{
				Case:            1,
				ExpressionList:  yyS[yypt-2].node.(*ExpressionList).reverse(),
				Token:           yyS[yypt-1].Token,
				ExpressionList2: yyS[yypt-0].node.(*ExpressionList).reverse(),
			}
		}
	case 10:
		{
			yyVAL.node = &Assignment{
				Case:            2,
				ExpressionList:  yyS[yypt-2].node.(*ExpressionList).reverse(),
				Token:           yyS[yypt-1].Token,
				ExpressionList2: yyS[yypt-0].node.(*ExpressionList).reverse(),
			}
		}
	case 11:
		{
			yyVAL.node = &Assignment{
				Case:            3,
				ExpressionList:  yyS[yypt-2].node.(*ExpressionList).reverse(),
				Token:           yyS[yypt-1].Token,
				ExpressionList2: yyS[yypt-0].node.(*ExpressionList).reverse(),
			}
		}
	case 12:
		{
			yyVAL.node = &Assignment{
				Case:            4,
				ExpressionList:  yyS[yypt-2].node.(*ExpressionList).reverse(),
				Token:           yyS[yypt-1].Token,
				ExpressionList2: yyS[yypt-0].node.(*ExpressionList).reverse(),
			}
		}
	case 13:
		{
			yyVAL.node = &Assignment{
				Case:            5,
				ExpressionList:  yyS[yypt-2].node.(*ExpressionList).reverse(),
				Token:           yyS[yypt-1].Token,
				ExpressionList2: yyS[yypt-0].node.(*ExpressionList).reverse(),
			}
		}
	case 14:
		{
			yyVAL.node = &Assignment{
				Case:            6,
				ExpressionList:  yyS[yypt-2].node.(*ExpressionList).reverse(),
				Token:           yyS[yypt-1].Token,
				ExpressionList2: yyS[yypt-0].node.(*ExpressionList).reverse(),
			}
		}
	case 15:
		{
			yyVAL.node = &Assignment{
				Case:            7,
				ExpressionList:  yyS[yypt-2].node.(*ExpressionList).reverse(),
				Token:           yyS[yypt-1].Token,
				ExpressionList2: yyS[yypt-0].node.(*ExpressionList).reverse(),
			}
		}
	case 16:
		{
			yyVAL.node = &Assignment{
				Case:            8,
				ExpressionList:  yyS[yypt-2].node.(*ExpressionList).reverse(),
				Token:           yyS[yypt-1].Token,
				ExpressionList2: yyS[yypt-0].node.(*ExpressionList).reverse(),
			}
		}
	case 17:
		{
			yyVAL.node = &Assignment{
				Case:            9,
				ExpressionList:  yyS[yypt-2].node.(*ExpressionList).reverse(),
				Token:           yyS[yypt-1].Token,
				ExpressionList2: yyS[yypt-0].node.(*ExpressionList).reverse(),
			}
		}
	case 18:
		{
			yyVAL.node = &Assignment{
				Case:            10,
				ExpressionList:  yyS[yypt-2].node.(*ExpressionList).reverse(),
				Token:           yyS[yypt-1].Token,
				ExpressionList2: yyS[yypt-0].node.(*ExpressionList).reverse(),
			}
		}
	case 19:
		{
			yyVAL.node = &Assignment{
				Case:            11,
				ExpressionList:  yyS[yypt-2].node.(*ExpressionList).reverse(),
				Token:           yyS[yypt-1].Token,
				ExpressionList2: yyS[yypt-0].node.(*ExpressionList).reverse(),
			}
		}
	case 20:
		{
			lx := yylex.(*lexer)
			lhs := &BasicLiteral{
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			t := lhs.Token
			s := string(t.S())
			s2 := s[1 : len(s)-1]
			r, _, _, err := strconv.UnquoteChar(s2, '\'')
			if err != nil {
				lx.err(t, "invalid character constant")
				break
			}

			if isSurrogate(r) || r > unicode.MaxRune {
				lx.err(t, "invalid Unicode code point in escape sequence: %#x", r)
				break
			}

			lhs.Value = newConstValue(newRuneConst(r, nil, lx.int32Type, true))
		}
	case 21:
		{
			lx := yylex.(*lexer)
			lhs := &BasicLiteral{
				Case:  1,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			t := lhs.Token
			s := string(t.S())
			if strings.Contains(s, "p") { // Error already reported by the scanner.
				break
			}

			f, _, err := big.ParseFloat(s, 10, lx.floatConstPrec, big.ToNearestEven)
			if err != nil {
				lx.err(t, "%s %s", err, t.S())
				break
			}

			lhs.Value = newConstValue(newFloatConst(0, f, lx.float64Type, true).normalize())
		}
	case 22:
		{
			lx := yylex.(*lexer)
			lhs := &BasicLiteral{
				Case:  2,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			t := lhs.Token
			s := string(t.S())
			if strings.Contains(s, "p") { // Error already reported by the scanner.
				break
			}

			s = s[:len(s)-1] // Remove final "i".
			f, _, err := big.ParseFloat(s, 10, lx.floatConstPrec, big.ToNearestEven)
			if err != nil {
				lx.err(t, "%s %s", err, t.S())
				break
			}

			lhs.Value = newConstValue(newComplexConst(0, &bigComplex{big.NewFloat(0).SetPrec(lx.floatConstPrec), f}, lx.float64Type, true).normalize())
		}
	case 23:
		{
			lx := yylex.(*lexer)
			lhs := &BasicLiteral{
				Case:  3,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			t := lhs.Token
			s := string(t.S())
			var z big.Int
			i, ok := z.SetString(s, 0)
			if !ok {
				lx.err(t, "invalid integer literal %s", s)
				break
			}

			if lhs.Value = newConstValue(newIntConst(0, i, lx.intType, true).normalize()); lhs.Value == nil {
				lx.err(t, "integer literal overflow %s", s)
			}
		}
	case 24:
		{
			lx := yylex.(*lexer)
			lhs := &BasicLiteral{
				Case:  4,
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.stringValue = stringLiteralValue(lx, lhs.Token)
			lhs.Value = newConstValue(newStringConst(lhs.stringValue, lx.stringType, true))
		}
	case 25:
		{
			lx := yylex.(*lexer)
			if !lx.scope.isMergeScope {
				lx.pushScope()
			}
			lx.scope.isMergeScope = false
		}
	case 26:
		{
			lx := yylex.(*lexer)
			yyVAL.node = &Block{
				Token:         yyS[yypt-3].Token,
				StatementList: yyS[yypt-1].node.(*StatementList).reverse(),
				Token2:        yyS[yypt-0].Token,
			}
			lx.popScope()
		}
	case 27:
		{
			lx := yylex.(*lexer)
			lx.pushScope()
		}
	case 28:
		{
			lx := yylex.(*lexer)
			yyVAL.node = &Body{
				Token:         yyS[yypt-3].Token,
				StatementList: yyS[yypt-1].node.(*StatementList).reverse(),
				Token2:        yyS[yypt-0].Token,
			}
			lx.popScope()
		}
	case 29:
		{
			yyVAL.node = &Call{
				Token:  yyS[yypt-1].Token,
				Token2: yyS[yypt-0].Token,
			}
		}
	case 30:
		{
			yyVAL.node = &Call{
				Case:         1,
				Token:        yyS[yypt-3].Token,
				ArgumentList: yyS[yypt-2].node.(*ArgumentList).reverse(),
				CommaOpt:     yyS[yypt-1].node.(*CommaOpt),
				Token2:       yyS[yypt-0].Token,
			}
		}
	case 31:
		{
			yyVAL.node = &Call{
				Case:         2,
				Token:        yyS[yypt-4].Token,
				ArgumentList: yyS[yypt-3].node.(*ArgumentList).reverse(),
				Token2:       yyS[yypt-2].Token,
				CommaOpt:     yyS[yypt-1].node.(*CommaOpt),
				Token3:       yyS[yypt-0].Token,
			}
		}
	case 32:
		{
			yyVAL.node = &ChanType{
				Token: yyS[yypt-1].Token,
				Typ:   yyS[yypt-0].node.(*Typ),
			}
		}
	case 33:
		{
			yyVAL.node = &ChanType{
				Case:   1,
				Token:  yyS[yypt-2].Token,
				Token2: yyS[yypt-1].Token,
				Typ:    yyS[yypt-0].node.(*Typ),
			}
		}
	case 34:
		{
			yyVAL.node = &ChanType{
				Case:   2,
				Token:  yyS[yypt-2].Token,
				Token2: yyS[yypt-1].Token,
				Typ:    yyS[yypt-0].node.(*Typ),
			}
		}
	case 35:
		{
			yyVAL.node = (*CommaOpt)(nil)
		}
	case 36:
		{
			yyVAL.node = &CommaOpt{
				Token: yyS[yypt-0].Token,
			}
		}
	case 37:
		{
			yyVAL.node = &CompLitItem{
				CompLitValue: yyS[yypt-0].node.(*CompLitValue),
			}
		}
	case 38:
		{
			yyVAL.node = &CompLitItem{
				Case:          1,
				CompLitValue:  yyS[yypt-2].node.(*CompLitValue),
				Token:         yyS[yypt-1].Token,
				CompLitValue2: yyS[yypt-0].node.(*CompLitValue),
			}
		}
	case 39:
		{
			yyVAL.node = &CompLitItem{
				Case:         2,
				CompLitValue: yyS[yypt-2].node.(*CompLitValue),
				Token:        yyS[yypt-1].Token,
				Expression:   yyS[yypt-0].node.(*Expression),
			}
		}
	case 40:
		{
			yyVAL.node = &CompLitItem{
				Case:       3,
				Expression: yyS[yypt-0].node.(*Expression),
			}
		}
	case 41:
		{
			yyVAL.node = &CompLitItem{
				Case:         4,
				Expression:   yyS[yypt-2].node.(*Expression),
				Token:        yyS[yypt-1].Token,
				CompLitValue: yyS[yypt-0].node.(*CompLitValue),
			}
		}
	case 42:
		{
			yyVAL.node = &CompLitItem{
				Case:        5,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 43:
		{
			yyVAL.node = &CompLitItemList{
				CompLitItem: yyS[yypt-0].node.(*CompLitItem),
			}
		}
	case 44:
		{
			yyVAL.node = &CompLitItemList{
				Case:            1,
				CompLitItemList: yyS[yypt-2].node.(*CompLitItemList),
				Token:           yyS[yypt-1].Token,
				CompLitItem:     yyS[yypt-0].node.(*CompLitItem),
			}
		}
	case 45:
		{
			yyVAL.node = &CompLitType{
				ArrayType: yyS[yypt-0].node.(*ArrayType),
			}
		}
	case 46:
		{
			yyVAL.node = &CompLitType{
				Case:    1,
				MapType: yyS[yypt-0].node.(*MapType),
			}
		}
	case 47:
		{
			yyVAL.node = &CompLitType{
				Case:      2,
				SliceType: yyS[yypt-0].node.(*SliceType),
			}
		}
	case 48:
		{
			yyVAL.node = &CompLitType{
				Case:       3,
				StructType: yyS[yypt-0].node.(*StructType),
			}
		}
	case 49:
		{
			yyVAL.node = &CompLitValue{
				Token:  yyS[yypt-1].Token,
				Token2: yyS[yypt-0].Token,
			}
		}
	case 50:
		{
			yyVAL.node = &CompLitValue{
				Case:            1,
				Token:           yyS[yypt-3].Token,
				CompLitItemList: yyS[yypt-2].node.(*CompLitItemList).reverse(),
				CommaOpt:        yyS[yypt-1].node.(*CommaOpt),
				Token2:          yyS[yypt-0].Token,
			}
		}
	case 51:
		{
			yyVAL.node = &ConstDecl{
				Token:  yyS[yypt-2].Token,
				Token2: yyS[yypt-1].Token,
				Token3: yyS[yypt-0].Token,
			}
		}
	case 52:
		{
			yyVAL.node = &ConstDecl{
				Case:          1,
				Token:         yyS[yypt-4].Token,
				Token2:        yyS[yypt-3].Token,
				ConstSpecList: yyS[yypt-2].node.(*ConstSpecList).reverse(),
				SemicolonOpt:  yyS[yypt-1].node.(*SemicolonOpt),
				Token3:        yyS[yypt-0].Token,
			}
		}
	case 53:
		{
			yyVAL.node = &ConstDecl{
				Case:      2,
				Token:     yyS[yypt-1].Token,
				ConstSpec: yyS[yypt-0].node.(*ConstSpec),
			}
		}
	case 54:
		{
			lx := yylex.(*lexer)
			lhs := &ConstSpec{
				IdentifierList: yyS[yypt-0].node.(*IdentifierList).reverse(),
			}
			yyVAL.node = lhs
			lhs.decl(lx, nil, nil)
		}
	case 55:
		{
			lx := yylex.(*lexer)
			lhs := &ConstSpec{
				Case:           1,
				IdentifierList: yyS[yypt-2].node.(*IdentifierList).reverse(),
				Token:          yyS[yypt-1].Token,
				ExpressionList: yyS[yypt-0].node.(*ExpressionList).reverse(),
			}
			yyVAL.node = lhs
			lhs.decl(lx, nil, lhs.ExpressionList)
		}
	case 56:
		{
			lx := yylex.(*lexer)
			lhs := &ConstSpec{
				Case:           2,
				IdentifierList: yyS[yypt-3].node.(*IdentifierList).reverse(),
				Typ:            yyS[yypt-2].node.(*Typ),
				Token:          yyS[yypt-1].Token,
				ExpressionList: yyS[yypt-0].node.(*ExpressionList).reverse(),
			}
			yyVAL.node = lhs
			lhs.decl(lx, lhs.Typ, lhs.ExpressionList)
		}
	case 57:
		{
			yyVAL.node = &ConstSpecList{
				ConstSpec: yyS[yypt-0].node.(*ConstSpec),
			}
		}
	case 58:
		{
			yyVAL.node = &ConstSpecList{
				Case:          1,
				ConstSpecList: yyS[yypt-2].node.(*ConstSpecList),
				Token:         yyS[yypt-1].Token,
				ConstSpec:     yyS[yypt-0].node.(*ConstSpec),
			}
		}
	case 59:
		{
			lx := yylex.(*lexer)
			yyVAL.node = &Elif{
				Token:    yyS[yypt-3].Token,
				Token2:   yyS[yypt-2].Token,
				IfHeader: yyS[yypt-1].node.(*IfHeader),
				Body:     yyS[yypt-0].node.(*Body),
			}
			lx.popScope() // Implicit "if" block.
		}
	case 60:
		{
			yyVAL.node = (*ElifList)(nil)
		}
	case 61:
		{
			yyVAL.node = &ElifList{
				ElifList: yyS[yypt-1].node.(*ElifList),
				Elif:     yyS[yypt-0].node.(*Elif),
			}
		}
	case 62:
		{
			yyVAL.node = (*ElseOpt)(nil)
		}
	case 63:
		{
			yyVAL.node = &ElseOpt{
				Token: yyS[yypt-1].Token,
				Block: yyS[yypt-0].node.(*Block),
			}
		}
	case 64:
		{
			yyVAL.node = &Expression{
				UnaryExpression: yyS[yypt-0].node.(*UnaryExpression),
			}
		}
	case 65:
		{
			yyVAL.node = &Expression{
				Case:        1,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 66:
		{
			yyVAL.node = &Expression{
				Case:        2,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 67:
		{
			yyVAL.node = &Expression{
				Case:        3,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 68:
		{
			yyVAL.node = &Expression{
				Case:        4,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 69:
		{
			yyVAL.node = &Expression{
				Case:        5,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 70:
		{
			yyVAL.node = &Expression{
				Case:        6,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 71:
		{
			yyVAL.node = &Expression{
				Case:        7,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 72:
		{
			yyVAL.node = &Expression{
				Case:        8,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 73:
		{
			yyVAL.node = &Expression{
				Case:        9,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 74:
		{
			yyVAL.node = &Expression{
				Case:        10,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 75:
		{
			yyVAL.node = &Expression{
				Case:        11,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 76:
		{
			yyVAL.node = &Expression{
				Case:        12,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 77:
		{
			yyVAL.node = &Expression{
				Case:        13,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 78:
		{
			yyVAL.node = &Expression{
				Case:        14,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 79:
		{
			yyVAL.node = &Expression{
				Case:        15,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 80:
		{
			yyVAL.node = &Expression{
				Case:        16,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 81:
		{
			yyVAL.node = &Expression{
				Case:        17,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 82:
		{
			yyVAL.node = &Expression{
				Case:        18,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 83:
		{
			yyVAL.node = &Expression{
				Case:        19,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 84:
		{
			yyVAL.node = &Expression{
				Case:        20,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 85:
		{
			yyVAL.node = (*ExpressionOpt)(nil)
		}
	case 86:
		{
			yyVAL.node = &ExpressionOpt{
				Expression: yyS[yypt-0].node.(*Expression),
			}
		}
	case 87:
		{
			yyVAL.node = &ExpressionList{
				Expression: yyS[yypt-0].node.(*Expression),
			}
		}
	case 88:
		{
			yyVAL.node = &ExpressionList{
				Case:           1,
				ExpressionList: yyS[yypt-2].node.(*ExpressionList),
				Token:          yyS[yypt-1].Token,
				Expression:     yyS[yypt-0].node.(*Expression),
			}
		}
	case 89:
		{
			yyVAL.node = (*ExpressionListOpt)(nil)
		}
	case 90:
		{
			yyVAL.node = &ExpressionListOpt{
				ExpressionList: yyS[yypt-0].node.(*ExpressionList).reverse(),
			}
		}
	case 91:
		{
			yyVAL.node = &ForHeader{
				Range: yyS[yypt-0].node.(*Range),
			}
		}
	case 92:
		{
			yyVAL.node = &ForHeader{
				Case:                1,
				SimpleStatementOpt:  yyS[yypt-4].node.(*SimpleStatementOpt),
				Token:               yyS[yypt-3].Token,
				SimpleStatementOpt2: yyS[yypt-2].node.(*SimpleStatementOpt),
				Token2:              yyS[yypt-1].Token,
				SimpleStatementOpt3: yyS[yypt-0].node.(*SimpleStatementOpt),
			}
		}
	case 93:
		{
			yyVAL.node = &ForHeader{
				Case:               2,
				SimpleStatementOpt: yyS[yypt-0].node.(*SimpleStatementOpt),
			}
		}
	case 94:
		{
			lx := yylex.(*lexer)
			yyVAL.node = &ForStatement{
				Token:     yyS[yypt-2].Token,
				ForHeader: yyS[yypt-1].node.(*ForHeader),
				Body:      yyS[yypt-0].node.(*Body),
			}
			lx.popScope() // Implicit "for" block.
		}
	case 95:
		{
			lx := yylex.(*lexer)
			yyVAL.node = (*FuncBodyOpt)(nil)
			lx.popScope()
		}
	case 96:
		{
			yyVAL.node = &FuncBodyOpt{
				Block: yyS[yypt-0].node.(*Block),
			}
		}
	case 97:
		{
			lx := yylex.(*lexer)
			p := lx.pkg
			r := yyS[yypt-3].node.(*ReceiverOpt)
			if r != nil {
				r.Parameters.post(lx)
			}
			sig := yyS[yypt-0].node.(*Signature)
			sig.post(lx)
			nm := yyS[yypt-2].Token
			if r == nil { // Function.
				if nm.Val == idInit {
					s := yyS[yypt-0].node.(*Signature)
					if l := s.Parameters.ParameterDeclList; l != nil {
						lx.err(l, "func init must have no arguments and no return values")
					} else if o := s.ResultOpt; o != nil {
						lx.err(o, "func init must have no arguments and no return values")
					}
				}
				p.Scope.declare(lx, newFuncDeclaration(nm, nil, sig, false))
				break
			}

			// Method.
			rx := r.nm
			if !rx.IsValid() {
				break
			}

			d := p.Scope.Bindings[rx.Val]
			switch x := d.(type) {
			case nil: // Forward type
				d = p.forwardTypes[rx.Val]
				if d == nil {
					d = newTypeDeclaration(lx, rx, nil)
					p.forwardTypes.declare(lx, d)
				}
				d.(*TypeDeclaration).declare(lx, newFuncDeclaration(nm, r, sig, false))
			case *TypeDeclaration:
				x.declare(lx, newFuncDeclaration(nm, r, sig, false))
			default:
				lx.err(rx, "%s is not a type", rx.S())
			}
		}
	case 98:
		{
			yyVAL.node = &FuncDecl{
				Token:                yyS[yypt-6].Token,
				ReceiverOpt:          yyS[yypt-5].node.(*ReceiverOpt),
				Token2:               yyS[yypt-4].Token,
				GenericParametersOpt: yyS[yypt-3].node.(*GenericParametersOpt),
				Signature:            yyS[yypt-2].node.(*Signature),
				FuncBodyOpt:          yyS[yypt-0].node.(*FuncBodyOpt),
			}
		}
	case 99:
		{
			lx := yylex.(*lexer)
			lhs := &FuncType{
				Token:     yyS[yypt-1].Token,
				Signature: yyS[yypt-0].node.(*Signature),
			}
			yyVAL.node = lhs
			lhs.Signature.post(lx)
		}
	case 100:
		{
			lx := yylex.(*lexer)
			lhs := &FuncType{
				Case:           1,
				Token:          yyS[yypt-4].Token,
				Token2:         yyS[yypt-3].Token,
				IdentifierList: yyS[yypt-2].node.(*IdentifierList).reverse(),
				Token3:         yyS[yypt-1].Token,
				Signature:      yyS[yypt-0].node.(*Signature),
			}
			yyVAL.node = lhs
			lx.err(lhs.Token2, "anonymous functions cannot have generic type parameters")
			lhs.Signature.post(lx)
		}
	case 101:
		{
			yyVAL.node = (*GenericArgumentsOpt)(nil)
		}
	case 102:
		{
			yyVAL.node = &GenericArgumentsOpt{
				Token:    yyS[yypt-2].Token,
				TypeList: yyS[yypt-1].node.(*TypeList).reverse(),
				Token2:   yyS[yypt-0].Token,
			}
		}
	case 103:
		{
			yyVAL.node = (*GenericParametersOpt)(nil)
		}
	case 104:
		{
			yyVAL.node = &GenericParametersOpt{
				Token:          yyS[yypt-2].Token,
				IdentifierList: yyS[yypt-1].node.(*IdentifierList).reverse(),
				Token2:         yyS[yypt-0].Token,
			}
		}
	case 105:
		{
			yyVAL.node = (*IdentifierOpt)(nil)
		}
	case 106:
		{
			yyVAL.node = &IdentifierOpt{
				Token: yyS[yypt-0].Token,
			}
		}
	case 107:
		{
			yyVAL.node = &IdentifierList{
				Token: yyS[yypt-0].Token,
			}
		}
	case 108:
		{
			yyVAL.node = &IdentifierList{
				Case:           1,
				IdentifierList: yyS[yypt-2].node.(*IdentifierList),
				Token:          yyS[yypt-1].Token,
				Token2:         yyS[yypt-0].Token,
			}
		}
	case 109:
		{
			yyVAL.node = &IfHeader{
				SimpleStatementOpt: yyS[yypt-0].node.(*SimpleStatementOpt),
			}
		}
	case 110:
		{
			yyVAL.node = &IfHeader{
				Case:                1,
				SimpleStatementOpt:  yyS[yypt-2].node.(*SimpleStatementOpt),
				Token:               yyS[yypt-1].Token,
				SimpleStatementOpt2: yyS[yypt-0].node.(*SimpleStatementOpt),
			}
		}
	case 111:
		{
			lx := yylex.(*lexer)
			yyVAL.node = &IfStatement{
				Token:    yyS[yypt-4].Token,
				IfHeader: yyS[yypt-3].node.(*IfHeader),
				Body:     yyS[yypt-2].node.(*Body),
				ElifList: yyS[yypt-1].node.(*ElifList).reverse(),
				ElseOpt:  yyS[yypt-0].node.(*ElseOpt),
			}
			lx.popScope() // Implicit "if" block.
		}
	case 112:
		{
			yyVAL.node = &ImportDecl{
				Token:  yyS[yypt-2].Token,
				Token2: yyS[yypt-1].Token,
				Token3: yyS[yypt-0].Token,
			}
		}
	case 113:
		{
			yyVAL.node = &ImportDecl{
				Case:           1,
				Token:          yyS[yypt-4].Token,
				Token2:         yyS[yypt-3].Token,
				ImportSpecList: yyS[yypt-2].node.(*ImportSpecList).reverse(),
				SemicolonOpt:   yyS[yypt-1].node.(*SemicolonOpt),
				Token3:         yyS[yypt-0].Token,
			}
		}
	case 114:
		{
			yyVAL.node = &ImportDecl{
				Case:       2,
				Token:      yyS[yypt-1].Token,
				ImportSpec: yyS[yypt-0].node.(*ImportSpec),
			}
		}
	case 115:
		{
			lx := yylex.(*lexer)
			lhs := &ImportSpec{
				Token:        yyS[yypt-1].Token,
				BasicLiteral: yyS[yypt-0].node.(*BasicLiteral),
			}
			yyVAL.node = lhs
			lhs.post(lx)
		}
	case 116:
		{
			lx := yylex.(*lexer)
			lhs := &ImportSpec{
				Case:          1,
				IdentifierOpt: yyS[yypt-1].node.(*IdentifierOpt),
				BasicLiteral:  yyS[yypt-0].node.(*BasicLiteral),
			}
			yyVAL.node = lhs
			lhs.post(lx)
		}
	case 117:
		{
			yyVAL.node = &ImportSpec{
				Case:         2,
				Token:        yyS[yypt-2].Token,
				BasicLiteral: yyS[yypt-1].node.(*BasicLiteral),
			}
		}
	case 118:
		{
			yyVAL.node = &ImportSpec{
				Case:          3,
				IdentifierOpt: yyS[yypt-2].node.(*IdentifierOpt),
				BasicLiteral:  yyS[yypt-1].node.(*BasicLiteral),
			}
		}
	case 119:
		{
			yyVAL.node = &ImportSpecList{
				ImportSpec: yyS[yypt-0].node.(*ImportSpec),
			}
		}
	case 120:
		{
			yyVAL.node = &ImportSpecList{
				Case:           1,
				ImportSpecList: yyS[yypt-2].node.(*ImportSpecList),
				Token:          yyS[yypt-1].Token,
				ImportSpec:     yyS[yypt-0].node.(*ImportSpec),
			}
		}
	case 121:
		{
			yyVAL.node = (*ImportList)(nil)
		}
	case 122:
		{
			yyVAL.node = &ImportList{
				ImportList: yyS[yypt-2].node.(*ImportList),
				ImportDecl: yyS[yypt-1].node.(*ImportDecl),
				Token:      yyS[yypt-0].Token,
			}
		}
	case 123:
		{
			yyVAL.node = &InterfaceType{
				Token:  yyS[yypt-2].Token,
				LBrace: yyS[yypt-1].node.(*LBrace),
				Token2: yyS[yypt-0].Token,
			}
		}
	case 124:
		{
			lx := yylex.(*lexer)
			s := lx.resolutionScope
			lx.pushScope()
			lx.resolutionScope = s
		}
	case 125:
		{
			lx := yylex.(*lexer)
			lhs := &InterfaceType{
				Case:                    1,
				Token:                   yyS[yypt-5].Token,
				LBrace:                  yyS[yypt-4].node.(*LBrace),
				InterfaceMethodDeclList: yyS[yypt-2].node.(*InterfaceMethodDeclList).reverse(),
				SemicolonOpt:            yyS[yypt-1].node.(*SemicolonOpt),
				Token2:                  yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.methods = lx.scope
			lx.popScope()
		}
	case 126:
		{
			lx := yylex.(*lexer)
			s := lx.resolutionScope
			lx.pushScope()
			lx.resolutionScope = s
		}
	case 127:
		{
			lx := yylex.(*lexer)
			lhs := &InterfaceMethodDecl{
				Token:     yyS[yypt-2].Token,
				Signature: yyS[yypt-0].node.(*Signature),
			}
			yyVAL.node = lhs
			lhs.Signature.post(lx)
			s := lx.resolutionScope
			lx.popScope()
			lx.resolutionScope = s
			lx.scope.declare(lx, newFuncDeclaration(lhs.Token, nil, lhs.Signature, true))
		}
	case 128:
		{
			yyVAL.node = &InterfaceMethodDecl{
				Case:           1,
				QualifiedIdent: yyS[yypt-0].node.(*QualifiedIdent),
			}
		}
	case 129:
		{
			yyVAL.node = &InterfaceMethodDeclList{
				InterfaceMethodDecl: yyS[yypt-0].node.(*InterfaceMethodDecl),
			}
		}
	case 130:
		{
			yyVAL.node = &InterfaceMethodDeclList{
				Case: 1,
				InterfaceMethodDeclList: yyS[yypt-2].node.(*InterfaceMethodDeclList),
				Token:               yyS[yypt-1].Token,
				InterfaceMethodDecl: yyS[yypt-0].node.(*InterfaceMethodDecl),
			}
		}
	case 131:
		{
			lx := yylex.(*lexer)
			yyVAL.node = &LBrace{
				Token: yyS[yypt-0].Token,
			}
			lx.fixLBR()
		}
	case 132:
		{
			yyVAL.node = &LBrace{
				Case:  1,
				Token: yyS[yypt-0].Token,
			}
		}
	case 133:
		{
			yyVAL.node = &LBraceCompLitItem{
				Expression: yyS[yypt-0].node.(*Expression),
			}
		}
	case 134:
		{
			yyVAL.node = &LBraceCompLitItem{
				Case:        1,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 135:
		{
			yyVAL.node = &LBraceCompLitItem{
				Case:               2,
				Expression:         yyS[yypt-2].node.(*Expression),
				Token:              yyS[yypt-1].Token,
				LBraceCompLitValue: yyS[yypt-0].node.(*LBraceCompLitValue),
			}
		}
	case 136:
		{
			yyVAL.node = &LBraceCompLitItem{
				Case:               3,
				LBraceCompLitValue: yyS[yypt-0].node.(*LBraceCompLitValue),
			}
		}
	case 137:
		{
			yyVAL.node = &LBraceCompLitItemList{
				LBraceCompLitItem: yyS[yypt-0].node.(*LBraceCompLitItem),
			}
		}
	case 138:
		{
			yyVAL.node = &LBraceCompLitItemList{
				Case: 1,
				LBraceCompLitItemList: yyS[yypt-2].node.(*LBraceCompLitItemList),
				Token:             yyS[yypt-1].Token,
				LBraceCompLitItem: yyS[yypt-0].node.(*LBraceCompLitItem),
			}
		}
	case 139:
		{
			yyVAL.node = &LBraceCompLitValue{
				LBrace: yyS[yypt-1].node.(*LBrace),
				Token:  yyS[yypt-0].Token,
			}
		}
	case 140:
		{
			yyVAL.node = &LBraceCompLitValue{
				Case:                  1,
				LBrace:                yyS[yypt-3].node.(*LBrace),
				LBraceCompLitItemList: yyS[yypt-2].node.(*LBraceCompLitItemList).reverse(),
				CommaOpt:              yyS[yypt-1].node.(*CommaOpt),
				Token:                 yyS[yypt-0].Token,
			}
		}
	case 141:
		{
			yyVAL.node = &MapType{
				Token:  yyS[yypt-4].Token,
				Token2: yyS[yypt-3].Token,
				Typ:    yyS[yypt-2].node.(*Typ),
				Token3: yyS[yypt-1].Token,
				Typ2:   yyS[yypt-0].node.(*Typ),
			}
		}
	case 142:
		{
			yyVAL.node = &Operand{
				Token:      yyS[yypt-2].Token,
				Expression: yyS[yypt-1].node.(*Expression),
				Token2:     yyS[yypt-0].Token,
			}
		}
	case 143:
		{
			yyVAL.node = &Operand{
				Case:        1,
				Token:       yyS[yypt-2].Token,
				TypeLiteral: yyS[yypt-1].node.(*TypeLiteral),
				Token2:      yyS[yypt-0].Token,
			}
		}
	case 144:
		{
			yyVAL.node = &Operand{
				Case:         2,
				BasicLiteral: yyS[yypt-0].node.(*BasicLiteral),
			}
		}
	case 145:
		{
			lx := yylex.(*lexer)
			lx.scope.isMergeScope = false
		}
	case 146:
		{
			lx := yylex.(*lexer)
			yyVAL.node = &Operand{
				Case:          3,
				FuncType:      yyS[yypt-4].node.(*FuncType),
				LBrace:        yyS[yypt-2].node.(*LBrace),
				StatementList: yyS[yypt-1].node.(*StatementList).reverse(),
				Token:         yyS[yypt-0].Token,
			}
			lx.popScope()
		}
	case 147:
		{
			lx := yylex.(*lexer)
			lhs := &Operand{
				Case:                4,
				Token:               yyS[yypt-1].Token,
				GenericArgumentsOpt: yyS[yypt-0].node.(*GenericArgumentsOpt),
			}
			yyVAL.node = lhs
			lhs.fileScope = lx.fileScope
			lhs.resolutionScope = lx.resolutionScope
		}
	case 148:
		{
			lx := yylex.(*lexer)
			lhs := &PackageClause{
				Token:  yyS[yypt-2].Token,
				Token2: yyS[yypt-1].Token,
				Token3: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			if lx.pkg.parseOnlyName {
				lx.pkg.Name = string(lhs.Token2.S())
				return 0
			}

			if !lx.build { // Build tags not satisfied
				return 0
			}

			lhs.post(lx)
		}
	case 149:
		{
			yyVAL.node = &ParameterDecl{
				Token: yyS[yypt-1].Token,
				Typ:   yyS[yypt-0].node.(*Typ),
			}
		}
	case 150:
		{
			yyVAL.node = &ParameterDecl{
				Case:   1,
				Token:  yyS[yypt-2].Token,
				Token2: yyS[yypt-1].Token,
				Typ:    yyS[yypt-0].node.(*Typ),
			}
		}
	case 151:
		{
			yyVAL.node = &ParameterDecl{
				Case:  2,
				Token: yyS[yypt-1].Token,
				Typ:   yyS[yypt-0].node.(*Typ),
			}
		}
	case 152:
		{
			yyVAL.node = &ParameterDecl{
				Case: 3,
				Typ:  yyS[yypt-0].node.(*Typ),
			}
		}
	case 153:
		{
			yyVAL.node = &ParameterDeclList{
				ParameterDecl: yyS[yypt-0].node.(*ParameterDecl),
			}
		}
	case 154:
		{
			yyVAL.node = &ParameterDeclList{
				Case:              1,
				ParameterDeclList: yyS[yypt-2].node.(*ParameterDeclList),
				Token:             yyS[yypt-1].Token,
				ParameterDecl:     yyS[yypt-0].node.(*ParameterDecl),
			}
		}
	case 155:
		{
			yyVAL.node = &Parameters{
				Token:  yyS[yypt-1].Token,
				Token2: yyS[yypt-0].Token,
			}
		}
	case 156:
		{
			yyVAL.node = &Parameters{
				Case:              1,
				Token:             yyS[yypt-3].Token,
				ParameterDeclList: yyS[yypt-2].node.(*ParameterDeclList).reverse(),
				CommaOpt:          yyS[yypt-1].node.(*CommaOpt),
				Token2:            yyS[yypt-0].Token,
			}
		}
	case 157:
		{
			yyVAL.node = &PrimaryExpression{
				Operand: yyS[yypt-0].node.(*Operand),
			}
		}
	case 158:
		{
			yyVAL.node = &PrimaryExpression{
				Case:               1,
				CompLitType:        yyS[yypt-1].node.(*CompLitType),
				LBraceCompLitValue: yyS[yypt-0].node.(*LBraceCompLitValue),
			}
		}
	case 159:
		{
			yyVAL.node = &PrimaryExpression{
				Case:              2,
				PrimaryExpression: yyS[yypt-4].node.(*PrimaryExpression),
				Token:             yyS[yypt-3].Token,
				Token2:            yyS[yypt-2].Token,
				Token3:            yyS[yypt-1].Token,
				Token4:            yyS[yypt-0].Token,
			}
		}
	case 160:
		{
			yyVAL.node = &PrimaryExpression{
				Case:              3,
				PrimaryExpression: yyS[yypt-4].node.(*PrimaryExpression),
				Token:             yyS[yypt-3].Token,
				Token2:            yyS[yypt-2].Token,
				Typ:               yyS[yypt-1].node.(*Typ),
				Token3:            yyS[yypt-0].Token,
			}
		}
	case 161:
		{
			yyVAL.node = &PrimaryExpression{
				Case:              4,
				PrimaryExpression: yyS[yypt-2].node.(*PrimaryExpression),
				Token:             yyS[yypt-1].Token,
				Token2:            yyS[yypt-0].Token,
			}
		}
	case 162:
		{
			yyVAL.node = &PrimaryExpression{
				Case:              5,
				PrimaryExpression: yyS[yypt-3].node.(*PrimaryExpression),
				Token:             yyS[yypt-2].Token,
				Expression:        yyS[yypt-1].node.(*Expression),
				Token2:            yyS[yypt-0].Token,
			}
		}
	case 163:
		{
			yyVAL.node = &PrimaryExpression{
				Case:              6,
				PrimaryExpression: yyS[yypt-7].node.(*PrimaryExpression),
				Token:             yyS[yypt-6].Token,
				ExpressionOpt:     yyS[yypt-5].node.(*ExpressionOpt),
				Token2:            yyS[yypt-4].Token,
				ExpressionOpt2:    yyS[yypt-3].node.(*ExpressionOpt),
				Token3:            yyS[yypt-2].Token,
				ExpressionOpt3:    yyS[yypt-1].node.(*ExpressionOpt),
				Token4:            yyS[yypt-0].Token,
			}
		}
	case 164:
		{
			yyVAL.node = &PrimaryExpression{
				Case:              7,
				PrimaryExpression: yyS[yypt-5].node.(*PrimaryExpression),
				Token:             yyS[yypt-4].Token,
				ExpressionOpt:     yyS[yypt-3].node.(*ExpressionOpt),
				Token2:            yyS[yypt-2].Token,
				ExpressionOpt2:    yyS[yypt-1].node.(*ExpressionOpt),
				Token3:            yyS[yypt-0].Token,
			}
		}
	case 165:
		{
			yyVAL.node = &PrimaryExpression{
				Case:              8,
				PrimaryExpression: yyS[yypt-1].node.(*PrimaryExpression),
				Call:              yyS[yypt-0].node.(*Call),
			}
		}
	case 166:
		{
			yyVAL.node = &PrimaryExpression{
				Case:              9,
				PrimaryExpression: yyS[yypt-1].node.(*PrimaryExpression),
				CompLitValue:      yyS[yypt-0].node.(*CompLitValue),
			}
		}
	case 167:
		{
			yyVAL.node = &PrimaryExpression{
				Case:        10,
				TypeLiteral: yyS[yypt-4].node.(*TypeLiteral),
				Token:       yyS[yypt-3].Token,
				Expression:  yyS[yypt-2].node.(*Expression),
				CommaOpt:    yyS[yypt-1].node.(*CommaOpt),
				Token2:      yyS[yypt-0].Token,
			}
		}
	case 168:
		{
			lx := yylex.(*lexer)
			lhs := &Prologue{
				PackageClause: yyS[yypt-1].node.(*PackageClause),
				ImportList:    yyS[yypt-0].node.(*ImportList).reverse(),
			}
			yyVAL.node = lhs
			lhs.post(lx)
		}
	case 169:
		{
			yyVAL.node = &QualifiedIdent{
				Token: yyS[yypt-0].Token,
			}
		}
	case 170:
		{
			yyVAL.node = &QualifiedIdent{
				Case:   1,
				Token:  yyS[yypt-2].Token,
				Token2: yyS[yypt-1].Token,
				Token3: yyS[yypt-0].Token,
			}
		}
	case 171:
		{
			yyVAL.node = &Range{
				ExpressionList: yyS[yypt-3].node.(*ExpressionList).reverse(),
				Token:          yyS[yypt-2].Token,
				Token2:         yyS[yypt-1].Token,
				Expression:     yyS[yypt-0].node.(*Expression),
			}
		}
	case 172:
		{
			lx := yylex.(*lexer)
			lhs := &Range{
				Case:           1,
				ExpressionList: yyS[yypt-3].node.(*ExpressionList).reverse(),
				Token:          yyS[yypt-2].Token,
				Token2:         yyS[yypt-1].Token,
				Expression:     yyS[yypt-0].node.(*Expression),
			}
			yyVAL.node = lhs
			varDecl(lx, lhs.ExpressionList, lhs.Expression, nil, ":=", 2, 1)
		}
	case 173:
		{
			yyVAL.node = &Range{
				Case:       2,
				Token:      yyS[yypt-1].Token,
				Expression: yyS[yypt-0].node.(*Expression),
			}
		}
	case 174:
		{
			yyVAL.node = (*ReceiverOpt)(nil)
		}
	case 175:
		{
			lx := yylex.(*lexer)
			lhs := &ReceiverOpt{
				Parameters: yyS[yypt-0].node.(*Parameters),
			}
			yyVAL.node = lhs
			lhs.resolutionScope = lx.resolutionScope
			lhs.post(lx)
		}
	case 176:
		{
			yyVAL.node = (*ResultOpt)(nil)
		}
	case 177:
		{
			yyVAL.node = &ResultOpt{
				Case:       1,
				Parameters: yyS[yypt-0].node.(*Parameters),
			}
		}
	case 178:
		{
			yyVAL.node = &ResultOpt{
				Case: 2,
				Typ:  yyS[yypt-0].node.(*Typ),
			}
		}
	case 179:
		{
			lx := yylex.(*lexer)
			lx.switchDecl = append(lx.switchDecl, xc.Token{})
		}
	case 180:
		{
			lx := yylex.(*lexer)
			yyVAL.node = &SelectStatement{
				Token:      yyS[yypt-2].Token,
				SwitchBody: yyS[yypt-0].node.(*SwitchBody),
			}
			lx.switchDecl = lx.switchDecl[:len(lx.switchDecl)-1]
		}
	case 181:
		{
			yyVAL.node = (*SemicolonOpt)(nil)
		}
	case 182:
		{
			yyVAL.node = &SemicolonOpt{
				Token: yyS[yypt-0].Token,
			}
		}
	case 183:
		{
			yyVAL.node = &Signature{
				Parameters: yyS[yypt-1].node.(*Parameters),
				ResultOpt:  yyS[yypt-0].node.(*ResultOpt),
			}
		}
	case 184:
		{
			yyVAL.node = &SimpleStatement{
				Assignment: yyS[yypt-0].node.(*Assignment),
			}
		}
	case 185:
		{
			yyVAL.node = &SimpleStatement{
				Case:       1,
				Expression: yyS[yypt-0].node.(*Expression),
			}
		}
	case 186:
		{
			yyVAL.node = &SimpleStatement{
				Case:       2,
				Expression: yyS[yypt-1].node.(*Expression),
				Token:      yyS[yypt-0].Token,
			}
		}
	case 187:
		{
			yyVAL.node = &SimpleStatement{
				Case:       3,
				Expression: yyS[yypt-1].node.(*Expression),
				Token:      yyS[yypt-0].Token,
			}
		}
	case 188:
		{
			lx := yylex.(*lexer)
			lhs := &SimpleStatement{
				Case:            4,
				ExpressionList:  yyS[yypt-2].node.(*ExpressionList).reverse(),
				Token:           yyS[yypt-1].Token,
				ExpressionList2: yyS[yypt-0].node.(*ExpressionList).reverse(),
			}
			yyVAL.node = lhs
			varDecl(lx, lhs.ExpressionList, lhs.ExpressionList2, nil, ":=", -1, -1)
		}
	case 189:
		{
			yyVAL.node = (*SimpleStatementOpt)(nil)
		}
	case 190:
		{
			yyVAL.node = &SimpleStatementOpt{
				SimpleStatement: yyS[yypt-0].node.(*SimpleStatement),
			}
		}
	case 191:
		{
			yyVAL.node = &SliceType{
				Token:  yyS[yypt-2].Token,
				Token2: yyS[yypt-1].Token,
				Typ:    yyS[yypt-0].node.(*Typ),
			}
		}
	case 192:
		{
			yyVAL.node = (*Statement)(nil)
		}
	case 193:
		{
			yyVAL.node = &Statement{
				Case:  1,
				Block: yyS[yypt-0].node.(*Block),
			}
		}
	case 194:
		{
			yyVAL.node = &Statement{
				Case:      2,
				ConstDecl: yyS[yypt-0].node.(*ConstDecl),
			}
		}
	case 195:
		{
			yyVAL.node = &Statement{
				Case:     3,
				TypeDecl: yyS[yypt-0].node.(*TypeDecl),
			}
		}
	case 196:
		{
			yyVAL.node = &Statement{
				Case:    4,
				VarDecl: yyS[yypt-0].node.(*VarDecl),
			}
		}
	case 197:
		{
			yyVAL.node = &Statement{
				Case:             5,
				StatementNonDecl: yyS[yypt-0].node.(*StatementNonDecl),
			}
		}
	case 198:
		{
			yyVAL.node = &Statement{
				Case: 6,
			}
		}
	case 199:
		{
			yyVAL.node = &StatementList{
				Statement: yyS[yypt-0].node.(*Statement),
			}
		}
	case 200:
		{
			yyVAL.node = &StatementList{
				Case:          1,
				StatementList: yyS[yypt-2].node.(*StatementList),
				Token:         yyS[yypt-1].Token,
				Statement:     yyS[yypt-0].node.(*Statement),
			}
		}
	case 201:
		{
			yyVAL.node = &StatementNonDecl{
				Token:         yyS[yypt-1].Token,
				IdentifierOpt: yyS[yypt-0].node.(*IdentifierOpt),
			}
		}
	case 202:
		{
			yyVAL.node = &StatementNonDecl{
				Case:          1,
				Token:         yyS[yypt-1].Token,
				IdentifierOpt: yyS[yypt-0].node.(*IdentifierOpt),
			}
		}
	case 203:
		{
			yyVAL.node = &StatementNonDecl{
				Case:       2,
				Token:      yyS[yypt-1].Token,
				Expression: yyS[yypt-0].node.(*Expression),
			}
		}
	case 204:
		{
			yyVAL.node = &StatementNonDecl{
				Case:  3,
				Token: yyS[yypt-0].Token,
			}
		}
	case 205:
		{
			yyVAL.node = &StatementNonDecl{
				Case:         4,
				ForStatement: yyS[yypt-0].node.(*ForStatement),
			}
		}
	case 206:
		{
			yyVAL.node = &StatementNonDecl{
				Case:       5,
				Token:      yyS[yypt-1].Token,
				Expression: yyS[yypt-0].node.(*Expression),
			}
		}
	case 207:
		{
			yyVAL.node = &StatementNonDecl{
				Case:   6,
				Token:  yyS[yypt-1].Token,
				Token2: yyS[yypt-0].Token,
			}
		}
	case 208:
		{
			lx := yylex.(*lexer)
			lhs := &StatementNonDecl{
				Case:      7,
				Token:     yyS[yypt-2].Token,
				Token2:    yyS[yypt-1].Token,
				Statement: yyS[yypt-0].node.(*Statement),
			}
			yyVAL.node = lhs
			lx.scope.declare(lx, newLabelDeclaration(lhs.Token))
		}
	case 209:
		{
			yyVAL.node = &StatementNonDecl{
				Case:        8,
				IfStatement: yyS[yypt-0].node.(*IfStatement),
			}
		}
	case 210:
		{
			yyVAL.node = &StatementNonDecl{
				Case:              9,
				Token:             yyS[yypt-1].Token,
				ExpressionListOpt: yyS[yypt-0].node.(*ExpressionListOpt),
			}
		}
	case 211:
		{
			yyVAL.node = &StatementNonDecl{
				Case:            10,
				SelectStatement: yyS[yypt-0].node.(*SelectStatement),
			}
		}
	case 212:
		{
			yyVAL.node = &StatementNonDecl{
				Case:            11,
				SimpleStatement: yyS[yypt-0].node.(*SimpleStatement),
			}
		}
	case 213:
		{
			yyVAL.node = &StatementNonDecl{
				Case:            12,
				SwitchStatement: yyS[yypt-0].node.(*SwitchStatement),
			}
		}
	case 214:
		{
			lx := yylex.(*lexer)
			lhs := &StructFieldDecl{
				Token:          yyS[yypt-2].Token,
				QualifiedIdent: yyS[yypt-1].node.(*QualifiedIdent),
				TagOpt:         yyS[yypt-0].node.(*TagOpt),
			}
			yyVAL.node = lhs
			lhs.decl(lx)
		}
	case 215:
		{
			lx := yylex.(*lexer)
			lhs := &StructFieldDecl{
				Case:           1,
				IdentifierList: yyS[yypt-2].node.(*IdentifierList).reverse(),
				Typ:            yyS[yypt-1].node.(*Typ),
				TagOpt:         yyS[yypt-0].node.(*TagOpt),
			}
			yyVAL.node = lhs
			lhs.decl(lx)
		}
	case 216:
		{
			lx := yylex.(*lexer)
			lhs := &StructFieldDecl{
				Case:           2,
				QualifiedIdent: yyS[yypt-1].node.(*QualifiedIdent),
				TagOpt:         yyS[yypt-0].node.(*TagOpt),
			}
			yyVAL.node = lhs
			lhs.decl(lx)
		}
	case 217:
		{
			lx := yylex.(*lexer)
			lhs := &StructFieldDecl{
				Case:           3,
				Token:          yyS[yypt-3].Token,
				QualifiedIdent: yyS[yypt-2].node.(*QualifiedIdent),
				Token2:         yyS[yypt-1].Token,
				TagOpt:         yyS[yypt-0].node.(*TagOpt),
			}
			yyVAL.node = lhs
			lhs.decl(lx)
		}
	case 218:
		{
			lx := yylex.(*lexer)
			lhs := &StructFieldDecl{
				Case:           4,
				Token:          yyS[yypt-4].Token,
				Token2:         yyS[yypt-3].Token,
				QualifiedIdent: yyS[yypt-2].node.(*QualifiedIdent),
				Token3:         yyS[yypt-1].Token,
				TagOpt:         yyS[yypt-0].node.(*TagOpt),
			}
			yyVAL.node = lhs
			lhs.decl(lx)
		}
	case 219:
		{
			lx := yylex.(*lexer)
			lhs := &StructFieldDecl{
				Case:           5,
				Token:          yyS[yypt-4].Token,
				Token2:         yyS[yypt-3].Token,
				QualifiedIdent: yyS[yypt-2].node.(*QualifiedIdent),
				Token3:         yyS[yypt-1].Token,
				TagOpt:         yyS[yypt-0].node.(*TagOpt),
			}
			yyVAL.node = lhs
			lhs.decl(lx)
		}
	case 220:
		{
			yyVAL.node = &StructFieldDeclList{
				StructFieldDecl: yyS[yypt-0].node.(*StructFieldDecl),
			}
		}
	case 221:
		{
			yyVAL.node = &StructFieldDeclList{
				Case:                1,
				StructFieldDeclList: yyS[yypt-2].node.(*StructFieldDeclList),
				Token:               yyS[yypt-1].Token,
				StructFieldDecl:     yyS[yypt-0].node.(*StructFieldDecl),
			}
		}
	case 222:
		{
			yyVAL.node = &StructType{
				Token:  yyS[yypt-2].Token,
				LBrace: yyS[yypt-1].node.(*LBrace),
				Token2: yyS[yypt-0].Token,
			}
		}
	case 223:
		{
			lx := yylex.(*lexer)
			lx.pushScope()
			lx.resolutionScope = lx.scope.Parent
		}
	case 224:
		{
			lx := yylex.(*lexer)
			lhs := &StructType{
				Case:                1,
				Token:               yyS[yypt-5].Token,
				LBrace:              yyS[yypt-4].node.(*LBrace),
				StructFieldDeclList: yyS[yypt-2].node.(*StructFieldDeclList).reverse(),
				SemicolonOpt:        yyS[yypt-1].node.(*SemicolonOpt),
				Token2:              yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.fields = lx.scope
			lx.popScope()
		}
	case 225:
		{
			yyVAL.node = &SwitchBody{
				Token:  yyS[yypt-1].Token,
				Token2: yyS[yypt-0].Token,
			}
		}
	case 226:
		{
			lx := yylex.(*lexer)
			lx.pushScope()
		}
	case 227:
		{
			lx := yylex.(*lexer)
			yyVAL.node = &SwitchBody{
				Case:           1,
				Token:          yyS[yypt-3].Token,
				SwitchCaseList: yyS[yypt-1].node.(*SwitchCaseList).reverse(),
				Token2:         yyS[yypt-0].Token,
			}
			lx.popScope()
		}
	case 228:
		{
			yyVAL.node = &SwitchCase{
				Token:        yyS[yypt-2].Token,
				ArgumentList: yyS[yypt-1].node.(*ArgumentList).reverse(),
				Token2:       yyS[yypt-0].Token,
			}
		}
	case 229:
		{
			yyVAL.node = &SwitchCase{
				Case:         1,
				Token:        yyS[yypt-4].Token,
				ArgumentList: yyS[yypt-3].node.(*ArgumentList).reverse(),
				Token2:       yyS[yypt-2].Token,
				Expression:   yyS[yypt-1].node.(*Expression),
				Token3:       yyS[yypt-0].Token,
			}
		}
	case 230:
		{
			lx := yylex.(*lexer)
			lhs := &SwitchCase{
				Case:         2,
				Token:        yyS[yypt-4].Token,
				ArgumentList: yyS[yypt-3].node.(*ArgumentList).reverse(),
				Token2:       yyS[yypt-2].Token,
				Expression:   yyS[yypt-1].node.(*Expression),
				Token3:       yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			varDecl(lx, lhs.ArgumentList, lhs.Expression, nil, ":=", 2, 1)
		}
	case 231:
		{
			yyVAL.node = &SwitchCase{
				Case:   3,
				Token:  yyS[yypt-1].Token,
				Token2: yyS[yypt-0].Token,
			}
		}
	case 232:
		{
			yyVAL.node = &SwitchCase{
				Case:  4,
				Token: yyS[yypt-1].Token,
			}
		}
	case 233:
		{
			yyVAL.node = &SwitchCase{
				Case:  5,
				Token: yyS[yypt-1].Token,
			}
		}
	case 234:
		{
			lx := yylex.(*lexer)
			yyVAL.node = &SwitchCaseBlock{
				SwitchCase:    yyS[yypt-1].node.(*SwitchCase),
				StatementList: yyS[yypt-0].node.(*StatementList).reverse(),
			}
			lx.popScope() // Implicit "case"/"default" block.
		}
	case 235:
		{
			yyVAL.node = &SwitchCaseList{
				SwitchCaseBlock: yyS[yypt-0].node.(*SwitchCaseBlock),
			}
		}
	case 236:
		{
			yyVAL.node = &SwitchCaseList{
				Case:            1,
				SwitchCaseList:  yyS[yypt-1].node.(*SwitchCaseList),
				SwitchCaseBlock: yyS[yypt-0].node.(*SwitchCaseBlock),
			}
		}
	case 237:
		{
			yyVAL.node = &SwitchHeader{
				SimpleStatementOpt: yyS[yypt-0].node.(*SimpleStatementOpt),
			}
		}
	case 238:
		{
			yyVAL.node = &SwitchHeader{
				Case:               1,
				SimpleStatementOpt: yyS[yypt-1].node.(*SimpleStatementOpt),
				Token:              yyS[yypt-0].Token,
			}
		}
	case 239:
		{
			yyVAL.node = &SwitchHeader{
				Case:               2,
				SimpleStatementOpt: yyS[yypt-2].node.(*SimpleStatementOpt),
				Token:              yyS[yypt-1].Token,
				Expression:         yyS[yypt-0].node.(*Expression),
			}
		}
	case 240:
		{
			yyVAL.node = &SwitchHeader{
				Case:               3,
				SimpleStatementOpt: yyS[yypt-8].node.(*SimpleStatementOpt),
				Token:              yyS[yypt-7].Token,
				Token2:             yyS[yypt-6].Token,
				Token3:             yyS[yypt-5].Token,
				PrimaryExpression:  yyS[yypt-4].node.(*PrimaryExpression),
				Token4:             yyS[yypt-3].Token,
				Token5:             yyS[yypt-2].Token,
				Token6:             yyS[yypt-1].Token,
				Token7:             yyS[yypt-0].Token,
			}
		}
	case 241:
		{
			lx := yylex.(*lexer)
			var t xc.Token
			if sh := yyS[yypt-0].node.(*SwitchHeader); sh.Case == 3 { // SimpleStatementOpt ';' IDENTIFIER ":=" PrimaryExpression '.' '(' "type" ')'
				t = sh.Token2 // IDENTIFIER
			}
			lx.switchDecl = append(lx.switchDecl, t)
		}
	case 242:
		{
			lx := yylex.(*lexer)
			yyVAL.node = &SwitchStatement{
				Token:        yyS[yypt-3].Token,
				SwitchHeader: yyS[yypt-2].node.(*SwitchHeader),
				SwitchBody:   yyS[yypt-0].node.(*SwitchBody),
			}
			lx.switchDecl = lx.switchDecl[:len(lx.switchDecl)-1]
			lx.popScope() // Implicit "switch" block.
		}
	case 243:
		{
			yyVAL.node = (*TagOpt)(nil)
		}
	case 244:
		{
			lx := yylex.(*lexer)
			lhs := &TagOpt{
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.stringValue = stringLiteralValue(lx, lhs.Token)
		}
	case 245:
		{
			yyVAL.node = &TopLevelDecl{
				ConstDecl: yyS[yypt-0].node.(*ConstDecl),
			}
		}
	case 246:
		{
			yyVAL.node = &TopLevelDecl{
				Case:     1,
				FuncDecl: yyS[yypt-0].node.(*FuncDecl),
			}
		}
	case 247:
		{
			yyVAL.node = &TopLevelDecl{
				Case:     2,
				TypeDecl: yyS[yypt-0].node.(*TypeDecl),
			}
		}
	case 248:
		{
			yyVAL.node = &TopLevelDecl{
				Case:    3,
				VarDecl: yyS[yypt-0].node.(*VarDecl),
			}
		}
	case 249:
		{
			yyVAL.node = &TopLevelDecl{
				Case:             4,
				StatementNonDecl: yyS[yypt-0].node.(*StatementNonDecl),
			}
		}
	case 250:
		{
			yyVAL.node = &TopLevelDecl{
				Case: 5,
			}
		}
	case 251:
		{
			yyVAL.node = (*TopLevelDeclList)(nil)
		}
	case 252:
		{
			yyVAL.node = &TopLevelDeclList{
				TopLevelDeclList: yyS[yypt-2].node.(*TopLevelDeclList),
				TopLevelDecl:     yyS[yypt-1].node.(*TopLevelDecl),
				Token:            yyS[yypt-0].Token,
			}
		}
	case 253:
		{
			yyVAL.node = &Typ{
				Token:  yyS[yypt-2].Token,
				Typ:    yyS[yypt-1].node.(*Typ),
				Token2: yyS[yypt-0].Token,
			}
		}
	case 254:
		{
			yyVAL.node = &Typ{
				Case:  1,
				Token: yyS[yypt-1].Token,
				Typ:   yyS[yypt-0].node.(*Typ),
			}
		}
	case 255:
		{
			yyVAL.node = &Typ{
				Case:      2,
				ArrayType: yyS[yypt-0].node.(*ArrayType),
			}
		}
	case 256:
		{
			yyVAL.node = &Typ{
				Case:     3,
				ChanType: yyS[yypt-0].node.(*ChanType),
			}
		}
	case 257:
		{
			lx := yylex.(*lexer)
			yyVAL.node = &Typ{
				Case:     4,
				FuncType: yyS[yypt-0].node.(*FuncType),
			}
			lx.popScope()
		}
	case 258:
		{
			yyVAL.node = &Typ{
				Case:          5,
				InterfaceType: yyS[yypt-0].node.(*InterfaceType),
			}
		}
	case 259:
		{
			yyVAL.node = &Typ{
				Case:    6,
				MapType: yyS[yypt-0].node.(*MapType),
			}
		}
	case 260:
		{
			lx := yylex.(*lexer)
			lhs := &Typ{
				Case:                7,
				QualifiedIdent:      yyS[yypt-1].node.(*QualifiedIdent),
				GenericArgumentsOpt: yyS[yypt-0].node.(*GenericArgumentsOpt),
			}
			yyVAL.node = lhs
			lhs.fileScope = lx.fileScope
			lhs.resolutionScope = lx.resolutionScope
		}
	case 261:
		{
			yyVAL.node = &Typ{
				Case:      8,
				SliceType: yyS[yypt-0].node.(*SliceType),
			}
		}
	case 262:
		{
			yyVAL.node = &Typ{
				Case:       9,
				StructType: yyS[yypt-0].node.(*StructType),
			}
		}
	case 263:
		{
			yyVAL.node = &TypeDecl{
				Token:  yyS[yypt-2].Token,
				Token2: yyS[yypt-1].Token,
				Token3: yyS[yypt-0].Token,
			}
		}
	case 264:
		{
			yyVAL.node = &TypeDecl{
				Case:         1,
				Token:        yyS[yypt-4].Token,
				Token2:       yyS[yypt-3].Token,
				TypeSpecList: yyS[yypt-2].node.(*TypeSpecList).reverse(),
				SemicolonOpt: yyS[yypt-1].node.(*SemicolonOpt),
				Token3:       yyS[yypt-0].Token,
			}
		}
	case 265:
		{
			yyVAL.node = &TypeDecl{
				Case:     2,
				Token:    yyS[yypt-1].Token,
				TypeSpec: yyS[yypt-0].node.(*TypeSpec),
			}
		}
	case 266:
		{
			yyVAL.node = &TypeList{
				Typ: yyS[yypt-0].node.(*Typ),
			}
		}
	case 267:
		{
			yyVAL.node = &TypeList{
				Case:     1,
				TypeList: yyS[yypt-2].node.(*TypeList),
				Token:    yyS[yypt-1].Token,
				Typ:      yyS[yypt-0].node.(*Typ),
			}
		}
	case 268:
		{
			yyVAL.node = &TypeLiteral{
				Token:       yyS[yypt-1].Token,
				TypeLiteral: yyS[yypt-0].node.(*TypeLiteral),
			}
		}
	case 269:
		{
			yyVAL.node = &TypeLiteral{
				Case:      1,
				ArrayType: yyS[yypt-0].node.(*ArrayType),
			}
		}
	case 270:
		{
			yyVAL.node = &TypeLiteral{
				Case:     2,
				ChanType: yyS[yypt-0].node.(*ChanType),
			}
		}
	case 271:
		{
			lx := yylex.(*lexer)
			yyVAL.node = &TypeLiteral{
				Case:     3,
				FuncType: yyS[yypt-0].node.(*FuncType),
			}
			lx.popScope()
		}
	case 272:
		{
			yyVAL.node = &TypeLiteral{
				Case:          4,
				InterfaceType: yyS[yypt-0].node.(*InterfaceType),
			}
		}
	case 273:
		{
			yyVAL.node = &TypeLiteral{
				Case:    5,
				MapType: yyS[yypt-0].node.(*MapType),
			}
		}
	case 274:
		{
			yyVAL.node = &TypeLiteral{
				Case:      6,
				SliceType: yyS[yypt-0].node.(*SliceType),
			}
		}
	case 275:
		{
			yyVAL.node = &TypeLiteral{
				Case:       7,
				StructType: yyS[yypt-0].node.(*StructType),
			}
		}
	case 276:
		{
			lx := yylex.(*lexer)
			lhs := &TypeSpec{
				Token:                yyS[yypt-2].Token,
				GenericParametersOpt: yyS[yypt-1].node.(*GenericParametersOpt),
				Typ:                  yyS[yypt-0].node.(*Typ),
			}
			yyVAL.node = lhs
			nm := lhs.Token
			if lx.scope.Kind == PackageScope {
				p := lx.pkg
				if d := p.forwardTypes[nm.Val]; d != nil {
					delete(p.forwardTypes, nm.Val)
					d.(*TypeDeclaration).typ0 = lhs.Typ
					lx.scope.declare(lx, d)
					break
				}
			}

			d := newTypeDeclaration(lx, nm, lhs.Typ)
			if t := lhs.Typ; t.Case == 5 { // InterfaceType
				d.methods = t.InterfaceType.methods
			}
			lx.scope.declare(lx, d)
		}
	case 277:
		{
			yyVAL.node = &TypeSpecList{
				TypeSpec: yyS[yypt-0].node.(*TypeSpec),
			}
		}
	case 278:
		{
			yyVAL.node = &TypeSpecList{
				Case:         1,
				TypeSpecList: yyS[yypt-2].node.(*TypeSpecList),
				Token:        yyS[yypt-1].Token,
				TypeSpec:     yyS[yypt-0].node.(*TypeSpec),
			}
		}
	case 279:
		{
			yyVAL.node = &UnaryExpression{
				Token:           yyS[yypt-1].Token,
				UnaryExpression: yyS[yypt-0].node.(*UnaryExpression),
			}
		}
	case 280:
		{
			yyVAL.node = &UnaryExpression{
				Case:            1,
				Token:           yyS[yypt-1].Token,
				UnaryExpression: yyS[yypt-0].node.(*UnaryExpression),
			}
		}
	case 281:
		{
			yyVAL.node = &UnaryExpression{
				Case:            2,
				Token:           yyS[yypt-1].Token,
				UnaryExpression: yyS[yypt-0].node.(*UnaryExpression),
			}
		}
	case 282:
		{
			yyVAL.node = &UnaryExpression{
				Case:            3,
				Token:           yyS[yypt-1].Token,
				UnaryExpression: yyS[yypt-0].node.(*UnaryExpression),
			}
		}
	case 283:
		{
			yyVAL.node = &UnaryExpression{
				Case:            4,
				Token:           yyS[yypt-1].Token,
				UnaryExpression: yyS[yypt-0].node.(*UnaryExpression),
			}
		}
	case 284:
		{
			yyVAL.node = &UnaryExpression{
				Case:            5,
				Token:           yyS[yypt-1].Token,
				UnaryExpression: yyS[yypt-0].node.(*UnaryExpression),
			}
		}
	case 285:
		{
			yyVAL.node = &UnaryExpression{
				Case:            6,
				Token:           yyS[yypt-1].Token,
				UnaryExpression: yyS[yypt-0].node.(*UnaryExpression),
			}
		}
	case 286:
		{
			yyVAL.node = &UnaryExpression{
				Case:              7,
				PrimaryExpression: yyS[yypt-0].node.(*PrimaryExpression),
			}
		}
	case 287:
		{
			yyVAL.node = &VarDecl{
				Token:  yyS[yypt-2].Token,
				Token2: yyS[yypt-1].Token,
				Token3: yyS[yypt-0].Token,
			}
		}
	case 288:
		{
			yyVAL.node = &VarDecl{
				Case:         1,
				Token:        yyS[yypt-4].Token,
				Token2:       yyS[yypt-3].Token,
				VarSpecList:  yyS[yypt-2].node.(*VarSpecList).reverse(),
				SemicolonOpt: yyS[yypt-1].node.(*SemicolonOpt),
				Token3:       yyS[yypt-0].Token,
			}
		}
	case 289:
		{
			yyVAL.node = &VarDecl{
				Case:    2,
				Token:   yyS[yypt-1].Token,
				VarSpec: yyS[yypt-0].node.(*VarSpec),
			}
		}
	case 290:
		{
			lx := yylex.(*lexer)
			lhs := &VarSpec{
				IdentifierList: yyS[yypt-2].node.(*IdentifierList).reverse(),
				Token:          yyS[yypt-1].Token,
				ExpressionList: yyS[yypt-0].node.(*ExpressionList).reverse(),
			}
			yyVAL.node = lhs
			varDecl(lx, lhs.IdentifierList, lhs.ExpressionList, nil, "=", -1, -1)
		}
	case 291:
		{
			lx := yylex.(*lexer)
			lhs := &VarSpec{
				Case:           1,
				IdentifierList: yyS[yypt-1].node.(*IdentifierList).reverse(),
				Typ:            yyS[yypt-0].node.(*Typ),
			}
			yyVAL.node = lhs
			varDecl(lx, lhs.IdentifierList, nil, lhs.Typ, "", -1, -1)
		}
	case 292:
		{
			lx := yylex.(*lexer)
			lhs := &VarSpec{
				Case:           2,
				IdentifierList: yyS[yypt-3].node.(*IdentifierList).reverse(),
				Typ:            yyS[yypt-2].node.(*Typ),
				Token:          yyS[yypt-1].Token,
				ExpressionList: yyS[yypt-0].node.(*ExpressionList).reverse(),
			}
			yyVAL.node = lhs
			varDecl(lx, lhs.IdentifierList, lhs.ExpressionList, lhs.Typ, "=", -1, -1)
		}
	case 293:
		{
			yyVAL.node = &VarSpecList{
				VarSpec: yyS[yypt-0].node.(*VarSpec),
			}
		}
	case 294:
		{
			yyVAL.node = &VarSpecList{
				Case:        1,
				VarSpecList: yyS[yypt-2].node.(*VarSpecList),
				Token:       yyS[yypt-1].Token,
				VarSpec:     yyS[yypt-0].node.(*VarSpec),
			}
		}

	}

	if yyEx != nil && yyEx.Reduced(r, exState, &yyVAL) {
		return -1
	}
	goto yystack /* stack new state and value */
}
