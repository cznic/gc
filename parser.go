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
	yyTabOfs   = -297
)

var (
	yyXLAT = map[int]int{
		59:    0,   // ';' (253x)
		125:   1,   // '}' (245x)
		40:    2,   // '(' (236x)
		42:    3,   // '*' (227x)
		38:    4,   // '&' (188x)
		43:    5,   // '+' (188x)
		45:    6,   // '-' (188x)
		94:    7,   // '^' (188x)
		57358: 8,   // COMM (188x)
		57354: 9,   // CASE (183x)
		57363: 10,  // DEFAULT (183x)
		91:    11,  // '[' (175x)
		41:    12,  // ')' (173x)
		57377: 13,  // IDENTIFIER (172x)
		44:    14,  // ',' (169x)
		57352: 15,  // BODY (153x)
		57403: 16,  // STRING_LIT (151x)
		57355: 17,  // CHAN (140x)
		57372: 18,  // FUNC (139x)
		57382: 19,  // INTERFACE (139x)
		57388: 20,  // MAP (139x)
		57401: 21,  // RXCHAN (139x)
		57404: 22,  // STRUCT (139x)
		58:    23,  // ':' (118x)
		57424: 24,  // ArrayType (116x)
		57430: 25,  // ChanType (116x)
		57451: 26,  // FuncType (116x)
		57464: 27,  // InterfaceType (116x)
		57469: 28,  // MapType (116x)
		57486: 29,  // SliceType (116x)
		57492: 30,  // StructType (116x)
		61:    31,  // '=' (115x)
		57356: 32,  // CHAR_LIT (111x)
		57370: 33,  // FLOAT_LIT (111x)
		57379: 34,  // IMAG_LIT (111x)
		57383: 35,  // INT_LIT (111x)
		57357: 36,  // COLAS (110x)
		57361: 37,  // DDD (109x)
		33:    38,  // '!' (104x)
		123:   39,  // '{' (104x)
		93:    40,  // ']' (101x)
		57426: 41,  // BasicLiteral (91x)
		57505: 42,  // TypeLiteral (90x)
		57434: 43,  // CompLitType (89x)
		57470: 44,  // Operand (89x)
		57475: 45,  // PrimaryExpression (89x)
		57508: 46,  // UnaryExpression (88x)
		37:    47,  // '%' (84x)
		47:    48,  // '/' (84x)
		60:    49,  // '<' (84x)
		62:    50,  // '>' (84x)
		124:   51,  // '|' (84x)
		57347: 52,  // ANDAND (84x)
		57348: 53,  // ANDNOT (84x)
		57367: 54,  // EQ (84x)
		57373: 55,  // GEQ (84x)
		57384: 56,  // LEQ (84x)
		57385: 57,  // LSH (84x)
		57391: 58,  // NEQ (84x)
		57393: 59,  // OROR (84x)
		57399: 60,  // RSH (84x)
		57442: 61,  // Expression (81x)
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
		57465: 101, // LBrace (8x)
		57436: 102, // ConstDecl (7x)
		57448: 103, // ForStatement (7x)
		57457: 104, // IfStatement (7x)
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
		57468: 119, // LBraceCompLitValue (5x)
		57483: 120, // Signature (5x)
		57422: 121, // Argument (4x)
		57452: 122, // GenericArgumentsOpt (4x)
		57380: 123, // IMPORT (4x)
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
		"ArrayType",
		"ChanType",
		"FuncType",
		"InterfaceType",
		"MapType",
		"SliceType",
		"StructType",
		"'='",
		"CHAR_LIT",
		"FLOAT_LIT",
		"IMAG_LIT",
		"INT_LIT",
		"COLAS",
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
		"LBrace",
		"ConstDecl",
		"ForStatement",
		"IfStatement",
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
		"LBraceCompLitValue",
		"Signature",
		"Argument",
		"GenericArgumentsOpt",
		"IMPORT",
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
		57370: "literal",
		57379: "literal",
		57383: "literal",
		57357: ":=",
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
		2:   {121, 1},
		3:   {121, 1},
		4:   {133, 1},
		5:   {133, 3},
		6:   {24, 4},
		7:   {24, 4},
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
		32:  {25, 2},
		33:  {25, 3},
		34:  {25, 3},
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
		51:  {102, 3},
		52:  {102, 5},
		53:  {102, 2},
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
		94:  {103, 3},
		95:  {164, 0},
		96:  {164, 1},
		97:  {149, 0},
		98:  {165, 7},
		99:  {26, 2},
		100: {26, 5},
		101: {122, 0},
		102: {122, 3},
		103: {136, 0},
		104: {136, 3},
		105: {118, 0},
		106: {118, 1},
		107: {98, 1},
		108: {98, 3},
		109: {137, 1},
		110: {137, 3},
		111: {104, 5},
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
		123: {27, 3},
		124: {150, 0},
		125: {27, 6},
		126: {151, 0},
		127: {138, 3},
		128: {138, 1},
		129: {169, 1},
		130: {169, 3},
		131: {101, 1},
		132: {101, 1},
		133: {139, 1},
		134: {139, 3},
		135: {139, 3},
		136: {139, 1},
		137: {139, 3},
		138: {139, 3},
		139: {170, 1},
		140: {170, 3},
		141: {119, 2},
		142: {119, 4},
		143: {28, 5},
		144: {44, 3},
		145: {44, 3},
		146: {44, 1},
		147: {152, 0},
		148: {44, 5},
		149: {44, 2},
		150: {172, 3},
		151: {129, 2},
		152: {129, 3},
		153: {129, 2},
		154: {129, 1},
		155: {140, 1},
		156: {140, 3},
		157: {105, 2},
		158: {105, 4},
		159: {45, 1},
		160: {45, 2},
		161: {45, 5},
		162: {45, 5},
		163: {45, 3},
		164: {45, 4},
		165: {45, 8},
		166: {45, 6},
		167: {45, 2},
		168: {45, 2},
		169: {45, 5},
		170: {173, 2},
		171: {78, 1},
		172: {78, 3},
		173: {174, 4},
		174: {174, 4},
		175: {174, 2},
		176: {175, 0},
		177: {175, 1},
		178: {141, 0},
		179: {141, 1},
		180: {141, 1},
		181: {153, 0},
		182: {106, 3},
		183: {114, 0},
		184: {114, 1},
		185: {120, 2},
		186: {96, 1},
		187: {96, 1},
		188: {96, 2},
		189: {96, 2},
		190: {96, 3},
		191: {107, 0},
		192: {107, 1},
		193: {29, 3},
		194: {115, 0},
		195: {115, 1},
		196: {115, 1},
		197: {115, 1},
		198: {115, 1},
		199: {115, 1},
		200: {115, 1},
		201: {124, 1},
		202: {124, 3},
		203: {108, 2},
		204: {108, 2},
		205: {108, 2},
		206: {108, 1},
		207: {108, 1},
		208: {108, 2},
		209: {108, 2},
		210: {108, 3},
		211: {108, 1},
		212: {108, 2},
		213: {108, 1},
		214: {108, 1},
		215: {108, 1},
		216: {142, 3},
		217: {142, 3},
		218: {142, 2},
		219: {142, 4},
		220: {142, 5},
		221: {142, 5},
		222: {176, 1},
		223: {176, 3},
		224: {30, 3},
		225: {154, 0},
		226: {30, 6},
		227: {143, 2},
		228: {155, 0},
		229: {143, 4},
		230: {144, 3},
		231: {144, 5},
		232: {144, 5},
		233: {144, 2},
		234: {144, 2},
		235: {144, 2},
		236: {145, 2},
		237: {177, 1},
		238: {177, 2},
		239: {178, 1},
		240: {178, 2},
		241: {178, 3},
		242: {178, 9},
		243: {147, 0},
		244: {109, 4},
		245: {116, 0},
		246: {116, 1},
		247: {179, 1},
		248: {179, 1},
		249: {179, 1},
		250: {179, 1},
		251: {179, 1},
		252: {179, 1},
		253: {180, 0},
		254: {180, 3},
		255: {80, 3},
		256: {80, 2},
		257: {80, 1},
		258: {80, 1},
		259: {80, 1},
		260: {80, 1},
		261: {80, 1},
		262: {80, 2},
		263: {80, 1},
		264: {80, 1},
		265: {110, 3},
		266: {110, 5},
		267: {110, 2},
		268: {182, 1},
		269: {182, 3},
		270: {42, 2},
		271: {42, 1},
		272: {42, 1},
		273: {42, 1},
		274: {42, 1},
		275: {42, 1},
		276: {42, 1},
		277: {42, 1},
		278: {131, 3},
		279: {183, 1},
		280: {183, 3},
		281: {46, 2},
		282: {46, 2},
		283: {46, 2},
		284: {46, 2},
		285: {46, 2},
		286: {46, 2},
		287: {46, 2},
		288: {46, 1},
		289: {111, 3},
		290: {111, 5},
		291: {111, 2},
		292: {132, 3},
		293: {132, 2},
		294: {132, 4},
		295: {184, 1},
		296: {184, 3},
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
		yyXError{504, -1}: "expected ')'",
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
		yyXError{448, -1}: "expected ';'",
		yyXError{449, -1}: "expected ';'",
		yyXError{458, -1}: "expected ';'",
		yyXError{497, -1}: "expected '='",
		yyXError{48, -1}:  "expected '['",
		yyXError{303, -1}: "expected ']'",
		yyXError{409, -1}: "expected ']'",
		yyXError{515, -1}: "expected ']'",
		yyXError{318, -1}: "expected '}'",
		yyXError{348, -1}: "expected '}'",
		yyXError{389, -1}: "expected '}'",
		yyXError{421, -1}: "expected '}'",
		yyXError{290, -1}: "expected argument list or one of ['!', '&', '(', ')', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{46, -1}:  "expected block statement or if statement header or one of ['!', '&', '(', '*', '+', '-', ';', '[', '^', <-, chan, func, identifier, interface, literal, map, struct, {]",
		yyXError{433, -1}: "expected block statement or if statement header or one of ['!', '&', '(', '*', '+', '-', ';', '[', '^', <-, chan, func, identifier, interface, literal, map, struct, {]",
		yyXError{430, -1}: "expected block statement or one of ['{', if]",
		yyXError{426, -1}: "expected block statement or {",
		yyXError{435, -1}: "expected block statement or {",
		yyXError{453, -1}: "expected block statement or {",
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
		yyXError{395, -1}: "expected composite literal value or expression or one of ['!', '&', '(', '*', '+', '-', '[', '^', '{', <-, chan, func, identifier, interface, literal, map, struct, {]",
		yyXError{321, -1}: "expected composite literal value or expression or one of ['!', '&', '(', '*', '+', '-', '[', '^', '{', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{324, -1}: "expected composite literal value or expression or one of ['!', '&', '(', '*', '+', '-', '[', '^', '{', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{54, -1}:  "expected composite literal value or one of ['{', {]",
		yyXError{493, -1}: "expected constant specification list or one of [')', identifier]",
		yyXError{41, -1}:  "expected constant specification or one of ['(', identifier]",
		yyXError{505, -1}: "expected constant specification or one of [')', identifier]",
		yyXError{429, -1}: "expected else if clause or optional else clause or one of [';', '}', case, default, else]",
		yyXError{428, -1}: "expected else if list clause or optional else clause or one of [';', '}', case, default, else]",
		yyXError{461, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, range, struct]",
		yyXError{473, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, range, struct]",
		yyXError{115, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{117, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{462, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{463, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{464, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{465, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{466, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{467, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{468, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{469, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{470, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{471, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{472, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{496, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{498, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{513, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{514, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
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
		yyXError{454, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{475, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{489, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{293, -1}: "expected expression or optional expression or one of ['!', '&', '(', '*', '+', '-', ':', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{49, -1}:  "expected expression or type literal or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{329, -1}: "expected expression/type literal or one of ['!', '&', '(', ')', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{267, -1}: "expected expression/type literal or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{44, -1}:  "expected for statement header or one of ['!', '&', '(', '*', '+', '-', ';', '[', '^', <-, chan, func, identifier, interface, literal, map, range, struct, {]",
		yyXError{202, -1}: "expected function/method signature or '('",
		yyXError{419, -1}: "expected function/method signature or '('",
		yyXError{445, -1}: "expected function/method signature or '('",
		yyXError{416, -1}: "expected function/method signature or one of ['(', '.', ';', '}']",
		yyXError{102, -1}: "expected function/method signature or one of ['(', «]",
		yyXError{444, -1}: "expected function/method signature or optional generic parameters or one of ['(', «]",
		yyXError{45, -1}:  "expected function/method signature or optional receiver or one of ['(', identifier, «]",
		yyXError{3, -1}:   "expected identifier",
		yyXError{65, -1}:  "expected identifier",
		yyXError{103, -1}: "expected identifier",
		yyXError{175, -1}: "expected identifier",
		yyXError{442, -1}: "expected identifier",
		yyXError{179, -1}: "expected identifier list or identifier",
		yyXError{220, -1}: "expected identifier list or identifier",
		yyXError{10, -1}:  "expected import specification list or one of [')', '.', identifier, literal]",
		yyXError{6, -1}:   "expected import specification or one of ['(', '.', identifier, literal]",
		yyXError{27, -1}:  "expected import specification or one of [')', '.', identifier, literal]",
		yyXError{414, -1}: "expected interface method declaration list or identifier",
		yyXError{412, -1}: "expected interface method declaration list or one of ['}', identifier]",
		yyXError{422, -1}: "expected interface method declaration or one of ['}', identifier]",
		yyXError{51, -1}:  "expected left brace or one of ['(', ')', ',', ':', '=', '{', ..., :=, {]",
		yyXError{47, -1}:  "expected left brace or one of ['{', {]",
		yyXError{71, -1}:  "expected left brace or one of ['{', {]",
		yyXError{400, -1}: "expected left brace or one of ['{', {]",
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
		yyXError{403, -1}: "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '[', ']', '^', '{', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{406, -1}: "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '[', ']', '^', '{', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{407, -1}: "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '[', ']', '^', '{', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
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
		yyXError{404, -1}: "expected one of [!=, &&, &^, '%', '&', ')', '*', '+', '-', '/', '<', '>', '^', '|', <-, <<, <=, ==, >=, >>, ||]",
		yyXError{313, -1}: "expected one of [!=, &&, &^, '%', '&', '*', '+', ',', '-', '/', ':', '<', '>', '^', '|', '}', <-, <<, <=, ==, >=, >>, ||]",
		yyXError{383, -1}: "expected one of [!=, &&, &^, '%', '&', '*', '+', ',', '-', '/', ':', '<', '>', '^', '|', '}', <-, <<, <=, ==, >=, >>, ||]",
		yyXError{323, -1}: "expected one of [!=, &&, &^, '%', '&', '*', '+', ',', '-', '/', '<', '>', '^', '|', '}', <-, <<, <=, ==, >=, >>, ||]",
		yyXError{326, -1}: "expected one of [!=, &&, &^, '%', '&', '*', '+', ',', '-', '/', '<', '>', '^', '|', '}', <-, <<, <=, ==, >=, >>, ||]",
		yyXError{393, -1}: "expected one of [!=, &&, &^, '%', '&', '*', '+', ',', '-', '/', '<', '>', '^', '|', '}', <-, <<, <=, ==, >=, >>, ||]",
		yyXError{396, -1}: "expected one of [!=, &&, &^, '%', '&', '*', '+', ',', '-', '/', '<', '>', '^', '|', '}', <-, <<, <=, ==, >=, >>, ||]",
		yyXError{296, -1}: "expected one of [!=, &&, &^, '%', '&', '*', '+', '-', '/', ':', '<', '>', ']', '^', '|', <-, <<, <=, ==, >=, >>, ||]",
		yyXError{299, -1}: "expected one of [!=, &&, &^, '%', '&', '*', '+', '-', '/', ':', '<', '>', ']', '^', '|', <-, <<, <=, ==, >=, >>, ||]",
		yyXError{271, -1}: "expected one of [!=, &&, &^, '%', '&', '*', '+', '-', '/', ':', '<', '>', '^', '|', <-, <<, <=, ==, >=, >>, ||]",
		yyXError{273, -1}: "expected one of [!=, &&, &^, '%', '&', '*', '+', '-', '/', ':', '<', '>', '^', '|', <-, <<, <=, ==, >=, >>, ||]",
		yyXError{371, -1}: "expected one of [!=, &&, &^, '%', '&', '*', '+', '-', '/', ';', '<', '>', '^', '|', '}', <-, <<, <=, ==, >=, >>, case, default, ||]",
		yyXError{372, -1}: "expected one of [!=, &&, &^, '%', '&', '*', '+', '-', '/', ';', '<', '>', '^', '|', '}', <-, <<, <=, ==, >=, >>, case, default, ||]",
		yyXError{516, -1}: "expected one of [!=, &&, &^, '%', '&', '*', '+', '-', '/', '<', '>', ']', '^', '|', <-, <<, <=, ==, >=, >>, ||]",
		yyXError{279, -1}: "expected one of [!=, &&, &^, '%', '&', '*', '+', '-', '/', '<', '>', '^', '|', <-, <<, <=, ==, >=, >>, {, ||]",
		yyXError{455, -1}: "expected one of [!=, &&, &^, '%', '&', '*', '+', '-', '/', '<', '>', '^', '|', <-, <<, <=, ==, >=, >>, {, ||]",
		yyXError{476, -1}: "expected one of [!=, &&, &^, '%', '&', '*', '+', '-', '/', '<', '>', '^', '|', <-, <<, <=, ==, >=, >>, {, ||]",
		yyXError{490, -1}: "expected one of [!=, &&, &^, '%', '&', '*', '+', '-', '/', '<', '>', '^', '|', <-, <<, <=, ==, >=, >>, {, ||]",
		yyXError{34, -1}:  "expected one of [%=, &=, &^=, ',', '=', *=, +=, -=, /=, :=, <<=, >>=, ^=, |=]",
		yyXError{450, -1}: "expected one of [%=, &=, &^=, ',', '=', *=, +=, -=, /=, :=, <<=, >>=, ^=, |=]",
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
		yyXError{411, -1}: "expected one of ['(', ')', ',', ':', ';', '=', ']', '{', '}', ..., :=, case, default, literal, {, »]",
		yyXError{413, -1}: "expected one of ['(', ')', ',', ':', ';', '=', ']', '{', '}', ..., :=, case, default, literal, {, »]",
		yyXError{424, -1}: "expected one of ['(', ')', ',', ':', ';', '=', ']', '{', '}', ..., :=, case, default, literal, {, »]",
		yyXError{509, -1}: "expected one of ['(', ')', ',', ':', ';', '=', ']', '{', '}', ..., :=, case, default, literal, {, »]",
		yyXError{510, -1}: "expected one of ['(', ')', ',', ':', ';', '=', ']', '{', '}', ..., :=, case, default, literal, {, »]",
		yyXError{512, -1}: "expected one of ['(', ')', ',', ':', ';', '=', ']', '{', '}', ..., :=, case, default, literal, {, »]",
		yyXError{518, -1}: "expected one of ['(', ')', ',', ':', ';', '=', ']', '{', '}', ..., :=, case, default, literal, {, »]",
		yyXError{520, -1}: "expected one of ['(', ')', ',', ':', ';', '=', ']', '{', '}', ..., :=, case, default, literal, {, »]",
		yyXError{522, -1}: "expected one of ['(', ')', ',', ':', ';', '=', ']', '{', '}', ..., :=, case, default, literal, {, »]",
		yyXError{37, -1}:  "expected one of ['(', ')', ',', ':', '=', '{', ..., :=, {]",
		yyXError{38, -1}:  "expected one of ['(', ')', ',', ':', '=', '{', ..., :=, {]",
		yyXError{39, -1}:  "expected one of ['(', ')', ',', ':', '=', '{', ..., :=, {]",
		yyXError{40, -1}:  "expected one of ['(', ')', ',', ':', '=', '{', ..., :=, {]",
		yyXError{82, -1}:  "expected one of ['(', ')', ',', ':', '=', ..., :=]",
		yyXError{83, -1}:  "expected one of ['(', ')', ',', ':', '=', ..., :=]",
		yyXError{210, -1}: "expected one of ['(', ')', ',', ':', '=', ..., :=]",
		yyXError{263, -1}: "expected one of ['(', ')', ',', ':', '=', ..., :=]",
		yyXError{405, -1}: "expected one of ['(', ')']",
		yyXError{340, -1}: "expected one of ['(', '*', ',', '.', ';', '[', '}', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{224, -1}: "expected one of ['(', '*', '[', <-, chan, func, identifier, interface, map, struct]",
		yyXError{311, -1}: "expected one of ['(', '.', '[', '{', {]",
		yyXError{292, -1}: "expected one of ['(', identifier]",
		yyXError{377, -1}: "expected one of ['(', identifier]",
		yyXError{264, -1}: "expected one of [')', ',', ':', '=', ..., :=]",
		yyXError{275, -1}: "expected one of [')', ',', ':', '=', ..., :=]",
		yyXError{119, -1}: "expected one of [')', ',', ';', '}', case, default]",
		yyXError{170, -1}: "expected one of [')', ',', ';', '}', case, default]",
		yyXError{499, -1}: "expected one of [')', ',', ';', '}', case, default]",
		yyXError{500, -1}: "expected one of [')', ',', ';', '}', case, default]",
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
		yyXError{503, -1}: "expected one of [')', ';']",
		yyXError{506, -1}: "expected one of [')', ';']",
		yyXError{265, -1}: "expected one of [',', ':', '=', :=]",
		yyXError{312, -1}: "expected one of [',', ':', '}']",
		yyXError{384, -1}: "expected one of [',', ':', '}']",
		yyXError{474, -1}: "expected one of [',', ';', '}', case, default, {]",
		yyXError{477, -1}: "expected one of [',', ';', '}', case, default, {]",
		yyXError{478, -1}: "expected one of [',', ';', '}', case, default, {]",
		yyXError{479, -1}: "expected one of [',', ';', '}', case, default, {]",
		yyXError{480, -1}: "expected one of [',', ';', '}', case, default, {]",
		yyXError{481, -1}: "expected one of [',', ';', '}', case, default, {]",
		yyXError{482, -1}: "expected one of [',', ';', '}', case, default, {]",
		yyXError{483, -1}: "expected one of [',', ';', '}', case, default, {]",
		yyXError{484, -1}: "expected one of [',', ';', '}', case, default, {]",
		yyXError{485, -1}: "expected one of [',', ';', '}', case, default, {]",
		yyXError{486, -1}: "expected one of [',', ';', '}', case, default, {]",
		yyXError{487, -1}: "expected one of [',', ';', '}', case, default, {]",
		yyXError{488, -1}: "expected one of [',', ';', '}', case, default, {]",
		yyXError{368, -1}: "expected one of [',', ';', '}', case, default]",
		yyXError{314, -1}: "expected one of [',', '}']",
		yyXError{320, -1}: "expected one of [',', '}']",
		yyXError{322, -1}: "expected one of [',', '}']",
		yyXError{325, -1}: "expected one of [',', '}']",
		yyXError{385, -1}: "expected one of [',', '}']",
		yyXError{391, -1}: "expected one of [',', '}']",
		yyXError{394, -1}: "expected one of [',', '}']",
		yyXError{397, -1}: "expected one of [',', '}']",
		yyXError{123, -1}: "expected one of [',', »]",
		yyXError{124, -1}: "expected one of [',', »]",
		yyXError{127, -1}: "expected one of [',', »]",
		yyXError{201, -1}: "expected one of [',', »]",
		yyXError{223, -1}: "expected one of [',', »]",
		yyXError{300, -1}: "expected one of [':', ']']",
		yyXError{431, -1}: "expected one of [';', '}', case, default, else]",
		yyXError{436, -1}: "expected one of [';', '}', case, default, else]",
		yyXError{439, -1}: "expected one of [';', '}', case, default, else]",
		yyXError{9, -1}:   "expected one of [';', '}', case, default, literal]",
		yyXError{58, -1}:  "expected one of [';', '}', case, default, {]",
		yyXError{491, -1}: "expected one of [';', '}', case, default, {]",
		yyXError{492, -1}: "expected one of [';', '}', case, default, {]",
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
		yyXError{399, -1}: "expected one of [';', '}', case, default]",
		yyXError{432, -1}: "expected one of [';', '}', case, default]",
		yyXError{434, -1}: "expected one of [';', '}', case, default]",
		yyXError{456, -1}: "expected one of [';', '}', case, default]",
		yyXError{494, -1}: "expected one of [';', '}', case, default]",
		yyXError{501, -1}: "expected one of [';', '}', case, default]",
		yyXError{507, -1}: "expected one of [';', '}', case, default]",
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
		yyXError{402, -1}: "expected one of [';', '}']",
		yyXError{417, -1}: "expected one of [';', '}']",
		yyXError{418, -1}: "expected one of [';', '}']",
		yyXError{420, -1}: "expected one of [';', '}']",
		yyXError{423, -1}: "expected one of [';', '}']",
		yyXError{438, -1}: "expected one of [';', '}']",
		yyXError{233, -1}: "expected one of [';', {]",
		yyXError{234, -1}: "expected one of [';', {]",
		yyXError{425, -1}: "expected one of [';', {]",
		yyXError{452, -1}: "expected one of [';', {]",
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
		yyXError{446, -1}: "expected optional function body or one of [';', '{']",
		yyXError{447, -1}: "expected optional function body or one of [';', '{']",
		yyXError{120, -1}: "expected optional generic arguments or one of [!=, %=, &&, &=, &^, &^=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '[', ']', '^', '{', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||, «]",
		yyXError{52, -1}:  "expected optional generic arguments or one of [!=, %=, &&, &=, &^, &^=, '%', '&', '(', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '[', '^', '{', '|', '}', *=, ++, +=, --, -=, /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, |=, ||, «]",
		yyXError{280, -1}: "expected optional generic arguments or one of [!=, &&, &^, '%', '&', '(', '*', '+', '-', '.', '/', '<', '>', '[', '^', '{', '|', :=, <-, <<, <=, ==, >=, >>, {, ||, «]",
		yyXError{112, -1}: "expected optional generic arguments or one of ['(', ')', ',', ':', ';', '=', ']', '{', '}', ..., :=, case, default, literal, {, «, »]",
		yyXError{219, -1}: "expected optional generic parameters or type or one of ['(', '*', '[', <-, chan, func, identifier, interface, map, struct, «]",
		yyXError{59, -1}:  "expected optional identifier or one of [';', '}', case, default, identifier]",
		yyXError{60, -1}:  "expected optional identifier or one of [';', '}', case, default, identifier]",
		yyXError{181, -1}: "expected optional result or one of ['(', ')', '*', ',', ':', ';', '=', '[', ']', '{', '}', ..., :=, <-, case, chan, default, func, identifier, interface, literal, map, struct, {, »]",
		yyXError{443, -1}: "expected optional result or one of ['(', '*', '[', '{', <-, chan, func, identifier, interface, map, struct, {]",
		yyXError{24, -1}:  "expected optional semicolon or one of [')', ';']",
		yyXError{96, -1}:  "expected optional semicolon or one of [')', ';']",
		yyXError{226, -1}: "expected optional semicolon or one of [')', ';']",
		yyXError{502, -1}: "expected optional semicolon or one of [')', ';']",
		yyXError{346, -1}: "expected optional semicolon or one of [';', '}']",
		yyXError{415, -1}: "expected optional semicolon or one of [';', '}']",
		yyXError{457, -1}: "expected optional simple statement or one of ['!', '&', '(', '*', '+', '-', ';', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{440, -1}: "expected optional simple statement or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct, {]",
		yyXError{459, -1}: "expected optional simple statement or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct, {]",
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
		yyXError{495, -1}: "expected type or one of ['(', ')', '*', ',', ';', '=', '[', '}', <-, case, chan, default, func, identifier, interface, map, struct]",
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
		yyXError{408, -1}: "expected type or one of ['(', '*', '[', <-, chan, func, identifier, interface, map, struct]",
		yyXError{410, -1}: "expected type or one of ['(', '*', '[', <-, chan, func, identifier, interface, map, struct]",
		yyXError{508, -1}: "expected type or one of ['(', '*', '[', <-, chan, func, identifier, interface, map, struct]",
		yyXError{511, -1}: "expected type or one of ['(', '*', '[', <-, chan, func, identifier, interface, map, struct]",
		yyXError{517, -1}: "expected type or one of ['(', '*', '[', <-, chan, func, identifier, interface, map, struct]",
		yyXError{519, -1}: "expected type or one of ['(', '*', '[', <-, chan, func, identifier, interface, map, struct]",
		yyXError{521, -1}: "expected type or one of ['(', '*', '[', <-, chan, func, identifier, interface, map, struct]",
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
		yyXError{441, -1}: "expected {",
		yyXError{451, -1}: "expected {",
		yyXError{460, -1}: "expected {",
	}

	yyParseTab = [523][]uint16{
		// 0
		{162: 298, 171: 300, 301, 299},
		{99: 297},
		{2: 44, 44, 44, 44, 44, 44, 44, 11: 44, 13: 44, 16: 44, 44, 44, 44, 44, 44, 44, 32: 44, 44, 44, 44, 38: 44, 79: 44, 81: 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 99: 44, 180: 329},
		{13: 327},
		{2: 176, 176, 176, 176, 176, 176, 176, 11: 176, 13: 176, 16: 176, 176, 176, 176, 176, 176, 176, 32: 176, 176, 176, 176, 38: 176, 79: 176, 81: 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 176, 99: 176, 123: 176, 167: 302},
		// 5
		{2: 127, 127, 127, 127, 127, 127, 127, 11: 127, 13: 127, 16: 127, 127, 127, 127, 127, 127, 127, 32: 127, 127, 127, 127, 38: 127, 79: 127, 81: 127, 127, 127, 127, 127, 127, 127, 127, 127, 127, 127, 127, 127, 127, 99: 127, 123: 303, 166: 304},
		{2: 307, 13: 306, 16: 192, 32: 192, 192, 192, 192, 75: 309, 118: 310, 128: 308},
		{305},
		{2: 175, 175, 175, 175, 175, 175, 175, 11: 175, 13: 175, 16: 175, 175, 175, 175, 175, 175, 175, 32: 175, 175, 175, 175, 38: 175, 79: 175, 81: 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 99: 175, 123: 175},
		{191, 191, 9: 191, 191, 16: 191, 32: 191, 191, 191, 191},
		// 10
		{12: 320, 306, 16: 192, 32: 192, 192, 192, 192, 75: 309, 118: 310, 128: 322, 168: 321},
		{183},
		{16: 315, 32: 311, 312, 313, 314, 41: 318},
		{16: 315, 32: 311, 312, 313, 314, 41: 316},
		{277, 277, 277, 277, 277, 277, 277, 277, 277, 277, 277, 277, 277, 14: 277, 277, 23: 277, 31: 277, 36: 277, 277, 39: 277, 277, 47: 277, 277, 277, 277, 277, 277, 277, 277, 277, 277, 277, 277, 277, 277, 62: 277, 277, 277, 277, 277, 277, 277, 277, 277, 277, 277, 277, 277, 277, 79: 277},
		// 15
		{276, 276, 276, 276, 276, 276, 276, 276, 276, 276, 276, 276, 276, 14: 276, 276, 23: 276, 31: 276, 36: 276, 276, 39: 276, 276, 47: 276, 276, 276, 276, 276, 276, 276, 276, 276, 276, 276, 276, 276, 276, 62: 276, 276, 276, 276, 276, 276, 276, 276, 276, 276, 276, 276, 276, 276, 79: 276},
		{275, 275, 275, 275, 275, 275, 275, 275, 275, 275, 275, 275, 275, 14: 275, 275, 23: 275, 31: 275, 36: 275, 275, 39: 275, 275, 47: 275, 275, 275, 275, 275, 275, 275, 275, 275, 275, 275, 275, 275, 275, 62: 275, 275, 275, 275, 275, 275, 275, 275, 275, 275, 275, 275, 275, 275, 79: 275},
		{274, 274, 274, 274, 274, 274, 274, 274, 274, 274, 274, 274, 274, 14: 274, 274, 23: 274, 31: 274, 36: 274, 274, 39: 274, 274, 47: 274, 274, 274, 274, 274, 274, 274, 274, 274, 274, 274, 274, 274, 274, 62: 274, 274, 274, 274, 274, 274, 274, 274, 274, 274, 274, 274, 274, 274, 79: 274},
		{273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 14: 273, 273, 23: 273, 31: 273, 36: 273, 273, 39: 273, 273, 47: 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 62: 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 273, 79: 273},
		{181, 12: 181, 79: 317},
		// 20
		{179, 12: 179},
		{182, 12: 182, 79: 319},
		{180, 12: 180},
		{185},
		{324, 12: 114, 114: 323},
		// 25
		{178, 12: 178},
		{12: 326},
		{12: 113, 306, 16: 192, 32: 192, 192, 192, 192, 75: 309, 118: 310, 128: 325},
		{177, 12: 177},
		{184},
		// 30
		{328},
		{2: 147, 147, 147, 147, 147, 147, 147, 11: 147, 13: 147, 16: 147, 147, 147, 147, 147, 147, 147, 32: 147, 147, 147, 147, 38: 147, 79: 147, 81: 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 147, 99: 147, 123: 147},
		{2: 346, 378, 382, 383, 384, 385, 386, 11: 330, 13: 349, 16: 315, 332, 342, 344, 345, 333, 368, 24: 334, 379, 348, 380, 335, 336, 337, 32: 311, 312, 313, 314, 38: 381, 41: 347, 353, 351, 350, 352, 339, 61: 340, 77: 331, 79: 375, 81: 377, 343, 356, 338, 357, 358, 359, 341, 361, 362, 364, 354, 369, 387, 355, 366, 99: 296, 102: 370, 360, 363, 106: 365, 108: 374, 367, 372, 373, 165: 371, 179: 376},
		{2: 346, 378, 382, 383, 384, 385, 386, 11: 330, 13: 417, 16: 315, 332, 399, 344, 345, 333, 368, 24: 334, 379, 348, 380, 335, 336, 337, 32: 311, 312, 313, 314, 37: 812, 381, 40: 814, 347, 353, 351, 350, 352, 339, 61: 813},
		{14: 425, 31: 810, 36: 811, 62: 759, 761, 760, 762, 763, 764, 765, 766, 767, 768, 769},
		// 35
		{2: 402, 403, 11: 330, 13: 401, 17: 332, 399, 344, 345, 333, 368, 24: 404, 405, 406, 407, 408, 410, 411, 78: 409, 80: 807, 181: 808},
		{17: 805},
		{2: 26, 12: 26, 14: 26, 252, 23: 26, 31: 26, 36: 26, 26, 39: 252},
		{2: 22, 12: 22, 14: 22, 251, 23: 22, 31: 22, 36: 22, 22, 39: 251},
		{2: 21, 12: 21, 14: 21, 250, 23: 21, 31: 21, 36: 21, 21, 39: 250},
		// 40
		{2: 20, 12: 20, 14: 20, 249, 23: 20, 31: 20, 36: 20, 20, 39: 249},
		{2: 790, 13: 388, 98: 792, 126: 791},
		{233, 233, 3: 233, 233, 233, 233, 233, 233, 233, 233, 12: 233, 14: 233, 233, 23: 233, 31: 233, 36: 233, 233, 40: 233, 47: 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 62: 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233, 233},
		{110, 110, 3: 429, 428, 430, 431, 435, 446, 110, 110, 14: 210, 110, 31: 210, 36: 210, 47: 427, 432, 433, 434, 436, 437, 438, 439, 440, 441, 442, 443, 444, 445, 62: 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 788, 789},
		{106, 2: 346, 378, 382, 383, 384, 385, 386, 11: 330, 13: 417, 15: 106, 315, 332, 399, 344, 345, 333, 368, 24: 334, 379, 348, 380, 335, 336, 337, 32: 311, 312, 313, 314, 38: 381, 41: 347, 353, 351, 350, 352, 339, 61: 340, 77: 747, 95: 355, 530, 107: 749, 130: 751, 163: 750, 174: 748},
		// 45
		{2: 477, 13: 121, 97: 476, 105: 740, 120: 475, 175: 739},
		{106, 2: 346, 378, 382, 383, 384, 385, 386, 11: 330, 13: 417, 15: 106, 315, 332, 399, 344, 345, 333, 368, 24: 334, 379, 348, 380, 335, 336, 337, 32: 311, 312, 313, 314, 38: 381, 41: 347, 353, 351, 350, 352, 339, 61: 340, 77: 331, 95: 355, 530, 107: 722, 137: 723},
		{15: 632, 39: 633, 101: 709},
		{11: 705},
		{2: 346, 378, 382, 383, 384, 385, 386, 11: 330, 13: 417, 16: 315, 332, 399, 344, 345, 333, 368, 24: 334, 379, 348, 380, 335, 336, 337, 32: 311, 312, 313, 314, 38: 381, 41: 347, 702, 351, 350, 352, 339, 61: 701},
		// 50
		{151, 151, 151, 151, 151, 151, 151, 151, 151, 151, 151, 151, 151, 14: 151, 151, 23: 151, 31: 151, 36: 151, 151, 39: 151, 151, 47: 151, 151, 151, 151, 151, 151, 151, 151, 151, 151, 151, 151, 151, 151, 62: 151, 151, 151, 151, 151, 151, 151, 151, 151, 151, 151, 151, 151, 151},
		{2: 24, 12: 24, 14: 24, 150, 23: 24, 31: 24, 36: 24, 24, 39: 150, 152: 697},
		{196, 196, 196, 196, 196, 196, 196, 196, 196, 196, 196, 196, 14: 196, 23: 695, 31: 196, 36: 196, 39: 196, 47: 196, 196, 196, 196, 196, 196, 196, 196, 196, 196, 196, 196, 196, 196, 62: 196, 196, 196, 196, 196, 196, 196, 196, 196, 196, 196, 196, 196, 196, 97: 418, 122: 419},
		{138, 138, 138, 138, 138, 138, 138, 138, 138, 138, 138, 138, 138, 14: 138, 138, 23: 138, 31: 138, 36: 138, 138, 39: 138, 138, 47: 138, 138, 138, 138, 138, 138, 138, 138, 138, 138, 138, 138, 138, 138, 62: 138, 138, 138, 138, 138, 138, 138, 138, 138, 138, 138, 138, 138, 138},
		{15: 632, 39: 633, 101: 678, 119: 679},
		// 55
		{9, 9, 587, 9, 9, 9, 9, 9, 9, 9, 9, 590, 9, 14: 9, 9, 23: 9, 31: 9, 36: 9, 9, 39: 588, 9, 47: 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 62: 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 674, 113: 592, 134: 591},
		{2: 509},
		{15: 116, 153: 672},
		{111, 111, 9: 111, 111, 15: 111},
		{192, 192, 9: 192, 192, 13: 306, 118: 671},
		// 60
		{192, 192, 9: 192, 192, 13: 306, 118: 670},
		{2: 346, 378, 382, 383, 384, 385, 386, 11: 330, 13: 417, 16: 315, 332, 399, 344, 345, 333, 368, 24: 334, 379, 348, 380, 335, 336, 337, 32: 311, 312, 313, 314, 38: 381, 41: 347, 353, 351, 350, 352, 339, 61: 669},
		{91, 91, 9: 91, 91},
		{90, 90, 9: 90, 90},
		{2: 346, 378, 382, 383, 384, 385, 386, 11: 330, 13: 417, 16: 315, 332, 399, 344, 345, 333, 368, 24: 334, 379, 348, 380, 335, 336, 337, 32: 311, 312, 313, 314, 38: 381, 41: 347, 353, 351, 350, 352, 339, 61: 668},
		// 65
		{13: 667},
		{86, 86, 9: 86, 86},
		{208, 208, 346, 378, 382, 383, 384, 385, 386, 208, 208, 330, 13: 417, 16: 315, 332, 399, 344, 345, 333, 368, 24: 334, 379, 348, 380, 335, 336, 337, 32: 311, 312, 313, 314, 38: 381, 41: 347, 353, 351, 350, 352, 339, 61: 415, 77: 665, 161: 666},
		{84, 84, 9: 84, 84},
		{83, 83, 9: 83, 83},
		// 70
		{82, 82, 9: 82, 82},
		{15: 632, 39: 633, 101: 634},
		{106, 2: 346, 378, 382, 383, 384, 385, 386, 11: 330, 13: 417, 15: 106, 315, 332, 399, 344, 345, 333, 368, 24: 334, 379, 348, 380, 335, 336, 337, 32: 311, 312, 313, 314, 38: 381, 41: 347, 353, 351, 350, 352, 339, 61: 340, 77: 331, 95: 355, 530, 107: 531, 178: 532},
		{50},
		{49},
		// 75
		{48},
		{47},
		{46},
		{45},
		{529},
		// 80
		{2: 514, 13: 516, 131: 515},
		{2: 346, 378, 382, 383, 384, 385, 386, 11: 330, 13: 417, 16: 315, 332, 399, 344, 345, 333, 368, 24: 334, 379, 348, 380, 335, 336, 337, 32: 311, 312, 313, 314, 38: 381, 41: 347, 507, 351, 350, 352, 508},
		{2: 25, 12: 25, 14: 25, 23: 25, 31: 25, 36: 25, 25},
		{2: 23, 12: 23, 14: 23, 23: 23, 31: 23, 36: 23, 23},
		{2: 346, 378, 382, 383, 384, 385, 386, 11: 330, 13: 417, 16: 315, 332, 399, 344, 345, 333, 368, 24: 334, 379, 348, 380, 335, 336, 337, 32: 311, 312, 313, 314, 38: 381, 41: 347, 353, 351, 350, 352, 506},
		// 85
		{2: 346, 378, 382, 383, 384, 385, 386, 11: 330, 13: 417, 16: 315, 332, 399, 344, 345, 333, 368, 24: 334, 379, 348, 380, 335, 336, 337, 32: 311, 312, 313, 314, 38: 381, 41: 347, 353, 351, 350, 352, 505},
		{2: 346, 378, 382, 383, 384, 385, 386, 11: 330, 13: 417, 16: 315, 332, 399, 344, 345, 333, 368, 24: 334, 379, 348, 380, 335, 336, 337, 32: 311, 312, 313, 314, 38: 381, 41: 347, 353, 351, 350, 352, 504},
		{2: 346, 378, 382, 383, 384, 385, 386, 11: 330, 13: 417, 16: 315, 332, 399, 344, 345, 333, 368, 24: 334, 379, 348, 380, 335, 336, 337, 32: 311, 312, 313, 314, 38: 381, 41: 347, 353, 351, 350, 352, 503},
		{2: 346, 378, 382, 383, 384, 385, 386, 11: 330, 13: 417, 16: 315, 332, 399, 344, 345, 333, 368, 24: 334, 379, 348, 380, 335, 336, 337, 32: 311, 312, 313, 314, 38: 381, 41: 347, 353, 351, 350, 352, 502},
		{2: 346, 378, 382, 383, 384, 385, 386, 11: 330, 13: 417, 16: 315, 332, 399, 344, 345, 333, 368, 24: 334, 379, 348, 380, 335, 336, 337, 32: 311, 312, 313, 314, 38: 381, 41: 347, 353, 351, 350, 352, 501},
		// 90
		{2: 390, 13: 388, 98: 389, 132: 391},
		{190, 190, 190, 190, 9: 190, 190, 190, 190, 190, 190, 17: 190, 190, 190, 190, 190, 190, 31: 190, 76: 190},
		{2: 402, 403, 11: 330, 13: 401, 400, 17: 332, 399, 344, 345, 333, 368, 24: 404, 405, 406, 407, 408, 410, 411, 412, 78: 409, 80: 413},
		{12: 392, 388, 98: 389, 132: 394, 184: 393},
		{6, 6, 9: 6, 6},
		// 95
		{8, 8, 9: 8, 8},
		{395, 12: 114, 114: 396},
		{2, 12: 2},
		{12: 113, 388, 98: 389, 132: 398},
		{12: 397},
		// 100
		{7, 7, 9: 7, 7},
		{1, 12: 1},
		{2: 477, 97: 476, 105: 478, 120: 475},
		{13: 474},
		{126, 126, 126, 9: 126, 126, 12: 126, 14: 126, 126, 126, 23: 126, 31: 126, 36: 126, 126, 39: 126, 126, 75: 472, 126, 97: 126},
		// 105
		{2: 402, 403, 11: 330, 13: 401, 17: 332, 399, 344, 345, 333, 368, 24: 404, 405, 406, 407, 408, 410, 411, 78: 409, 80: 470},
		{2: 402, 403, 11: 330, 13: 401, 17: 332, 399, 344, 345, 333, 368, 24: 404, 405, 406, 407, 408, 410, 411, 78: 409, 80: 469},
		{40, 40, 40, 9: 40, 40, 12: 40, 14: 40, 40, 40, 23: 40, 31: 40, 36: 40, 40, 39: 40, 40, 76: 40},
		{39, 39, 39, 9: 39, 39, 12: 39, 14: 39, 39, 39, 23: 39, 31: 39, 36: 39, 39, 39: 39, 39, 76: 39},
		{38, 38, 38, 9: 38, 38, 12: 38, 14: 38, 38, 38, 23: 38, 31: 38, 36: 38, 38, 39: 38, 38, 76: 38},
		// 110
		{37, 37, 37, 9: 37, 37, 12: 37, 14: 37, 37, 37, 23: 37, 31: 37, 36: 37, 37, 39: 37, 37, 76: 37},
		{36, 36, 36, 9: 36, 36, 12: 36, 14: 36, 36, 36, 23: 36, 31: 36, 36: 36, 36, 39: 36, 36, 76: 36},
		{196, 196, 196, 9: 196, 196, 12: 196, 14: 196, 196, 196, 23: 196, 31: 196, 36: 196, 196, 39: 196, 196, 76: 196, 97: 418, 122: 468},
		{34, 34, 34, 9: 34, 34, 12: 34, 14: 34, 34, 34, 23: 34, 31: 34, 36: 34, 34, 39: 34, 34, 76: 34},
		{33, 33, 33, 9: 33, 33, 12: 33, 14: 33, 33, 33, 23: 33, 31: 33, 36: 33, 33, 39: 33, 33, 76: 33},
		// 115
		{2: 346, 378, 382, 383, 384, 385, 386, 11: 330, 13: 417, 16: 315, 332, 399, 344, 345, 333, 368, 24: 334, 379, 348, 380, 335, 336, 337, 32: 311, 312, 313, 314, 38: 381, 41: 347, 353, 351, 350, 352, 339, 61: 415, 77: 467},
		{4, 4, 9: 4, 4, 12: 4, 31: 414},
		{2: 346, 378, 382, 383, 384, 385, 386, 11: 330, 13: 417, 16: 315, 332, 399, 344, 345, 333, 368, 24: 334, 379, 348, 380, 335, 336, 337, 32: 311, 312, 313, 314, 38: 381, 41: 347, 353, 351, 350, 352, 339, 61: 415, 77: 416},
		{210, 210, 3: 429, 428, 430, 431, 435, 446, 210, 210, 12: 210, 14: 210, 210, 47: 427, 432, 433, 434, 436, 437, 438, 439, 440, 441, 442, 443, 444, 445},
		{3, 3, 9: 3, 3, 12: 3, 14: 425},
		// 120
		{196, 196, 196, 196, 196, 196, 196, 196, 196, 196, 196, 196, 196, 14: 196, 196, 23: 196, 31: 196, 36: 196, 196, 39: 196, 196, 47: 196, 196, 196, 196, 196, 196, 196, 196, 196, 196, 196, 196, 196, 196, 62: 196, 196, 196, 196, 196, 196, 196, 196, 196, 196, 196, 196, 196, 196, 97: 418, 122: 419},
		{2: 402, 403, 11: 330, 13: 401, 17: 332, 399, 344, 345, 333, 368, 24: 404, 405, 406, 407, 408, 410, 411, 78: 409, 80: 421, 182: 420},
		{148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 14: 148, 148, 23: 148, 31: 148, 36: 148, 148, 39: 148, 148, 47: 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 62: 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148},
		{14: 423, 76: 422},
		{14: 29, 76: 29},
		// 125
		{195, 195, 195, 195, 195, 195, 195, 195, 195, 195, 195, 195, 195, 14: 195, 195, 195, 23: 195, 31: 195, 36: 195, 195, 39: 195, 195, 47: 195, 195, 195, 195, 195, 195, 195, 195, 195, 195, 195, 195, 195, 195, 62: 195, 195, 195, 195, 195, 195, 195, 195, 195, 195, 195, 195, 195, 195, 195},
		{2: 402, 403, 11: 330, 13: 401, 17: 332, 399, 344, 345, 333, 368, 24: 404, 405, 406, 407, 408, 410, 411, 78: 409, 80: 424},
		{14: 28, 76: 28},
		{2: 346, 378, 382, 383, 384, 385, 386, 11: 330, 13: 417, 16: 315, 332, 399, 344, 345, 333, 368, 24: 334, 379, 348, 380, 335, 336, 337, 32: 311, 312, 313, 314, 38: 381, 41: 347, 353, 351, 350, 352, 339, 61: 426},
		{209, 209, 3: 429, 428, 430, 431, 435, 446, 209, 209, 12: 209, 14: 209, 209, 31: 209, 36: 209, 47: 427, 432, 433, 434, 436, 437, 438, 439, 440, 441, 442, 443, 444, 445, 62: 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209},
		// 130
		{2: 346, 378, 382, 383, 384, 385, 386, 11: 330, 13: 417, 16: 315, 332, 399, 344, 345, 333, 368, 24: 334, 379, 348, 380, 335, 336, 337, 32: 311, 312, 313, 314, 38: 381, 41: 347, 353, 351, 350, 352, 339, 61: 466},
		{2: 346, 378, 382, 383, 384, 385, 386, 11: 330, 13: 417, 16: 315, 332, 399, 344, 345, 333, 368, 24: 334, 379, 348, 380, 335, 336, 337, 32: 311, 312, 313, 314, 38: 381, 41: 347, 353, 351, 350, 352, 339, 61: 465},
		{2: 346, 378, 382, 383, 384, 385, 386, 11: 330, 13: 417, 16: 315, 332, 399, 344, 345, 333, 368, 24: 334, 379, 348, 380, 335, 336, 337, 32: 311, 312, 313, 314, 38: 381, 41: 347, 353, 351, 350, 352, 339, 61: 464},
		{2: 346, 378, 382, 383, 384, 385, 386, 11: 330, 13: 417, 16: 315, 332, 399, 344, 345, 333, 368, 24: 334, 379, 348, 380, 335, 336, 337, 32: 311, 312, 313, 314, 38: 381, 41: 347, 353, 351, 350, 352, 339, 61: 463},
		{2: 346, 378, 382, 383, 384, 385, 386, 11: 330, 13: 417, 16: 315, 332, 399, 344, 345, 333, 368, 24: 334, 379, 348, 380, 335, 336, 337, 32: 311, 312, 313, 314, 38: 381, 41: 347, 353, 351, 350, 352, 339, 61: 462},
		// 135
		{2: 346, 378, 382, 383, 384, 385, 386, 11: 330, 13: 417, 16: 315, 332, 399, 344, 345, 333, 368, 24: 334, 379, 348, 380, 335, 336, 337, 32: 311, 312, 313, 314, 38: 381, 41: 347, 353, 351, 350, 352, 339, 61: 461},
		{2: 346, 378, 382, 383, 384, 385, 386, 11: 330, 13: 417, 16: 315, 332, 399, 344, 345, 333, 368, 24: 334, 379, 348, 380, 335, 336, 337, 32: 311, 312, 313, 314, 38: 381, 41: 347, 353, 351, 350, 352, 339, 61: 460},
		{2: 346, 378, 382, 383, 384, 385, 386, 11: 330, 13: 417, 16: 315, 332, 399, 344, 345, 333, 368, 24: 334, 379, 348, 380, 335, 336, 337, 32: 311, 312, 313, 314, 38: 381, 41: 347, 353, 351, 350, 352, 339, 61: 459},
		{2: 346, 378, 382, 383, 384, 385, 386, 11: 330, 13: 417, 16: 315, 332, 399, 344, 345, 333, 368, 24: 334, 379, 348, 380, 335, 336, 337, 32: 311, 312, 313, 314, 38: 381, 41: 347, 353, 351, 350, 352, 339, 61: 458},
		{2: 346, 378, 382, 383, 384, 385, 386, 11: 330, 13: 417, 16: 315, 332, 399, 344, 345, 333, 368, 24: 334, 379, 348, 380, 335, 336, 337, 32: 311, 312, 313, 314, 38: 381, 41: 347, 353, 351, 350, 352, 339, 61: 457},
		// 140
		{2: 346, 378, 382, 383, 384, 385, 386, 11: 330, 13: 417, 16: 315, 332, 399, 344, 345, 333, 368, 24: 334, 379, 348, 380, 335, 336, 337, 32: 311, 312, 313, 314, 38: 381, 41: 347, 353, 351, 350, 352, 339, 61: 456},
		{2: 346, 378, 382, 383, 384, 385, 386, 11: 330, 13: 417, 16: 315, 332, 399, 344, 345, 333, 368, 24: 334, 379, 348, 380, 335, 336, 337, 32: 311, 312, 313, 314, 38: 381, 41: 347, 353, 351, 350, 352, 339, 61: 455},
		{2: 346, 378, 382, 383, 384, 385, 386, 11: 330, 13: 417, 16: 315, 332, 399, 344, 345, 333, 368, 24: 334, 379, 348, 380, 335, 336, 337, 32: 311, 312, 313, 314, 38: 381, 41: 347, 353, 351, 350, 352, 339, 61: 454},
		{2: 346, 378, 382, 383, 384, 385, 386, 11: 330, 13: 417, 16: 315, 332, 399, 344, 345, 333, 368, 24: 334, 379, 348, 380, 335, 336, 337, 32: 311, 312, 313, 314, 38: 381, 41: 347, 353, 351, 350, 352, 339, 61: 453},
		{2: 346, 378, 382, 383, 384, 385, 386, 11: 330, 13: 417, 16: 315, 332, 399, 344, 345, 333, 368, 24: 334, 379, 348, 380, 335, 336, 337, 32: 311, 312, 313, 314, 38: 381, 41: 347, 353, 351, 350, 352, 339, 61: 452},
		// 145
		{2: 346, 378, 382, 383, 384, 385, 386, 11: 330, 13: 417, 16: 315, 332, 399, 344, 345, 333, 368, 24: 334, 379, 348, 380, 335, 336, 337, 32: 311, 312, 313, 314, 38: 381, 41: 347, 353, 351, 350, 352, 339, 61: 451},
		{2: 346, 378, 382, 383, 384, 385, 386, 11: 330, 13: 417, 16: 315, 332, 399, 344, 345, 333, 368, 24: 334, 379, 348, 380, 335, 336, 337, 32: 311, 312, 313, 314, 38: 381, 41: 347, 353, 351, 350, 352, 339, 61: 450},
		{2: 346, 378, 382, 383, 384, 385, 386, 11: 330, 13: 417, 16: 315, 332, 399, 344, 345, 333, 368, 24: 334, 379, 348, 380, 335, 336, 337, 32: 311, 312, 313, 314, 38: 381, 41: 347, 353, 351, 350, 352, 339, 61: 449},
		{2: 346, 378, 382, 383, 384, 385, 386, 11: 330, 13: 417, 16: 315, 332, 399, 344, 345, 333, 368, 24: 334, 379, 348, 380, 335, 336, 337, 32: 311, 312, 313, 314, 38: 381, 41: 347, 353, 351, 350, 352, 339, 61: 448},
		{2: 346, 378, 382, 383, 384, 385, 386, 11: 330, 13: 417, 16: 315, 332, 399, 344, 345, 333, 368, 24: 334, 379, 348, 380, 335, 336, 337, 32: 311, 312, 313, 314, 38: 381, 41: 347, 353, 351, 350, 352, 339, 61: 447},
		// 150
		{213, 213, 3: 429, 428, 430, 431, 435, 213, 213, 213, 12: 213, 14: 213, 213, 23: 213, 31: 213, 36: 213, 213, 40: 213, 47: 427, 432, 433, 434, 436, 437, 438, 439, 440, 441, 442, 443, 444, 445, 62: 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213},
		{214, 214, 3: 214, 214, 214, 214, 214, 214, 214, 214, 12: 214, 14: 214, 214, 23: 214, 31: 214, 36: 214, 214, 40: 214, 47: 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 62: 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214},
		{215, 215, 3: 429, 428, 430, 431, 435, 215, 215, 215, 12: 215, 14: 215, 215, 23: 215, 31: 215, 36: 215, 215, 40: 215, 47: 427, 432, 433, 434, 436, 437, 438, 439, 440, 441, 442, 443, 215, 445, 62: 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215},
		{216, 216, 3: 429, 428, 430, 431, 435, 216, 216, 216, 12: 216, 14: 216, 216, 23: 216, 31: 216, 36: 216, 216, 40: 216, 47: 427, 432, 216, 216, 436, 216, 438, 216, 216, 216, 442, 216, 216, 445, 62: 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216},
		{217, 217, 3: 217, 217, 217, 217, 217, 217, 217, 217, 12: 217, 14: 217, 217, 23: 217, 31: 217, 36: 217, 217, 40: 217, 47: 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 62: 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 217},
		// 155
		{218, 218, 3: 429, 428, 430, 431, 435, 218, 218, 218, 12: 218, 14: 218, 218, 23: 218, 31: 218, 36: 218, 218, 40: 218, 47: 427, 432, 218, 218, 436, 218, 438, 218, 218, 218, 442, 218, 218, 445, 62: 218, 218, 218, 218, 218, 218, 218, 218, 218, 218, 218, 218, 218},
		{219, 219, 3: 429, 428, 430, 431, 435, 219, 219, 219, 12: 219, 14: 219, 219, 23: 219, 31: 219, 36: 219, 219, 40: 219, 47: 427, 432, 219, 219, 436, 219, 438, 219, 219, 219, 442, 219, 219, 445, 62: 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219},
		{220, 220, 3: 429, 428, 430, 431, 435, 220, 220, 220, 12: 220, 14: 220, 220, 23: 220, 31: 220, 36: 220, 220, 40: 220, 47: 427, 432, 220, 220, 436, 220, 438, 220, 220, 220, 442, 220, 220, 445, 62: 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 220},
		{221, 221, 3: 221, 221, 221, 221, 221, 221, 221, 221, 12: 221, 14: 221, 221, 23: 221, 31: 221, 36: 221, 221, 40: 221, 47: 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 62: 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221},
		{222, 222, 3: 429, 428, 430, 431, 435, 222, 222, 222, 12: 222, 14: 222, 222, 23: 222, 31: 222, 36: 222, 222, 40: 222, 47: 427, 432, 433, 434, 436, 222, 438, 439, 440, 441, 442, 443, 222, 445, 62: 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222},
		// 160
		{223, 223, 3: 429, 428, 223, 223, 223, 223, 223, 223, 12: 223, 14: 223, 223, 23: 223, 31: 223, 36: 223, 223, 40: 223, 47: 427, 432, 223, 223, 223, 223, 438, 223, 223, 223, 442, 223, 223, 445, 62: 223, 223, 223, 223, 223, 223, 223, 223, 223, 223, 223, 223, 223},
		{224, 224, 3: 429, 428, 224, 224, 224, 224, 224, 224, 12: 224, 14: 224, 224, 23: 224, 31: 224, 36: 224, 224, 40: 224, 47: 427, 432, 224, 224, 224, 224, 438, 224, 224, 224, 442, 224, 224, 445, 62: 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224},
		{225, 225, 3: 429, 428, 430, 431, 435, 225, 225, 225, 12: 225, 14: 225, 225, 23: 225, 31: 225, 36: 225, 225, 40: 225, 47: 427, 432, 225, 225, 436, 225, 438, 225, 225, 225, 442, 225, 225, 445, 62: 225, 225, 225, 225, 225, 225, 225, 225, 225, 225, 225, 225, 225},
		{226, 226, 3: 429, 428, 430, 431, 435, 226, 226, 226, 12: 226, 14: 226, 226, 23: 226, 31: 226, 36: 226, 226, 40: 226, 47: 427, 432, 226, 226, 436, 226, 438, 226, 226, 226, 442, 226, 226, 445, 62: 226, 226, 226, 226, 226, 226, 226, 226, 226, 226, 226, 226, 226},
		{227, 227, 3: 227, 227, 227, 227, 227, 227, 227, 227, 12: 227, 14: 227, 227, 23: 227, 31: 227, 36: 227, 227, 40: 227, 47: 227, 227, 227, 227, 227, 227, 227, 227, 227, 227, 227, 227, 227, 227, 62: 227, 227, 227, 227, 227, 227, 227, 227, 227, 227, 227, 227, 227},
		// 165
		{228, 228, 3: 429, 428, 228, 228, 228, 228, 228, 228, 12: 228, 14: 228, 228, 23: 228, 31: 228, 36: 228, 228, 40: 228, 47: 427, 432, 228, 228, 228, 228, 438, 228, 228, 228, 442, 228, 228, 445, 62: 228, 228, 228, 228, 228, 228, 228, 228, 228, 228, 228, 228, 228},
		{229, 229, 3: 429, 428, 229, 229, 229, 229, 229, 229, 12: 229, 14: 229, 229, 23: 229, 31: 229, 36: 229, 229, 40: 229, 47: 427, 432, 229, 229, 229, 229, 438, 229, 229, 229, 442, 229, 229, 445, 62: 229, 229, 229, 229, 229, 229, 229, 229, 229, 229, 229, 229, 229},
		{230, 230, 3: 230, 230, 230, 230, 230, 230, 230, 230, 12: 230, 14: 230, 230, 23: 230, 31: 230, 36: 230, 230, 40: 230, 47: 230, 230, 230, 230, 230, 230, 230, 230, 230, 230, 230, 230, 230, 230, 62: 230, 230, 230, 230, 230, 230, 230, 230, 230, 230, 230, 230, 230},
		{231, 231, 3: 231, 231, 231, 231, 231, 231, 231, 231, 12: 231, 14: 231, 231, 23: 231, 31: 231, 36: 231, 231, 40: 231, 47: 231, 231, 231, 231, 231, 231, 231, 231, 231, 231, 231, 231, 231, 231, 62: 231, 231, 231, 231, 231, 231, 231, 231, 231, 231, 231, 231, 231},
		{232, 232, 3: 232, 232, 232, 232, 232, 232, 232, 232, 12: 232, 14: 232, 232, 23: 232, 31: 232, 36: 232, 232, 40: 232, 47: 232, 232, 232, 232, 232, 232, 232, 232, 232, 232, 232, 232, 232, 232, 62: 232, 232, 232, 232, 232, 232, 232, 232, 232, 232, 232, 232, 232},
		// 170
		{5, 5, 9: 5, 5, 12: 5, 14: 425},
		{35, 35, 35, 9: 35, 35, 12: 35, 14: 35, 35, 35, 23: 35, 31: 35, 36: 35, 35, 39: 35, 35, 76: 35},
		{41, 41, 41, 9: 41, 41, 12: 41, 14: 41, 41, 41, 23: 41, 31: 41, 36: 41, 41, 39: 41, 41, 76: 41},
		{12: 471},
		{42, 42, 42, 9: 42, 42, 12: 42, 14: 42, 42, 42, 23: 42, 31: 42, 36: 42, 42, 39: 42, 42, 76: 42},
		// 175
		{13: 473},
		{125, 125, 125, 9: 125, 125, 12: 125, 14: 125, 125, 125, 23: 125, 31: 125, 36: 125, 125, 39: 125, 125, 76: 125, 97: 125},
		{189, 189, 189, 189, 9: 189, 189, 189, 189, 189, 189, 17: 189, 189, 189, 189, 189, 189, 31: 189, 76: 189},
		{198, 198, 198, 9: 198, 198, 12: 198, 14: 198, 198, 198, 23: 198, 31: 198, 36: 198, 198, 39: 198, 198, 76: 198},
		{13: 388, 98: 498},
		// 180
		{2: 402, 403, 11: 330, 488, 484, 17: 332, 399, 344, 345, 333, 368, 24: 404, 405, 406, 407, 408, 410, 411, 37: 483, 78: 409, 80: 492, 129: 486, 140: 487},
		{119, 119, 479, 403, 9: 119, 119, 330, 119, 401, 119, 119, 119, 332, 399, 344, 345, 333, 368, 119, 404, 405, 406, 407, 408, 410, 411, 119, 36: 119, 119, 39: 119, 119, 76: 119, 78: 409, 80: 481, 105: 480, 141: 482},
		{2: 402, 403, 11: 330, 488, 484, 17: 332, 399, 344, 345, 333, 368, 24: 404, 405, 406, 407, 408, 410, 411, 37: 483, 78: 409, 80: 485, 129: 486, 140: 487},
		{118, 118, 118, 9: 118, 118, 12: 118, 14: 118, 118, 118, 23: 118, 31: 118, 36: 118, 118, 39: 118, 118, 76: 118},
		{117, 117, 117, 9: 117, 117, 12: 117, 14: 117, 117, 117, 23: 117, 31: 117, 36: 117, 117, 39: 117, 117, 76: 117},
		// 185
		{112, 112, 112, 9: 112, 112, 12: 112, 14: 112, 112, 112, 23: 112, 31: 112, 36: 112, 112, 39: 112, 112, 76: 112},
		{2: 402, 403, 11: 330, 13: 401, 17: 332, 399, 344, 345, 333, 368, 24: 404, 405, 406, 407, 408, 410, 411, 78: 409, 80: 497},
		{2: 402, 403, 11: 330, 126, 401, 126, 17: 332, 399, 344, 345, 333, 368, 24: 404, 405, 406, 407, 408, 410, 411, 37: 494, 75: 472, 78: 409, 80: 495, 97: 126},
		{12: 471, 14: 143},
		{12: 142, 14: 142},
		// 190
		{12: 262, 14: 489, 112: 490},
		{140, 140, 140, 140, 9: 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 31: 140, 36: 140, 140, 39: 140, 140, 76: 140},
		{2: 402, 403, 11: 330, 261, 484, 17: 332, 399, 344, 345, 333, 368, 24: 404, 405, 406, 407, 408, 410, 411, 37: 483, 78: 409, 80: 492, 129: 493},
		{12: 491},
		{139, 139, 139, 139, 9: 139, 139, 139, 139, 139, 139, 139, 139, 139, 139, 139, 139, 139, 139, 139, 31: 139, 36: 139, 139, 39: 139, 139, 76: 139},
		// 195
		{12: 143, 14: 143},
		{12: 141, 14: 141},
		{2: 402, 403, 11: 330, 13: 401, 17: 332, 399, 344, 345, 333, 368, 24: 404, 405, 406, 407, 408, 410, 411, 78: 409, 80: 496},
		{12: 144, 14: 144},
		{12: 145, 14: 145},
		// 200
		{12: 146, 14: 146},
		{14: 400, 76: 499},
		{2: 477, 105: 478, 120: 500},
		{197, 197, 197, 9: 197, 197, 12: 197, 14: 197, 197, 197, 23: 197, 31: 197, 36: 197, 197, 39: 197, 197, 76: 197},
		{10, 10, 3: 10, 10, 10, 10, 10, 10, 10, 10, 12: 10, 14: 10, 10, 23: 10, 31: 10, 36: 10, 10, 40: 10, 47: 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 62: 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10},
		// 205
		{11, 11, 3: 11, 11, 11, 11, 11, 11, 11, 11, 12: 11, 14: 11, 11, 23: 11, 31: 11, 36: 11, 11, 40: 11, 47: 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 62: 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11},
		{12, 12, 3: 12, 12, 12, 12, 12, 12, 12, 12, 12: 12, 14: 12, 12, 23: 12, 31: 12, 36: 12, 12, 40: 12, 47: 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 62: 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12},
		{13, 13, 3: 13, 13, 13, 13, 13, 13, 13, 13, 12: 13, 14: 13, 13, 23: 13, 31: 13, 36: 13, 13, 40: 13, 47: 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 62: 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13},
		{15, 15, 3: 15, 15, 15, 15, 15, 15, 15, 15, 12: 15, 14: 15, 15, 23: 15, 31: 15, 36: 15, 15, 40: 15, 47: 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 62: 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15},
		{16, 16, 3: 16, 16, 16, 16, 16, 16, 16, 16, 12: 16, 14: 16, 16, 23: 16, 31: 16, 36: 16, 16, 40: 16, 47: 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 62: 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16},
		// 210
		{2: 509, 12: 27, 14: 27, 23: 27, 31: 27, 36: 27, 27},
		{14, 14, 3: 14, 14, 14, 14, 14, 14, 14, 14, 12: 14, 14: 14, 14, 23: 14, 31: 14, 36: 14, 14, 40: 14, 47: 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 62: 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14},
		{2: 346, 378, 382, 383, 384, 385, 386, 11: 330, 13: 417, 16: 315, 332, 399, 344, 345, 333, 368, 24: 334, 379, 348, 380, 335, 336, 337, 32: 311, 312, 313, 314, 38: 381, 41: 347, 353, 351, 350, 352, 339, 61: 510},
		{3: 429, 428, 430, 431, 435, 446, 12: 262, 14: 511, 47: 427, 432, 433, 434, 436, 437, 438, 439, 440, 441, 442, 443, 444, 445, 112: 512},
		{12: 261},
		// 215
		{12: 513},
		{128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 14: 128, 128, 23: 128, 31: 128, 36: 128, 128, 39: 128, 128, 47: 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 62: 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128},
		{12: 522, 516, 131: 524, 183: 523},
		{30, 30, 9: 30, 30},
		{2: 194, 194, 11: 194, 13: 194, 17: 194, 194, 194, 194, 194, 194, 97: 517, 136: 518},
		// 220
		{13: 388, 98: 520},
		{2: 402, 403, 11: 330, 13: 401, 17: 332, 399, 344, 345, 333, 368, 24: 404, 405, 406, 407, 408, 410, 411, 78: 409, 80: 519},
		{19, 19, 9: 19, 19, 12: 19},
		{14: 400, 76: 521},
		{2: 193, 193, 11: 193, 13: 193, 17: 193, 193, 193, 193, 193, 193},
		// 225
		{32, 32, 9: 32, 32},
		{525, 12: 114, 114: 526},
		{18, 12: 18},
		{12: 113, 516, 131: 528},
		{12: 527},
		// 230
		{31, 31, 9: 31, 31},
		{17, 12: 17},
		{2: 43, 43, 43, 43, 43, 43, 43, 11: 43, 13: 43, 16: 43, 43, 43, 43, 43, 43, 43, 32: 43, 43, 43, 43, 38: 43, 79: 43, 81: 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 43, 99: 43},
		{105, 15: 105},
		{575, 15: 58},
		// 235
		{15: 54, 147: 533},
		{15: 534, 143: 535},
		{1: 536, 9: 69, 69, 155: 537},
		{53, 53, 9: 53, 53},
		{70, 70, 9: 70, 70},
		// 240
		{9: 539, 540, 144: 541, 542, 177: 538},
		{1: 573, 9: 539, 540, 144: 541, 574},
		{2: 346, 378, 382, 383, 384, 385, 386, 11: 330, 13: 417, 16: 315, 332, 399, 344, 345, 333, 368, 24: 334, 379, 348, 380, 335, 336, 337, 32: 311, 312, 313, 314, 38: 381, 41: 347, 560, 351, 350, 352, 339, 61: 559, 79: 563, 121: 561, 133: 562},
		{23: 557, 79: 558},
		{103, 103, 346, 378, 382, 383, 384, 385, 386, 103, 103, 330, 13: 349, 16: 315, 332, 399, 344, 345, 333, 368, 24: 334, 379, 348, 380, 335, 336, 337, 32: 311, 312, 313, 314, 38: 381, 543, 41: 347, 353, 351, 350, 352, 339, 61: 340, 77: 331, 79: 549, 81: 377, 343, 356, 338, 357, 358, 359, 341, 361, 362, 364, 354, 369, 387, 355, 366, 100: 544, 102: 545, 360, 363, 106: 365, 108: 548, 367, 546, 547, 115: 550, 124: 551},
		// 245
		{1: 60, 9: 60, 60},
		{272, 272, 272, 272, 272, 272, 272, 272, 272, 11: 272, 13: 272, 16: 272, 272, 272, 272, 272, 272, 272, 32: 272, 272, 272, 272, 38: 272, 272, 79: 272, 81: 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 272, 146: 554},
		{102, 102, 9: 102, 102},
		{101, 101, 9: 101, 101},
		{100, 100, 9: 100, 100},
		// 250
		{99, 99, 9: 99, 99},
		{98, 98, 9: 98, 98},
		{97, 97, 9: 97, 97},
		{96, 96, 9: 96, 96},
		{552, 61, 9: 61, 61},
		// 255
		{103, 103, 346, 378, 382, 383, 384, 385, 386, 103, 103, 330, 13: 349, 16: 315, 332, 399, 344, 345, 333, 368, 24: 334, 379, 348, 380, 335, 336, 337, 32: 311, 312, 313, 314, 38: 381, 543, 41: 347, 353, 351, 350, 352, 339, 61: 340, 77: 331, 79: 549, 81: 377, 343, 356, 338, 357, 358, 359, 341, 361, 362, 364, 354, 369, 387, 355, 366, 100: 544, 102: 545, 360, 363, 106: 365, 108: 548, 367, 546, 547, 115: 553},
		{95, 95, 9: 95, 95},
		{103, 103, 346, 378, 382, 383, 384, 385, 386, 11: 330, 13: 349, 16: 315, 332, 399, 344, 345, 333, 368, 24: 334, 379, 348, 380, 335, 336, 337, 32: 311, 312, 313, 314, 38: 381, 543, 41: 347, 353, 351, 350, 352, 339, 61: 340, 77: 331, 79: 549, 81: 377, 343, 356, 338, 357, 358, 359, 341, 361, 362, 364, 354, 369, 387, 355, 366, 100: 544, 102: 545, 360, 363, 106: 365, 108: 548, 367, 546, 547, 115: 550, 124: 555},
		{552, 556},
		{271, 271, 9: 271, 271},
		// 260
		{64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 13: 64, 16: 64, 64, 64, 64, 64, 64, 64, 32: 64, 64, 64, 64, 38: 64, 64, 79: 64, 81: 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64},
		{62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 13: 62, 16: 62, 62, 62, 62, 62, 62, 62, 32: 62, 62, 62, 62, 38: 62, 62, 79: 62, 81: 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62},
		{3: 429, 428, 430, 431, 435, 446, 12: 295, 14: 295, 23: 295, 31: 295, 36: 295, 295, 47: 427, 432, 433, 434, 436, 437, 438, 439, 440, 441, 442, 443, 444, 445},
		{2: 509, 12: 294, 14: 294, 23: 294, 31: 294, 36: 294, 294},
		{12: 293, 14: 293, 23: 293, 31: 293, 36: 293, 293},
		// 265
		{14: 564, 23: 565, 31: 566, 36: 567},
		{63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 13: 63, 16: 63, 63, 63, 63, 63, 63, 63, 32: 63, 63, 63, 63, 38: 63, 63, 79: 63, 81: 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63},
		{2: 346, 378, 382, 383, 384, 385, 386, 11: 330, 13: 417, 16: 315, 332, 399, 344, 345, 333, 368, 24: 334, 379, 348, 380, 335, 336, 337, 32: 311, 312, 313, 314, 38: 381, 41: 347, 560, 351, 350, 352, 339, 61: 559, 121: 572},
		{67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 13: 67, 16: 67, 67, 67, 67, 67, 67, 67, 32: 67, 67, 67, 67, 38: 67, 67, 79: 67, 81: 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67},
		{2: 346, 378, 382, 383, 384, 385, 386, 11: 330, 13: 417, 16: 315, 332, 399, 344, 345, 333, 368, 24: 334, 379, 348, 380, 335, 336, 337, 32: 311, 312, 313, 314, 38: 381, 41: 347, 353, 351, 350, 352, 339, 61: 570},
		// 270
		{2: 346, 378, 382, 383, 384, 385, 386, 11: 330, 13: 417, 16: 315, 332, 399, 344, 345, 333, 368, 24: 334, 379, 348, 380, 335, 336, 337, 32: 311, 312, 313, 314, 38: 381, 41: 347, 353, 351, 350, 352, 339, 61: 568},
		{3: 429, 428, 430, 431, 435, 446, 23: 569, 47: 427, 432, 433, 434, 436, 437, 438, 439, 440, 441, 442, 443, 444, 445},
		{65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 13: 65, 16: 65, 65, 65, 65, 65, 65, 65, 32: 65, 65, 65, 65, 38: 65, 65, 79: 65, 81: 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65},
		{3: 429, 428, 430, 431, 435, 446, 23: 571, 47: 427, 432, 433, 434, 436, 437, 438, 439, 440, 441, 442, 443, 444, 445},
		{66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 13: 66, 16: 66, 66, 66, 66, 66, 66, 66, 32: 66, 66, 66, 66, 38: 66, 66, 79: 66, 81: 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66},
		// 275
		{12: 292, 14: 292, 23: 292, 31: 292, 36: 292, 292},
		{68, 68, 9: 68, 68},
		{1: 59, 9: 59, 59},
		{2: 346, 378, 382, 383, 384, 385, 386, 11: 330, 13: 577, 15: 57, 315, 332, 399, 344, 345, 333, 368, 24: 334, 379, 348, 380, 335, 336, 337, 32: 311, 312, 313, 314, 38: 381, 41: 347, 353, 351, 350, 352, 339, 61: 576},
		{3: 429, 428, 430, 431, 435, 446, 15: 56, 47: 427, 432, 433, 434, 436, 437, 438, 439, 440, 441, 442, 443, 444, 445},
		// 280
		{2: 196, 196, 196, 196, 196, 196, 196, 11: 196, 15: 196, 36: 578, 39: 196, 47: 196, 196, 196, 196, 196, 196, 196, 196, 196, 196, 196, 196, 196, 196, 75: 196, 97: 418, 122: 419},
		{2: 346, 580, 11: 330, 13: 417, 16: 315, 332, 399, 344, 345, 333, 368, 24: 334, 379, 348, 380, 335, 336, 337, 32: 311, 312, 313, 314, 41: 347, 353, 351, 350, 579},
		{2: 587, 11: 590, 39: 588, 75: 589, 113: 592, 134: 591},
		{3: 580, 11: 330, 17: 332, 399, 344, 345, 333, 368, 24: 582, 379, 583, 380, 584, 585, 586, 42: 581},
		{2: 27},
		// 285
		{2: 26},
		{2: 24},
		{2: 22},
		{2: 21},
		{2: 20},
		// 290
		{2: 346, 378, 382, 383, 384, 385, 386, 11: 330, 625, 417, 16: 315, 332, 399, 344, 345, 333, 368, 24: 334, 379, 348, 380, 335, 336, 337, 32: 311, 312, 313, 314, 38: 381, 41: 347, 560, 351, 350, 352, 339, 61: 559, 121: 561, 133: 624},
		{1: 613, 346, 378, 382, 383, 384, 385, 386, 11: 330, 13: 417, 16: 315, 332, 399, 344, 345, 333, 368, 24: 334, 379, 348, 380, 335, 336, 337, 32: 311, 312, 313, 314, 38: 381, 588, 41: 347, 353, 351, 350, 352, 339, 61: 610, 113: 609, 135: 611, 156: 612},
		{2: 603, 13: 604},
		{2: 346, 378, 382, 383, 384, 385, 386, 11: 330, 13: 417, 16: 315, 332, 399, 344, 345, 333, 368, 212, 334, 379, 348, 380, 335, 336, 337, 32: 311, 312, 313, 314, 38: 381, 41: 347, 353, 351, 350, 352, 339, 61: 593, 127: 594},
		{130, 130, 130, 130, 130, 130, 130, 130, 130, 130, 130, 130, 130, 14: 130, 130, 23: 130, 31: 130, 36: 130, 130, 39: 130, 130, 47: 130, 130, 130, 130, 130, 130, 130, 130, 130, 130, 130, 130, 130, 130, 62: 130, 130, 130, 130, 130, 130, 130, 130, 130, 130, 130, 130, 130, 130},
		// 295
		{129, 129, 129, 129, 129, 129, 129, 129, 129, 129, 129, 129, 129, 14: 129, 129, 23: 129, 31: 129, 36: 129, 129, 39: 129, 129, 47: 129, 129, 129, 129, 129, 129, 129, 129, 129, 129, 129, 129, 129, 129, 62: 129, 129, 129, 129, 129, 129, 129, 129, 129, 129, 129, 129, 129, 129},
		{3: 429, 428, 430, 431, 435, 446, 23: 211, 40: 602, 47: 427, 432, 433, 434, 436, 437, 438, 439, 440, 441, 442, 443, 444, 445},
		{23: 595},
		{2: 346, 378, 382, 383, 384, 385, 386, 11: 330, 13: 417, 16: 315, 332, 399, 344, 345, 333, 368, 212, 334, 379, 348, 380, 335, 336, 337, 32: 311, 312, 313, 314, 38: 381, 40: 212, 347, 353, 351, 350, 352, 339, 61: 596, 127: 597},
		{3: 429, 428, 430, 431, 435, 446, 23: 211, 40: 211, 47: 427, 432, 433, 434, 436, 437, 438, 439, 440, 441, 442, 443, 444, 445},
		// 300
		{23: 598, 40: 599},
		{2: 346, 378, 382, 383, 384, 385, 386, 11: 330, 13: 417, 16: 315, 332, 399, 344, 345, 333, 368, 24: 334, 379, 348, 380, 335, 336, 337, 32: 311, 312, 313, 314, 38: 381, 40: 212, 347, 353, 351, 350, 352, 339, 61: 596, 127: 600},
		{131, 131, 131, 131, 131, 131, 131, 131, 131, 131, 131, 131, 131, 14: 131, 131, 23: 131, 31: 131, 36: 131, 131, 39: 131, 131, 47: 131, 131, 131, 131, 131, 131, 131, 131, 131, 131, 131, 131, 131, 131, 62: 131, 131, 131, 131, 131, 131, 131, 131, 131, 131, 131, 131, 131, 131},
		{40: 601},
		{132, 132, 132, 132, 132, 132, 132, 132, 132, 132, 132, 132, 132, 14: 132, 132, 23: 132, 31: 132, 36: 132, 132, 39: 132, 132, 47: 132, 132, 132, 132, 132, 132, 132, 132, 132, 132, 132, 132, 132, 132, 62: 132, 132, 132, 132, 132, 132, 132, 132, 132, 132, 132, 132, 132, 132},
		// 305
		{133, 133, 133, 133, 133, 133, 133, 133, 133, 133, 133, 133, 133, 14: 133, 133, 23: 133, 31: 133, 36: 133, 133, 39: 133, 133, 47: 133, 133, 133, 133, 133, 133, 133, 133, 133, 133, 133, 133, 133, 133, 62: 133, 133, 133, 133, 133, 133, 133, 133, 133, 133, 133, 133, 133, 133},
		{2: 402, 403, 11: 330, 13: 401, 17: 332, 399, 344, 345, 333, 368, 24: 404, 405, 406, 407, 408, 410, 411, 78: 409, 80: 606, 605},
		{134, 134, 134, 134, 134, 134, 134, 134, 134, 134, 134, 134, 134, 14: 134, 134, 23: 134, 31: 134, 36: 134, 134, 39: 134, 134, 47: 134, 134, 134, 134, 134, 134, 134, 134, 134, 134, 134, 134, 134, 134, 62: 134, 134, 134, 134, 134, 134, 134, 134, 134, 134, 134, 134, 134, 134},
		{12: 608},
		{12: 607},
		// 310
		{135, 135, 135, 135, 135, 135, 135, 135, 135, 135, 135, 135, 135, 14: 135, 135, 23: 135, 31: 135, 36: 135, 135, 39: 135, 135, 47: 135, 135, 135, 135, 135, 135, 135, 135, 135, 135, 135, 135, 135, 135, 62: 135, 135, 135, 135, 135, 135, 135, 135, 135, 135, 135, 135, 135, 135},
		{2: 136, 11: 136, 15: 55, 39: 136, 75: 136},
		{1: 260, 14: 260, 23: 621},
		{1: 257, 3: 429, 428, 430, 431, 435, 446, 14: 257, 23: 618, 47: 427, 432, 433, 434, 436, 437, 438, 439, 440, 441, 442, 443, 444, 445},
		{1: 254, 14: 254},
		// 315
		{1: 262, 14: 614, 112: 615},
		{248, 248, 248, 248, 248, 248, 248, 248, 248, 248, 248, 248, 248, 14: 248, 248, 23: 248, 31: 248, 36: 248, 248, 39: 248, 248, 47: 248, 248, 248, 248, 248, 248, 248, 248, 248, 248, 248, 248, 248, 248, 62: 248, 248, 248, 248, 248, 248, 248, 248, 248, 248, 248, 248, 248, 248},
		{1: 261, 346, 378, 382, 383, 384, 385, 386, 11: 330, 13: 417, 16: 315, 332, 399, 344, 345, 333, 368, 24: 334, 379, 348, 380, 335, 336, 337, 32: 311, 312, 313, 314, 38: 381, 588, 41: 347, 353, 351, 350, 352, 339, 61: 610, 113: 609, 135: 617},
		{1: 616},
		{247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 14: 247, 247, 23: 247, 31: 247, 36: 247, 247, 39: 247, 247, 47: 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 62: 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247, 247},
		// 320
		{1: 253, 14: 253},
		{2: 346, 378, 382, 383, 384, 385, 386, 11: 330, 13: 417, 16: 315, 332, 399, 344, 345, 333, 368, 24: 334, 379, 348, 380, 335, 336, 337, 32: 311, 312, 313, 314, 38: 381, 588, 41: 347, 353, 351, 350, 352, 339, 61: 620, 113: 619},
		{1: 256, 14: 256},
		{1: 255, 3: 429, 428, 430, 431, 435, 446, 14: 255, 47: 427, 432, 433, 434, 436, 437, 438, 439, 440, 441, 442, 443, 444, 445},
		{2: 346, 378, 382, 383, 384, 385, 386, 11: 330, 13: 417, 16: 315, 332, 399, 344, 345, 333, 368, 24: 334, 379, 348, 380, 335, 336, 337, 32: 311, 312, 313, 314, 38: 381, 588, 41: 347, 353, 351, 350, 352, 339, 61: 623, 113: 622},
		// 325
		{1: 259, 14: 259},
		{1: 258, 3: 429, 428, 430, 431, 435, 446, 14: 258, 47: 427, 432, 433, 434, 436, 437, 438, 439, 440, 441, 442, 443, 444, 445},
		{12: 262, 14: 626, 37: 628, 112: 627},
		{268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 14: 268, 268, 23: 268, 31: 268, 36: 268, 268, 39: 268, 268, 47: 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 62: 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268},
		{2: 346, 378, 382, 383, 384, 385, 386, 11: 330, 261, 417, 16: 315, 332, 399, 344, 345, 333, 368, 24: 334, 379, 348, 380, 335, 336, 337, 32: 311, 312, 313, 314, 38: 381, 41: 347, 560, 351, 350, 352, 339, 61: 559, 121: 572},
		// 330
		{12: 631},
		{12: 262, 14: 511, 112: 629},
		{12: 630},
		{266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 14: 266, 266, 23: 266, 31: 266, 36: 266, 266, 39: 266, 266, 47: 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 62: 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266},
		{267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 14: 267, 267, 23: 267, 31: 267, 36: 267, 267, 39: 267, 267, 47: 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 62: 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267},
		// 335
		{166, 166, 166, 166, 166, 166, 166, 166, 166, 11: 166, 13: 166, 15: 166, 166, 166, 166, 166, 166, 166, 166, 32: 166, 166, 166, 166, 38: 166, 166, 79: 166, 81: 166, 166, 166, 166, 166, 166, 166, 166, 166, 166, 166, 166, 166, 166},
		{165, 165, 165, 165, 165, 165, 165, 165, 165, 11: 165, 13: 165, 15: 165, 165, 165, 165, 165, 165, 165, 165, 32: 165, 165, 165, 165, 38: 165, 165, 79: 165, 81: 165, 165, 165, 165, 165, 165, 165, 165, 165, 165, 165, 165, 165, 165},
		{1: 635, 72, 72, 13: 72, 154: 636},
		{73, 73, 73, 9: 73, 73, 12: 73, 14: 73, 73, 73, 23: 73, 31: 73, 36: 73, 73, 39: 73, 73, 76: 73},
		{2: 641, 639, 13: 637, 78: 640, 98: 638, 142: 642, 176: 643},
		// 340
		{126, 126, 190, 190, 11: 190, 13: 190, 190, 16: 126, 190, 190, 190, 190, 190, 190, 75: 472},
		{2: 402, 403, 11: 330, 13: 401, 400, 17: 332, 399, 344, 345, 333, 368, 24: 404, 405, 406, 407, 408, 410, 411, 78: 409, 80: 663},
		{2: 658, 13: 401, 78: 657},
		{52, 52, 16: 653, 116: 656},
		{3: 649, 13: 401, 78: 648},
		// 345
		{75, 75},
		{644, 114, 114: 645},
		{1: 113, 641, 639, 13: 637, 78: 640, 98: 638, 142: 647},
		{1: 646},
		{71, 71, 71, 9: 71, 71, 12: 71, 14: 71, 71, 71, 23: 71, 31: 71, 36: 71, 71, 39: 71, 71, 76: 71},
		// 350
		{74, 74},
		{12: 654},
		{13: 401, 78: 650},
		{12: 651},
		{52, 52, 16: 653, 116: 652},
		// 355
		{77, 77},
		{51, 51},
		{52, 52, 16: 653, 116: 655},
		{78, 78},
		{79, 79},
		// 360
		{52, 52, 16: 653, 116: 662},
		{13: 401, 78: 659},
		{12: 660},
		{52, 52, 16: 653, 116: 661},
		{76, 76},
		// 365
		{81, 81},
		{52, 52, 16: 653, 116: 664},
		{80, 80},
		{207, 207, 9: 207, 207, 14: 425},
		{85, 85, 9: 85, 85},
		// 370
		{88, 88, 9: 88, 88},
		{89, 89, 3: 429, 428, 430, 431, 435, 446, 89, 89, 47: 427, 432, 433, 434, 436, 437, 438, 439, 440, 441, 442, 443, 444, 445},
		{92, 92, 3: 429, 428, 430, 431, 435, 446, 92, 92, 47: 427, 432, 433, 434, 436, 437, 438, 439, 440, 441, 442, 443, 444, 445},
		{93, 93, 9: 93, 93},
		{94, 94, 9: 94, 94},
		// 375
		{15: 534, 143: 673},
		{115, 115, 9: 115, 115},
		{2: 675, 13: 604},
		{2: 402, 403, 11: 330, 13: 401, 17: 332, 399, 344, 345, 333, 368, 24: 404, 405, 406, 407, 408, 410, 411, 78: 409, 80: 606, 676},
		{12: 677},
		// 380
		{136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 14: 136, 136, 23: 136, 31: 136, 36: 136, 136, 39: 136, 136, 47: 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 62: 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136},
		{1: 684, 346, 378, 382, 383, 384, 385, 386, 11: 330, 13: 417, 15: 632, 315, 332, 399, 344, 345, 333, 368, 24: 334, 379, 348, 380, 335, 336, 337, 32: 311, 312, 313, 314, 38: 381, 633, 41: 347, 353, 351, 350, 352, 339, 61: 680, 101: 678, 119: 681, 139: 682, 170: 683},
		{137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 14: 137, 137, 23: 137, 31: 137, 36: 137, 137, 39: 137, 137, 47: 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 62: 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137},
		{1: 164, 3: 429, 428, 430, 431, 435, 446, 14: 164, 23: 692, 47: 427, 432, 433, 434, 436, 437, 438, 439, 440, 441, 442, 443, 444, 445},
		{1: 161, 14: 161, 23: 689},
		// 385
		{1: 158, 14: 158},
		{1: 262, 14: 685, 112: 686},
		{156, 156, 156, 156, 156, 156, 156, 156, 156, 156, 156, 156, 156, 14: 156, 156, 23: 156, 31: 156, 36: 156, 156, 39: 156, 156, 47: 156, 156, 156, 156, 156, 156, 156, 156, 156, 156, 156, 156, 156, 156, 62: 156, 156, 156, 156, 156, 156, 156, 156, 156, 156, 156, 156, 156, 156},
		{1: 261, 346, 378, 382, 383, 384, 385, 386, 11: 330, 13: 417, 15: 632, 315, 332, 399, 344, 345, 333, 368, 24: 334, 379, 348, 380, 335, 336, 337, 32: 311, 312, 313, 314, 38: 381, 633, 41: 347, 353, 351, 350, 352, 339, 61: 680, 101: 678, 119: 681, 139: 688},
		{1: 687},
		// 390
		{155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 14: 155, 155, 23: 155, 31: 155, 36: 155, 155, 39: 155, 155, 47: 155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 62: 155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 155, 155},
		{1: 157, 14: 157},
		{2: 346, 378, 382, 383, 384, 385, 386, 11: 330, 13: 417, 15: 632, 315, 332, 399, 344, 345, 333, 368, 24: 334, 379, 348, 380, 335, 336, 337, 32: 311, 312, 313, 314, 38: 381, 633, 41: 347, 353, 351, 350, 352, 339, 61: 690, 101: 678, 119: 691},
		{1: 160, 3: 429, 428, 430, 431, 435, 446, 14: 160, 47: 427, 432, 433, 434, 436, 437, 438, 439, 440, 441, 442, 443, 444, 445},
		{1: 159, 14: 159},
		// 395
		{2: 346, 378, 382, 383, 384, 385, 386, 11: 330, 13: 417, 15: 632, 315, 332, 399, 344, 345, 333, 368, 24: 334, 379, 348, 380, 335, 336, 337, 32: 311, 312, 313, 314, 38: 381, 633, 41: 347, 353, 351, 350, 352, 339, 61: 693, 101: 678, 119: 694},
		{1: 163, 3: 429, 428, 430, 431, 435, 446, 14: 163, 47: 427, 432, 433, 434, 436, 437, 438, 439, 440, 441, 442, 443, 444, 445},
		{1: 162, 14: 162},
		{103, 103, 346, 378, 382, 383, 384, 385, 386, 103, 103, 330, 13: 349, 16: 315, 332, 399, 344, 345, 333, 368, 24: 334, 379, 348, 380, 335, 336, 337, 32: 311, 312, 313, 314, 38: 381, 543, 41: 347, 353, 351, 350, 352, 339, 61: 340, 77: 331, 79: 549, 81: 377, 343, 356, 338, 357, 358, 359, 341, 361, 362, 364, 354, 369, 387, 355, 366, 100: 544, 102: 545, 360, 363, 106: 365, 108: 548, 367, 546, 547, 115: 696},
		{87, 87, 9: 87, 87},
		// 400
		{15: 632, 39: 633, 101: 698},
		{103, 103, 346, 378, 382, 383, 384, 385, 386, 11: 330, 13: 349, 16: 315, 332, 399, 344, 345, 333, 368, 24: 334, 379, 348, 380, 335, 336, 337, 32: 311, 312, 313, 314, 38: 381, 543, 41: 347, 353, 351, 350, 352, 339, 61: 340, 77: 331, 79: 549, 81: 377, 343, 356, 338, 357, 358, 359, 341, 361, 362, 364, 354, 369, 387, 355, 366, 100: 544, 102: 545, 360, 363, 106: 365, 108: 548, 367, 546, 547, 115: 550, 124: 699},
		{552, 700},
		{149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 14: 149, 149, 23: 149, 31: 149, 36: 149, 149, 39: 149, 149, 47: 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 62: 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149},
		{3: 429, 428, 430, 431, 435, 446, 12: 704, 47: 427, 432, 433, 434, 436, 437, 438, 439, 440, 441, 442, 443, 444, 445},
		// 405
		{2: 509, 12: 703},
		{152, 152, 152, 152, 152, 152, 152, 152, 152, 152, 152, 152, 152, 14: 152, 152, 23: 152, 31: 152, 36: 152, 152, 39: 152, 152, 47: 152, 152, 152, 152, 152, 152, 152, 152, 152, 152, 152, 152, 152, 152, 62: 152, 152, 152, 152, 152, 152, 152, 152, 152, 152, 152, 152, 152, 152},
		{153, 153, 153, 153, 153, 153, 153, 153, 153, 153, 153, 153, 153, 14: 153, 153, 23: 153, 31: 153, 36: 153, 153, 39: 153, 153, 47: 153, 153, 153, 153, 153, 153, 153, 153, 153, 153, 153, 153, 153, 153, 62: 153, 153, 153, 153, 153, 153, 153, 153, 153, 153, 153, 153, 153, 153},
		{2: 402, 403, 11: 330, 13: 401, 17: 332, 399, 344, 345, 333, 368, 24: 404, 405, 406, 407, 408, 410, 411, 78: 409, 80: 706},
		{40: 707},
		// 410
		{2: 402, 403, 11: 330, 13: 401, 17: 332, 399, 344, 345, 333, 368, 24: 404, 405, 406, 407, 408, 410, 411, 78: 409, 80: 708},
		{154, 154, 154, 9: 154, 154, 12: 154, 14: 154, 154, 154, 23: 154, 31: 154, 36: 154, 154, 39: 154, 154, 76: 154},
		{1: 710, 13: 173, 150: 711},
		{174, 174, 174, 9: 174, 174, 12: 174, 14: 174, 174, 174, 23: 174, 31: 174, 36: 174, 174, 39: 174, 174, 76: 174},
		{13: 713, 78: 714, 138: 715, 169: 712},
		// 415
		{719, 114, 114: 718},
		{126, 126, 171, 75: 472, 151: 716},
		{169, 169},
		{168, 168},
		{2: 477, 105: 478, 120: 717},
		// 420
		{170, 170},
		{1: 721},
		{1: 113, 13: 713, 78: 714, 138: 720},
		{167, 167},
		{172, 172, 172, 9: 172, 172, 12: 172, 14: 172, 172, 172, 23: 172, 31: 172, 36: 172, 172, 39: 172, 172, 76: 172},
		// 425
		{737, 15: 188},
		{15: 724, 125: 725},
		{270, 270, 270, 270, 270, 270, 270, 270, 270, 11: 270, 13: 270, 16: 270, 270, 270, 270, 270, 270, 270, 32: 270, 270, 270, 270, 38: 270, 270, 79: 270, 81: 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 148: 734},
		{237, 237, 9: 237, 237, 117: 237, 159: 726},
		{235, 235, 9: 235, 235, 117: 727, 158: 728, 160: 729},
		// 430
		{39: 543, 82: 730, 100: 731},
		{236, 236, 9: 236, 236, 117: 236},
		{186, 186, 9: 186, 186},
		{106, 2: 346, 378, 382, 383, 384, 385, 386, 11: 330, 13: 417, 15: 106, 315, 332, 399, 344, 345, 333, 368, 24: 334, 379, 348, 380, 335, 336, 337, 32: 311, 312, 313, 314, 38: 381, 41: 347, 353, 351, 350, 352, 339, 61: 340, 77: 331, 95: 355, 530, 107: 722, 137: 732},
		{234, 234, 9: 234, 234},
		// 435
		{15: 724, 125: 733},
		{238, 238, 9: 238, 238, 117: 238},
		{103, 103, 346, 378, 382, 383, 384, 385, 386, 11: 330, 13: 349, 16: 315, 332, 399, 344, 345, 333, 368, 24: 334, 379, 348, 380, 335, 336, 337, 32: 311, 312, 313, 314, 38: 381, 543, 41: 347, 353, 351, 350, 352, 339, 61: 340, 77: 331, 79: 549, 81: 377, 343, 356, 338, 357, 358, 359, 341, 361, 362, 364, 354, 369, 387, 355, 366, 100: 544, 102: 545, 360, 363, 106: 365, 108: 548, 367, 546, 547, 115: 550, 124: 735},
		{552, 736},
		{269, 269, 9: 269, 269, 117: 269},
		// 440
		{2: 346, 378, 382, 383, 384, 385, 386, 11: 330, 13: 417, 15: 106, 315, 332, 399, 344, 345, 333, 368, 24: 334, 379, 348, 380, 335, 336, 337, 32: 311, 312, 313, 314, 38: 381, 41: 347, 353, 351, 350, 352, 339, 61: 340, 77: 331, 95: 355, 530, 107: 738},
		{15: 187},
		{13: 741},
		{2: 479, 403, 11: 330, 13: 120, 15: 119, 17: 332, 399, 344, 345, 333, 368, 24: 404, 405, 406, 407, 408, 410, 411, 39: 119, 78: 409, 80: 481, 105: 480, 141: 482},
		{2: 194, 97: 517, 136: 742},
		// 445
		{2: 477, 105: 478, 120: 743},
		{200, 39: 200, 149: 744},
		{202, 39: 543, 100: 745, 164: 746},
		{201},
		{199},
		// 450
		{14: 425, 31: 758, 36: 770, 62: 759, 761, 760, 762, 763, 764, 765, 766, 767, 768, 769},
		{15: 206},
		{754, 15: 204},
		{15: 724, 125: 753},
		{2: 346, 378, 382, 383, 384, 385, 386, 11: 330, 13: 417, 16: 315, 332, 399, 344, 345, 333, 368, 24: 334, 379, 348, 380, 335, 336, 337, 32: 311, 312, 313, 314, 38: 381, 41: 347, 353, 351, 350, 352, 339, 61: 752},
		// 455
		{3: 429, 428, 430, 431, 435, 446, 15: 122, 47: 427, 432, 433, 434, 436, 437, 438, 439, 440, 441, 442, 443, 444, 445},
		{203, 203, 9: 203, 203},
		{106, 2: 346, 378, 382, 383, 384, 385, 386, 11: 330, 13: 417, 16: 315, 332, 399, 344, 345, 333, 368, 24: 334, 379, 348, 380, 335, 336, 337, 32: 311, 312, 313, 314, 38: 381, 41: 347, 353, 351, 350, 352, 339, 61: 340, 77: 331, 95: 355, 530, 107: 755},
		{756},
		{2: 346, 378, 382, 383, 384, 385, 386, 11: 330, 13: 417, 15: 106, 315, 332, 399, 344, 345, 333, 368, 24: 334, 379, 348, 380, 335, 336, 337, 32: 311, 312, 313, 314, 38: 381, 41: 347, 353, 351, 350, 352, 339, 61: 340, 77: 331, 95: 355, 530, 107: 757},
		// 460
		{15: 205},
		{2: 346, 378, 382, 383, 384, 385, 386, 11: 330, 13: 417, 16: 315, 332, 399, 344, 345, 333, 368, 24: 334, 379, 348, 380, 335, 336, 337, 32: 311, 312, 313, 314, 38: 381, 41: 347, 353, 351, 350, 352, 339, 61: 415, 77: 785, 130: 786},
		{2: 346, 378, 382, 383, 384, 385, 386, 11: 330, 13: 417, 16: 315, 332, 399, 344, 345, 333, 368, 24: 334, 379, 348, 380, 335, 336, 337, 32: 311, 312, 313, 314, 38: 381, 41: 347, 353, 351, 350, 352, 339, 61: 415, 77: 784},
		{2: 346, 378, 382, 383, 384, 385, 386, 11: 330, 13: 417, 16: 315, 332, 399, 344, 345, 333, 368, 24: 334, 379, 348, 380, 335, 336, 337, 32: 311, 312, 313, 314, 38: 381, 41: 347, 353, 351, 350, 352, 339, 61: 415, 77: 783},
		{2: 346, 378, 382, 383, 384, 385, 386, 11: 330, 13: 417, 16: 315, 332, 399, 344, 345, 333, 368, 24: 334, 379, 348, 380, 335, 336, 337, 32: 311, 312, 313, 314, 38: 381, 41: 347, 353, 351, 350, 352, 339, 61: 415, 77: 782},
		// 465
		{2: 346, 378, 382, 383, 384, 385, 386, 11: 330, 13: 417, 16: 315, 332, 399, 344, 345, 333, 368, 24: 334, 379, 348, 380, 335, 336, 337, 32: 311, 312, 313, 314, 38: 381, 41: 347, 353, 351, 350, 352, 339, 61: 415, 77: 781},
		{2: 346, 378, 382, 383, 384, 385, 386, 11: 330, 13: 417, 16: 315, 332, 399, 344, 345, 333, 368, 24: 334, 379, 348, 380, 335, 336, 337, 32: 311, 312, 313, 314, 38: 381, 41: 347, 353, 351, 350, 352, 339, 61: 415, 77: 780},
		{2: 346, 378, 382, 383, 384, 385, 386, 11: 330, 13: 417, 16: 315, 332, 399, 344, 345, 333, 368, 24: 334, 379, 348, 380, 335, 336, 337, 32: 311, 312, 313, 314, 38: 381, 41: 347, 353, 351, 350, 352, 339, 61: 415, 77: 779},
		{2: 346, 378, 382, 383, 384, 385, 386, 11: 330, 13: 417, 16: 315, 332, 399, 344, 345, 333, 368, 24: 334, 379, 348, 380, 335, 336, 337, 32: 311, 312, 313, 314, 38: 381, 41: 347, 353, 351, 350, 352, 339, 61: 415, 77: 778},
		{2: 346, 378, 382, 383, 384, 385, 386, 11: 330, 13: 417, 16: 315, 332, 399, 344, 345, 333, 368, 24: 334, 379, 348, 380, 335, 336, 337, 32: 311, 312, 313, 314, 38: 381, 41: 347, 353, 351, 350, 352, 339, 61: 415, 77: 777},
		// 470
		{2: 346, 378, 382, 383, 384, 385, 386, 11: 330, 13: 417, 16: 315, 332, 399, 344, 345, 333, 368, 24: 334, 379, 348, 380, 335, 336, 337, 32: 311, 312, 313, 314, 38: 381, 41: 347, 353, 351, 350, 352, 339, 61: 415, 77: 776},
		{2: 346, 378, 382, 383, 384, 385, 386, 11: 330, 13: 417, 16: 315, 332, 399, 344, 345, 333, 368, 24: 334, 379, 348, 380, 335, 336, 337, 32: 311, 312, 313, 314, 38: 381, 41: 347, 353, 351, 350, 352, 339, 61: 415, 77: 775},
		{2: 346, 378, 382, 383, 384, 385, 386, 11: 330, 13: 417, 16: 315, 332, 399, 344, 345, 333, 368, 24: 334, 379, 348, 380, 335, 336, 337, 32: 311, 312, 313, 314, 38: 381, 41: 347, 353, 351, 350, 352, 339, 61: 415, 77: 774},
		{2: 346, 378, 382, 383, 384, 385, 386, 11: 330, 13: 417, 16: 315, 332, 399, 344, 345, 333, 368, 24: 334, 379, 348, 380, 335, 336, 337, 32: 311, 312, 313, 314, 38: 381, 41: 347, 353, 351, 350, 352, 339, 61: 415, 77: 771, 130: 772},
		{107, 107, 9: 107, 107, 14: 425, 107},
		// 475
		{2: 346, 378, 382, 383, 384, 385, 386, 11: 330, 13: 417, 16: 315, 332, 399, 344, 345, 333, 368, 24: 334, 379, 348, 380, 335, 336, 337, 32: 311, 312, 313, 314, 38: 381, 41: 347, 353, 351, 350, 352, 339, 61: 773},
		{3: 429, 428, 430, 431, 435, 446, 15: 123, 47: 427, 432, 433, 434, 436, 437, 438, 439, 440, 441, 442, 443, 444, 445},
		{278, 278, 9: 278, 278, 14: 425, 278},
		{279, 279, 9: 279, 279, 14: 425, 279},
		{280, 280, 9: 280, 280, 14: 425, 280},
		// 480
		{281, 281, 9: 281, 281, 14: 425, 281},
		{282, 282, 9: 282, 282, 14: 425, 282},
		{283, 283, 9: 283, 283, 14: 425, 283},
		{284, 284, 9: 284, 284, 14: 425, 284},
		{285, 285, 9: 285, 285, 14: 425, 285},
		// 485
		{286, 286, 9: 286, 286, 14: 425, 286},
		{287, 287, 9: 287, 287, 14: 425, 287},
		{288, 288, 9: 288, 288, 14: 425, 288},
		{289, 289, 9: 289, 289, 14: 425, 289},
		{2: 346, 378, 382, 383, 384, 385, 386, 11: 330, 13: 417, 16: 315, 332, 399, 344, 345, 333, 368, 24: 334, 379, 348, 380, 335, 336, 337, 32: 311, 312, 313, 314, 38: 381, 41: 347, 353, 351, 350, 352, 339, 61: 787},
		// 490
		{3: 429, 428, 430, 431, 435, 446, 15: 124, 47: 427, 432, 433, 434, 436, 437, 438, 439, 440, 441, 442, 443, 444, 445},
		{109, 109, 9: 109, 109, 15: 109},
		{108, 108, 9: 108, 108, 15: 108},
		{12: 798, 388, 98: 792, 126: 800, 157: 799},
		{244, 244, 9: 244, 244},
		// 495
		{243, 243, 402, 403, 9: 243, 243, 330, 243, 401, 400, 17: 332, 399, 344, 345, 333, 368, 24: 404, 405, 406, 407, 408, 410, 411, 793, 78: 409, 80: 794},
		{2: 346, 378, 382, 383, 384, 385, 386, 11: 330, 13: 417, 16: 315, 332, 399, 344, 345, 333, 368, 24: 334, 379, 348, 380, 335, 336, 337, 32: 311, 312, 313, 314, 38: 381, 41: 347, 353, 351, 350, 352, 339, 61: 415, 77: 797},
		{31: 795},
		{2: 346, 378, 382, 383, 384, 385, 386, 11: 330, 13: 417, 16: 315, 332, 399, 344, 345, 333, 368, 24: 334, 379, 348, 380, 335, 336, 337, 32: 311, 312, 313, 314, 38: 381, 41: 347, 353, 351, 350, 352, 339, 61: 415, 77: 796},
		{241, 241, 9: 241, 241, 12: 241, 14: 425},
		// 500
		{242, 242, 9: 242, 242, 12: 242, 14: 425},
		{246, 246, 9: 246, 246},
		{802, 12: 114, 114: 801},
		{240, 12: 240},
		{12: 804},
		// 505
		{12: 113, 388, 98: 792, 126: 803},
		{239, 12: 239},
		{245, 245, 9: 245, 245},
		{2: 402, 403, 11: 330, 13: 401, 17: 332, 399, 344, 345, 333, 368, 24: 404, 405, 406, 407, 408, 410, 411, 78: 409, 80: 806},
		{263, 263, 263, 9: 263, 263, 12: 263, 14: 263, 263, 263, 23: 263, 31: 263, 36: 263, 263, 39: 263, 263, 76: 263},
		// 510
		{265, 265, 265, 9: 265, 265, 12: 265, 14: 265, 265, 265, 23: 265, 31: 265, 36: 265, 265, 39: 265, 265, 76: 265},
		{2: 402, 403, 11: 330, 13: 401, 17: 332, 399, 344, 345, 333, 368, 24: 404, 405, 406, 407, 408, 410, 411, 78: 409, 80: 809},
		{264, 264, 264, 9: 264, 264, 12: 264, 14: 264, 264, 264, 23: 264, 31: 264, 36: 264, 264, 39: 264, 264, 76: 264},
		{2: 346, 378, 382, 383, 384, 385, 386, 11: 330, 13: 417, 16: 315, 332, 399, 344, 345, 333, 368, 24: 334, 379, 348, 380, 335, 336, 337, 32: 311, 312, 313, 314, 38: 381, 41: 347, 353, 351, 350, 352, 339, 61: 415, 77: 785},
		{2: 346, 378, 382, 383, 384, 385, 386, 11: 330, 13: 417, 16: 315, 332, 399, 344, 345, 333, 368, 24: 334, 379, 348, 380, 335, 336, 337, 32: 311, 312, 313, 314, 38: 381, 41: 347, 353, 351, 350, 352, 339, 61: 415, 77: 771},
		// 515
		{40: 818},
		{3: 429, 428, 430, 431, 435, 446, 40: 816, 47: 427, 432, 433, 434, 436, 437, 438, 439, 440, 441, 442, 443, 444, 445},
		{2: 402, 403, 11: 330, 13: 401, 17: 332, 399, 344, 345, 333, 368, 24: 404, 405, 406, 407, 408, 410, 411, 78: 409, 80: 815},
		{104, 104, 104, 9: 104, 104, 12: 104, 14: 104, 104, 104, 23: 104, 31: 104, 36: 104, 104, 39: 104, 104, 76: 104},
		{2: 402, 403, 11: 330, 13: 401, 17: 332, 399, 344, 345, 333, 368, 24: 404, 405, 406, 407, 408, 410, 411, 78: 409, 80: 817},
		// 520
		{290, 290, 290, 9: 290, 290, 12: 290, 14: 290, 290, 290, 23: 290, 31: 290, 36: 290, 290, 39: 290, 290, 76: 290},
		{2: 402, 403, 11: 330, 13: 401, 17: 332, 399, 344, 345, 333, 368, 24: 404, 405, 406, 407, 408, 410, 411, 78: 409, 80: 819},
		{291, 291, 291, 9: 291, 291, 12: 291, 14: 291, 291, 291, 23: 291, 31: 291, 36: 291, 291, 39: 291, 291, 76: 291},
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

			lhs.Value = newConstValue(newFloatConst(0, f, lx.float64Type, true).normalize(lx.ctx))
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

			lhs.Value = newConstValue(newComplexConst(0, &bigComplex{big.NewFloat(0).SetPrec(lx.floatConstPrec), f}, lx.float64Type, true).normalize(lx.ctx))
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

			if lhs.Value = newConstValue(newIntConst(0, i, lx.intType, true).normalize(lx.ctx)); lhs.Value == nil {
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
			if !lx.declarationScope.isMergeScope {
				lx.pushScope()
			}
			lx.declarationScope.isMergeScope = false
			if f := lx.lastFuncDeclaration; f != nil {
				lx.declarationScope.fnType = &f.Type
				lx.lastFuncDeclaration = nil
			}
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
			lx.lastFuncDeclaration = nil
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
				p.Scope.declare(lx, newFuncDeclaration(nm, nil, sig, false, lx.pkg.unsafe))
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
				d.(*TypeDeclaration).declare(lx, newFuncDeclaration(nm, r, sig, false, lx.pkg.unsafe))
			case *TypeDeclaration:
				x.declare(lx, newFuncDeclaration(nm, r, sig, false, lx.pkg.unsafe))
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
			lx.pushScope()
			lx.declarationScope.skip = true
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
			lhs.methods = lx.declarationScope
			lhs.pkgPath = lx.pkg.importPath
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
			lx.popScope()
			lx.declarationScope.declare(lx, newFuncDeclaration(lhs.Token, nil, lhs.Signature, true, false))
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
			yyVAL.node = &LBraceCompLitItem{
				Case:               4,
				LBraceCompLitValue: yyS[yypt-2].node.(*LBraceCompLitValue),
				Token:              yyS[yypt-1].Token,
				Expression:         yyS[yypt-0].node.(*Expression),
			}
		}
	case 138:
		{
			yyVAL.node = &LBraceCompLitItem{
				Case:                5,
				LBraceCompLitValue:  yyS[yypt-2].node.(*LBraceCompLitValue),
				Token:               yyS[yypt-1].Token,
				LBraceCompLitValue2: yyS[yypt-0].node.(*LBraceCompLitValue),
			}
		}
	case 139:
		{
			yyVAL.node = &LBraceCompLitItemList{
				LBraceCompLitItem: yyS[yypt-0].node.(*LBraceCompLitItem),
			}
		}
	case 140:
		{
			yyVAL.node = &LBraceCompLitItemList{
				Case: 1,
				LBraceCompLitItemList: yyS[yypt-2].node.(*LBraceCompLitItemList),
				Token:             yyS[yypt-1].Token,
				LBraceCompLitItem: yyS[yypt-0].node.(*LBraceCompLitItem),
			}
		}
	case 141:
		{
			yyVAL.node = &LBraceCompLitValue{
				LBrace: yyS[yypt-1].node.(*LBrace),
				Token:  yyS[yypt-0].Token,
			}
		}
	case 142:
		{
			yyVAL.node = &LBraceCompLitValue{
				Case:                  1,
				LBrace:                yyS[yypt-3].node.(*LBrace),
				LBraceCompLitItemList: yyS[yypt-2].node.(*LBraceCompLitItemList).reverse(),
				CommaOpt:              yyS[yypt-1].node.(*CommaOpt),
				Token:                 yyS[yypt-0].Token,
			}
		}
	case 143:
		{
			yyVAL.node = &MapType{
				Token:  yyS[yypt-4].Token,
				Token2: yyS[yypt-3].Token,
				Typ:    yyS[yypt-2].node.(*Typ),
				Token3: yyS[yypt-1].Token,
				Typ2:   yyS[yypt-0].node.(*Typ),
			}
		}
	case 144:
		{
			yyVAL.node = &Operand{
				Token:      yyS[yypt-2].Token,
				Expression: yyS[yypt-1].node.(*Expression),
				Token2:     yyS[yypt-0].Token,
			}
		}
	case 145:
		{
			yyVAL.node = &Operand{
				Case:        1,
				Token:       yyS[yypt-2].Token,
				TypeLiteral: yyS[yypt-1].node.(*TypeLiteral),
				Token2:      yyS[yypt-0].Token,
			}
		}
	case 146:
		{
			yyVAL.node = &Operand{
				Case:         2,
				BasicLiteral: yyS[yypt-0].node.(*BasicLiteral),
			}
		}
	case 147:
		{
			lx := yylex.(*lexer)
			lx.declarationScope.isMergeScope = false
			lx.declarationScope.fnType = &yyS[yypt-0].node.(*FuncType).Type
		}
	case 148:
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
	case 149:
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
	case 150:
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
	case 151:
		{
			yyVAL.node = &ParameterDecl{
				Token: yyS[yypt-1].Token,
				Typ:   yyS[yypt-0].node.(*Typ),
			}
		}
	case 152:
		{
			yyVAL.node = &ParameterDecl{
				Case:   1,
				Token:  yyS[yypt-2].Token,
				Token2: yyS[yypt-1].Token,
				Typ:    yyS[yypt-0].node.(*Typ),
			}
		}
	case 153:
		{
			yyVAL.node = &ParameterDecl{
				Case:  2,
				Token: yyS[yypt-1].Token,
				Typ:   yyS[yypt-0].node.(*Typ),
			}
		}
	case 154:
		{
			yyVAL.node = &ParameterDecl{
				Case: 3,
				Typ:  yyS[yypt-0].node.(*Typ),
			}
		}
	case 155:
		{
			yyVAL.node = &ParameterDeclList{
				ParameterDecl: yyS[yypt-0].node.(*ParameterDecl),
			}
		}
	case 156:
		{
			yyVAL.node = &ParameterDeclList{
				Case:              1,
				ParameterDeclList: yyS[yypt-2].node.(*ParameterDeclList),
				Token:             yyS[yypt-1].Token,
				ParameterDecl:     yyS[yypt-0].node.(*ParameterDecl),
			}
		}
	case 157:
		{
			yyVAL.node = &Parameters{
				Token:  yyS[yypt-1].Token,
				Token2: yyS[yypt-0].Token,
			}
		}
	case 158:
		{
			yyVAL.node = &Parameters{
				Case:              1,
				Token:             yyS[yypt-3].Token,
				ParameterDeclList: yyS[yypt-2].node.(*ParameterDeclList).reverse(),
				CommaOpt:          yyS[yypt-1].node.(*CommaOpt),
				Token2:            yyS[yypt-0].Token,
			}
		}
	case 159:
		{
			yyVAL.node = &PrimaryExpression{
				Operand: yyS[yypt-0].node.(*Operand),
			}
		}
	case 160:
		{
			lhs := &PrimaryExpression{
				Case:               1,
				CompLitType:        yyS[yypt-1].node.(*CompLitType),
				LBraceCompLitValue: yyS[yypt-0].node.(*LBraceCompLitValue),
			}
			yyVAL.node = lhs
			if a := lhs.CompLitType.ArrayType; a != nil {
				a.compLitValue = lhs.LBraceCompLitValue
			}
		}
	case 161:
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
	case 162:
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
	case 163:
		{
			yyVAL.node = &PrimaryExpression{
				Case:              4,
				PrimaryExpression: yyS[yypt-2].node.(*PrimaryExpression),
				Token:             yyS[yypt-1].Token,
				Token2:            yyS[yypt-0].Token,
			}
		}
	case 164:
		{
			yyVAL.node = &PrimaryExpression{
				Case:              5,
				PrimaryExpression: yyS[yypt-3].node.(*PrimaryExpression),
				Token:             yyS[yypt-2].Token,
				Expression:        yyS[yypt-1].node.(*Expression),
				Token2:            yyS[yypt-0].Token,
			}
		}
	case 165:
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
	case 166:
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
	case 167:
		{
			yyVAL.node = &PrimaryExpression{
				Case:              8,
				PrimaryExpression: yyS[yypt-1].node.(*PrimaryExpression),
				Call:              yyS[yypt-0].node.(*Call),
			}
		}
	case 168:
		{
			lx := yylex.(*lexer)
			lhs := &PrimaryExpression{
				Case:              9,
				PrimaryExpression: yyS[yypt-1].node.(*PrimaryExpression),
				CompLitValue:      yyS[yypt-0].node.(*CompLitValue),
			}
			yyVAL.node = lhs
			if o := lhs.PrimaryExpression.Operand; o != nil {
				switch o.Case {
				case
					0, // '(' Expression ')'
					1: // '(' TypeLiteral ')'
					lx.err(lhs, "syntax error: cannot parenthesize type in composite literal")
				}
			}
		}
	case 169:
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
	case 170:
		{
			lx := yylex.(*lexer)
			lhs := &Prologue{
				PackageClause: yyS[yypt-1].node.(*PackageClause),
				ImportList:    yyS[yypt-0].node.(*ImportList).reverse(),
			}
			yyVAL.node = lhs
			for _, v := range lx.imports {
				if v.ip == idC {
					return 0
				}
			}

			if lhs.post(lx) {
				return 0
			}
		}
	case 171:
		{
			lx := yylex.(*lexer)
			lhs := &QualifiedIdent{
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.fileScope = lx.fileScope
			lhs.resolutionScope = lx.resolutionScope
		}
	case 172:
		{
			lx := yylex.(*lexer)
			lhs := &QualifiedIdent{
				Case:   1,
				Token:  yyS[yypt-2].Token,
				Token2: yyS[yypt-1].Token,
				Token3: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.fileScope = lx.fileScope
			lhs.resolutionScope = lx.resolutionScope
		}
	case 173:
		{
			yyVAL.node = &Range{
				ExpressionList: yyS[yypt-3].node.(*ExpressionList).reverse(),
				Token:          yyS[yypt-2].Token,
				Token2:         yyS[yypt-1].Token,
				Expression:     yyS[yypt-0].node.(*Expression),
			}
		}
	case 174:
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
	case 175:
		{
			yyVAL.node = &Range{
				Case:       2,
				Token:      yyS[yypt-1].Token,
				Expression: yyS[yypt-0].node.(*Expression),
			}
		}
	case 176:
		{
			yyVAL.node = (*ReceiverOpt)(nil)
		}
	case 177:
		{
			lx := yylex.(*lexer)
			lhs := &ReceiverOpt{
				Parameters: yyS[yypt-0].node.(*Parameters),
			}
			yyVAL.node = lhs
			lhs.resolutionScope = lx.resolutionScope
			lhs.post(lx)
		}
	case 178:
		{
			yyVAL.node = (*ResultOpt)(nil)
		}
	case 179:
		{
			yyVAL.node = &ResultOpt{
				Case:       1,
				Parameters: yyS[yypt-0].node.(*Parameters),
			}
		}
	case 180:
		{
			yyVAL.node = &ResultOpt{
				Case: 2,
				Typ:  yyS[yypt-0].node.(*Typ),
			}
		}
	case 181:
		{
			lx := yylex.(*lexer)
			lx.switchDecl = append(lx.switchDecl, xc.Token{})
		}
	case 182:
		{
			lx := yylex.(*lexer)
			yyVAL.node = &SelectStatement{
				Token:      yyS[yypt-2].Token,
				SwitchBody: yyS[yypt-0].node.(*SwitchBody),
			}
			lx.switchDecl = lx.switchDecl[:len(lx.switchDecl)-1]
		}
	case 183:
		{
			yyVAL.node = (*SemicolonOpt)(nil)
		}
	case 184:
		{
			yyVAL.node = &SemicolonOpt{
				Token: yyS[yypt-0].Token,
			}
		}
	case 185:
		{
			yyVAL.node = &Signature{
				Parameters: yyS[yypt-1].node.(*Parameters),
				ResultOpt:  yyS[yypt-0].node.(*ResultOpt),
			}
		}
	case 186:
		{
			yyVAL.node = &SimpleStatement{
				Assignment: yyS[yypt-0].node.(*Assignment),
			}
		}
	case 187:
		{
			yyVAL.node = &SimpleStatement{
				Case:       1,
				Expression: yyS[yypt-0].node.(*Expression),
			}
		}
	case 188:
		{
			yyVAL.node = &SimpleStatement{
				Case:       2,
				Expression: yyS[yypt-1].node.(*Expression),
				Token:      yyS[yypt-0].Token,
			}
		}
	case 189:
		{
			yyVAL.node = &SimpleStatement{
				Case:       3,
				Expression: yyS[yypt-1].node.(*Expression),
				Token:      yyS[yypt-0].Token,
			}
		}
	case 190:
		{
			lx := yylex.(*lexer)
			lhs := &SimpleStatement{
				Case:            4,
				ExpressionList:  yyS[yypt-2].node.(*ExpressionList).reverse(),
				Token:           yyS[yypt-1].Token,
				ExpressionList2: yyS[yypt-0].node.(*ExpressionList).reverse(),
			}
			yyVAL.node = lhs
			lhs.resolutionScope = lx.resolutionScope
			lhs.idlist = varDecl(lx, lhs.ExpressionList, lhs.ExpressionList2, nil, ":=", -1, -1)
		}
	case 191:
		{
			yyVAL.node = (*SimpleStatementOpt)(nil)
		}
	case 192:
		{
			yyVAL.node = &SimpleStatementOpt{
				SimpleStatement: yyS[yypt-0].node.(*SimpleStatement),
			}
		}
	case 193:
		{
			yyVAL.node = &SliceType{
				Token:  yyS[yypt-2].Token,
				Token2: yyS[yypt-1].Token,
				Typ:    yyS[yypt-0].node.(*Typ),
			}
		}
	case 194:
		{
			yyVAL.node = (*Statement)(nil)
		}
	case 195:
		{
			yyVAL.node = &Statement{
				Case:  1,
				Block: yyS[yypt-0].node.(*Block),
			}
		}
	case 196:
		{
			yyVAL.node = &Statement{
				Case:      2,
				ConstDecl: yyS[yypt-0].node.(*ConstDecl),
			}
		}
	case 197:
		{
			yyVAL.node = &Statement{
				Case:     3,
				TypeDecl: yyS[yypt-0].node.(*TypeDecl),
			}
		}
	case 198:
		{
			yyVAL.node = &Statement{
				Case:    4,
				VarDecl: yyS[yypt-0].node.(*VarDecl),
			}
		}
	case 199:
		{
			yyVAL.node = &Statement{
				Case:             5,
				StatementNonDecl: yyS[yypt-0].node.(*StatementNonDecl),
			}
		}
	case 200:
		{
			yyVAL.node = &Statement{
				Case: 6,
			}
		}
	case 201:
		{
			yyVAL.node = &StatementList{
				Statement: yyS[yypt-0].node.(*Statement),
			}
		}
	case 202:
		{
			yyVAL.node = &StatementList{
				Case:          1,
				StatementList: yyS[yypt-2].node.(*StatementList),
				Token:         yyS[yypt-1].Token,
				Statement:     yyS[yypt-0].node.(*Statement),
			}
		}
	case 203:
		{
			lx := yylex.(*lexer)
			lhs := &StatementNonDecl{
				Token:         yyS[yypt-1].Token,
				IdentifierOpt: yyS[yypt-0].node.(*IdentifierOpt),
			}
			yyVAL.node = lhs
			lhs.resolutionScope = lx.resolutionScope
		}
	case 204:
		{
			lx := yylex.(*lexer)
			lhs := &StatementNonDecl{
				Case:          1,
				Token:         yyS[yypt-1].Token,
				IdentifierOpt: yyS[yypt-0].node.(*IdentifierOpt),
			}
			yyVAL.node = lhs
			lhs.resolutionScope = lx.resolutionScope
		}
	case 205:
		{
			yyVAL.node = &StatementNonDecl{
				Case:       2,
				Token:      yyS[yypt-1].Token,
				Expression: yyS[yypt-0].node.(*Expression),
			}
		}
	case 206:
		{
			yyVAL.node = &StatementNonDecl{
				Case:  3,
				Token: yyS[yypt-0].Token,
			}
		}
	case 207:
		{
			yyVAL.node = &StatementNonDecl{
				Case:         4,
				ForStatement: yyS[yypt-0].node.(*ForStatement),
			}
		}
	case 208:
		{
			yyVAL.node = &StatementNonDecl{
				Case:       5,
				Token:      yyS[yypt-1].Token,
				Expression: yyS[yypt-0].node.(*Expression),
			}
		}
	case 209:
		{
			lx := yylex.(*lexer)
			lhs := &StatementNonDecl{
				Case:   6,
				Token:  yyS[yypt-1].Token,
				Token2: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.resolutionScope = lx.resolutionScope
		}
	case 210:
		{
			lx := yylex.(*lexer)
			lhs := &StatementNonDecl{
				Case:      7,
				Token:     yyS[yypt-2].Token,
				Token2:    yyS[yypt-1].Token,
				Statement: yyS[yypt-0].node.(*Statement),
			}
			yyVAL.node = lhs
			lx.declarationScope.declare(lx, newLabelDeclaration(lhs.Token))
		}
	case 211:
		{
			yyVAL.node = &StatementNonDecl{
				Case:        8,
				IfStatement: yyS[yypt-0].node.(*IfStatement),
			}
		}
	case 212:
		{
			lx := yylex.(*lexer)
			lhs := &StatementNonDecl{
				Case:              9,
				Token:             yyS[yypt-1].Token,
				ExpressionListOpt: yyS[yypt-0].node.(*ExpressionListOpt),
			}
			yyVAL.node = lhs
			lhs.resolutionScope = lx.resolutionScope
		}
	case 213:
		{
			yyVAL.node = &StatementNonDecl{
				Case:            10,
				SelectStatement: yyS[yypt-0].node.(*SelectStatement),
			}
		}
	case 214:
		{
			yyVAL.node = &StatementNonDecl{
				Case:            11,
				SimpleStatement: yyS[yypt-0].node.(*SimpleStatement),
			}
		}
	case 215:
		{
			yyVAL.node = &StatementNonDecl{
				Case:            12,
				SwitchStatement: yyS[yypt-0].node.(*SwitchStatement),
			}
		}
	case 216:
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
	case 217:
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
	case 218:
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
	case 219:
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
	case 220:
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
	case 221:
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
	case 222:
		{
			yyVAL.node = &StructFieldDeclList{
				StructFieldDecl: yyS[yypt-0].node.(*StructFieldDecl),
			}
		}
	case 223:
		{
			yyVAL.node = &StructFieldDeclList{
				Case:                1,
				StructFieldDeclList: yyS[yypt-2].node.(*StructFieldDeclList),
				Token:               yyS[yypt-1].Token,
				StructFieldDecl:     yyS[yypt-0].node.(*StructFieldDecl),
			}
		}
	case 224:
		{
			yyVAL.node = &StructType{
				Token:  yyS[yypt-2].Token,
				LBrace: yyS[yypt-1].node.(*LBrace),
				Token2: yyS[yypt-0].Token,
			}
		}
	case 225:
		{
			lx := yylex.(*lexer)
			lx.pushScope()
			lx.declarationScope.skip = true
		}
	case 226:
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
			lhs.fields = lx.declarationScope
			for _, v := range lhs.fields.Bindings {
				v.(*FieldDeclaration).parent = lhs
			}
			lhs.pkgPath = lx.pkg.importPath
			lx.popScope()
		}
	case 227:
		{
			yyVAL.node = &SwitchBody{
				Token:  yyS[yypt-1].Token,
				Token2: yyS[yypt-0].Token,
			}
		}
	case 228:
		{
			lx := yylex.(*lexer)
			lx.pushScope()
		}
	case 229:
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
	case 230:
		{
			yyVAL.node = &SwitchCase{
				Token:        yyS[yypt-2].Token,
				ArgumentList: yyS[yypt-1].node.(*ArgumentList).reverse(),
				Token2:       yyS[yypt-0].Token,
			}
		}
	case 231:
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
	case 232:
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
	case 233:
		{
			yyVAL.node = &SwitchCase{
				Case:   3,
				Token:  yyS[yypt-1].Token,
				Token2: yyS[yypt-0].Token,
			}
		}
	case 234:
		{
			yyVAL.node = &SwitchCase{
				Case:  4,
				Token: yyS[yypt-1].Token,
			}
		}
	case 235:
		{
			yyVAL.node = &SwitchCase{
				Case:  5,
				Token: yyS[yypt-1].Token,
			}
		}
	case 236:
		{
			lx := yylex.(*lexer)
			yyVAL.node = &SwitchCaseBlock{
				SwitchCase:    yyS[yypt-1].node.(*SwitchCase),
				StatementList: yyS[yypt-0].node.(*StatementList).reverse(),
			}
			lx.popScope() // Implicit "case"/"default" block.
		}
	case 237:
		{
			yyVAL.node = &SwitchCaseList{
				SwitchCaseBlock: yyS[yypt-0].node.(*SwitchCaseBlock),
			}
		}
	case 238:
		{
			yyVAL.node = &SwitchCaseList{
				Case:            1,
				SwitchCaseList:  yyS[yypt-1].node.(*SwitchCaseList),
				SwitchCaseBlock: yyS[yypt-0].node.(*SwitchCaseBlock),
			}
		}
	case 239:
		{
			yyVAL.node = &SwitchHeader{
				SimpleStatementOpt: yyS[yypt-0].node.(*SimpleStatementOpt),
			}
		}
	case 240:
		{
			yyVAL.node = &SwitchHeader{
				Case:               1,
				SimpleStatementOpt: yyS[yypt-1].node.(*SimpleStatementOpt),
				Token:              yyS[yypt-0].Token,
			}
		}
	case 241:
		{
			yyVAL.node = &SwitchHeader{
				Case:               2,
				SimpleStatementOpt: yyS[yypt-2].node.(*SimpleStatementOpt),
				Token:              yyS[yypt-1].Token,
				Expression:         yyS[yypt-0].node.(*Expression),
			}
		}
	case 242:
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
	case 243:
		{
			lx := yylex.(*lexer)
			var t xc.Token
			if sh := yyS[yypt-0].node.(*SwitchHeader); sh.Case == 3 { // SimpleStatementOpt ';' IDENTIFIER ":=" PrimaryExpression '.' '(' "type" ')'
				t = sh.Token2 // IDENTIFIER
			}
			lx.switchDecl = append(lx.switchDecl, t)
		}
	case 244:
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
	case 245:
		{
			yyVAL.node = (*TagOpt)(nil)
		}
	case 246:
		{
			lx := yylex.(*lexer)
			lhs := &TagOpt{
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.stringValue = stringLiteralValue(lx, lhs.Token)
		}
	case 247:
		{
			yyVAL.node = &TopLevelDecl{
				ConstDecl: yyS[yypt-0].node.(*ConstDecl),
			}
		}
	case 248:
		{
			yyVAL.node = &TopLevelDecl{
				Case:     1,
				FuncDecl: yyS[yypt-0].node.(*FuncDecl),
			}
		}
	case 249:
		{
			yyVAL.node = &TopLevelDecl{
				Case:     2,
				TypeDecl: yyS[yypt-0].node.(*TypeDecl),
			}
		}
	case 250:
		{
			yyVAL.node = &TopLevelDecl{
				Case:    3,
				VarDecl: yyS[yypt-0].node.(*VarDecl),
			}
		}
	case 251:
		{
			yyVAL.node = &TopLevelDecl{
				Case:             4,
				StatementNonDecl: yyS[yypt-0].node.(*StatementNonDecl),
			}
		}
	case 252:
		{
			yyVAL.node = &TopLevelDecl{
				Case: 5,
			}
		}
	case 253:
		{
			yyVAL.node = (*TopLevelDeclList)(nil)
		}
	case 254:
		{
			yyVAL.node = &TopLevelDeclList{
				TopLevelDeclList: yyS[yypt-2].node.(*TopLevelDeclList),
				TopLevelDecl:     yyS[yypt-1].node.(*TopLevelDecl),
				Token:            yyS[yypt-0].Token,
			}
		}
	case 255:
		{
			yyVAL.node = &Typ{
				Token:  yyS[yypt-2].Token,
				Typ:    yyS[yypt-1].node.(*Typ),
				Token2: yyS[yypt-0].Token,
			}
		}
	case 256:
		{
			yyVAL.node = &Typ{
				Case:  1,
				Token: yyS[yypt-1].Token,
				Typ:   yyS[yypt-0].node.(*Typ),
			}
		}
	case 257:
		{
			yyVAL.node = &Typ{
				Case:      2,
				ArrayType: yyS[yypt-0].node.(*ArrayType),
			}
		}
	case 258:
		{
			yyVAL.node = &Typ{
				Case:     3,
				ChanType: yyS[yypt-0].node.(*ChanType),
			}
		}
	case 259:
		{
			lx := yylex.(*lexer)
			yyVAL.node = &Typ{
				Case:     4,
				FuncType: yyS[yypt-0].node.(*FuncType),
			}
			lx.popScope()
		}
	case 260:
		{
			yyVAL.node = &Typ{
				Case:          5,
				InterfaceType: yyS[yypt-0].node.(*InterfaceType),
			}
		}
	case 261:
		{
			yyVAL.node = &Typ{
				Case:    6,
				MapType: yyS[yypt-0].node.(*MapType),
			}
		}
	case 262:
		{
			yyVAL.node = &Typ{
				Case:                7,
				QualifiedIdent:      yyS[yypt-1].node.(*QualifiedIdent),
				GenericArgumentsOpt: yyS[yypt-0].node.(*GenericArgumentsOpt),
			}
		}
	case 263:
		{
			yyVAL.node = &Typ{
				Case:      8,
				SliceType: yyS[yypt-0].node.(*SliceType),
			}
		}
	case 264:
		{
			yyVAL.node = &Typ{
				Case:       9,
				StructType: yyS[yypt-0].node.(*StructType),
			}
		}
	case 265:
		{
			yyVAL.node = &TypeDecl{
				Token:  yyS[yypt-2].Token,
				Token2: yyS[yypt-1].Token,
				Token3: yyS[yypt-0].Token,
			}
		}
	case 266:
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
	case 267:
		{
			yyVAL.node = &TypeDecl{
				Case:     2,
				Token:    yyS[yypt-1].Token,
				TypeSpec: yyS[yypt-0].node.(*TypeSpec),
			}
		}
	case 268:
		{
			yyVAL.node = &TypeList{
				Typ: yyS[yypt-0].node.(*Typ),
			}
		}
	case 269:
		{
			yyVAL.node = &TypeList{
				Case:     1,
				TypeList: yyS[yypt-2].node.(*TypeList),
				Token:    yyS[yypt-1].Token,
				Typ:      yyS[yypt-0].node.(*Typ),
			}
		}
	case 270:
		{
			yyVAL.node = &TypeLiteral{
				Token:       yyS[yypt-1].Token,
				TypeLiteral: yyS[yypt-0].node.(*TypeLiteral),
			}
		}
	case 271:
		{
			yyVAL.node = &TypeLiteral{
				Case:      1,
				ArrayType: yyS[yypt-0].node.(*ArrayType),
			}
		}
	case 272:
		{
			yyVAL.node = &TypeLiteral{
				Case:     2,
				ChanType: yyS[yypt-0].node.(*ChanType),
			}
		}
	case 273:
		{
			lx := yylex.(*lexer)
			yyVAL.node = &TypeLiteral{
				Case:     3,
				FuncType: yyS[yypt-0].node.(*FuncType),
			}
			lx.popScope()
		}
	case 274:
		{
			yyVAL.node = &TypeLiteral{
				Case:          4,
				InterfaceType: yyS[yypt-0].node.(*InterfaceType),
			}
		}
	case 275:
		{
			yyVAL.node = &TypeLiteral{
				Case:    5,
				MapType: yyS[yypt-0].node.(*MapType),
			}
		}
	case 276:
		{
			yyVAL.node = &TypeLiteral{
				Case:      6,
				SliceType: yyS[yypt-0].node.(*SliceType),
			}
		}
	case 277:
		{
			yyVAL.node = &TypeLiteral{
				Case:       7,
				StructType: yyS[yypt-0].node.(*StructType),
			}
		}
	case 278:
		{
			lx := yylex.(*lexer)
			lhs := &TypeSpec{
				Token:                yyS[yypt-2].Token,
				GenericParametersOpt: yyS[yypt-1].node.(*GenericParametersOpt),
				Typ:                  yyS[yypt-0].node.(*Typ),
			}
			yyVAL.node = lhs
			nm := lhs.Token
			if lx.declarationScope.Kind == PackageScope {
				p := lx.pkg
				if d := p.forwardTypes[nm.Val]; d != nil {
					delete(p.forwardTypes, nm.Val)
					d.(*TypeDeclaration).typ0 = lhs.Typ
					lx.declarationScope.declare(lx, d)
					break
				}
			}

			d := newTypeDeclaration(lx, nm, lhs.Typ)
			if t := lhs.Typ; t.Case == 5 { // InterfaceType
				d.methods = t.InterfaceType.methods
			}
			lx.declarationScope.declare(lx, d)
		}
	case 279:
		{
			yyVAL.node = &TypeSpecList{
				TypeSpec: yyS[yypt-0].node.(*TypeSpec),
			}
		}
	case 280:
		{
			yyVAL.node = &TypeSpecList{
				Case:         1,
				TypeSpecList: yyS[yypt-2].node.(*TypeSpecList),
				Token:        yyS[yypt-1].Token,
				TypeSpec:     yyS[yypt-0].node.(*TypeSpec),
			}
		}
	case 281:
		{
			yyVAL.node = &UnaryExpression{
				Token:           yyS[yypt-1].Token,
				UnaryExpression: yyS[yypt-0].node.(*UnaryExpression),
			}
		}
	case 282:
		{
			yyVAL.node = &UnaryExpression{
				Case:            1,
				Token:           yyS[yypt-1].Token,
				UnaryExpression: yyS[yypt-0].node.(*UnaryExpression),
			}
		}
	case 283:
		{
			yyVAL.node = &UnaryExpression{
				Case:            2,
				Token:           yyS[yypt-1].Token,
				UnaryExpression: yyS[yypt-0].node.(*UnaryExpression),
			}
		}
	case 284:
		{
			yyVAL.node = &UnaryExpression{
				Case:            3,
				Token:           yyS[yypt-1].Token,
				UnaryExpression: yyS[yypt-0].node.(*UnaryExpression),
			}
		}
	case 285:
		{
			yyVAL.node = &UnaryExpression{
				Case:            4,
				Token:           yyS[yypt-1].Token,
				UnaryExpression: yyS[yypt-0].node.(*UnaryExpression),
			}
		}
	case 286:
		{
			yyVAL.node = &UnaryExpression{
				Case:            5,
				Token:           yyS[yypt-1].Token,
				UnaryExpression: yyS[yypt-0].node.(*UnaryExpression),
			}
		}
	case 287:
		{
			yyVAL.node = &UnaryExpression{
				Case:            6,
				Token:           yyS[yypt-1].Token,
				UnaryExpression: yyS[yypt-0].node.(*UnaryExpression),
			}
		}
	case 288:
		{
			yyVAL.node = &UnaryExpression{
				Case:              7,
				PrimaryExpression: yyS[yypt-0].node.(*PrimaryExpression),
			}
		}
	case 289:
		{
			yyVAL.node = &VarDecl{
				Token:  yyS[yypt-2].Token,
				Token2: yyS[yypt-1].Token,
				Token3: yyS[yypt-0].Token,
			}
		}
	case 290:
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
	case 291:
		{
			yyVAL.node = &VarDecl{
				Case:    2,
				Token:   yyS[yypt-1].Token,
				VarSpec: yyS[yypt-0].node.(*VarSpec),
			}
		}
	case 292:
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
	case 293:
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
	case 294:
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
	case 295:
		{
			yyVAL.node = &VarSpecList{
				VarSpec: yyS[yypt-0].node.(*VarSpec),
			}
		}
	case 296:
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
