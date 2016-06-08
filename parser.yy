%{
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

import (
	"strconv"
	"strings"

	"github.com/cznic/xc"
)
%}

%union {
	node    Node
	Token   xc.Token
}

%token
	/*yy:token "'%c'"   */  CHAR_LIT        "literal"
	/*yy:token "1.%d"   */  FLOAT_LIT       "literal "
	/*yy:token "%c"     */  IDENTIFIER      "identifier"
	/*yy:token "%di"    */  IMAG_LIT        "literal  "
	/*yy:token "%d"     */  INT_LIT         "literal   "
	/*yy:token "\"%c\"" */  STRING_LIT      "literal    "

%token
	ADD_ASSIGN      "+="
	ANDAND          "&&"
	ANDNOT          "&^"
	ANDNOT_ASSIGN   "&^="
	AND_ASSIGN      "&="
	BAD_FLOAT_LIT	"0x1p10"
	BODY            "{"
	BREAK           "break"
	CASE            "case"
	CHAN            "chan"
	COLAS           ":="
	COMM            "<-"
	CONST           "const"
	CONTINUE        "continue"
	DDD             "..."
	DEC             "--"
	DEFAULT         "default"
	DEFER           "defer"
	DIV_ASSIGN      "/="
	ELSE            "else"
	EQ              "=="
	ERRCHECK	"// ERROR"
	FALLTHROUGH     "fallthrough"
	FOR             "for"
	FUNC            "func"
	GEQ             ">="
	GO              "go"
	GOTO            "goto"
	GTGT		"»"
	IF              "if"
	IMPORT          "import"
	INC             "++"
	INTERFACE       "interface"
	LEQ             "<="
	LSH             "<<"
	LSH_ASSIGN      "<<="
	LTLT		"«"
	MAP             "map"
	MOD_ASSIGN      "%="
	MUL_ASSIGN      "*="
	NEQ             "!="
	OROR            "||"
	OR_ASSIGN       "|="
	PACKAGE         "package"
	RANGE           "range"
	RETURN          "return"
	RSH             ">>"
	RSH_ASSIGN      ">>="
	RXCHAN          "<-  "
	SELECT          "select"
	STRUCT          "struct"
	SUB_ASSIGN      "-="
	SWITCH          "switch"
	TXCHAN          "<- "
	TYPE            "type"
	VAR             "var"
	XOR_ASSIGN      "^="

%type
	Argument		"expression/type literal"
	ArgumentList		"argument list"
	ArrayType		"array type"
	Assignment		"assignment"
	BasicLiteral		"literal      "
	Block			"block statement"
	BlockOpt		"optional block"
	Body			"block statement "
	Call			"call"
	ChanType		"channel type"
	CommaOpt		"optional comma"
	CompLitItem		"composite literal item"
	CompLitItemList		"composite literal item list"
	CompLitType		"composite literal type"
	CompLitValue		"composite literal value"
	ConstDecl		"constant declaration"
	ConstSpec		"constant specification"
	ConstSpecList		"constant specification list"
	Elif			"else if clause"
	ElifList		"else if list clause"
	ElseOpt			"optional else clause"
	Expression		"expression"
	ExpressionList		"expression list"
	ExpressionListOpt	"optional expression list"
	ExpressionOpt		"optional expression"
	File			"source file"
	ForHeader		"for statement header"
	ForStatement		"for statement"
	FuncDecl		"function declaration"
	FuncOrMethod		"function/method declaration"
	FuncType		"function type"
	Function		"function"
	GenericArgumentList	"generic argument list"
	GenericArgumentListItem	"generic argument list item"
	GenericArguments	"generic arguments"
	GenericParameterList	"generic parameter list"
	GenericParameterListItem "generic parameter list item"
	GenericParams		"generic parameters"
	GenericParamsOpt	"optional generic parameters"
	IdentifierList		"identifier list"
	IdentifierOpt		"optional identifier"
	If			"if "
	IfHeader		"if statement header"
	IfStatement		"if statement"
	ImportDecl		"import declaration"
	ImportList		"import declaration list"
	ImportSpec		"import specification"
	ImportSpecList		"import specification list"
	InterfaceMethodDecl	"interface method declaration"
	InterfaceMethodDeclList	"interface method declaration list"
	InterfaceType		"interface type"
	LBrace			"left brace"
	LBraceCompLitItem	"composite literal item "
	LBraceCompLitItemList	"composite literal item list "
	LBraceCompLitValue	"composite literal value "
	MapType			"map type"
	Operand			"operand"
	PackageClause		"package clause"
	ParameterDecl		"parameter declaration"
	ParameterDeclList	"parameter declaration list"
	Parameters		"parameters"
	PrimaryExpression	"primary expression"
	Prologue		"package clause and import list"
	QualifiedIdent		"qualified identifier"
	Range			"range clause"
	ReceiverOpt		"optional receiver"
	ResultOpt		"optional result"
	SelectStatement		"select statement"
	SemicolonOpt		"optional semicolon"
	Signature		"function/method signature"
	SimpleStatement		"simple statement"
	SimpleStatementOpt	"optional simple statement"
	SliceType		"slice type"
	Statement		"statement"
	StatementList		"statement list"
	StatementNonDecl	"non declarative statement"
	StringLitOpt		"optional string literal"
	StructFieldDecl		"struct field declaration"
	StructFieldDeclList	"struct field declaration list"
	StructType		"struct type"
	SwitchBody		"body of the switch statement"
	SwitchCase		"switch case/default clause"
	SwitchCaseBlock		"switch case/default clause statement block"
	SwitchCaseList		"switch case/default clause list"
	SwitchStatement		"switch statement"
	TopLevelDecl		"top level declaration"
	TopLevelDeclList	"top level declaration list"
	Typ			"type "
	TypeDecl		"type declaration"
	TypeLiteral		"type literal"
	TypeSpec		"type specification"
	TypeSpecList		"type specification list"
	UnaryExpression		"unary expression"
	VarDecl			"variable declaration"
	VarSpec			"variable specification"
	VarSpecList		"variable specification list"

%left   COMM

%left   OROR
%left   ANDAND
%left   EQ NEQ '<' LEQ '>' GEQ
%left   '+' '-' '|' '^'
%left   '*' '/' '%' LSH RSH '&' ANDNOT

%precedence     NO_RESULT
%precedence     '('

%precedence     TYPE
%precedence     ')'

%precedence	IDENTIFIER
%precedence	PARAMS

%start  File

%%

//yy:field	DotImports	[]*ImportDeclaration	// import . "foo"
//yy:field	UnboundImports	[]*ImportDeclaration	// import _ "foo"
//yy:field	Scope		*Scope			// import "foo" and import foo "bar"
//yy:field	Path		string			// The source file path.
File:
	Prologue TopLevelDeclList
	{
		lhs.DotImports = lx.dotImports
		lhs.Path = lx.name
		lhs.Scope = lx.fileScope
		lhs.UnboundImports = lx.unboundImports
		lx.pkg.Files = append(lx.pkg.Files, lhs)
	}

Argument:
	Expression
|       TypeLiteral

ArgumentList:
	Argument
|       ArgumentList ',' Argument

ArrayType:
	'[' "..." ']' Typ
|       '[' Expression ']' Typ

Assignment:
	ExpressionList '='   ExpressionList
|       ExpressionList "+="  ExpressionList
|       ExpressionList "&^=" ExpressionList
|       ExpressionList "&="  ExpressionList
|       ExpressionList "/="  ExpressionList
|       ExpressionList "<<=" ExpressionList
|       ExpressionList "%="  ExpressionList
|       ExpressionList "*="  ExpressionList
|       ExpressionList "|="  ExpressionList
|       ExpressionList ">>=" ExpressionList
|       ExpressionList "-="  ExpressionList
|       ExpressionList "^="  ExpressionList

//yy:field	val	interface{}	//TODO -> Value
BasicLiteral:
	CHAR_LIT
|       FLOAT_LIT
|       IMAG_LIT
|       INT_LIT
|       STRING_LIT
	{
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

Block:
	'{' StatementList '}'

BlockOpt:
	/* empty */
|        Block

Body:
	BODY StatementList '}'


Call:
	'(' ')'
|       '(' ArgumentList CommaOpt ')'
|       '(' ArgumentList "..." CommaOpt ')'

ChanType:
	"chan" Typ
|       "chan" TXCHAN Typ
|       RXCHAN "chan" Typ

CommaOpt:
	/* empty */
|       ','

CompLitItem:
	CompLitValue
|       CompLitValue ':' CompLitValue
|       CompLitValue ':' Expression
|       Expression
|       Expression ':' CompLitValue
|       Expression ':' Expression

CompLitItemList:
	CompLitItem
|       CompLitItemList ',' CompLitItem

CompLitType:
	ArrayType
|       MapType
|       SliceType
|       StructType

CompLitValue:
	'{' '}'
|       '{' CompLitItemList CommaOpt '}'

ConstDecl:
	"const" '(' ')'
|       "const" '(' ConstSpecList SemicolonOpt ')'
|       "const" ConstSpec

ConstSpec:
	IdentifierList
|       IdentifierList '=' ExpressionList
|       IdentifierList Typ '=' ExpressionList

ConstSpecList:
	ConstSpec
|       ConstSpecList ';' ConstSpec

Elif:
	"else" If IfHeader Body

ElifList:
	/* empty */
|       ElifList Elif

ElseOpt:
	/* empty */
|       "else" Block

Expression:
	UnaryExpression
|       Expression '%' Expression
|       Expression '&' Expression
|       Expression '*' Expression
|       Expression '+' Expression
|       Expression '-' Expression
|       Expression '/' Expression
|       Expression '<' Expression
|       Expression '>' Expression
|       Expression '^' Expression
|       Expression '|' Expression
|       Expression "&&" Expression
|       Expression "&^" Expression
|       Expression "==" Expression
|       Expression ">=" Expression
|       Expression "<=" Expression
|       Expression "<<" Expression
|       Expression "!=" Expression
|       Expression "||" Expression
|       Expression ">>" Expression
|       Expression "<-" Expression

ExpressionOpt:
	/* empty */
|       Expression

ExpressionList:
	Expression
|       ExpressionList ',' Expression

ExpressionListOpt:
	/* empty */
|       ExpressionList

ForHeader:
	Range
|       SimpleStatementOpt ';' SimpleStatementOpt ';' SimpleStatementOpt
|       SimpleStatementOpt

ForStatement:
	"for"
	ForHeader Body

FuncDecl:
	FuncOrMethod BlockOpt

FuncOrMethod:
	Function ReceiverOpt IDENTIFIER GenericArgumentsOpt Signature

/*yy:example "package a ; var b func()" */
FuncType:
	Function Signature

Function:
	"func"

GenericArgumentList:
	GenericArgumentListItem
|	GenericArgumentList ',' GenericArgumentListItem

GenericArgumentListItem:
	Typ

GenericArguments:
	"«" GenericArgumentList "»"

GenericArgumentsOpt:
	/* empty */
|	GenericArguments

GenericParameterList:
	GenericParameterListItem
|	GenericParameterList ',' GenericParameterListItem

GenericParameterListItem:
	IDENTIFIER

GenericParams:
	"«" GenericParameterList "»"

GenericParamsOpt:
	/* empty */
|	GenericParams

IdentifierOpt:
	/* empty */
|       IDENTIFIER

IdentifierList:
	IDENTIFIER
|       IdentifierList ',' IDENTIFIER

If:
	"if"

/*yy:example "package a ; switch b {" */
IfHeader:
	SimpleStatementOpt
|       SimpleStatementOpt ';' SimpleStatementOpt

IfStatement:
	If IfHeader Body ElifList ElseOpt

ImportDecl:
	"import" '(' ')'
|       "import" '(' ImportSpecList SemicolonOpt ')'
|       "import" ImportSpec

ImportSpec:
	'.' BasicLiteral
	{
		lhs.post(lx)
	}
|       IdentifierOpt BasicLiteral
	{
		lhs.post(lx)
	}
|       '.' BasicLiteral error
|       IdentifierOpt BasicLiteral error

ImportSpecList:
	ImportSpec
|       ImportSpecList ';' ImportSpec

ImportList:
	/* empty */
|       ImportList ImportDecl ';'

InterfaceType:
	"interface" LBrace '}'
|       "interface" LBrace InterfaceMethodDeclList SemicolonOpt '}'

InterfaceMethodDecl:
	IDENTIFIER
	Signature
|       QualifiedIdent

InterfaceMethodDeclList:
	InterfaceMethodDecl
|       InterfaceMethodDeclList ';' InterfaceMethodDecl

/*yy:example "package a ; if interface { !" */
LBrace:
	BODY 
	{
		lx.fixLBR()
	}
|      '{'

LBraceCompLitItem:
	Expression
|       Expression ':' Expression
|       Expression ':' LBraceCompLitValue
|       LBraceCompLitValue

LBraceCompLitItemList:
	LBraceCompLitItem
|       LBraceCompLitItemList ',' LBraceCompLitItem

LBraceCompLitValue:
	LBrace '}'
|       LBrace LBraceCompLitItemList CommaOpt '}'

MapType:
	"map" '[' Typ ']' Typ

Operand:
	'(' Expression ')'
|       '(' TypeLiteral ')'
|       BasicLiteral
|       FuncType LBrace StatementList '}'
|       IDENTIFIER GenericArgumentsOpt

QualifiedIdent:
	IDENTIFIER
|       IDENTIFIER '.' IDENTIFIER

PackageClause:
	"package"
	{
		if !lx.build { // Build tags not satisfied
			return 0
		}
	}
	IDENTIFIER ';'
	{
		lx.pkgName = lhs.Token2.Val
		if lx.parseOnlyPackageClause {
			return 0
		}

		lhs.post(lx)
	}

ParameterDecl:
	"..." Typ
|       IDENTIFIER "..." Typ
|       IDENTIFIER Typ
|       Typ %prec TYPE

ParameterDeclList:
	ParameterDecl
|       ParameterDeclList ',' ParameterDecl

Parameters:
	'(' ')'
|       '(' ParameterDeclList CommaOpt ')'

PrimaryExpression:
	Operand
|       CompLitType LBraceCompLitValue
|       PrimaryExpression '.' '(' "type" ')'
|       PrimaryExpression '.' '(' Typ ')'
|       PrimaryExpression '.' IDENTIFIER
|       PrimaryExpression '[' Expression ']'
|       PrimaryExpression '[' ExpressionOpt ':' ExpressionOpt ':' ExpressionOpt ']'
|       PrimaryExpression '[' ExpressionOpt ':' ExpressionOpt ']'
|       PrimaryExpression Call
|       PrimaryExpression CompLitValue
|       TypeLiteral '(' Expression CommaOpt ')'

Prologue:
	PackageClause ImportList
	{
		lhs.post(lx)
	}

Range:
	ExpressionList '=' "range" Expression
|       ExpressionList ":=" "range" Expression
|       "range" Expression

ReceiverOpt:
	/* empty */
|       Parameters %prec PARAMS

ResultOpt:
	/* empty */ %prec NO_RESULT
|       Parameters
/*yy:example "package a ; func ( ) []b (" */
|       Typ      

SelectStatement:
	"select" SwitchBody

SemicolonOpt:
	/* empty */
|       ';'

/*yy:example "package a ; var b func ( )" */
Signature:
	Parameters ResultOpt

SimpleStatement:
	Assignment
|       Expression
|       Expression "--"
|       Expression "++"
|       ExpressionList ":=" ExpressionList

SimpleStatementOpt:
	/* empty */
|       SimpleStatement

SliceType:
	'[' ']' Typ

Statement:
	/* empty */
|       Block
|       ConstDecl
|       TypeDecl
|       VarDecl
|       StatementNonDecl
|	error

/*yy:example "package a ; if { b ;" */
StatementList:
	Statement
|       StatementList ';' Statement

StatementNonDecl:
	"break" IdentifierOpt
|       "continue" IdentifierOpt
|       "defer" Expression
|       "fallthrough"
|       ForStatement
|       "go" Expression
|       "goto" IDENTIFIER
|       IDENTIFIER ':' Statement
|       IfStatement
|       "return" ExpressionListOpt
|       SelectStatement
|       SimpleStatement
|       SwitchStatement

StringLitOpt:
	/* empty */
|       STRING_LIT

StructFieldDecl:
	'*' QualifiedIdent StringLitOpt
|       IdentifierList Typ StringLitOpt
|       QualifiedIdent StringLitOpt
|       '(' QualifiedIdent ')' StringLitOpt
|       '(' '*' QualifiedIdent ')' StringLitOpt
|       '*' '(' QualifiedIdent ')' StringLitOpt

StructFieldDeclList:
	StructFieldDecl
|       StructFieldDeclList ';' StructFieldDecl

StructType:
	"struct" LBrace '}'
|       "struct" LBrace StructFieldDeclList SemicolonOpt '}'

SwitchBody:
	BODY '}'
|       BODY SwitchCaseList '}'

SwitchCase:
	"case" ArgumentList ':'
|       "case" ArgumentList '=' Expression ':'
|       "case" ArgumentList ":=" Expression ':'
|       "default" ':'
|	"case" error
|	"default" error

SwitchCaseBlock:
	SwitchCase StatementList

SwitchCaseList:
	SwitchCaseBlock
|       SwitchCaseList SwitchCaseBlock

SwitchStatement:
	"switch" IfHeader SwitchBody

TopLevelDecl:
	ConstDecl
|       FuncDecl
|       TypeDecl
|       VarDecl
|	StatementNonDecl
|	error

TopLevelDeclList:
	/* empty */
|       TopLevelDeclList TopLevelDecl ';'

Typ:
	'(' Typ ')'
|       '*' Typ
|       ArrayType
|       ChanType
/*yy:example "package a ; var b func ( )" */
|       FuncType
|       InterfaceType
|       MapType
|       QualifiedIdent GenericArgumentsOpt
|       SliceType
|       StructType

TypeDecl:
	"type" '(' ')'
|       "type" '(' TypeSpecList SemicolonOpt ')'
|       "type" TypeSpec

TypeLiteral:
	'*' TypeLiteral
|       ArrayType
|       ChanType
/*yy:example "package a ; b(func())" */
|       FuncType
|       InterfaceType
|       MapType
|       SliceType
|       StructType

TypeSpec:
	IDENTIFIER GenericParamsOpt Typ

TypeSpecList:
	TypeSpec
|       TypeSpecList ';' TypeSpec

UnaryExpression:
	'!' UnaryExpression
|       '&' UnaryExpression
|       '*' UnaryExpression
|       '+' UnaryExpression
|       '-' UnaryExpression
|       '^' UnaryExpression
|       "<-" UnaryExpression
|       PrimaryExpression

VarDecl:
	"var" '(' ')'
|       "var" '(' VarSpecList SemicolonOpt ')'
|       "var" VarSpec

VarSpec:
	IdentifierList '=' ExpressionList
|       IdentifierList Typ
|       IdentifierList Typ '=' ExpressionList

VarSpecList:
	VarSpec
|       VarSpecList ';' VarSpec
