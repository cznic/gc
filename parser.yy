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
	"math/big"
	"strconv"
	"strings"
	"unicode"

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
	FuncBodyOpt		"optional function body"
	FuncDecl		"function/method declaration"
	FuncType		"function type"
	GenericArgumentsOpt	"optional generic arguments"
	GenericParametersOpt	"optional generic parameters"
	IdentifierList		"identifier list"
	IdentifierOpt		"optional identifier"
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
	StructFieldDecl		"struct field declaration"
	StructFieldDeclList	"struct field declaration list"
	StructType		"struct type"
	SwitchBody		"body of the switch statement"
	SwitchCase		"switch case/default clause"
	SwitchCaseBlock		"switch case/default clause statement block"
	SwitchCaseList		"switch case/default clause list"
	SwitchHeader		"switch statement header"
	SwitchStatement		"switch statement"
	TagOpt			"optional tag"
	TopLevelDecl		"top level declaration"
	TopLevelDeclList	"top level declaration list"
	Typ			"type "
	TypeDecl		"type declaration"
	TypeList		"type list"
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

//yy:field	Value	Value
Argument:
	Expression
|       TypeLiteral

ArgumentList:
	Argument
|       ArgumentList ',' Argument

//yy:field	guard	gate
//yy:field	Type	Type
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

//yy:field	Value	Value
//yy:field	stringValue	stringValue
BasicLiteral:
	CHAR_LIT
	{
		t := lhs.Token
		s := string(t.S())
		s2 := s[1:len(s)-1]
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
|       FLOAT_LIT
	{
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
|       IMAG_LIT
	{
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
|       INT_LIT
	{
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
|       STRING_LIT
	{
		lhs.stringValue = stringLiteralValue(lx, lhs.Token)
		lhs.Value = newConstValue(newStringConst(lhs.stringValue, lx.stringType, true))
	}

Block:
	'{'
	{
		if !lx.scope.isMergeScope {
			lx.pushScope()
		}
		lx.scope.isMergeScope = false
	}
	StatementList '}'
	{
		lx.popScope()
	}

Body:
	BODY
	{
		lx.pushScope()
	}
	StatementList '}'
	{
		lx.popScope()
	}


Call:
	'(' ')'
|       '(' ArgumentList CommaOpt ')'
|       '(' ArgumentList "..." CommaOpt ')'

//yy:field	guard	gate
//yy:field	Type	Type
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

//yy:field	guard	gate
//yy:field	Type	Type
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
	{
		lhs.decl(lx, nil, nil)
	}
|       IdentifierList '=' ExpressionList
	{
		lhs.decl(lx, nil, lhs.ExpressionList)
	}
|       IdentifierList Typ '=' ExpressionList
	{
		lhs.decl(lx, lhs.Typ, lhs.ExpressionList)
	}

ConstSpecList:
	ConstSpec
|       ConstSpecList ';' ConstSpec

Elif:
	"else" "if" IfHeader Body
	{
		lx.popScope() // Implicit "if" block.
	}

ElifList:
	/* empty */
|       ElifList Elif

ElseOpt:
	/* empty */
|       "else" Block

//yy:field	Value	Value
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
	"for" ForHeader Body
	{
		lx.popScope() // Implicit "for" block.
	}

FuncBodyOpt:
	/* empty */
	{
		lx.popScope()
	}
|        Block

FuncDecl:
	"func" ReceiverOpt IDENTIFIER GenericParametersOpt Signature
	{
		p := lx.pkg
		r := $2.(*ReceiverOpt)
		if r != nil {
			r.Parameters.post(lx)
		}
		sig := $5.(*Signature)
		sig.post(lx)
		nm := $3
		if r == nil {// Function.
			if nm.Val == idInit {
				s := $5.(*Signature)
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
	FuncBodyOpt

/*yy:example "package a ; var b func()" */
//yy:field	guard	gate
//yy:field	Type	Type
FuncType:
	"func" Signature
	{
		lhs.Signature.post(lx)
	}
//yy:example "package a ; var b func « b » ( )"
|	"func" "«" IdentifierList "»" Signature
	{
		lx.err(lhs.Token2, "anonymous functions cannot have generic type parameters")
		lhs.Signature.post(lx)
	}

GenericArgumentsOpt:
	/* empty */
|	"«" TypeList "»"

GenericParametersOpt:
	/* empty */
|	"«" IdentifierList "»"

IdentifierOpt:
	/* empty */
|       IDENTIFIER

IdentifierList:
	IDENTIFIER
|       IdentifierList ',' IDENTIFIER

/*yy:example "package a ; if b {" */
IfHeader:
	SimpleStatementOpt
|       SimpleStatementOpt ';' SimpleStatementOpt

IfStatement:
	"if" IfHeader Body ElifList ElseOpt
	{
		lx.popScope() // Implicit "if" block.
	}

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

//yy:field	guard	gate
//yy:field	methods	*Scope
//yy:field	Type	Type
InterfaceType:
	"interface" LBrace '}'
|       "interface" LBrace
	{
		s := lx.resolutionScope
		lx.pushScope()
		lx.resolutionScope = s
	}
	InterfaceMethodDeclList SemicolonOpt '}'
	{
		lhs.methods = lx.scope
		lx.popScope()
	}

InterfaceMethodDecl:
	IDENTIFIER
	{
		s := lx.resolutionScope
		lx.pushScope()
		lx.resolutionScope = s
	}
	Signature
	{
		lhs.Signature.post(lx)
		s := lx.resolutionScope
		lx.popScope()
		lx.resolutionScope = s
		lx.scope.declare(lx, newFuncDeclaration(lhs.Token, nil, lhs.Signature, true))
	}
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

//yy:field	guard	gate
//yy:field	Type	Type
MapType:
	"map" '[' Typ ']' Typ

//yy:field	Value		Value
//yy:field	fileScope	*Scope
//yy:field	resolutionScope	*Scope	// Where to search for case 4: IDENTIFIER.
Operand:
	'(' Expression ')'
|       '(' TypeLiteral ')'
|       BasicLiteral
|       FuncType
	{
		lx.scope.isMergeScope = false
	}
	LBrace StatementList '}'
	{
		lx.popScope()
	}
|       IDENTIFIER GenericArgumentsOpt
	{
		lhs.fileScope = lx.fileScope
		lhs.resolutionScope = lx.resolutionScope
	}

PackageClause:
	"package" IDENTIFIER ';'
	{
		if lx.pkg.parseOnlyName {
			lx.pkg.Name = string(lhs.Token2.S())
			return 0
		}

		if !lx.build { // Build tags not satisfied
			return 0
		}

		lhs.post(lx)
	}

//yy:field	isParamName	bool
//yy:field	isVariadic	bool
//yy:field	nm		xc.Token
//yy:field	typ		*Typ
ParameterDecl:
	"..." Typ
|       IDENTIFIER "..." Typ
|       IDENTIFIER Typ
|       Typ %prec TYPE

ParameterDeclList:
	ParameterDecl
|       ParameterDeclList ',' ParameterDecl

//yy:field	list	[]*ParameterDecl
Parameters:
	'(' ')'
|       '(' ParameterDeclList CommaOpt ')'

//yy:field	Value	Value
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

QualifiedIdent:
	IDENTIFIER
|       IDENTIFIER '.' IDENTIFIER

Range:
	ExpressionList '=' "range" Expression
|       ExpressionList ":=" "range" Expression
	{
		varDecl(lx, lhs.ExpressionList, lhs.Expression, nil, ":=", 2, 1)
	}
|       "range" Expression

//yy:field	Type		Type
//yy:field	isPtr		bool
//yy:field	nm		xc.Token
//yy:field	resolutionScope	*Scope
ReceiverOpt:
	/* empty */
|       Parameters %prec PARAMS
	{
		lhs.resolutionScope = lx.resolutionScope
		lhs.post(lx)
	}

//yy:field	Type	Type
ResultOpt:
	/* empty */ %prec NO_RESULT
|       Parameters
/*yy:example "package a ; func ( ) []b (" */
|       Typ      

SelectStatement:
	"select"
	{
		lx.switchDecl = append(lx.switchDecl, xc.Token{})
	}
	SwitchBody
	{
		lx.switchDecl = lx.switchDecl[:len(lx.switchDecl)-1]
	}

SemicolonOpt:
	/* empty */
|       ';'

/*yy:example "package a ; var b func ( )" */
//yy:field	Type	Type
Signature:
	Parameters ResultOpt

SimpleStatement:
	Assignment
|       Expression
|       Expression "--"
|       Expression "++"
|       ExpressionList ":=" ExpressionList
	{
		varDecl(lx, lhs.ExpressionList, lhs.ExpressionList2, nil, ":=", -1, -1)
	}

SimpleStatementOpt:
	/* empty */
|       SimpleStatement

//yy:field	guard	gate
//yy:field	Type	Type
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
	{
		lx.scope.declare(lx, newLabelDeclaration(lhs.Token))
	}
|       IfStatement
|       "return" ExpressionListOpt
|       SelectStatement
|       SimpleStatement
|       SwitchStatement

StructFieldDecl:
	'*' QualifiedIdent TagOpt
	{
		lhs.decl(lx)
	}
|       IdentifierList Typ TagOpt
	{
		lhs.decl(lx)
	}
|       QualifiedIdent TagOpt
	{
		lhs.decl(lx)
	}
|       '(' QualifiedIdent ')' TagOpt
	{
		lhs.decl(lx)
	}
|       '(' '*' QualifiedIdent ')' TagOpt
	{
		lhs.decl(lx)
	}
|       '*' '(' QualifiedIdent ')' TagOpt
	{
		lhs.decl(lx)
	}

StructFieldDeclList:
	StructFieldDecl
|       StructFieldDeclList ';' StructFieldDecl

//yy:field	fields	*Scope
//yy:field	guard	gate
//yy:field	Type	Type
StructType:
	"struct" LBrace '}'
|       "struct" LBrace
	{
		lx.pushScope()
		lx.resolutionScope = lx.scope.Parent
	}
	StructFieldDeclList SemicolonOpt '}'
	{
		lhs.fields = lx.scope
		lx.popScope()
	}

SwitchBody:
	BODY '}'
|       BODY
	{
		lx.pushScope()
	}
	SwitchCaseList '}'
	{
		lx.popScope()
	}

SwitchCase:
	"case" ArgumentList ':'
|       "case" ArgumentList '=' Expression ':'
|       "case" ArgumentList ":=" Expression ':'
	{
		varDecl(lx, lhs.ArgumentList, lhs.Expression, nil, ":=", 2, 1)
	}
|       "default" ':'	
|	"case" error
|	"default" error

SwitchCaseBlock:
	SwitchCase StatementList
	{
		lx.popScope() // Implicit "case"/"default" block.
	}

SwitchCaseList:
	SwitchCaseBlock
|       SwitchCaseList SwitchCaseBlock

/*yy:example "package a ; switch b {" */
SwitchHeader:
	SimpleStatementOpt
|       SimpleStatementOpt ';'
|       SimpleStatementOpt ';' Expression
|       SimpleStatementOpt ';' IDENTIFIER ":=" PrimaryExpression '.' '(' "type" ')'

SwitchStatement:
	"switch" SwitchHeader
	{
		var t xc.Token
		if sh := $2.(*SwitchHeader); sh.Case == 3 { // SimpleStatementOpt ';' IDENTIFIER ":=" PrimaryExpression '.' '(' "type" ')'
			t = sh.Token2 // IDENTIFIER
		}
		lx.switchDecl = append(lx.switchDecl, t)
	}
	SwitchBody
	{
		lx.switchDecl = lx.switchDecl[:len(lx.switchDecl)-1]
		lx.popScope() // Implicit "switch" block.
	}

//yy:field	stringValue	stringValue
TagOpt:
	/* empty */
|       STRING_LIT
	{
		lhs.stringValue = stringLiteralValue(lx, lhs.Token)
	}

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

//yy:field	Type		Type
//yy:field	fileScope	*Scope
//yy:field	guard		gate
//yy:field	resolutionScope	*Scope	// Where to search for case 7: QualifiedIdent.
Typ:
	'(' Typ ')'
|       '*' Typ
|       ArrayType
|       ChanType
/*yy:example "package a ; var b func ( )" */
|       FuncType
	{
		lx.popScope()
	}
|       InterfaceType
|       MapType
|       QualifiedIdent GenericArgumentsOpt
	{
		lhs.fileScope = lx.fileScope
		lhs.resolutionScope = lx.resolutionScope
	}
|       SliceType
|       StructType

TypeDecl:
	"type" '(' ')'
|       "type" '(' TypeSpecList SemicolonOpt ')'
|       "type" TypeSpec

TypeList:
	Typ
|	TypeList ',' Typ

//yy:field	guard	gate
//yy:field	Type	Type
TypeLiteral:
	'*' TypeLiteral
|       ArrayType
|       ChanType
/*yy:example "package a ; b(func())" */
|       FuncType
	{
		lx.popScope()
	}
|       InterfaceType
|       MapType
|       SliceType
|       StructType

TypeSpec:
	IDENTIFIER GenericParametersOpt Typ
	{
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

TypeSpecList:
	TypeSpec
|       TypeSpecList ';' TypeSpec

//yy:field	Value	Value
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
	{
		varDecl(lx, lhs.IdentifierList, lhs.ExpressionList, nil, "=", -1, -1)
	}
|       IdentifierList Typ
	{
		varDecl(lx, lhs.IdentifierList, nil, lhs.Typ, "", -1, -1)
	}
|       IdentifierList Typ '=' ExpressionList
	{
		varDecl(lx, lhs.IdentifierList, lhs.ExpressionList, lhs.Typ, "=", -1, -1)
	}

VarSpecList:
	VarSpec
|       VarSpecList ';' VarSpec
