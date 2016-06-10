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
	yyDefault     = 57420
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
	yyTabOfs   = -291
)

var (
	yyXLAT = map[int]int{
		59:    0,   // ';' (251x)
		125:   1,   // '}' (242x)
		40:    2,   // '(' (220x)
		42:    3,   // '*' (219x)
		38:    4,   // '&' (183x)
		43:    5,   // '+' (183x)
		45:    6,   // '-' (183x)
		94:    7,   // '^' (183x)
		57358: 8,   // COMM (183x)
		57354: 9,   // CASE (182x)
		57363: 10,  // DEFAULT (182x)
		41:    11,  // ')' (171x)
		57377: 12,  // IDENTIFIER (168x)
		44:    13,  // ',' (167x)
		91:    14,  // '[' (167x)
		57403: 15,  // STRING_LIT (147x)
		57352: 16,  // BODY (144x)
		57355: 17,  // CHAN (135x)
		57372: 18,  // FUNC (134x)
		57382: 19,  // INTERFACE (134x)
		57388: 20,  // MAP (134x)
		57401: 21,  // RXCHAN (134x)
		57404: 22,  // STRUCT (134x)
		58:    23,  // ':' (116x)
		61:    24,  // '=' (114x)
		57423: 25,  // ArrayType (111x)
		57429: 26,  // ChanType (111x)
		57450: 27,  // FuncType (111x)
		57465: 28,  // InterfaceType (111x)
		57470: 29,  // MapType (111x)
		57487: 30,  // SliceType (111x)
		57494: 31,  // StructType (111x)
		57356: 32,  // CHAR_LIT (108x)
		57357: 33,  // COLAS (108x)
		57361: 34,  // DDD (108x)
		57370: 35,  // FLOAT_LIT (108x)
		57379: 36,  // IMAG_LIT (108x)
		57383: 37,  // INT_LIT (108x)
		33:    38,  // '!' (102x)
		93:    39,  // ']' (100x)
		123:   40,  // '{' (99x)
		57425: 41,  // BasicLiteral (88x)
		57433: 42,  // CompLitType (86x)
		57471: 43,  // Operand (86x)
		57476: 44,  // PrimaryExpression (86x)
		57504: 45,  // TypeLiteral (86x)
		57507: 46,  // UnaryExpression (86x)
		37:    47,  // '%' (81x)
		47:    48,  // '/' (81x)
		60:    49,  // '<' (81x)
		62:    50,  // '>' (81x)
		124:   51,  // '|' (81x)
		57347: 52,  // ANDAND (81x)
		57348: 53,  // ANDNOT (81x)
		57367: 54,  // EQ (81x)
		57373: 55,  // GEQ (81x)
		57384: 56,  // LEQ (81x)
		57385: 57,  // LSH (81x)
		57391: 58,  // NEQ (81x)
		57393: 59,  // OROR (81x)
		57399: 60,  // RSH (81x)
		57441: 61,  // Expression (79x)
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
		46:    75,  // '.' (39x)
		57376: 76,  // GTGT (38x)
		57442: 77,  // ExpressionList (34x)
		57478: 78,  // QualifiedIdent (33x)
		57345: 79,  // error (32x)
		57502: 80,  // Typ (25x)
		57378: 81,  // IF (24x)
		57408: 82,  // TYPE (24x)
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
		57424: 95,  // Assignment (14x)
		57485: 96,  // SimpleStatement (14x)
		57344: 97,  // $end (8x)
		57426: 98,  // Block (8x)
		57455: 99,  // IdentifierList (8x)
		57387: 100, // LTLT (8x)
		57435: 101, // ConstDecl (7x)
		57447: 102, // ForStatement (7x)
		57458: 103, // IfStatement (7x)
		57466: 104, // LBrace (7x)
		57482: 105, // SelectStatement (7x)
		57486: 106, // SimpleStatementOpt (7x)
		57490: 107, // StatementNonDecl (7x)
		57499: 108, // SwitchStatement (7x)
		57503: 109, // TypeDecl (7x)
		57508: 110, // VarDecl (7x)
		57430: 111, // CommaOpt (6x)
		57475: 112, // Parameters (6x)
		57483: 113, // SemicolonOpt (6x)
		57488: 114, // Statement (6x)
		57491: 115, // StringLitOpt (6x)
		57434: 116, // CompLitValue (5x)
		57366: 117, // ELSE (5x)
		57456: 118, // IdentifierOpt (5x)
		57421: 119, // Argument (4x)
		57380: 120, // IMPORT (4x)
		57469: 121, // LBraceCompLitValue (4x)
		57484: 122, // Signature (4x)
		57489: 123, // StatementList (4x)
		57427: 124, // Body (3x)
		57436: 125, // ConstSpec (3x)
		57444: 126, // ExpressionOpt (3x)
		57452: 127, // GenericArgumentsOpt (3x)
		57457: 128, // IfHeader (3x)
		57461: 129, // ImportSpec (3x)
		57473: 130, // ParameterDecl (3x)
		57397: 131, // RANGE (3x)
		57505: 132, // TypeSpec (3x)
		57509: 133, // VarSpec (3x)
		57422: 134, // ArgumentList (2x)
		57431: 135, // CompLitItem (2x)
		57454: 136, // GenericParametersOpt (2x)
		57463: 137, // InterfaceMethodDecl (2x)
		57467: 138, // LBraceCompLitItem (2x)
		57474: 139, // ParameterDeclList (2x)
		57481: 140, // ResultOpt (2x)
		57492: 141, // StructFieldDecl (2x)
		57495: 142, // SwitchBody (2x)
		57496: 143, // SwitchCase (2x)
		57497: 144, // SwitchCaseBlock (2x)
		57411: 145, // $@1 (1x)
		57412: 146, // $@2 (1x)
		57413: 147, // $@3 (1x)
		57414: 148, // $@4 (1x)
		57415: 149, // $@5 (1x)
		57416: 150, // $@6 (1x)
		57417: 151, // $@7 (1x)
		57418: 152, // $@8 (1x)
		57419: 153, // $@9 (1x)
		57428: 154, // Call (1x)
		57432: 155, // CompLitItemList (1x)
		57437: 156, // ConstSpecList (1x)
		57438: 157, // Elif (1x)
		57439: 158, // ElifList (1x)
		57440: 159, // ElseOpt (1x)
		57443: 160, // ExpressionListOpt (1x)
		57445: 161, // File (1x)
		57446: 162, // ForHeader (1x)
		57448: 163, // FuncBodyOpt (1x)
		57449: 164, // FuncDecl (1x)
		57451: 165, // GenericArgumentList (1x)
		57453: 166, // GenericParameterList (1x)
		57459: 167, // ImportDecl (1x)
		57460: 168, // ImportList (1x)
		57462: 169, // ImportSpecList (1x)
		57464: 170, // InterfaceMethodDeclList (1x)
		57468: 171, // LBraceCompLitItemList (1x)
		57395: 172, // PACKAGE (1x)
		57472: 173, // PackageClause (1x)
		57477: 174, // Prologue (1x)
		57479: 175, // Range (1x)
		57480: 176, // ReceiverOpt (1x)
		57493: 177, // StructFieldDeclList (1x)
		57498: 178, // SwitchCaseList (1x)
		57500: 179, // TopLevelDecl (1x)
		57501: 180, // TopLevelDeclList (1x)
		57407: 181, // TXCHAN (1x)
		57506: 182, // TypeSpecList (1x)
		57510: 183, // VarSpecList (1x)
		57420: 184, // $default (0x)
		57351: 185, // BAD_FLOAT_LIT (0x)
		57368: 186, // ERRCHECK (0x)
		57392: 187, // NO_RESULT (0x)
		57396: 188, // PARAMS (0x)
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
		"IDENTIFIER",
		"','",
		"'['",
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
		"FuncType",
		"InterfaceType",
		"MapType",
		"SliceType",
		"StructType",
		"CHAR_LIT",
		"COLAS",
		"DDD",
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
		"'.'",
		"GTGT",
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
		"IMPORT",
		"LBraceCompLitValue",
		"Signature",
		"StatementList",
		"Body",
		"ConstSpec",
		"ExpressionOpt",
		"GenericArgumentsOpt",
		"IfHeader",
		"ImportSpec",
		"ParameterDecl",
		"RANGE",
		"TypeSpec",
		"VarSpec",
		"ArgumentList",
		"CompLitItem",
		"GenericParametersOpt",
		"InterfaceMethodDecl",
		"LBraceCompLitItem",
		"ParameterDeclList",
		"ResultOpt",
		"StructFieldDecl",
		"SwitchBody",
		"SwitchCase",
		"SwitchCaseBlock",
		"$@1",
		"$@2",
		"$@3",
		"$@4",
		"$@5",
		"$@6",
		"$@7",
		"$@8",
		"$@9",
		"Call",
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
		"GenericArgumentList",
		"GenericParameterList",
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
		57356: "literal",
		57357: ":=",
		57361: "...",
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
		1:   {161, 2},
		2:   {119, 1},
		3:   {119, 1},
		4:   {134, 1},
		5:   {134, 3},
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
		25:  {145, 0},
		26:  {98, 4},
		27:  {146, 0},
		28:  {124, 4},
		29:  {154, 2},
		30:  {154, 4},
		31:  {154, 5},
		32:  {26, 2},
		33:  {26, 3},
		34:  {26, 3},
		35:  {111, 0},
		36:  {111, 1},
		37:  {135, 1},
		38:  {135, 3},
		39:  {135, 3},
		40:  {135, 1},
		41:  {135, 3},
		42:  {135, 3},
		43:  {155, 1},
		44:  {155, 3},
		45:  {42, 1},
		46:  {42, 1},
		47:  {42, 1},
		48:  {42, 1},
		49:  {116, 2},
		50:  {116, 4},
		51:  {101, 3},
		52:  {101, 5},
		53:  {101, 2},
		54:  {125, 1},
		55:  {125, 3},
		56:  {125, 4},
		57:  {156, 1},
		58:  {156, 3},
		59:  {157, 4},
		60:  {158, 0},
		61:  {158, 2},
		62:  {159, 0},
		63:  {159, 2},
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
		85:  {126, 0},
		86:  {126, 1},
		87:  {77, 1},
		88:  {77, 3},
		89:  {160, 0},
		90:  {160, 1},
		91:  {162, 1},
		92:  {162, 5},
		93:  {162, 1},
		94:  {102, 3},
		95:  {163, 0},
		96:  {163, 1},
		97:  {147, 0},
		98:  {164, 7},
		99:  {27, 2},
		100: {165, 1},
		101: {165, 3},
		102: {127, 0},
		103: {127, 3},
		104: {166, 1},
		105: {166, 3},
		106: {136, 0},
		107: {136, 3},
		108: {118, 0},
		109: {118, 1},
		110: {99, 1},
		111: {99, 3},
		112: {128, 1},
		113: {128, 3},
		114: {103, 5},
		115: {167, 3},
		116: {167, 5},
		117: {167, 2},
		118: {129, 2},
		119: {129, 2},
		120: {129, 3},
		121: {129, 3},
		122: {169, 1},
		123: {169, 3},
		124: {168, 0},
		125: {168, 3},
		126: {28, 3},
		127: {148, 0},
		128: {28, 6},
		129: {149, 0},
		130: {137, 3},
		131: {137, 1},
		132: {170, 1},
		133: {170, 3},
		134: {104, 1},
		135: {104, 1},
		136: {138, 1},
		137: {138, 3},
		138: {138, 3},
		139: {138, 1},
		140: {171, 1},
		141: {171, 3},
		142: {121, 2},
		143: {121, 4},
		144: {29, 5},
		145: {43, 3},
		146: {43, 3},
		147: {43, 1},
		148: {150, 0},
		149: {43, 5},
		150: {43, 2},
		151: {78, 1},
		152: {78, 3},
		153: {151, 0},
		154: {173, 4},
		155: {130, 2},
		156: {130, 3},
		157: {130, 2},
		158: {130, 1},
		159: {139, 1},
		160: {139, 3},
		161: {112, 2},
		162: {112, 4},
		163: {44, 1},
		164: {44, 2},
		165: {44, 5},
		166: {44, 5},
		167: {44, 3},
		168: {44, 4},
		169: {44, 8},
		170: {44, 6},
		171: {44, 2},
		172: {44, 2},
		173: {44, 5},
		174: {174, 2},
		175: {175, 4},
		176: {175, 4},
		177: {175, 2},
		178: {176, 0},
		179: {176, 1},
		180: {140, 0},
		181: {140, 1},
		182: {140, 1},
		183: {105, 2},
		184: {113, 0},
		185: {113, 1},
		186: {122, 2},
		187: {96, 1},
		188: {96, 1},
		189: {96, 2},
		190: {96, 2},
		191: {96, 3},
		192: {106, 0},
		193: {106, 1},
		194: {30, 3},
		195: {114, 0},
		196: {114, 1},
		197: {114, 1},
		198: {114, 1},
		199: {114, 1},
		200: {114, 1},
		201: {114, 1},
		202: {123, 1},
		203: {123, 3},
		204: {107, 2},
		205: {107, 2},
		206: {107, 2},
		207: {107, 1},
		208: {107, 1},
		209: {107, 2},
		210: {107, 2},
		211: {107, 3},
		212: {107, 1},
		213: {107, 2},
		214: {107, 1},
		215: {107, 1},
		216: {107, 1},
		217: {115, 0},
		218: {115, 1},
		219: {141, 3},
		220: {141, 3},
		221: {141, 2},
		222: {141, 4},
		223: {141, 5},
		224: {141, 5},
		225: {177, 1},
		226: {177, 3},
		227: {31, 3},
		228: {152, 0},
		229: {31, 6},
		230: {142, 2},
		231: {153, 0},
		232: {142, 4},
		233: {143, 3},
		234: {143, 5},
		235: {143, 5},
		236: {143, 2},
		237: {143, 2},
		238: {143, 2},
		239: {144, 2},
		240: {178, 1},
		241: {178, 2},
		242: {108, 3},
		243: {179, 1},
		244: {179, 1},
		245: {179, 1},
		246: {179, 1},
		247: {179, 1},
		248: {179, 1},
		249: {180, 0},
		250: {180, 3},
		251: {80, 3},
		252: {80, 2},
		253: {80, 1},
		254: {80, 1},
		255: {80, 1},
		256: {80, 1},
		257: {80, 1},
		258: {80, 2},
		259: {80, 1},
		260: {80, 1},
		261: {109, 3},
		262: {109, 5},
		263: {109, 2},
		264: {45, 2},
		265: {45, 1},
		266: {45, 1},
		267: {45, 1},
		268: {45, 1},
		269: {45, 1},
		270: {45, 1},
		271: {45, 1},
		272: {132, 3},
		273: {182, 1},
		274: {182, 3},
		275: {46, 2},
		276: {46, 2},
		277: {46, 2},
		278: {46, 2},
		279: {46, 2},
		280: {46, 2},
		281: {46, 2},
		282: {46, 1},
		283: {110, 3},
		284: {110, 5},
		285: {110, 2},
		286: {133, 3},
		287: {133, 2},
		288: {133, 4},
		289: {183, 1},
		290: {183, 3},
	}

	yyXErrors = map[yyXError]string{
		yyXError{0, 97}:    "invalid empty input",
		yyXError{1, -1}:    "expected $end",
		yyXError{57, -1}:   "expected '('",
		yyXError{26, -1}:   "expected ')'",
		yyXError{100, -1}:  "expected ')'",
		yyXError{174, -1}:  "expected ')'",
		yyXError{193, -1}:  "expected ')'",
		yyXError{211, -1}:  "expected ')'",
		yyXError{212, -1}:  "expected ')'",
		yyXError{229, -1}:  "expected ')'",
		yyXError{295, -1}:  "expected ')'",
		yyXError{297, -1}:  "expected ')'",
		yyXError{306, -1}:  "expected ')'",
		yyXError{338, -1}:  "expected ')'",
		yyXError{339, -1}:  "expected ')'",
		yyXError{360, -1}:  "expected ')'",
		yyXError{362, -1}:  "expected ')'",
		yyXError{482, -1}:  "expected ')'",
		yyXError{327, -1}:  "expected ':'",
		yyXError{7, -1}:    "expected ';'",
		yyXError{11, -1}:   "expected ';'",
		yyXError{23, -1}:   "expected ';'",
		yyXError{29, -1}:   "expected ';'",
		yyXError{31, -1}:   "expected ';'",
		yyXError{74, -1}:   "expected ';'",
		yyXError{75, -1}:   "expected ';'",
		yyXError{76, -1}:   "expected ';'",
		yyXError{77, -1}:   "expected ';'",
		yyXError{78, -1}:   "expected ';'",
		yyXError{79, -1}:   "expected ';'",
		yyXError{80, -1}:   "expected ';'",
		yyXError{426, -1}:  "expected ';'",
		yyXError{427, -1}:  "expected ';'",
		yyXError{436, -1}:  "expected ';'",
		yyXError{475, -1}:  "expected '='",
		yyXError{49, -1}:   "expected '['",
		yyXError{333, -1}:  "expected ']'",
		yyXError{390, -1}:  "expected ']'",
		yyXError{493, -1}:  "expected ']'",
		yyXError{292, -1}:  "expected '}'",
		yyXError{348, -1}:  "expected '}'",
		yyXError{373, -1}:  "expected '}'",
		yyXError{402, -1}:  "expected '}'",
		yyXError{320, -1}:  "expected argument list or one of ['!', '&', '(', ')', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{47, -1}:   "expected block statement or if statement header or one of ['!', '&', '(', '*', '+', '-', ';', '[', '^', <-, chan, func, identifier, interface, literal, map, struct, {]",
		yyXError{413, -1}:  "expected block statement or if statement header or one of ['!', '&', '(', '*', '+', '-', ';', '[', '^', <-, chan, func, identifier, interface, literal, map, struct, {]",
		yyXError{410, -1}:  "expected block statement or one of ['{', if]",
		yyXError{406, -1}:  "expected block statement or {",
		yyXError{415, -1}:  "expected block statement or {",
		yyXError{431, -1}:  "expected block statement or {",
		yyXError{73, -1}:   "expected body of the switch statement or if statement header or one of ['!', '&', '(', '*', '+', '-', ';', '[', '^', <-, chan, func, identifier, interface, literal, map, struct, {]",
		yyXError{58, -1}:   "expected body of the switch statement or {",
		yyXError{235, -1}:  "expected body of the switch statement or {",
		yyXError{56, -1}:   "expected call or composite literal value or one of [!=, %=, &&, &=, &^, &^=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '[', ']', '^', '{', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{37, -1}:   "expected chan",
		yyXError{365, -1}:  "expected composite literal item list or one of ['!', '&', '(', '*', '+', '-', '[', '^', '{', '}', <-, chan, func, identifier, interface, literal, map, struct, {]",
		yyXError{321, -1}:  "expected composite literal item list or one of ['!', '&', '(', '*', '+', '-', '[', '^', '{', '}', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{372, -1}:  "expected composite literal item or one of ['!', '&', '(', '*', '+', '-', '[', '^', '{', '}', <-, chan, func, identifier, interface, literal, map, struct, {]",
		yyXError{347, -1}:  "expected composite literal item or one of ['!', '&', '(', '*', '+', '-', '[', '^', '{', '}', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{376, -1}:  "expected composite literal value or expression or one of ['!', '&', '(', '*', '+', '-', '[', '^', '{', <-, chan, func, identifier, interface, literal, map, struct, {]",
		yyXError{351, -1}:  "expected composite literal value or expression or one of ['!', '&', '(', '*', '+', '-', '[', '^', '{', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{354, -1}:  "expected composite literal value or expression or one of ['!', '&', '(', '*', '+', '-', '[', '^', '{', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{55, -1}:   "expected composite literal value or one of ['{', {]",
		yyXError{471, -1}:  "expected constant specification list or one of [')', identifier]",
		yyXError{42, -1}:   "expected constant specification or one of ['(', identifier]",
		yyXError{483, -1}:  "expected constant specification or one of [')', identifier]",
		yyXError{409, -1}:  "expected else if clause or optional else clause or one of [';', '}', case, default, else]",
		yyXError{408, -1}:  "expected else if list clause or optional else clause or one of [';', '}', case, default, else]",
		yyXError{439, -1}:  "expected expression list or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, range, struct]",
		yyXError{451, -1}:  "expected expression list or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, range, struct]",
		yyXError{116, -1}:  "expected expression list or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{118, -1}:  "expected expression list or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{440, -1}:  "expected expression list or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{441, -1}:  "expected expression list or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{442, -1}:  "expected expression list or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{443, -1}:  "expected expression list or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{444, -1}:  "expected expression list or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{445, -1}:  "expected expression list or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{446, -1}:  "expected expression list or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{447, -1}:  "expected expression list or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{448, -1}:  "expected expression list or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{449, -1}:  "expected expression list or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{450, -1}:  "expected expression list or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{474, -1}:  "expected expression list or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{476, -1}:  "expected expression list or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{491, -1}:  "expected expression list or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{492, -1}:  "expected expression list or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{34, -1}:   "expected expression or one of ['!', '&', '(', '*', '+', '-', '[', ']', '^', ..., <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{62, -1}:   "expected expression or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{65, -1}:   "expected expression or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{129, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{131, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{132, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{133, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{134, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{135, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{136, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{137, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{138, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{139, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{140, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{141, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{142, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{143, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{144, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{145, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{146, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{147, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{148, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{149, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{150, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{209, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{268, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{269, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{432, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{453, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{467, -1}:  "expected expression or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{323, -1}:  "expected expression or optional expression or one of ['!', '&', '(', '*', '+', '-', ':', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{50, -1}:   "expected expression or type literal or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{359, -1}:  "expected expression/type literal or one of ['!', '&', '(', ')', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{266, -1}:  "expected expression/type literal or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{45, -1}:   "expected for statement header or one of ['!', '&', '(', '*', '+', '-', ';', '[', '^', <-, chan, func, identifier, interface, literal, map, range, struct, {]",
		yyXError{103, -1}:  "expected function/method signature or '('",
		yyXError{400, -1}:  "expected function/method signature or '('",
		yyXError{423, -1}:  "expected function/method signature or '('",
		yyXError{397, -1}:  "expected function/method signature or one of ['(', '.', ';', '}']",
		yyXError{422, -1}:  "expected function/method signature or optional generic parameters or one of ['(', «]",
		yyXError{46, -1}:   "expected function/method signature or optional receiver or one of ['(', identifier]",
		yyXError{122, -1}:  "expected generic argument list or one of ['(', '*', '[', <-, chan, func, identifier, interface, map, struct]",
		yyXError{217, -1}:  "expected generic parameter list or identifier",
		yyXError{3, -1}:    "expected identifier",
		yyXError{30, -1}:   "expected identifier",
		yyXError{66, -1}:   "expected identifier",
		yyXError{104, -1}:  "expected identifier",
		yyXError{176, -1}:  "expected identifier",
		yyXError{222, -1}:  "expected identifier",
		yyXError{420, -1}:  "expected identifier",
		yyXError{10, -1}:   "expected import specification list or one of [')', '.', identifier, literal]",
		yyXError{6, -1}:    "expected import specification or one of ['(', '.', identifier, literal]",
		yyXError{27, -1}:   "expected import specification or one of [')', '.', identifier, literal]",
		yyXError{395, -1}:  "expected interface method declaration list or identifier",
		yyXError{393, -1}:  "expected interface method declaration list or one of ['}', identifier]",
		yyXError{403, -1}:  "expected interface method declaration or one of ['}', identifier]",
		yyXError{52, -1}:   "expected left brace or one of ['(', ')', ',', ':', '=', '{', ..., :=, {]",
		yyXError{48, -1}:   "expected left brace or one of ['{', {]",
		yyXError{72, -1}:   "expected left brace or one of ['{', {]",
		yyXError{381, -1}:  "expected left brace or one of ['{', {]",
		yyXError{12, -1}:   "expected literal or ",
		yyXError{13, -1}:   "expected literal or ",
		yyXError{127, -1}:  "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '[', ']', '^', '{', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, literal, {, |=, ||, »]",
		yyXError{51, -1}:   "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '[', ']', '^', '{', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{54, -1}:   "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '[', ']', '^', '{', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{123, -1}:  "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '[', ']', '^', '{', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{213, -1}:  "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '[', ']', '^', '{', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{324, -1}:  "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '[', ']', '^', '{', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{325, -1}:  "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '[', ']', '^', '{', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{332, -1}:  "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '[', ']', '^', '{', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{334, -1}:  "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '[', ']', '^', '{', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{335, -1}:  "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '[', ']', '^', '{', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{337, -1}:  "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '[', ']', '^', '{', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{340, -1}:  "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '[', ']', '^', '{', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{341, -1}:  "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '[', ']', '^', '{', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{346, -1}:  "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '[', ']', '^', '{', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{349, -1}:  "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '[', ']', '^', '{', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{358, -1}:  "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '[', ']', '^', '{', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{363, -1}:  "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '[', ']', '^', '{', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{364, -1}:  "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '[', ']', '^', '{', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{366, -1}:  "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '[', ']', '^', '{', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{371, -1}:  "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '[', ']', '^', '{', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{374, -1}:  "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '[', ']', '^', '{', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{384, -1}:  "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '[', ']', '^', '{', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{387, -1}:  "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '[', ']', '^', '{', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{388, -1}:  "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '[', ']', '^', '{', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{43, -1}:   "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', ')', '*', '+', ',', '-', '/', ':', ';', '<', '=', '>', ']', '^', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{151, -1}:  "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', ')', '*', '+', ',', '-', '/', ':', ';', '<', '=', '>', ']', '^', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{152, -1}:  "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', ')', '*', '+', ',', '-', '/', ':', ';', '<', '=', '>', ']', '^', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{153, -1}:  "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', ')', '*', '+', ',', '-', '/', ':', ';', '<', '=', '>', ']', '^', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{154, -1}:  "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', ')', '*', '+', ',', '-', '/', ':', ';', '<', '=', '>', ']', '^', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{155, -1}:  "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', ')', '*', '+', ',', '-', '/', ':', ';', '<', '=', '>', ']', '^', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{156, -1}:  "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', ')', '*', '+', ',', '-', '/', ':', ';', '<', '=', '>', ']', '^', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{157, -1}:  "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', ')', '*', '+', ',', '-', '/', ':', ';', '<', '=', '>', ']', '^', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{158, -1}:  "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', ')', '*', '+', ',', '-', '/', ':', ';', '<', '=', '>', ']', '^', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{159, -1}:  "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', ')', '*', '+', ',', '-', '/', ':', ';', '<', '=', '>', ']', '^', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{160, -1}:  "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', ')', '*', '+', ',', '-', '/', ':', ';', '<', '=', '>', ']', '^', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{161, -1}:  "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', ')', '*', '+', ',', '-', '/', ':', ';', '<', '=', '>', ']', '^', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{162, -1}:  "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', ')', '*', '+', ',', '-', '/', ':', ';', '<', '=', '>', ']', '^', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{163, -1}:  "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', ')', '*', '+', ',', '-', '/', ':', ';', '<', '=', '>', ']', '^', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{164, -1}:  "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', ')', '*', '+', ',', '-', '/', ':', ';', '<', '=', '>', ']', '^', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{165, -1}:  "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', ')', '*', '+', ',', '-', '/', ':', ';', '<', '=', '>', ']', '^', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{166, -1}:  "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', ')', '*', '+', ',', '-', '/', ':', ';', '<', '=', '>', ']', '^', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{167, -1}:  "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', ')', '*', '+', ',', '-', '/', ':', ';', '<', '=', '>', ']', '^', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{168, -1}:  "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', ')', '*', '+', ',', '-', '/', ':', ';', '<', '=', '>', ']', '^', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{169, -1}:  "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', ')', '*', '+', ',', '-', '/', ':', ';', '<', '=', '>', ']', '^', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{170, -1}:  "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', ')', '*', '+', ',', '-', '/', ':', ';', '<', '=', '>', ']', '^', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{201, -1}:  "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', ')', '*', '+', ',', '-', '/', ':', ';', '<', '=', '>', ']', '^', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{202, -1}:  "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', ')', '*', '+', ',', '-', '/', ':', ';', '<', '=', '>', ']', '^', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{203, -1}:  "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', ')', '*', '+', ',', '-', '/', ':', ';', '<', '=', '>', ']', '^', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{204, -1}:  "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', ')', '*', '+', ',', '-', '/', ':', ';', '<', '=', '>', ']', '^', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{205, -1}:  "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', ')', '*', '+', ',', '-', '/', ':', ';', '<', '=', '>', ']', '^', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{206, -1}:  "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', ')', '*', '+', ',', '-', '/', ':', ';', '<', '=', '>', ']', '^', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{208, -1}:  "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', ')', '*', '+', ',', '-', '/', ':', ';', '<', '=', '>', ']', '^', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{130, -1}:  "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', ')', '*', '+', ',', '-', '/', ';', '<', '=', '>', '^', '|', '}', *=, +=, -=, /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{44, -1}:   "expected one of [!=, %=, &&, &=, &^, &^=, '%', '&', '*', '+', ',', '-', '/', ';', '<', '=', '>', '^', '|', '}', *=, ++, +=, --, -=, /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||]",
		yyXError{261, -1}:  "expected one of [!=, &&, &^, '%', '&', ')', '*', '+', ',', '-', '/', ':', '<', '=', '>', '^', '|', ..., :=, <-, <<, <=, ==, >=, >>, ||]",
		yyXError{119, -1}:  "expected one of [!=, &&, &^, '%', '&', ')', '*', '+', ',', '-', '/', ';', '<', '>', '^', '|', '}', <-, <<, <=, ==, >=, >>, case, default, {, ||]",
		yyXError{385, -1}:  "expected one of [!=, &&, &^, '%', '&', ')', '*', '+', '-', '/', '<', '>', '^', '|', <-, <<, <=, ==, >=, >>, ||]",
		yyXError{343, -1}:  "expected one of [!=, &&, &^, '%', '&', '*', '+', ',', '-', '/', ':', '<', '>', '^', '|', '}', <-, <<, <=, ==, >=, >>, ||]",
		yyXError{367, -1}:  "expected one of [!=, &&, &^, '%', '&', '*', '+', ',', '-', '/', ':', '<', '>', '^', '|', '}', <-, <<, <=, ==, >=, >>, ||]",
		yyXError{353, -1}:  "expected one of [!=, &&, &^, '%', '&', '*', '+', ',', '-', '/', '<', '>', '^', '|', '}', <-, <<, <=, ==, >=, >>, ||]",
		yyXError{356, -1}:  "expected one of [!=, &&, &^, '%', '&', '*', '+', ',', '-', '/', '<', '>', '^', '|', '}', <-, <<, <=, ==, >=, >>, ||]",
		yyXError{377, -1}:  "expected one of [!=, &&, &^, '%', '&', '*', '+', ',', '-', '/', '<', '>', '^', '|', '}', <-, <<, <=, ==, >=, >>, ||]",
		yyXError{326, -1}:  "expected one of [!=, &&, &^, '%', '&', '*', '+', '-', '/', ':', '<', '>', ']', '^', '|', <-, <<, <=, ==, >=, >>, ||]",
		yyXError{329, -1}:  "expected one of [!=, &&, &^, '%', '&', '*', '+', '-', '/', ':', '<', '>', ']', '^', '|', <-, <<, <=, ==, >=, >>, ||]",
		yyXError{270, -1}:  "expected one of [!=, &&, &^, '%', '&', '*', '+', '-', '/', ':', '<', '>', '^', '|', <-, <<, <=, ==, >=, >>, ||]",
		yyXError{272, -1}:  "expected one of [!=, &&, &^, '%', '&', '*', '+', '-', '/', ':', '<', '>', '^', '|', <-, <<, <=, ==, >=, >>, ||]",
		yyXError{315, -1}:  "expected one of [!=, &&, &^, '%', '&', '*', '+', '-', '/', ';', '<', '>', '^', '|', '}', <-, <<, <=, ==, >=, >>, case, default, ||]",
		yyXError{316, -1}:  "expected one of [!=, &&, &^, '%', '&', '*', '+', '-', '/', ';', '<', '>', '^', '|', '}', <-, <<, <=, ==, >=, >>, case, default, ||]",
		yyXError{494, -1}:  "expected one of [!=, &&, &^, '%', '&', '*', '+', '-', '/', '<', '>', ']', '^', '|', <-, <<, <=, ==, >=, >>, ||]",
		yyXError{433, -1}:  "expected one of [!=, &&, &^, '%', '&', '*', '+', '-', '/', '<', '>', '^', '|', <-, <<, <=, ==, >=, >>, {, ||]",
		yyXError{454, -1}:  "expected one of [!=, &&, &^, '%', '&', '*', '+', '-', '/', '<', '>', '^', '|', <-, <<, <=, ==, >=, >>, {, ||]",
		yyXError{468, -1}:  "expected one of [!=, &&, &^, '%', '&', '*', '+', '-', '/', '<', '>', '^', '|', <-, <<, <=, ==, >=, >>, {, ||]",
		yyXError{35, -1}:   "expected one of [%=, &=, &^=, ',', '=', *=, +=, -=, /=, :=, <<=, >>=, ^=, |=]",
		yyXError{428, -1}:  "expected one of [%=, &=, &^=, ',', '=', *=, +=, -=, /=, :=, <<=, >>=, ^=, |=]",
		yyXError{191, -1}:  "expected one of ['(', ')', '*', ',', ':', ';', '=', '[', ']', '{', '}', ..., :=, <-, case, chan, default, func, identifier, interface, literal, map, struct, {, »]",
		yyXError{194, -1}:  "expected one of ['(', ')', '*', ',', ':', ';', '=', '[', ']', '{', '}', ..., :=, <-, case, chan, default, func, identifier, interface, literal, map, struct, {, »]",
		yyXError{92, -1}:   "expected one of ['(', ')', '*', ',', ';', '=', '[', '}', <-, case, chan, default, func, identifier, interface, map, struct]",
		yyXError{178, -1}:  "expected one of ['(', ')', '*', ',', ';', '=', '[', '}', <-, case, chan, default, func, identifier, interface, map, struct]",
		yyXError{105, -1}:  "expected one of ['(', ')', ',', '.', ':', ';', '=', ']', '{', '}', ..., :=, case, default, literal, {, «, »]",
		yyXError{177, -1}:  "expected one of ['(', ')', ',', ':', ';', '=', ']', '{', '}', ..., :=, case, default, literal, {, «, »]",
		yyXError{108, -1}:  "expected one of ['(', ')', ',', ':', ';', '=', ']', '{', '}', ..., :=, case, default, literal, {, »]",
		yyXError{109, -1}:  "expected one of ['(', ')', ',', ':', ';', '=', ']', '{', '}', ..., :=, case, default, literal, {, »]",
		yyXError{110, -1}:  "expected one of ['(', ')', ',', ':', ';', '=', ']', '{', '}', ..., :=, case, default, literal, {, »]",
		yyXError{111, -1}:  "expected one of ['(', ')', ',', ':', ';', '=', ']', '{', '}', ..., :=, case, default, literal, {, »]",
		yyXError{112, -1}:  "expected one of ['(', ')', ',', ':', ';', '=', ']', '{', '}', ..., :=, case, default, literal, {, »]",
		yyXError{114, -1}:  "expected one of ['(', ')', ',', ':', ';', '=', ']', '{', '}', ..., :=, case, default, literal, {, »]",
		yyXError{115, -1}:  "expected one of ['(', ')', ',', ':', ';', '=', ']', '{', '}', ..., :=, case, default, literal, {, »]",
		yyXError{172, -1}:  "expected one of ['(', ')', ',', ':', ';', '=', ']', '{', '}', ..., :=, case, default, literal, {, »]",
		yyXError{173, -1}:  "expected one of ['(', ')', ',', ':', ';', '=', ']', '{', '}', ..., :=, case, default, literal, {, »]",
		yyXError{175, -1}:  "expected one of ['(', ')', ',', ':', ';', '=', ']', '{', '}', ..., :=, case, default, literal, {, »]",
		yyXError{179, -1}:  "expected one of ['(', ')', ',', ':', ';', '=', ']', '{', '}', ..., :=, case, default, literal, {, »]",
		yyXError{183, -1}:  "expected one of ['(', ')', ',', ':', ';', '=', ']', '{', '}', ..., :=, case, default, literal, {, »]",
		yyXError{184, -1}:  "expected one of ['(', ')', ',', ':', ';', '=', ']', '{', '}', ..., :=, case, default, literal, {, »]",
		yyXError{185, -1}:  "expected one of ['(', ')', ',', ':', ';', '=', ']', '{', '}', ..., :=, case, default, literal, {, »]",
		yyXError{282, -1}:  "expected one of ['(', ')', ',', ':', ';', '=', ']', '{', '}', ..., :=, case, default, literal, {, »]",
		yyXError{293, -1}:  "expected one of ['(', ')', ',', ':', ';', '=', ']', '{', '}', ..., :=, case, default, literal, {, »]",
		yyXError{392, -1}:  "expected one of ['(', ')', ',', ':', ';', '=', ']', '{', '}', ..., :=, case, default, literal, {, »]",
		yyXError{394, -1}:  "expected one of ['(', ')', ',', ':', ';', '=', ']', '{', '}', ..., :=, case, default, literal, {, »]",
		yyXError{405, -1}:  "expected one of ['(', ')', ',', ':', ';', '=', ']', '{', '}', ..., :=, case, default, literal, {, »]",
		yyXError{487, -1}:  "expected one of ['(', ')', ',', ':', ';', '=', ']', '{', '}', ..., :=, case, default, literal, {, »]",
		yyXError{488, -1}:  "expected one of ['(', ')', ',', ':', ';', '=', ']', '{', '}', ..., :=, case, default, literal, {, »]",
		yyXError{490, -1}:  "expected one of ['(', ')', ',', ':', ';', '=', ']', '{', '}', ..., :=, case, default, literal, {, »]",
		yyXError{496, -1}:  "expected one of ['(', ')', ',', ':', ';', '=', ']', '{', '}', ..., :=, case, default, literal, {, »]",
		yyXError{498, -1}:  "expected one of ['(', ')', ',', ':', ';', '=', ']', '{', '}', ..., :=, case, default, literal, {, »]",
		yyXError{500, -1}:  "expected one of ['(', ')', ',', ':', ';', '=', ']', '{', '}', ..., :=, case, default, literal, {, »]",
		yyXError{38, -1}:   "expected one of ['(', ')', ',', ':', '=', '{', ..., :=, {]",
		yyXError{39, -1}:   "expected one of ['(', ')', ',', ':', '=', '{', ..., :=, {]",
		yyXError{40, -1}:   "expected one of ['(', ')', ',', ':', '=', '{', ..., :=, {]",
		yyXError{41, -1}:   "expected one of ['(', ')', ',', ':', '=', '{', ..., :=, {]",
		yyXError{83, -1}:   "expected one of ['(', ')', ',', ':', '=', ..., :=]",
		yyXError{84, -1}:   "expected one of ['(', ')', ',', ':', '=', ..., :=]",
		yyXError{207, -1}:  "expected one of ['(', ')', ',', ':', '=', ..., :=]",
		yyXError{262, -1}:  "expected one of ['(', ')', ',', ':', '=', ..., :=]",
		yyXError{386, -1}:  "expected one of ['(', ')']",
		yyXError{284, -1}:  "expected one of ['(', '*', ',', '.', ';', '[', '}', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{223, -1}:  "expected one of ['(', '*', '[', <-, chan, func, identifier, interface, map, struct]",
		yyXError{322, -1}:  "expected one of ['(', identifier]",
		yyXError{263, -1}:  "expected one of [')', ',', ':', '=', ..., :=]",
		yyXError{274, -1}:  "expected one of [')', ',', ':', '=', ..., :=]",
		yyXError{120, -1}:  "expected one of [')', ',', ';', '}', case, default]",
		yyXError{171, -1}:  "expected one of [')', ',', ';', '}', case, default]",
		yyXError{477, -1}:  "expected one of [')', ',', ';', '}', case, default]",
		yyXError{478, -1}:  "expected one of [')', ',', ';', '}', case, default]",
		yyXError{188, -1}:  "expected one of [')', ',']",
		yyXError{189, -1}:  "expected one of [')', ',']",
		yyXError{195, -1}:  "expected one of [')', ',']",
		yyXError{196, -1}:  "expected one of [')', ',']",
		yyXError{197, -1}:  "expected one of [')', ',']",
		yyXError{199, -1}:  "expected one of [')', ',']",
		yyXError{200, -1}:  "expected one of [')', ',']",
		yyXError{117, -1}:  "expected one of [')', ';', '=', '}', case, default]",
		yyXError{219, -1}:  "expected one of [')', ';', '}', case, default]",
		yyXError{20, -1}:   "expected one of [')', ';']",
		yyXError{22, -1}:   "expected one of [')', ';']",
		yyXError{25, -1}:   "expected one of [')', ';']",
		yyXError{28, -1}:   "expected one of [')', ';']",
		yyXError{98, -1}:   "expected one of [')', ';']",
		yyXError{102, -1}:  "expected one of [')', ';']",
		yyXError{227, -1}:  "expected one of [')', ';']",
		yyXError{231, -1}:  "expected one of [')', ';']",
		yyXError{481, -1}:  "expected one of [')', ';']",
		yyXError{484, -1}:  "expected one of [')', ';']",
		yyXError{264, -1}:  "expected one of [',', ':', '=', :=]",
		yyXError{342, -1}:  "expected one of [',', ':', '}']",
		yyXError{452, -1}:  "expected one of [',', ';', '}', case, default, {]",
		yyXError{455, -1}:  "expected one of [',', ';', '}', case, default, {]",
		yyXError{456, -1}:  "expected one of [',', ';', '}', case, default, {]",
		yyXError{457, -1}:  "expected one of [',', ';', '}', case, default, {]",
		yyXError{458, -1}:  "expected one of [',', ';', '}', case, default, {]",
		yyXError{459, -1}:  "expected one of [',', ';', '}', case, default, {]",
		yyXError{460, -1}:  "expected one of [',', ';', '}', case, default, {]",
		yyXError{461, -1}:  "expected one of [',', ';', '}', case, default, {]",
		yyXError{462, -1}:  "expected one of [',', ';', '}', case, default, {]",
		yyXError{463, -1}:  "expected one of [',', ';', '}', case, default, {]",
		yyXError{464, -1}:  "expected one of [',', ';', '}', case, default, {]",
		yyXError{465, -1}:  "expected one of [',', ';', '}', case, default, {]",
		yyXError{466, -1}:  "expected one of [',', ';', '}', case, default, {]",
		yyXError{312, -1}:  "expected one of [',', ';', '}', case, default]",
		yyXError{344, -1}:  "expected one of [',', '}']",
		yyXError{350, -1}:  "expected one of [',', '}']",
		yyXError{352, -1}:  "expected one of [',', '}']",
		yyXError{355, -1}:  "expected one of [',', '}']",
		yyXError{368, -1}:  "expected one of [',', '}']",
		yyXError{369, -1}:  "expected one of [',', '}']",
		yyXError{375, -1}:  "expected one of [',', '}']",
		yyXError{378, -1}:  "expected one of [',', '}']",
		yyXError{124, -1}:  "expected one of [',', »]",
		yyXError{125, -1}:  "expected one of [',', »]",
		yyXError{128, -1}:  "expected one of [',', »]",
		yyXError{220, -1}:  "expected one of [',', »]",
		yyXError{221, -1}:  "expected one of [',', »]",
		yyXError{224, -1}:  "expected one of [',', »]",
		yyXError{330, -1}:  "expected one of [':', ']']",
		yyXError{411, -1}:  "expected one of [';', '}', case, default, else]",
		yyXError{416, -1}:  "expected one of [';', '}', case, default, else]",
		yyXError{419, -1}:  "expected one of [';', '}', case, default, else]",
		yyXError{9, -1}:    "expected one of [';', '}', case, default, literal]",
		yyXError{59, -1}:   "expected one of [';', '}', case, default, {]",
		yyXError{469, -1}:  "expected one of [';', '}', case, default, {]",
		yyXError{470, -1}:  "expected one of [';', '}', case, default, {]",
		yyXError{63, -1}:   "expected one of [';', '}', case, default]",
		yyXError{64, -1}:   "expected one of [';', '}', case, default]",
		yyXError{67, -1}:   "expected one of [';', '}', case, default]",
		yyXError{69, -1}:   "expected one of [';', '}', case, default]",
		yyXError{70, -1}:   "expected one of [';', '}', case, default]",
		yyXError{71, -1}:   "expected one of [';', '}', case, default]",
		yyXError{95, -1}:   "expected one of [';', '}', case, default]",
		yyXError{96, -1}:   "expected one of [';', '}', case, default]",
		yyXError{101, -1}:  "expected one of [';', '}', case, default]",
		yyXError{215, -1}:  "expected one of [';', '}', case, default]",
		yyXError{225, -1}:  "expected one of [';', '}', case, default]",
		yyXError{230, -1}:  "expected one of [';', '}', case, default]",
		yyXError{237, -1}:  "expected one of [';', '}', case, default]",
		yyXError{238, -1}:  "expected one of [';', '}', case, default]",
		yyXError{246, -1}:  "expected one of [';', '}', case, default]",
		yyXError{247, -1}:  "expected one of [';', '}', case, default]",
		yyXError{248, -1}:  "expected one of [';', '}', case, default]",
		yyXError{249, -1}:  "expected one of [';', '}', case, default]",
		yyXError{250, -1}:  "expected one of [';', '}', case, default]",
		yyXError{251, -1}:  "expected one of [';', '}', case, default]",
		yyXError{252, -1}:  "expected one of [';', '}', case, default]",
		yyXError{253, -1}:  "expected one of [';', '}', case, default]",
		yyXError{255, -1}:  "expected one of [';', '}', case, default]",
		yyXError{258, -1}:  "expected one of [';', '}', case, default]",
		yyXError{275, -1}:  "expected one of [';', '}', case, default]",
		yyXError{313, -1}:  "expected one of [';', '}', case, default]",
		yyXError{314, -1}:  "expected one of [';', '}', case, default]",
		yyXError{317, -1}:  "expected one of [';', '}', case, default]",
		yyXError{318, -1}:  "expected one of [';', '}', case, default]",
		yyXError{319, -1}:  "expected one of [';', '}', case, default]",
		yyXError{380, -1}:  "expected one of [';', '}', case, default]",
		yyXError{412, -1}:  "expected one of [';', '}', case, default]",
		yyXError{414, -1}:  "expected one of [';', '}', case, default]",
		yyXError{434, -1}:  "expected one of [';', '}', case, default]",
		yyXError{472, -1}:  "expected one of [';', '}', case, default]",
		yyXError{479, -1}:  "expected one of [';', '}', case, default]",
		yyXError{485, -1}:  "expected one of [';', '}', case, default]",
		yyXError{257, -1}:  "expected one of [';', '}']",
		yyXError{289, -1}:  "expected one of [';', '}']",
		yyXError{294, -1}:  "expected one of [';', '}']",
		yyXError{299, -1}:  "expected one of [';', '}']",
		yyXError{300, -1}:  "expected one of [';', '}']",
		yyXError{302, -1}:  "expected one of [';', '}']",
		yyXError{303, -1}:  "expected one of [';', '}']",
		yyXError{308, -1}:  "expected one of [';', '}']",
		yyXError{309, -1}:  "expected one of [';', '}']",
		yyXError{311, -1}:  "expected one of [';', '}']",
		yyXError{383, -1}:  "expected one of [';', '}']",
		yyXError{398, -1}:  "expected one of [';', '}']",
		yyXError{399, -1}:  "expected one of [';', '}']",
		yyXError{401, -1}:  "expected one of [';', '}']",
		yyXError{404, -1}:  "expected one of [';', '}']",
		yyXError{418, -1}:  "expected one of [';', '}']",
		yyXError{233, -1}:  "expected one of [';', {]",
		yyXError{234, -1}:  "expected one of [';', {]",
		yyXError{430, -1}:  "expected one of [';', {]",
		yyXError{244, -1}:  "expected one of ['}', case, default]",
		yyXError{276, -1}:  "expected one of ['}', case, default]",
		yyXError{210, -1}:  "expected optional comma or one of [!=, &&, &^, '%', '&', ')', '*', '+', ',', '-', '/', '<', '>', '^', '|', <-, <<, <=, ==, >=, >>, ||]",
		yyXError{357, -1}:  "expected optional comma or one of [')', ',', ...]",
		yyXError{190, -1}:  "expected optional comma or one of [')', ',']",
		yyXError{361, -1}:  "expected optional comma or one of [')', ',']",
		yyXError{345, -1}:  "expected optional comma or one of [',', '}']",
		yyXError{370, -1}:  "expected optional comma or one of [',', '}']",
		yyXError{68, -1}:   "expected optional expression list or one of ['!', '&', '(', '*', '+', '-', ';', '[', '^', '}', <-, case, chan, default, func, identifier, interface, literal, map, struct]",
		yyXError{328, -1}:  "expected optional expression or one of ['!', '&', '(', '*', '+', '-', ':', '[', ']', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{331, -1}:  "expected optional expression or one of ['!', '&', '(', '*', '+', '-', '[', ']', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{424, -1}:  "expected optional function body or one of [';', '{']",
		yyXError{425, -1}:  "expected optional function body or one of [';', '{']",
		yyXError{121, -1}:  "expected optional generic arguments or one of [!=, %=, &&, &=, &^, &^=, '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '[', ']', '^', '{', '|', '}', *=, ++, +=, --, -=, ..., /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, {, |=, ||, «]",
		yyXError{53, -1}:   "expected optional generic arguments or one of [!=, %=, &&, &=, &^, &^=, '%', '&', '(', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '[', '^', '{', '|', '}', *=, ++, +=, --, -=, /=, :=, <-, <<, <<=, <=, ==, >=, >>, >>=, ^=, case, default, |=, ||, «]",
		yyXError{113, -1}:  "expected optional generic arguments or one of ['(', ')', ',', ':', ';', '=', ']', '{', '}', ..., :=, case, default, literal, {, «, »]",
		yyXError{216, -1}:  "expected optional generic parameters or type or one of ['(', '*', '[', <-, chan, func, identifier, interface, map, struct, «]",
		yyXError{60, -1}:   "expected optional identifier or one of [';', '}', case, default, identifier]",
		yyXError{61, -1}:   "expected optional identifier or one of [';', '}', case, default, identifier]",
		yyXError{181, -1}:  "expected optional result or one of ['(', ')', '*', ',', ':', ';', '=', '[', ']', '{', '}', ..., :=, <-, case, chan, default, func, identifier, interface, literal, map, struct, {, »]",
		yyXError{421, -1}:  "expected optional result or one of ['(', '*', '[', '{', <-, chan, func, identifier, interface, map, struct, {]",
		yyXError{24, -1}:   "expected optional semicolon or one of [')', ';']",
		yyXError{97, -1}:   "expected optional semicolon or one of [')', ';']",
		yyXError{226, -1}:  "expected optional semicolon or one of [')', ';']",
		yyXError{480, -1}:  "expected optional semicolon or one of [')', ';']",
		yyXError{290, -1}:  "expected optional semicolon or one of [';', '}']",
		yyXError{396, -1}:  "expected optional semicolon or one of [';', '}']",
		yyXError{435, -1}:  "expected optional simple statement or one of ['!', '&', '(', '*', '+', '-', ';', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{277, -1}:  "expected optional simple statement or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct, {]",
		yyXError{437, -1}:  "expected optional simple statement or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct, {]",
		yyXError{287, -1}:  "expected optional string literal or one of [';', '}', literal]",
		yyXError{298, -1}:  "expected optional string literal or one of [';', '}', literal]",
		yyXError{301, -1}:  "expected optional string literal or one of [';', '}', literal]",
		yyXError{304, -1}:  "expected optional string literal or one of [';', '}', literal]",
		yyXError{307, -1}:  "expected optional string literal or one of [';', '}', literal]",
		yyXError{310, -1}:  "expected optional string literal or one of [';', '}', literal]",
		yyXError{180, -1}:  "expected parameter declaration list or one of ['(', ')', '*', '[', ..., <-, chan, func, identifier, interface, map, struct]",
		yyXError{182, -1}:  "expected parameter declaration list or type or one of ['(', ')', '*', '[', ..., <-, chan, func, identifier, interface, map, struct]",
		yyXError{192, -1}:  "expected parameter declaration or one of ['(', ')', '*', '[', ..., <-, chan, func, identifier, interface, map, struct]",
		yyXError{296, -1}:  "expected qualified identifier or identifier",
		yyXError{305, -1}:  "expected qualified identifier or identifier",
		yyXError{286, -1}:  "expected qualified identifier or one of ['(', identifier]",
		yyXError{288, -1}:  "expected qualified identifier or one of ['*', identifier]",
		yyXError{0, -1}:    "expected source file or package",
		yyXError{281, -1}:  "expected struct field declaration list or one of ['(', '*', '}', identifier]",
		yyXError{283, -1}:  "expected struct field declaration list or one of ['(', '*', identifier]",
		yyXError{291, -1}:  "expected struct field declaration or one of ['(', '*', '}', identifier]",
		yyXError{236, -1}:  "expected switch case/default clause list or one of ['}', case, default]",
		yyXError{239, -1}:  "expected switch case/default clause list or one of [case, default]",
		yyXError{240, -1}:  "expected switch case/default clause statement block or one of ['}', case, default]",
		yyXError{82, -1}:   "expected type literal or unary expression or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{186, -1}:  "expected type or one of ['(', ')', '*', ',', '.', '[', ..., <-, chan, func, identifier, interface, map, struct, «]",
		yyXError{473, -1}:  "expected type or one of ['(', ')', '*', ',', ';', '=', '[', '}', <-, case, chan, default, func, identifier, interface, map, struct]",
		yyXError{93, -1}:   "expected type or one of ['(', '*', ',', '=', '[', <-, chan, func, identifier, interface, map, struct]",
		yyXError{285, -1}:  "expected type or one of ['(', '*', ',', '[', <-, chan, func, identifier, interface, map, struct]",
		yyXError{36, -1}:   "expected type or one of ['(', '*', '[', <-, chan, func, identifier, interface, map, struct]",
		yyXError{106, -1}:  "expected type or one of ['(', '*', '[', <-, chan, func, identifier, interface, map, struct]",
		yyXError{107, -1}:  "expected type or one of ['(', '*', '[', <-, chan, func, identifier, interface, map, struct]",
		yyXError{126, -1}:  "expected type or one of ['(', '*', '[', <-, chan, func, identifier, interface, map, struct]",
		yyXError{187, -1}:  "expected type or one of ['(', '*', '[', <-, chan, func, identifier, interface, map, struct]",
		yyXError{198, -1}:  "expected type or one of ['(', '*', '[', <-, chan, func, identifier, interface, map, struct]",
		yyXError{218, -1}:  "expected type or one of ['(', '*', '[', <-, chan, func, identifier, interface, map, struct]",
		yyXError{336, -1}:  "expected type or one of ['(', '*', '[', <-, chan, func, identifier, interface, map, struct]",
		yyXError{389, -1}:  "expected type or one of ['(', '*', '[', <-, chan, func, identifier, interface, map, struct]",
		yyXError{391, -1}:  "expected type or one of ['(', '*', '[', <-, chan, func, identifier, interface, map, struct]",
		yyXError{486, -1}:  "expected type or one of ['(', '*', '[', <-, chan, func, identifier, interface, map, struct]",
		yyXError{489, -1}:  "expected type or one of ['(', '*', '[', <-, chan, func, identifier, interface, map, struct]",
		yyXError{495, -1}:  "expected type or one of ['(', '*', '[', <-, chan, func, identifier, interface, map, struct]",
		yyXError{497, -1}:  "expected type or one of ['(', '*', '[', <-, chan, func, identifier, interface, map, struct]",
		yyXError{499, -1}:  "expected type or one of ['(', '*', '[', <-, chan, func, identifier, interface, map, struct]",
		yyXError{214, -1}:  "expected type specification list or one of [')', identifier]",
		yyXError{81, -1}:   "expected type specification or one of ['(', identifier]",
		yyXError{228, -1}:  "expected type specification or one of [')', identifier]",
		yyXError{85, -1}:   "expected unary expression or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{86, -1}:   "expected unary expression or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{87, -1}:   "expected unary expression or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{88, -1}:   "expected unary expression or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{89, -1}:   "expected unary expression or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{90, -1}:   "expected unary expression or one of ['!', '&', '(', '*', '+', '-', '[', '^', <-, chan, func, identifier, interface, literal, map, struct]",
		yyXError{94, -1}:   "expected variable specification list or one of [')', identifier]",
		yyXError{91, -1}:   "expected variable specification or one of ['(', identifier]",
		yyXError{99, -1}:   "expected variable specification or one of [')', identifier]",
		yyXError{278, -1}:  "expected {",
		yyXError{429, -1}:  "expected {",
		yyXError{438, -1}:  "expected {",
		yyXError{103, 100}: "anonymous functions cannot have generic type parameters",
	}

	yyParseTab = [501][]uint16{
		// 0
		{161: 292, 172: 294, 295, 293},
		{97: 291},
		{2: 42, 42, 42, 42, 42, 42, 42, 12: 42, 14: 42, 42, 17: 42, 42, 42, 42, 42, 42, 32: 42, 35: 42, 42, 42, 42, 79: 42, 81: 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 97: 42, 180: 324},
		{12: 138, 151: 321},
		{2: 167, 167, 167, 167, 167, 167, 167, 12: 167, 14: 167, 167, 17: 167, 167, 167, 167, 167, 167, 32: 167, 35: 167, 167, 167, 167, 79: 167, 81: 167, 167, 167, 167, 167, 167, 167, 167, 167, 167, 167, 167, 167, 167, 97: 167, 120: 167, 168: 296},
		// 5
		{2: 117, 117, 117, 117, 117, 117, 117, 12: 117, 14: 117, 117, 17: 117, 117, 117, 117, 117, 117, 32: 117, 35: 117, 117, 117, 117, 79: 117, 81: 117, 117, 117, 117, 117, 117, 117, 117, 117, 117, 117, 117, 117, 117, 97: 117, 120: 297, 167: 298},
		{2: 301, 12: 300, 15: 183, 32: 183, 35: 183, 183, 183, 75: 303, 118: 304, 129: 302},
		{299},
		{2: 166, 166, 166, 166, 166, 166, 166, 12: 166, 14: 166, 166, 17: 166, 166, 166, 166, 166, 166, 32: 166, 35: 166, 166, 166, 166, 79: 166, 81: 166, 166, 166, 166, 166, 166, 166, 166, 166, 166, 166, 166, 166, 166, 97: 166, 120: 166},
		{182, 182, 9: 182, 182, 15: 182, 32: 182, 35: 182, 182, 182},
		// 10
		{11: 314, 300, 15: 183, 32: 183, 35: 183, 183, 183, 75: 303, 118: 304, 129: 316, 169: 315},
		{174},
		{15: 309, 32: 305, 35: 306, 307, 308, 41: 312},
		{15: 309, 32: 305, 35: 306, 307, 308, 41: 310},
		{271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 13: 271, 271, 16: 271, 23: 271, 271, 33: 271, 271, 39: 271, 271, 47: 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 62: 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 271, 79: 271},
		// 15
		{270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 13: 270, 270, 16: 270, 23: 270, 270, 33: 270, 270, 39: 270, 270, 47: 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 62: 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 270, 79: 270},
		{269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 13: 269, 269, 16: 269, 23: 269, 269, 33: 269, 269, 39: 269, 269, 47: 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 62: 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 269, 79: 269},
		{268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 13: 268, 268, 16: 268, 23: 268, 268, 33: 268, 268, 39: 268, 268, 47: 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 62: 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 79: 268},
		{267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 13: 267, 267, 16: 267, 23: 267, 267, 33: 267, 267, 39: 267, 267, 47: 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 62: 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 267, 79: 267},
		{172, 11: 172, 79: 311},
		// 20
		{170, 11: 170},
		{173, 11: 173, 79: 313},
		{171, 11: 171},
		{176},
		{318, 11: 107, 113: 317},
		// 25
		{169, 11: 169},
		{11: 320},
		{11: 106, 300, 15: 183, 32: 183, 35: 183, 183, 183, 75: 303, 118: 304, 129: 319},
		{168, 11: 168},
		{175},
		// 30
		{12: 322},
		{323},
		{2: 137, 137, 137, 137, 137, 137, 137, 12: 137, 14: 137, 137, 17: 137, 137, 137, 137, 137, 137, 32: 137, 35: 137, 137, 137, 137, 79: 137, 81: 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 97: 137, 120: 137},
		{2: 341, 373, 377, 378, 379, 380, 381, 12: 344, 14: 325, 309, 17: 327, 337, 339, 340, 328, 363, 25: 329, 374, 343, 375, 330, 331, 332, 305, 35: 306, 307, 308, 376, 41: 342, 346, 345, 347, 348, 334, 61: 335, 77: 326, 79: 370, 81: 338, 372, 351, 333, 352, 353, 354, 336, 356, 357, 359, 349, 364, 382, 350, 361, 290, 101: 365, 355, 358, 105: 360, 107: 369, 362, 367, 368, 164: 366, 179: 371},
		{2: 341, 373, 377, 378, 379, 380, 381, 12: 412, 14: 325, 309, 17: 327, 394, 339, 340, 328, 363, 25: 329, 374, 343, 375, 330, 331, 332, 305, 34: 784, 306, 307, 308, 376, 786, 41: 342, 346, 345, 347, 348, 334, 61: 785},
		// 35
		{13: 420, 24: 782, 33: 783, 62: 731, 733, 732, 734, 735, 736, 737, 738, 739, 740, 741},
		{2: 397, 398, 12: 396, 14: 325, 17: 327, 394, 339, 340, 328, 363, 25: 399, 400, 401, 402, 403, 405, 406, 78: 404, 80: 779, 181: 780},
		{17: 777},
		{2: 26, 11: 26, 13: 26, 16: 246, 23: 26, 26, 33: 26, 26, 40: 246},
		{2: 22, 11: 22, 13: 22, 16: 245, 23: 22, 22, 33: 22, 22, 40: 245},
		// 40
		{2: 21, 11: 21, 13: 21, 16: 244, 23: 21, 21, 33: 21, 21, 40: 244},
		{2: 20, 11: 20, 13: 20, 16: 243, 23: 20, 20, 33: 20, 20, 40: 243},
		{2: 762, 12: 383, 99: 764, 125: 763},
		{227, 227, 3: 227, 227, 227, 227, 227, 227, 227, 227, 227, 13: 227, 16: 227, 23: 227, 227, 33: 227, 227, 39: 227, 47: 227, 227, 227, 227, 227, 227, 227, 227, 227, 227, 227, 227, 227, 227, 62: 227, 227, 227, 227, 227, 227, 227, 227, 227, 227, 227, 227, 227},
		{103, 103, 3: 424, 423, 425, 426, 430, 441, 103, 103, 13: 204, 16: 103, 24: 204, 33: 204, 47: 422, 427, 428, 429, 431, 432, 433, 434, 435, 436, 437, 438, 439, 440, 62: 204, 204, 204, 204, 204, 204, 204, 204, 204, 204, 204, 760, 761},
		// 45
		{99, 2: 341, 373, 377, 378, 379, 380, 381, 12: 412, 14: 325, 309, 99, 327, 394, 339, 340, 328, 363, 25: 329, 374, 343, 375, 330, 331, 332, 305, 35: 306, 307, 308, 376, 41: 342, 346, 345, 347, 348, 334, 61: 335, 77: 719, 95: 350, 525, 106: 721, 131: 723, 162: 722, 175: 720},
		{2: 471, 12: 113, 112: 712, 122: 470, 176: 711},
		{99, 2: 341, 373, 377, 378, 379, 380, 381, 12: 412, 14: 325, 309, 99, 327, 394, 339, 340, 328, 363, 25: 329, 374, 343, 375, 330, 331, 332, 305, 35: 306, 307, 308, 376, 41: 342, 346, 345, 347, 348, 334, 61: 335, 77: 326, 95: 350, 525, 106: 524, 128: 697},
		{16: 570, 40: 571, 104: 684},
		{14: 680},
		// 50
		{2: 341, 373, 377, 378, 379, 380, 381, 12: 412, 14: 325, 309, 17: 327, 394, 339, 340, 328, 363, 25: 329, 374, 343, 375, 330, 331, 332, 305, 35: 306, 307, 308, 376, 41: 342, 346, 345, 347, 677, 334, 61: 676},
		{144, 144, 144, 144, 144, 144, 144, 144, 144, 144, 144, 144, 13: 144, 144, 16: 144, 23: 144, 144, 33: 144, 144, 39: 144, 144, 47: 144, 144, 144, 144, 144, 144, 144, 144, 144, 144, 144, 144, 144, 144, 62: 144, 144, 144, 144, 144, 144, 144, 144, 144, 144, 144, 144, 144, 144},
		{2: 24, 11: 24, 13: 24, 16: 143, 23: 24, 24, 33: 24, 24, 40: 143, 150: 672},
		{189, 189, 189, 189, 189, 189, 189, 189, 189, 189, 189, 13: 189, 189, 23: 670, 189, 33: 189, 40: 189, 47: 189, 189, 189, 189, 189, 189, 189, 189, 189, 189, 189, 189, 189, 189, 62: 189, 189, 189, 189, 189, 189, 189, 189, 189, 189, 189, 189, 189, 189, 100: 413, 127: 414},
		{128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 13: 128, 128, 16: 128, 23: 128, 128, 33: 128, 128, 39: 128, 128, 47: 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 62: 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128},
		// 55
		{16: 570, 40: 571, 104: 656, 121: 657},
		{9, 9, 611, 9, 9, 9, 9, 9, 9, 9, 9, 9, 13: 9, 614, 16: 9, 23: 9, 9, 33: 9, 9, 39: 9, 612, 47: 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 62: 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 613, 116: 616, 154: 615},
		{2: 500},
		{16: 527, 142: 610},
		{104, 104, 9: 104, 104, 16: 104},
		// 60
		{183, 183, 9: 183, 183, 12: 300, 118: 609},
		{183, 183, 9: 183, 183, 12: 300, 118: 608},
		{2: 341, 373, 377, 378, 379, 380, 381, 12: 412, 14: 325, 309, 17: 327, 394, 339, 340, 328, 363, 25: 329, 374, 343, 375, 330, 331, 332, 305, 35: 306, 307, 308, 376, 41: 342, 346, 345, 347, 348, 334, 61: 607},
		{84, 84, 9: 84, 84},
		{83, 83, 9: 83, 83},
		// 65
		{2: 341, 373, 377, 378, 379, 380, 381, 12: 412, 14: 325, 309, 17: 327, 394, 339, 340, 328, 363, 25: 329, 374, 343, 375, 330, 331, 332, 305, 35: 306, 307, 308, 376, 41: 342, 346, 345, 347, 348, 334, 61: 606},
		{12: 605},
		{79, 79, 9: 79, 79},
		{202, 202, 341, 373, 377, 378, 379, 380, 381, 202, 202, 12: 412, 14: 325, 309, 17: 327, 394, 339, 340, 328, 363, 25: 329, 374, 343, 375, 330, 331, 332, 305, 35: 306, 307, 308, 376, 41: 342, 346, 345, 347, 348, 334, 61: 410, 77: 603, 160: 604},
		{77, 77, 9: 77, 77},
		// 70
		{76, 76, 9: 76, 76},
		{75, 75, 9: 75, 75},
		{16: 570, 40: 571, 104: 572},
		{99, 2: 341, 373, 377, 378, 379, 380, 381, 12: 412, 14: 325, 309, 99, 327, 394, 339, 340, 328, 363, 25: 329, 374, 343, 375, 330, 331, 332, 305, 35: 306, 307, 308, 376, 41: 342, 346, 345, 347, 348, 334, 61: 335, 77: 326, 95: 350, 525, 106: 524, 128: 526},
		{48},
		// 75
		{47},
		{46},
		{45},
		{44},
		{43},
		// 80
		{523},
		{2: 505, 12: 507, 132: 506},
		{2: 341, 373, 377, 378, 379, 380, 381, 12: 412, 14: 325, 309, 17: 327, 394, 339, 340, 328, 363, 25: 329, 374, 343, 375, 330, 331, 332, 305, 35: 306, 307, 308, 376, 41: 342, 346, 345, 347, 498, 499},
		{2: 25, 11: 25, 13: 25, 23: 25, 25, 33: 25, 25},
		{2: 23, 11: 23, 13: 23, 23: 23, 23, 33: 23, 23},
		// 85
		{2: 341, 373, 377, 378, 379, 380, 381, 12: 412, 14: 325, 309, 17: 327, 394, 339, 340, 328, 363, 25: 329, 374, 343, 375, 330, 331, 332, 305, 35: 306, 307, 308, 376, 41: 342, 346, 345, 347, 348, 497},
		{2: 341, 373, 377, 378, 379, 380, 381, 12: 412, 14: 325, 309, 17: 327, 394, 339, 340, 328, 363, 25: 329, 374, 343, 375, 330, 331, 332, 305, 35: 306, 307, 308, 376, 41: 342, 346, 345, 347, 348, 496},
		{2: 341, 373, 377, 378, 379, 380, 381, 12: 412, 14: 325, 309, 17: 327, 394, 339, 340, 328, 363, 25: 329, 374, 343, 375, 330, 331, 332, 305, 35: 306, 307, 308, 376, 41: 342, 346, 345, 347, 348, 495},
		{2: 341, 373, 377, 378, 379, 380, 381, 12: 412, 14: 325, 309, 17: 327, 394, 339, 340, 328, 363, 25: 329, 374, 343, 375, 330, 331, 332, 305, 35: 306, 307, 308, 376, 41: 342, 346, 345, 347, 348, 494},
		{2: 341, 373, 377, 378, 379, 380, 381, 12: 412, 14: 325, 309, 17: 327, 394, 339, 340, 328, 363, 25: 329, 374, 343, 375, 330, 331, 332, 305, 35: 306, 307, 308, 376, 41: 342, 346, 345, 347, 348, 493},
		// 90
		{2: 341, 373, 377, 378, 379, 380, 381, 12: 412, 14: 325, 309, 17: 327, 394, 339, 340, 328, 363, 25: 329, 374, 343, 375, 330, 331, 332, 305, 35: 306, 307, 308, 376, 41: 342, 346, 345, 347, 348, 492},
		{2: 385, 12: 383, 99: 384, 133: 386},
		{181, 181, 181, 181, 9: 181, 181, 181, 181, 181, 181, 17: 181, 181, 181, 181, 181, 181, 24: 181},
		{2: 397, 398, 12: 396, 395, 325, 17: 327, 394, 339, 340, 328, 363, 24: 407, 399, 400, 401, 402, 403, 405, 406, 78: 404, 80: 408},
		{11: 387, 383, 99: 384, 133: 389, 183: 388},
		// 95
		{6, 6, 9: 6, 6},
		{8, 8, 9: 8, 8},
		{390, 11: 107, 113: 391},
		{2, 11: 2},
		{11: 106, 383, 99: 384, 133: 393},
		// 100
		{11: 392},
		{7, 7, 9: 7, 7},
		{1, 11: 1},
		{2: 471, 112: 472, 122: 470},
		{12: 469},
		// 105
		{140, 140, 140, 9: 140, 140, 140, 13: 140, 15: 140, 140, 23: 140, 140, 33: 140, 140, 39: 140, 140, 75: 467, 140, 100: 140},
		{2: 397, 398, 12: 396, 14: 325, 17: 327, 394, 339, 340, 328, 363, 25: 399, 400, 401, 402, 403, 405, 406, 78: 404, 80: 465},
		{2: 397, 398, 12: 396, 14: 325, 17: 327, 394, 339, 340, 328, 363, 25: 399, 400, 401, 402, 403, 405, 406, 78: 404, 80: 464},
		{38, 38, 38, 9: 38, 38, 38, 13: 38, 15: 38, 38, 23: 38, 38, 33: 38, 38, 39: 38, 38, 76: 38},
		{37, 37, 37, 9: 37, 37, 37, 13: 37, 15: 37, 37, 23: 37, 37, 33: 37, 37, 39: 37, 37, 76: 37},
		// 110
		{36, 36, 36, 9: 36, 36, 36, 13: 36, 15: 36, 36, 23: 36, 36, 33: 36, 36, 39: 36, 36, 76: 36},
		{35, 35, 35, 9: 35, 35, 35, 13: 35, 15: 35, 35, 23: 35, 35, 33: 35, 35, 39: 35, 35, 76: 35},
		{34, 34, 34, 9: 34, 34, 34, 13: 34, 15: 34, 34, 23: 34, 34, 33: 34, 34, 39: 34, 34, 76: 34},
		{189, 189, 189, 9: 189, 189, 189, 13: 189, 15: 189, 189, 23: 189, 189, 33: 189, 189, 39: 189, 189, 76: 189, 100: 413, 127: 463},
		{32, 32, 32, 9: 32, 32, 32, 13: 32, 15: 32, 32, 23: 32, 32, 33: 32, 32, 39: 32, 32, 76: 32},
		// 115
		{31, 31, 31, 9: 31, 31, 31, 13: 31, 15: 31, 31, 23: 31, 31, 33: 31, 31, 39: 31, 31, 76: 31},
		{2: 341, 373, 377, 378, 379, 380, 381, 12: 412, 14: 325, 309, 17: 327, 394, 339, 340, 328, 363, 25: 329, 374, 343, 375, 330, 331, 332, 305, 35: 306, 307, 308, 376, 41: 342, 346, 345, 347, 348, 334, 61: 410, 77: 462},
		{4, 4, 9: 4, 4, 4, 24: 409},
		{2: 341, 373, 377, 378, 379, 380, 381, 12: 412, 14: 325, 309, 17: 327, 394, 339, 340, 328, 363, 25: 329, 374, 343, 375, 330, 331, 332, 305, 35: 306, 307, 308, 376, 41: 342, 346, 345, 347, 348, 334, 61: 410, 77: 411},
		{204, 204, 3: 424, 423, 425, 426, 430, 441, 204, 204, 204, 13: 204, 16: 204, 47: 422, 427, 428, 429, 431, 432, 433, 434, 435, 436, 437, 438, 439, 440},
		// 120
		{3, 3, 9: 3, 3, 3, 13: 420},
		{189, 189, 189, 189, 189, 189, 189, 189, 189, 189, 189, 189, 13: 189, 189, 16: 189, 23: 189, 189, 33: 189, 189, 39: 189, 189, 47: 189, 189, 189, 189, 189, 189, 189, 189, 189, 189, 189, 189, 189, 189, 62: 189, 189, 189, 189, 189, 189, 189, 189, 189, 189, 189, 189, 189, 189, 100: 413, 127: 414},
		{2: 397, 398, 12: 396, 14: 325, 17: 327, 394, 339, 340, 328, 363, 25: 399, 400, 401, 402, 403, 405, 406, 78: 404, 80: 415, 165: 416},
		{141, 141, 141, 141, 141, 141, 141, 141, 141, 141, 141, 141, 13: 141, 141, 16: 141, 23: 141, 141, 33: 141, 141, 39: 141, 141, 47: 141, 141, 141, 141, 141, 141, 141, 141, 141, 141, 141, 141, 141, 141, 62: 141, 141, 141, 141, 141, 141, 141, 141, 141, 141, 141, 141, 141, 141},
		{13: 191, 76: 191},
		// 125
		{13: 417, 76: 418},
		{2: 397, 398, 12: 396, 14: 325, 17: 327, 394, 339, 340, 328, 363, 25: 399, 400, 401, 402, 403, 405, 406, 78: 404, 80: 419},
		{188, 188, 188, 188, 188, 188, 188, 188, 188, 188, 188, 188, 13: 188, 188, 188, 188, 23: 188, 188, 33: 188, 188, 39: 188, 188, 47: 188, 188, 188, 188, 188, 188, 188, 188, 188, 188, 188, 188, 188, 188, 62: 188, 188, 188, 188, 188, 188, 188, 188, 188, 188, 188, 188, 188, 188, 188},
		{13: 190, 76: 190},
		{2: 341, 373, 377, 378, 379, 380, 381, 12: 412, 14: 325, 309, 17: 327, 394, 339, 340, 328, 363, 25: 329, 374, 343, 375, 330, 331, 332, 305, 35: 306, 307, 308, 376, 41: 342, 346, 345, 347, 348, 334, 61: 421},
		// 130
		{203, 203, 3: 424, 423, 425, 426, 430, 441, 203, 203, 203, 13: 203, 16: 203, 24: 203, 33: 203, 47: 422, 427, 428, 429, 431, 432, 433, 434, 435, 436, 437, 438, 439, 440, 62: 203, 203, 203, 203, 203, 203, 203, 203, 203, 203, 203},
		{2: 341, 373, 377, 378, 379, 380, 381, 12: 412, 14: 325, 309, 17: 327, 394, 339, 340, 328, 363, 25: 329, 374, 343, 375, 330, 331, 332, 305, 35: 306, 307, 308, 376, 41: 342, 346, 345, 347, 348, 334, 61: 461},
		{2: 341, 373, 377, 378, 379, 380, 381, 12: 412, 14: 325, 309, 17: 327, 394, 339, 340, 328, 363, 25: 329, 374, 343, 375, 330, 331, 332, 305, 35: 306, 307, 308, 376, 41: 342, 346, 345, 347, 348, 334, 61: 460},
		{2: 341, 373, 377, 378, 379, 380, 381, 12: 412, 14: 325, 309, 17: 327, 394, 339, 340, 328, 363, 25: 329, 374, 343, 375, 330, 331, 332, 305, 35: 306, 307, 308, 376, 41: 342, 346, 345, 347, 348, 334, 61: 459},
		{2: 341, 373, 377, 378, 379, 380, 381, 12: 412, 14: 325, 309, 17: 327, 394, 339, 340, 328, 363, 25: 329, 374, 343, 375, 330, 331, 332, 305, 35: 306, 307, 308, 376, 41: 342, 346, 345, 347, 348, 334, 61: 458},
		// 135
		{2: 341, 373, 377, 378, 379, 380, 381, 12: 412, 14: 325, 309, 17: 327, 394, 339, 340, 328, 363, 25: 329, 374, 343, 375, 330, 331, 332, 305, 35: 306, 307, 308, 376, 41: 342, 346, 345, 347, 348, 334, 61: 457},
		{2: 341, 373, 377, 378, 379, 380, 381, 12: 412, 14: 325, 309, 17: 327, 394, 339, 340, 328, 363, 25: 329, 374, 343, 375, 330, 331, 332, 305, 35: 306, 307, 308, 376, 41: 342, 346, 345, 347, 348, 334, 61: 456},
		{2: 341, 373, 377, 378, 379, 380, 381, 12: 412, 14: 325, 309, 17: 327, 394, 339, 340, 328, 363, 25: 329, 374, 343, 375, 330, 331, 332, 305, 35: 306, 307, 308, 376, 41: 342, 346, 345, 347, 348, 334, 61: 455},
		{2: 341, 373, 377, 378, 379, 380, 381, 12: 412, 14: 325, 309, 17: 327, 394, 339, 340, 328, 363, 25: 329, 374, 343, 375, 330, 331, 332, 305, 35: 306, 307, 308, 376, 41: 342, 346, 345, 347, 348, 334, 61: 454},
		{2: 341, 373, 377, 378, 379, 380, 381, 12: 412, 14: 325, 309, 17: 327, 394, 339, 340, 328, 363, 25: 329, 374, 343, 375, 330, 331, 332, 305, 35: 306, 307, 308, 376, 41: 342, 346, 345, 347, 348, 334, 61: 453},
		// 140
		{2: 341, 373, 377, 378, 379, 380, 381, 12: 412, 14: 325, 309, 17: 327, 394, 339, 340, 328, 363, 25: 329, 374, 343, 375, 330, 331, 332, 305, 35: 306, 307, 308, 376, 41: 342, 346, 345, 347, 348, 334, 61: 452},
		{2: 341, 373, 377, 378, 379, 380, 381, 12: 412, 14: 325, 309, 17: 327, 394, 339, 340, 328, 363, 25: 329, 374, 343, 375, 330, 331, 332, 305, 35: 306, 307, 308, 376, 41: 342, 346, 345, 347, 348, 334, 61: 451},
		{2: 341, 373, 377, 378, 379, 380, 381, 12: 412, 14: 325, 309, 17: 327, 394, 339, 340, 328, 363, 25: 329, 374, 343, 375, 330, 331, 332, 305, 35: 306, 307, 308, 376, 41: 342, 346, 345, 347, 348, 334, 61: 450},
		{2: 341, 373, 377, 378, 379, 380, 381, 12: 412, 14: 325, 309, 17: 327, 394, 339, 340, 328, 363, 25: 329, 374, 343, 375, 330, 331, 332, 305, 35: 306, 307, 308, 376, 41: 342, 346, 345, 347, 348, 334, 61: 449},
		{2: 341, 373, 377, 378, 379, 380, 381, 12: 412, 14: 325, 309, 17: 327, 394, 339, 340, 328, 363, 25: 329, 374, 343, 375, 330, 331, 332, 305, 35: 306, 307, 308, 376, 41: 342, 346, 345, 347, 348, 334, 61: 448},
		// 145
		{2: 341, 373, 377, 378, 379, 380, 381, 12: 412, 14: 325, 309, 17: 327, 394, 339, 340, 328, 363, 25: 329, 374, 343, 375, 330, 331, 332, 305, 35: 306, 307, 308, 376, 41: 342, 346, 345, 347, 348, 334, 61: 447},
		{2: 341, 373, 377, 378, 379, 380, 381, 12: 412, 14: 325, 309, 17: 327, 394, 339, 340, 328, 363, 25: 329, 374, 343, 375, 330, 331, 332, 305, 35: 306, 307, 308, 376, 41: 342, 346, 345, 347, 348, 334, 61: 446},
		{2: 341, 373, 377, 378, 379, 380, 381, 12: 412, 14: 325, 309, 17: 327, 394, 339, 340, 328, 363, 25: 329, 374, 343, 375, 330, 331, 332, 305, 35: 306, 307, 308, 376, 41: 342, 346, 345, 347, 348, 334, 61: 445},
		{2: 341, 373, 377, 378, 379, 380, 381, 12: 412, 14: 325, 309, 17: 327, 394, 339, 340, 328, 363, 25: 329, 374, 343, 375, 330, 331, 332, 305, 35: 306, 307, 308, 376, 41: 342, 346, 345, 347, 348, 334, 61: 444},
		{2: 341, 373, 377, 378, 379, 380, 381, 12: 412, 14: 325, 309, 17: 327, 394, 339, 340, 328, 363, 25: 329, 374, 343, 375, 330, 331, 332, 305, 35: 306, 307, 308, 376, 41: 342, 346, 345, 347, 348, 334, 61: 443},
		// 150
		{2: 341, 373, 377, 378, 379, 380, 381, 12: 412, 14: 325, 309, 17: 327, 394, 339, 340, 328, 363, 25: 329, 374, 343, 375, 330, 331, 332, 305, 35: 306, 307, 308, 376, 41: 342, 346, 345, 347, 348, 334, 61: 442},
		{207, 207, 3: 424, 423, 425, 426, 430, 207, 207, 207, 207, 13: 207, 16: 207, 23: 207, 207, 33: 207, 207, 39: 207, 47: 422, 427, 428, 429, 431, 432, 433, 434, 435, 436, 437, 438, 439, 440, 62: 207, 207, 207, 207, 207, 207, 207, 207, 207, 207, 207, 207, 207},
		{208, 208, 3: 208, 208, 208, 208, 208, 208, 208, 208, 208, 13: 208, 16: 208, 23: 208, 208, 33: 208, 208, 39: 208, 47: 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 62: 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208, 208},
		{209, 209, 3: 424, 423, 425, 426, 430, 209, 209, 209, 209, 13: 209, 16: 209, 23: 209, 209, 33: 209, 209, 39: 209, 47: 422, 427, 428, 429, 431, 432, 433, 434, 435, 436, 437, 438, 209, 440, 62: 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209, 209},
		{210, 210, 3: 424, 423, 425, 426, 430, 210, 210, 210, 210, 13: 210, 16: 210, 23: 210, 210, 33: 210, 210, 39: 210, 47: 422, 427, 210, 210, 431, 210, 433, 210, 210, 210, 437, 210, 210, 440, 62: 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210, 210},
		// 155
		{211, 211, 3: 211, 211, 211, 211, 211, 211, 211, 211, 211, 13: 211, 16: 211, 23: 211, 211, 33: 211, 211, 39: 211, 47: 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 62: 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211, 211},
		{212, 212, 3: 424, 423, 425, 426, 430, 212, 212, 212, 212, 13: 212, 16: 212, 23: 212, 212, 33: 212, 212, 39: 212, 47: 422, 427, 212, 212, 431, 212, 433, 212, 212, 212, 437, 212, 212, 440, 62: 212, 212, 212, 212, 212, 212, 212, 212, 212, 212, 212, 212, 212},
		{213, 213, 3: 424, 423, 425, 426, 430, 213, 213, 213, 213, 13: 213, 16: 213, 23: 213, 213, 33: 213, 213, 39: 213, 47: 422, 427, 213, 213, 431, 213, 433, 213, 213, 213, 437, 213, 213, 440, 62: 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213, 213},
		{214, 214, 3: 424, 423, 425, 426, 430, 214, 214, 214, 214, 13: 214, 16: 214, 23: 214, 214, 33: 214, 214, 39: 214, 47: 422, 427, 214, 214, 431, 214, 433, 214, 214, 214, 437, 214, 214, 440, 62: 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214, 214},
		{215, 215, 3: 215, 215, 215, 215, 215, 215, 215, 215, 215, 13: 215, 16: 215, 23: 215, 215, 33: 215, 215, 39: 215, 47: 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 62: 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215, 215},
		// 160
		{216, 216, 3: 424, 423, 425, 426, 430, 216, 216, 216, 216, 13: 216, 16: 216, 23: 216, 216, 33: 216, 216, 39: 216, 47: 422, 427, 428, 429, 431, 216, 433, 434, 435, 436, 437, 438, 216, 440, 62: 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216, 216},
		{217, 217, 3: 424, 423, 217, 217, 217, 217, 217, 217, 217, 13: 217, 16: 217, 23: 217, 217, 33: 217, 217, 39: 217, 47: 422, 427, 217, 217, 217, 217, 433, 217, 217, 217, 437, 217, 217, 440, 62: 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 217, 217},
		{218, 218, 3: 424, 423, 218, 218, 218, 218, 218, 218, 218, 13: 218, 16: 218, 23: 218, 218, 33: 218, 218, 39: 218, 47: 422, 427, 218, 218, 218, 218, 433, 218, 218, 218, 437, 218, 218, 440, 62: 218, 218, 218, 218, 218, 218, 218, 218, 218, 218, 218, 218, 218},
		{219, 219, 3: 424, 423, 425, 426, 430, 219, 219, 219, 219, 13: 219, 16: 219, 23: 219, 219, 33: 219, 219, 39: 219, 47: 422, 427, 219, 219, 431, 219, 433, 219, 219, 219, 437, 219, 219, 440, 62: 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219, 219},
		{220, 220, 3: 424, 423, 425, 426, 430, 220, 220, 220, 220, 13: 220, 16: 220, 23: 220, 220, 33: 220, 220, 39: 220, 47: 422, 427, 220, 220, 431, 220, 433, 220, 220, 220, 437, 220, 220, 440, 62: 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 220, 220},
		// 165
		{221, 221, 3: 221, 221, 221, 221, 221, 221, 221, 221, 221, 13: 221, 16: 221, 23: 221, 221, 33: 221, 221, 39: 221, 47: 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 62: 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221, 221},
		{222, 222, 3: 424, 423, 222, 222, 222, 222, 222, 222, 222, 13: 222, 16: 222, 23: 222, 222, 33: 222, 222, 39: 222, 47: 422, 427, 222, 222, 222, 222, 433, 222, 222, 222, 437, 222, 222, 440, 62: 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222},
		{223, 223, 3: 424, 423, 223, 223, 223, 223, 223, 223, 223, 13: 223, 16: 223, 23: 223, 223, 33: 223, 223, 39: 223, 47: 422, 427, 223, 223, 223, 223, 433, 223, 223, 223, 437, 223, 223, 440, 62: 223, 223, 223, 223, 223, 223, 223, 223, 223, 223, 223, 223, 223},
		{224, 224, 3: 224, 224, 224, 224, 224, 224, 224, 224, 224, 13: 224, 16: 224, 23: 224, 224, 33: 224, 224, 39: 224, 47: 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 62: 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224, 224},
		{225, 225, 3: 225, 225, 225, 225, 225, 225, 225, 225, 225, 13: 225, 16: 225, 23: 225, 225, 33: 225, 225, 39: 225, 47: 225, 225, 225, 225, 225, 225, 225, 225, 225, 225, 225, 225, 225, 225, 62: 225, 225, 225, 225, 225, 225, 225, 225, 225, 225, 225, 225, 225},
		// 170
		{226, 226, 3: 226, 226, 226, 226, 226, 226, 226, 226, 226, 13: 226, 16: 226, 23: 226, 226, 33: 226, 226, 39: 226, 47: 226, 226, 226, 226, 226, 226, 226, 226, 226, 226, 226, 226, 226, 226, 62: 226, 226, 226, 226, 226, 226, 226, 226, 226, 226, 226, 226, 226},
		{5, 5, 9: 5, 5, 5, 13: 420},
		{33, 33, 33, 9: 33, 33, 33, 13: 33, 15: 33, 33, 23: 33, 33, 33: 33, 33, 39: 33, 33, 76: 33},
		{39, 39, 39, 9: 39, 39, 39, 13: 39, 15: 39, 39, 23: 39, 39, 33: 39, 39, 39: 39, 39, 76: 39},
		{11: 466},
		// 175
		{40, 40, 40, 9: 40, 40, 40, 13: 40, 15: 40, 40, 23: 40, 40, 33: 40, 40, 39: 40, 40, 76: 40},
		{12: 468},
		{139, 139, 139, 9: 139, 139, 139, 13: 139, 15: 139, 139, 23: 139, 139, 33: 139, 139, 39: 139, 139, 76: 139, 100: 139},
		{180, 180, 180, 180, 9: 180, 180, 180, 180, 180, 180, 17: 180, 180, 180, 180, 180, 180, 24: 180},
		{192, 192, 192, 9: 192, 192, 192, 13: 192, 15: 192, 192, 23: 192, 192, 33: 192, 192, 39: 192, 192, 76: 192},
		// 180
		{2: 397, 398, 11: 482, 477, 14: 325, 17: 327, 394, 339, 340, 328, 363, 25: 399, 400, 401, 402, 403, 405, 406, 34: 478, 78: 404, 80: 486, 130: 480, 139: 481},
		{111, 111, 473, 398, 9: 111, 111, 111, 396, 111, 325, 111, 111, 327, 394, 339, 340, 328, 363, 111, 111, 399, 400, 401, 402, 403, 405, 406, 33: 111, 111, 39: 111, 111, 76: 111, 78: 404, 80: 475, 112: 474, 140: 476},
		{2: 397, 398, 11: 482, 477, 14: 325, 17: 327, 394, 339, 340, 328, 363, 25: 399, 400, 401, 402, 403, 405, 406, 34: 478, 78: 404, 80: 479, 130: 480, 139: 481},
		{110, 110, 110, 9: 110, 110, 110, 13: 110, 15: 110, 110, 23: 110, 110, 33: 110, 110, 39: 110, 110, 76: 110},
		{109, 109, 109, 9: 109, 109, 109, 13: 109, 15: 109, 109, 23: 109, 109, 33: 109, 109, 39: 109, 109, 76: 109},
		// 185
		{105, 105, 105, 9: 105, 105, 105, 13: 105, 15: 105, 105, 23: 105, 105, 33: 105, 105, 39: 105, 105, 76: 105},
		{2: 397, 398, 11: 140, 396, 140, 325, 17: 327, 394, 339, 340, 328, 363, 25: 399, 400, 401, 402, 403, 405, 406, 34: 489, 75: 467, 78: 404, 80: 490, 100: 140},
		{2: 397, 398, 12: 396, 14: 325, 17: 327, 394, 339, 340, 328, 363, 25: 399, 400, 401, 402, 403, 405, 406, 78: 404, 80: 488},
		{11: 466, 13: 133},
		{11: 132, 13: 132},
		// 190
		{11: 256, 13: 483, 111: 484},
		{130, 130, 130, 130, 9: 130, 130, 130, 130, 130, 130, 130, 130, 130, 130, 130, 130, 130, 130, 130, 130, 33: 130, 130, 39: 130, 130, 76: 130},
		{2: 397, 398, 11: 255, 477, 14: 325, 17: 327, 394, 339, 340, 328, 363, 25: 399, 400, 401, 402, 403, 405, 406, 34: 478, 78: 404, 80: 486, 130: 487},
		{11: 485},
		{129, 129, 129, 129, 9: 129, 129, 129, 129, 129, 129, 129, 129, 129, 129, 129, 129, 129, 129, 129, 129, 33: 129, 129, 39: 129, 129, 76: 129},
		// 195
		{11: 133, 13: 133},
		{11: 131, 13: 131},
		{11: 136, 13: 136},
		{2: 397, 398, 12: 396, 14: 325, 17: 327, 394, 339, 340, 328, 363, 25: 399, 400, 401, 402, 403, 405, 406, 78: 404, 80: 491},
		{11: 134, 13: 134},
		// 200
		{11: 135, 13: 135},
		{10, 10, 3: 10, 10, 10, 10, 10, 10, 10, 10, 10, 13: 10, 16: 10, 23: 10, 10, 33: 10, 10, 39: 10, 47: 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 62: 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10},
		{11, 11, 3: 11, 11, 11, 11, 11, 11, 11, 11, 11, 13: 11, 16: 11, 23: 11, 11, 33: 11, 11, 39: 11, 47: 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 62: 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11},
		{12, 12, 3: 12, 12, 12, 12, 12, 12, 12, 12, 12, 13: 12, 16: 12, 23: 12, 12, 33: 12, 12, 39: 12, 47: 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 62: 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12},
		{13, 13, 3: 13, 13, 13, 13, 13, 13, 13, 13, 13, 13: 13, 16: 13, 23: 13, 13, 33: 13, 13, 39: 13, 47: 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 62: 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13},
		// 205
		{15, 15, 3: 15, 15, 15, 15, 15, 15, 15, 15, 15, 13: 15, 16: 15, 23: 15, 15, 33: 15, 15, 39: 15, 47: 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 62: 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15},
		{16, 16, 3: 16, 16, 16, 16, 16, 16, 16, 16, 16, 13: 16, 16: 16, 23: 16, 16, 33: 16, 16, 39: 16, 47: 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 62: 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16},
		{2: 500, 11: 27, 13: 27, 23: 27, 27, 33: 27, 27},
		{14, 14, 3: 14, 14, 14, 14, 14, 14, 14, 14, 14, 13: 14, 16: 14, 23: 14, 14, 33: 14, 14, 39: 14, 47: 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 62: 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14},
		{2: 341, 373, 377, 378, 379, 380, 381, 12: 412, 14: 325, 309, 17: 327, 394, 339, 340, 328, 363, 25: 329, 374, 343, 375, 330, 331, 332, 305, 35: 306, 307, 308, 376, 41: 342, 346, 345, 347, 348, 334, 61: 501},
		// 210
		{3: 424, 423, 425, 426, 430, 441, 11: 256, 13: 502, 47: 422, 427, 428, 429, 431, 432, 433, 434, 435, 436, 437, 438, 439, 440, 111: 503},
		{11: 255},
		{11: 504},
		{118, 118, 118, 118, 118, 118, 118, 118, 118, 118, 118, 118, 13: 118, 118, 16: 118, 23: 118, 118, 33: 118, 118, 39: 118, 118, 47: 118, 118, 118, 118, 118, 118, 118, 118, 118, 118, 118, 118, 118, 118, 62: 118, 118, 118, 118, 118, 118, 118, 118, 118, 118, 118, 118, 118, 118},
		{11: 516, 507, 132: 518, 182: 517},
		// 215
		{28, 28, 9: 28, 28},
		{2: 185, 185, 12: 185, 14: 185, 17: 185, 185, 185, 185, 185, 185, 100: 508, 136: 509},
		{12: 511, 166: 512},
		{2: 397, 398, 12: 396, 14: 325, 17: 327, 394, 339, 340, 328, 363, 25: 399, 400, 401, 402, 403, 405, 406, 78: 404, 80: 510},
		{19, 19, 9: 19, 19, 19},
		// 220
		{13: 187, 76: 187},
		{13: 513, 76: 514},
		{12: 515},
		{2: 184, 184, 12: 184, 14: 184, 17: 184, 184, 184, 184, 184, 184},
		{13: 186, 76: 186},
		// 225
		{30, 30, 9: 30, 30},
		{519, 11: 107, 113: 520},
		{18, 11: 18},
		{11: 106, 507, 132: 522},
		{11: 521},
		// 230
		{29, 29, 9: 29, 29},
		{17, 11: 17},
		{2: 41, 41, 41, 41, 41, 41, 41, 12: 41, 14: 41, 41, 17: 41, 41, 41, 41, 41, 41, 32: 41, 35: 41, 41, 41, 41, 79: 41, 81: 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 97: 41},
		{568, 16: 179},
		{98, 16: 98},
		// 235
		{16: 527, 142: 528},
		{1: 529, 9: 60, 60, 153: 530},
		{49, 49, 9: 49, 49},
		{61, 61, 9: 61, 61},
		{9: 532, 533, 143: 534, 535, 178: 531},
		// 240
		{1: 566, 9: 532, 533, 143: 534, 567},
		{2: 341, 373, 377, 378, 379, 380, 381, 12: 412, 14: 325, 309, 17: 327, 394, 339, 340, 328, 363, 25: 329, 374, 343, 375, 330, 331, 332, 305, 35: 306, 307, 308, 376, 41: 342, 346, 345, 347, 553, 334, 61: 552, 79: 556, 119: 554, 134: 555},
		{23: 550, 79: 551},
		{96, 96, 341, 373, 377, 378, 379, 380, 381, 96, 96, 12: 344, 14: 325, 309, 17: 327, 394, 339, 340, 328, 363, 25: 329, 374, 343, 375, 330, 331, 332, 305, 35: 306, 307, 308, 376, 40: 536, 342, 346, 345, 347, 348, 334, 61: 335, 77: 326, 79: 542, 81: 338, 372, 351, 333, 352, 353, 354, 336, 356, 357, 359, 349, 364, 382, 350, 361, 98: 537, 101: 538, 355, 358, 105: 360, 107: 541, 362, 539, 540, 114: 543, 123: 544},
		{1: 51, 9: 51, 51},
		// 245
		{266, 266, 266, 266, 266, 266, 266, 266, 266, 12: 266, 14: 266, 266, 17: 266, 266, 266, 266, 266, 266, 32: 266, 35: 266, 266, 266, 266, 40: 266, 79: 266, 81: 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 266, 145: 547},
		{95, 95, 9: 95, 95},
		{94, 94, 9: 94, 94},
		{93, 93, 9: 93, 93},
		{92, 92, 9: 92, 92},
		// 250
		{91, 91, 9: 91, 91},
		{90, 90, 9: 90, 90},
		{89, 89, 9: 89, 89},
		{545, 52, 9: 52, 52},
		{96, 96, 341, 373, 377, 378, 379, 380, 381, 96, 96, 12: 344, 14: 325, 309, 17: 327, 394, 339, 340, 328, 363, 25: 329, 374, 343, 375, 330, 331, 332, 305, 35: 306, 307, 308, 376, 40: 536, 342, 346, 345, 347, 348, 334, 61: 335, 77: 326, 79: 542, 81: 338, 372, 351, 333, 352, 353, 354, 336, 356, 357, 359, 349, 364, 382, 350, 361, 98: 537, 101: 538, 355, 358, 105: 360, 107: 541, 362, 539, 540, 114: 546},
		// 255
		{88, 88, 9: 88, 88},
		{96, 96, 341, 373, 377, 378, 379, 380, 381, 12: 344, 14: 325, 309, 17: 327, 394, 339, 340, 328, 363, 25: 329, 374, 343, 375, 330, 331, 332, 305, 35: 306, 307, 308, 376, 40: 536, 342, 346, 345, 347, 348, 334, 61: 335, 77: 326, 79: 542, 81: 338, 372, 351, 333, 352, 353, 354, 336, 356, 357, 359, 349, 364, 382, 350, 361, 98: 537, 101: 538, 355, 358, 105: 360, 107: 541, 362, 539, 540, 114: 543, 123: 548},
		{545, 549},
		{265, 265, 9: 265, 265},
		{55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 12: 55, 14: 55, 55, 17: 55, 55, 55, 55, 55, 55, 32: 55, 35: 55, 55, 55, 55, 40: 55, 79: 55, 81: 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55},
		// 260
		{53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 12: 53, 14: 53, 53, 17: 53, 53, 53, 53, 53, 53, 32: 53, 35: 53, 53, 53, 53, 40: 53, 79: 53, 81: 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53},
		{3: 424, 423, 425, 426, 430, 441, 11: 289, 13: 289, 23: 289, 289, 33: 289, 289, 47: 422, 427, 428, 429, 431, 432, 433, 434, 435, 436, 437, 438, 439, 440},
		{2: 500, 11: 288, 13: 288, 23: 288, 288, 33: 288, 288},
		{11: 287, 13: 287, 23: 287, 287, 33: 287, 287},
		{13: 557, 23: 558, 559, 33: 560},
		// 265
		{54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 12: 54, 14: 54, 54, 17: 54, 54, 54, 54, 54, 54, 32: 54, 35: 54, 54, 54, 54, 40: 54, 79: 54, 81: 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 54},
		{2: 341, 373, 377, 378, 379, 380, 381, 12: 412, 14: 325, 309, 17: 327, 394, 339, 340, 328, 363, 25: 329, 374, 343, 375, 330, 331, 332, 305, 35: 306, 307, 308, 376, 41: 342, 346, 345, 347, 553, 334, 61: 552, 119: 565},
		{58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 12: 58, 14: 58, 58, 17: 58, 58, 58, 58, 58, 58, 32: 58, 35: 58, 58, 58, 58, 40: 58, 79: 58, 81: 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58},
		{2: 341, 373, 377, 378, 379, 380, 381, 12: 412, 14: 325, 309, 17: 327, 394, 339, 340, 328, 363, 25: 329, 374, 343, 375, 330, 331, 332, 305, 35: 306, 307, 308, 376, 41: 342, 346, 345, 347, 348, 334, 61: 563},
		{2: 341, 373, 377, 378, 379, 380, 381, 12: 412, 14: 325, 309, 17: 327, 394, 339, 340, 328, 363, 25: 329, 374, 343, 375, 330, 331, 332, 305, 35: 306, 307, 308, 376, 41: 342, 346, 345, 347, 348, 334, 61: 561},
		// 270
		{3: 424, 423, 425, 426, 430, 441, 23: 562, 47: 422, 427, 428, 429, 431, 432, 433, 434, 435, 436, 437, 438, 439, 440},
		{56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 12: 56, 14: 56, 56, 17: 56, 56, 56, 56, 56, 56, 32: 56, 35: 56, 56, 56, 56, 40: 56, 79: 56, 81: 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56},
		{3: 424, 423, 425, 426, 430, 441, 23: 564, 47: 422, 427, 428, 429, 431, 432, 433, 434, 435, 436, 437, 438, 439, 440},
		{57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 12: 57, 14: 57, 57, 17: 57, 57, 57, 57, 57, 57, 32: 57, 35: 57, 57, 57, 57, 40: 57, 79: 57, 81: 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57},
		{11: 286, 13: 286, 23: 286, 286, 33: 286, 286},
		// 275
		{59, 59, 9: 59, 59},
		{1: 50, 9: 50, 50},
		{2: 341, 373, 377, 378, 379, 380, 381, 12: 412, 14: 325, 309, 99, 327, 394, 339, 340, 328, 363, 25: 329, 374, 343, 375, 330, 331, 332, 305, 35: 306, 307, 308, 376, 41: 342, 346, 345, 347, 348, 334, 61: 335, 77: 326, 95: 350, 525, 106: 569},
		{16: 178},
		{157, 157, 157, 157, 157, 157, 157, 157, 157, 12: 157, 14: 157, 157, 157, 157, 157, 157, 157, 157, 157, 32: 157, 35: 157, 157, 157, 157, 40: 157, 79: 157, 81: 157, 157, 157, 157, 157, 157, 157, 157, 157, 157, 157, 157, 157, 157},
		// 280
		{156, 156, 156, 156, 156, 156, 156, 156, 156, 12: 156, 14: 156, 156, 156, 156, 156, 156, 156, 156, 156, 32: 156, 35: 156, 156, 156, 156, 40: 156, 79: 156, 81: 156, 156, 156, 156, 156, 156, 156, 156, 156, 156, 156, 156, 156, 156},
		{1: 573, 63, 63, 12: 63, 152: 574},
		{64, 64, 64, 9: 64, 64, 64, 13: 64, 15: 64, 64, 23: 64, 64, 33: 64, 64, 39: 64, 64, 76: 64},
		{2: 579, 577, 12: 575, 78: 578, 99: 576, 141: 580, 177: 581},
		{140, 140, 181, 181, 12: 181, 181, 181, 140, 17: 181, 181, 181, 181, 181, 181, 75: 467},
		// 285
		{2: 397, 398, 12: 396, 395, 325, 17: 327, 394, 339, 340, 328, 363, 25: 399, 400, 401, 402, 403, 405, 406, 78: 404, 80: 601},
		{2: 596, 12: 396, 78: 595},
		{74, 74, 15: 590, 115: 594},
		{3: 587, 12: 396, 78: 586},
		{66, 66},
		// 290
		{582, 107, 113: 583},
		{1: 106, 579, 577, 12: 575, 78: 578, 99: 576, 141: 585},
		{1: 584},
		{62, 62, 62, 9: 62, 62, 62, 13: 62, 15: 62, 62, 23: 62, 62, 33: 62, 62, 39: 62, 62, 76: 62},
		{65, 65},
		// 295
		{11: 592},
		{12: 396, 78: 588},
		{11: 589},
		{74, 74, 15: 590, 115: 591},
		{73, 73},
		// 300
		{68, 68},
		{74, 74, 15: 590, 115: 593},
		{69, 69},
		{70, 70},
		{74, 74, 15: 590, 115: 600},
		// 305
		{12: 396, 78: 597},
		{11: 598},
		{74, 74, 15: 590, 115: 599},
		{67, 67},
		{72, 72},
		// 310
		{74, 74, 15: 590, 115: 602},
		{71, 71},
		{201, 201, 9: 201, 201, 13: 420},
		{78, 78, 9: 78, 78},
		{81, 81, 9: 81, 81},
		// 315
		{82, 82, 3: 424, 423, 425, 426, 430, 441, 82, 82, 47: 422, 427, 428, 429, 431, 432, 433, 434, 435, 436, 437, 438, 439, 440},
		{85, 85, 3: 424, 423, 425, 426, 430, 441, 85, 85, 47: 422, 427, 428, 429, 431, 432, 433, 434, 435, 436, 437, 438, 439, 440},
		{86, 86, 9: 86, 86},
		{87, 87, 9: 87, 87},
		{108, 108, 9: 108, 108},
		// 320
		{2: 341, 373, 377, 378, 379, 380, 381, 11: 649, 412, 14: 325, 309, 17: 327, 394, 339, 340, 328, 363, 25: 329, 374, 343, 375, 330, 331, 332, 305, 35: 306, 307, 308, 376, 41: 342, 346, 345, 347, 553, 334, 61: 552, 119: 554, 134: 648},
		{1: 637, 341, 373, 377, 378, 379, 380, 381, 12: 412, 14: 325, 309, 17: 327, 394, 339, 340, 328, 363, 25: 329, 374, 343, 375, 330, 331, 332, 305, 35: 306, 307, 308, 376, 40: 612, 342, 346, 345, 347, 348, 334, 61: 634, 116: 633, 135: 635, 155: 636},
		{2: 627, 12: 628},
		{2: 341, 373, 377, 378, 379, 380, 381, 12: 412, 14: 325, 309, 17: 327, 394, 339, 340, 328, 363, 206, 25: 329, 374, 343, 375, 330, 331, 332, 305, 35: 306, 307, 308, 376, 41: 342, 346, 345, 347, 348, 334, 61: 617, 126: 618},
		{120, 120, 120, 120, 120, 120, 120, 120, 120, 120, 120, 120, 13: 120, 120, 16: 120, 23: 120, 120, 33: 120, 120, 39: 120, 120, 47: 120, 120, 120, 120, 120, 120, 120, 120, 120, 120, 120, 120, 120, 120, 62: 120, 120, 120, 120, 120, 120, 120, 120, 120, 120, 120, 120, 120, 120},
		// 325
		{119, 119, 119, 119, 119, 119, 119, 119, 119, 119, 119, 119, 13: 119, 119, 16: 119, 23: 119, 119, 33: 119, 119, 39: 119, 119, 47: 119, 119, 119, 119, 119, 119, 119, 119, 119, 119, 119, 119, 119, 119, 62: 119, 119, 119, 119, 119, 119, 119, 119, 119, 119, 119, 119, 119, 119},
		{3: 424, 423, 425, 426, 430, 441, 23: 205, 39: 626, 47: 422, 427, 428, 429, 431, 432, 433, 434, 435, 436, 437, 438, 439, 440},
		{23: 619},
		{2: 341, 373, 377, 378, 379, 380, 381, 12: 412, 14: 325, 309, 17: 327, 394, 339, 340, 328, 363, 206, 25: 329, 374, 343, 375, 330, 331, 332, 305, 35: 306, 307, 308, 376, 206, 41: 342, 346, 345, 347, 348, 334, 61: 620, 126: 621},
		{3: 424, 423, 425, 426, 430, 441, 23: 205, 39: 205, 47: 422, 427, 428, 429, 431, 432, 433, 434, 435, 436, 437, 438, 439, 440},
		// 330
		{23: 622, 39: 623},
		{2: 341, 373, 377, 378, 379, 380, 381, 12: 412, 14: 325, 309, 17: 327, 394, 339, 340, 328, 363, 25: 329, 374, 343, 375, 330, 331, 332, 305, 35: 306, 307, 308, 376, 206, 41: 342, 346, 345, 347, 348, 334, 61: 620, 126: 624},
		{121, 121, 121, 121, 121, 121, 121, 121, 121, 121, 121, 121, 13: 121, 121, 16: 121, 23: 121, 121, 33: 121, 121, 39: 121, 121, 47: 121, 121, 121, 121, 121, 121, 121, 121, 121, 121, 121, 121, 121, 121, 62: 121, 121, 121, 121, 121, 121, 121, 121, 121, 121, 121, 121, 121, 121},
		{39: 625},
		{122, 122, 122, 122, 122, 122, 122, 122, 122, 122, 122, 122, 13: 122, 122, 16: 122, 23: 122, 122, 33: 122, 122, 39: 122, 122, 47: 122, 122, 122, 122, 122, 122, 122, 122, 122, 122, 122, 122, 122, 122, 62: 122, 122, 122, 122, 122, 122, 122, 122, 122, 122, 122, 122, 122, 122},
		// 335
		{123, 123, 123, 123, 123, 123, 123, 123, 123, 123, 123, 123, 13: 123, 123, 16: 123, 23: 123, 123, 33: 123, 123, 39: 123, 123, 47: 123, 123, 123, 123, 123, 123, 123, 123, 123, 123, 123, 123, 123, 123, 62: 123, 123, 123, 123, 123, 123, 123, 123, 123, 123, 123, 123, 123, 123},
		{2: 397, 398, 12: 396, 14: 325, 17: 327, 394, 339, 340, 328, 363, 25: 399, 400, 401, 402, 403, 405, 406, 78: 404, 80: 630, 82: 629},
		{124, 124, 124, 124, 124, 124, 124, 124, 124, 124, 124, 124, 13: 124, 124, 16: 124, 23: 124, 124, 33: 124, 124, 39: 124, 124, 47: 124, 124, 124, 124, 124, 124, 124, 124, 124, 124, 124, 124, 124, 124, 62: 124, 124, 124, 124, 124, 124, 124, 124, 124, 124, 124, 124, 124, 124},
		{11: 632},
		{11: 631},
		// 340
		{125, 125, 125, 125, 125, 125, 125, 125, 125, 125, 125, 125, 13: 125, 125, 16: 125, 23: 125, 125, 33: 125, 125, 39: 125, 125, 47: 125, 125, 125, 125, 125, 125, 125, 125, 125, 125, 125, 125, 125, 125, 62: 125, 125, 125, 125, 125, 125, 125, 125, 125, 125, 125, 125, 125, 125},
		{126, 126, 126, 126, 126, 126, 126, 126, 126, 126, 126, 126, 13: 126, 126, 16: 126, 23: 126, 126, 33: 126, 126, 39: 126, 126, 47: 126, 126, 126, 126, 126, 126, 126, 126, 126, 126, 126, 126, 126, 126, 62: 126, 126, 126, 126, 126, 126, 126, 126, 126, 126, 126, 126, 126, 126},
		{1: 254, 13: 254, 23: 645},
		{1: 251, 3: 424, 423, 425, 426, 430, 441, 13: 251, 23: 642, 47: 422, 427, 428, 429, 431, 432, 433, 434, 435, 436, 437, 438, 439, 440},
		{1: 248, 13: 248},
		// 345
		{1: 256, 13: 638, 111: 639},
		{242, 242, 242, 242, 242, 242, 242, 242, 242, 242, 242, 242, 13: 242, 242, 16: 242, 23: 242, 242, 33: 242, 242, 39: 242, 242, 47: 242, 242, 242, 242, 242, 242, 242, 242, 242, 242, 242, 242, 242, 242, 62: 242, 242, 242, 242, 242, 242, 242, 242, 242, 242, 242, 242, 242, 242},
		{1: 255, 341, 373, 377, 378, 379, 380, 381, 12: 412, 14: 325, 309, 17: 327, 394, 339, 340, 328, 363, 25: 329, 374, 343, 375, 330, 331, 332, 305, 35: 306, 307, 308, 376, 40: 612, 342, 346, 345, 347, 348, 334, 61: 634, 116: 633, 135: 641},
		{1: 640},
		{241, 241, 241, 241, 241, 241, 241, 241, 241, 241, 241, 241, 13: 241, 241, 16: 241, 23: 241, 241, 33: 241, 241, 39: 241, 241, 47: 241, 241, 241, 241, 241, 241, 241, 241, 241, 241, 241, 241, 241, 241, 62: 241, 241, 241, 241, 241, 241, 241, 241, 241, 241, 241, 241, 241, 241},
		// 350
		{1: 247, 13: 247},
		{2: 341, 373, 377, 378, 379, 380, 381, 12: 412, 14: 325, 309, 17: 327, 394, 339, 340, 328, 363, 25: 329, 374, 343, 375, 330, 331, 332, 305, 35: 306, 307, 308, 376, 40: 612, 342, 346, 345, 347, 348, 334, 61: 644, 116: 643},
		{1: 250, 13: 250},
		{1: 249, 3: 424, 423, 425, 426, 430, 441, 13: 249, 47: 422, 427, 428, 429, 431, 432, 433, 434, 435, 436, 437, 438, 439, 440},
		{2: 341, 373, 377, 378, 379, 380, 381, 12: 412, 14: 325, 309, 17: 327, 394, 339, 340, 328, 363, 25: 329, 374, 343, 375, 330, 331, 332, 305, 35: 306, 307, 308, 376, 40: 612, 342, 346, 345, 347, 348, 334, 61: 647, 116: 646},
		// 355
		{1: 253, 13: 253},
		{1: 252, 3: 424, 423, 425, 426, 430, 441, 13: 252, 47: 422, 427, 428, 429, 431, 432, 433, 434, 435, 436, 437, 438, 439, 440},
		{11: 256, 13: 650, 34: 652, 111: 651},
		{262, 262, 262, 262, 262, 262, 262, 262, 262, 262, 262, 262, 13: 262, 262, 16: 262, 23: 262, 262, 33: 262, 262, 39: 262, 262, 47: 262, 262, 262, 262, 262, 262, 262, 262, 262, 262, 262, 262, 262, 262, 62: 262, 262, 262, 262, 262, 262, 262, 262, 262, 262, 262, 262, 262, 262},
		{2: 341, 373, 377, 378, 379, 380, 381, 11: 255, 412, 14: 325, 309, 17: 327, 394, 339, 340, 328, 363, 25: 329, 374, 343, 375, 330, 331, 332, 305, 35: 306, 307, 308, 376, 41: 342, 346, 345, 347, 553, 334, 61: 552, 119: 565},
		// 360
		{11: 655},
		{11: 256, 13: 502, 111: 653},
		{11: 654},
		{260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 13: 260, 260, 16: 260, 23: 260, 260, 33: 260, 260, 39: 260, 260, 47: 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 62: 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 260, 260},
		{261, 261, 261, 261, 261, 261, 261, 261, 261, 261, 261, 261, 13: 261, 261, 16: 261, 23: 261, 261, 33: 261, 261, 39: 261, 261, 47: 261, 261, 261, 261, 261, 261, 261, 261, 261, 261, 261, 261, 261, 261, 62: 261, 261, 261, 261, 261, 261, 261, 261, 261, 261, 261, 261, 261, 261},
		// 365
		{1: 662, 341, 373, 377, 378, 379, 380, 381, 12: 412, 14: 325, 309, 570, 327, 394, 339, 340, 328, 363, 25: 329, 374, 343, 375, 330, 331, 332, 305, 35: 306, 307, 308, 376, 40: 571, 342, 346, 345, 347, 348, 334, 61: 658, 104: 656, 121: 659, 138: 660, 171: 661},
		{127, 127, 127, 127, 127, 127, 127, 127, 127, 127, 127, 127, 13: 127, 127, 16: 127, 23: 127, 127, 33: 127, 127, 39: 127, 127, 47: 127, 127, 127, 127, 127, 127, 127, 127, 127, 127, 127, 127, 127, 127, 62: 127, 127, 127, 127, 127, 127, 127, 127, 127, 127, 127, 127, 127, 127},
		{1: 155, 3: 424, 423, 425, 426, 430, 441, 13: 155, 23: 667, 47: 422, 427, 428, 429, 431, 432, 433, 434, 435, 436, 437, 438, 439, 440},
		{1: 152, 13: 152},
		{1: 151, 13: 151},
		// 370
		{1: 256, 13: 663, 111: 664},
		{149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 13: 149, 149, 16: 149, 23: 149, 149, 33: 149, 149, 39: 149, 149, 47: 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 62: 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149, 149},
		{1: 255, 341, 373, 377, 378, 379, 380, 381, 12: 412, 14: 325, 309, 570, 327, 394, 339, 340, 328, 363, 25: 329, 374, 343, 375, 330, 331, 332, 305, 35: 306, 307, 308, 376, 40: 571, 342, 346, 345, 347, 348, 334, 61: 658, 104: 656, 121: 659, 138: 666},
		{1: 665},
		{148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 13: 148, 148, 16: 148, 23: 148, 148, 33: 148, 148, 39: 148, 148, 47: 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 62: 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148},
		// 375
		{1: 150, 13: 150},
		{2: 341, 373, 377, 378, 379, 380, 381, 12: 412, 14: 325, 309, 570, 327, 394, 339, 340, 328, 363, 25: 329, 374, 343, 375, 330, 331, 332, 305, 35: 306, 307, 308, 376, 40: 571, 342, 346, 345, 347, 348, 334, 61: 668, 104: 656, 121: 669},
		{1: 154, 3: 424, 423, 425, 426, 430, 441, 13: 154, 47: 422, 427, 428, 429, 431, 432, 433, 434, 435, 436, 437, 438, 439, 440},
		{1: 153, 13: 153},
		{96, 96, 341, 373, 377, 378, 379, 380, 381, 96, 96, 12: 344, 14: 325, 309, 17: 327, 394, 339, 340, 328, 363, 25: 329, 374, 343, 375, 330, 331, 332, 305, 35: 306, 307, 308, 376, 40: 536, 342, 346, 345, 347, 348, 334, 61: 335, 77: 326, 79: 542, 81: 338, 372, 351, 333, 352, 353, 354, 336, 356, 357, 359, 349, 364, 382, 350, 361, 98: 537, 101: 538, 355, 358, 105: 360, 107: 541, 362, 539, 540, 114: 671},
		// 380
		{80, 80, 9: 80, 80},
		{16: 570, 40: 571, 104: 673},
		{96, 96, 341, 373, 377, 378, 379, 380, 381, 12: 344, 14: 325, 309, 17: 327, 394, 339, 340, 328, 363, 25: 329, 374, 343, 375, 330, 331, 332, 305, 35: 306, 307, 308, 376, 40: 536, 342, 346, 345, 347, 348, 334, 61: 335, 77: 326, 79: 542, 81: 338, 372, 351, 333, 352, 353, 354, 336, 356, 357, 359, 349, 364, 382, 350, 361, 98: 537, 101: 538, 355, 358, 105: 360, 107: 541, 362, 539, 540, 114: 543, 123: 674},
		{545, 675},
		{142, 142, 142, 142, 142, 142, 142, 142, 142, 142, 142, 142, 13: 142, 142, 16: 142, 23: 142, 142, 33: 142, 142, 39: 142, 142, 47: 142, 142, 142, 142, 142, 142, 142, 142, 142, 142, 142, 142, 142, 142, 62: 142, 142, 142, 142, 142, 142, 142, 142, 142, 142, 142, 142, 142, 142},
		// 385
		{3: 424, 423, 425, 426, 430, 441, 11: 679, 47: 422, 427, 428, 429, 431, 432, 433, 434, 435, 436, 437, 438, 439, 440},
		{2: 500, 11: 678},
		{145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 13: 145, 145, 16: 145, 23: 145, 145, 33: 145, 145, 39: 145, 145, 47: 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 62: 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 145, 145},
		{146, 146, 146, 146, 146, 146, 146, 146, 146, 146, 146, 146, 13: 146, 146, 16: 146, 23: 146, 146, 33: 146, 146, 39: 146, 146, 47: 146, 146, 146, 146, 146, 146, 146, 146, 146, 146, 146, 146, 146, 146, 62: 146, 146, 146, 146, 146, 146, 146, 146, 146, 146, 146, 146, 146, 146},
		{2: 397, 398, 12: 396, 14: 325, 17: 327, 394, 339, 340, 328, 363, 25: 399, 400, 401, 402, 403, 405, 406, 78: 404, 80: 681},
		// 390
		{39: 682},
		{2: 397, 398, 12: 396, 14: 325, 17: 327, 394, 339, 340, 328, 363, 25: 399, 400, 401, 402, 403, 405, 406, 78: 404, 80: 683},
		{147, 147, 147, 9: 147, 147, 147, 13: 147, 15: 147, 147, 23: 147, 147, 33: 147, 147, 39: 147, 147, 76: 147},
		{1: 685, 12: 164, 148: 686},
		{165, 165, 165, 9: 165, 165, 165, 13: 165, 15: 165, 165, 23: 165, 165, 33: 165, 165, 39: 165, 165, 76: 165},
		// 395
		{12: 688, 78: 689, 137: 690, 170: 687},
		{694, 107, 113: 693},
		{140, 140, 162, 75: 467, 149: 691},
		{160, 160},
		{159, 159},
		// 400
		{2: 471, 112: 472, 122: 692},
		{161, 161},
		{1: 696},
		{1: 106, 12: 688, 78: 689, 137: 695},
		{158, 158},
		// 405
		{163, 163, 163, 9: 163, 163, 163, 13: 163, 15: 163, 163, 23: 163, 163, 33: 163, 163, 39: 163, 163, 76: 163},
		{16: 698, 124: 699},
		{264, 264, 264, 264, 264, 264, 264, 264, 264, 12: 264, 14: 264, 264, 17: 264, 264, 264, 264, 264, 264, 32: 264, 35: 264, 264, 264, 264, 40: 264, 79: 264, 81: 264, 264, 264, 264, 264, 264, 264, 264, 264, 264, 264, 264, 264, 264, 146: 708},
		{231, 231, 9: 231, 231, 117: 231, 158: 700},
		{229, 229, 9: 229, 229, 117: 701, 157: 702, 159: 703},
		// 410
		{40: 536, 81: 704, 98: 705},
		{230, 230, 9: 230, 230, 117: 230},
		{177, 177, 9: 177, 177},
		{99, 2: 341, 373, 377, 378, 379, 380, 381, 12: 412, 14: 325, 309, 99, 327, 394, 339, 340, 328, 363, 25: 329, 374, 343, 375, 330, 331, 332, 305, 35: 306, 307, 308, 376, 41: 342, 346, 345, 347, 348, 334, 61: 335, 77: 326, 95: 350, 525, 106: 524, 128: 706},
		{228, 228, 9: 228, 228},
		// 415
		{16: 698, 124: 707},
		{232, 232, 9: 232, 232, 117: 232},
		{96, 96, 341, 373, 377, 378, 379, 380, 381, 12: 344, 14: 325, 309, 17: 327, 394, 339, 340, 328, 363, 25: 329, 374, 343, 375, 330, 331, 332, 305, 35: 306, 307, 308, 376, 40: 536, 342, 346, 345, 347, 348, 334, 61: 335, 77: 326, 79: 542, 81: 338, 372, 351, 333, 352, 353, 354, 336, 356, 357, 359, 349, 364, 382, 350, 361, 98: 537, 101: 538, 355, 358, 105: 360, 107: 541, 362, 539, 540, 114: 543, 123: 709},
		{545, 710},
		{263, 263, 9: 263, 263, 117: 263},
		// 420
		{12: 713},
		{2: 473, 398, 12: 112, 14: 325, 16: 111, 327, 394, 339, 340, 328, 363, 25: 399, 400, 401, 402, 403, 405, 406, 40: 111, 78: 404, 80: 475, 112: 474, 140: 476},
		{2: 185, 100: 508, 136: 714},
		{2: 471, 112: 472, 122: 715},
		{194, 40: 194, 147: 716},
		// 425
		{196, 40: 536, 98: 717, 163: 718},
		{195},
		{193},
		{13: 420, 24: 730, 33: 742, 62: 731, 733, 732, 734, 735, 736, 737, 738, 739, 740, 741},
		{16: 200},
		// 430
		{726, 16: 198},
		{16: 698, 124: 725},
		{2: 341, 373, 377, 378, 379, 380, 381, 12: 412, 14: 325, 309, 17: 327, 394, 339, 340, 328, 363, 25: 329, 374, 343, 375, 330, 331, 332, 305, 35: 306, 307, 308, 376, 41: 342, 346, 345, 347, 348, 334, 61: 724},
		{3: 424, 423, 425, 426, 430, 441, 16: 114, 47: 422, 427, 428, 429, 431, 432, 433, 434, 435, 436, 437, 438, 439, 440},
		{197, 197, 9: 197, 197},
		// 435
		{99, 2: 341, 373, 377, 378, 379, 380, 381, 12: 412, 14: 325, 309, 17: 327, 394, 339, 340, 328, 363, 25: 329, 374, 343, 375, 330, 331, 332, 305, 35: 306, 307, 308, 376, 41: 342, 346, 345, 347, 348, 334, 61: 335, 77: 326, 95: 350, 525, 106: 727},
		{728},
		{2: 341, 373, 377, 378, 379, 380, 381, 12: 412, 14: 325, 309, 99, 327, 394, 339, 340, 328, 363, 25: 329, 374, 343, 375, 330, 331, 332, 305, 35: 306, 307, 308, 376, 41: 342, 346, 345, 347, 348, 334, 61: 335, 77: 326, 95: 350, 525, 106: 729},
		{16: 199},
		{2: 341, 373, 377, 378, 379, 380, 381, 12: 412, 14: 325, 309, 17: 327, 394, 339, 340, 328, 363, 25: 329, 374, 343, 375, 330, 331, 332, 305, 35: 306, 307, 308, 376, 41: 342, 346, 345, 347, 348, 334, 61: 410, 77: 757, 131: 758},
		// 440
		{2: 341, 373, 377, 378, 379, 380, 381, 12: 412, 14: 325, 309, 17: 327, 394, 339, 340, 328, 363, 25: 329, 374, 343, 375, 330, 331, 332, 305, 35: 306, 307, 308, 376, 41: 342, 346, 345, 347, 348, 334, 61: 410, 77: 756},
		{2: 341, 373, 377, 378, 379, 380, 381, 12: 412, 14: 325, 309, 17: 327, 394, 339, 340, 328, 363, 25: 329, 374, 343, 375, 330, 331, 332, 305, 35: 306, 307, 308, 376, 41: 342, 346, 345, 347, 348, 334, 61: 410, 77: 755},
		{2: 341, 373, 377, 378, 379, 380, 381, 12: 412, 14: 325, 309, 17: 327, 394, 339, 340, 328, 363, 25: 329, 374, 343, 375, 330, 331, 332, 305, 35: 306, 307, 308, 376, 41: 342, 346, 345, 347, 348, 334, 61: 410, 77: 754},
		{2: 341, 373, 377, 378, 379, 380, 381, 12: 412, 14: 325, 309, 17: 327, 394, 339, 340, 328, 363, 25: 329, 374, 343, 375, 330, 331, 332, 305, 35: 306, 307, 308, 376, 41: 342, 346, 345, 347, 348, 334, 61: 410, 77: 753},
		{2: 341, 373, 377, 378, 379, 380, 381, 12: 412, 14: 325, 309, 17: 327, 394, 339, 340, 328, 363, 25: 329, 374, 343, 375, 330, 331, 332, 305, 35: 306, 307, 308, 376, 41: 342, 346, 345, 347, 348, 334, 61: 410, 77: 752},
		// 445
		{2: 341, 373, 377, 378, 379, 380, 381, 12: 412, 14: 325, 309, 17: 327, 394, 339, 340, 328, 363, 25: 329, 374, 343, 375, 330, 331, 332, 305, 35: 306, 307, 308, 376, 41: 342, 346, 345, 347, 348, 334, 61: 410, 77: 751},
		{2: 341, 373, 377, 378, 379, 380, 381, 12: 412, 14: 325, 309, 17: 327, 394, 339, 340, 328, 363, 25: 329, 374, 343, 375, 330, 331, 332, 305, 35: 306, 307, 308, 376, 41: 342, 346, 345, 347, 348, 334, 61: 410, 77: 750},
		{2: 341, 373, 377, 378, 379, 380, 381, 12: 412, 14: 325, 309, 17: 327, 394, 339, 340, 328, 363, 25: 329, 374, 343, 375, 330, 331, 332, 305, 35: 306, 307, 308, 376, 41: 342, 346, 345, 347, 348, 334, 61: 410, 77: 749},
		{2: 341, 373, 377, 378, 379, 380, 381, 12: 412, 14: 325, 309, 17: 327, 394, 339, 340, 328, 363, 25: 329, 374, 343, 375, 330, 331, 332, 305, 35: 306, 307, 308, 376, 41: 342, 346, 345, 347, 348, 334, 61: 410, 77: 748},
		{2: 341, 373, 377, 378, 379, 380, 381, 12: 412, 14: 325, 309, 17: 327, 394, 339, 340, 328, 363, 25: 329, 374, 343, 375, 330, 331, 332, 305, 35: 306, 307, 308, 376, 41: 342, 346, 345, 347, 348, 334, 61: 410, 77: 747},
		// 450
		{2: 341, 373, 377, 378, 379, 380, 381, 12: 412, 14: 325, 309, 17: 327, 394, 339, 340, 328, 363, 25: 329, 374, 343, 375, 330, 331, 332, 305, 35: 306, 307, 308, 376, 41: 342, 346, 345, 347, 348, 334, 61: 410, 77: 746},
		{2: 341, 373, 377, 378, 379, 380, 381, 12: 412, 14: 325, 309, 17: 327, 394, 339, 340, 328, 363, 25: 329, 374, 343, 375, 330, 331, 332, 305, 35: 306, 307, 308, 376, 41: 342, 346, 345, 347, 348, 334, 61: 410, 77: 743, 131: 744},
		{100, 100, 9: 100, 100, 13: 420, 16: 100},
		{2: 341, 373, 377, 378, 379, 380, 381, 12: 412, 14: 325, 309, 17: 327, 394, 339, 340, 328, 363, 25: 329, 374, 343, 375, 330, 331, 332, 305, 35: 306, 307, 308, 376, 41: 342, 346, 345, 347, 348, 334, 61: 745},
		{3: 424, 423, 425, 426, 430, 441, 16: 115, 47: 422, 427, 428, 429, 431, 432, 433, 434, 435, 436, 437, 438, 439, 440},
		// 455
		{272, 272, 9: 272, 272, 13: 420, 16: 272},
		{273, 273, 9: 273, 273, 13: 420, 16: 273},
		{274, 274, 9: 274, 274, 13: 420, 16: 274},
		{275, 275, 9: 275, 275, 13: 420, 16: 275},
		{276, 276, 9: 276, 276, 13: 420, 16: 276},
		// 460
		{277, 277, 9: 277, 277, 13: 420, 16: 277},
		{278, 278, 9: 278, 278, 13: 420, 16: 278},
		{279, 279, 9: 279, 279, 13: 420, 16: 279},
		{280, 280, 9: 280, 280, 13: 420, 16: 280},
		{281, 281, 9: 281, 281, 13: 420, 16: 281},
		// 465
		{282, 282, 9: 282, 282, 13: 420, 16: 282},
		{283, 283, 9: 283, 283, 13: 420, 16: 283},
		{2: 341, 373, 377, 378, 379, 380, 381, 12: 412, 14: 325, 309, 17: 327, 394, 339, 340, 328, 363, 25: 329, 374, 343, 375, 330, 331, 332, 305, 35: 306, 307, 308, 376, 41: 342, 346, 345, 347, 348, 334, 61: 759},
		{3: 424, 423, 425, 426, 430, 441, 16: 116, 47: 422, 427, 428, 429, 431, 432, 433, 434, 435, 436, 437, 438, 439, 440},
		{102, 102, 9: 102, 102, 16: 102},
		// 470
		{101, 101, 9: 101, 101, 16: 101},
		{11: 770, 383, 99: 764, 125: 772, 156: 771},
		{238, 238, 9: 238, 238},
		{237, 237, 397, 398, 9: 237, 237, 237, 396, 395, 325, 17: 327, 394, 339, 340, 328, 363, 24: 765, 399, 400, 401, 402, 403, 405, 406, 78: 404, 80: 766},
		{2: 341, 373, 377, 378, 379, 380, 381, 12: 412, 14: 325, 309, 17: 327, 394, 339, 340, 328, 363, 25: 329, 374, 343, 375, 330, 331, 332, 305, 35: 306, 307, 308, 376, 41: 342, 346, 345, 347, 348, 334, 61: 410, 77: 769},
		// 475
		{24: 767},
		{2: 341, 373, 377, 378, 379, 380, 381, 12: 412, 14: 325, 309, 17: 327, 394, 339, 340, 328, 363, 25: 329, 374, 343, 375, 330, 331, 332, 305, 35: 306, 307, 308, 376, 41: 342, 346, 345, 347, 348, 334, 61: 410, 77: 768},
		{235, 235, 9: 235, 235, 235, 13: 420},
		{236, 236, 9: 236, 236, 236, 13: 420},
		{240, 240, 9: 240, 240},
		// 480
		{774, 11: 107, 113: 773},
		{234, 11: 234},
		{11: 776},
		{11: 106, 383, 99: 764, 125: 775},
		{233, 11: 233},
		// 485
		{239, 239, 9: 239, 239},
		{2: 397, 398, 12: 396, 14: 325, 17: 327, 394, 339, 340, 328, 363, 25: 399, 400, 401, 402, 403, 405, 406, 78: 404, 80: 778},
		{257, 257, 257, 9: 257, 257, 257, 13: 257, 15: 257, 257, 23: 257, 257, 33: 257, 257, 39: 257, 257, 76: 257},
		{259, 259, 259, 9: 259, 259, 259, 13: 259, 15: 259, 259, 23: 259, 259, 33: 259, 259, 39: 259, 259, 76: 259},
		{2: 397, 398, 12: 396, 14: 325, 17: 327, 394, 339, 340, 328, 363, 25: 399, 400, 401, 402, 403, 405, 406, 78: 404, 80: 781},
		// 490
		{258, 258, 258, 9: 258, 258, 258, 13: 258, 15: 258, 258, 23: 258, 258, 33: 258, 258, 39: 258, 258, 76: 258},
		{2: 341, 373, 377, 378, 379, 380, 381, 12: 412, 14: 325, 309, 17: 327, 394, 339, 340, 328, 363, 25: 329, 374, 343, 375, 330, 331, 332, 305, 35: 306, 307, 308, 376, 41: 342, 346, 345, 347, 348, 334, 61: 410, 77: 757},
		{2: 341, 373, 377, 378, 379, 380, 381, 12: 412, 14: 325, 309, 17: 327, 394, 339, 340, 328, 363, 25: 329, 374, 343, 375, 330, 331, 332, 305, 35: 306, 307, 308, 376, 41: 342, 346, 345, 347, 348, 334, 61: 410, 77: 743},
		{39: 790},
		{3: 424, 423, 425, 426, 430, 441, 39: 788, 47: 422, 427, 428, 429, 431, 432, 433, 434, 435, 436, 437, 438, 439, 440},
		// 495
		{2: 397, 398, 12: 396, 14: 325, 17: 327, 394, 339, 340, 328, 363, 25: 399, 400, 401, 402, 403, 405, 406, 78: 404, 80: 787},
		{97, 97, 97, 9: 97, 97, 97, 13: 97, 15: 97, 97, 23: 97, 97, 33: 97, 97, 39: 97, 97, 76: 97},
		{2: 397, 398, 12: 396, 14: 325, 17: 327, 394, 339, 340, 328, 363, 25: 399, 400, 401, 402, 403, 405, 406, 78: 404, 80: 789},
		{284, 284, 284, 9: 284, 284, 284, 13: 284, 15: 284, 284, 23: 284, 284, 33: 284, 284, 39: 284, 284, 76: 284},
		{2: 397, 398, 12: 396, 14: 325, 17: 327, 394, 339, 340, 328, 363, 25: 399, 400, 401, 402, 403, 405, 406, 78: 404, 80: 791},
		// 500
		{285, 285, 285, 9: 285, 285, 285, 13: 285, 15: 285, 285, 23: 285, 285, 33: 285, 285, 39: 285, 285, 76: 285},
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
			lx.popScope() // Implicit block.
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
			lx.popScope() // Implicit block.
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
			switch r := yyS[yypt-3].node.(*ReceiverOpt); {
			case r == nil: // Function.
				lx.scope.Parent.declare(lx, newFuncDeclaration(yyS[yypt-2].Token))
			default: // Method.
				//TODO
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
			yyVAL.node = &FuncType{
				Token:     yyS[yypt-1].Token,
				Signature: yyS[yypt-0].node.(*Signature),
			}
		}
	case 100:
		{
			yyVAL.node = &GenericArgumentList{
				Typ: yyS[yypt-0].node.(*Typ),
			}
		}
	case 101:
		{
			yyVAL.node = &GenericArgumentList{
				Case:                1,
				GenericArgumentList: yyS[yypt-2].node.(*GenericArgumentList),
				Token:               yyS[yypt-1].Token,
				Typ:                 yyS[yypt-0].node.(*Typ),
			}
		}
	case 102:
		{
			yyVAL.node = (*GenericArgumentsOpt)(nil)
		}
	case 103:
		{
			yyVAL.node = &GenericArgumentsOpt{
				Token:               yyS[yypt-2].Token,
				GenericArgumentList: yyS[yypt-1].node.(*GenericArgumentList).reverse(),
				Token2:              yyS[yypt-0].Token,
			}
		}
	case 104:
		{
			yyVAL.node = &GenericParameterList{
				Token: yyS[yypt-0].Token,
			}
		}
	case 105:
		{
			yyVAL.node = &GenericParameterList{
				Case:                 1,
				GenericParameterList: yyS[yypt-2].node.(*GenericParameterList),
				Token:                yyS[yypt-1].Token,
				Token2:               yyS[yypt-0].Token,
			}
		}
	case 106:
		{
			yyVAL.node = (*GenericParametersOpt)(nil)
		}
	case 107:
		{
			yyVAL.node = &GenericParametersOpt{
				Token:                yyS[yypt-2].Token,
				GenericParameterList: yyS[yypt-1].node.(*GenericParameterList).reverse(),
				Token2:               yyS[yypt-0].Token,
			}
		}
	case 108:
		{
			yyVAL.node = (*IdentifierOpt)(nil)
		}
	case 109:
		{
			yyVAL.node = &IdentifierOpt{
				Token: yyS[yypt-0].Token,
			}
		}
	case 110:
		{
			yyVAL.node = &IdentifierList{
				Token: yyS[yypt-0].Token,
			}
		}
	case 111:
		{
			yyVAL.node = &IdentifierList{
				Case:           1,
				IdentifierList: yyS[yypt-2].node.(*IdentifierList),
				Token:          yyS[yypt-1].Token,
				Token2:         yyS[yypt-0].Token,
			}
		}
	case 112:
		{
			yyVAL.node = &IfHeader{
				SimpleStatementOpt: yyS[yypt-0].node.(*SimpleStatementOpt),
			}
		}
	case 113:
		{
			yyVAL.node = &IfHeader{
				Case:                1,
				SimpleStatementOpt:  yyS[yypt-2].node.(*SimpleStatementOpt),
				Token:               yyS[yypt-1].Token,
				SimpleStatementOpt2: yyS[yypt-0].node.(*SimpleStatementOpt),
			}
		}
	case 114:
		{
			lx := yylex.(*lexer)
			yyVAL.node = &IfStatement{
				Token:    yyS[yypt-4].Token,
				IfHeader: yyS[yypt-3].node.(*IfHeader),
				Body:     yyS[yypt-2].node.(*Body),
				ElifList: yyS[yypt-1].node.(*ElifList).reverse(),
				ElseOpt:  yyS[yypt-0].node.(*ElseOpt),
			}
			lx.popScope() // Implicit block.
		}
	case 115:
		{
			yyVAL.node = &ImportDecl{
				Token:  yyS[yypt-2].Token,
				Token2: yyS[yypt-1].Token,
				Token3: yyS[yypt-0].Token,
			}
		}
	case 116:
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
	case 117:
		{
			yyVAL.node = &ImportDecl{
				Case:       2,
				Token:      yyS[yypt-1].Token,
				ImportSpec: yyS[yypt-0].node.(*ImportSpec),
			}
		}
	case 118:
		{
			lx := yylex.(*lexer)
			lhs := &ImportSpec{
				Token:        yyS[yypt-1].Token,
				BasicLiteral: yyS[yypt-0].node.(*BasicLiteral),
			}
			yyVAL.node = lhs
			lhs.post(lx)
		}
	case 119:
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
	case 120:
		{
			yyVAL.node = &ImportSpec{
				Case:         2,
				Token:        yyS[yypt-2].Token,
				BasicLiteral: yyS[yypt-1].node.(*BasicLiteral),
			}
		}
	case 121:
		{
			yyVAL.node = &ImportSpec{
				Case:          3,
				IdentifierOpt: yyS[yypt-2].node.(*IdentifierOpt),
				BasicLiteral:  yyS[yypt-1].node.(*BasicLiteral),
			}
		}
	case 122:
		{
			yyVAL.node = &ImportSpecList{
				ImportSpec: yyS[yypt-0].node.(*ImportSpec),
			}
		}
	case 123:
		{
			yyVAL.node = &ImportSpecList{
				Case:           1,
				ImportSpecList: yyS[yypt-2].node.(*ImportSpecList),
				Token:          yyS[yypt-1].Token,
				ImportSpec:     yyS[yypt-0].node.(*ImportSpec),
			}
		}
	case 124:
		{
			yyVAL.node = (*ImportList)(nil)
		}
	case 125:
		{
			yyVAL.node = &ImportList{
				ImportList: yyS[yypt-2].node.(*ImportList),
				ImportDecl: yyS[yypt-1].node.(*ImportDecl),
				Token:      yyS[yypt-0].Token,
			}
		}
	case 126:
		{
			yyVAL.node = &InterfaceType{
				Token:  yyS[yypt-2].Token,
				LBrace: yyS[yypt-1].node.(*LBrace),
				Token2: yyS[yypt-0].Token,
			}
		}
	case 127:
		{
			lx := yylex.(*lexer)
			lx.pushScope()
		}
	case 128:
		{
			lx := yylex.(*lexer)
			yyVAL.node = &InterfaceType{
				Case:                    1,
				Token:                   yyS[yypt-5].Token,
				LBrace:                  yyS[yypt-4].node.(*LBrace),
				InterfaceMethodDeclList: yyS[yypt-2].node.(*InterfaceMethodDeclList).reverse(),
				SemicolonOpt:            yyS[yypt-1].node.(*SemicolonOpt),
				Token2:                  yyS[yypt-0].Token,
			}
			lx.popScope()
		}
	case 129:
		{
			lx := yylex.(*lexer)
			lx.pushScope()
		}
	case 130:
		{
			lx := yylex.(*lexer)
			yyVAL.node = &InterfaceMethodDecl{
				Token:     yyS[yypt-2].Token,
				Signature: yyS[yypt-0].node.(*Signature),
			}
			lx.popScope()
		}
	case 131:
		{
			yyVAL.node = &InterfaceMethodDecl{
				Case:           1,
				QualifiedIdent: yyS[yypt-0].node.(*QualifiedIdent),
			}
		}
	case 132:
		{
			yyVAL.node = &InterfaceMethodDeclList{
				InterfaceMethodDecl: yyS[yypt-0].node.(*InterfaceMethodDecl),
			}
		}
	case 133:
		{
			yyVAL.node = &InterfaceMethodDeclList{
				Case: 1,
				InterfaceMethodDeclList: yyS[yypt-2].node.(*InterfaceMethodDeclList),
				Token:               yyS[yypt-1].Token,
				InterfaceMethodDecl: yyS[yypt-0].node.(*InterfaceMethodDecl),
			}
		}
	case 134:
		{
			lx := yylex.(*lexer)
			yyVAL.node = &LBrace{
				Token: yyS[yypt-0].Token,
			}
			lx.fixLBR()
		}
	case 135:
		{
			yyVAL.node = &LBrace{
				Case:  1,
				Token: yyS[yypt-0].Token,
			}
		}
	case 136:
		{
			yyVAL.node = &LBraceCompLitItem{
				Expression: yyS[yypt-0].node.(*Expression),
			}
		}
	case 137:
		{
			yyVAL.node = &LBraceCompLitItem{
				Case:        1,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
		}
	case 138:
		{
			yyVAL.node = &LBraceCompLitItem{
				Case:               2,
				Expression:         yyS[yypt-2].node.(*Expression),
				Token:              yyS[yypt-1].Token,
				LBraceCompLitValue: yyS[yypt-0].node.(*LBraceCompLitValue),
			}
		}
	case 139:
		{
			yyVAL.node = &LBraceCompLitItem{
				Case:               3,
				LBraceCompLitValue: yyS[yypt-0].node.(*LBraceCompLitValue),
			}
		}
	case 140:
		{
			yyVAL.node = &LBraceCompLitItemList{
				LBraceCompLitItem: yyS[yypt-0].node.(*LBraceCompLitItem),
			}
		}
	case 141:
		{
			yyVAL.node = &LBraceCompLitItemList{
				Case: 1,
				LBraceCompLitItemList: yyS[yypt-2].node.(*LBraceCompLitItemList),
				Token:             yyS[yypt-1].Token,
				LBraceCompLitItem: yyS[yypt-0].node.(*LBraceCompLitItem),
			}
		}
	case 142:
		{
			yyVAL.node = &LBraceCompLitValue{
				LBrace: yyS[yypt-1].node.(*LBrace),
				Token:  yyS[yypt-0].Token,
			}
		}
	case 143:
		{
			yyVAL.node = &LBraceCompLitValue{
				Case:                  1,
				LBrace:                yyS[yypt-3].node.(*LBrace),
				LBraceCompLitItemList: yyS[yypt-2].node.(*LBraceCompLitItemList).reverse(),
				CommaOpt:              yyS[yypt-1].node.(*CommaOpt),
				Token:                 yyS[yypt-0].Token,
			}
		}
	case 144:
		{
			yyVAL.node = &MapType{
				Token:  yyS[yypt-4].Token,
				Token2: yyS[yypt-3].Token,
				Typ:    yyS[yypt-2].node.(*Typ),
				Token3: yyS[yypt-1].Token,
				Typ2:   yyS[yypt-0].node.(*Typ),
			}
		}
	case 145:
		{
			yyVAL.node = &Operand{
				Token:      yyS[yypt-2].Token,
				Expression: yyS[yypt-1].node.(*Expression),
				Token2:     yyS[yypt-0].Token,
			}
		}
	case 146:
		{
			yyVAL.node = &Operand{
				Case:        1,
				Token:       yyS[yypt-2].Token,
				TypeLiteral: yyS[yypt-1].node.(*TypeLiteral),
				Token2:      yyS[yypt-0].Token,
			}
		}
	case 147:
		{
			yyVAL.node = &Operand{
				Case:         2,
				BasicLiteral: yyS[yypt-0].node.(*BasicLiteral),
			}
		}
	case 148:
		{
			lx := yylex.(*lexer)
			lx.scope.isMergeScope = false
		}
	case 149:
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
	case 150:
		{
			yyVAL.node = &Operand{
				Case:                4,
				Token:               yyS[yypt-1].Token,
				GenericArgumentsOpt: yyS[yypt-0].node.(*GenericArgumentsOpt),
			}
		}
	case 151:
		{
			yyVAL.node = &QualifiedIdent{
				Token: yyS[yypt-0].Token,
			}
		}
	case 152:
		{
			yyVAL.node = &QualifiedIdent{
				Case:   1,
				Token:  yyS[yypt-2].Token,
				Token2: yyS[yypt-1].Token,
				Token3: yyS[yypt-0].Token,
			}
		}
	case 153:
		{
			lx := yylex.(*lexer)
			if !lx.build { // Build tags not satisfied
				return 0
			}
		}
	case 154:
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
	case 155:
		{
			yyVAL.node = &ParameterDecl{
				Token: yyS[yypt-1].Token,
				Typ:   yyS[yypt-0].node.(*Typ),
			}
		}
	case 156:
		{
			yyVAL.node = &ParameterDecl{
				Case:   1,
				Token:  yyS[yypt-2].Token,
				Token2: yyS[yypt-1].Token,
				Typ:    yyS[yypt-0].node.(*Typ),
			}
		}
	case 157:
		{
			yyVAL.node = &ParameterDecl{
				Case:  2,
				Token: yyS[yypt-1].Token,
				Typ:   yyS[yypt-0].node.(*Typ),
			}
		}
	case 158:
		{
			yyVAL.node = &ParameterDecl{
				Case: 3,
				Typ:  yyS[yypt-0].node.(*Typ),
			}
		}
	case 159:
		{
			yyVAL.node = &ParameterDeclList{
				ParameterDecl: yyS[yypt-0].node.(*ParameterDecl),
			}
		}
	case 160:
		{
			yyVAL.node = &ParameterDeclList{
				Case:              1,
				ParameterDeclList: yyS[yypt-2].node.(*ParameterDeclList),
				Token:             yyS[yypt-1].Token,
				ParameterDecl:     yyS[yypt-0].node.(*ParameterDecl),
			}
		}
	case 161:
		{
			yyVAL.node = &Parameters{
				Token:  yyS[yypt-1].Token,
				Token2: yyS[yypt-0].Token,
			}
		}
	case 162:
		{
			lx := yylex.(*lexer)
			lhs := &Parameters{
				Case:              1,
				Token:             yyS[yypt-3].Token,
				ParameterDeclList: yyS[yypt-2].node.(*ParameterDeclList).reverse(),
				CommaOpt:          yyS[yypt-1].node.(*CommaOpt),
				Token2:            yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.post(lx)
		}
	case 163:
		{
			yyVAL.node = &PrimaryExpression{
				Operand: yyS[yypt-0].node.(*Operand),
			}
		}
	case 164:
		{
			yyVAL.node = &PrimaryExpression{
				Case:               1,
				CompLitType:        yyS[yypt-1].node.(*CompLitType),
				LBraceCompLitValue: yyS[yypt-0].node.(*LBraceCompLitValue),
			}
		}
	case 165:
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
	case 166:
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
	case 167:
		{
			yyVAL.node = &PrimaryExpression{
				Case:              4,
				PrimaryExpression: yyS[yypt-2].node.(*PrimaryExpression),
				Token:             yyS[yypt-1].Token,
				Token2:            yyS[yypt-0].Token,
			}
		}
	case 168:
		{
			yyVAL.node = &PrimaryExpression{
				Case:              5,
				PrimaryExpression: yyS[yypt-3].node.(*PrimaryExpression),
				Token:             yyS[yypt-2].Token,
				Expression:        yyS[yypt-1].node.(*Expression),
				Token2:            yyS[yypt-0].Token,
			}
		}
	case 169:
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
	case 170:
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
	case 171:
		{
			yyVAL.node = &PrimaryExpression{
				Case:              8,
				PrimaryExpression: yyS[yypt-1].node.(*PrimaryExpression),
				Call:              yyS[yypt-0].node.(*Call),
			}
		}
	case 172:
		{
			yyVAL.node = &PrimaryExpression{
				Case:              9,
				PrimaryExpression: yyS[yypt-1].node.(*PrimaryExpression),
				CompLitValue:      yyS[yypt-0].node.(*CompLitValue),
			}
		}
	case 173:
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
	case 174:
		{
			lx := yylex.(*lexer)
			lhs := &Prologue{
				PackageClause: yyS[yypt-1].node.(*PackageClause),
				ImportList:    yyS[yypt-0].node.(*ImportList).reverse(),
			}
			yyVAL.node = lhs
			lhs.post(lx)
		}
	case 175:
		{
			yyVAL.node = &Range{
				ExpressionList: yyS[yypt-3].node.(*ExpressionList).reverse(),
				Token:          yyS[yypt-2].Token,
				Token2:         yyS[yypt-1].Token,
				Expression:     yyS[yypt-0].node.(*Expression),
			}
		}
	case 176:
		{
			yyVAL.node = &Range{
				Case:           1,
				ExpressionList: yyS[yypt-3].node.(*ExpressionList).reverse(),
				Token:          yyS[yypt-2].Token,
				Token2:         yyS[yypt-1].Token,
				Expression:     yyS[yypt-0].node.(*Expression),
			}
		}
	case 177:
		{
			yyVAL.node = &Range{
				Case:       2,
				Token:      yyS[yypt-1].Token,
				Expression: yyS[yypt-0].node.(*Expression),
			}
		}
	case 178:
		{
			yyVAL.node = (*ReceiverOpt)(nil)
		}
	case 179:
		{
			yyVAL.node = &ReceiverOpt{
				Parameters: yyS[yypt-0].node.(*Parameters),
			}
		}
	case 180:
		{
			yyVAL.node = (*ResultOpt)(nil)
		}
	case 181:
		{
			yyVAL.node = &ResultOpt{
				Case:       1,
				Parameters: yyS[yypt-0].node.(*Parameters),
			}
		}
	case 182:
		{
			yyVAL.node = &ResultOpt{
				Case: 2,
				Typ:  yyS[yypt-0].node.(*Typ),
			}
		}
	case 183:
		{
			yyVAL.node = &SelectStatement{
				Token:      yyS[yypt-1].Token,
				SwitchBody: yyS[yypt-0].node.(*SwitchBody),
			}
		}
	case 184:
		{
			yyVAL.node = (*SemicolonOpt)(nil)
		}
	case 185:
		{
			yyVAL.node = &SemicolonOpt{
				Token: yyS[yypt-0].Token,
			}
		}
	case 186:
		{
			yyVAL.node = &Signature{
				Parameters: yyS[yypt-1].node.(*Parameters),
				ResultOpt:  yyS[yypt-0].node.(*ResultOpt),
			}
		}
	case 187:
		{
			yyVAL.node = &SimpleStatement{
				Assignment: yyS[yypt-0].node.(*Assignment),
			}
		}
	case 188:
		{
			yyVAL.node = &SimpleStatement{
				Case:       1,
				Expression: yyS[yypt-0].node.(*Expression),
			}
		}
	case 189:
		{
			yyVAL.node = &SimpleStatement{
				Case:       2,
				Expression: yyS[yypt-1].node.(*Expression),
				Token:      yyS[yypt-0].Token,
			}
		}
	case 190:
		{
			yyVAL.node = &SimpleStatement{
				Case:       3,
				Expression: yyS[yypt-1].node.(*Expression),
				Token:      yyS[yypt-0].Token,
			}
		}
	case 191:
		{
			yyVAL.node = &SimpleStatement{
				Case:            4,
				ExpressionList:  yyS[yypt-2].node.(*ExpressionList).reverse(),
				Token:           yyS[yypt-1].Token,
				ExpressionList2: yyS[yypt-0].node.(*ExpressionList).reverse(),
			}
		}
	case 192:
		{
			yyVAL.node = (*SimpleStatementOpt)(nil)
		}
	case 193:
		{
			yyVAL.node = &SimpleStatementOpt{
				SimpleStatement: yyS[yypt-0].node.(*SimpleStatement),
			}
		}
	case 194:
		{
			yyVAL.node = &SliceType{
				Token:  yyS[yypt-2].Token,
				Token2: yyS[yypt-1].Token,
				Typ:    yyS[yypt-0].node.(*Typ),
			}
		}
	case 195:
		{
			yyVAL.node = (*Statement)(nil)
		}
	case 196:
		{
			yyVAL.node = &Statement{
				Case:  1,
				Block: yyS[yypt-0].node.(*Block),
			}
		}
	case 197:
		{
			yyVAL.node = &Statement{
				Case:      2,
				ConstDecl: yyS[yypt-0].node.(*ConstDecl),
			}
		}
	case 198:
		{
			yyVAL.node = &Statement{
				Case:     3,
				TypeDecl: yyS[yypt-0].node.(*TypeDecl),
			}
		}
	case 199:
		{
			yyVAL.node = &Statement{
				Case:    4,
				VarDecl: yyS[yypt-0].node.(*VarDecl),
			}
		}
	case 200:
		{
			yyVAL.node = &Statement{
				Case:             5,
				StatementNonDecl: yyS[yypt-0].node.(*StatementNonDecl),
			}
		}
	case 201:
		{
			yyVAL.node = &Statement{
				Case: 6,
			}
		}
	case 202:
		{
			yyVAL.node = &StatementList{
				Statement: yyS[yypt-0].node.(*Statement),
			}
		}
	case 203:
		{
			yyVAL.node = &StatementList{
				Case:          1,
				StatementList: yyS[yypt-2].node.(*StatementList),
				Token:         yyS[yypt-1].Token,
				Statement:     yyS[yypt-0].node.(*Statement),
			}
		}
	case 204:
		{
			yyVAL.node = &StatementNonDecl{
				Token:         yyS[yypt-1].Token,
				IdentifierOpt: yyS[yypt-0].node.(*IdentifierOpt),
			}
		}
	case 205:
		{
			yyVAL.node = &StatementNonDecl{
				Case:          1,
				Token:         yyS[yypt-1].Token,
				IdentifierOpt: yyS[yypt-0].node.(*IdentifierOpt),
			}
		}
	case 206:
		{
			yyVAL.node = &StatementNonDecl{
				Case:       2,
				Token:      yyS[yypt-1].Token,
				Expression: yyS[yypt-0].node.(*Expression),
			}
		}
	case 207:
		{
			yyVAL.node = &StatementNonDecl{
				Case:  3,
				Token: yyS[yypt-0].Token,
			}
		}
	case 208:
		{
			yyVAL.node = &StatementNonDecl{
				Case:         4,
				ForStatement: yyS[yypt-0].node.(*ForStatement),
			}
		}
	case 209:
		{
			yyVAL.node = &StatementNonDecl{
				Case:       5,
				Token:      yyS[yypt-1].Token,
				Expression: yyS[yypt-0].node.(*Expression),
			}
		}
	case 210:
		{
			yyVAL.node = &StatementNonDecl{
				Case:   6,
				Token:  yyS[yypt-1].Token,
				Token2: yyS[yypt-0].Token,
			}
		}
	case 211:
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
	case 212:
		{
			yyVAL.node = &StatementNonDecl{
				Case:        8,
				IfStatement: yyS[yypt-0].node.(*IfStatement),
			}
		}
	case 213:
		{
			yyVAL.node = &StatementNonDecl{
				Case:              9,
				Token:             yyS[yypt-1].Token,
				ExpressionListOpt: yyS[yypt-0].node.(*ExpressionListOpt),
			}
		}
	case 214:
		{
			yyVAL.node = &StatementNonDecl{
				Case:            10,
				SelectStatement: yyS[yypt-0].node.(*SelectStatement),
			}
		}
	case 215:
		{
			yyVAL.node = &StatementNonDecl{
				Case:            11,
				SimpleStatement: yyS[yypt-0].node.(*SimpleStatement),
			}
		}
	case 216:
		{
			yyVAL.node = &StatementNonDecl{
				Case:            12,
				SwitchStatement: yyS[yypt-0].node.(*SwitchStatement),
			}
		}
	case 217:
		{
			yyVAL.node = (*StringLitOpt)(nil)
		}
	case 218:
		{
			yyVAL.node = &StringLitOpt{
				Token: yyS[yypt-0].Token,
			}
		}
	case 219:
		{
			yyVAL.node = &StructFieldDecl{
				Token:          yyS[yypt-2].Token,
				QualifiedIdent: yyS[yypt-1].node.(*QualifiedIdent),
				StringLitOpt:   yyS[yypt-0].node.(*StringLitOpt),
			}
		}
	case 220:
		{
			yyVAL.node = &StructFieldDecl{
				Case:           1,
				IdentifierList: yyS[yypt-2].node.(*IdentifierList).reverse(),
				Typ:            yyS[yypt-1].node.(*Typ),
				StringLitOpt:   yyS[yypt-0].node.(*StringLitOpt),
			}
		}
	case 221:
		{
			yyVAL.node = &StructFieldDecl{
				Case:           2,
				QualifiedIdent: yyS[yypt-1].node.(*QualifiedIdent),
				StringLitOpt:   yyS[yypt-0].node.(*StringLitOpt),
			}
		}
	case 222:
		{
			yyVAL.node = &StructFieldDecl{
				Case:           3,
				Token:          yyS[yypt-3].Token,
				QualifiedIdent: yyS[yypt-2].node.(*QualifiedIdent),
				Token2:         yyS[yypt-1].Token,
				StringLitOpt:   yyS[yypt-0].node.(*StringLitOpt),
			}
		}
	case 223:
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
	case 224:
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
	case 225:
		{
			yyVAL.node = &StructFieldDeclList{
				StructFieldDecl: yyS[yypt-0].node.(*StructFieldDecl),
			}
		}
	case 226:
		{
			yyVAL.node = &StructFieldDeclList{
				Case:                1,
				StructFieldDeclList: yyS[yypt-2].node.(*StructFieldDeclList),
				Token:               yyS[yypt-1].Token,
				StructFieldDecl:     yyS[yypt-0].node.(*StructFieldDecl),
			}
		}
	case 227:
		{
			yyVAL.node = &StructType{
				Token:  yyS[yypt-2].Token,
				LBrace: yyS[yypt-1].node.(*LBrace),
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
			yyVAL.node = &StructType{
				Case:                1,
				Token:               yyS[yypt-5].Token,
				LBrace:              yyS[yypt-4].node.(*LBrace),
				StructFieldDeclList: yyS[yypt-2].node.(*StructFieldDeclList).reverse(),
				SemicolonOpt:        yyS[yypt-1].node.(*SemicolonOpt),
				Token2:              yyS[yypt-0].Token,
			}
			lx.popScope()
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
			lx := yylex.(*lexer)
			lx.pushScope()
		}
	case 232:
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
	case 233:
		{
			yyVAL.node = &SwitchCase{
				Token:        yyS[yypt-2].Token,
				ArgumentList: yyS[yypt-1].node.(*ArgumentList).reverse(),
				Token2:       yyS[yypt-0].Token,
			}
		}
	case 234:
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
	case 235:
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
	case 236:
		{
			yyVAL.node = &SwitchCase{
				Case:   3,
				Token:  yyS[yypt-1].Token,
				Token2: yyS[yypt-0].Token,
			}
		}
	case 237:
		{
			yyVAL.node = &SwitchCase{
				Case:  4,
				Token: yyS[yypt-1].Token,
			}
		}
	case 238:
		{
			yyVAL.node = &SwitchCase{
				Case:  5,
				Token: yyS[yypt-1].Token,
			}
		}
	case 239:
		{
			lx := yylex.(*lexer)
			yyVAL.node = &SwitchCaseBlock{
				SwitchCase:    yyS[yypt-1].node.(*SwitchCase),
				StatementList: yyS[yypt-0].node.(*StatementList).reverse(),
			}
			lx.popScope() // Implicit block.
		}
	case 240:
		{
			yyVAL.node = &SwitchCaseList{
				SwitchCaseBlock: yyS[yypt-0].node.(*SwitchCaseBlock),
			}
		}
	case 241:
		{
			yyVAL.node = &SwitchCaseList{
				Case:            1,
				SwitchCaseList:  yyS[yypt-1].node.(*SwitchCaseList),
				SwitchCaseBlock: yyS[yypt-0].node.(*SwitchCaseBlock),
			}
		}
	case 242:
		{
			lx := yylex.(*lexer)
			yyVAL.node = &SwitchStatement{
				Token:      yyS[yypt-2].Token,
				IfHeader:   yyS[yypt-1].node.(*IfHeader),
				SwitchBody: yyS[yypt-0].node.(*SwitchBody),
			}
			lx.popScope() // Implicit block.
		}
	case 243:
		{
			yyVAL.node = &TopLevelDecl{
				ConstDecl: yyS[yypt-0].node.(*ConstDecl),
			}
		}
	case 244:
		{
			yyVAL.node = &TopLevelDecl{
				Case:     1,
				FuncDecl: yyS[yypt-0].node.(*FuncDecl),
			}
		}
	case 245:
		{
			yyVAL.node = &TopLevelDecl{
				Case:     2,
				TypeDecl: yyS[yypt-0].node.(*TypeDecl),
			}
		}
	case 246:
		{
			yyVAL.node = &TopLevelDecl{
				Case:    3,
				VarDecl: yyS[yypt-0].node.(*VarDecl),
			}
		}
	case 247:
		{
			yyVAL.node = &TopLevelDecl{
				Case:             4,
				StatementNonDecl: yyS[yypt-0].node.(*StatementNonDecl),
			}
		}
	case 248:
		{
			yyVAL.node = &TopLevelDecl{
				Case: 5,
			}
		}
	case 249:
		{
			yyVAL.node = (*TopLevelDeclList)(nil)
		}
	case 250:
		{
			yyVAL.node = &TopLevelDeclList{
				TopLevelDeclList: yyS[yypt-2].node.(*TopLevelDeclList),
				TopLevelDecl:     yyS[yypt-1].node.(*TopLevelDecl),
				Token:            yyS[yypt-0].Token,
			}
		}
	case 251:
		{
			yyVAL.node = &Typ{
				Token:  yyS[yypt-2].Token,
				Typ:    yyS[yypt-1].node.(*Typ),
				Token2: yyS[yypt-0].Token,
			}
		}
	case 252:
		{
			yyVAL.node = &Typ{
				Case:  1,
				Token: yyS[yypt-1].Token,
				Typ:   yyS[yypt-0].node.(*Typ),
			}
		}
	case 253:
		{
			yyVAL.node = &Typ{
				Case:      2,
				ArrayType: yyS[yypt-0].node.(*ArrayType),
			}
		}
	case 254:
		{
			yyVAL.node = &Typ{
				Case:     3,
				ChanType: yyS[yypt-0].node.(*ChanType),
			}
		}
	case 255:
		{
			lx := yylex.(*lexer)
			yyVAL.node = &Typ{
				Case:     4,
				FuncType: yyS[yypt-0].node.(*FuncType),
			}
			lx.popScope()
		}
	case 256:
		{
			yyVAL.node = &Typ{
				Case:          5,
				InterfaceType: yyS[yypt-0].node.(*InterfaceType),
			}
		}
	case 257:
		{
			yyVAL.node = &Typ{
				Case:    6,
				MapType: yyS[yypt-0].node.(*MapType),
			}
		}
	case 258:
		{
			yyVAL.node = &Typ{
				Case:                7,
				QualifiedIdent:      yyS[yypt-1].node.(*QualifiedIdent),
				GenericArgumentsOpt: yyS[yypt-0].node.(*GenericArgumentsOpt),
			}
		}
	case 259:
		{
			yyVAL.node = &Typ{
				Case:      8,
				SliceType: yyS[yypt-0].node.(*SliceType),
			}
		}
	case 260:
		{
			yyVAL.node = &Typ{
				Case:       9,
				StructType: yyS[yypt-0].node.(*StructType),
			}
		}
	case 261:
		{
			yyVAL.node = &TypeDecl{
				Token:  yyS[yypt-2].Token,
				Token2: yyS[yypt-1].Token,
				Token3: yyS[yypt-0].Token,
			}
		}
	case 262:
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
	case 263:
		{
			yyVAL.node = &TypeDecl{
				Case:     2,
				Token:    yyS[yypt-1].Token,
				TypeSpec: yyS[yypt-0].node.(*TypeSpec),
			}
		}
	case 264:
		{
			yyVAL.node = &TypeLiteral{
				Token:       yyS[yypt-1].Token,
				TypeLiteral: yyS[yypt-0].node.(*TypeLiteral),
			}
		}
	case 265:
		{
			yyVAL.node = &TypeLiteral{
				Case:      1,
				ArrayType: yyS[yypt-0].node.(*ArrayType),
			}
		}
	case 266:
		{
			yyVAL.node = &TypeLiteral{
				Case:     2,
				ChanType: yyS[yypt-0].node.(*ChanType),
			}
		}
	case 267:
		{
			lx := yylex.(*lexer)
			yyVAL.node = &TypeLiteral{
				Case:     3,
				FuncType: yyS[yypt-0].node.(*FuncType),
			}
			lx.popScope()
		}
	case 268:
		{
			yyVAL.node = &TypeLiteral{
				Case:          4,
				InterfaceType: yyS[yypt-0].node.(*InterfaceType),
			}
		}
	case 269:
		{
			yyVAL.node = &TypeLiteral{
				Case:    5,
				MapType: yyS[yypt-0].node.(*MapType),
			}
		}
	case 270:
		{
			yyVAL.node = &TypeLiteral{
				Case:      6,
				SliceType: yyS[yypt-0].node.(*SliceType),
			}
		}
	case 271:
		{
			yyVAL.node = &TypeLiteral{
				Case:       7,
				StructType: yyS[yypt-0].node.(*StructType),
			}
		}
	case 272:
		{
			lx := yylex.(*lexer)
			lhs := &TypeSpec{
				Token:                yyS[yypt-2].Token,
				GenericParametersOpt: yyS[yypt-1].node.(*GenericParametersOpt),
				Typ:                  yyS[yypt-0].node.(*Typ),
			}
			yyVAL.node = lhs
			lx.scope.declare(lx, newTypeDeclaration(lhs.Token))
		}
	case 273:
		{
			yyVAL.node = &TypeSpecList{
				TypeSpec: yyS[yypt-0].node.(*TypeSpec),
			}
		}
	case 274:
		{
			yyVAL.node = &TypeSpecList{
				Case:         1,
				TypeSpecList: yyS[yypt-2].node.(*TypeSpecList),
				Token:        yyS[yypt-1].Token,
				TypeSpec:     yyS[yypt-0].node.(*TypeSpec),
			}
		}
	case 275:
		{
			yyVAL.node = &UnaryExpression{
				Token:           yyS[yypt-1].Token,
				UnaryExpression: yyS[yypt-0].node.(*UnaryExpression),
			}
		}
	case 276:
		{
			yyVAL.node = &UnaryExpression{
				Case:            1,
				Token:           yyS[yypt-1].Token,
				UnaryExpression: yyS[yypt-0].node.(*UnaryExpression),
			}
		}
	case 277:
		{
			yyVAL.node = &UnaryExpression{
				Case:            2,
				Token:           yyS[yypt-1].Token,
				UnaryExpression: yyS[yypt-0].node.(*UnaryExpression),
			}
		}
	case 278:
		{
			yyVAL.node = &UnaryExpression{
				Case:            3,
				Token:           yyS[yypt-1].Token,
				UnaryExpression: yyS[yypt-0].node.(*UnaryExpression),
			}
		}
	case 279:
		{
			yyVAL.node = &UnaryExpression{
				Case:            4,
				Token:           yyS[yypt-1].Token,
				UnaryExpression: yyS[yypt-0].node.(*UnaryExpression),
			}
		}
	case 280:
		{
			yyVAL.node = &UnaryExpression{
				Case:            5,
				Token:           yyS[yypt-1].Token,
				UnaryExpression: yyS[yypt-0].node.(*UnaryExpression),
			}
		}
	case 281:
		{
			yyVAL.node = &UnaryExpression{
				Case:            6,
				Token:           yyS[yypt-1].Token,
				UnaryExpression: yyS[yypt-0].node.(*UnaryExpression),
			}
		}
	case 282:
		{
			yyVAL.node = &UnaryExpression{
				Case:              7,
				PrimaryExpression: yyS[yypt-0].node.(*PrimaryExpression),
			}
		}
	case 283:
		{
			yyVAL.node = &VarDecl{
				Token:  yyS[yypt-2].Token,
				Token2: yyS[yypt-1].Token,
				Token3: yyS[yypt-0].Token,
			}
		}
	case 284:
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
	case 285:
		{
			yyVAL.node = &VarDecl{
				Case:    2,
				Token:   yyS[yypt-1].Token,
				VarSpec: yyS[yypt-0].node.(*VarSpec),
			}
		}
	case 286:
		{
			lx := yylex.(*lexer)
			lhs := &VarSpec{
				IdentifierList: yyS[yypt-2].node.(*IdentifierList).reverse(),
				Token:          yyS[yypt-1].Token,
				ExpressionList: yyS[yypt-0].node.(*ExpressionList).reverse(),
			}
			yyVAL.node = lhs
			lhs.decl(lx, nil, lhs.ExpressionList)
		}
	case 287:
		{
			lx := yylex.(*lexer)
			lhs := &VarSpec{
				Case:           1,
				IdentifierList: yyS[yypt-1].node.(*IdentifierList).reverse(),
				Typ:            yyS[yypt-0].node.(*Typ),
			}
			yyVAL.node = lhs
			lhs.decl(lx, lhs.Typ, nil)
		}
	case 288:
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
			lhs.decl(lx, lhs.Typ, lhs.ExpressionList)
		}
	case 289:
		{
			yyVAL.node = &VarSpecList{
				VarSpec: yyS[yypt-0].node.(*VarSpec),
			}
		}
	case 290:
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
