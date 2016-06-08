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
	"strconv"
	"strings"

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
	yyDefault     = 57412
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
	yyTabOfs   = -290
)

var (
	yyXLAT = map[int]int{
		59:    0,   // ';' (251x)
		125:   1,   // '}' (241x)
		40:    2,   // '(' (220x)
		42:    3,   // '*' (219x)
		38:    4,   // '&' (183x)
		43:    5,   // '+' (183x)
		45:    6,   // '-' (183x)
		94:    7,   // '^' (183x)
		57358: 8,   // COMM (183x)
		57354: 9,   // CASE (182x)
		57363: 10,  // DEFAULT (182x)
		41:    11,  // ')' (172x)
		44:    12,  // ',' (170x)
		91:    13,  // '[' (168x)
		57377: 14,  // IDENTIFIER (167x)
		57403: 15,  // STRING_LIT (147x)
		57352: 16,  // BODY (145x)
		57355: 17,  // CHAN (135x)
		57372: 18,  // FUNC (134x)
		57382: 19,  // INTERFACE (134x)
		57388: 20,  // MAP (134x)
		57401: 21,  // RXCHAN (134x)
		57404: 22,  // STRUCT (134x)
		58:    23,  // ':' (117x)
		61:    24,  // '=' (115x)
		57415: 25,  // ArrayType (111x)
		57422: 26,  // ChanType (111x)
		57444: 27,  // Function (111x)
		57443: 28,  // FuncType (111x)
		57464: 29,  // InterfaceType (111x)
		57469: 30,  // MapType (111x)
		57486: 31,  // SliceType (111x)
		57493: 32,  // StructType (111x)
		57357: 33,  // COLAS (109x)
		57361: 34,  // DDD (109x)
		57356: 35,  // CHAR_LIT (107x)
		57370: 36,  // FLOAT_LIT (107x)
		57379: 37,  // IMAG_LIT (107x)
		57383: 38,  // INT_LIT (107x)
		33:    39,  // '!' (101x)
		93:    40,  // ']' (101x)
		123:   41,  // '{' (97x)
		57417: 42,  // BasicLiteral (88x)
		57426: 43,  // CompLitType (86x)
		57470: 44,  // Operand (86x)
		57475: 45,  // PrimaryExpression (86x)
		57503: 46,  // TypeLiteral (86x)
		57506: 47,  // UnaryExpression (86x)
		37:    48,  // '%' (82x)
		47:    49,  // '/' (82x)
		60:    50,  // '<' (82x)
		62:    51,  // '>' (82x)
		124:   52,  // '|' (82x)
		57347: 53,  // ANDAND (82x)
		57348: 54,  // ANDNOT (82x)
		57367: 55,  // EQ (82x)
		57373: 56,  // GEQ (82x)
		57384: 57,  // LEQ (82x)
		57385: 58,  // LSH (82x)
		57391: 59,  // NEQ (82x)
		57393: 60,  // OROR (82x)
		57399: 61,  // RSH (82x)
		57434: 62,  // Expression (79x)
		57346: 63,  // ADD_ASSIGN (65x)
		57350: 64,  // AND_ASSIGN (65x)
		57349: 65,  // ANDNOT_ASSIGN (65x)
		57365: 66,  // DIV_ASSIGN (65x)
		57386: 67,  // LSH_ASSIGN (65x)
		57389: 68,  // MOD_ASSIGN (65x)
		57390: 69,  // MUL_ASSIGN (65x)
		57394: 70,  // OR_ASSIGN (65x)
		57400: 71,  // RSH_ASSIGN (65x)
		57405: 72,  // SUB_ASSIGN (65x)
		57410: 73,  // XOR_ASSIGN (65x)
		57362: 74,  // DEC (62x)
		57381: 75,  // INC (62x)
		57376: 76,  // GTGT (41x)
		46:    77,  // '.' (40x)
		57435: 78,  // ExpressionList (34x)
		57477: 79,  // QualifiedIdent (33x)
		57345: 80,  // error (30x)
		57501: 81,  // Typ (25x)
		57378: 82,  // IF (22x)
		57408: 83,  // TYPE (22x)
		57353: 84,  // BREAK (21x)
		57359: 85,  // CONST (21x)
		57360: 86,  // CONTINUE (21x)
		57364: 87,  // DEFER (21x)
		57369: 88,  // FALLTHROUGH (21x)
		57371: 89,  // FOR (21x)
		57374: 90,  // GO (21x)
		57375: 91,  // GOTO (21x)
		57398: 92,  // RETURN (21x)
		57402: 93,  // SELECT (21x)
		57406: 94,  // SWITCH (21x)
		57409: 95,  // VAR (21x)
		57416: 96,  // Assignment (14x)
		57484: 97,  // SimpleStatement (14x)
		57344: 98,  // $end (8x)
		57418: 99,  // Block (8x)
		57453: 100, // IdentifierList (8x)
		57455: 101, // If (8x)
		57387: 102, // LTLT (8x)
		57428: 103, // ConstDecl (7x)
		57440: 104, // ForStatement (7x)
		57457: 105, // IfStatement (7x)
		57465: 106, // LBrace (7x)
		57481: 107, // SelectStatement (7x)
		57485: 108, // SimpleStatementOpt (7x)
		57489: 109, // StatementNonDecl (7x)
		57498: 110, // SwitchStatement (7x)
		57502: 111, // TypeDecl (7x)
		57507: 112, // VarDecl (7x)
		57423: 113, // CommaOpt (6x)
		57474: 114, // Parameters (6x)
		57482: 115, // SemicolonOpt (6x)
		57487: 116, // Statement (6x)
		57490: 117, // StringLitOpt (6x)
		57427: 118, // CompLitValue (5x)
		57366: 119, // ELSE (5x)
		57454: 120, // IdentifierOpt (5x)
		57413: 121, // Argument (4x)
		57447: 122, // GenericArguments (4x)
		57448: 123, // GenericArgumentsOpt (4x)
		57380: 124, // IMPORT (4x)
		57468: 125, // LBraceCompLitValue (4x)
		57483: 126, // Signature (4x)
		57488: 127, // StatementList (4x)
		57420: 128, // Body (3x)
		57429: 129, // ConstSpec (3x)
		57437: 130, // ExpressionOpt (3x)
		57456: 131, // IfHeader (3x)
		57460: 132, // ImportSpec (3x)
		57472: 133, // ParameterDecl (3x)
		57397: 134, // RANGE (3x)
		57504: 135, // TypeSpec (3x)
		57508: 136, // VarSpec (3x)
		57414: 137, // ArgumentList (2x)
		57424: 138, // CompLitItem (2x)
		57446: 139, // GenericArgumentListItem (2x)
		57450: 140, // GenericParameterListItem (2x)
		57462: 141, // InterfaceMethodDecl (2x)
		57466: 142, // LBraceCompLitItem (2x)
		57473: 143, // ParameterDeclList (2x)
		57480: 144, // ResultOpt (2x)
		57491: 145, // StructFieldDecl (2x)
		57494: 146, // SwitchBody (2x)
		57495: 147, // SwitchCase (2x)
		57496: 148, // SwitchCaseBlock (2x)
		57411: 149, // $@1 (1x)
		57419: 150, // BlockOpt (1x)
		57421: 151, // Call (1x)
		57425: 152, // CompLitItemList (1x)
		57430: 153, // ConstSpecList (1x)
		57431: 154, // Elif (1x)
		57432: 155, // ElifList (1x)
		57433: 156, // ElseOpt (1x)
		57436: 157, // ExpressionListOpt (1x)
		57438: 158, // File (1x)
		57439: 159, // ForHeader (1x)
		57441: 160, // FuncDecl (1x)
		57442: 161, // FuncOrMethod (1x)
		57445: 162, // GenericArgumentList (1x)
		57449: 163, // GenericParameterList (1x)
		57451: 164, // GenericParams (1x)
		57452: 165, // GenericParamsOpt (1x)
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
		57492: 176, // StructFieldDeclList (1x)
		57497: 177, // SwitchCaseList (1x)
		57499: 178, // TopLevelDecl (1x)
		57500: 179, // TopLevelDeclList (1x)
		57407: 180, // TXCHAN (1x)
		57505: 181, // TypeSpecList (1x)
		57509: 182, // VarSpecList (1x)
		57412: 183, // $default (0x)
		57351: 184, // BAD_FLOAT_LIT (0x)
		57368: 185, // ERRCHECK (0x)
		57392: 186, // NO_RESULT (0x)
		57396: 187, // PARAMS (0x)
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
		"')'",
		"','",
		"'['",
		"IDENTIFIER",
		"STRING_LIT",
		"BODY",
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
		"Function",
		"FuncType",
		"InterfaceType",
		"MapType",
		"SliceType",
		"StructType",
		"COLAS",
		"DDD",
		"CHAR_LIT",
		"FLOAT_LIT",
		"IMAG_LIT",
		"INT_LIT",
		"'!'",
		"']'",
		"'{'",
		"BasicLiteral",
		"CompLitType",
		"Operand",
		"PrimaryExpression",
		"TypeLiteral",
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
		"GTGT",
		"'.'",
		"ExpressionList",
		"QualifiedIdent",
		"error",
		"Typ",
		"IF",
		"TYPE",
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
		"$end",
		"Block",
		"IdentifierList",
		"If",
		"LTLT",
		"ConstDecl",
		"ForStatement",
		"IfStatement",
		"LBrace",
		"SelectStatement",
		"SimpleStatementOpt",
		"StatementNonDecl",
		"SwitchStatement",
		"TypeDecl",
		"VarDecl",
		"CommaOpt",
		"Parameters",
		"SemicolonOpt",
		"Statement",
		"StringLitOpt",
		"CompLitValue",
		"ELSE",
		"IdentifierOpt",
		"Argument",
		"GenericArguments",
		"GenericArgumentsOpt",
		"IMPORT",
		"LBraceCompLitValue",
		"Signature",
		"StatementList",
		"Body",
		"ConstSpec",
		"ExpressionOpt",
		"IfHeader",
		"ImportSpec",
		"ParameterDecl",
		"RANGE",
		"TypeSpec",
		"VarSpec",
		"ArgumentList",
		"CompLitItem",
		"GenericArgumentListItem",
		"GenericParameterListItem",
		"InterfaceMethodDecl",
		"LBraceCompLitItem",
		"ParameterDeclList",
		"ResultOpt",
		"StructFieldDecl",
		"SwitchBody",
		"SwitchCase",
		"SwitchCaseBlock",
		"$@1",
		"BlockOpt",
		"Call",
		"CompLitItemList",
		"ConstSpecList",
		"Elif",
		"ElifList",
		"ElseOpt",
		"ExpressionListOpt",
		"File",
		"ForHeader",
		"FuncDecl",
		"FuncOrMethod",
		"GenericArgumentList",
		"GenericParameterList",
		"GenericParams",
		"GenericParamsOpt",
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
		"TopLevelDecl",
		"TopLevelDeclList",
		"TXCHAN",
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
		57403: "literal",
		57352: "{",
		57355: "chan",
		57372: "func",
		57382: "interface",
		57388: "map",
		57401: "<-",
		57404: "struct",
		57357: ":=",
		57361: "...",
		57356: "literal",
		57370: "literal",
		57379: "literal",
		57383: "literal",
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
		57378: "if",
		57408: "type",
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
		1:   {158, 2},
		2:   {121, 1},
		3:   {121, 1},
		4:   {137, 1},
		5:   {137, 3},
		6:   {25, 4},
		7:   {25, 4},
		8:   {96, 3},
		9:   {96, 3},
		10:  {96, 3},
		11:  {96, 3},
		12:  {96, 3},
		13:  {96, 3},
		14:  {96, 3},
		15:  {96, 3},
		16:  {96, 3},
		17:  {96, 3},
		18:  {96, 3},
		19:  {96, 3},
		20:  {42, 1},
		21:  {42, 1},
		22:  {42, 1},
		23:  {42, 1},
		24:  {42, 1},
		25:  {99, 3},
		26:  {150, 0},
		27:  {150, 1},
		28:  {128, 3},
		29:  {151, 2},
		30:  {151, 4},
		31:  {151, 5},
		32:  {26, 2},
		33:  {26, 3},
		34:  {26, 3},
		35:  {113, 0},
		36:  {113, 1},
		37:  {138, 1},
		38:  {138, 3},
		39:  {138, 3},
		40:  {138, 1},
		41:  {138, 3},
		42:  {138, 3},
		43:  {152, 1},
		44:  {152, 3},
		45:  {43, 1},
		46:  {43, 1},
		47:  {43, 1},
		48:  {43, 1},
		49:  {118, 2},
		50:  {118, 4},
		51:  {103, 3},
		52:  {103, 5},
		53:  {103, 2},
		54:  {129, 1},
		55:  {129, 3},
		56:  {129, 4},
		57:  {153, 1},
		58:  {153, 3},
		59:  {154, 4},
		60:  {155, 0},
		61:  {155, 2},
		62:  {156, 0},
		63:  {156, 2},
		64:  {62, 1},
		65:  {62, 3},
		66:  {62, 3},
		67:  {62, 3},
		68:  {62, 3},
		69:  {62, 3},
		70:  {62, 3},
		71:  {62, 3},
		72:  {62, 3},
		73:  {62, 3},
		74:  {62, 3},
		75:  {62, 3},
		76:  {62, 3},
		77:  {62, 3},
		78:  {62, 3},
		79:  {62, 3},
		80:  {62, 3},
		81:  {62, 3},
		82:  {62, 3},
		83:  {62, 3},
		84:  {62, 3},
		85:  {130, 0},
		86:  {130, 1},
		87:  {78, 1},
		88:  {78, 3},
		89:  {157, 0},
		90:  {157, 1},
		91:  {159, 1},
		92:  {159, 5},
		93:  {159, 1},
		94:  {104, 3},
		95:  {160, 2},
		96:  {161, 5},
		97:  {28, 2},
		98:  {27, 1},
		99:  {162, 1},
		100: {162, 3},
		101: {139, 1},
		102: {122, 3},
		103: {123, 0},
		104: {123, 1},
		105: {163, 1},
		106: {163, 3},
		107: {140, 1},
		108: {164, 3},
		109: {165, 0},
		110: {165, 1},
		111: {120, 0},
		112: {120, 1},
		113: {100, 1},
		114: {100, 3},
		115: {101, 1},
		116: {131, 1},
		117: {131, 3},
		118: {105, 5},
		119: {166, 3},
		120: {166, 5},
		121: {166, 2},
		122: {132, 2},
		123: {132, 2},
		124: {132, 3},
		125: {132, 3},
		126: {168, 1},
		127: {168, 3},
		128: {167, 0},
		129: {167, 3},
		130: {29, 3},
		131: {29, 5},
		132: {141, 2},
		133: {141, 1},
		134: {169, 1},
		135: {169, 3},
		136: {106, 1},
		137: {106, 1},
		138: {142, 1},
		139: {142, 3},
		140: {142, 3},
		141: {142, 1},
		142: {170, 1},
		143: {170, 3},
		144: {125, 2},
		145: {125, 4},
		146: {30, 5},
		147: {44, 3},
		148: {44, 3},
		149: {44, 1},
		150: {44, 4},
		151: {44, 2},
		152: {79, 1},
		153: {79, 3},
		154: {149, 0},
		155: {172, 4},
		156: {133, 2},
		157: {133, 3},
		158: {133, 2},
		159: {133, 1},
		160: {143, 1},
		161: {143, 3},
		162: {114, 2},
		163: {114, 4},
		164: {45, 1},
		165: {45, 2},
		166: {45, 5},
		167: {45, 5},
		168: {45, 3},
		169: {45, 4},
		170: {45, 8},
		171: {45, 6},
		172: {45, 2},
		173: {45, 2},
		174: {45, 5},
		175: {173, 2},
		176: {174, 4},
		177: {174, 4},
		178: {174, 2},
		179: {175, 0},
		180: {175, 1},
		181: {144, 0},
		182: {144, 1},
		183: {144, 1},
		184: {107, 2},
		185: {115, 0},
		186: {115, 1},
		187: {126, 2},
		188: {97, 1},
		189: {97, 1},
		190: {97, 2},
		191: {97, 2},
		192: {97, 3},
		193: {108, 0},
		194: {108, 1},
		195: {31, 3},
		196: {116, 0},
		197: {116, 1},
		198: {116, 1},
		199: {116, 1},
		200: {116, 1},
		201: {116, 1},
		202: {116, 1},
		203: {127, 1},
		204: {127, 3},
		205: {109, 2},
		206: {109, 2},
		207: {109, 2},
		208: {109, 1},
		209: {109, 1},
		210: {109, 2},
		211: {109, 2},
		212: {109, 3},
		213: {109, 1},
		214: {109, 2},
		215: {109, 1},
		216: {109, 1},
		217: {109, 1},
		218: {117, 0},
		219: {117, 1},
		220: {145, 3},
		221: {145, 3},
		222: {145, 2},
		223: {145, 4},
		224: {145, 5},
		225: {145, 5},
		226: {176, 1},
		227: {176, 3},
		228: {32, 3},
		229: {32, 5},
		230: {146, 2},
		231: {146, 3},
		232: {147, 3},
		233: {147, 5},
		234: {147, 5},
		235: {147, 2},
		236: {147, 2},
		237: {147, 2},
		238: {148, 2},
		239: {177, 1},
		240: {177, 2},
		241: {110, 3},
		242: {178, 1},
		243: {178, 1},
		244: {178, 1},
		245: {178, 1},
		246: {178, 1},
		247: {178, 1},
		248: {179, 0},
		249: {179, 3},
		250: {81, 3},
		251: {81, 2},
		252: {81, 1},
		253: {81, 1},
		254: {81, 1},
		255: {81, 1},
		256: {81, 1},
		257: {81, 2},
		258: {81, 1},
		259: {81, 1},
		260: {111, 3},
		261: {111, 5},
		262: {111, 2},
		263: {46, 2},
		264: {46, 1},
		265: {46, 1},
		266: {46, 1},
		267: {46, 1},
		268: {46, 1},
		269: {46, 1},
		270: {46, 1},
		271: {135, 3},
		272: {181, 1},
		273: {181, 3},
		274: {47, 2},
		275: {47, 2},
		276: {47, 2},
		277: {47, 2},
		278: {47, 2},
		279: {47, 2},
		280: {47, 2},
		281: {47, 1},
		282: {112, 3},
		283: {112, 5},
		284: {112, 2},
		285: {136, 3},
		286: {136, 2},
		287: {136, 4},
		288: {182, 1},
		289: {182, 3},
	}

	yyXErrors = map[yyXError]string{
		yyXError{0, 98}:   "invalid empty input",
		yyXError{1, -1}:   "expected $end",
		yyXError{60, -1}:  "expected '('",
		yyXError{26, -1}:  "expected ')'",
		yyXError{103, -1}: "expected ')'",
		yyXError{179, -1}: "expected ')'",
		yyXError{198, -1}: "expected ')'",
		yyXError{216, -1}: "expected ')'",
		yyXError{217, -1}: "expected ')'",
		yyXError{236, -1}: "expected ')'",
		yyXError{299, -1}: "expected ')'",
		yyXError{301, -1}: "expected ')'",
		yyXError{310, -1}: "expected ')'",
		yyXError{342, -1}: "expected ')'",
		yyXError{343, -1}: "expected ')'",
		yyXError{364, -1}: "expected ')'",
		yyXError{366, -1}: "expected ')'",
		yyXError{481, -1}: "expected ')'",
		yyXError{331, -1}: "expected ':'",
		yyXError{7, -1}:   "expected ';'",
		yyXError{11, -1}:  "expected ';'",
		yyXError{23, -1}:  "expected ';'",
		yyXError{29, -1}:  "expected ';'",
		yyXError{31, -1}:  "expected ';'",
		yyXError{77, -1}:  "expected ';'",
		yyXError{78, -1}:  "expected ';'",
		yyXError{79, -1}:  "expected ';'",
		yyXError{80, -1}:  "expected ';'",
		yyXError{81, -1}:  "expected ';'",
		yyXError{82, -1}:  "expected ';'",
		yyXError{83, -1}:  "expected ';'",
		yyXError{425, -1}: "expected ';'",
		yyXError{426, -1}: "expected ';'",
		yyXError{435, -1}: "expected ';'",
		yyXError{474, -1}: "expected '='",
		yyXError{52, -1}:  "expected '['",
		yyXError{337, -1}: "expected ']'",
		yyXError{393, -1}: "expected ']'",
		yyXError{492, -1}: "expected ']'",
		yyXError{296, -1}: "expected '}'",
		yyXError{352, -1}: "expected '}'",
		yyXError{377, -1}: "expected '}'",
		yyXError{403, -1}: "expected '}'",
		yyXError{422, -1}: "expected GenericArgumentsOpt or function/method signature or one of ['(', «]",
		yyXError{124, -1}: "expected GenericArgumentsOpt or one of [!=, %=, &&, &=, &^, &^=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '[', ']', '^', '{', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||, «]",
		yyXError{56, -1}:  "expected GenericArgumentsOpt or one of [!=, %=, &&, &=, &^, &^=, '%', '&', '(', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '[', '^', '{', '|', '}', *=, ++, +=, --, -=, /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, |=, ||, «]",
		yyXError{116, -1}: "expected GenericArgumentsOpt or one of ['(', ')', ',', ':', ';', '=', ']', '{', '}', ..., :=, case, default, literal, {, «, »]",
		yyXError{324, -1}: "expected argument list or one of ['!', '&', '(', ')', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{411, -1}: "expected block statement or if or '{'",
		yyXError{50, -1}:  "expected block statement or if statement header or one of ['!', '&', '(', '*', '+', '-', ';', '[', '^', <-, chan, func, identifier, interface, literal, map, struct, {]",
		yyXError{414, -1}: "expected block statement or if statement header or one of ['!', '&', '(', '*', '+', '-', ';', '[', '^', <-, chan, func, identifier, interface, literal, map, struct, {]",
		yyXError{407, -1}: "expected block statement or {",
		yyXError{416, -1}: "expected block statement or {",
		yyXError{430, -1}: "expected block statement or {",
		yyXError{76, -1}:  "expected body of the switch statement or if statement header or one of ['!', '&', '(', '*', '+', '-', ';', '[', '^', <-, chan, func, identifier, interface, literal, map, struct, {]",
		yyXError{61, -1}:  "expected body of the switch statement or {",
		yyXError{242, -1}: "expected body of the switch statement or {",
		yyXError{59, -1}:  "expected call or composite literal value or one of [!=, %=, &&, &=, &^, &^=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '[', ']', '^', '{', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{37, -1}:  "expected chan",
		yyXError{369, -1}: "expected composite literal item list or one of ['!', '&', '(', '*', '+', '-', '[', '^', '{', '}', <-, chan, func, identifier, interface, literal, map, struct, {]",
		yyXError{325, -1}: "expected composite literal item list or one of ['!', '&', '(', '*', '+', '-', '[', '^', '{', '}', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{376, -1}: "expected composite literal item or one of ['!', '&', '(', '*', '+', '-', '[', '^', '{', '}', <-, chan, func, identifier, interface, literal, map, struct, {]",
		yyXError{351, -1}: "expected composite literal item or one of ['!', '&', '(', '*', '+', '-', '[', '^', '{', '}', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{380, -1}: "expected composite literal value or expression or one of ['!', '&', '(', '*', '+', '-', '[', '^', '{', <-, chan, func, identifier, interface, literal, map, struct, {]",
		yyXError{355, -1}: "expected composite literal value or expression or one of ['!', '&', '(', '*', '+', '-', '[', '^', '{', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{358, -1}: "expected composite literal value or expression or one of ['!', '&', '(', '*', '+', '-', '[', '^', '{', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{58, -1}:  "expected composite literal value or one of ['{', {]",
		yyXError{470, -1}: "expected constant specification list or one of [')', identifier]",
		yyXError{42, -1}:  "expected constant specification or one of ['(', identifier]",
		yyXError{482, -1}: "expected constant specification or one of [')', identifier]",
		yyXError{410, -1}: "expected else if clause or optional else clause or one of [';', '}', case, default, else]",
		yyXError{409, -1}: "expected else if list clause or optional else clause or one of [';', '}', case, default, else]",
		yyXError{438, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, range, struct]",
		yyXError{450, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, range, struct]",
		yyXError{119, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{121, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{439, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{440, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{441, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{442, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{443, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{444, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{445, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{446, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{447, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{448, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{449, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{473, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{475, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{490, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{491, -1}: "expected expression list or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{34, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '[', ']', '^', ..., <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{65, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{68, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{134, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
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
		yyXError{150, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{151, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{152, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{153, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{154, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{155, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{214, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{273, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{274, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{431, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{452, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{466, -1}: "expected expression or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{327, -1}: "expected expression or optional expression or one of ['!', '&', '(', '*', '+', '-', ':', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{53, -1}:  "expected expression or type literal or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{363, -1}: "expected expression/type literal or one of ['!', '&', '(', ')', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{271, -1}: "expected expression/type literal or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{45, -1}:  "expected for statement header or one of ['!', '&', '(', '*', '+', '-', ';', '[', '^', <-, chan, func, identifier, interface, literal, map, range, struct, {]",
		yyXError{106, -1}: "expected function/method signature or '('",
		yyXError{423, -1}: "expected function/method signature or '('",
		yyXError{399, -1}: "expected function/method signature or one of ['(', '.', ';', '}']",
		yyXError{47, -1}:  "expected function/method signature or optional receiver or one of ['(', identifier]",
		yyXError{131, -1}: "expected generic argument list item or one of ['(', '*', '[', <-, chan, func, identifier, interface, map, struct]",
		yyXError{125, -1}: "expected generic argument list or one of ['(', '*', '[', <-, chan, func, identifier, interface, map, struct]",
		yyXError{229, -1}: "expected generic parameter list item or identifier",
		yyXError{222, -1}: "expected generic parameter list or identifier",
		yyXError{3, -1}:   "expected identifier",
		yyXError{30, -1}:  "expected identifier",
		yyXError{69, -1}:  "expected identifier",
		yyXError{107, -1}: "expected identifier",
		yyXError{181, -1}: "expected identifier",
		yyXError{420, -1}: "expected identifier",
		yyXError{10, -1}:  "expected import specification list or one of [')', '.', identifier, literal]",
		yyXError{6, -1}:   "expected import specification or one of ['(', '.', identifier, literal]",
		yyXError{27, -1}:  "expected import specification or one of [')', '.', identifier, literal]",
		yyXError{396, -1}: "expected interface method declaration list or one of ['}', identifier]",
		yyXError{404, -1}: "expected interface method declaration or one of ['}', identifier]",
		yyXError{55, -1}:  "expected left brace or one of ['(', ')', ',', ':', '=', '{', ..., :=, {]",
		yyXError{51, -1}:  "expected left brace or one of ['{', {]",
		yyXError{75, -1}:  "expected left brace or one of ['{', {]",
		yyXError{12, -1}:  "expected literal or ",
		yyXError{13, -1}:  "expected literal or ",
		yyXError{126, -1}: "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '[', ']', '^', '{', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, literal, {, |=, ||, »]",
		yyXError{132, -1}: "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '[', ']', '^', '{', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, literal, {, |=, ||, »]",
		yyXError{54, -1}:  "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '[', ']', '^', '{', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{57, -1}:  "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '[', ']', '^', '{', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{127, -1}: "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '[', ']', '^', '{', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{218, -1}: "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '[', ']', '^', '{', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{328, -1}: "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '[', ']', '^', '{', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{329, -1}: "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '[', ']', '^', '{', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{336, -1}: "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '[', ']', '^', '{', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{338, -1}: "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '[', ']', '^', '{', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{339, -1}: "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '[', ']', '^', '{', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{341, -1}: "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '[', ']', '^', '{', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{344, -1}: "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '[', ']', '^', '{', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{345, -1}: "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '[', ']', '^', '{', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{350, -1}: "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '[', ']', '^', '{', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{353, -1}: "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '[', ']', '^', '{', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{362, -1}: "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '[', ']', '^', '{', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{367, -1}: "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '[', ']', '^', '{', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{368, -1}: "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '[', ']', '^', '{', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{370, -1}: "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '[', ']', '^', '{', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{375, -1}: "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '[', ']', '^', '{', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{378, -1}: "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '[', ']', '^', '{', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{387, -1}: "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '[', ']', '^', '{', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{390, -1}: "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '[', ']', '^', '{', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{391, -1}: "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '[', ']', '^', '{', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{43, -1}:  "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', ')', '*', '+', ',', '-', '/', ':', ';', '<', '=', '>', ']', '^', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
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
		yyXError{170, -1}: "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', ')', '*', '+', ',', '-', '/', ':', ';', '<', '=', '>', ']', '^', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{171, -1}: "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', ')', '*', '+', ',', '-', '/', ':', ';', '<', '=', '>', ']', '^', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{172, -1}: "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', ')', '*', '+', ',', '-', '/', ':', ';', '<', '=', '>', ']', '^', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{173, -1}: "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', ')', '*', '+', ',', '-', '/', ':', ';', '<', '=', '>', ']', '^', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{174, -1}: "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', ')', '*', '+', ',', '-', '/', ':', ';', '<', '=', '>', ']', '^', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{175, -1}: "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', ')', '*', '+', ',', '-', '/', ':', ';', '<', '=', '>', ']', '^', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{206, -1}: "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', ')', '*', '+', ',', '-', '/', ':', ';', '<', '=', '>', ']', '^', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{207, -1}: "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', ')', '*', '+', ',', '-', '/', ':', ';', '<', '=', '>', ']', '^', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{208, -1}: "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', ')', '*', '+', ',', '-', '/', ':', ';', '<', '=', '>', ']', '^', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{209, -1}: "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', ')', '*', '+', ',', '-', '/', ':', ';', '<', '=', '>', ']', '^', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{210, -1}: "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', ')', '*', '+', ',', '-', '/', ':', ';', '<', '=', '>', ']', '^', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{211, -1}: "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', ')', '*', '+', ',', '-', '/', ':', ';', '<', '=', '>', ']', '^', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{213, -1}: "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', ')', '*', '+', ',', '-', '/', ':', ';', '<', '=', '>', ']', '^', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{135, -1}: "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', ')', '*', '+', ',', '-', '/', ';', '<', '=', '>', '^', '|', '}', *=, +=, -=, /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{44, -1}:  "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', '*', '+', ',', '-', '/', ';', '<', '=', '>', '^', '|', '}', *=, ++, +=, --, -=, /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{266, -1}: "expected one of [!=, &&, &^, '%', '&', ')', '*', '+', ',', '-', '/', ':', '<', '=', '>', '^', '|', ..., :=, <-, <<, <=, ==, >=, >>, ||]",
		yyXError{122, -1}: "expected one of [!=, &&, &^, '%', '&', ')', '*', '+', ',', '-', '/', ';', '<', '>', '^', '|', '}', <-, <<, <=, ==, >=, >>, case, default, {, ||]",
		yyXError{388, -1}: "expected one of [!=, &&, &^, '%', '&', ')', '*', '+', '-', '/', '<', '>', '^', '|', <-, <<, <=, ==, >=, >>, ||]",
		yyXError{347, -1}: "expected one of [!=, &&, &^, '%', '&', '*', '+', ',', '-', '/', ':', '<', '>', '^', '|', '}', <-, <<, <=, ==, >=, >>, ||]",
		yyXError{371, -1}: "expected one of [!=, &&, &^, '%', '&', '*', '+', ',', '-', '/', ':', '<', '>', '^', '|', '}', <-, <<, <=, ==, >=, >>, ||]",
		yyXError{357, -1}: "expected one of [!=, &&, &^, '%', '&', '*', '+', ',', '-', '/', '<', '>', '^', '|', '}', <-, <<, <=, ==, >=, >>, ||]",
		yyXError{360, -1}: "expected one of [!=, &&, &^, '%', '&', '*', '+', ',', '-', '/', '<', '>', '^', '|', '}', <-, <<, <=, ==, >=, >>, ||]",
		yyXError{381, -1}: "expected one of [!=, &&, &^, '%', '&', '*', '+', ',', '-', '/', '<', '>', '^', '|', '}', <-, <<, <=, ==, >=, >>, ||]",
		yyXError{330, -1}: "expected one of [!=, &&, &^, '%', '&', '*', '+', '-', '/', ':', '<', '>', ']', '^', '|', <-, <<, <=, ==, >=, >>, ||]",
		yyXError{333, -1}: "expected one of [!=, &&, &^, '%', '&', '*', '+', '-', '/', ':', '<', '>', ']', '^', '|', <-, <<, <=, ==, >=, >>, ||]",
		yyXError{275, -1}: "expected one of [!=, &&, &^, '%', '&', '*', '+', '-', '/', ':', '<', '>', '^', '|', <-, <<, <=, ==, >=, >>, ||]",
		yyXError{277, -1}: "expected one of [!=, &&, &^, '%', '&', '*', '+', '-', '/', ':', '<', '>', '^', '|', <-, <<, <=, ==, >=, >>, ||]",
		yyXError{319, -1}: "expected one of [!=, &&, &^, '%', '&', '*', '+', '-', '/', ';', '<', '>', '^', '|', '}', <-, <<, <=, ==, >=, >>, case, default, ||]",
		yyXError{320, -1}: "expected one of [!=, &&, &^, '%', '&', '*', '+', '-', '/', ';', '<', '>', '^', '|', '}', <-, <<, <=, ==, >=, >>, case, default, ||]",
		yyXError{493, -1}: "expected one of [!=, &&, &^, '%', '&', '*', '+', '-', '/', '<', '>', ']', '^', '|', <-, <<, <=, ==, >=, >>, ||]",
		yyXError{432, -1}: "expected one of [!=, &&, &^, '%', '&', '*', '+', '-', '/', '<', '>', '^', '|', <-, <<, <=, ==, >=, >>, {, ||]",
		yyXError{453, -1}: "expected one of [!=, &&, &^, '%', '&', '*', '+', '-', '/', '<', '>', '^', '|', <-, <<, <=, ==, >=, >>, {, ||]",
		yyXError{467, -1}: "expected one of [!=, &&, &^, '%', '&', '*', '+', '-', '/', '<', '>', '^', '|', <-, <<, <=, ==, >=, >>, {, ||]",
		yyXError{35, -1}:  "expected one of [%=, &=, &^=, ',', '=', *=, +=, -=, /=, :=, <<=, >>=, ^=, |=]",
		yyXError{427, -1}: "expected one of [%=, &=, &^=, ',', '=', *=, +=, -=, /=, :=, <<=, >>=, ^=, |=]",
		yyXError{49, -1}:  "expected one of ['!', '&', '(', '*', '+', '-', ';', '[', '^', <-, chan, func, identifier, interface, literal, map, struct, {]",
		yyXError{196, -1}: "expected one of ['(', ')', '*', ',', ':', ';', '=', '[', ']', '{', '}', ..., :=, <-, case, chan, default, func, identifier, interface, literal, map, struct, {, »]",
		yyXError{199, -1}: "expected one of ['(', ')', '*', ',', ':', ';', '=', '[', ']', '{', '}', ..., :=, <-, case, chan, default, func, identifier, interface, literal, map, struct, {, »]",
		yyXError{95, -1}:  "expected one of ['(', ')', '*', ',', ';', '=', '[', '}', <-, case, chan, default, func, identifier, interface, map, struct]",
		yyXError{183, -1}: "expected one of ['(', ')', '*', ',', ';', '=', '[', '}', <-, case, chan, default, func, identifier, interface, map, struct]",
		yyXError{108, -1}: "expected one of ['(', ')', ',', '.', ':', ';', '=', ']', '{', '}', ..., :=, case, default, literal, {, «, »]",
		yyXError{182, -1}: "expected one of ['(', ')', ',', ':', ';', '=', ']', '{', '}', ..., :=, case, default, literal, {, «, »]",
		yyXError{111, -1}: "expected one of ['(', ')', ',', ':', ';', '=', ']', '{', '}', ..., :=, case, default, literal, {, »]",
		yyXError{112, -1}: "expected one of ['(', ')', ',', ':', ';', '=', ']', '{', '}', ..., :=, case, default, literal, {, »]",
		yyXError{113, -1}: "expected one of ['(', ')', ',', ':', ';', '=', ']', '{', '}', ..., :=, case, default, literal, {, »]",
		yyXError{114, -1}: "expected one of ['(', ')', ',', ':', ';', '=', ']', '{', '}', ..., :=, case, default, literal, {, »]",
		yyXError{115, -1}: "expected one of ['(', ')', ',', ':', ';', '=', ']', '{', '}', ..., :=, case, default, literal, {, »]",
		yyXError{117, -1}: "expected one of ['(', ')', ',', ':', ';', '=', ']', '{', '}', ..., :=, case, default, literal, {, »]",
		yyXError{118, -1}: "expected one of ['(', ')', ',', ':', ';', '=', ']', '{', '}', ..., :=, case, default, literal, {, »]",
		yyXError{177, -1}: "expected one of ['(', ')', ',', ':', ';', '=', ']', '{', '}', ..., :=, case, default, literal, {, »]",
		yyXError{178, -1}: "expected one of ['(', ')', ',', ':', ';', '=', ']', '{', '}', ..., :=, case, default, literal, {, »]",
		yyXError{180, -1}: "expected one of ['(', ')', ',', ':', ';', '=', ']', '{', '}', ..., :=, case, default, literal, {, »]",
		yyXError{184, -1}: "expected one of ['(', ')', ',', ':', ';', '=', ']', '{', '}', ..., :=, case, default, literal, {, »]",
		yyXError{188, -1}: "expected one of ['(', ')', ',', ':', ';', '=', ']', '{', '}', ..., :=, case, default, literal, {, »]",
		yyXError{189, -1}: "expected one of ['(', ')', ',', ':', ';', '=', ']', '{', '}', ..., :=, case, default, literal, {, »]",
		yyXError{190, -1}: "expected one of ['(', ')', ',', ':', ';', '=', ']', '{', '}', ..., :=, case, default, literal, {, »]",
		yyXError{294, -1}: "expected one of ['(', ')', ',', ':', ';', '=', ']', '{', '}', ..., :=, case, default, literal, {, »]",
		yyXError{297, -1}: "expected one of ['(', ')', ',', ':', ';', '=', ']', '{', '}', ..., :=, case, default, literal, {, »]",
		yyXError{395, -1}: "expected one of ['(', ')', ',', ':', ';', '=', ']', '{', '}', ..., :=, case, default, literal, {, »]",
		yyXError{397, -1}: "expected one of ['(', ')', ',', ':', ';', '=', ']', '{', '}', ..., :=, case, default, literal, {, »]",
		yyXError{406, -1}: "expected one of ['(', ')', ',', ':', ';', '=', ']', '{', '}', ..., :=, case, default, literal, {, »]",
		yyXError{486, -1}: "expected one of ['(', ')', ',', ':', ';', '=', ']', '{', '}', ..., :=, case, default, literal, {, »]",
		yyXError{487, -1}: "expected one of ['(', ')', ',', ':', ';', '=', ']', '{', '}', ..., :=, case, default, literal, {, »]",
		yyXError{489, -1}: "expected one of ['(', ')', ',', ':', ';', '=', ']', '{', '}', ..., :=, case, default, literal, {, »]",
		yyXError{495, -1}: "expected one of ['(', ')', ',', ':', ';', '=', ']', '{', '}', ..., :=, case, default, literal, {, »]",
		yyXError{497, -1}: "expected one of ['(', ')', ',', ':', ';', '=', ']', '{', '}', ..., :=, case, default, literal, {, »]",
		yyXError{499, -1}: "expected one of ['(', ')', ',', ':', ';', '=', ']', '{', '}', ..., :=, case, default, literal, {, »]",
		yyXError{38, -1}:  "expected one of ['(', ')', ',', ':', '=', '{', ..., :=, {]",
		yyXError{39, -1}:  "expected one of ['(', ')', ',', ':', '=', '{', ..., :=, {]",
		yyXError{40, -1}:  "expected one of ['(', ')', ',', ':', '=', '{', ..., :=, {]",
		yyXError{41, -1}:  "expected one of ['(', ')', ',', ':', '=', '{', ..., :=, {]",
		yyXError{86, -1}:  "expected one of ['(', ')', ',', ':', '=', ..., :=]",
		yyXError{87, -1}:  "expected one of ['(', ')', ',', ':', '=', ..., :=]",
		yyXError{212, -1}: "expected one of ['(', ')', ',', ':', '=', ..., :=]",
		yyXError{267, -1}: "expected one of ['(', ')', ',', ':', '=', ..., :=]",
		yyXError{389, -1}: "expected one of ['(', ')']",
		yyXError{287, -1}: "expected one of ['(', '*', ',', '.', ';', '[', '}', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{223, -1}: "expected one of ['(', '*', '[', <-, chan, func, identifier, interface, map, struct]",
		yyXError{230, -1}: "expected one of ['(', '*', '[', <-, chan, func, identifier, interface, map, struct]",
		yyXError{48, -1}:  "expected one of ['(', identifier]",
		yyXError{326, -1}: "expected one of ['(', identifier]",
		yyXError{268, -1}: "expected one of [')', ',', ':', '=', ..., :=]",
		yyXError{279, -1}: "expected one of [')', ',', ':', '=', ..., :=]",
		yyXError{123, -1}: "expected one of [')', ',', ';', '}', case, default]",
		yyXError{176, -1}: "expected one of [')', ',', ';', '}', case, default]",
		yyXError{476, -1}: "expected one of [')', ',', ';', '}', case, default]",
		yyXError{477, -1}: "expected one of [')', ',', ';', '}', case, default]",
		yyXError{193, -1}: "expected one of [')', ',']",
		yyXError{194, -1}: "expected one of [')', ',']",
		yyXError{200, -1}: "expected one of [')', ',']",
		yyXError{201, -1}: "expected one of [')', ',']",
		yyXError{202, -1}: "expected one of [')', ',']",
		yyXError{204, -1}: "expected one of [')', ',']",
		yyXError{205, -1}: "expected one of [')', ',']",
		yyXError{120, -1}: "expected one of [')', ';', '=', '}', case, default]",
		yyXError{225, -1}: "expected one of [')', ';', '}', case, default]",
		yyXError{20, -1}:  "expected one of [')', ';']",
		yyXError{22, -1}:  "expected one of [')', ';']",
		yyXError{25, -1}:  "expected one of [')', ';']",
		yyXError{28, -1}:  "expected one of [')', ';']",
		yyXError{101, -1}: "expected one of [')', ';']",
		yyXError{105, -1}: "expected one of [')', ';']",
		yyXError{234, -1}: "expected one of [')', ';']",
		yyXError{238, -1}: "expected one of [')', ';']",
		yyXError{480, -1}: "expected one of [')', ';']",
		yyXError{483, -1}: "expected one of [')', ';']",
		yyXError{269, -1}: "expected one of [',', ':', '=', :=]",
		yyXError{346, -1}: "expected one of [',', ':', '}']",
		yyXError{451, -1}: "expected one of [',', ';', '}', case, default, {]",
		yyXError{454, -1}: "expected one of [',', ';', '}', case, default, {]",
		yyXError{455, -1}: "expected one of [',', ';', '}', case, default, {]",
		yyXError{456, -1}: "expected one of [',', ';', '}', case, default, {]",
		yyXError{457, -1}: "expected one of [',', ';', '}', case, default, {]",
		yyXError{458, -1}: "expected one of [',', ';', '}', case, default, {]",
		yyXError{459, -1}: "expected one of [',', ';', '}', case, default, {]",
		yyXError{460, -1}: "expected one of [',', ';', '}', case, default, {]",
		yyXError{461, -1}: "expected one of [',', ';', '}', case, default, {]",
		yyXError{462, -1}: "expected one of [',', ';', '}', case, default, {]",
		yyXError{463, -1}: "expected one of [',', ';', '}', case, default, {]",
		yyXError{464, -1}: "expected one of [',', ';', '}', case, default, {]",
		yyXError{465, -1}: "expected one of [',', ';', '}', case, default, {]",
		yyXError{316, -1}: "expected one of [',', ';', '}', case, default]",
		yyXError{348, -1}: "expected one of [',', '}']",
		yyXError{354, -1}: "expected one of [',', '}']",
		yyXError{356, -1}: "expected one of [',', '}']",
		yyXError{359, -1}: "expected one of [',', '}']",
		yyXError{372, -1}: "expected one of [',', '}']",
		yyXError{373, -1}: "expected one of [',', '}']",
		yyXError{379, -1}: "expected one of [',', '}']",
		yyXError{382, -1}: "expected one of [',', '}']",
		yyXError{128, -1}: "expected one of [',', »]",
		yyXError{129, -1}: "expected one of [',', »]",
		yyXError{130, -1}: "expected one of [',', »]",
		yyXError{133, -1}: "expected one of [',', »]",
		yyXError{226, -1}: "expected one of [',', »]",
		yyXError{227, -1}: "expected one of [',', »]",
		yyXError{228, -1}: "expected one of [',', »]",
		yyXError{231, -1}: "expected one of [',', »]",
		yyXError{334, -1}: "expected one of [':', ']']",
		yyXError{424, -1}: "expected one of [';', '{']",
		yyXError{412, -1}: "expected one of [';', '}', case, default, else]",
		yyXError{417, -1}: "expected one of [';', '}', case, default, else]",
		yyXError{419, -1}: "expected one of [';', '}', case, default, else]",
		yyXError{9, -1}:   "expected one of [';', '}', case, default, literal]",
		yyXError{62, -1}:  "expected one of [';', '}', case, default, {]",
		yyXError{468, -1}: "expected one of [';', '}', case, default, {]",
		yyXError{469, -1}: "expected one of [';', '}', case, default, {]",
		yyXError{66, -1}:  "expected one of [';', '}', case, default]",
		yyXError{67, -1}:  "expected one of [';', '}', case, default]",
		yyXError{70, -1}:  "expected one of [';', '}', case, default]",
		yyXError{72, -1}:  "expected one of [';', '}', case, default]",
		yyXError{73, -1}:  "expected one of [';', '}', case, default]",
		yyXError{74, -1}:  "expected one of [';', '}', case, default]",
		yyXError{98, -1}:  "expected one of [';', '}', case, default]",
		yyXError{99, -1}:  "expected one of [';', '}', case, default]",
		yyXError{104, -1}: "expected one of [';', '}', case, default]",
		yyXError{220, -1}: "expected one of [';', '}', case, default]",
		yyXError{232, -1}: "expected one of [';', '}', case, default]",
		yyXError{237, -1}: "expected one of [';', '}', case, default]",
		yyXError{244, -1}: "expected one of [';', '}', case, default]",
		yyXError{245, -1}: "expected one of [';', '}', case, default]",
		yyXError{252, -1}: "expected one of [';', '}', case, default]",
		yyXError{253, -1}: "expected one of [';', '}', case, default]",
		yyXError{254, -1}: "expected one of [';', '}', case, default]",
		yyXError{255, -1}: "expected one of [';', '}', case, default]",
		yyXError{256, -1}: "expected one of [';', '}', case, default]",
		yyXError{257, -1}: "expected one of [';', '}', case, default]",
		yyXError{258, -1}: "expected one of [';', '}', case, default]",
		yyXError{259, -1}: "expected one of [';', '}', case, default]",
		yyXError{261, -1}: "expected one of [';', '}', case, default]",
		yyXError{263, -1}: "expected one of [';', '}', case, default]",
		yyXError{280, -1}: "expected one of [';', '}', case, default]",
		yyXError{317, -1}: "expected one of [';', '}', case, default]",
		yyXError{318, -1}: "expected one of [';', '}', case, default]",
		yyXError{321, -1}: "expected one of [';', '}', case, default]",
		yyXError{322, -1}: "expected one of [';', '}', case, default]",
		yyXError{323, -1}: "expected one of [';', '}', case, default]",
		yyXError{384, -1}: "expected one of [';', '}', case, default]",
		yyXError{413, -1}: "expected one of [';', '}', case, default]",
		yyXError{415, -1}: "expected one of [';', '}', case, default]",
		yyXError{433, -1}: "expected one of [';', '}', case, default]",
		yyXError{471, -1}: "expected one of [';', '}', case, default]",
		yyXError{478, -1}: "expected one of [';', '}', case, default]",
		yyXError{484, -1}: "expected one of [';', '}', case, default]",
		yyXError{262, -1}: "expected one of [';', '}']",
		yyXError{292, -1}: "expected one of [';', '}']",
		yyXError{298, -1}: "expected one of [';', '}']",
		yyXError{303, -1}: "expected one of [';', '}']",
		yyXError{304, -1}: "expected one of [';', '}']",
		yyXError{306, -1}: "expected one of [';', '}']",
		yyXError{307, -1}: "expected one of [';', '}']",
		yyXError{312, -1}: "expected one of [';', '}']",
		yyXError{313, -1}: "expected one of [';', '}']",
		yyXError{315, -1}: "expected one of [';', '}']",
		yyXError{386, -1}: "expected one of [';', '}']",
		yyXError{400, -1}: "expected one of [';', '}']",
		yyXError{401, -1}: "expected one of [';', '}']",
		yyXError{402, -1}: "expected one of [';', '}']",
		yyXError{405, -1}: "expected one of [';', '}']",
		yyXError{418, -1}: "expected one of [';', '}']",
		yyXError{240, -1}: "expected one of [';', {]",
		yyXError{241, -1}: "expected one of [';', {]",
		yyXError{429, -1}: "expected one of [';', {]",
		yyXError{250, -1}: "expected one of ['}', case, default]",
		yyXError{281, -1}: "expected one of ['}', case, default]",
		yyXError{46, -1}:  "expected optional block or one of [';', '{']",
		yyXError{215, -1}: "expected optional comma or one of [!=, &&, &^, '%', '&', ')', '*', '+', ',', '-', '/', '<', '>', '^', '|', <-, <<, <=, ==, >=, >>, ||]",
		yyXError{361, -1}: "expected optional comma or one of [')', ',', ...]",
		yyXError{195, -1}: "expected optional comma or one of [')', ',']",
		yyXError{365, -1}: "expected optional comma or one of [')', ',']",
		yyXError{349, -1}: "expected optional comma or one of [',', '}']",
		yyXError{374, -1}: "expected optional comma or one of [',', '}']",
		yyXError{71, -1}:  "expected optional expression list or one of ['!', '&', '(', '*', '+', '-', ';', '[', '^', '}', <-, case, chan, default, func, identifier, interface, literal, map, struct]",
		yyXError{332, -1}: "expected optional expression or one of ['!', '&', '(', '*', '+', '-', ':', '[', ']', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{335, -1}: "expected optional expression or one of ['!', '&', '(', '*', '+', '-', '[', ']', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{221, -1}: "expected optional generic parameters or type or one of ['(', '*', '[', <-, chan, func, identifier, interface, map, struct, «]",
		yyXError{63, -1}:  "expected optional identifier or one of [';', '}', case, default, identifier]",
		yyXError{64, -1}:  "expected optional identifier or one of [';', '}', case, default, identifier]",
		yyXError{186, -1}: "expected optional result or one of ['(', ')', '*', ',', ':', ';', '=', '[', ']', '{', '}', ..., :=, <-, case, chan, default, func, identifier, interface, literal, map, struct, {, »]",
		yyXError{421, -1}: "expected optional result or one of ['(', '*', '[', '{', <-, chan, func, identifier, interface, map, struct, {]",
		yyXError{24, -1}:  "expected optional semicolon or one of [')', ';']",
		yyXError{100, -1}: "expected optional semicolon or one of [')', ';']",
		yyXError{233, -1}: "expected optional semicolon or one of [')', ';']",
		yyXError{479, -1}: "expected optional semicolon or one of [')', ';']",
		yyXError{293, -1}: "expected optional semicolon or one of [';', '}']",
		yyXError{398, -1}: "expected optional semicolon or one of [';', '}']",
		yyXError{434, -1}: "expected optional simple statement or one of ['!', '&', '(', '*', '+', '-', ';', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{282, -1}: "expected optional simple statement or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct, {]",
		yyXError{436, -1}: "expected optional simple statement or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct, {]",
		yyXError{290, -1}: "expected optional string literal or one of [';', '}', literal]",
		yyXError{302, -1}: "expected optional string literal or one of [';', '}', literal]",
		yyXError{305, -1}: "expected optional string literal or one of [';', '}', literal]",
		yyXError{308, -1}: "expected optional string literal or one of [';', '}', literal]",
		yyXError{311, -1}: "expected optional string literal or one of [';', '}', literal]",
		yyXError{314, -1}: "expected optional string literal or one of [';', '}', literal]",
		yyXError{185, -1}: "expected parameter declaration list or one of ['(', ')', '*', '[', ..., <-, chan, func, identifier, interface, map, struct]",
		yyXError{187, -1}: "expected parameter declaration list or type or one of ['(', ')', '*', '[', ..., <-, chan, func, identifier, interface, map, struct]",
		yyXError{197, -1}: "expected parameter declaration or one of ['(', ')', '*', '[', ..., <-, chan, func, identifier, interface, map, struct]",
		yyXError{300, -1}: "expected qualified identifier or identifier",
		yyXError{309, -1}: "expected qualified identifier or identifier",
		yyXError{289, -1}: "expected qualified identifier or one of ['(', identifier]",
		yyXError{291, -1}: "expected qualified identifier or one of ['*', identifier]",
		yyXError{0, -1}:   "expected source file or package",
		yyXError{286, -1}: "expected struct field declaration list or one of ['(', '*', '}', identifier]",
		yyXError{295, -1}: "expected struct field declaration or one of ['(', '*', '}', identifier]",
		yyXError{243, -1}: "expected switch case/default clause list or one of ['}', case, default]",
		yyXError{246, -1}: "expected switch case/default clause statement block or one of ['}', case, default]",
		yyXError{85, -1}:  "expected type literal or unary expression or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{191, -1}: "expected type or one of ['(', ')', '*', ',', '.', '[', ..., <-, chan, func, identifier, interface, map, struct, «]",
		yyXError{472, -1}: "expected type or one of ['(', ')', '*', ',', ';', '=', '[', '}', <-, case, chan, default, func, identifier, interface, map, struct]",
		yyXError{96, -1}:  "expected type or one of ['(', '*', ',', '=', '[', <-, chan, func, identifier, interface, map, struct]",
		yyXError{288, -1}: "expected type or one of ['(', '*', ',', '[', <-, chan, func, identifier, interface, map, struct]",
		yyXError{36, -1}:  "expected type or one of ['(', '*', '[', <-, chan, func, identifier, interface, map, struct]",
		yyXError{109, -1}: "expected type or one of ['(', '*', '[', <-, chan, func, identifier, interface, map, struct]",
		yyXError{110, -1}: "expected type or one of ['(', '*', '[', <-, chan, func, identifier, interface, map, struct]",
		yyXError{192, -1}: "expected type or one of ['(', '*', '[', <-, chan, func, identifier, interface, map, struct]",
		yyXError{203, -1}: "expected type or one of ['(', '*', '[', <-, chan, func, identifier, interface, map, struct]",
		yyXError{224, -1}: "expected type or one of ['(', '*', '[', <-, chan, func, identifier, interface, map, struct]",
		yyXError{340, -1}: "expected type or one of ['(', '*', '[', <-, chan, func, identifier, interface, map, struct]",
		yyXError{392, -1}: "expected type or one of ['(', '*', '[', <-, chan, func, identifier, interface, map, struct]",
		yyXError{394, -1}: "expected type or one of ['(', '*', '[', <-, chan, func, identifier, interface, map, struct]",
		yyXError{485, -1}: "expected type or one of ['(', '*', '[', <-, chan, func, identifier, interface, map, struct]",
		yyXError{488, -1}: "expected type or one of ['(', '*', '[', <-, chan, func, identifier, interface, map, struct]",
		yyXError{494, -1}: "expected type or one of ['(', '*', '[', <-, chan, func, identifier, interface, map, struct]",
		yyXError{496, -1}: "expected type or one of ['(', '*', '[', <-, chan, func, identifier, interface, map, struct]",
		yyXError{498, -1}: "expected type or one of ['(', '*', '[', <-, chan, func, identifier, interface, map, struct]",
		yyXError{219, -1}: "expected type specification list or one of [')', identifier]",
		yyXError{84, -1}:  "expected type specification or one of ['(', identifier]",
		yyXError{235, -1}: "expected type specification or one of [')', identifier]",
		yyXError{88, -1}:  "expected unary expression or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{89, -1}:  "expected unary expression or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{90, -1}:  "expected unary expression or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{91, -1}:  "expected unary expression or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{92, -1}:  "expected unary expression or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{93, -1}:  "expected unary expression or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{97, -1}:  "expected variable specification list or one of [')', identifier]",
		yyXError{94, -1}:  "expected variable specification or one of ['(', identifier]",
		yyXError{102, -1}: "expected variable specification or one of [')', identifier]",
		yyXError{283, -1}: "expected {",
		yyXError{428, -1}: "expected {",
		yyXError{437, -1}: "expected {",
		yyXError{48, 102}: "anonymous functions cannot have generic type parameters",
	}

	yyParseTab = [500][]uint16{
		// 0
		{158: 291, 171: 293, 294, 292},
		{98: 290},
		{2: 42, 42, 42, 42, 42, 42, 42, 13: 42, 42, 42, 17: 42, 42, 42, 42, 42, 42, 35: 42, 42, 42, 42, 42, 80: 42, 82: 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 98: 42, 179: 323},
		{14: 136, 149: 320},
		{2: 162, 162, 162, 162, 162, 162, 162, 13: 162, 162, 162, 17: 162, 162, 162, 162, 162, 162, 35: 162, 162, 162, 162, 162, 80: 162, 82: 162, 162, 162, 162, 162, 162, 162, 162, 162, 162, 162, 162, 162, 162, 98: 162, 124: 162, 167: 295},
		// 5
		{2: 115, 115, 115, 115, 115, 115, 115, 13: 115, 115, 115, 17: 115, 115, 115, 115, 115, 115, 35: 115, 115, 115, 115, 115, 80: 115, 82: 115, 115, 115, 115, 115, 115, 115, 115, 115, 115, 115, 115, 115, 115, 98: 115, 124: 296, 166: 297},
		{2: 300, 14: 299, 179, 35: 179, 179, 179, 179, 77: 302, 120: 303, 132: 301},
		{298},
		{2: 161, 161, 161, 161, 161, 161, 161, 13: 161, 161, 161, 17: 161, 161, 161, 161, 161, 161, 35: 161, 161, 161, 161, 161, 80: 161, 82: 161, 161, 161, 161, 161, 161, 161, 161, 161, 161, 161, 161, 161, 161, 98: 161, 124: 161},
		{178, 178, 9: 178, 178, 15: 178, 35: 178, 178, 178, 178},
		// 10
		{11: 313, 14: 299, 179, 35: 179, 179, 179, 179, 77: 302, 120: 303, 132: 315, 168: 314},
		{169},
		{15: 308, 35: 304, 305, 306, 307, 42: 311},
		{15: 308, 35: 304, 305, 306, 307, 42: 309},
		{270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 16: 270, 23: 270, 270, 33: 270, 270, 40: 270, 270, 48: 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 63: 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 77: 270, 80: 270},
		// 15
		{269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 16: 269, 23: 269, 269, 33: 269, 269, 40: 269, 269, 48: 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 63: 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 77: 269, 80: 269},
		{268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 16: 268, 23: 268, 268, 33: 268, 268, 40: 268, 268, 48: 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 63: 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 77: 268, 80: 268},
		{267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 16: 267, 23: 267, 267, 33: 267, 267, 40: 267, 267, 48: 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 63: 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 77: 267, 80: 267},
		{266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 16: 266, 23: 266, 266, 33: 266, 266, 40: 266, 266, 48: 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 63: 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 77: 266, 80: 266},
		{167, 11: 167, 80: 310},
		// 20
		{165, 11: 165},
		{168, 11: 168, 80: 312},
		{166, 11: 166},
		{171},
		{317, 11: 105, 115: 316},
		// 25
		{164, 11: 164},
		{11: 319},
		{11: 104, 14: 299, 179, 35: 179, 179, 179, 179, 77: 302, 120: 303, 132: 318},
		{163, 11: 163},
		{170},
		// 30
		{14: 321},
		{322},
		{2: 135, 135, 135, 135, 135, 135, 135, 13: 135, 135, 135, 17: 135, 135, 135, 135, 135, 135, 35: 135, 135, 135, 135, 135, 80: 135, 82: 135, 135, 135, 135, 135, 135, 135, 135, 135, 135, 135, 135, 135, 135, 98: 135, 124: 135},
		{2: 343, 375, 379, 380, 381, 382, 383, 13: 324, 346, 308, 17: 326, 338, 341, 342, 327, 365, 25: 328, 376, 337, 345, 377, 329, 330, 331, 35: 304, 305, 306, 307, 378, 42: 344, 348, 347, 349, 350, 333, 62: 334, 78: 325, 80: 372, 82: 339, 374, 353, 332, 354, 355, 356, 335, 358, 359, 361, 351, 366, 384, 352, 363, 289, 101: 340, 103: 367, 357, 360, 107: 362, 109: 371, 364, 369, 370, 160: 368, 336, 178: 373},
		{2: 343, 375, 379, 380, 381, 382, 383, 13: 324, 414, 308, 17: 326, 338, 341, 342, 327, 365, 25: 328, 376, 396, 345, 377, 329, 330, 331, 34: 782, 304, 305, 306, 307, 378, 784, 42: 344, 348, 347, 349, 350, 333, 62: 783},
		// 35
		{12: 424, 24: 780, 33: 781, 63: 729, 731, 730, 732, 733, 734, 735, 736, 737, 738, 739},
		{2: 399, 400, 13: 324, 398, 17: 326, 338, 341, 342, 327, 365, 25: 401, 402, 396, 403, 404, 405, 407, 408, 79: 406, 81: 777, 180: 778},
		{17: 775},
		{2: 26, 11: 26, 26, 16: 245, 23: 26, 26, 33: 26, 26, 41: 245},
		{2: 22, 11: 22, 22, 16: 244, 23: 22, 22, 33: 22, 22, 41: 244},
		// 40
		{2: 21, 11: 21, 21, 16: 243, 23: 21, 21, 33: 21, 21, 41: 243},
		{2: 20, 11: 20, 20, 16: 242, 23: 20, 20, 33: 20, 20, 41: 242},
		{2: 760, 14: 385, 100: 762, 129: 761},
		{226, 226, 3: 226, 226, 226, 226, 226, 226, 226, 226, 226, 226, 16: 226, 23: 226, 226, 33: 226, 226, 40: 226, 48: 226, 226, 226, 226, 226, 226, 226, 226, 226, 226, 226, 226, 226, 226, 63: 226, 226, 226, 226, 226, 226, 226, 226, 226, 226, 226, 226, 226},
		{101, 101, 3: 428, 427, 429, 430, 434, 445, 101, 101, 12: 203, 16: 101, 24: 203, 33: 203, 48: 426, 431, 432, 433, 435, 436, 437, 438, 439, 440, 441, 442, 443, 444, 63: 203, 203, 203, 203, 203, 203, 203, 203, 203, 203, 203, 758, 759},
		// 45
		{97, 2: 343, 375, 379, 380, 381, 382, 383, 13: 324, 414, 308, 97, 326, 338, 341, 342, 327, 365, 25: 328, 376, 396, 345, 377, 329, 330, 331, 35: 304, 305, 306, 307, 378, 42: 344, 348, 347, 349, 350, 333, 62: 334, 78: 717, 96: 352, 531, 108: 719, 134: 721, 159: 720, 174: 718},
		{264, 41: 541, 99: 715, 150: 716},
		{2: 475, 14: 111, 114: 711, 126: 474, 175: 710},
		{2: 192, 14: 192},
		{175, 2: 175, 175, 175, 175, 175, 175, 175, 13: 175, 175, 175, 175, 175, 175, 175, 175, 175, 175, 35: 175, 175, 175, 175, 175},
		// 50
		{97, 2: 343, 375, 379, 380, 381, 382, 383, 13: 324, 414, 308, 97, 326, 338, 341, 342, 327, 365, 25: 328, 376, 396, 345, 377, 329, 330, 331, 35: 304, 305, 306, 307, 378, 42: 344, 348, 347, 349, 350, 333, 62: 334, 78: 325, 96: 352, 531, 108: 530, 131: 697},
		{16: 574, 41: 575, 106: 686},
		{13: 682},
		{2: 343, 375, 379, 380, 381, 382, 383, 13: 324, 414, 308, 17: 326, 338, 341, 342, 327, 365, 25: 328, 376, 396, 345, 377, 329, 330, 331, 35: 304, 305, 306, 307, 378, 42: 344, 348, 347, 349, 679, 333, 62: 678},
		{141, 141, 141, 141, 141, 141, 141, 141, 141, 141, 141, 141, 141, 141, 16: 141, 23: 141, 141, 33: 141, 141, 40: 141, 141, 48: 141, 141, 141, 141, 141, 141, 141, 141, 141, 141, 141, 141, 141, 141, 63: 141, 141, 141, 141, 141, 141, 141, 141, 141, 141, 141, 141, 141, 77: 141},
		// 55
		{2: 24, 11: 24, 24, 16: 574, 23: 24, 24, 33: 24, 24, 41: 575, 106: 675},
		{187, 187, 187, 187, 187, 187, 187, 187, 187, 187, 187, 12: 187, 187, 23: 673, 187, 33: 187, 41: 187, 48: 187, 187, 187, 187, 187, 187, 187, 187, 187, 187, 187, 187, 187, 187, 63: 187, 187, 187, 187, 187, 187, 187, 187, 187, 187, 187, 187, 187, 77: 187, 102: 415, 122: 416, 417},
		{126, 126, 126, 126, 126, 126, 126, 126, 126, 126, 126, 126, 126, 126, 16: 126, 23: 126, 126, 33: 126, 126, 40: 126, 126, 48: 126, 126, 126, 126, 126, 126, 126, 126, 126, 126, 126, 126, 126, 126, 63: 126, 126, 126, 126, 126, 126, 126, 126, 126, 126, 126, 126, 126, 77: 126},
		{16: 574, 41: 575, 106: 659, 125: 660},
		{9, 9, 614, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 617, 16: 9, 23: 9, 9, 33: 9, 9, 40: 9, 615, 48: 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 63: 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 77: 616, 118: 619, 151: 618},
		// 60
		{2: 504},
		{16: 533, 146: 613},
		{102, 102, 9: 102, 102, 16: 102},
		{179, 179, 9: 179, 179, 14: 299, 120: 612},
		{179, 179, 9: 179, 179, 14: 299, 120: 611},
		// 65
		{2: 343, 375, 379, 380, 381, 382, 383, 13: 324, 414, 308, 17: 326, 338, 341, 342, 327, 365, 25: 328, 376, 396, 345, 377, 329, 330, 331, 35: 304, 305, 306, 307, 378, 42: 344, 348, 347, 349, 350, 333, 62: 610},
		{82, 82, 9: 82, 82},
		{81, 81, 9: 81, 81},
		{2: 343, 375, 379, 380, 381, 382, 383, 13: 324, 414, 308, 17: 326, 338, 341, 342, 327, 365, 25: 328, 376, 396, 345, 377, 329, 330, 331, 35: 304, 305, 306, 307, 378, 42: 344, 348, 347, 349, 350, 333, 62: 609},
		{14: 608},
		// 70
		{77, 77, 9: 77, 77},
		{201, 201, 343, 375, 379, 380, 381, 382, 383, 201, 201, 13: 324, 414, 308, 17: 326, 338, 341, 342, 327, 365, 25: 328, 376, 396, 345, 377, 329, 330, 331, 35: 304, 305, 306, 307, 378, 42: 344, 348, 347, 349, 350, 333, 62: 412, 78: 606, 157: 607},
		{75, 75, 9: 75, 75},
		{74, 74, 9: 74, 74},
		{73, 73, 9: 73, 73},
		// 75
		{16: 574, 41: 575, 106: 576},
		{97, 2: 343, 375, 379, 380, 381, 382, 383, 13: 324, 414, 308, 97, 326, 338, 341, 342, 327, 365, 25: 328, 376, 396, 345, 377, 329, 330, 331, 35: 304, 305, 306, 307, 378, 42: 344, 348, 347, 349, 350, 333, 62: 334, 78: 325, 96: 352, 531, 108: 530, 131: 532},
		{48},
		{47},
		{46},
		// 80
		{45},
		{44},
		{43},
		{529},
		{2: 509, 14: 511, 135: 510},
		// 85
		{2: 343, 375, 379, 380, 381, 382, 383, 13: 324, 414, 308, 17: 326, 338, 341, 342, 327, 365, 25: 328, 376, 396, 345, 377, 329, 330, 331, 35: 304, 305, 306, 307, 378, 42: 344, 348, 347, 349, 502, 503},
		{2: 25, 11: 25, 25, 23: 25, 25, 33: 25, 25},
		{2: 23, 11: 23, 23, 23: 23, 23, 33: 23, 23},
		{2: 343, 375, 379, 380, 381, 382, 383, 13: 324, 414, 308, 17: 326, 338, 341, 342, 327, 365, 25: 328, 376, 396, 345, 377, 329, 330, 331, 35: 304, 305, 306, 307, 378, 42: 344, 348, 347, 349, 350, 501},
		{2: 343, 375, 379, 380, 381, 382, 383, 13: 324, 414, 308, 17: 326, 338, 341, 342, 327, 365, 25: 328, 376, 396, 345, 377, 329, 330, 331, 35: 304, 305, 306, 307, 378, 42: 344, 348, 347, 349, 350, 500},
		// 90
		{2: 343, 375, 379, 380, 381, 382, 383, 13: 324, 414, 308, 17: 326, 338, 341, 342, 327, 365, 25: 328, 376, 396, 345, 377, 329, 330, 331, 35: 304, 305, 306, 307, 378, 42: 344, 348, 347, 349, 350, 499},
		{2: 343, 375, 379, 380, 381, 382, 383, 13: 324, 414, 308, 17: 326, 338, 341, 342, 327, 365, 25: 328, 376, 396, 345, 377, 329, 330, 331, 35: 304, 305, 306, 307, 378, 42: 344, 348, 347, 349, 350, 498},
		{2: 343, 375, 379, 380, 381, 382, 383, 13: 324, 414, 308, 17: 326, 338, 341, 342, 327, 365, 25: 328, 376, 396, 345, 377, 329, 330, 331, 35: 304, 305, 306, 307, 378, 42: 344, 348, 347, 349, 350, 497},
		{2: 343, 375, 379, 380, 381, 382, 383, 13: 324, 414, 308, 17: 326, 338, 341, 342, 327, 365, 25: 328, 376, 396, 345, 377, 329, 330, 331, 35: 304, 305, 306, 307, 378, 42: 344, 348, 347, 349, 350, 496},
		{2: 387, 14: 385, 100: 386, 136: 388},
		// 95
		{177, 177, 177, 177, 9: 177, 177, 177, 177, 177, 177, 17: 177, 177, 177, 177, 177, 177, 24: 177},
		{2: 399, 400, 12: 397, 324, 398, 17: 326, 338, 341, 342, 327, 365, 24: 409, 401, 402, 396, 403, 404, 405, 407, 408, 79: 406, 81: 410},
		{11: 389, 14: 385, 100: 386, 136: 391, 182: 390},
		{6, 6, 9: 6, 6},
		{8, 8, 9: 8, 8},
		// 100
		{392, 11: 105, 115: 393},
		{2, 11: 2},
		{11: 104, 14: 385, 100: 386, 136: 395},
		{11: 394},
		{7, 7, 9: 7, 7},
		// 105
		{1, 11: 1},
		{2: 475, 114: 476, 126: 474},
		{14: 473},
		{138, 138, 138, 9: 138, 138, 138, 138, 15: 138, 138, 23: 138, 138, 33: 138, 138, 40: 138, 138, 76: 138, 471, 102: 138},
		{2: 399, 400, 13: 324, 398, 17: 326, 338, 341, 342, 327, 365, 25: 401, 402, 396, 403, 404, 405, 407, 408, 79: 406, 81: 469},
		// 110
		{2: 399, 400, 13: 324, 398, 17: 326, 338, 341, 342, 327, 365, 25: 401, 402, 396, 403, 404, 405, 407, 408, 79: 406, 81: 468},
		{38, 38, 38, 9: 38, 38, 38, 38, 15: 38, 38, 23: 38, 38, 33: 38, 38, 40: 38, 38, 76: 38},
		{37, 37, 37, 9: 37, 37, 37, 37, 15: 37, 37, 23: 37, 37, 33: 37, 37, 40: 37, 37, 76: 37},
		{36, 36, 36, 9: 36, 36, 36, 36, 15: 36, 36, 23: 36, 36, 33: 36, 36, 40: 36, 36, 76: 36},
		{35, 35, 35, 9: 35, 35, 35, 35, 15: 35, 35, 23: 35, 35, 33: 35, 35, 40: 35, 35, 76: 35},
		// 115
		{34, 34, 34, 9: 34, 34, 34, 34, 15: 34, 34, 23: 34, 34, 33: 34, 34, 40: 34, 34, 76: 34},
		{187, 187, 187, 9: 187, 187, 187, 187, 15: 187, 187, 23: 187, 187, 33: 187, 187, 40: 187, 187, 76: 187, 102: 415, 122: 416, 467},
		{32, 32, 32, 9: 32, 32, 32, 32, 15: 32, 32, 23: 32, 32, 33: 32, 32, 40: 32, 32, 76: 32},
		{31, 31, 31, 9: 31, 31, 31, 31, 15: 31, 31, 23: 31, 31, 33: 31, 31, 40: 31, 31, 76: 31},
		{2: 343, 375, 379, 380, 381, 382, 383, 13: 324, 414, 308, 17: 326, 338, 341, 342, 327, 365, 25: 328, 376, 396, 345, 377, 329, 330, 331, 35: 304, 305, 306, 307, 378, 42: 344, 348, 347, 349, 350, 333, 62: 412, 78: 466},
		// 120
		{4, 4, 9: 4, 4, 4, 24: 411},
		{2: 343, 375, 379, 380, 381, 382, 383, 13: 324, 414, 308, 17: 326, 338, 341, 342, 327, 365, 25: 328, 376, 396, 345, 377, 329, 330, 331, 35: 304, 305, 306, 307, 378, 42: 344, 348, 347, 349, 350, 333, 62: 412, 78: 413},
		{203, 203, 3: 428, 427, 429, 430, 434, 445, 203, 203, 203, 203, 16: 203, 48: 426, 431, 432, 433, 435, 436, 437, 438, 439, 440, 441, 442, 443, 444},
		{3, 3, 9: 3, 3, 3, 424},
		{187, 187, 187, 187, 187, 187, 187, 187, 187, 187, 187, 187, 187, 187, 16: 187, 23: 187, 187, 33: 187, 187, 40: 187, 187, 48: 187, 187, 187, 187, 187, 187, 187, 187, 187, 187, 187, 187, 187, 187, 63: 187, 187, 187, 187, 187, 187, 187, 187, 187, 187, 187, 187, 187, 77: 187, 102: 415, 122: 416, 417},
		// 125
		{2: 399, 400, 13: 324, 398, 17: 326, 338, 341, 342, 327, 365, 25: 401, 402, 396, 403, 404, 405, 407, 408, 79: 406, 81: 420, 139: 418, 162: 419},
		{186, 186, 186, 186, 186, 186, 186, 186, 186, 186, 186, 186, 186, 186, 15: 186, 186, 23: 186, 186, 33: 186, 186, 40: 186, 186, 48: 186, 186, 186, 186, 186, 186, 186, 186, 186, 186, 186, 186, 186, 186, 63: 186, 186, 186, 186, 186, 186, 186, 186, 186, 186, 186, 186, 186, 186, 186},
		{139, 139, 139, 139, 139, 139, 139, 139, 139, 139, 139, 139, 139, 139, 16: 139, 23: 139, 139, 33: 139, 139, 40: 139, 139, 48: 139, 139, 139, 139, 139, 139, 139, 139, 139, 139, 139, 139, 139, 139, 63: 139, 139, 139, 139, 139, 139, 139, 139, 139, 139, 139, 139, 139, 77: 139},
		{12: 191, 76: 191},
		{12: 421, 76: 422},
		// 130
		{12: 189, 76: 189},
		{2: 399, 400, 13: 324, 398, 17: 326, 338, 341, 342, 327, 365, 25: 401, 402, 396, 403, 404, 405, 407, 408, 79: 406, 81: 420, 139: 423},
		{188, 188, 188, 188, 188, 188, 188, 188, 188, 188, 188, 188, 188, 188, 15: 188, 188, 23: 188, 188, 33: 188, 188, 40: 188, 188, 48: 188, 188, 188, 188, 188, 188, 188, 188, 188, 188, 188, 188, 188, 188, 63: 188, 188, 188, 188, 188, 188, 188, 188, 188, 188, 188, 188, 188, 188, 188},
		{12: 190, 76: 190},
		{2: 343, 375, 379, 380, 381, 382, 383, 13: 324, 414, 308, 17: 326, 338, 341, 342, 327, 365, 25: 328, 376, 396, 345, 377, 329, 330, 331, 35: 304, 305, 306, 307, 378, 42: 344, 348, 347, 349, 350, 333, 62: 425},
		// 135
		{202, 202, 3: 428, 427, 429, 430, 434, 445, 202, 202, 202, 202, 16: 202, 24: 202, 33: 202, 48: 426, 431, 432, 433, 435, 436, 437, 438, 439, 440, 441, 442, 443, 444, 63: 202, 202, 202, 202, 202, 202, 202, 202, 202, 202, 202},
		{2: 343, 375, 379, 380, 381, 382, 383, 13: 324, 414, 308, 17: 326, 338, 341, 342, 327, 365, 25: 328, 376, 396, 345, 377, 329, 330, 331, 35: 304, 305, 306, 307, 378, 42: 344, 348, 347, 349, 350, 333, 62: 465},
		{2: 343, 375, 379, 380, 381, 382, 383, 13: 324, 414, 308, 17: 326, 338, 341, 342, 327, 365, 25: 328, 376, 396, 345, 377, 329, 330, 331, 35: 304, 305, 306, 307, 378, 42: 344, 348, 347, 349, 350, 333, 62: 464},
		{2: 343, 375, 379, 380, 381, 382, 383, 13: 324, 414, 308, 17: 326, 338, 341, 342, 327, 365, 25: 328, 376, 396, 345, 377, 329, 330, 331, 35: 304, 305, 306, 307, 378, 42: 344, 348, 347, 349, 350, 333, 62: 463},
		{2: 343, 375, 379, 380, 381, 382, 383, 13: 324, 414, 308, 17: 326, 338, 341, 342, 327, 365, 25: 328, 376, 396, 345, 377, 329, 330, 331, 35: 304, 305, 306, 307, 378, 42: 344, 348, 347, 349, 350, 333, 62: 462},
		// 140
		{2: 343, 375, 379, 380, 381, 382, 383, 13: 324, 414, 308, 17: 326, 338, 341, 342, 327, 365, 25: 328, 376, 396, 345, 377, 329, 330, 331, 35: 304, 305, 306, 307, 378, 42: 344, 348, 347, 349, 350, 333, 62: 461},
		{2: 343, 375, 379, 380, 381, 382, 383, 13: 324, 414, 308, 17: 326, 338, 341, 342, 327, 365, 25: 328, 376, 396, 345, 377, 329, 330, 331, 35: 304, 305, 306, 307, 378, 42: 344, 348, 347, 349, 350, 333, 62: 460},
		{2: 343, 375, 379, 380, 381, 382, 383, 13: 324, 414, 308, 17: 326, 338, 341, 342, 327, 365, 25: 328, 376, 396, 345, 377, 329, 330, 331, 35: 304, 305, 306, 307, 378, 42: 344, 348, 347, 349, 350, 333, 62: 459},
		{2: 343, 375, 379, 380, 381, 382, 383, 13: 324, 414, 308, 17: 326, 338, 341, 342, 327, 365, 25: 328, 376, 396, 345, 377, 329, 330, 331, 35: 304, 305, 306, 307, 378, 42: 344, 348, 347, 349, 350, 333, 62: 458},
		{2: 343, 375, 379, 380, 381, 382, 383, 13: 324, 414, 308, 17: 326, 338, 341, 342, 327, 365, 25: 328, 376, 396, 345, 377, 329, 330, 331, 35: 304, 305, 306, 307, 378, 42: 344, 348, 347, 349, 350, 333, 62: 457},
		// 145
		{2: 343, 375, 379, 380, 381, 382, 383, 13: 324, 414, 308, 17: 326, 338, 341, 342, 327, 365, 25: 328, 376, 396, 345, 377, 329, 330, 331, 35: 304, 305, 306, 307, 378, 42: 344, 348, 347, 349, 350, 333, 62: 456},
		{2: 343, 375, 379, 380, 381, 382, 383, 13: 324, 414, 308, 17: 326, 338, 341, 342, 327, 365, 25: 328, 376, 396, 345, 377, 329, 330, 331, 35: 304, 305, 306, 307, 378, 42: 344, 348, 347, 349, 350, 333, 62: 455},
		{2: 343, 375, 379, 380, 381, 382, 383, 13: 324, 414, 308, 17: 326, 338, 341, 342, 327, 365, 25: 328, 376, 396, 345, 377, 329, 330, 331, 35: 304, 305, 306, 307, 378, 42: 344, 348, 347, 349, 350, 333, 62: 454},
		{2: 343, 375, 379, 380, 381, 382, 383, 13: 324, 414, 308, 17: 326, 338, 341, 342, 327, 365, 25: 328, 376, 396, 345, 377, 329, 330, 331, 35: 304, 305, 306, 307, 378, 42: 344, 348, 347, 349, 350, 333, 62: 453},
		{2: 343, 375, 379, 380, 381, 382, 383, 13: 324, 414, 308, 17: 326, 338, 341, 342, 327, 365, 25: 328, 376, 396, 345, 377, 329, 330, 331, 35: 304, 305, 306, 307, 378, 42: 344, 348, 347, 349, 350, 333, 62: 452},
		// 150
		{2: 343, 375, 379, 380, 381, 382, 383, 13: 324, 414, 308, 17: 326, 338, 341, 342, 327, 365, 25: 328, 376, 396, 345, 377, 329, 330, 331, 35: 304, 305, 306, 307, 378, 42: 344, 348, 347, 349, 350, 333, 62: 451},
		{2: 343, 375, 379, 380, 381, 382, 383, 13: 324, 414, 308, 17: 326, 338, 341, 342, 327, 365, 25: 328, 376, 396, 345, 377, 329, 330, 331, 35: 304, 305, 306, 307, 378, 42: 344, 348, 347, 349, 350, 333, 62: 450},
		{2: 343, 375, 379, 380, 381, 382, 383, 13: 324, 414, 308, 17: 326, 338, 341, 342, 327, 365, 25: 328, 376, 396, 345, 377, 329, 330, 331, 35: 304, 305, 306, 307, 378, 42: 344, 348, 347, 349, 350, 333, 62: 449},
		{2: 343, 375, 379, 380, 381, 382, 383, 13: 324, 414, 308, 17: 326, 338, 341, 342, 327, 365, 25: 328, 376, 396, 345, 377, 329, 330, 331, 35: 304, 305, 306, 307, 378, 42: 344, 348, 347, 349, 350, 333, 62: 448},
		{2: 343, 375, 379, 380, 381, 382, 383, 13: 324, 414, 308, 17: 326, 338, 341, 342, 327, 365, 25: 328, 376, 396, 345, 377, 329, 330, 331, 35: 304, 305, 306, 307, 378, 42: 344, 348, 347, 349, 350, 333, 62: 447},
		// 155
		{2: 343, 375, 379, 380, 381, 382, 383, 13: 324, 414, 308, 17: 326, 338, 341, 342, 327, 365, 25: 328, 376, 396, 345, 377, 329, 330, 331, 35: 304, 305, 306, 307, 378, 42: 344, 348, 347, 349, 350, 333, 62: 446},
		{206, 206, 3: 428, 427, 429, 430, 434, 206, 206, 206, 206, 206, 16: 206, 23: 206, 206, 33: 206, 206, 40: 206, 48: 426, 431, 432, 433, 435, 436, 437, 438, 439, 440, 441, 442, 443, 444, 63: 206, 206, 206, 206, 206, 206, 206, 206, 206, 206, 206, 206, 206},
		{207, 207, 3: 207, 207, 207, 207, 207, 207, 207, 207, 207, 207, 16: 207, 23: 207, 207, 33: 207, 207, 40: 207, 48: 207, 207, 207, 207, 207, 207, 207, 207, 207, 207, 207, 207, 207, 207, 63: 207, 207, 207, 207, 207, 207, 207, 207, 207, 207, 207, 207, 207},
		{208, 208, 3: 428, 427, 429, 430, 434, 208, 208, 208, 208, 208, 16: 208, 23: 208, 208, 33: 208, 208, 40: 208, 48: 426, 431, 432, 433, 435, 436, 437, 438, 439, 440, 441, 442, 208, 444, 63: 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208},
		{209, 209, 3: 428, 427, 429, 430, 434, 209, 209, 209, 209, 209, 16: 209, 23: 209, 209, 33: 209, 209, 40: 209, 48: 426, 431, 209, 209, 435, 209, 437, 209, 209, 209, 441, 209, 209, 444, 63: 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209},
		// 160
		{210, 210, 3: 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 16: 210, 23: 210, 210, 33: 210, 210, 40: 210, 48: 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 63: 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210},
		{211, 211, 3: 428, 427, 429, 430, 434, 211, 211, 211, 211, 211, 16: 211, 23: 211, 211, 33: 211, 211, 40: 211, 48: 426, 431, 211, 211, 435, 211, 437, 211, 211, 211, 441, 211, 211, 444, 63: 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211},
		{212, 212, 3: 428, 427, 429, 430, 434, 212, 212, 212, 212, 212, 16: 212, 23: 212, 212, 33: 212, 212, 40: 212, 48: 426, 431, 212, 212, 435, 212, 437, 212, 212, 212, 441, 212, 212, 444, 63: 212, 212, 212, 212, 212, 212, 212, 212, 212, 212, 212, 212, 212},
		{213, 213, 3: 428, 427, 429, 430, 434, 213, 213, 213, 213, 213, 16: 213, 23: 213, 213, 33: 213, 213, 40: 213, 48: 426, 431, 213, 213, 435, 213, 437, 213, 213, 213, 441, 213, 213, 444, 63: 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213},
		{214, 214, 3: 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 16: 214, 23: 214, 214, 33: 214, 214, 40: 214, 48: 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 63: 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214},
		// 165
		{215, 215, 3: 428, 427, 429, 430, 434, 215, 215, 215, 215, 215, 16: 215, 23: 215, 215, 33: 215, 215, 40: 215, 48: 426, 431, 432, 433, 435, 215, 437, 438, 439, 440, 441, 442, 215, 444, 63: 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215},
		{216, 216, 3: 428, 427, 216, 216, 216, 216, 216, 216, 216, 216, 16: 216, 23: 216, 216, 33: 216, 216, 40: 216, 48: 426, 431, 216, 216, 216, 216, 437, 216, 216, 216, 441, 216, 216, 444, 63: 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216},
		{217, 217, 3: 428, 427, 217, 217, 217, 217, 217, 217, 217, 217, 16: 217, 23: 217, 217, 33: 217, 217, 40: 217, 48: 426, 431, 217, 217, 217, 217, 437, 217, 217, 217, 441, 217, 217, 444, 63: 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 217},
		{218, 218, 3: 428, 427, 429, 430, 434, 218, 218, 218, 218, 218, 16: 218, 23: 218, 218, 33: 218, 218, 40: 218, 48: 426, 431, 218, 218, 435, 218, 437, 218, 218, 218, 441, 218, 218, 444, 63: 218, 218, 218, 218, 218, 218, 218, 218, 218, 218, 218, 218, 218},
		{219, 219, 3: 428, 427, 429, 430, 434, 219, 219, 219, 219, 219, 16: 219, 23: 219, 219, 33: 219, 219, 40: 219, 48: 426, 431, 219, 219, 435, 219, 437, 219, 219, 219, 441, 219, 219, 444, 63: 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219},
		// 170
		{220, 220, 3: 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 16: 220, 23: 220, 220, 33: 220, 220, 40: 220, 48: 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 63: 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 220},
		{221, 221, 3: 428, 427, 221, 221, 221, 221, 221, 221, 221, 221, 16: 221, 23: 221, 221, 33: 221, 221, 40: 221, 48: 426, 431, 221, 221, 221, 221, 437, 221, 221, 221, 441, 221, 221, 444, 63: 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221},
		{222, 222, 3: 428, 427, 222, 222, 222, 222, 222, 222, 222, 222, 16: 222, 23: 222, 222, 33: 222, 222, 40: 222, 48: 426, 431, 222, 222, 222, 222, 437, 222, 222, 222, 441, 222, 222, 444, 63: 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222},
		{223, 223, 3: 223, 223, 223, 223, 223, 223, 223, 223, 223, 223, 16: 223, 23: 223, 223, 33: 223, 223, 40: 223, 48: 223, 223, 223, 223, 223, 223, 223, 223, 223, 223, 223, 223, 223, 223, 63: 223, 223, 223, 223, 223, 223, 223, 223, 223, 223, 223, 223, 223},
		{224, 224, 3: 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 16: 224, 23: 224, 224, 33: 224, 224, 40: 224, 48: 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 63: 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224},
		// 175
		{225, 225, 3: 225, 225, 225, 225, 225, 225, 225, 225, 225, 225, 16: 225, 23: 225, 225, 33: 225, 225, 40: 225, 48: 225, 225, 225, 225, 225, 225, 225, 225, 225, 225, 225, 225, 225, 225, 63: 225, 225, 225, 225, 225, 225, 225, 225, 225, 225, 225, 225, 225},
		{5, 5, 9: 5, 5, 5, 424},
		{33, 33, 33, 9: 33, 33, 33, 33, 15: 33, 33, 23: 33, 33, 33: 33, 33, 40: 33, 33, 76: 33},
		{39, 39, 39, 9: 39, 39, 39, 39, 15: 39, 39, 23: 39, 39, 33: 39, 39, 40: 39, 39, 76: 39},
		{11: 470},
		// 180
		{40, 40, 40, 9: 40, 40, 40, 40, 15: 40, 40, 23: 40, 40, 33: 40, 40, 40: 40, 40, 76: 40},
		{14: 472},
		{137, 137, 137, 9: 137, 137, 137, 137, 15: 137, 137, 23: 137, 137, 33: 137, 137, 40: 137, 137, 76: 137, 102: 137},
		{176, 176, 176, 176, 9: 176, 176, 176, 176, 176, 176, 17: 176, 176, 176, 176, 176, 176, 24: 176},
		{193, 193, 193, 9: 193, 193, 193, 193, 15: 193, 193, 23: 193, 193, 33: 193, 193, 40: 193, 193, 76: 193},
		// 185
		{2: 399, 400, 11: 486, 13: 324, 481, 17: 326, 338, 341, 342, 327, 365, 25: 401, 402, 396, 403, 404, 405, 407, 408, 34: 482, 79: 406, 81: 490, 133: 484, 143: 485},
		{109, 109, 477, 400, 9: 109, 109, 109, 109, 324, 398, 109, 109, 326, 338, 341, 342, 327, 365, 109, 109, 401, 402, 396, 403, 404, 405, 407, 408, 109, 109, 40: 109, 109, 76: 109, 79: 406, 81: 479, 114: 478, 144: 480},
		{2: 399, 400, 11: 486, 13: 324, 481, 17: 326, 338, 341, 342, 327, 365, 25: 401, 402, 396, 403, 404, 405, 407, 408, 34: 482, 79: 406, 81: 483, 133: 484, 143: 485},
		{108, 108, 108, 9: 108, 108, 108, 108, 15: 108, 108, 23: 108, 108, 33: 108, 108, 40: 108, 108, 76: 108},
		{107, 107, 107, 9: 107, 107, 107, 107, 15: 107, 107, 23: 107, 107, 33: 107, 107, 40: 107, 107, 76: 107},
		// 190
		{103, 103, 103, 9: 103, 103, 103, 103, 15: 103, 103, 23: 103, 103, 33: 103, 103, 40: 103, 103, 76: 103},
		{2: 399, 400, 11: 138, 138, 324, 398, 17: 326, 338, 341, 342, 327, 365, 25: 401, 402, 396, 403, 404, 405, 407, 408, 34: 493, 77: 471, 79: 406, 81: 494, 102: 138},
		{2: 399, 400, 13: 324, 398, 17: 326, 338, 341, 342, 327, 365, 25: 401, 402, 396, 403, 404, 405, 407, 408, 79: 406, 81: 492},
		{11: 470, 131},
		{11: 130, 130},
		// 195
		{11: 255, 487, 113: 488},
		{128, 128, 128, 128, 9: 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 33: 128, 128, 40: 128, 128, 76: 128},
		{2: 399, 400, 11: 254, 13: 324, 481, 17: 326, 338, 341, 342, 327, 365, 25: 401, 402, 396, 403, 404, 405, 407, 408, 34: 482, 79: 406, 81: 490, 133: 491},
		{11: 489},
		{127, 127, 127, 127, 9: 127, 127, 127, 127, 127, 127, 127, 127, 127, 127, 127, 127, 127, 127, 127, 127, 33: 127, 127, 40: 127, 127, 76: 127},
		// 200
		{11: 131, 131},
		{11: 129, 129},
		{11: 134, 134},
		{2: 399, 400, 13: 324, 398, 17: 326, 338, 341, 342, 327, 365, 25: 401, 402, 396, 403, 404, 405, 407, 408, 79: 406, 81: 495},
		{11: 132, 132},
		// 205
		{11: 133, 133},
		{10, 10, 3: 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 16: 10, 23: 10, 10, 33: 10, 10, 40: 10, 48: 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 63: 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10},
		{11, 11, 3: 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 16: 11, 23: 11, 11, 33: 11, 11, 40: 11, 48: 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 63: 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11},
		{12, 12, 3: 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 16: 12, 23: 12, 12, 33: 12, 12, 40: 12, 48: 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 63: 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12},
		{13, 13, 3: 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 16: 13, 23: 13, 13, 33: 13, 13, 40: 13, 48: 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 63: 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13},
		// 210
		{15, 15, 3: 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 16: 15, 23: 15, 15, 33: 15, 15, 40: 15, 48: 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 63: 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15},
		{16, 16, 3: 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16: 16, 23: 16, 16, 33: 16, 16, 40: 16, 48: 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 63: 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16},
		{2: 504, 11: 27, 27, 23: 27, 27, 33: 27, 27},
		{14, 14, 3: 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 16: 14, 23: 14, 14, 33: 14, 14, 40: 14, 48: 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 63: 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14},
		{2: 343, 375, 379, 380, 381, 382, 383, 13: 324, 414, 308, 17: 326, 338, 341, 342, 327, 365, 25: 328, 376, 396, 345, 377, 329, 330, 331, 35: 304, 305, 306, 307, 378, 42: 344, 348, 347, 349, 350, 333, 62: 505},
		// 215
		{3: 428, 427, 429, 430, 434, 445, 11: 255, 506, 48: 426, 431, 432, 433, 435, 436, 437, 438, 439, 440, 441, 442, 443, 444, 113: 507},
		{11: 254},
		{11: 508},
		{116, 116, 116, 116, 116, 116, 116, 116, 116, 116, 116, 116, 116, 116, 16: 116, 23: 116, 116, 33: 116, 116, 40: 116, 116, 48: 116, 116, 116, 116, 116, 116, 116, 116, 116, 116, 116, 116, 116, 116, 63: 116, 116, 116, 116, 116, 116, 116, 116, 116, 116, 116, 116, 116, 77: 116},
		{11: 522, 14: 511, 135: 524, 181: 523},
		// 220
		{28, 28, 9: 28, 28},
		{2: 181, 181, 13: 181, 181, 17: 181, 181, 181, 181, 181, 181, 102: 512, 164: 513, 514},
		{14: 518, 140: 516, 163: 517},
		{2: 180, 180, 13: 180, 180, 17: 180, 180, 180, 180, 180, 180},
		{2: 399, 400, 13: 324, 398, 17: 326, 338, 341, 342, 327, 365, 25: 401, 402, 396, 403, 404, 405, 407, 408, 79: 406, 81: 515},
		// 225
		{19, 19, 9: 19, 19, 19},
		{12: 185, 76: 185},
		{12: 519, 76: 520},
		{12: 183, 76: 183},
		{14: 518, 140: 521},
		// 230
		{2: 182, 182, 13: 182, 182, 17: 182, 182, 182, 182, 182, 182},
		{12: 184, 76: 184},
		{30, 30, 9: 30, 30},
		{525, 11: 105, 115: 526},
		{18, 11: 18},
		// 235
		{11: 104, 14: 511, 135: 528},
		{11: 527},
		{29, 29, 9: 29, 29},
		{17, 11: 17},
		{2: 41, 41, 41, 41, 41, 41, 41, 13: 41, 41, 41, 17: 41, 41, 41, 41, 41, 41, 35: 41, 41, 41, 41, 41, 80: 41, 82: 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 98: 41},
		// 240
		{572, 16: 174},
		{96, 16: 96},
		{16: 533, 146: 534},
		{1: 535, 9: 537, 538, 147: 539, 540, 177: 536},
		{49, 49, 9: 49, 49},
		// 245
		{60, 60, 9: 60, 60},
		{1: 570, 9: 537, 538, 147: 539, 571},
		{2: 343, 375, 379, 380, 381, 382, 383, 13: 324, 414, 308, 17: 326, 338, 341, 342, 327, 365, 25: 328, 376, 396, 345, 377, 329, 330, 331, 35: 304, 305, 306, 307, 378, 42: 344, 348, 347, 349, 557, 333, 62: 556, 80: 560, 121: 558, 137: 559},
		{23: 554, 80: 555},
		{94, 94, 343, 375, 379, 380, 381, 382, 383, 94, 94, 13: 324, 346, 308, 17: 326, 338, 341, 342, 327, 365, 25: 328, 376, 396, 345, 377, 329, 330, 331, 35: 304, 305, 306, 307, 378, 41: 541, 344, 348, 347, 349, 350, 333, 62: 334, 78: 325, 80: 547, 82: 339, 374, 353, 332, 354, 355, 356, 335, 358, 359, 361, 351, 366, 384, 352, 363, 99: 542, 101: 340, 103: 543, 357, 360, 107: 362, 109: 546, 364, 544, 545, 116: 548, 127: 549},
		// 250
		{1: 51, 9: 51, 51},
		{94, 94, 343, 375, 379, 380, 381, 382, 383, 13: 324, 346, 308, 17: 326, 338, 341, 342, 327, 365, 25: 328, 376, 396, 345, 377, 329, 330, 331, 35: 304, 305, 306, 307, 378, 41: 541, 344, 348, 347, 349, 350, 333, 62: 334, 78: 325, 80: 547, 82: 339, 374, 353, 332, 354, 355, 356, 335, 358, 359, 361, 351, 366, 384, 352, 363, 99: 542, 101: 340, 103: 543, 357, 360, 107: 362, 109: 546, 364, 544, 545, 116: 548, 127: 552},
		{93, 93, 9: 93, 93},
		{92, 92, 9: 92, 92},
		{91, 91, 9: 91, 91},
		// 255
		{90, 90, 9: 90, 90},
		{89, 89, 9: 89, 89},
		{88, 88, 9: 88, 88},
		{87, 87, 9: 87, 87},
		{550, 52, 9: 52, 52},
		// 260
		{94, 94, 343, 375, 379, 380, 381, 382, 383, 94, 94, 13: 324, 346, 308, 17: 326, 338, 341, 342, 327, 365, 25: 328, 376, 396, 345, 377, 329, 330, 331, 35: 304, 305, 306, 307, 378, 41: 541, 344, 348, 347, 349, 350, 333, 62: 334, 78: 325, 80: 547, 82: 339, 374, 353, 332, 354, 355, 356, 335, 358, 359, 361, 351, 366, 384, 352, 363, 99: 542, 101: 340, 103: 543, 357, 360, 107: 362, 109: 546, 364, 544, 545, 116: 551},
		{86, 86, 9: 86, 86},
		{550, 553},
		{265, 265, 9: 265, 265},
		{55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 13: 55, 55, 55, 17: 55, 55, 55, 55, 55, 55, 35: 55, 55, 55, 55, 55, 41: 55, 80: 55, 82: 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55},
		// 265
		{53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 13: 53, 53, 53, 17: 53, 53, 53, 53, 53, 53, 35: 53, 53, 53, 53, 53, 41: 53, 80: 53, 82: 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53},
		{3: 428, 427, 429, 430, 434, 445, 11: 288, 288, 23: 288, 288, 33: 288, 288, 48: 426, 431, 432, 433, 435, 436, 437, 438, 439, 440, 441, 442, 443, 444},
		{2: 504, 11: 287, 287, 23: 287, 287, 33: 287, 287},
		{11: 286, 286, 23: 286, 286, 33: 286, 286},
		{12: 561, 23: 562, 563, 33: 564},
		// 270
		{54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 13: 54, 54, 54, 17: 54, 54, 54, 54, 54, 54, 35: 54, 54, 54, 54, 54, 41: 54, 80: 54, 82: 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 54},
		{2: 343, 375, 379, 380, 381, 382, 383, 13: 324, 414, 308, 17: 326, 338, 341, 342, 327, 365, 25: 328, 376, 396, 345, 377, 329, 330, 331, 35: 304, 305, 306, 307, 378, 42: 344, 348, 347, 349, 557, 333, 62: 556, 121: 569},
		{58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 13: 58, 58, 58, 17: 58, 58, 58, 58, 58, 58, 35: 58, 58, 58, 58, 58, 41: 58, 80: 58, 82: 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58},
		{2: 343, 375, 379, 380, 381, 382, 383, 13: 324, 414, 308, 17: 326, 338, 341, 342, 327, 365, 25: 328, 376, 396, 345, 377, 329, 330, 331, 35: 304, 305, 306, 307, 378, 42: 344, 348, 347, 349, 350, 333, 62: 567},
		{2: 343, 375, 379, 380, 381, 382, 383, 13: 324, 414, 308, 17: 326, 338, 341, 342, 327, 365, 25: 328, 376, 396, 345, 377, 329, 330, 331, 35: 304, 305, 306, 307, 378, 42: 344, 348, 347, 349, 350, 333, 62: 565},
		// 275
		{3: 428, 427, 429, 430, 434, 445, 23: 566, 48: 426, 431, 432, 433, 435, 436, 437, 438, 439, 440, 441, 442, 443, 444},
		{56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 13: 56, 56, 56, 17: 56, 56, 56, 56, 56, 56, 35: 56, 56, 56, 56, 56, 41: 56, 80: 56, 82: 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56},
		{3: 428, 427, 429, 430, 434, 445, 23: 568, 48: 426, 431, 432, 433, 435, 436, 437, 438, 439, 440, 441, 442, 443, 444},
		{57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 13: 57, 57, 57, 17: 57, 57, 57, 57, 57, 57, 35: 57, 57, 57, 57, 57, 41: 57, 80: 57, 82: 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57},
		{11: 285, 285, 23: 285, 285, 33: 285, 285},
		// 280
		{59, 59, 9: 59, 59},
		{1: 50, 9: 50, 50},
		{2: 343, 375, 379, 380, 381, 382, 383, 13: 324, 414, 308, 97, 326, 338, 341, 342, 327, 365, 25: 328, 376, 396, 345, 377, 329, 330, 331, 35: 304, 305, 306, 307, 378, 42: 344, 348, 347, 349, 350, 333, 62: 334, 78: 325, 96: 352, 531, 108: 573},
		{16: 173},
		{154, 154, 154, 154, 154, 154, 154, 154, 154, 13: 154, 154, 154, 154, 154, 154, 154, 154, 154, 154, 35: 154, 154, 154, 154, 154, 41: 154, 80: 154, 82: 154, 154, 154, 154, 154, 154, 154, 154, 154, 154, 154, 154, 154, 154},
		// 285
		{153, 153, 153, 153, 153, 153, 153, 153, 153, 13: 153, 153, 153, 153, 153, 153, 153, 153, 153, 153, 35: 153, 153, 153, 153, 153, 41: 153, 80: 153, 82: 153, 153, 153, 153, 153, 153, 153, 153, 153, 153, 153, 153, 153, 153},
		{1: 584, 581, 579, 14: 577, 79: 580, 100: 578, 145: 582, 176: 583},
		{138, 138, 177, 177, 12: 177, 177, 177, 138, 17: 177, 177, 177, 177, 177, 177, 77: 471},
		{2: 399, 400, 12: 397, 324, 398, 17: 326, 338, 341, 342, 327, 365, 25: 401, 402, 396, 403, 404, 405, 407, 408, 79: 406, 81: 604},
		{2: 599, 14: 398, 79: 598},
		// 290
		{72, 72, 15: 593, 117: 597},
		{3: 590, 14: 398, 79: 589},
		{64, 64},
		{585, 105, 115: 586},
		{62, 62, 62, 9: 62, 62, 62, 62, 15: 62, 62, 23: 62, 62, 33: 62, 62, 40: 62, 62, 76: 62},
		// 295
		{1: 104, 581, 579, 14: 577, 79: 580, 100: 578, 145: 588},
		{1: 587},
		{61, 61, 61, 9: 61, 61, 61, 61, 15: 61, 61, 23: 61, 61, 33: 61, 61, 40: 61, 61, 76: 61},
		{63, 63},
		{11: 595},
		// 300
		{14: 398, 79: 591},
		{11: 592},
		{72, 72, 15: 593, 117: 594},
		{71, 71},
		{66, 66},
		// 305
		{72, 72, 15: 593, 117: 596},
		{67, 67},
		{68, 68},
		{72, 72, 15: 593, 117: 603},
		{14: 398, 79: 600},
		// 310
		{11: 601},
		{72, 72, 15: 593, 117: 602},
		{65, 65},
		{70, 70},
		{72, 72, 15: 593, 117: 605},
		// 315
		{69, 69},
		{200, 200, 9: 200, 200, 12: 424},
		{76, 76, 9: 76, 76},
		{79, 79, 9: 79, 79},
		{80, 80, 3: 428, 427, 429, 430, 434, 445, 80, 80, 48: 426, 431, 432, 433, 435, 436, 437, 438, 439, 440, 441, 442, 443, 444},
		// 320
		{83, 83, 3: 428, 427, 429, 430, 434, 445, 83, 83, 48: 426, 431, 432, 433, 435, 436, 437, 438, 439, 440, 441, 442, 443, 444},
		{84, 84, 9: 84, 84},
		{85, 85, 9: 85, 85},
		{106, 106, 9: 106, 106},
		{2: 343, 375, 379, 380, 381, 382, 383, 11: 652, 13: 324, 414, 308, 17: 326, 338, 341, 342, 327, 365, 25: 328, 376, 396, 345, 377, 329, 330, 331, 35: 304, 305, 306, 307, 378, 42: 344, 348, 347, 349, 557, 333, 62: 556, 121: 558, 137: 651},
		// 325
		{1: 640, 343, 375, 379, 380, 381, 382, 383, 13: 324, 414, 308, 17: 326, 338, 341, 342, 327, 365, 25: 328, 376, 396, 345, 377, 329, 330, 331, 35: 304, 305, 306, 307, 378, 41: 615, 344, 348, 347, 349, 350, 333, 62: 637, 118: 636, 138: 638, 152: 639},
		{2: 630, 14: 631},
		{2: 343, 375, 379, 380, 381, 382, 383, 13: 324, 414, 308, 17: 326, 338, 341, 342, 327, 365, 205, 25: 328, 376, 396, 345, 377, 329, 330, 331, 35: 304, 305, 306, 307, 378, 42: 344, 348, 347, 349, 350, 333, 62: 620, 130: 621},
		{118, 118, 118, 118, 118, 118, 118, 118, 118, 118, 118, 118, 118, 118, 16: 118, 23: 118, 118, 33: 118, 118, 40: 118, 118, 48: 118, 118, 118, 118, 118, 118, 118, 118, 118, 118, 118, 118, 118, 118, 63: 118, 118, 118, 118, 118, 118, 118, 118, 118, 118, 118, 118, 118, 77: 118},
		{117, 117, 117, 117, 117, 117, 117, 117, 117, 117, 117, 117, 117, 117, 16: 117, 23: 117, 117, 33: 117, 117, 40: 117, 117, 48: 117, 117, 117, 117, 117, 117, 117, 117, 117, 117, 117, 117, 117, 117, 63: 117, 117, 117, 117, 117, 117, 117, 117, 117, 117, 117, 117, 117, 77: 117},
		// 330
		{3: 428, 427, 429, 430, 434, 445, 23: 204, 40: 629, 48: 426, 431, 432, 433, 435, 436, 437, 438, 439, 440, 441, 442, 443, 444},
		{23: 622},
		{2: 343, 375, 379, 380, 381, 382, 383, 13: 324, 414, 308, 17: 326, 338, 341, 342, 327, 365, 205, 25: 328, 376, 396, 345, 377, 329, 330, 331, 35: 304, 305, 306, 307, 378, 205, 42: 344, 348, 347, 349, 350, 333, 62: 623, 130: 624},
		{3: 428, 427, 429, 430, 434, 445, 23: 204, 40: 204, 48: 426, 431, 432, 433, 435, 436, 437, 438, 439, 440, 441, 442, 443, 444},
		{23: 625, 40: 626},
		// 335
		{2: 343, 375, 379, 380, 381, 382, 383, 13: 324, 414, 308, 17: 326, 338, 341, 342, 327, 365, 25: 328, 376, 396, 345, 377, 329, 330, 331, 35: 304, 305, 306, 307, 378, 205, 42: 344, 348, 347, 349, 350, 333, 62: 623, 130: 627},
		{119, 119, 119, 119, 119, 119, 119, 119, 119, 119, 119, 119, 119, 119, 16: 119, 23: 119, 119, 33: 119, 119, 40: 119, 119, 48: 119, 119, 119, 119, 119, 119, 119, 119, 119, 119, 119, 119, 119, 119, 63: 119, 119, 119, 119, 119, 119, 119, 119, 119, 119, 119, 119, 119, 77: 119},
		{40: 628},
		{120, 120, 120, 120, 120, 120, 120, 120, 120, 120, 120, 120, 120, 120, 16: 120, 23: 120, 120, 33: 120, 120, 40: 120, 120, 48: 120, 120, 120, 120, 120, 120, 120, 120, 120, 120, 120, 120, 120, 120, 63: 120, 120, 120, 120, 120, 120, 120, 120, 120, 120, 120, 120, 120, 77: 120},
		{121, 121, 121, 121, 121, 121, 121, 121, 121, 121, 121, 121, 121, 121, 16: 121, 23: 121, 121, 33: 121, 121, 40: 121, 121, 48: 121, 121, 121, 121, 121, 121, 121, 121, 121, 121, 121, 121, 121, 121, 63: 121, 121, 121, 121, 121, 121, 121, 121, 121, 121, 121, 121, 121, 77: 121},
		// 340
		{2: 399, 400, 13: 324, 398, 17: 326, 338, 341, 342, 327, 365, 25: 401, 402, 396, 403, 404, 405, 407, 408, 79: 406, 81: 633, 83: 632},
		{122, 122, 122, 122, 122, 122, 122, 122, 122, 122, 122, 122, 122, 122, 16: 122, 23: 122, 122, 33: 122, 122, 40: 122, 122, 48: 122, 122, 122, 122, 122, 122, 122, 122, 122, 122, 122, 122, 122, 122, 63: 122, 122, 122, 122, 122, 122, 122, 122, 122, 122, 122, 122, 122, 77: 122},
		{11: 635},
		{11: 634},
		{123, 123, 123, 123, 123, 123, 123, 123, 123, 123, 123, 123, 123, 123, 16: 123, 23: 123, 123, 33: 123, 123, 40: 123, 123, 48: 123, 123, 123, 123, 123, 123, 123, 123, 123, 123, 123, 123, 123, 123, 63: 123, 123, 123, 123, 123, 123, 123, 123, 123, 123, 123, 123, 123, 77: 123},
		// 345
		{124, 124, 124, 124, 124, 124, 124, 124, 124, 124, 124, 124, 124, 124, 16: 124, 23: 124, 124, 33: 124, 124, 40: 124, 124, 48: 124, 124, 124, 124, 124, 124, 124, 124, 124, 124, 124, 124, 124, 124, 63: 124, 124, 124, 124, 124, 124, 124, 124, 124, 124, 124, 124, 124, 77: 124},
		{1: 253, 12: 253, 23: 648},
		{1: 250, 3: 428, 427, 429, 430, 434, 445, 12: 250, 23: 645, 48: 426, 431, 432, 433, 435, 436, 437, 438, 439, 440, 441, 442, 443, 444},
		{1: 247, 12: 247},
		{1: 255, 12: 641, 113: 642},
		// 350
		{241, 241, 241, 241, 241, 241, 241, 241, 241, 241, 241, 241, 241, 241, 16: 241, 23: 241, 241, 33: 241, 241, 40: 241, 241, 48: 241, 241, 241, 241, 241, 241, 241, 241, 241, 241, 241, 241, 241, 241, 63: 241, 241, 241, 241, 241, 241, 241, 241, 241, 241, 241, 241, 241, 77: 241},
		{1: 254, 343, 375, 379, 380, 381, 382, 383, 13: 324, 414, 308, 17: 326, 338, 341, 342, 327, 365, 25: 328, 376, 396, 345, 377, 329, 330, 331, 35: 304, 305, 306, 307, 378, 41: 615, 344, 348, 347, 349, 350, 333, 62: 637, 118: 636, 138: 644},
		{1: 643},
		{240, 240, 240, 240, 240, 240, 240, 240, 240, 240, 240, 240, 240, 240, 16: 240, 23: 240, 240, 33: 240, 240, 40: 240, 240, 48: 240, 240, 240, 240, 240, 240, 240, 240, 240, 240, 240, 240, 240, 240, 63: 240, 240, 240, 240, 240, 240, 240, 240, 240, 240, 240, 240, 240, 77: 240},
		{1: 246, 12: 246},
		// 355
		{2: 343, 375, 379, 380, 381, 382, 383, 13: 324, 414, 308, 17: 326, 338, 341, 342, 327, 365, 25: 328, 376, 396, 345, 377, 329, 330, 331, 35: 304, 305, 306, 307, 378, 41: 615, 344, 348, 347, 349, 350, 333, 62: 647, 118: 646},
		{1: 249, 12: 249},
		{1: 248, 3: 428, 427, 429, 430, 434, 445, 12: 248, 48: 426, 431, 432, 433, 435, 436, 437, 438, 439, 440, 441, 442, 443, 444},
		{2: 343, 375, 379, 380, 381, 382, 383, 13: 324, 414, 308, 17: 326, 338, 341, 342, 327, 365, 25: 328, 376, 396, 345, 377, 329, 330, 331, 35: 304, 305, 306, 307, 378, 41: 615, 344, 348, 347, 349, 350, 333, 62: 650, 118: 649},
		{1: 252, 12: 252},
		// 360
		{1: 251, 3: 428, 427, 429, 430, 434, 445, 12: 251, 48: 426, 431, 432, 433, 435, 436, 437, 438, 439, 440, 441, 442, 443, 444},
		{11: 255, 653, 34: 655, 113: 654},
		{261, 261, 261, 261, 261, 261, 261, 261, 261, 261, 261, 261, 261, 261, 16: 261, 23: 261, 261, 33: 261, 261, 40: 261, 261, 48: 261, 261, 261, 261, 261, 261, 261, 261, 261, 261, 261, 261, 261, 261, 63: 261, 261, 261, 261, 261, 261, 261, 261, 261, 261, 261, 261, 261, 77: 261},
		{2: 343, 375, 379, 380, 381, 382, 383, 11: 254, 13: 324, 414, 308, 17: 326, 338, 341, 342, 327, 365, 25: 328, 376, 396, 345, 377, 329, 330, 331, 35: 304, 305, 306, 307, 378, 42: 344, 348, 347, 349, 557, 333, 62: 556, 121: 569},
		{11: 658},
		// 365
		{11: 255, 506, 113: 656},
		{11: 657},
		{259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 16: 259, 23: 259, 259, 33: 259, 259, 40: 259, 259, 48: 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 63: 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 259, 77: 259},
		{260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 16: 260, 23: 260, 260, 33: 260, 260, 40: 260, 260, 48: 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 63: 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 77: 260},
		{1: 665, 343, 375, 379, 380, 381, 382, 383, 13: 324, 414, 308, 574, 326, 338, 341, 342, 327, 365, 25: 328, 376, 396, 345, 377, 329, 330, 331, 35: 304, 305, 306, 307, 378, 41: 575, 344, 348, 347, 349, 350, 333, 62: 661, 106: 659, 125: 662, 142: 663, 170: 664},
		// 370
		{125, 125, 125, 125, 125, 125, 125, 125, 125, 125, 125, 125, 125, 125, 16: 125, 23: 125, 125, 33: 125, 125, 40: 125, 125, 48: 125, 125, 125, 125, 125, 125, 125, 125, 125, 125, 125, 125, 125, 125, 63: 125, 125, 125, 125, 125, 125, 125, 125, 125, 125, 125, 125, 125, 77: 125},
		{1: 152, 3: 428, 427, 429, 430, 434, 445, 12: 152, 23: 670, 48: 426, 431, 432, 433, 435, 436, 437, 438, 439, 440, 441, 442, 443, 444},
		{1: 149, 12: 149},
		{1: 148, 12: 148},
		{1: 255, 12: 666, 113: 667},
		// 375
		{146, 146, 146, 146, 146, 146, 146, 146, 146, 146, 146, 146, 146, 146, 16: 146, 23: 146, 146, 33: 146, 146, 40: 146, 146, 48: 146, 146, 146, 146, 146, 146, 146, 146, 146, 146, 146, 146, 146, 146, 63: 146, 146, 146, 146, 146, 146, 146, 146, 146, 146, 146, 146, 146, 77: 146},
		{1: 254, 343, 375, 379, 380, 381, 382, 383, 13: 324, 414, 308, 574, 326, 338, 341, 342, 327, 365, 25: 328, 376, 396, 345, 377, 329, 330, 331, 35: 304, 305, 306, 307, 378, 41: 575, 344, 348, 347, 349, 350, 333, 62: 661, 106: 659, 125: 662, 142: 669},
		{1: 668},
		{145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 16: 145, 23: 145, 145, 33: 145, 145, 40: 145, 145, 48: 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 63: 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 77: 145},
		{1: 147, 12: 147},
		// 380
		{2: 343, 375, 379, 380, 381, 382, 383, 13: 324, 414, 308, 574, 326, 338, 341, 342, 327, 365, 25: 328, 376, 396, 345, 377, 329, 330, 331, 35: 304, 305, 306, 307, 378, 41: 575, 344, 348, 347, 349, 350, 333, 62: 671, 106: 659, 125: 672},
		{1: 151, 3: 428, 427, 429, 430, 434, 445, 12: 151, 48: 426, 431, 432, 433, 435, 436, 437, 438, 439, 440, 441, 442, 443, 444},
		{1: 150, 12: 150},
		{94, 94, 343, 375, 379, 380, 381, 382, 383, 94, 94, 13: 324, 346, 308, 17: 326, 338, 341, 342, 327, 365, 25: 328, 376, 396, 345, 377, 329, 330, 331, 35: 304, 305, 306, 307, 378, 41: 541, 344, 348, 347, 349, 350, 333, 62: 334, 78: 325, 80: 547, 82: 339, 374, 353, 332, 354, 355, 356, 335, 358, 359, 361, 351, 366, 384, 352, 363, 99: 542, 101: 340, 103: 543, 357, 360, 107: 362, 109: 546, 364, 544, 545, 116: 674},
		{78, 78, 9: 78, 78},
		// 385
		{94, 94, 343, 375, 379, 380, 381, 382, 383, 13: 324, 346, 308, 17: 326, 338, 341, 342, 327, 365, 25: 328, 376, 396, 345, 377, 329, 330, 331, 35: 304, 305, 306, 307, 378, 41: 541, 344, 348, 347, 349, 350, 333, 62: 334, 78: 325, 80: 547, 82: 339, 374, 353, 332, 354, 355, 356, 335, 358, 359, 361, 351, 366, 384, 352, 363, 99: 542, 101: 340, 103: 543, 357, 360, 107: 362, 109: 546, 364, 544, 545, 116: 548, 127: 676},
		{550, 677},
		{140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 16: 140, 23: 140, 140, 33: 140, 140, 40: 140, 140, 48: 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 63: 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 77: 140},
		{3: 428, 427, 429, 430, 434, 445, 11: 681, 48: 426, 431, 432, 433, 435, 436, 437, 438, 439, 440, 441, 442, 443, 444},
		{2: 504, 11: 680},
		// 390
		{142, 142, 142, 142, 142, 142, 142, 142, 142, 142, 142, 142, 142, 142, 16: 142, 23: 142, 142, 33: 142, 142, 40: 142, 142, 48: 142, 142, 142, 142, 142, 142, 142, 142, 142, 142, 142, 142, 142, 142, 63: 142, 142, 142, 142, 142, 142, 142, 142, 142, 142, 142, 142, 142, 77: 142},
		{143, 143, 143, 143, 143, 143, 143, 143, 143, 143, 143, 143, 143, 143, 16: 143, 23: 143, 143, 33: 143, 143, 40: 143, 143, 48: 143, 143, 143, 143, 143, 143, 143, 143, 143, 143, 143, 143, 143, 143, 63: 143, 143, 143, 143, 143, 143, 143, 143, 143, 143, 143, 143, 143, 77: 143},
		{2: 399, 400, 13: 324, 398, 17: 326, 338, 341, 342, 327, 365, 25: 401, 402, 396, 403, 404, 405, 407, 408, 79: 406, 81: 683},
		{40: 684},
		{2: 399, 400, 13: 324, 398, 17: 326, 338, 341, 342, 327, 365, 25: 401, 402, 396, 403, 404, 405, 407, 408, 79: 406, 81: 685},
		// 395
		{144, 144, 144, 9: 144, 144, 144, 144, 15: 144, 144, 23: 144, 144, 33: 144, 144, 40: 144, 144, 76: 144},
		{1: 687, 14: 689, 79: 690, 141: 691, 169: 688},
		{160, 160, 160, 9: 160, 160, 160, 160, 15: 160, 160, 23: 160, 160, 33: 160, 160, 40: 160, 160, 76: 160},
		{694, 105, 115: 693},
		{138, 138, 475, 77: 471, 114: 476, 126: 692},
		// 400
		{157, 157},
		{156, 156},
		{158, 158},
		{1: 696},
		{1: 104, 14: 689, 79: 690, 141: 695},
		// 405
		{155, 155},
		{159, 159, 159, 9: 159, 159, 159, 159, 15: 159, 159, 23: 159, 159, 33: 159, 159, 40: 159, 159, 76: 159},
		{16: 698, 128: 699},
		{94, 94, 343, 375, 379, 380, 381, 382, 383, 13: 324, 346, 308, 17: 326, 338, 341, 342, 327, 365, 25: 328, 376, 396, 345, 377, 329, 330, 331, 35: 304, 305, 306, 307, 378, 41: 541, 344, 348, 347, 349, 350, 333, 62: 334, 78: 325, 80: 547, 82: 339, 374, 353, 332, 354, 355, 356, 335, 358, 359, 361, 351, 366, 384, 352, 363, 99: 542, 101: 340, 103: 543, 357, 360, 107: 362, 109: 546, 364, 544, 545, 116: 548, 127: 708},
		{230, 230, 9: 230, 230, 119: 230, 155: 700},
		// 410
		{228, 228, 9: 228, 228, 119: 701, 154: 702, 156: 703},
		{41: 541, 82: 339, 99: 705, 101: 704},
		{229, 229, 9: 229, 229, 119: 229},
		{172, 172, 9: 172, 172},
		{97, 2: 343, 375, 379, 380, 381, 382, 383, 13: 324, 414, 308, 97, 326, 338, 341, 342, 327, 365, 25: 328, 376, 396, 345, 377, 329, 330, 331, 35: 304, 305, 306, 307, 378, 42: 344, 348, 347, 349, 350, 333, 62: 334, 78: 325, 96: 352, 531, 108: 530, 131: 706},
		// 415
		{227, 227, 9: 227, 227},
		{16: 698, 128: 707},
		{231, 231, 9: 231, 231, 119: 231},
		{550, 709},
		{262, 262, 9: 262, 262, 119: 262},
		// 420
		{14: 712},
		{2: 477, 400, 13: 324, 110, 16: 109, 326, 338, 341, 342, 327, 365, 25: 401, 402, 396, 403, 404, 405, 407, 408, 41: 109, 79: 406, 81: 479, 114: 478, 144: 480},
		{2: 187, 102: 415, 122: 416, 713},
		{2: 475, 114: 476, 126: 714},
		{194, 41: 194},
		// 425
		{263},
		{195},
		{12: 424, 24: 728, 33: 740, 63: 729, 731, 730, 732, 733, 734, 735, 736, 737, 738, 739},
		{16: 199},
		{724, 16: 197},
		// 430
		{16: 698, 128: 723},
		{2: 343, 375, 379, 380, 381, 382, 383, 13: 324, 414, 308, 17: 326, 338, 341, 342, 327, 365, 25: 328, 376, 396, 345, 377, 329, 330, 331, 35: 304, 305, 306, 307, 378, 42: 344, 348, 347, 349, 350, 333, 62: 722},
		{3: 428, 427, 429, 430, 434, 445, 16: 112, 48: 426, 431, 432, 433, 435, 436, 437, 438, 439, 440, 441, 442, 443, 444},
		{196, 196, 9: 196, 196},
		{97, 2: 343, 375, 379, 380, 381, 382, 383, 13: 324, 414, 308, 17: 326, 338, 341, 342, 327, 365, 25: 328, 376, 396, 345, 377, 329, 330, 331, 35: 304, 305, 306, 307, 378, 42: 344, 348, 347, 349, 350, 333, 62: 334, 78: 325, 96: 352, 531, 108: 725},
		// 435
		{726},
		{2: 343, 375, 379, 380, 381, 382, 383, 13: 324, 414, 308, 97, 326, 338, 341, 342, 327, 365, 25: 328, 376, 396, 345, 377, 329, 330, 331, 35: 304, 305, 306, 307, 378, 42: 344, 348, 347, 349, 350, 333, 62: 334, 78: 325, 96: 352, 531, 108: 727},
		{16: 198},
		{2: 343, 375, 379, 380, 381, 382, 383, 13: 324, 414, 308, 17: 326, 338, 341, 342, 327, 365, 25: 328, 376, 396, 345, 377, 329, 330, 331, 35: 304, 305, 306, 307, 378, 42: 344, 348, 347, 349, 350, 333, 62: 412, 78: 755, 134: 756},
		{2: 343, 375, 379, 380, 381, 382, 383, 13: 324, 414, 308, 17: 326, 338, 341, 342, 327, 365, 25: 328, 376, 396, 345, 377, 329, 330, 331, 35: 304, 305, 306, 307, 378, 42: 344, 348, 347, 349, 350, 333, 62: 412, 78: 754},
		// 440
		{2: 343, 375, 379, 380, 381, 382, 383, 13: 324, 414, 308, 17: 326, 338, 341, 342, 327, 365, 25: 328, 376, 396, 345, 377, 329, 330, 331, 35: 304, 305, 306, 307, 378, 42: 344, 348, 347, 349, 350, 333, 62: 412, 78: 753},
		{2: 343, 375, 379, 380, 381, 382, 383, 13: 324, 414, 308, 17: 326, 338, 341, 342, 327, 365, 25: 328, 376, 396, 345, 377, 329, 330, 331, 35: 304, 305, 306, 307, 378, 42: 344, 348, 347, 349, 350, 333, 62: 412, 78: 752},
		{2: 343, 375, 379, 380, 381, 382, 383, 13: 324, 414, 308, 17: 326, 338, 341, 342, 327, 365, 25: 328, 376, 396, 345, 377, 329, 330, 331, 35: 304, 305, 306, 307, 378, 42: 344, 348, 347, 349, 350, 333, 62: 412, 78: 751},
		{2: 343, 375, 379, 380, 381, 382, 383, 13: 324, 414, 308, 17: 326, 338, 341, 342, 327, 365, 25: 328, 376, 396, 345, 377, 329, 330, 331, 35: 304, 305, 306, 307, 378, 42: 344, 348, 347, 349, 350, 333, 62: 412, 78: 750},
		{2: 343, 375, 379, 380, 381, 382, 383, 13: 324, 414, 308, 17: 326, 338, 341, 342, 327, 365, 25: 328, 376, 396, 345, 377, 329, 330, 331, 35: 304, 305, 306, 307, 378, 42: 344, 348, 347, 349, 350, 333, 62: 412, 78: 749},
		// 445
		{2: 343, 375, 379, 380, 381, 382, 383, 13: 324, 414, 308, 17: 326, 338, 341, 342, 327, 365, 25: 328, 376, 396, 345, 377, 329, 330, 331, 35: 304, 305, 306, 307, 378, 42: 344, 348, 347, 349, 350, 333, 62: 412, 78: 748},
		{2: 343, 375, 379, 380, 381, 382, 383, 13: 324, 414, 308, 17: 326, 338, 341, 342, 327, 365, 25: 328, 376, 396, 345, 377, 329, 330, 331, 35: 304, 305, 306, 307, 378, 42: 344, 348, 347, 349, 350, 333, 62: 412, 78: 747},
		{2: 343, 375, 379, 380, 381, 382, 383, 13: 324, 414, 308, 17: 326, 338, 341, 342, 327, 365, 25: 328, 376, 396, 345, 377, 329, 330, 331, 35: 304, 305, 306, 307, 378, 42: 344, 348, 347, 349, 350, 333, 62: 412, 78: 746},
		{2: 343, 375, 379, 380, 381, 382, 383, 13: 324, 414, 308, 17: 326, 338, 341, 342, 327, 365, 25: 328, 376, 396, 345, 377, 329, 330, 331, 35: 304, 305, 306, 307, 378, 42: 344, 348, 347, 349, 350, 333, 62: 412, 78: 745},
		{2: 343, 375, 379, 380, 381, 382, 383, 13: 324, 414, 308, 17: 326, 338, 341, 342, 327, 365, 25: 328, 376, 396, 345, 377, 329, 330, 331, 35: 304, 305, 306, 307, 378, 42: 344, 348, 347, 349, 350, 333, 62: 412, 78: 744},
		// 450
		{2: 343, 375, 379, 380, 381, 382, 383, 13: 324, 414, 308, 17: 326, 338, 341, 342, 327, 365, 25: 328, 376, 396, 345, 377, 329, 330, 331, 35: 304, 305, 306, 307, 378, 42: 344, 348, 347, 349, 350, 333, 62: 412, 78: 741, 134: 742},
		{98, 98, 9: 98, 98, 12: 424, 16: 98},
		{2: 343, 375, 379, 380, 381, 382, 383, 13: 324, 414, 308, 17: 326, 338, 341, 342, 327, 365, 25: 328, 376, 396, 345, 377, 329, 330, 331, 35: 304, 305, 306, 307, 378, 42: 344, 348, 347, 349, 350, 333, 62: 743},
		{3: 428, 427, 429, 430, 434, 445, 16: 113, 48: 426, 431, 432, 433, 435, 436, 437, 438, 439, 440, 441, 442, 443, 444},
		{271, 271, 9: 271, 271, 12: 424, 16: 271},
		// 455
		{272, 272, 9: 272, 272, 12: 424, 16: 272},
		{273, 273, 9: 273, 273, 12: 424, 16: 273},
		{274, 274, 9: 274, 274, 12: 424, 16: 274},
		{275, 275, 9: 275, 275, 12: 424, 16: 275},
		{276, 276, 9: 276, 276, 12: 424, 16: 276},
		// 460
		{277, 277, 9: 277, 277, 12: 424, 16: 277},
		{278, 278, 9: 278, 278, 12: 424, 16: 278},
		{279, 279, 9: 279, 279, 12: 424, 16: 279},
		{280, 280, 9: 280, 280, 12: 424, 16: 280},
		{281, 281, 9: 281, 281, 12: 424, 16: 281},
		// 465
		{282, 282, 9: 282, 282, 12: 424, 16: 282},
		{2: 343, 375, 379, 380, 381, 382, 383, 13: 324, 414, 308, 17: 326, 338, 341, 342, 327, 365, 25: 328, 376, 396, 345, 377, 329, 330, 331, 35: 304, 305, 306, 307, 378, 42: 344, 348, 347, 349, 350, 333, 62: 757},
		{3: 428, 427, 429, 430, 434, 445, 16: 114, 48: 426, 431, 432, 433, 435, 436, 437, 438, 439, 440, 441, 442, 443, 444},
		{100, 100, 9: 100, 100, 16: 100},
		{99, 99, 9: 99, 99, 16: 99},
		// 470
		{11: 768, 14: 385, 100: 762, 129: 770, 153: 769},
		{237, 237, 9: 237, 237},
		{236, 236, 399, 400, 9: 236, 236, 236, 397, 324, 398, 17: 326, 338, 341, 342, 327, 365, 24: 763, 401, 402, 396, 403, 404, 405, 407, 408, 79: 406, 81: 764},
		{2: 343, 375, 379, 380, 381, 382, 383, 13: 324, 414, 308, 17: 326, 338, 341, 342, 327, 365, 25: 328, 376, 396, 345, 377, 329, 330, 331, 35: 304, 305, 306, 307, 378, 42: 344, 348, 347, 349, 350, 333, 62: 412, 78: 767},
		{24: 765},
		// 475
		{2: 343, 375, 379, 380, 381, 382, 383, 13: 324, 414, 308, 17: 326, 338, 341, 342, 327, 365, 25: 328, 376, 396, 345, 377, 329, 330, 331, 35: 304, 305, 306, 307, 378, 42: 344, 348, 347, 349, 350, 333, 62: 412, 78: 766},
		{234, 234, 9: 234, 234, 234, 424},
		{235, 235, 9: 235, 235, 235, 424},
		{239, 239, 9: 239, 239},
		{772, 11: 105, 115: 771},
		// 480
		{233, 11: 233},
		{11: 774},
		{11: 104, 14: 385, 100: 762, 129: 773},
		{232, 11: 232},
		{238, 238, 9: 238, 238},
		// 485
		{2: 399, 400, 13: 324, 398, 17: 326, 338, 341, 342, 327, 365, 25: 401, 402, 396, 403, 404, 405, 407, 408, 79: 406, 81: 776},
		{256, 256, 256, 9: 256, 256, 256, 256, 15: 256, 256, 23: 256, 256, 33: 256, 256, 40: 256, 256, 76: 256},
		{258, 258, 258, 9: 258, 258, 258, 258, 15: 258, 258, 23: 258, 258, 33: 258, 258, 40: 258, 258, 76: 258},
		{2: 399, 400, 13: 324, 398, 17: 326, 338, 341, 342, 327, 365, 25: 401, 402, 396, 403, 404, 405, 407, 408, 79: 406, 81: 779},
		{257, 257, 257, 9: 257, 257, 257, 257, 15: 257, 257, 23: 257, 257, 33: 257, 257, 40: 257, 257, 76: 257},
		// 490
		{2: 343, 375, 379, 380, 381, 382, 383, 13: 324, 414, 308, 17: 326, 338, 341, 342, 327, 365, 25: 328, 376, 396, 345, 377, 329, 330, 331, 35: 304, 305, 306, 307, 378, 42: 344, 348, 347, 349, 350, 333, 62: 412, 78: 755},
		{2: 343, 375, 379, 380, 381, 382, 383, 13: 324, 414, 308, 17: 326, 338, 341, 342, 327, 365, 25: 328, 376, 396, 345, 377, 329, 330, 331, 35: 304, 305, 306, 307, 378, 42: 344, 348, 347, 349, 350, 333, 62: 412, 78: 741},
		{40: 788},
		{3: 428, 427, 429, 430, 434, 445, 40: 786, 48: 426, 431, 432, 433, 435, 436, 437, 438, 439, 440, 441, 442, 443, 444},
		{2: 399, 400, 13: 324, 398, 17: 326, 338, 341, 342, 327, 365, 25: 401, 402, 396, 403, 404, 405, 407, 408, 79: 406, 81: 785},
		// 495
		{95, 95, 95, 9: 95, 95, 95, 95, 15: 95, 95, 23: 95, 95, 33: 95, 95, 40: 95, 95, 76: 95},
		{2: 399, 400, 13: 324, 398, 17: 326, 338, 341, 342, 327, 365, 25: 401, 402, 396, 403, 404, 405, 407, 408, 79: 406, 81: 787},
		{283, 283, 283, 9: 283, 283, 283, 283, 15: 283, 283, 23: 283, 283, 33: 283, 283, 40: 283, 283, 76: 283},
		{2: 399, 400, 13: 324, 398, 17: 326, 338, 341, 342, 327, 365, 25: 401, 402, 396, 403, 404, 405, 407, 408, 79: 406, 81: 789},
		{284, 284, 284, 9: 284, 284, 284, 284, 15: 284, 284, 23: 284, 284, 33: 284, 284, 40: 284, 284, 76: 284},
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
	const yyError = 80

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
			yyVAL.node = &BasicLiteral{
				Token: yyS[yypt-0].Token,
			}
		}
	case 21:
		{
			yyVAL.node = &BasicLiteral{
				Case:  1,
				Token: yyS[yypt-0].Token,
			}
		}
	case 22:
		{
			yyVAL.node = &BasicLiteral{
				Case:  2,
				Token: yyS[yypt-0].Token,
			}
		}
	case 23:
		{
			yyVAL.node = &BasicLiteral{
				Case:  3,
				Token: yyS[yypt-0].Token,
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
			t := lhs.Token
			s := string(t.S())
			value, err := strconv.Unquote(s)
			if err != nil {
				lx.err(t, "%s: %q", err, t.S())
				break
			}

			// https://github.com/golang/go/issues/15997
			if b := lhs.Token.S(); len(b) != 0 && b[0] == '`' {
				value = strings.Replace(value, "\r", "", -1)
			}
			lhs.val = StringID(dict.SID(value))
		}
	case 25:
		{
			yyVAL.node = &Block{
				Token:         yyS[yypt-2].Token,
				StatementList: yyS[yypt-1].node.(*StatementList).reverse(),
				Token2:        yyS[yypt-0].Token,
			}
		}
	case 26:
		{
			yyVAL.node = (*BlockOpt)(nil)
		}
	case 27:
		{
			yyVAL.node = &BlockOpt{
				Block: yyS[yypt-0].node.(*Block),
			}
		}
	case 28:
		{
			yyVAL.node = &Body{
				Token:         yyS[yypt-2].Token,
				StatementList: yyS[yypt-1].node.(*StatementList).reverse(),
				Token2:        yyS[yypt-0].Token,
			}
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
			yyVAL.node = &ConstSpec{
				IdentifierList: yyS[yypt-0].node.(*IdentifierList).reverse(),
			}
		}
	case 55:
		{
			yyVAL.node = &ConstSpec{
				Case:           1,
				IdentifierList: yyS[yypt-2].node.(*IdentifierList).reverse(),
				Token:          yyS[yypt-1].Token,
				ExpressionList: yyS[yypt-0].node.(*ExpressionList).reverse(),
			}
		}
	case 56:
		{
			yyVAL.node = &ConstSpec{
				Case:           2,
				IdentifierList: yyS[yypt-3].node.(*IdentifierList).reverse(),
				Typ:            yyS[yypt-2].node.(*Typ),
				Token:          yyS[yypt-1].Token,
				ExpressionList: yyS[yypt-0].node.(*ExpressionList).reverse(),
			}
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
			yyVAL.node = &Elif{
				Token:    yyS[yypt-3].Token,
				If:       yyS[yypt-2].node.(*If),
				IfHeader: yyS[yypt-1].node.(*IfHeader),
				Body:     yyS[yypt-0].node.(*Body),
			}
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
			yyVAL.node = &ForStatement{
				Token:     yyS[yypt-2].Token,
				ForHeader: yyS[yypt-1].node.(*ForHeader),
				Body:      yyS[yypt-0].node.(*Body),
			}
		}
	case 95:
		{
			yyVAL.node = &FuncDecl{
				FuncOrMethod: yyS[yypt-1].node.(*FuncOrMethod),
				BlockOpt:     yyS[yypt-0].node.(*BlockOpt),
			}
		}
	case 96:
		{
			yyVAL.node = &FuncOrMethod{
				Function:            yyS[yypt-4].node.(*Function),
				ReceiverOpt:         yyS[yypt-3].node.(*ReceiverOpt),
				Token:               yyS[yypt-2].Token,
				GenericArgumentsOpt: yyS[yypt-1].node.(*GenericArgumentsOpt),
				Signature:           yyS[yypt-0].node.(*Signature),
			}
		}
	case 97:
		{
			yyVAL.node = &FuncType{
				Function:  yyS[yypt-1].node.(*Function),
				Signature: yyS[yypt-0].node.(*Signature),
			}
		}
	case 98:
		{
			yyVAL.node = &Function{
				Token: yyS[yypt-0].Token,
			}
		}
	case 99:
		{
			yyVAL.node = &GenericArgumentList{
				GenericArgumentListItem: yyS[yypt-0].node.(*GenericArgumentListItem),
			}
		}
	case 100:
		{
			yyVAL.node = &GenericArgumentList{
				Case:                1,
				GenericArgumentList: yyS[yypt-2].node.(*GenericArgumentList),
				Token:               yyS[yypt-1].Token,
				GenericArgumentListItem: yyS[yypt-0].node.(*GenericArgumentListItem),
			}
		}
	case 101:
		{
			yyVAL.node = &GenericArgumentListItem{
				Typ: yyS[yypt-0].node.(*Typ),
			}
		}
	case 102:
		{
			yyVAL.node = &GenericArguments{
				Token:               yyS[yypt-2].Token,
				GenericArgumentList: yyS[yypt-1].node.(*GenericArgumentList).reverse(),
				Token2:              yyS[yypt-0].Token,
			}
		}
	case 103:
		{
			yyVAL.node = (*GenericArgumentsOpt)(nil)
		}
	case 104:
		{
			yyVAL.node = &GenericArgumentsOpt{
				GenericArguments: yyS[yypt-0].node.(*GenericArguments),
			}
		}
	case 105:
		{
			yyVAL.node = &GenericParameterList{
				GenericParameterListItem: yyS[yypt-0].node.(*GenericParameterListItem),
			}
		}
	case 106:
		{
			yyVAL.node = &GenericParameterList{
				Case:                 1,
				GenericParameterList: yyS[yypt-2].node.(*GenericParameterList),
				Token:                yyS[yypt-1].Token,
				GenericParameterListItem: yyS[yypt-0].node.(*GenericParameterListItem),
			}
		}
	case 107:
		{
			yyVAL.node = &GenericParameterListItem{
				Token: yyS[yypt-0].Token,
			}
		}
	case 108:
		{
			yyVAL.node = &GenericParams{
				Token:                yyS[yypt-2].Token,
				GenericParameterList: yyS[yypt-1].node.(*GenericParameterList).reverse(),
				Token2:               yyS[yypt-0].Token,
			}
		}
	case 109:
		{
			yyVAL.node = (*GenericParamsOpt)(nil)
		}
	case 110:
		{
			yyVAL.node = &GenericParamsOpt{
				GenericParams: yyS[yypt-0].node.(*GenericParams),
			}
		}
	case 111:
		{
			yyVAL.node = (*IdentifierOpt)(nil)
		}
	case 112:
		{
			yyVAL.node = &IdentifierOpt{
				Token: yyS[yypt-0].Token,
			}
		}
	case 113:
		{
			yyVAL.node = &IdentifierList{
				Token: yyS[yypt-0].Token,
			}
		}
	case 114:
		{
			yyVAL.node = &IdentifierList{
				Case:           1,
				IdentifierList: yyS[yypt-2].node.(*IdentifierList),
				Token:          yyS[yypt-1].Token,
				Token2:         yyS[yypt-0].Token,
			}
		}
	case 115:
		{
			yyVAL.node = &If{
				Token: yyS[yypt-0].Token,
			}
		}
	case 116:
		{
			yyVAL.node = &IfHeader{
				SimpleStatementOpt: yyS[yypt-0].node.(*SimpleStatementOpt),
			}
		}
	case 117:
		{
			yyVAL.node = &IfHeader{
				Case:                1,
				SimpleStatementOpt:  yyS[yypt-2].node.(*SimpleStatementOpt),
				Token:               yyS[yypt-1].Token,
				SimpleStatementOpt2: yyS[yypt-0].node.(*SimpleStatementOpt),
			}
		}
	case 118:
		{
			yyVAL.node = &IfStatement{
				If:       yyS[yypt-4].node.(*If),
				IfHeader: yyS[yypt-3].node.(*IfHeader),
				Body:     yyS[yypt-2].node.(*Body),
				ElifList: yyS[yypt-1].node.(*ElifList).reverse(),
				ElseOpt:  yyS[yypt-0].node.(*ElseOpt),
			}
		}
	case 119:
		{
			yyVAL.node = &ImportDecl{
				Token:  yyS[yypt-2].Token,
				Token2: yyS[yypt-1].Token,
				Token3: yyS[yypt-0].Token,
			}
		}
	case 120:
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
	case 121:
		{
			yyVAL.node = &ImportDecl{
				Case:       2,
				Token:      yyS[yypt-1].Token,
				ImportSpec: yyS[yypt-0].node.(*ImportSpec),
			}
		}
	case 122:
		{
			lx := yylex.(*lexer)
			lhs := &ImportSpec{
				Token:        yyS[yypt-1].Token,
				BasicLiteral: yyS[yypt-0].node.(*BasicLiteral),
			}
			yyVAL.node = lhs
			lhs.post(lx)
		}
	case 123:
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
	case 124:
		{
			yyVAL.node = &ImportSpec{
				Case:         2,
				Token:        yyS[yypt-2].Token,
				BasicLiteral: yyS[yypt-1].node.(*BasicLiteral),
			}
		}
	case 125:
		{
			yyVAL.node = &ImportSpec{
				Case:          3,
				IdentifierOpt: yyS[yypt-2].node.(*IdentifierOpt),
				BasicLiteral:  yyS[yypt-1].node.(*BasicLiteral),
			}
		}
	case 126:
		{
			yyVAL.node = &ImportSpecList{
				ImportSpec: yyS[yypt-0].node.(*ImportSpec),
			}
		}
	case 127:
		{
			yyVAL.node = &ImportSpecList{
				Case:           1,
				ImportSpecList: yyS[yypt-2].node.(*ImportSpecList),
				Token:          yyS[yypt-1].Token,
				ImportSpec:     yyS[yypt-0].node.(*ImportSpec),
			}
		}
	case 128:
		{
			yyVAL.node = (*ImportList)(nil)
		}
	case 129:
		{
			yyVAL.node = &ImportList{
				ImportList: yyS[yypt-2].node.(*ImportList),
				ImportDecl: yyS[yypt-1].node.(*ImportDecl),
				Token:      yyS[yypt-0].Token,
			}
		}
	case 130:
		{
			yyVAL.node = &InterfaceType{
				Token:  yyS[yypt-2].Token,
				LBrace: yyS[yypt-1].node.(*LBrace),
				Token2: yyS[yypt-0].Token,
			}
		}
	case 131:
		{
			yyVAL.node = &InterfaceType{
				Case:                    1,
				Token:                   yyS[yypt-4].Token,
				LBrace:                  yyS[yypt-3].node.(*LBrace),
				InterfaceMethodDeclList: yyS[yypt-2].node.(*InterfaceMethodDeclList).reverse(),
				SemicolonOpt:            yyS[yypt-1].node.(*SemicolonOpt),
				Token2:                  yyS[yypt-0].Token,
			}
		}
	case 132:
		{
			yyVAL.node = &InterfaceMethodDecl{
				Token:     yyS[yypt-1].Token,
				Signature: yyS[yypt-0].node.(*Signature),
			}
		}
	case 133:
		{
			yyVAL.node = &InterfaceMethodDecl{
				Case:           1,
				QualifiedIdent: yyS[yypt-0].node.(*QualifiedIdent),
			}
		}
	case 134:
		{
			yyVAL.node = &InterfaceMethodDeclList{
				InterfaceMethodDecl: yyS[yypt-0].node.(*InterfaceMethodDecl),
			}
		}
	case 135:
		{
			yyVAL.node = &InterfaceMethodDeclList{
				Case: 1,
				InterfaceMethodDeclList: yyS[yypt-2].node.(*InterfaceMethodDeclList),
				Token:               yyS[yypt-1].Token,
				InterfaceMethodDecl: yyS[yypt-0].node.(*InterfaceMethodDecl),
			}
		}
	case 136:
		{
			lx := yylex.(*lexer)
			yyVAL.node = &LBrace{
				Token: yyS[yypt-0].Token,
			}
			lx.fixLBR()
		}
	case 137:
		{
			yyVAL.node = &LBrace{
				Case:  1,
				Token: yyS[yypt-0].Token,
			}
		}
	case 138:
		{
			yyVAL.node = &LBraceCompLitItem{
				Expression: yyS[yypt-0].node.(*Expression),
			}
		}
	case 139:
		{
			yyVAL.node = &LBraceCompLitItem{
				Case:        1,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 140:
		{
			yyVAL.node = &LBraceCompLitItem{
				Case:               2,
				Expression:         yyS[yypt-2].node.(*Expression),
				Token:              yyS[yypt-1].Token,
				LBraceCompLitValue: yyS[yypt-0].node.(*LBraceCompLitValue),
			}
		}
	case 141:
		{
			yyVAL.node = &LBraceCompLitItem{
				Case:               3,
				LBraceCompLitValue: yyS[yypt-0].node.(*LBraceCompLitValue),
			}
		}
	case 142:
		{
			yyVAL.node = &LBraceCompLitItemList{
				LBraceCompLitItem: yyS[yypt-0].node.(*LBraceCompLitItem),
			}
		}
	case 143:
		{
			yyVAL.node = &LBraceCompLitItemList{
				Case: 1,
				LBraceCompLitItemList: yyS[yypt-2].node.(*LBraceCompLitItemList),
				Token:             yyS[yypt-1].Token,
				LBraceCompLitItem: yyS[yypt-0].node.(*LBraceCompLitItem),
			}
		}
	case 144:
		{
			yyVAL.node = &LBraceCompLitValue{
				LBrace: yyS[yypt-1].node.(*LBrace),
				Token:  yyS[yypt-0].Token,
			}
		}
	case 145:
		{
			yyVAL.node = &LBraceCompLitValue{
				Case:                  1,
				LBrace:                yyS[yypt-3].node.(*LBrace),
				LBraceCompLitItemList: yyS[yypt-2].node.(*LBraceCompLitItemList).reverse(),
				CommaOpt:              yyS[yypt-1].node.(*CommaOpt),
				Token:                 yyS[yypt-0].Token,
			}
		}
	case 146:
		{
			yyVAL.node = &MapType{
				Token:  yyS[yypt-4].Token,
				Token2: yyS[yypt-3].Token,
				Typ:    yyS[yypt-2].node.(*Typ),
				Token3: yyS[yypt-1].Token,
				Typ2:   yyS[yypt-0].node.(*Typ),
			}
		}
	case 147:
		{
			yyVAL.node = &Operand{
				Token:      yyS[yypt-2].Token,
				Expression: yyS[yypt-1].node.(*Expression),
				Token2:     yyS[yypt-0].Token,
			}
		}
	case 148:
		{
			yyVAL.node = &Operand{
				Case:        1,
				Token:       yyS[yypt-2].Token,
				TypeLiteral: yyS[yypt-1].node.(*TypeLiteral),
				Token2:      yyS[yypt-0].Token,
			}
		}
	case 149:
		{
			yyVAL.node = &Operand{
				Case:         2,
				BasicLiteral: yyS[yypt-0].node.(*BasicLiteral),
			}
		}
	case 150:
		{
			yyVAL.node = &Operand{
				Case:          3,
				FuncType:      yyS[yypt-3].node.(*FuncType),
				LBrace:        yyS[yypt-2].node.(*LBrace),
				StatementList: yyS[yypt-1].node.(*StatementList).reverse(),
				Token:         yyS[yypt-0].Token,
			}
		}
	case 151:
		{
			yyVAL.node = &Operand{
				Case:                4,
				Token:               yyS[yypt-1].Token,
				GenericArgumentsOpt: yyS[yypt-0].node.(*GenericArgumentsOpt),
			}
		}
	case 152:
		{
			yyVAL.node = &QualifiedIdent{
				Token: yyS[yypt-0].Token,
			}
		}
	case 153:
		{
			yyVAL.node = &QualifiedIdent{
				Case:   1,
				Token:  yyS[yypt-2].Token,
				Token2: yyS[yypt-1].Token,
				Token3: yyS[yypt-0].Token,
			}
		}
	case 154:
		{
			lx := yylex.(*lexer)
			if !lx.build { // Build tags not satisfied
				return 0
			}
		}
	case 155:
		{
			lx := yylex.(*lexer)
			lhs := &PackageClause{
				Token:  yyS[yypt-3].Token,
				Token2: yyS[yypt-1].Token,
				Token3: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lx.pkgName = lhs.Token2.Val
			if lx.parseOnlyPackageClause {
				return 0
			}

			lhs.post(lx)
		}
	case 156:
		{
			yyVAL.node = &ParameterDecl{
				Token: yyS[yypt-1].Token,
				Typ:   yyS[yypt-0].node.(*Typ),
			}
		}
	case 157:
		{
			yyVAL.node = &ParameterDecl{
				Case:   1,
				Token:  yyS[yypt-2].Token,
				Token2: yyS[yypt-1].Token,
				Typ:    yyS[yypt-0].node.(*Typ),
			}
		}
	case 158:
		{
			yyVAL.node = &ParameterDecl{
				Case:  2,
				Token: yyS[yypt-1].Token,
				Typ:   yyS[yypt-0].node.(*Typ),
			}
		}
	case 159:
		{
			yyVAL.node = &ParameterDecl{
				Case: 3,
				Typ:  yyS[yypt-0].node.(*Typ),
			}
		}
	case 160:
		{
			yyVAL.node = &ParameterDeclList{
				ParameterDecl: yyS[yypt-0].node.(*ParameterDecl),
			}
		}
	case 161:
		{
			yyVAL.node = &ParameterDeclList{
				Case:              1,
				ParameterDeclList: yyS[yypt-2].node.(*ParameterDeclList),
				Token:             yyS[yypt-1].Token,
				ParameterDecl:     yyS[yypt-0].node.(*ParameterDecl),
			}
		}
	case 162:
		{
			yyVAL.node = &Parameters{
				Token:  yyS[yypt-1].Token,
				Token2: yyS[yypt-0].Token,
			}
		}
	case 163:
		{
			yyVAL.node = &Parameters{
				Case:              1,
				Token:             yyS[yypt-3].Token,
				ParameterDeclList: yyS[yypt-2].node.(*ParameterDeclList).reverse(),
				CommaOpt:          yyS[yypt-1].node.(*CommaOpt),
				Token2:            yyS[yypt-0].Token,
			}
		}
	case 164:
		{
			yyVAL.node = &PrimaryExpression{
				Operand: yyS[yypt-0].node.(*Operand),
			}
		}
	case 165:
		{
			yyVAL.node = &PrimaryExpression{
				Case:               1,
				CompLitType:        yyS[yypt-1].node.(*CompLitType),
				LBraceCompLitValue: yyS[yypt-0].node.(*LBraceCompLitValue),
			}
		}
	case 166:
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
	case 167:
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
	case 168:
		{
			yyVAL.node = &PrimaryExpression{
				Case:              4,
				PrimaryExpression: yyS[yypt-2].node.(*PrimaryExpression),
				Token:             yyS[yypt-1].Token,
				Token2:            yyS[yypt-0].Token,
			}
		}
	case 169:
		{
			yyVAL.node = &PrimaryExpression{
				Case:              5,
				PrimaryExpression: yyS[yypt-3].node.(*PrimaryExpression),
				Token:             yyS[yypt-2].Token,
				Expression:        yyS[yypt-1].node.(*Expression),
				Token2:            yyS[yypt-0].Token,
			}
		}
	case 170:
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
	case 171:
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
	case 172:
		{
			yyVAL.node = &PrimaryExpression{
				Case:              8,
				PrimaryExpression: yyS[yypt-1].node.(*PrimaryExpression),
				Call:              yyS[yypt-0].node.(*Call),
			}
		}
	case 173:
		{
			yyVAL.node = &PrimaryExpression{
				Case:              9,
				PrimaryExpression: yyS[yypt-1].node.(*PrimaryExpression),
				CompLitValue:      yyS[yypt-0].node.(*CompLitValue),
			}
		}
	case 174:
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
	case 175:
		{
			lx := yylex.(*lexer)
			lhs := &Prologue{
				PackageClause: yyS[yypt-1].node.(*PackageClause),
				ImportList:    yyS[yypt-0].node.(*ImportList).reverse(),
			}
			yyVAL.node = lhs
			lhs.post(lx)
		}
	case 176:
		{
			yyVAL.node = &Range{
				ExpressionList: yyS[yypt-3].node.(*ExpressionList).reverse(),
				Token:          yyS[yypt-2].Token,
				Token2:         yyS[yypt-1].Token,
				Expression:     yyS[yypt-0].node.(*Expression),
			}
		}
	case 177:
		{
			yyVAL.node = &Range{
				Case:           1,
				ExpressionList: yyS[yypt-3].node.(*ExpressionList).reverse(),
				Token:          yyS[yypt-2].Token,
				Token2:         yyS[yypt-1].Token,
				Expression:     yyS[yypt-0].node.(*Expression),
			}
		}
	case 178:
		{
			yyVAL.node = &Range{
				Case:       2,
				Token:      yyS[yypt-1].Token,
				Expression: yyS[yypt-0].node.(*Expression),
			}
		}
	case 179:
		{
			yyVAL.node = (*ReceiverOpt)(nil)
		}
	case 180:
		{
			yyVAL.node = &ReceiverOpt{
				Parameters: yyS[yypt-0].node.(*Parameters),
			}
		}
	case 181:
		{
			yyVAL.node = (*ResultOpt)(nil)
		}
	case 182:
		{
			yyVAL.node = &ResultOpt{
				Case:       1,
				Parameters: yyS[yypt-0].node.(*Parameters),
			}
		}
	case 183:
		{
			yyVAL.node = &ResultOpt{
				Case: 2,
				Typ:  yyS[yypt-0].node.(*Typ),
			}
		}
	case 184:
		{
			yyVAL.node = &SelectStatement{
				Token:      yyS[yypt-1].Token,
				SwitchBody: yyS[yypt-0].node.(*SwitchBody),
			}
		}
	case 185:
		{
			yyVAL.node = (*SemicolonOpt)(nil)
		}
	case 186:
		{
			yyVAL.node = &SemicolonOpt{
				Token: yyS[yypt-0].Token,
			}
		}
	case 187:
		{
			yyVAL.node = &Signature{
				Parameters: yyS[yypt-1].node.(*Parameters),
				ResultOpt:  yyS[yypt-0].node.(*ResultOpt),
			}
		}
	case 188:
		{
			yyVAL.node = &SimpleStatement{
				Assignment: yyS[yypt-0].node.(*Assignment),
			}
		}
	case 189:
		{
			yyVAL.node = &SimpleStatement{
				Case:       1,
				Expression: yyS[yypt-0].node.(*Expression),
			}
		}
	case 190:
		{
			yyVAL.node = &SimpleStatement{
				Case:       2,
				Expression: yyS[yypt-1].node.(*Expression),
				Token:      yyS[yypt-0].Token,
			}
		}
	case 191:
		{
			yyVAL.node = &SimpleStatement{
				Case:       3,
				Expression: yyS[yypt-1].node.(*Expression),
				Token:      yyS[yypt-0].Token,
			}
		}
	case 192:
		{
			yyVAL.node = &SimpleStatement{
				Case:            4,
				ExpressionList:  yyS[yypt-2].node.(*ExpressionList).reverse(),
				Token:           yyS[yypt-1].Token,
				ExpressionList2: yyS[yypt-0].node.(*ExpressionList).reverse(),
			}
		}
	case 193:
		{
			yyVAL.node = (*SimpleStatementOpt)(nil)
		}
	case 194:
		{
			yyVAL.node = &SimpleStatementOpt{
				SimpleStatement: yyS[yypt-0].node.(*SimpleStatement),
			}
		}
	case 195:
		{
			yyVAL.node = &SliceType{
				Token:  yyS[yypt-2].Token,
				Token2: yyS[yypt-1].Token,
				Typ:    yyS[yypt-0].node.(*Typ),
			}
		}
	case 196:
		{
			yyVAL.node = (*Statement)(nil)
		}
	case 197:
		{
			yyVAL.node = &Statement{
				Case:  1,
				Block: yyS[yypt-0].node.(*Block),
			}
		}
	case 198:
		{
			yyVAL.node = &Statement{
				Case:      2,
				ConstDecl: yyS[yypt-0].node.(*ConstDecl),
			}
		}
	case 199:
		{
			yyVAL.node = &Statement{
				Case:     3,
				TypeDecl: yyS[yypt-0].node.(*TypeDecl),
			}
		}
	case 200:
		{
			yyVAL.node = &Statement{
				Case:    4,
				VarDecl: yyS[yypt-0].node.(*VarDecl),
			}
		}
	case 201:
		{
			yyVAL.node = &Statement{
				Case:             5,
				StatementNonDecl: yyS[yypt-0].node.(*StatementNonDecl),
			}
		}
	case 202:
		{
			yyVAL.node = &Statement{
				Case: 6,
			}
		}
	case 203:
		{
			yyVAL.node = &StatementList{
				Statement: yyS[yypt-0].node.(*Statement),
			}
		}
	case 204:
		{
			yyVAL.node = &StatementList{
				Case:          1,
				StatementList: yyS[yypt-2].node.(*StatementList),
				Token:         yyS[yypt-1].Token,
				Statement:     yyS[yypt-0].node.(*Statement),
			}
		}
	case 205:
		{
			yyVAL.node = &StatementNonDecl{
				Token:         yyS[yypt-1].Token,
				IdentifierOpt: yyS[yypt-0].node.(*IdentifierOpt),
			}
		}
	case 206:
		{
			yyVAL.node = &StatementNonDecl{
				Case:          1,
				Token:         yyS[yypt-1].Token,
				IdentifierOpt: yyS[yypt-0].node.(*IdentifierOpt),
			}
		}
	case 207:
		{
			yyVAL.node = &StatementNonDecl{
				Case:       2,
				Token:      yyS[yypt-1].Token,
				Expression: yyS[yypt-0].node.(*Expression),
			}
		}
	case 208:
		{
			yyVAL.node = &StatementNonDecl{
				Case:  3,
				Token: yyS[yypt-0].Token,
			}
		}
	case 209:
		{
			yyVAL.node = &StatementNonDecl{
				Case:         4,
				ForStatement: yyS[yypt-0].node.(*ForStatement),
			}
		}
	case 210:
		{
			yyVAL.node = &StatementNonDecl{
				Case:       5,
				Token:      yyS[yypt-1].Token,
				Expression: yyS[yypt-0].node.(*Expression),
			}
		}
	case 211:
		{
			yyVAL.node = &StatementNonDecl{
				Case:   6,
				Token:  yyS[yypt-1].Token,
				Token2: yyS[yypt-0].Token,
			}
		}
	case 212:
		{
			yyVAL.node = &StatementNonDecl{
				Case:      7,
				Token:     yyS[yypt-2].Token,
				Token2:    yyS[yypt-1].Token,
				Statement: yyS[yypt-0].node.(*Statement),
			}
		}
	case 213:
		{
			yyVAL.node = &StatementNonDecl{
				Case:        8,
				IfStatement: yyS[yypt-0].node.(*IfStatement),
			}
		}
	case 214:
		{
			yyVAL.node = &StatementNonDecl{
				Case:              9,
				Token:             yyS[yypt-1].Token,
				ExpressionListOpt: yyS[yypt-0].node.(*ExpressionListOpt),
			}
		}
	case 215:
		{
			yyVAL.node = &StatementNonDecl{
				Case:            10,
				SelectStatement: yyS[yypt-0].node.(*SelectStatement),
			}
		}
	case 216:
		{
			yyVAL.node = &StatementNonDecl{
				Case:            11,
				SimpleStatement: yyS[yypt-0].node.(*SimpleStatement),
			}
		}
	case 217:
		{
			yyVAL.node = &StatementNonDecl{
				Case:            12,
				SwitchStatement: yyS[yypt-0].node.(*SwitchStatement),
			}
		}
	case 218:
		{
			yyVAL.node = (*StringLitOpt)(nil)
		}
	case 219:
		{
			yyVAL.node = &StringLitOpt{
				Token: yyS[yypt-0].Token,
			}
		}
	case 220:
		{
			yyVAL.node = &StructFieldDecl{
				Token:          yyS[yypt-2].Token,
				QualifiedIdent: yyS[yypt-1].node.(*QualifiedIdent),
				StringLitOpt:   yyS[yypt-0].node.(*StringLitOpt),
			}
		}
	case 221:
		{
			yyVAL.node = &StructFieldDecl{
				Case:           1,
				IdentifierList: yyS[yypt-2].node.(*IdentifierList).reverse(),
				Typ:            yyS[yypt-1].node.(*Typ),
				StringLitOpt:   yyS[yypt-0].node.(*StringLitOpt),
			}
		}
	case 222:
		{
			yyVAL.node = &StructFieldDecl{
				Case:           2,
				QualifiedIdent: yyS[yypt-1].node.(*QualifiedIdent),
				StringLitOpt:   yyS[yypt-0].node.(*StringLitOpt),
			}
		}
	case 223:
		{
			yyVAL.node = &StructFieldDecl{
				Case:           3,
				Token:          yyS[yypt-3].Token,
				QualifiedIdent: yyS[yypt-2].node.(*QualifiedIdent),
				Token2:         yyS[yypt-1].Token,
				StringLitOpt:   yyS[yypt-0].node.(*StringLitOpt),
			}
		}
	case 224:
		{
			yyVAL.node = &StructFieldDecl{
				Case:           4,
				Token:          yyS[yypt-4].Token,
				Token2:         yyS[yypt-3].Token,
				QualifiedIdent: yyS[yypt-2].node.(*QualifiedIdent),
				Token3:         yyS[yypt-1].Token,
				StringLitOpt:   yyS[yypt-0].node.(*StringLitOpt),
			}
		}
	case 225:
		{
			yyVAL.node = &StructFieldDecl{
				Case:           5,
				Token:          yyS[yypt-4].Token,
				Token2:         yyS[yypt-3].Token,
				QualifiedIdent: yyS[yypt-2].node.(*QualifiedIdent),
				Token3:         yyS[yypt-1].Token,
				StringLitOpt:   yyS[yypt-0].node.(*StringLitOpt),
			}
		}
	case 226:
		{
			yyVAL.node = &StructFieldDeclList{
				StructFieldDecl: yyS[yypt-0].node.(*StructFieldDecl),
			}
		}
	case 227:
		{
			yyVAL.node = &StructFieldDeclList{
				Case:                1,
				StructFieldDeclList: yyS[yypt-2].node.(*StructFieldDeclList),
				Token:               yyS[yypt-1].Token,
				StructFieldDecl:     yyS[yypt-0].node.(*StructFieldDecl),
			}
		}
	case 228:
		{
			yyVAL.node = &StructType{
				Token:  yyS[yypt-2].Token,
				LBrace: yyS[yypt-1].node.(*LBrace),
				Token2: yyS[yypt-0].Token,
			}
		}
	case 229:
		{
			yyVAL.node = &StructType{
				Case:                1,
				Token:               yyS[yypt-4].Token,
				LBrace:              yyS[yypt-3].node.(*LBrace),
				StructFieldDeclList: yyS[yypt-2].node.(*StructFieldDeclList).reverse(),
				SemicolonOpt:        yyS[yypt-1].node.(*SemicolonOpt),
				Token2:              yyS[yypt-0].Token,
			}
		}
	case 230:
		{
			yyVAL.node = &SwitchBody{
				Token:  yyS[yypt-1].Token,
				Token2: yyS[yypt-0].Token,
			}
		}
	case 231:
		{
			yyVAL.node = &SwitchBody{
				Case:           1,
				Token:          yyS[yypt-2].Token,
				SwitchCaseList: yyS[yypt-1].node.(*SwitchCaseList).reverse(),
				Token2:         yyS[yypt-0].Token,
			}
		}
	case 232:
		{
			yyVAL.node = &SwitchCase{
				Token:        yyS[yypt-2].Token,
				ArgumentList: yyS[yypt-1].node.(*ArgumentList).reverse(),
				Token2:       yyS[yypt-0].Token,
			}
		}
	case 233:
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
	case 234:
		{
			yyVAL.node = &SwitchCase{
				Case:         2,
				Token:        yyS[yypt-4].Token,
				ArgumentList: yyS[yypt-3].node.(*ArgumentList).reverse(),
				Token2:       yyS[yypt-2].Token,
				Expression:   yyS[yypt-1].node.(*Expression),
				Token3:       yyS[yypt-0].Token,
			}
		}
	case 235:
		{
			yyVAL.node = &SwitchCase{
				Case:   3,
				Token:  yyS[yypt-1].Token,
				Token2: yyS[yypt-0].Token,
			}
		}
	case 236:
		{
			yyVAL.node = &SwitchCase{
				Case:  4,
				Token: yyS[yypt-1].Token,
			}
		}
	case 237:
		{
			yyVAL.node = &SwitchCase{
				Case:  5,
				Token: yyS[yypt-1].Token,
			}
		}
	case 238:
		{
			yyVAL.node = &SwitchCaseBlock{
				SwitchCase:    yyS[yypt-1].node.(*SwitchCase),
				StatementList: yyS[yypt-0].node.(*StatementList).reverse(),
			}
		}
	case 239:
		{
			yyVAL.node = &SwitchCaseList{
				SwitchCaseBlock: yyS[yypt-0].node.(*SwitchCaseBlock),
			}
		}
	case 240:
		{
			yyVAL.node = &SwitchCaseList{
				Case:            1,
				SwitchCaseList:  yyS[yypt-1].node.(*SwitchCaseList),
				SwitchCaseBlock: yyS[yypt-0].node.(*SwitchCaseBlock),
			}
		}
	case 241:
		{
			yyVAL.node = &SwitchStatement{
				Token:      yyS[yypt-2].Token,
				IfHeader:   yyS[yypt-1].node.(*IfHeader),
				SwitchBody: yyS[yypt-0].node.(*SwitchBody),
			}
		}
	case 242:
		{
			yyVAL.node = &TopLevelDecl{
				ConstDecl: yyS[yypt-0].node.(*ConstDecl),
			}
		}
	case 243:
		{
			yyVAL.node = &TopLevelDecl{
				Case:     1,
				FuncDecl: yyS[yypt-0].node.(*FuncDecl),
			}
		}
	case 244:
		{
			yyVAL.node = &TopLevelDecl{
				Case:     2,
				TypeDecl: yyS[yypt-0].node.(*TypeDecl),
			}
		}
	case 245:
		{
			yyVAL.node = &TopLevelDecl{
				Case:    3,
				VarDecl: yyS[yypt-0].node.(*VarDecl),
			}
		}
	case 246:
		{
			yyVAL.node = &TopLevelDecl{
				Case:             4,
				StatementNonDecl: yyS[yypt-0].node.(*StatementNonDecl),
			}
		}
	case 247:
		{
			yyVAL.node = &TopLevelDecl{
				Case: 5,
			}
		}
	case 248:
		{
			yyVAL.node = (*TopLevelDeclList)(nil)
		}
	case 249:
		{
			yyVAL.node = &TopLevelDeclList{
				TopLevelDeclList: yyS[yypt-2].node.(*TopLevelDeclList),
				TopLevelDecl:     yyS[yypt-1].node.(*TopLevelDecl),
				Token:            yyS[yypt-0].Token,
			}
		}
	case 250:
		{
			yyVAL.node = &Typ{
				Token:  yyS[yypt-2].Token,
				Typ:    yyS[yypt-1].node.(*Typ),
				Token2: yyS[yypt-0].Token,
			}
		}
	case 251:
		{
			yyVAL.node = &Typ{
				Case:  1,
				Token: yyS[yypt-1].Token,
				Typ:   yyS[yypt-0].node.(*Typ),
			}
		}
	case 252:
		{
			yyVAL.node = &Typ{
				Case:      2,
				ArrayType: yyS[yypt-0].node.(*ArrayType),
			}
		}
	case 253:
		{
			yyVAL.node = &Typ{
				Case:     3,
				ChanType: yyS[yypt-0].node.(*ChanType),
			}
		}
	case 254:
		{
			yyVAL.node = &Typ{
				Case:     4,
				FuncType: yyS[yypt-0].node.(*FuncType),
			}
		}
	case 255:
		{
			yyVAL.node = &Typ{
				Case:          5,
				InterfaceType: yyS[yypt-0].node.(*InterfaceType),
			}
		}
	case 256:
		{
			yyVAL.node = &Typ{
				Case:    6,
				MapType: yyS[yypt-0].node.(*MapType),
			}
		}
	case 257:
		{
			yyVAL.node = &Typ{
				Case:                7,
				QualifiedIdent:      yyS[yypt-1].node.(*QualifiedIdent),
				GenericArgumentsOpt: yyS[yypt-0].node.(*GenericArgumentsOpt),
			}
		}
	case 258:
		{
			yyVAL.node = &Typ{
				Case:      8,
				SliceType: yyS[yypt-0].node.(*SliceType),
			}
		}
	case 259:
		{
			yyVAL.node = &Typ{
				Case:       9,
				StructType: yyS[yypt-0].node.(*StructType),
			}
		}
	case 260:
		{
			yyVAL.node = &TypeDecl{
				Token:  yyS[yypt-2].Token,
				Token2: yyS[yypt-1].Token,
				Token3: yyS[yypt-0].Token,
			}
		}
	case 261:
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
	case 262:
		{
			yyVAL.node = &TypeDecl{
				Case:     2,
				Token:    yyS[yypt-1].Token,
				TypeSpec: yyS[yypt-0].node.(*TypeSpec),
			}
		}
	case 263:
		{
			yyVAL.node = &TypeLiteral{
				Token:       yyS[yypt-1].Token,
				TypeLiteral: yyS[yypt-0].node.(*TypeLiteral),
			}
		}
	case 264:
		{
			yyVAL.node = &TypeLiteral{
				Case:      1,
				ArrayType: yyS[yypt-0].node.(*ArrayType),
			}
		}
	case 265:
		{
			yyVAL.node = &TypeLiteral{
				Case:     2,
				ChanType: yyS[yypt-0].node.(*ChanType),
			}
		}
	case 266:
		{
			yyVAL.node = &TypeLiteral{
				Case:     3,
				FuncType: yyS[yypt-0].node.(*FuncType),
			}
		}
	case 267:
		{
			yyVAL.node = &TypeLiteral{
				Case:          4,
				InterfaceType: yyS[yypt-0].node.(*InterfaceType),
			}
		}
	case 268:
		{
			yyVAL.node = &TypeLiteral{
				Case:    5,
				MapType: yyS[yypt-0].node.(*MapType),
			}
		}
	case 269:
		{
			yyVAL.node = &TypeLiteral{
				Case:      6,
				SliceType: yyS[yypt-0].node.(*SliceType),
			}
		}
	case 270:
		{
			yyVAL.node = &TypeLiteral{
				Case:       7,
				StructType: yyS[yypt-0].node.(*StructType),
			}
		}
	case 271:
		{
			yyVAL.node = &TypeSpec{
				Token:            yyS[yypt-2].Token,
				GenericParamsOpt: yyS[yypt-1].node.(*GenericParamsOpt),
				Typ:              yyS[yypt-0].node.(*Typ),
			}
		}
	case 272:
		{
			yyVAL.node = &TypeSpecList{
				TypeSpec: yyS[yypt-0].node.(*TypeSpec),
			}
		}
	case 273:
		{
			yyVAL.node = &TypeSpecList{
				Case:         1,
				TypeSpecList: yyS[yypt-2].node.(*TypeSpecList),
				Token:        yyS[yypt-1].Token,
				TypeSpec:     yyS[yypt-0].node.(*TypeSpec),
			}
		}
	case 274:
		{
			yyVAL.node = &UnaryExpression{
				Token:           yyS[yypt-1].Token,
				UnaryExpression: yyS[yypt-0].node.(*UnaryExpression),
			}
		}
	case 275:
		{
			yyVAL.node = &UnaryExpression{
				Case:            1,
				Token:           yyS[yypt-1].Token,
				UnaryExpression: yyS[yypt-0].node.(*UnaryExpression),
			}
		}
	case 276:
		{
			yyVAL.node = &UnaryExpression{
				Case:            2,
				Token:           yyS[yypt-1].Token,
				UnaryExpression: yyS[yypt-0].node.(*UnaryExpression),
			}
		}
	case 277:
		{
			yyVAL.node = &UnaryExpression{
				Case:            3,
				Token:           yyS[yypt-1].Token,
				UnaryExpression: yyS[yypt-0].node.(*UnaryExpression),
			}
		}
	case 278:
		{
			yyVAL.node = &UnaryExpression{
				Case:            4,
				Token:           yyS[yypt-1].Token,
				UnaryExpression: yyS[yypt-0].node.(*UnaryExpression),
			}
		}
	case 279:
		{
			yyVAL.node = &UnaryExpression{
				Case:            5,
				Token:           yyS[yypt-1].Token,
				UnaryExpression: yyS[yypt-0].node.(*UnaryExpression),
			}
		}
	case 280:
		{
			yyVAL.node = &UnaryExpression{
				Case:            6,
				Token:           yyS[yypt-1].Token,
				UnaryExpression: yyS[yypt-0].node.(*UnaryExpression),
			}
		}
	case 281:
		{
			yyVAL.node = &UnaryExpression{
				Case:              7,
				PrimaryExpression: yyS[yypt-0].node.(*PrimaryExpression),
			}
		}
	case 282:
		{
			yyVAL.node = &VarDecl{
				Token:  yyS[yypt-2].Token,
				Token2: yyS[yypt-1].Token,
				Token3: yyS[yypt-0].Token,
			}
		}
	case 283:
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
	case 284:
		{
			yyVAL.node = &VarDecl{
				Case:    2,
				Token:   yyS[yypt-1].Token,
				VarSpec: yyS[yypt-0].node.(*VarSpec),
			}
		}
	case 285:
		{
			yyVAL.node = &VarSpec{
				IdentifierList: yyS[yypt-2].node.(*IdentifierList).reverse(),
				Token:          yyS[yypt-1].Token,
				ExpressionList: yyS[yypt-0].node.(*ExpressionList).reverse(),
			}
		}
	case 286:
		{
			yyVAL.node = &VarSpec{
				Case:           1,
				IdentifierList: yyS[yypt-1].node.(*IdentifierList).reverse(),
				Typ:            yyS[yypt-0].node.(*Typ),
			}
		}
	case 287:
		{
			yyVAL.node = &VarSpec{
				Case:           2,
				IdentifierList: yyS[yypt-3].node.(*IdentifierList).reverse(),
				Typ:            yyS[yypt-2].node.(*Typ),
				Token:          yyS[yypt-1].Token,
				ExpressionList: yyS[yypt-0].node.(*ExpressionList).reverse(),
			}
		}
	case 288:
		{
			yyVAL.node = &VarSpecList{
				VarSpec: yyS[yypt-0].node.(*VarSpec),
			}
		}
	case 289:
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
