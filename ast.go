// Copyright 2016 The GC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//TODO declare labels

//TODO s/n.sourceFile.Package.errorlist.add/n.err/g

package gc

import (
	"bytes"
	"fmt"
	gotoken "go/token"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"unicode"

	"github.com/cznic/token"
)

var (
	_ Node = Declaration(nil)
	//TODO _ Node = Expr(nil)
	//TODO _ Node = Statement(nil)
	_ Node = Type(nil)

	buildMark  = []byte("// +build")
	buildMark2 = []byte("//+build")

	emptyParamTypeList ParamTypeList = &ParamTypeListType{}
	invalidType        Type          = Invalid
)

type Node interface {
	Package() *Package
	Position() token.Position
	SourceFile() *SourceFile
	String() string

	checker
	err(string, ...interface{})
	err0(token.Position, string, ...interface{})
	//TODO? trap()
}

type node struct {
	npos
	sentinel
}

type npos struct {
	sf   *SourceFile
	tpos token.Pos
}

func newPos(t Token) npos { return t.npos } //TODO-

func (n *npos) Package() *Package                 { return n.sf.Package }
func (n *npos) Pos() token.Pos                    { return n.tpos }
func (n *npos) SourceFile() *SourceFile           { return n.sf }
func (n *npos) err(s string, args ...interface{}) { n.err0(n.Position(), s, args...) }
func (n *npos) pos() npos                         { return *n }
func (n *npos) set(p npos)                        { *n = p }

func (n *npos) Position() (r token.Position) {
	if s := n.sf; s != nil && s.File != nil {
		r = s.File.PositionFor(n.tpos, true)
	}
	return r
}

func (p *npos) err0(position token.Position, s string, args ...interface{}) {
	p.Package().errorList.add(position, fmt.Sprintf(s, args...))
}

// Token represents the position and value of a token.
type Token struct {
	npos
	Val string
}

// Ident represents an identifier and its resolution scope
type Ident struct {
	Token
	scope *Scope
}

// Scope returns n's resolution scope.
func (n *Ident) Scope() *Scope { return n.scope }

type parser struct {
	c             gotoken.Token
	l             *lexer
	loophackStack []bool
	off           int
	scope         *Scope
	sourceFile    *SourceFile
	syntaxError   func(*parser)

	loophack          bool
	noSyntaxErrorFunc bool //TODO later delete
	errAtEOF          bool
}

func newParser(src *SourceFile, l *lexer) *parser {
	p := &parser{
		l:          l,
		sourceFile: src,
	}
	return p
}

func (p *parser) declare(d Declaration) { p.scope.declare(p.sourceFile, d) }

func (p *parser) init(src *SourceFile, l *lexer) {
	p.l = l
	p.loophack = false
	p.loophackStack = p.loophackStack[:0]
	p.sourceFile = src
}

func (p *parser) err0(position token.Position, msg string, args ...interface{}) {
	if p.c == gotoken.EOF {
		if p.errAtEOF {
			return
		}

		p.errAtEOF = true
	}
	p.sourceFile.Package.errorList.add(position, fmt.Sprintf(msg, args...))
}

func (p *parser) err(msg string, args ...interface{}) { p.err0(p.position(), msg, args...) }

func (p *parser) n() gotoken.Token {
more:
	switch p.off, p.c = p.l.Scan(); p.c {
	case gotoken.FOR, gotoken.IF, gotoken.SELECT, gotoken.SWITCH:
		p.loophack = true
	case gotoken.LPAREN, gotoken.LBRACK:
		if p.loophack || len(p.loophackStack) != 0 {
			p.loophackStack = append(p.loophackStack, p.loophack)
			p.loophack = false
		}
	case gotoken.RPAREN, gotoken.RBRACK:
		if n := len(p.loophackStack); n != 0 {
			p.loophack = p.loophackStack[n-1]
			p.loophackStack = p.loophackStack[:n-1]
		}
	case gotoken.LBRACE:
		if p.loophack {
			p.c = tokenBODY
			p.loophack = false
		}
	case tokenBOM:
		goto more
	case gotoken.ILLEGAL:
		if p.noSyntaxErrorFunc {
			goto more
		}
	}
	return p.c
}

func (p *parser) ident() Ident { return Ident{p.l.Token(p.off), p.scope} }
func (p *parser) tok() Token   { return p.l.Token(p.off) }

func (p *parser) opt(tok gotoken.Token) bool {
	if p.c == tok {
		p.n()
		return true
	}

	return false
}

func (p *parser) skip(toks ...gotoken.Token) {
	for p.n() != gotoken.EOF {
		for _, v := range toks {
			if p.c == v {
				return
			}
		}
	}
}

func (p *parser) must(tok gotoken.Token) (ok bool) {
	ok = true
	if p.c != tok {
		p.syntaxError(p)
		if p.c != gotoken.EOF {
			p.err("syntax error: unexpected %v, expecting %v", p.unexpected(), tok)
		}
		ok = false
	}
	p.n()
	return ok
}

func (p *parser) mustTok(tok gotoken.Token) (t Token, ok bool) {
	ok = true
	if p.c != tok {
		p.syntaxError(p)
		if p.c != gotoken.EOF {
			p.err("syntax error: unexpected %v, expecting %v", p.unexpected(), tok)
		}
		p.n()
		return t, false
	}

	t = p.tok()
	p.n()
	return t, true
}

func (p *parser) must2(toks ...gotoken.Token) (ok bool) {
	ok = true
	for _, tok := range toks {
		ok = p.must(tok) && ok
	}
	return ok
}

func (p *parser) not2(toks ...gotoken.Token) bool {
	for _, tok := range toks {
		if p.c == tok {
			return false
		}
	}
	return true
}

func (p *parser) unexpected() string {
	lit := p.l.lit
	if len(lit) == 1 && lit[0] == '\n' {
		if p.l.c == classEOF {
			return "EOF"
		}

		return "newline"
	}

	switch p.c {
	case gotoken.IDENT:
		return string(lit)
	case gotoken.INT, gotoken.FLOAT:
		return fmt.Sprintf("literal %s", lit)
	}

	return p.c.String()
}

func (p *parser) pos() token.Pos           { return p.l.file.Pos(p.off) }
func (p *parser) position() token.Position { return p.l.file.Position(p.pos()) }

func (p *parser) strLit(s string) string {
	value, err := strconv.Unquote(s)
	if err != nil {
		p.err("%s: %q", err, s)
		return ""
	}

	// https://github.com/golang/go/issues/15997
	if s[0] == '`' {
		value = strings.Replace(value, "\r", "", -1)
	}
	return value
}

func (p *parser) commentHandler(_ int, lit []byte) {
	if p.sourceFile.build {
		if bytes.HasPrefix(lit, buildMark) {
			p.buildDirective(lit[len(buildMark):])
			return
		}

		if bytes.HasPrefix(lit, buildMark2) {
			p.buildDirective(lit[len(buildMark2):])
			return
		}
	}
}

func (p *parser) buildDirective(b []byte) {
	ctx := p.sourceFile.Package.ctx
	s := string(b)
	s = strings.Replace(s, "\t", " ", -1)
	for _, term := range strings.Split(s, " ") { // term || term
		if term = strings.TrimSpace(term); term == "" {
			continue
		}

		val := true
		for _, factor := range strings.Split(term, ",") { // factor && factor
			if factor = strings.TrimSpace(factor); factor == "" {
				continue
			}

			not := factor[0] == '!'
			if not {
				factor = strings.TrimSpace(factor[1:])
			}

			if factor == "" {
				continue
			}

			_, ok := ctx.tags[factor]
			if not {
				ok = !ok
			}

			if !ok {
				val = false
				break
			}

		}
		if val {
			return
		}
	}

	p.sourceFile.build = false
}

func (p *parser) push() *Scope {
	s := newScope(BlockScope, p.scope)
	p.scope = s
	return s
}

func (p *parser) pop() {
	if p.scope.Kind == PackageScope {
		panic("internal error 013")
	}

	p.scope = p.scope.Parent
}

// ImportSpec implements Declaration.
func (n *ImportDecl) ImportSpec() *ImportDecl { return n }

// Kind implements Declaration.
func (n *ImportDecl) Kind() DeclarationKind { return ImportDeclaration }

// Name implements Declaration.
func (n *ImportDecl) Name() string {
	if n.Qualifier != "" {
		return n.Qualifier
	}

	return n.Package().Name
}

// importSpec:
// 	'.' STRING
//	|	IDENT STRING
//	|	STRING
func (p *parser) importSpec() {
	var decl, qualifier Token
	var dot bool
	switch p.c {
	case gotoken.IDENT:
		qualifier = p.tok()
		decl = qualifier
		p.n()
	case gotoken.PERIOD:
		dot = true
		p.n()
	}
	switch p.c {
	case gotoken.STRING:
		if decl.Val == "" {
			decl = p.tok()
		}
		ip := p.strLit(string(p.l.lit))
		if ip == "C" { //TODO
			p.n()
			return
		}

		if ip == "" {
			p.err("import path is empty")
			p.n()
			return
		}

		if strings.Contains(ip, "\x00") {
			p.err("import path contains NUL")
			p.n()
			return
		}

		if strings.Contains(ip, "\\") {
			p.err("import path contains backslash; use slash: %q", ip)
			p.n()
			return
		}

		if filepath.IsAbs(ip) {
			p.err("import path cannot be absolute path")
			p.n()
			return
		}

		for _, v := range ip {
			if unicode.IsControl(v) {
				p.err("import path contains control character: %q", string(v))
				p.n()
				return
			}

			switch v {
			case ' ':
				p.err("import path contains space character: %q", ip)
				p.n()
				return
			case '!', '"', '#', '$', '%', '&', '\'', '(', ')', '*', ',', ':', ';', '<', '=', '>', '?', '[', '\\', ']', '^', '`', '{', '|', '}', '\ufffd':
				p.err("import path contains invalid character %c: %q", v, ip)
				p.n()
				return
			}

			if unicode.IsLetter(v) || unicode.IsMark(v) || unicode.IsNumber(v) || unicode.IsPunct(v) || unicode.IsSymbol(v) {
				continue
			}

			p.err("import path contains invalid character %c: %q", v, ip)
			p.n()
			return
		}

		if !p.sourceFile.Package.ctx.tweaks.ignoreImports {
			if strings.HasPrefix(ip, "./") {
				dir, err := p.sourceFile.Package.ctx.dirForImportPath(decl.Position(), ip)
				if err == nil {
					ip = filepath.Join(dir, ip)
				}
			}
			spec := newImportDecl(decl, p.off, dot, qualifier.Val, ip)
			spec.pkg = p.sourceFile.Package.ctx.load(p.position(), ip, nil, p.sourceFile.Package.errorList).waitFor()
			p.sourceFile.ImportSpecs = append(p.sourceFile.ImportSpecs, spec)
			switch {
			case dot:
				for _, v := range spec.Package().SourceFiles {
					for _, v := range v.TopLevelDecls {
						if _, ok := v.(*MethodDecl); ok {
							continue
						}

						p.sourceFile.Scope.declare(p.sourceFile, v)
					}
				}
			default:
				if ex, ok := spec.Package().fileScopeNames[spec.Name()]; ok {
					_ = ex
					panic("TODO414") // declared in pkg and file scope at the same time.
					break
				}

				if spec.Name() != "" {
					p.sourceFile.Scope.declare(p.sourceFile, spec)
				}
			}
		}
		p.n()
	default:
		p.syntaxError(p)
		p.err("import path must be a string")
		if p.noSyntaxErrorFunc {
			p.skip(gotoken.SEMICOLON)
		}
	}
}

// importSpecList:
//		importSpec
//	|	importSpecList ';' importSpec
func (p *parser) importSpecList() {
	for p.importSpec(); p.opt(gotoken.SEMICOLON) && p.c != gotoken.RPAREN; {
		p.importSpec()
	}
}

// imports:
//	|	imports "import" '(' ')' ';'
//	|	imports "import" '(' importSpecList semiOpt ')' ';'
//	|	imports "import" importSpec ';'
func (p *parser) imports() {
	for p.opt(gotoken.IMPORT) {
		switch {
		case p.opt(gotoken.LPAREN):
			if !p.opt(gotoken.RPAREN) {
				p.importSpecList()
				if p.c == gotoken.COMMA {
					p.syntaxError(p)
					p.err("syntax error: unexpected comma, expecting semicolon, newline, or )")
					p.skip(gotoken.RPAREN)
				}
				p.must(gotoken.RPAREN)
			}
		default:
			p.importSpec()
		}
		p.must(gotoken.SEMICOLON)
	}
}

// identList:
//		IDENT
//	|	identList ',' IDENT
func (p *parser) identList() (l []Token) {
	switch p.c {
	case gotoken.IDENT:
		l = []Token{p.tok()}
		p.n()
		for p.opt(gotoken.COMMA) && p.c != tokenGTGT {
			switch p.c {
			case gotoken.IDENT:
				l = append(l, p.tok())
				p.n()
			default:
				//TODO p.err()
				p.syntaxError(p)
			}
		}
	default:
		p.syntaxError(p)
		p.err("syntax error: unexpected %v, expecting name", p.c)
	}
	return l
}

// CompLitExpr represents data reduced by productions
//
//	compLitExpr:
//		'{' bracedKeyValList '}'
//	|	expr
type CompLitExpr interface {
	isCompLitExpr()
	str() string
}

type compLitExpr struct{}

func (compLitExpr) isCompLitExpr() {}

// CompLitExprKeyValList is a CompLitExpr.
type CompLitExprKeyValList struct {
	compLitExpr
	List []KeyVal
}

func (n *CompLitExprKeyValList) str() string {
	var a []string
	for _, v := range n.List {
		a = append(a, v.str())
	}
	return fmt.Sprintf("{%s}", strings.Join(a, ", "))
}

// CompLitExprExpr is a CompLitExpr.
type CompLitExprExpr struct {
	compLitExpr
	Expr Expr
}

func (n *CompLitExprExpr) str() string { return n.Expr.str() }

func (p *parser) compLitExpr() (r CompLitExpr) {
	if p.opt(gotoken.LBRACE) {
		r = &CompLitExprKeyValList{List: p.bracedKeyValList()}
		p.must(gotoken.RBRACE)
		return r
	}

	return &CompLitExprExpr{Expr: p.expr()}
}

// KeyVal represents data reduced by productions
//
//	keyVal:
//		compLitExpr
//	|	compLitExpr ':' compLitExpr
type KeyVal interface {
	isKeyVal()
	str() string
}

type keyVal struct{}

func (keyVal) isKeyVal() {}

// KeyValExpr is a KeyVal
type KeyValExpr struct {
	keyVal
	Expr CompLitExpr
}

func (n *KeyValExpr) str() string { return n.Expr.str() }

// KeyValKeyedExpr is a KeyVal
type KeyValKeyedExpr struct {
	keyVal
	Key  CompLitExpr
	Expr CompLitExpr
}

func (n *KeyValKeyedExpr) str() string { return "TODO KeyValKeyedExpr.str()" }

func (p *parser) keyVal() KeyVal {
	e := p.compLitExpr()
	if !p.opt(gotoken.COLON) {
		return &KeyValExpr{Expr: e}
	}

	return &KeyValKeyedExpr{Key: e, Expr: p.compLitExpr()}
}

// keyValList:
//		keyVal
//	|	keyValList ',' keyVal
func (p *parser) keyValList() (r []KeyVal) {
	for r = []KeyVal{p.keyVal()}; p.opt(gotoken.COMMA) && p.c != gotoken.RBRACE; {
		r = append(r, p.keyVal())
	}
	if p.c == gotoken.SEMICOLON {
		p.syntaxError(p)
		p.err("syntax error: unexpected %v, expecting comma or }", p.unexpected())
		p.n()
	}
	return r
}

// bracedKeyValList:
//	|	keyValList commaOpt
func (p *parser) bracedKeyValList() (r []KeyVal) {
	if p.c != gotoken.RBRACE {
		r = p.keyValList()
		p.opt(gotoken.COMMA)
	}
	return r
}

// ExprOrType represents data reduced by productions
//
//	exprOrType:
//		expr
//	|	nonExprType %prec _PreferToRightParen
type ExprOrType interface {
	Position() token.Position
	Type() Type
	checker
	isExprOrType()
	str() string
}

type exprOrType struct {
	sentinel
}

func (e *exprOrType) Position() (r token.Position) { return r }
func (e *exprOrType) Type() Type                   { return invalidType }
func (e *exprOrType) check(s *cstack)              {}
func (e *exprOrType) isExprOrType()                {}
func (e *exprOrType) str() string                  { panic("internal error 028") }

// ExprOrTypeExpr is an ExprOrType.
type ExprOrTypeExpr struct {
	exprOrType
	Expr Expr
}

func (n *ExprOrTypeExpr) str() string { return n.Expr.str() }

// Position implemenst Expr.
func (n *ExprOrTypeExpr) Position() token.Position { return n.Expr.Position() }

func (n *ExprOrTypeExpr) check(s *cstack) { s.check(n.Expr, n) }

// Type implements Expr.
func (n *ExprOrTypeExpr) Type() Type { return n.Expr.Type() }

// ExprOrTypeType is an ExprOrType.
type ExprOrTypeType struct {
	exprOrType
	typ Type
}

func (n *ExprOrTypeType) str() string { return n.typ.String() }

func (n *ExprOrTypeType) check(s *cstack) { s.check(n.typ, n) }

// Type implements Expr.
func (n *ExprOrTypeType) Type() Type { return n.typ }

func (p *parser) exprOrType(fs string) (r ExprOrType) {
	var mul []Token

	defer func() {
		if len(mul) == 0 {
			return
		}

		switch x := r.(type) {
		case *ExprOrTypeExpr:
			for i := 0; i < len(mul); i++ {
				x.Expr = &UnaryExprDeref{npos: x.Expr.pos(), Expr: x.Expr}
			}
		case *ExprOrTypeType:
			for i := len(mul) - 1; i >= 0; i-- {
				x.typ = newPointerType(mul[i].npos, x.Type())
			}
		case *exprOrType: // syntax error
			// nop
		default:
			panic("internal error 014")
		}
	}()

more:
	var fix bool
	switch p.c {
	case gotoken.IDENT, gotoken.INT, gotoken.FLOAT, gotoken.IMAG, gotoken.CHAR, gotoken.STRING,
		gotoken.NOT, gotoken.AND, gotoken.ADD, gotoken.SUB, gotoken.XOR, gotoken.LPAREN:
		return &ExprOrTypeExpr{Expr: p.expr()}
	case gotoken.ARROW:
		tok := p.tok()
		switch p.n() {
		case gotoken.CHAN:
			p.n()
			return &ExprOrTypeType{typ: newChannelType(tok.npos, RxChan, p.typ())}
		default:
			return &ExprOrTypeExpr{Expr: p.expr()}
		}
	case gotoken.MUL:
		mul = append(mul, p.tok())
		p.n()
		goto more
	case gotoken.CHAN, gotoken.INTERFACE, gotoken.MAP, gotoken.STRUCT, gotoken.LBRACK:
		t := p.otherType(p.c)
		tok := p.tok()
		switch p.c {
		case gotoken.LBRACE:
			r = &ExprOrTypeExpr{Expr: p.expr2(p.primaryExpr2(&PrimaryExprCompositeLiteral{npos: newPos(tok), expr: expr{typ: t}}))}
		case gotoken.LPAREN:
			r = &ExprOrTypeExpr{Expr: p.expr2(p.primaryExpr2(&PrimaryExprConversion{npos: newPos(tok), expr: expr{typ: t}}))}
		default:
			r = &ExprOrTypeType{typ: t}
		}
	case gotoken.FUNC:
		t := p.fnType()
		tok := p.tok()
		switch p.c {
		case tokenBODY:
			fix = true
			fallthrough
		case gotoken.LBRACE:
			p.n()
			body := &StmtBlock{s: p.scope, List: p.stmtList()}
			f := &PrimaryExprFuncLiteral{npos: newPos(tok), expr: expr{typ: t}, Body: body}
			p.loophack = fix
			p.must(gotoken.RBRACE)
			p.pop()
			switch {
			case p.c == gotoken.LPAREN:
				r = &ExprOrTypeExpr{Expr: p.expr2(p.primaryExpr2(f))}
			default:
				r = &ExprOrTypeExpr{Expr: f}
			}
		default:
			r = &ExprOrTypeType{typ: t}
			p.pop()
		}
	default:
		r = &exprOrType{}
		p.syntaxError(p)
		s := "syntax error: unexpected %v"
		if fs != "" {
			s = s + ", expecting " + fs
		}
		p.err(s, p.unexpected())

	}
	return r
}

// exprOrTypeList:
//		exprOrType
//	|	exprOrTypeList ',' exprOrType
func (p *parser) exprOrTypeList() (r []ExprOrType) {
	r = []ExprOrType{p.exprOrType("")}
	for p.opt(gotoken.COMMA) && p.not2(gotoken.RPAREN, gotoken.ELLIPSIS) {
		r = append(r, p.exprOrType(") or ..."))
	}
	return r
}

// exprOpt:
//	|	expr
func (p *parser) exprOpt() (r Expr) {
	if p.c == gotoken.COLON || p.c == gotoken.RBRACK {
		return nil
	}

	return p.expr()
}

// PrimaryExprParenExpr is an Expr.
type PrimaryExprParenExpr struct {
	npos
	expr
	Expr Expr
}

func (n *PrimaryExprParenExpr) str() string { return fmt.Sprintf("(%s)", n.Expr.str()) }

func (n *PrimaryExprParenExpr) check(s *cstack) { s.check(n.unparen(), n) }

func (n *PrimaryExprParenExpr) unparen() Expr {
	for {
		switch x := n.Expr.(type) {
		case *PrimaryExprParenExpr:
			n.Expr = x.Expr
		default:
			return n.Expr
		}
	}
}

// primaryExprParenType is an Expr.
type primaryExprParenType struct {
	npos
	expr
}

func (n *primaryExprParenType) str() string { return "TODO primaryExprParenType.str()" }

func (n *primaryExprParenType) check(s *cstack) {
	//TODO
}

// PrimaryExprIdent is an Expr.
type PrimaryExprIdent struct {
	expr
	Ident Ident
}

func (n *PrimaryExprIdent) str() string { return n.Ident.Val }

// Position implemenst Expr.
func (n *PrimaryExprIdent) Position() token.Position { return n.Ident.Position() }

func (n *PrimaryExprIdent) pos() npos { return newPos(n.Ident.Token) }

func (n *PrimaryExprIdent) check(s *cstack) {
	sc := n.Ident.Scope()
	sf := n.Ident.SourceFile()
	if n.Ident.Val == "_" {
		sf.Package.errorList.add(n.Ident.Position(), "cannot use _ as value")
		return
	}

	d := sc.Lookup(sf.Package, sf.Scope, n.Ident.Token)
	if d == nil {
		sf.Package.errorList.add(n.Ident.Position(), fmt.Sprintf("undefined: %s", n.Ident.Val))
		return
	}

	switch x := d.(type) {
	case
		*ConstDecl,
		*FuncDecl,
		*VarDecl:

		s.check(x, n)
		n.typ = x.Type()
	case *TypeDecl:
		s.check(x, n)
		n.typ = newNamedType(n.Ident)
		s.check(n.typ, n)
	case *ImportDecl:
		sf.Package.errorList.add(n.Ident.Position(), fmt.Sprintf("use of package %s without selector", n.Ident.Val))
	default:
		panic("internal error 042")
	}
}

func (n *PrimaryExprIdent) lookup() Declaration {
	sf := n.Ident.SourceFile()
	return n.Ident.Scope().Lookup(sf.Package, sf.Scope, n.Ident.Token)
}

// PrimaryExprFuncLiteral is an Expr.
type PrimaryExprFuncLiteral struct {
	npos
	expr
	Body *StmtBlock
}

func (n *PrimaryExprFuncLiteral) str() string { return "TODO PrimaryExprFuncLiteral.str()" }

func (n *PrimaryExprFuncLiteral) check(s *cstack) {
	//TODO
}

// PrimaryExprConversion is an Expr.
type PrimaryExprConversion struct {
	npos
	expr
	Expr Expr
}

func (n *PrimaryExprConversion) str() string { return "TODO PrimaryExprConversion.str()" }

func (n *PrimaryExprConversion) check(s *cstack) {
	//TODO
}

// PrimaryExprIntLiteral is an Expr.
type PrimaryExprIntLiteral struct {
	expr
	Literal Token
}

func (n *PrimaryExprIntLiteral) str() string { return n.Literal.Val }

// Position implemenst Expr.
func (n *PrimaryExprIntLiteral) Position() token.Position { return n.Literal.Position() }

func (n *PrimaryExprIntLiteral) pos() npos { return newPos(n.Literal) }

func (n *PrimaryExprIntLiteral) check(s *cstack) {
	n.expr.typ = UntypedInt
	//TODO
}

// PrimaryExprFloatLiteral is an Expr.
type PrimaryExprFloatLiteral struct {
	expr
	Literal Token
}

func (n *PrimaryExprFloatLiteral) str() string { return "TODO PrimaryExprFloatLiteral.str()" }

// Position implemenst Expr.
func (n *PrimaryExprFloatLiteral) Position() token.Position { return n.Literal.Position() }

func (n *PrimaryExprFloatLiteral) pos() npos { return newPos(n.Literal) }

func (n *PrimaryExprFloatLiteral) check(s *cstack) {
	//TODO
}

// PrimaryExprImagLiteral is an Expr.
type PrimaryExprImagLiteral struct {
	expr
	Literal Token
}

func (n *PrimaryExprImagLiteral) str() string { return "TODO PrimaryExprImagLiteral.str()" }

// Position implemenst Expr.
func (n *PrimaryExprImagLiteral) Position() token.Position { return n.Literal.Position() }

func (n *PrimaryExprImagLiteral) pos() npos { return newPos(n.Literal) }

func (n *PrimaryExprImagLiteral) check(s *cstack) {
	//TODO
}

// PrimaryExprRuneLiteral is an Expr.
type PrimaryExprRuneLiteral struct {
	expr
	Literal Token
}

func (n *PrimaryExprRuneLiteral) str() string { return "TODO PrimaryExprRuneLiteral.str()" }

// Position implemenst Expr.
func (n *PrimaryExprRuneLiteral) Position() token.Position { return n.Literal.Position() }

func (n *PrimaryExprRuneLiteral) pos() npos { return newPos(n.Literal) }

func (n *PrimaryExprRuneLiteral) check(s *cstack) {
	//TODO
}

// PrimaryExprStringLiteral is an Expr.
type PrimaryExprStringLiteral struct {
	expr
	Literal Token
}

func (n *PrimaryExprStringLiteral) str() string { return n.Literal.Val }

// Position implemenst Expr.
func (n *PrimaryExprStringLiteral) Position() token.Position { return n.Literal.Position() }

func (n *PrimaryExprStringLiteral) pos() npos { return newPos(n.Literal) }

func (n *PrimaryExprStringLiteral) check(s *cstack) {
	n.expr.typ = UntypedString
	//TODO
}

// PrimaryExprCompositeLiteral is an Expr.
type PrimaryExprCompositeLiteral struct {
	npos
	expr
	List []KeyVal
}

func (n *PrimaryExprCompositeLiteral) str() string {
	return fmt.Sprintf("%s literal", n.Type().alias())
}

func (n *PrimaryExprCompositeLiteral) check(s *cstack) {
	s.check(n.Type(), n)
	//TODO
}

// PrimaryExprCall is an Expr.
type PrimaryExprCall struct {
	npos
	expr
	PrimaryExpr Expr
	Args        []ExprOrType

	DDD          bool
	IsConversion bool
}

func (n *PrimaryExprCall) str() string {
	var a []string
	for _, v := range n.Args {
		a = append(a, v.str())
	}
	return fmt.Sprintf("%s(%s)", n.PrimaryExpr.str(), strings.Join(a, ", "))
}

func (n *PrimaryExprCall) check(s *cstack) {
	s.check(n.PrimaryExpr, n)
	t := n.PrimaryExpr.Type()
	//TODO if t != nil && t.state() == 0 {
	//TODO 	panic(fmt.Errorf("%v: TODO1011", n.Position()))
	//TODO }
	if t != nil && t != Invalid && !t.isGeneric(map[Type]bool{}) {
		switch u := t.UnderlyingType(); f := u.(type) {
		case *FunctionType:
			n.expr.typ = f.Result()
			switch x := n.PrimaryExpr.(type) {
			case *PrimaryExprIdent:
				switch y := x.lookup().(type) {
				case *FuncDecl, *VarDecl:
					n.checkCall(s, f)
				case *TypeDecl:
					u = newNamedType(x.Ident)
					s.check(u, n)
					n.checkConversion(s, u)
				default:
					panic(fmt.Sprintf("%v: TODO915 %T", n.PrimaryExpr.Position(), y))
				}
			case *PrimaryExprCall, *PrimaryExprFuncLiteral:
				n.checkCall(s, f)
			case *PrimaryExprSelector:
				if _, ok := x.Declaration.(*TypeDecl); ok {
					n.checkConversion(s, t)
					break
				}

				n.checkCall(s, f)
			default:
				panic(fmt.Sprintf("%v: TODO899 %T", n.PrimaryExpr.Position(), x))
			}
		default:
			n.checkConversion(s, t)
		}
	}
}

func (n *PrimaryExprCall) wrongArgCount(have []ExprOrType, want *FunctionType) string {
	var a []string
	for i, v := range have {
		var t Type
		switch x := v.(type) {
		case *ExprOrTypeExpr:
			t = x.Type()
		case *ExprOrTypeType:
			t = x.Type()
		default:
			panic(fmt.Sprintf("unexpected type %T", x))
		}
		s := fmt.Sprint(t)
		if i == len(have)-1 && n.DDD {
			s += "..."
		}
		a = append(a, s)
	}
	return fmt.Sprintf("\n\thave (%s)\n\twant (%s)", strings.Join(a, ", "), want.ParameterList.types())
}

func (n *PrimaryExprCall) checkCall(s *cstack, f *FunctionType) {
	params := f.ParameterList.tuple()
	n.IsConversion = true
	for _, v := range n.Args {
		s.check(v, n)
	}

	args := n.Args
	if len(args) == 1 {
		var t Type
		switch x := args[0].(type) {
		case *ExprOrTypeExpr:
			t = x.Type()
		case *ExprOrTypeType:
			t = x.Type()
		default:
			panic(fmt.Sprintf("unexpected type %T", x))
		}
		if t == nil || t == Invalid {
			return
		}

		if x, ok := t.(*TupleType); ok {
			args = make([]ExprOrType, len(x.List))
			for i, v := range x.List {
				args[i] = &ExprOrTypeType{typ: v}
			}
		}
	}

	nparam := len(params.List)
	if !n.DDD {
		switch g, e := len(args), len(params.List); {
		case (g < e && !f.Variadic) || (g < e-1 && f.Variadic):
			n.err("not enough arguments in call to %s%s", n.PrimaryExpr.str(), n.wrongArgCount(n.Args, f))
			return
		case g > e && (e == 0 || !f.Variadic):
			n.err("too many arguments in call to %s%s", n.PrimaryExpr.str(), n.wrongArgCount(n.Args, f))
			return
		}

		var t Type
		for i, v := range args {
			if i < nparam {
				t = params.List[i]
			}
			if t == nil || t == Invalid || t.isGeneric(map[Type]bool{}) { //TODO- isGeneric
				continue
			}

			if i == nparam-1 && f.Variadic {
				t = t.(*SliceType).Element
			}
			switch x := v.(type) {
			case *ExprOrTypeExpr:
				if y := x.Type(); y != nil && y != Invalid && /*TODO- */ !y.isGeneric(map[Type]bool{}) && !y.AssignableTo(t) {
					// dbg("%v: %v variadic %v", n.Position(), n.PrimaryExpr.Type(), f.Variadic)
					// dbg("%v: %T(%s)", y.Position(), y, y)
					// dbg("%v: %T(%s)", t.Position(), t, t)
					n.err0(x.Expr.Position(),
						fmt.Sprintf("cannot use %s (type %s) as type %s in argument to %s",
							x.Expr.str(), y, t, n.PrimaryExpr.str(),
						))
				}
			case *ExprOrTypeType:
				//TODO
			default:
				panic(fmt.Sprintf("%v: TODO937 %T", n.Position(), x))
			}
		}
		return
	}

	if !f.Variadic {
		n.err("invalid use of ... in call to %s", n.PrimaryExpr.str())
	}

	switch g, e := len(args), len(params.List); {
	case g < e:
		n.err("not enough arguments in call to %s%s", n.PrimaryExpr.str(), n.wrongArgCount(n.Args, f))
		return
	case g > e:
		n.err("too many arguments in call to %s%s", n.PrimaryExpr.str(), n.wrongArgCount(n.Args, f))
		return
	}

	for i, v := range args {
		t := params.List[i]
		if t == nil || t == Invalid || t.isGeneric(map[Type]bool{}) { //TODO- isGeneric
			continue
		}

		switch x := v.(type) {
		case *ExprOrTypeExpr:
			if y := x.Type(); y != nil && y != Invalid && /*TODO- */ !y.isGeneric(map[Type]bool{}) && !y.AssignableTo(t) {
				// dbg("%v: %v variadic %v", n.Position(), n.PrimaryExpr.Type(), f.Variadic)
				// dbg("%v: %T(%s)", y.Position(), y, y)
				// dbg("%v: %T(%s)", t.Position(), t, t)
				n.err0(x.Expr.Position(),
					fmt.Sprintf("cannot use %s (type %s) as type %s in argument to %s",
						x.Expr.str(), y, t, n.PrimaryExpr.str(),
					))
			}
		case *ExprOrTypeType:
			//TODO
		default:
			panic(fmt.Sprintf("%v: TODO937 %T", n.Position(), x))
		}
	}

	//TODO
}

func (n *PrimaryExprCall) checkConversion(s *cstack, t Type) {
	n.expr.typ = t
	if n.DDD {
		panic(fmt.Sprintf("%v: TODO933", n.PrimaryExpr.Position())) // invalid ...
	}

	switch len(n.Args) {
	case 0:
		panic(fmt.Sprintf("%v: TODO938 %v(%v)", n.PrimaryExpr.Position(), t, t.UnderlyingType())) // missing arg
	case 1:
		s.check(n.Args[0], n)
		//TODO check conversion possible
	default:
		panic(fmt.Sprintf("%v: TODO943 %v", n.PrimaryExpr.Position(), t)) // too many args
	}
}

// PrimaryExprTypeSwitchGuard is an Expr.
type PrimaryExprTypeSwitchGuard struct {
	npos
	expr
	PrimaryExpr Expr
}

func (n *PrimaryExprTypeSwitchGuard) str() string { return "TODO PrimaryExprTypeSwitchGuard.str()" }

func (n *PrimaryExprTypeSwitchGuard) check(s *cstack) {
	//TODO
}

// PrimaryExprTypeAssertion is an Expr.
type PrimaryExprTypeAssertion struct {
	npos
	expr
	PrimaryExpr Expr
	Assert      ExprOrType
}

func (n *PrimaryExprTypeAssertion) str() string { return "TODO PrimaryExprTypeAssertion.str()" }

func (n *PrimaryExprTypeAssertion) check(s *cstack) {
	//TODO
}

// PrimaryExprSelector is an Expr.
type PrimaryExprSelector struct {
	npos
	expr
	PrimaryExpr Expr
	Selector    Ident
	Declaration Declaration // Non-nil if the selector denotes a declaration.
}

func (n *PrimaryExprSelector) str() string {
	return fmt.Sprintf("%s.%s", n.PrimaryExpr.str(), n.Selector.Val)
}

func (n *PrimaryExprSelector) check(s *cstack) {
	for {
		switch x := n.PrimaryExpr.(type) {
		case *PrimaryExprParenExpr:
			n.PrimaryExpr = x.unparen()
		case *PrimaryExprIdent:
			if x.Ident.Val == "C" { //TODO
				return
			}

			sf := x.Ident.SourceFile()
			n.Declaration = x.Ident.Scope().Lookup(sf.Package, sf.Scope, x.Ident.Token)
			switch y := n.Declaration.(type) {
			case nil:
				n.err0(x.Ident.Position(), fmt.Sprintf("undefined: %s", x.Ident.Val))
				return
			case *ConstDecl:
				s.check(y, n)
				if n.Selector.Val == "_" {
					n.err0(n.Selector.Position(), "TODO1168")
					return
				}

				//TODO
				return
			case *ImportDecl:
				s.check(y, n)
				if n.Selector.Val == "_" {
					n.err0(n.Selector.Position(), "TODO1177")
					return
				}

				if !isExported(n.Selector.Val) {
					n.err0(x.Ident.Position(), fmt.Sprintf("cannot refer to unexported name %s.%s", x.Ident.Val, n.Selector.Val))
					return
				}

				d := y.Package().Scope.Lookup(y.Package(), nil, n.Selector.Token)
				n.Declaration = d
				if d == nil {
					n.err0(x.Ident.Position(), fmt.Sprintf("undefined: %s.%s", x.Ident.Val, n.Selector.Val))
					return
				}

				switch z := d.(type) {
				case
					*ConstDecl,
					*FuncDecl,
					*VarDecl:

					n.expr.typ = z.Type()
				case *TypeDecl:
					t := newQualifiedNamedType(x.Ident, n.Selector.Val)
					s.check(t, n)
					n.expr.typ = t
				default:
					panic(fmt.Sprintf("unexpected type %T", z))
				}
				return
			case *TypeDecl:
				s.check(y, n)
				if n.Selector.Val == "_" {
					n.err0(n.Selector.Position(), "TODO1211")
					return
				}

				//TODO
				return
			case *VarDecl:
				s.check(y, n)
				if n.Selector.Val == "_" {
					n.err0(n.Selector.Position(), "cannot refer to blank field or method")
					return
				}

				if t := y.Type(); t != nil && t != Invalid {
					n.expr.typ = t.selector(n.Selector.Token)
					if n.expr.typ == Invalid {
						//TODO n.err("%s.%s undefined (type %s has no field or method %[2]s)", n.PrimaryExpr.str(), n.Selector.Val, t)
						return
					}
				}

				return
			default:
				n.err0(x.Ident.Position(), fmt.Sprintf("unexpected type %T", y))
				return
			}
		default:
			//TODO
			return
		}
	}
}

// PrimaryExprIndex is an Expr.
type PrimaryExprIndex struct {
	npos
	expr
	PrimaryExpr Expr
	Index       Expr
}

func (n *PrimaryExprIndex) str() string { return "TODO PrimaryExprIndex.str()" }

func (n *PrimaryExprIndex) check(s *cstack) {
	//TODO
}

// PrimaryExprSimpleSlice is an Expr.
type PrimaryExprSimpleSlice struct {
	npos
	expr
	PrimaryExpr Expr
	Lo          Expr
	Hi          Expr
}

func (n *PrimaryExprSimpleSlice) str() string { return "TODO PrimaryExprSimpleSlice.str()" }

func (n *PrimaryExprSimpleSlice) check(s *cstack) {
	//TODO
}

// PrimaryExprFullSlice is an Expr.
type PrimaryExprFullSlice struct {
	npos
	expr
	PrimaryExpr Expr
	Lo          Expr
	Hi          Expr
	Max         Expr
}

func (n *PrimaryExprFullSlice) str() string { return "TODO PrimaryExprFullSlice.str()" }

func (n *PrimaryExprFullSlice) check(s *cstack) {
	//TODO
}

// PrimaryExpr represents data reduced by productions
//
// 	primaryExpr:
//		'(' exprOrType ')'
//	|	IDENT genericArgsOpt %prec _NotParen
//	|	convType '(' expr commaOpt ')'
//	|	fnType lbrace stmtList '}'
//	|	literal
//	|	otherType lbrace bracedKeyValList '}'
//	|	primaryExpr '(' ')'
//	|	primaryExpr '(' exprOrTypeList "..." commaOpt ')'
//	|	primaryExpr '(' exprOrTypeList commaOpt ')'
//	|	primaryExpr '.' '(' "type" ')'
//	|	primaryExpr '.' '(' exprOrType ')'
//	|	primaryExpr '.' IDENT
//	|	primaryExpr '[' expr ']'
//	|	primaryExpr '[' exprOpt ':' exprOpt ':' exprOpt ']'
//	|	primaryExpr '[' exprOpt ':' exprOpt ']'
//	|	primaryExpr '{' bracedKeyValList '}'
type PrimaryExpr Expr

func (p *parser) primaryExpr() (r Expr) {
	var fix bool
	tok := p.tok()
	switch ch := p.c; ch {
	case gotoken.LPAREN:
		p.n()
		switch x := p.exprOrType(")").(type) {
		case *ExprOrTypeExpr:
			r = &PrimaryExprParenExpr{npos: newPos(tok), Expr: x.Expr}
		case *ExprOrTypeType:
			r = &primaryExprParenType{npos: newPos(tok), expr: expr{typ: x.Type()}}
		case *exprOrType:
			// nop
		default:
			panic("internal error 015")
		}
		p.must(gotoken.RPAREN)
	case gotoken.IDENT:
		r = &PrimaryExprIdent{Ident: p.ident()}
		p.n()
		if p.c == gotoken.COLON {
			return r
		}

		p.genericArgsOpt()
	case gotoken.FUNC:
		t := p.fnType()
		tok := p.tok()
		switch p.c {
		case tokenBODY:
			fix = true
			fallthrough
		case gotoken.LBRACE:
			p.n()
			r = &PrimaryExprFuncLiteral{npos: newPos(tok), expr: expr{typ: t}, Body: &StmtBlock{s: p.scope, List: p.stmtList()}}
			p.loophack = fix
			p.must(gotoken.RBRACE)
			p.pop()
		case gotoken.LPAREN:
			p.pop()
			p.n()
			r = &PrimaryExprConversion{npos: newPos(tok), expr: expr{typ: t}, Expr: p.expr()}
			p.opt(gotoken.COMMA)
			p.must(gotoken.RPAREN)
		default:
			p.pop()
			p.err("type %s is not an expression", t)
			p.syntaxError(p)
		}
	case gotoken.INT:
		r = &PrimaryExprIntLiteral{Literal: tok}
		p.n()
	case gotoken.FLOAT:
		r = &PrimaryExprFloatLiteral{Literal: tok}
		p.n()
	case gotoken.IMAG:
		r = &PrimaryExprImagLiteral{Literal: tok}
		p.n()
	case gotoken.CHAR:
		r = &PrimaryExprRuneLiteral{Literal: tok}
		p.n()
	case gotoken.STRING:
		r = &PrimaryExprStringLiteral{Literal: tok}
		p.n()
	case gotoken.CHAN, gotoken.INTERFACE, gotoken.MAP, gotoken.STRUCT, gotoken.LBRACK:
		t := p.otherType(ch)
		tok := p.tok()
		switch p.c {
		case gotoken.LPAREN:
			p.n()
			r = &PrimaryExprConversion{npos: newPos(p.tok()), expr: expr{typ: t}, Expr: p.expr()}
			p.opt(gotoken.COMMA)
			p.must(gotoken.RPAREN)
		case tokenBODY:
			fix = true
			fallthrough
		case gotoken.LBRACE:
			p.n()
			r = &PrimaryExprCompositeLiteral{npos: newPos(tok), expr: expr{typ: t}, List: p.bracedKeyValList()}
			p.loophack = fix
			p.must(gotoken.RBRACE)
		default:
			p.err("syntax error: unexpected %v, expecting { or (", p.unexpected())
			p.syntaxError(p)
		}
	default:
		//TODO p.err()
		p.syntaxError(p)
	}
	return p.primaryExpr2(r)
}

func (p *parser) primaryExpr2(r0 Expr) (r Expr) {
	defer func() {
		if r0 == r {
			return
		}

		for {
			switch x := r0.(type) {
			case *PrimaryExprParenExpr:
				r0 = x.unparen()
			case *PrimaryExprIdent:
				if x.Ident.Val == "_" {
					p.err0(x.Ident.Position(), "cannot use _ as value")
				}
				return
			default:
				return
			}
		}
	}()

	r = r0
	for {
		tok := p.tok()
		switch p.c {
		case gotoken.LPAREN:
			switch x := r.(type) {
			case *PrimaryExprConversion:
				p.n()
				if p.opt(gotoken.RPAREN) {
					//TODO p.err() // missing conversion operand
					break
				}

				list := p.exprOrTypeList()
				switch {
				case len(list) == 1:
					switch y := list[0].(type) {
					case *ExprOrTypeExpr:
						x.Expr = y.Expr
					case *exprOrType:
						// nop
					default:
						panic("internal error 016")
					}
				default:
					//TODO p.err() // too many conversion arguments
				}
				ddd := p.opt(gotoken.ELLIPSIS)
				if ddd {
					//TODO p.syntaxError()
				}
				_ = ddd && p.opt(gotoken.COMMA) //TODOOK
				p.must(gotoken.RPAREN)
			case *primaryExprParenType:
				p.n()
				if p.opt(gotoken.RPAREN) {
					//TODO p.err() // missing conversion operand
					break
				}

				list := p.exprOrTypeList()
				switch {
				case len(list) == 1:
					switch y := list[0].(type) {
					case *ExprOrTypeExpr:
						r = &PrimaryExprConversion{npos: newPos(tok), expr: expr{typ: x.Type()}, Expr: y.Expr}
					case *exprOrType:
						// nop
					default:
						panic("internal error 017")
					}
				default:
					//TODO p.err() // too many conversion arguments
				}
				ddd := p.opt(gotoken.ELLIPSIS)
				if ddd {
					//TODO p.syntaxError()
				}
				_ = ddd && p.opt(gotoken.COMMA) //TODOOK
				p.must(gotoken.RPAREN)
			case
				*PrimaryExprCall,
				*PrimaryExprFuncLiteral,
				*PrimaryExprIdent,
				*PrimaryExprIndex,
				*PrimaryExprParenExpr,
				*PrimaryExprSelector,
				*PrimaryExprTypeAssertion:

				c := &PrimaryExprCall{npos: newPos(tok), PrimaryExpr: r}
				p.n()
				if p.opt(gotoken.RPAREN) {
					r = c
					break
				}

				c.Args = p.exprOrTypeList()
				c.DDD = p.opt(gotoken.ELLIPSIS)
				r = c
				_ = c.DDD && p.opt(gotoken.COMMA) //TODOOK
				p.must(gotoken.RPAREN)
			default:
				c := &PrimaryExprCall{npos: newPos(tok), PrimaryExpr: r}
				p.n()
				if p.opt(gotoken.RPAREN) {
					r = c
					break
				}

				c.Args = p.exprOrTypeList()
				c.DDD = p.opt(gotoken.ELLIPSIS)
				r = c
				_ = c.DDD && p.opt(gotoken.COMMA) //TODOOK
				p.must(gotoken.RPAREN)
			}
		case gotoken.PERIOD:
			pos := newPos(p.tok())
			switch p.n() {
			case gotoken.IDENT:
				r = &PrimaryExprSelector{npos: pos, PrimaryExpr: r, Selector: p.ident()} //TODO proper Scope
				p.n()
			case gotoken.LPAREN:
				p.n()
				switch {
				case p.opt(gotoken.TYPE):
					r = &PrimaryExprTypeSwitchGuard{npos: pos, PrimaryExpr: r}
				default:
					r = &PrimaryExprTypeAssertion{npos: pos, PrimaryExpr: r, Assert: p.exprOrType(")")}
				}
				p.must(gotoken.RPAREN)
			default:
				p.syntaxError(p)
				p.err("syntax error: unexpected %v, expecting name or (", p.unexpected())
				p.n()
			}
		case gotoken.LBRACK:
			pos := newPos(p.tok())
			p.n()
			if p.c == gotoken.RBRACK {
				p.syntaxError(p)
				//TODO p.err()
				break
			}

			e := p.exprOpt()
			switch {
			case p.opt(gotoken.COLON):
				e2 := p.exprOpt()
				switch {
				case p.opt(gotoken.COLON):
					e3 := p.exprOpt()
					r = &PrimaryExprFullSlice{npos: pos, PrimaryExpr: r, Lo: e, Hi: e2, Max: e3}
				default:
					r = &PrimaryExprSimpleSlice{npos: pos, PrimaryExpr: r, Lo: e, Hi: e2}
				}
			default:
				r = &PrimaryExprIndex{npos: pos, PrimaryExpr: r, Index: e}
			}
			p.must(gotoken.RBRACK)
		case gotoken.LBRACE:
			tok := p.tok()
			p.n()
			list := p.bracedKeyValList()
			switch x := r.(type) {
			case *PrimaryExprIdent:
				r = &PrimaryExprCompositeLiteral{npos: newPos(tok), expr: expr{typ: newNamedType(x.Ident)}, List: list}
			case *PrimaryExprCompositeLiteral:
				x.List = list
			case *PrimaryExprSelector:
				switch y := x.PrimaryExpr.(type) {
				case *PrimaryExprIdent:
					// TODO x.y.z{...} is properly rejected
					if !isExported(x.Selector.Val) {
						//TODO p.err unexported
					}
					r = &PrimaryExprCompositeLiteral{npos: newPos(tok), expr: expr{typ: newQualifiedNamedType(y.Ident, x.Selector.Val)}, List: list}
				default:
					panic("internal error 018")
				}
			}
			p.must(gotoken.RBRACE)
		default:
			return r
		}
	}
}

// UnaryExprReceive is an Expr.
type UnaryExprReceive struct {
	npos
	expr
	Expr Expr
}

func (n *UnaryExprReceive) str() string { return "TODO UnaryExprReceive.str()" }

func (n *UnaryExprReceive) check(s *cstack) {
	//TODO
}

func (n *UnaryExprReceive) setExpr(e Expr) { n.Expr = e }

// UnaryExprNot is an Expr.
type UnaryExprNot struct {
	npos
	expr
	Expr Expr
}

func (n *UnaryExprNot) str() string { return "TODO UnaryExprNot.str()" }

func (n *UnaryExprNot) check(s *cstack) {
	//TODO
}

func (n *UnaryExprNot) setExpr(e Expr) { n.Expr = e }

// UnaryExprAddr is an Expr.
type UnaryExprAddr struct {
	npos
	expr
	Expr Expr
}

func (n *UnaryExprAddr) str() string { return "TODO UnaryExprAddr.str()" }

func (n *UnaryExprAddr) check(s *cstack) {
	//TODO
}

func (n *UnaryExprAddr) setExpr(e Expr) { n.Expr = e }

// UnaryExprDeref is an Expr.
type UnaryExprDeref struct {
	npos
	expr
	Expr Expr
}

func (n *UnaryExprDeref) str() string { return fmt.Sprintf("*%s", n.Expr.str()) }

func (n *UnaryExprDeref) check(s *cstack) {
	//TODO
}

func (n *UnaryExprDeref) setExpr(e Expr) { n.Expr = e }

// UnaryExprPos is an Expr.
type UnaryExprPos struct {
	npos
	expr
	Expr Expr
}

func (n *UnaryExprPos) str() string { return "TODO UnaryExprPos.str()" }

func (n *UnaryExprPos) check(s *cstack) {
	//TODO
}

func (n *UnaryExprPos) setExpr(e Expr) { n.Expr = e }

// UnaryExprNeg is an Expr.
type UnaryExprNeg struct {
	npos
	expr
	Expr Expr
}

func (n *UnaryExprNeg) str() string { return "TODO UnaryExprNeg.str()" }

func (n *UnaryExprNeg) check(s *cstack) {
	//TODO
}

func (n *UnaryExprNeg) setExpr(e Expr) { n.Expr = e }

// UnaryExprXor is an Expr.
type UnaryExprXor struct {
	npos
	expr
	Expr Expr
}

func (n *UnaryExprXor) str() string { return "TODO UnaryExprXor.str()" }

func (n *UnaryExprXor) check(s *cstack) {
	//TODO
}

func (n *UnaryExprXor) setExpr(e Expr) { n.Expr = e }

// UnaryExpr represents data reduced by productions
//
//	unaryExpr:
//		"<-" unaryExpr
//	|	'!' unaryExpr
//	|	'&' unaryExpr
//	|	'*' unaryExpr
//	|	'+' unaryExpr
//	|	'-' unaryExpr
//	|	'^' unaryExpr
//	|	primaryExpr
type UnaryExpr Expr

func (p *parser) unaryExpr() (r Expr) {
	for {
		pos := newPos(p.tok())
		switch p.c {
		case gotoken.ARROW:
			r = &UnaryExprReceive{npos: pos, Expr: r}
		case gotoken.NOT:
			r = &UnaryExprNot{npos: pos, Expr: r}
		case gotoken.AND:
			r = &UnaryExprAddr{npos: pos, Expr: r}
		case gotoken.MUL:
			r = &UnaryExprDeref{npos: pos, Expr: r}
		case gotoken.ADD:
			r = &UnaryExprPos{npos: pos, Expr: r}
		case gotoken.SUB:
			r = &UnaryExprNeg{npos: pos, Expr: r}
		case gotoken.XOR:
			r = &UnaryExprXor{npos: pos, Expr: r}
		default:
			e := p.primaryExpr()
			if r == nil {
				return e
			}

			r.setExpr(e)
			return r
		}
		p.n()
	}
}

// Expr represents data reduced by productions
//
//	expr:
//		expr "!=" expr
//	|	expr "&&" expr
//	|	expr "&^" expr
//	|	expr "<-" expr
//	|	expr "<<" expr
//	|	expr "<=" expr
//	|	expr "==" expr
//	|	expr ">=" expr
//	|	expr ">>" expr
//	|	expr "||" expr
//	|	expr '%' expr
//	|	expr '&' expr
//	|	expr '*' expr
//	|	expr '+' expr
//	|	expr '-' expr
//	|	expr '/' expr
//	|	expr '<' expr
//	|	expr '>' expr
//	|	expr '^' expr
//	|	expr '|' expr
//	|	unaryExpr
type Expr interface {
	Operand
	Position() token.Position
	checker
	isExpr()
	lhs() Expr
	pos() npos
	prec() int
	rot()
	setExpr(Expr)
	setLHS(Expr)
	setRHS(Expr)
	str() string
}

type expr struct {
	typ Type
	val Value

	sentinel
}

func (e *expr) Type() Type   { return e.typ }
func (e *expr) Value() Value { return e.val }
func (e *expr) isExpr()      {}
func (e *expr) lhs() Expr    { panic("internal error 019") }
func (e *expr) prec() int    { return 0 }
func (e *expr) rot()         { panic("internal error 020") }
func (e *expr) setExpr(Expr) { panic("internal error 021") }
func (e *expr) setLHS(Expr)  {}
func (e *expr) setRHS(Expr)  {}

// ExprNEQ represents LHS != RHS.
type ExprNEQ struct {
	npos
	expr
	LHS Expr
	RHS Expr
}

func (n *ExprNEQ) str() string { return fmt.Sprintf("%s != %s", n.LHS.str(), n.RHS.str()) }

func (n *ExprNEQ) check(s *cstack) {
	s.check(n.LHS, n)
	s.check(n.RHS, n)
	if t := n.LHS.Type(); t != nil {
		if ok, s := t.ComparableWith(n.RHS.Type()); !ok {
			if s != "" {
				s = fmt.Sprintf(" (%s)", s)
			}
			// dbg("LHS %s", n.LHS.str())
			// dbg("LHS.Type %s %s", n.LHS.Type(), n.LHS.Type().Kind())
			// dbg("RHS %s", n.RHS.str())
			// dbg("RHS.Type %s %s", n.RHS.Type(), n.RHS.Type().Kind())
			n.sf.Package.errorList.add(n.Position(), "invalid operation: %s%s", n.str(), s)
		}
	}
	n.expr.typ = UntypedBool
}

func (n *ExprNEQ) lhs() Expr     { return n.LHS }
func (n *ExprNEQ) setLHS(e Expr) { n.LHS = e }
func (n *ExprNEQ) setRHS(e Expr) { n.RHS = e }
func (*ExprNEQ) prec() int       { return gotoken.NEQ.Precedence() }

// ExprLAND represents LHS && RHS.
type ExprLAND struct {
	npos
	expr
	LHS Expr
	RHS Expr
}

func (n *ExprLAND) str() string { return "TODO ExprLAND.str()" }

func (n *ExprLAND) check(s *cstack) {
	//TODO
}

func (n *ExprLAND) lhs() Expr     { return n.LHS }
func (n *ExprLAND) setLHS(e Expr) { n.LHS = e }
func (n *ExprLAND) setRHS(e Expr) { n.RHS = e }
func (*ExprLAND) prec() int       { return gotoken.LAND.Precedence() }

// ExprANDNOT represents LHS &^ RHS.
type ExprANDNOT struct {
	npos
	expr
	LHS Expr
	RHS Expr
}

func (n *ExprANDNOT) str() string { return "TODO ExprANDNOT.str()" }

func (n *ExprANDNOT) check(s *cstack) {
	//TODO
}

func (n *ExprANDNOT) lhs() Expr     { return n.LHS }
func (n *ExprANDNOT) setLHS(e Expr) { n.LHS = e }
func (n *ExprANDNOT) setRHS(e Expr) { n.RHS = e }
func (*ExprANDNOT) prec() int       { return gotoken.AND_NOT.Precedence() }

// ExprARROW represents LHS <- RHS.
type ExprARROW struct {
	npos
	expr
	LHS Expr
	RHS Expr
}

func (n *ExprARROW) str() string { return "TODO ExprARROW.str()" }

func (n *ExprARROW) check(s *cstack) {
	//TODO
}

func (n *ExprARROW) lhs() Expr     { return n.LHS }
func (n *ExprARROW) setLHS(e Expr) { n.LHS = e }
func (n *ExprARROW) setRHS(e Expr) { n.RHS = e }
func (*ExprARROW) prec() int       { return gotoken.ARROW.Precedence() }

// ExprSHL represents LHS << RHS.
type ExprSHL struct {
	npos
	expr
	LHS Expr
	RHS Expr
}

func (n *ExprSHL) str() string { return "TODO ExprSHL.str()" }

func (n *ExprSHL) check(s *cstack) {
	//TODO
}

func (n *ExprSHL) lhs() Expr     { return n.LHS }
func (n *ExprSHL) setLHS(e Expr) { n.LHS = e }
func (n *ExprSHL) setRHS(e Expr) { n.RHS = e }
func (*ExprSHL) prec() int       { return gotoken.SHL.Precedence() }

// ExprLEQ represents LHS <= RHS.
type ExprLEQ struct {
	npos
	expr
	LHS Expr
	RHS Expr
}

func (n *ExprLEQ) str() string { return "TODO ExprLEQ.str()" }

func (n *ExprLEQ) check(s *cstack) {
	//TODO
}

func (n *ExprLEQ) lhs() Expr     { return n.LHS }
func (n *ExprLEQ) setLHS(e Expr) { n.LHS = e }
func (n *ExprLEQ) setRHS(e Expr) { n.RHS = e }
func (*ExprLEQ) prec() int       { return gotoken.LEQ.Precedence() }

// ExprEQL represents LHS == RHS.
type ExprEQL struct {
	npos
	expr
	LHS Expr
	RHS Expr
}

func (n *ExprEQL) str() string { return "TODO ExprEQL.str()" }

func (n *ExprEQL) check(s *cstack) {
	s.check(n.LHS, n)
	s.check(n.RHS, n)
	if t := n.LHS.Type(); t != nil {
		if ok, s := t.ComparableWith(n.RHS.Type()); !ok {
			if s != "" {
				s = fmt.Sprintf(" (%s)", s)
			}
			//dbg("", n.LHS.str())
			//dbg("%T(%[1]s)", n.LHS.Type())
			//dbg("", n.RHS.str())
			//dbg("%T(%[1]s)", n.RHS.Type())
			//panic("TODO2045")
			n.sf.Package.errorList.add(n.Position(), "invalid operation: %s == %s%s", n.LHS.str(), n.RHS.str(), s)
		}
	}
	n.expr.typ = UntypedBool
}

func (n *ExprEQL) lhs() Expr     { return n.LHS }
func (n *ExprEQL) setLHS(e Expr) { n.LHS = e }
func (n *ExprEQL) setRHS(e Expr) { n.RHS = e }
func (*ExprEQL) prec() int       { return gotoken.EQL.Precedence() }

// ExprGEQ represents LHS >= RHS.
type ExprGEQ struct {
	npos
	expr
	LHS Expr
	RHS Expr
}

func (n *ExprGEQ) str() string { return "TODO ExprGEQ.str()" }

func (n *ExprGEQ) check(s *cstack) {
	//TODO
}

func (n *ExprGEQ) lhs() Expr     { return n.LHS }
func (n *ExprGEQ) setLHS(e Expr) { n.LHS = e }
func (n *ExprGEQ) setRHS(e Expr) { n.RHS = e }
func (*ExprGEQ) prec() int       { return gotoken.GEQ.Precedence() }

// ExprSHR represents LHS >> RHS.
type ExprSHR struct {
	npos
	expr
	LHS Expr
	RHS Expr
}

func (n *ExprSHR) str() string { return "TODO ExprSHR.str()" }

func (n *ExprSHR) check(s *cstack) {
	//TODO
}

func (n *ExprSHR) lhs() Expr     { return n.LHS }
func (n *ExprSHR) setLHS(e Expr) { n.LHS = e }
func (n *ExprSHR) setRHS(e Expr) { n.RHS = e }
func (*ExprSHR) prec() int       { return gotoken.SHR.Precedence() }

// ExprLOR represents LHS || RHS.
type ExprLOR struct {
	npos
	expr
	LHS Expr
	RHS Expr
}

func (n *ExprLOR) str() string { return "TODO ExprLOR.str()" }

func (n *ExprLOR) check(s *cstack) {
	//TODO
}

func (n *ExprLOR) lhs() Expr     { return n.LHS }
func (n *ExprLOR) setLHS(e Expr) { n.LHS = e }
func (n *ExprLOR) setRHS(e Expr) { n.RHS = e }
func (*ExprLOR) prec() int       { return gotoken.LOR.Precedence() }

// ExprREM represents LHS % RHS.
type ExprREM struct {
	npos
	expr
	LHS Expr
	RHS Expr
}

func (n *ExprREM) str() string { return "TODO ExprREM.str()" }

func (n *ExprREM) check(s *cstack) {
	//TODO
}

func (n *ExprREM) lhs() Expr     { return n.LHS }
func (n *ExprREM) setLHS(e Expr) { n.LHS = e }
func (n *ExprREM) setRHS(e Expr) { n.RHS = e }
func (*ExprREM) prec() int       { return gotoken.REM.Precedence() }

// ExprAND represents LHS & RHS.
type ExprAND struct {
	npos
	expr
	LHS Expr
	RHS Expr
}

func (n *ExprAND) str() string { return "TODO ExprAND.str()" }

func (n *ExprAND) check(s *cstack) {
	//TODO
}

func (n *ExprAND) lhs() Expr     { return n.LHS }
func (n *ExprAND) setLHS(e Expr) { n.LHS = e }
func (n *ExprAND) setRHS(e Expr) { n.RHS = e }
func (*ExprAND) prec() int       { return gotoken.AND.Precedence() }

// ExprMUL represents LHS * RHS.
type ExprMUL struct {
	npos
	expr
	LHS Expr
	RHS Expr
}

func (n *ExprMUL) str() string { return "TODO ExprMUL.str()" }

func (n *ExprMUL) check(s *cstack) {
	//TODO
}

func (n *ExprMUL) lhs() Expr     { return n.LHS }
func (n *ExprMUL) setLHS(e Expr) { n.LHS = e }
func (n *ExprMUL) setRHS(e Expr) { n.RHS = e }
func (*ExprMUL) prec() int       { return gotoken.MUL.Precedence() }

// ExprADD represents LHS + RHS.
type ExprADD struct {
	npos
	expr
	LHS Expr
	RHS Expr
}

func (n *ExprADD) str() string { return "TODO ExprADD.str()" }

func (n *ExprADD) check(s *cstack) {
	s.check(n.LHS, n)
	s.check(n.RHS, n)
	//TODO
}

func (n *ExprADD) lhs() Expr     { return n.LHS }
func (n *ExprADD) setLHS(e Expr) { n.LHS = e }
func (n *ExprADD) setRHS(e Expr) { n.RHS = e }
func (*ExprADD) prec() int       { return gotoken.ADD.Precedence() }

// ExprSUB represents LHS - RHS.
type ExprSUB struct {
	npos
	expr
	LHS Expr
	RHS Expr
}

func (n *ExprSUB) str() string { return "TODO ExprSUB.str()" }

func (n *ExprSUB) check(s *cstack) {
	//TODO
}

func (n *ExprSUB) lhs() Expr     { return n.LHS }
func (n *ExprSUB) setLHS(e Expr) { n.LHS = e }
func (n *ExprSUB) setRHS(e Expr) { n.RHS = e }
func (*ExprSUB) prec() int       { return gotoken.SUB.Precedence() }

// ExprQUO represents LHS / RHS.
type ExprQUO struct {
	npos
	expr
	LHS Expr
	RHS Expr
}

func (n *ExprQUO) str() string { return "TODO ExprQUO.str()" }

func (n *ExprQUO) check(s *cstack) {
	//TODO
}

func (n *ExprQUO) lhs() Expr     { return n.LHS }
func (n *ExprQUO) setLHS(e Expr) { n.LHS = e }
func (n *ExprQUO) setRHS(e Expr) { n.RHS = e }
func (*ExprQUO) prec() int       { return gotoken.QUO.Precedence() }

// ExprLSS represents LHS < RHS.
type ExprLSS struct {
	npos
	expr
	LHS Expr
	RHS Expr
}

func (n *ExprLSS) str() string { return "TODO ExprLSS.str()" }

func (n *ExprLSS) check(s *cstack) {
	//TODO
}

func (n *ExprLSS) lhs() Expr     { return n.LHS }
func (n *ExprLSS) setLHS(e Expr) { n.LHS = e }
func (n *ExprLSS) setRHS(e Expr) { n.RHS = e }
func (*ExprLSS) prec() int       { return gotoken.LSS.Precedence() }

// ExprGTR represents LHS > RHS.
type ExprGTR struct {
	npos
	expr
	LHS Expr
	RHS Expr
}

func (n *ExprGTR) str() string { return "TODO ExprGTR.str()" }

func (n *ExprGTR) check(s *cstack) {
	//TODO
}

func (n *ExprGTR) lhs() Expr     { return n.LHS }
func (n *ExprGTR) setLHS(e Expr) { n.LHS = e }
func (n *ExprGTR) setRHS(e Expr) { n.RHS = e }
func (*ExprGTR) prec() int       { return gotoken.GTR.Precedence() }

// ExprXOR represents LHS 6 RHS.
type ExprXOR struct {
	npos
	expr
	LHS Expr
	RHS Expr
}

func (n *ExprXOR) str() string { return "TODO ExprXOR.str()" }

func (n *ExprXOR) check(s *cstack) {
	//TODO
}

func (n *ExprXOR) lhs() Expr     { return n.LHS }
func (n *ExprXOR) setLHS(e Expr) { n.LHS = e }
func (n *ExprXOR) setRHS(e Expr) { n.RHS = e }
func (*ExprXOR) prec() int       { return gotoken.XOR.Precedence() }

// ExprOR represents LHS | RHS.
type ExprOR struct {
	npos
	expr
	LHS Expr
	RHS Expr
}

func (n *ExprOR) str() string { return "TODO ExprOR.str()" }

func (n *ExprOR) check(s *cstack) {
	//TODO
}

func (n *ExprOR) lhs() Expr     { return n.LHS }
func (n *ExprOR) setLHS(e Expr) { n.LHS = e }
func (n *ExprOR) setRHS(e Expr) { n.RHS = e }
func (*ExprOR) prec() int       { return gotoken.OR.Precedence() }

func (p *parser) expr() (r Expr) {
	r = p.unaryExpr()
	if _, ok := r.(*PrimaryExprIdent); ok && p.c == gotoken.COLON {
		return r
	}

	return p.expr2(r)
}

func (p *parser) expr2(lhs Expr) (r Expr) {
	r = lhs
	for {
		pos := newPos(p.tok())
		switch p.c {
		case gotoken.LOR:
			r = &ExprLOR{npos: pos}
		case gotoken.LAND:
			r = &ExprLAND{npos: pos}
		case gotoken.EQL:
			r = &ExprEQL{npos: pos}
		case gotoken.NEQ:
			r = &ExprNEQ{npos: pos}
		case gotoken.LSS:
			r = &ExprLSS{npos: pos}
		case gotoken.LEQ:
			r = &ExprLEQ{npos: pos}
		case gotoken.GTR:
			r = &ExprGTR{npos: pos}
		case gotoken.GEQ:
			r = &ExprGEQ{npos: pos}
		case gotoken.ADD:
			r = &ExprADD{npos: pos}
		case gotoken.SUB:
			r = &ExprSUB{npos: pos}
		case gotoken.OR:
			r = &ExprOR{npos: pos}
		case gotoken.XOR:
			r = &ExprXOR{npos: pos}
		case gotoken.MUL:
			r = &ExprMUL{npos: pos}
		case gotoken.QUO:
			r = &ExprQUO{npos: pos}
		case gotoken.REM:
			r = &ExprREM{npos: pos}
		case gotoken.SHL:
			r = &ExprSHL{npos: pos}
		case gotoken.SHR:
			r = &ExprSHR{npos: pos}
		case gotoken.AND:
			r = &ExprAND{npos: pos}
		case gotoken.AND_NOT:
			r = &ExprANDNOT{npos: pos}
		case gotoken.ARROW:
			r = &ExprARROW{npos: pos}
		default:
			return r
		}

		p.n()
		rhs := p.expr()
		switch {
		case rhs != nil && rhs.prec() != 0 && r.prec() > rhs.prec():
			r.setLHS(lhs)
			r.setRHS(rhs.lhs())
			r0 := r
			r = rhs
			r.setLHS(r0)
		default:
			r.setLHS(lhs)
			r.setRHS(rhs)
		}
		lhs = r
	}
}

// exprList:
//		expr
//	|	exprList ',' expr
func (p *parser) exprList() (r []Expr) {
	r = []Expr{p.expr()}
	for p.opt(gotoken.COMMA) {
		r = append(r, p.expr())
	}
	return r
}

// constSpec:
//		identList
//	|	identList '=' exprList
//	|	identList typ
//	|	identList typ '=' exprList
func (p *parser) constSpec() (r []Declaration) {
	l := p.identList()
	var e []Expr
	var t Type

	defer func() {
		pos := token.NoPos
		if p.scope.Kind != PackageScope {
			pos = p.pos()
		}
		for i, v := range l {
			d := newConstDecl(v, pos)
			d.index = i
			d.list = e
			d.setType(t)
			p.declare(d)
			r = append(r, d)
		}
		_ = e //TODO use the list
	}()

	switch p.c {
	case gotoken.RPAREN, gotoken.SEMICOLON:
		return
	case gotoken.ASSIGN:
		p.n()
		p.exprList()
		return
	}

	t = p.typ()
	if p.opt(gotoken.ASSIGN) {
		e = p.exprList()
	}
	if p.not2(gotoken.SEMICOLON, gotoken.RPAREN) {
		p.syntaxError(p)
		//TODO p.err()
		p.skip(gotoken.SEMICOLON, gotoken.RPAREN)
	}
	return r
}

// constSpecList:
//		constSpec
//	|	constSpecList ';' constSpec
func (p *parser) constSpecList() (r []Declaration) {
	for r = p.constSpec(); p.opt(gotoken.SEMICOLON) && p.c != gotoken.RPAREN; {
		r = append(r, p.constSpec()...)
	}
	return r
}

// FieldDeclaration represents data reduced by productions
//
// 	fieldDecl:
//		'*' embeddedName literalOpt
//	|	identList typ literalOpt
//	|	embeddedName literalOpt
type FieldDeclaration interface{ isFieldDeclaration() }

type fieldDeclaration struct{}

func (fieldDeclaration) isFieldDeclaration() {}

// FieldDeclarationEmbedded is a FieldDeclaration.
type FieldDeclarationEmbedded struct {
	fieldDeclaration
	Type Type
	Tag  string

	Ptr bool
}

// FieldDeclarationNamed is a FieldDeclaration.
type FieldDeclarationNamed struct {
	fieldDeclaration
	Name Token
	Type Type
	Tag  string
}

func (p *parser) fieldDecl() (r []FieldDeclaration) {
	var l []Token
	var t Type
	switch p.c {
	case gotoken.IDENT:
		ident := p.ident()
		l = []Token{ident.Token}
		switch p.n() {
		case gotoken.INT, gotoken.FLOAT, gotoken.IMAG, gotoken.CHAR:
			//TODO p.err()
			p.n()
			return []FieldDeclaration{&FieldDeclarationEmbedded{Type: newNamedType(ident)}}
		case gotoken.STRING:
			lit := p.tok()
			p.n()
			return []FieldDeclaration{&FieldDeclarationEmbedded{Type: newNamedType(ident), Tag: lit.Val}}
		case gotoken.SEMICOLON, gotoken.RBRACE:
			return []FieldDeclaration{&FieldDeclarationEmbedded{Type: newNamedType(ident)}}
		case gotoken.PERIOD:
			p.n()
			ident2 := p.ident()
			p.must(gotoken.IDENT)
			if !isExported(ident2.Val) {
				//TODO p.err unexported
			}
			r = append(r, &FieldDeclarationEmbedded{Type: newQualifiedNamedType(ident, ident2.Val)})
		case gotoken.COMMA:
			p.n()
			l = append(l, p.identList()...)
			fallthrough
		default:
			t = p.typ()
			for _, v := range l {
				r = append(r, &FieldDeclarationNamed{Name: v, Type: t})
			}
		}
	case gotoken.MUL:
		tok := p.tok()
		if p.n() == gotoken.LPAREN {
			p.syntaxError(p)
			p.err("syntax error: cannot parenthesize embedded type")
			p.skip(gotoken.SEMICOLON, gotoken.RBRACE)
			return nil
		}

		r = []FieldDeclaration{&FieldDeclarationEmbedded{Type: newPointerType(tok.npos, p.qualifiedIdent()), Ptr: true}}
	case gotoken.LPAREN:
		p.syntaxError(p)
		p.err("syntax error: cannot parenthesize embedded type")
		p.skip(gotoken.SEMICOLON, gotoken.RBRACE)
		return nil
	default:
		p.syntaxError(p)
		//TODO p.err()
		p.skip(gotoken.SEMICOLON, gotoken.RBRACE)
		return nil
	}

	switch p.c {
	case gotoken.INT, gotoken.FLOAT, gotoken.IMAG, gotoken.CHAR:
		p.n()
	case gotoken.STRING:
		lit := p.tok().Val
		for _, v := range r {
			switch x := v.(type) {
			case *FieldDeclarationEmbedded:
				x.Tag = lit
			case *FieldDeclarationNamed:
				x.Tag = lit
			default:
				panic("internal error 032")
			}
		}
		p.n()
	}
	return r
}

// fieldDeclList:
//		fieldDecl
//	|	fieldDeclList ';' fieldDecl
func (p *parser) fieldDeclList() (r []FieldDeclaration) {
	for r = p.fieldDecl(); p.opt(gotoken.SEMICOLON) && p.c != gotoken.RBRACE; {
		r = append(r, p.fieldDecl()...)
	}
	return r
}

type interfaceDecl interface{ isInterfaceDecl() }

type aIinterfaceDecl struct{}

func (aIinterfaceDecl) isInterfaceDecl() {}

type interfaceDeclMethod struct {
	aIinterfaceDecl
	Name Token
	Type *FunctionType
}

type interfaceDeclEmbedded struct {
	aIinterfaceDecl
	Type Type
}

// interfaceDecl:
//		IDENT '(' paramTypeListCommaOptOpt ')' result
//	|	embeddedName
func (p *parser) interfaceDecl() (r interfaceDecl) {
	ident := p.ident()
	p.must(gotoken.IDENT)
	switch p.c {
	case gotoken.LPAREN:
		p.n()
		params := p.paramTypeListCommaOptOpt()
		p.must(gotoken.RPAREN)
		result := p.result()
		s := newScope(BlockScope, p.scope)
		pos := p.pos()
		params.declare(p, s, pos)
		result.declare(p, s, pos)
		return &interfaceDeclMethod{Name: ident.Token, Type: newFunctionType(ident.npos, params, result)}
	case gotoken.PERIOD:
		p.n()
		ident2 := p.ident()
		p.must(gotoken.IDENT)
		if !isExported(ident2.Val) {
			//TODO p.err unexported
		}
		return &interfaceDeclEmbedded{Type: newQualifiedNamedType(ident, ident2.Val)}
	case gotoken.SEMICOLON, gotoken.RBRACE:
		// nop
		return &interfaceDeclEmbedded{Type: newNamedType(ident)}
	case gotoken.COMMA:
		p.syntaxError(p)
		p.err("syntax error: name list not allowed in interface type")
		p.skip(gotoken.SEMICOLON, gotoken.RBRACE)
	default:
		p.syntaxError(p)
		p.err("syntax error: unexpected %v, expecting semicolon, newline, or }", p.unexpected())
		p.skip(gotoken.SEMICOLON, gotoken.RBRACE)
	}
	return nil
}

// interfaceDeclList:
//		interfaceDecl
//	|	interfaceDeclList ';' interfaceDecl
func (p *parser) interfaceDeclList() (r []interfaceDecl) {
	if n := p.interfaceDecl(); n != nil {
		r = append(r, n)
	}
	for p.opt(gotoken.SEMICOLON) && p.c != gotoken.RBRACE {
		if n := p.interfaceDecl(); n != nil {
			r = append(r, n)
		}
	}
	return r
}

// otherType:
//		"chan" "<-" typ
//	|	"chan" '(' typ ')'
//	|	"chan" qualifiedIdent
//	|	"chan" fnType
//	|	"chan" otherType
//	|	"chan" ptrType
//	|	"interface" lbrace '}'
//	|	"interface" lbrace interfaceDeclList semiOpt '}'
//	|	"map" '[' typ ']' typ
//	|	"struct" lbrace '}'
//	|	"struct" lbrace fieldDeclList semiOpt '}'
//	|	'[' "..." ']' typ
//	|	'[' exprOpt ']' typ
func (p *parser) otherType(ch gotoken.Token) (r Type) {
	var fix bool
	switch p.c {
	case gotoken.CHAN:
		tok := p.tok()
		switch p.n() {
		case gotoken.ARROW:
			p.n()
			return newChannelType(tok.npos, TxChan, p.typ())
		case gotoken.LPAREN:
			p.n()
			r = newChannelType(tok.npos, RxChan|TxChan, p.typ())
			p.must(gotoken.RPAREN)
		case gotoken.IDENT:
			return newChannelType(tok.npos, RxChan|TxChan, p.qualifiedIdent())
		case gotoken.FUNC:
			r = newChannelType(tok.npos, RxChan|TxChan, p.fnType())
			p.pop()
		case gotoken.MUL:
			return newChannelType(tok.npos, RxChan|TxChan, p.ptrType())
		default:
			return newChannelType(tok.npos, RxChan|TxChan, p.otherType(gotoken.CHAN))
		}
	case gotoken.INTERFACE:
		tok := p.tok()
		switch p.n() {
		case tokenBODY:
			fix = true
			fallthrough
		case gotoken.LBRACE:
			var l []interfaceDecl
			if p.n() != gotoken.RBRACE {
				l = p.interfaceDeclList()
			}
			p.loophack = fix
			p.must(gotoken.RBRACE)
			n := newInterfaceType(tok.npos)
			r = n
			for _, v := range l {
				switch x := v.(type) {
				case *interfaceDeclEmbedded:
					n.Embedded = append(n.Embedded, x.Type)
				case *interfaceDeclMethod:
					if n.Methods == nil {
						n.Methods = map[string]*FunctionType{}
					}
					nm := x.Name.Val
					if nm == "_" {
						p.err0(x.Name.Position(), "methods must have a unique non-blank name")
						continue
					}

					if _, ok := n.Methods[nm]; ok {
						//TODO p.err redeclared
						break
					}

					n.Methods[nm] = x.Type
					n.names = append(n.names, nm)
					n.methods = append(n.methods, x.Type)
				default:
					panic("internal error 030")
				}
			}
		default:
			p.syntaxError(p)
			//TODO p.err()
		}
	case gotoken.MAP:
		tok := p.tok()
		p.n()
		p.must(gotoken.LBRACK)
		k := p.typ()
		p.must(gotoken.RBRACK)
		return newMapType(tok.npos, k, p.typ())
	case gotoken.STRUCT:
		tok := p.tok()
		switch p.n() {
		case tokenBODY:
			fix = true
			fallthrough
		case gotoken.LBRACE:
			var l []FieldDeclaration
			if p.n() != gotoken.RBRACE {
				l = p.fieldDeclList()
				m := map[string]struct{}{}
				for _, v := range l {
					var nm string
					switch x := v.(type) {
					case *FieldDeclarationEmbedded:
						nm = x.Type.String()
					case *FieldDeclarationNamed:
						nm = x.Name.Val
					}
					if _, ok := m[nm]; ok && nm != "_" {
						//TODO p.err redeclared
					}
					m[nm] = struct{}{}
				}
			}
			p.loophack = fix
			p.must(gotoken.RBRACE)
			n := newStructType(tok.npos)
			n.Fields = l
			return n
		default:
			p.syntaxError(p)
			//TODO p.err()
		}
	case gotoken.LBRACK:
		tok := p.tok()
		p.n()
		var opt Expr
		ddd := true
		if !p.opt(gotoken.ELLIPSIS) {
			opt = p.exprOpt()
			ddd = false
		}
		p.must(gotoken.RBRACK)
		t := p.typ()
		switch {
		case ddd || opt != nil:
			return newArrayType(tok.npos, opt, t)
		default:
			return newSliceType(tok.npos, t)
		}
	default:
		p.syntaxError(p)
		switch ch {
		case gotoken.CHAN:
			p.err("syntax error: missing channel element type")
		default:
			//TODO p.err()
		}
	}
	return r
}

// qualifiedIdent:
//		IDENT %prec _NotParen
//	|	IDENT '.' IDENT
func (p *parser) qualifiedIdent() Type {
	tok, _ := p.mustTok(gotoken.IDENT)
	switch {
	case p.opt(gotoken.PERIOD):
		tok2, _ := p.mustTok(gotoken.IDENT)
		if !isExported(tok2.Val) {
			//TODO p.err unexported
		}
		r := newQualifiedNamedType(Ident{tok, p.scope}, tok2.Val)
		return r
	default:
		return newNamedType(Ident{tok, p.scope})
	}
}

// ptrType:
// 	'*' typ
func (p *parser) ptrType() *PointerType {
	tok := p.tok()
	p.must(gotoken.MUL)
	return newPointerType(tok.npos, p.typ())
}

// fnType:
// 	"func" '(' paramTypeListCommaOptOpt ')' result
func (p *parser) fnType() *FunctionType {
	tok := p.tok()
	p.n() // "func"
	p.must(gotoken.LPAREN)
	params := p.paramTypeListCommaOptOpt()
	p.must(gotoken.RPAREN)
	result := p.result()
	s := p.push()
	pos := p.pos()
	params.declare(p, s, pos)
	result.declare(p, s, pos)
	return newFunctionType(tok.npos, params, result)
}

// rxChanType:
// 	"<-" "chan" typ
func (p *parser) rxChanType() *ChannelType {
	tok := p.tok()
	p.n() // "<-"
	p.must(gotoken.CHAN)
	return newChannelType(tok.npos, RxChan, p.typ())
}

// typeList:
//		typ
//	|	typeList ',' typ
func (p *parser) typeList() (r []Type) {
	for r = []Type{p.typ()}; p.opt(gotoken.COMMA) && p.c != tokenGTGT; {
		r = append(r, p.typ())
	}
	return r
}

// genericArgsOpt:
//	|	"«" typeList commaOpt "»"
func (p *parser) genericArgsOpt() (r []Type) {
	if p.opt(tokenLTLT) {
		r = p.typeList()
		p.opt(gotoken.COMMA)
		p.must(tokenGTGT)
	}
	return r
}

// typ:
//		'(' typ ')'
//	|	qualifiedIdent genericArgsOpt
//	|	fnType
//	|	otherType
//	|	ptrType
//	|	rxChanType
func (p *parser) typ() (r Type) {
	switch ch := p.c; ch {
	case gotoken.LPAREN:
		p.n()
		r = p.typ()
		p.must(gotoken.RPAREN)
	case gotoken.IDENT:
		r = p.qualifiedIdent()
		p.genericArgsOpt()
	case gotoken.FUNC:
		r = p.fnType()
		p.pop()
	case gotoken.CHAN, gotoken.INTERFACE, gotoken.MAP, gotoken.STRUCT, gotoken.LBRACK:
		r = p.otherType(ch)
	case gotoken.MUL:
		r = p.ptrType()
	case gotoken.ARROW:
		r = p.rxChanType()
	default:
		p.syntaxError(p)
		p.err("syntax error: unexpected %v in type declaration", p.unexpected())
	}
	return r
}

//genericParamsOpt:
//	|	"«" identList "»"
func (p *parser) genericParamsOpt() (r []Token) {
	if p.opt(tokenLTLT) {
		r = p.identList()
		p.opt(gotoken.COMMA)
		p.must(tokenGTGT)
	}
	return r
}

// typeSpec:
//		IDENT genericParamsOpt typ
//	|	IDENT '=' typ
func (p *parser) typeSpec() Declaration {
	var d *TypeDecl
	if tok, ok := p.mustTok(gotoken.IDENT); ok {
		pos := token.NoPos
		if p.scope.Kind != PackageScope {
			pos = p.pos()
		}
		d = newTypeDecl(tok, pos)
	}
	switch p.c {
	case gotoken.ASSIGN:
		p.n()
		d.alias = true
	default:
		p.genericParamsOpt()
	}
	if d != nil {
		d.typ = p.typ()
		p.declare(d)
	}
	return d
}

// typeSpecList:
//		typeSpec
//	|	typeSpecList ';' typeSpec
func (p *parser) typeSpecList() (r []Declaration) {
	if s := p.typeSpec(); s != nil {
		r = append(r, s)
	}

	for p.opt(gotoken.SEMICOLON) && p.c != gotoken.RPAREN {
		if s := p.typeSpec(); s != nil {
			r = append(r, s)
		}
	}
	return r
}

// varSpec:
//		identList '=' exprList
//	|	identList typ
//	|	identList typ '=' exprList
func (p *parser) varSpec() (r []Declaration) {
	l := p.identList()
	var e []Expr
	var t Type

	defer func() {
		pos := token.NoPos
		if p.scope.Kind != PackageScope {
			pos = p.pos()
		}
		switch len(e) {
		case 0:
			for _, v := range l {
				d := newVarDecl(v, pos, false)
				d.setType(t)
				d.tupleIndex = -1
				p.declare(d)
				r = append(r, d)
			}
		case 1:
			for i, v := range l {
				d := newVarDecl(v, pos, false)
				d.Initializer = e[0]
				switch {
				case len(l) == 1:
					d.tupleIndex = -1
				default:
					d.tupleIndex = i
				}
				d.setType(t)
				p.declare(d)
				r = append(r, d)
			}
		default:
			if len(l) != len(e) {
				p.err("TODO2793")
				break
			}

			for i, v := range l {
				d := newVarDecl(v, pos, false)
				d.Initializer = e[i]
				d.setType(t)
				d.tupleIndex = -1
				p.declare(d)
				r = append(r, d)
			}
		}
	}()

	switch p.c {
	case gotoken.ASSIGN:
		p.n()
		e = p.exprList()
		return
	case gotoken.PERIOD:
		p.syntaxError(p)
		p.err("syntax error: unexpected %v, expecting type", p.unexpected())
		p.skip(gotoken.SEMICOLON, gotoken.RPAREN)
		return nil
	}

	t = p.typ()
	if p.opt(gotoken.ASSIGN) {
		e = p.exprList()
	}
	return nil
}

// varSpecList:
//		varSpec
//	|	varSpecList ';' varSpec
func (p *parser) varSpecList() (r []Declaration) {
	for r = p.varSpec(); p.opt(gotoken.SEMICOLON) && p.c != gotoken.RPAREN; {
		r = append(r, p.varSpec()...)
	}
	return r
}

// commonDecl:
//		"const" '(' ')'
//	|	"const" '(' constSpec ';' constSpecList semiOpt ')'
//	|	"const" '(' constSpec semiOpt ')'
//	|	"const" constSpec
//	|	"type" '(' ')'
//	|	"type" '(' typeSpecList semiOpt ')'
//	|	"type" typeSpec
//	|	"var" '(' ')'
//	|	"var" '(' varSpecList semiOpt ')'
//	|	"var" varSpec
func (p *parser) commonDecl() (r []Declaration) {
	switch p.c {
	case gotoken.CONST:
		p.n()
		switch {
		case p.opt(gotoken.LPAREN):
			if !p.opt(gotoken.RPAREN) {
				r = p.constSpecList()
				p.must(gotoken.RPAREN)
			}
		default:
			r = p.constSpec()
		}
	case gotoken.TYPE:
		p.n()
		switch {
		case p.opt(gotoken.LPAREN):
			if !p.opt(gotoken.RPAREN) {
				r = p.typeSpecList()
				p.must(gotoken.RPAREN)
			}
		default:
			r = []Declaration{p.typeSpec()}
		}
	case gotoken.VAR:
		p.n()
		switch {
		case p.opt(gotoken.LPAREN):
			if !p.opt(gotoken.RPAREN) {
				r = p.varSpecList()
				p.must(gotoken.RPAREN)
			}
		default:
			r = p.varSpec()
		}
	}
	return r
}

// ParamType represents data reduced by productions
//
//	paramType:
//		IDENT dddType
//	|	IDENT typ
//	|	dddType
//	|	typ
type ParamType interface{ isParamType() }

type paramType struct{}

func (paramType) isParamType() {}

// ParamTypeIdentType is a ParamType that includes the name of the parameter.
type ParamTypeIdentType struct {
	paramType

	Ident Token
	Type  Type

	Variadic bool
}

// ParamTypeType is a ParamType without a name of the parameter.
type ParamTypeType struct {
	paramType

	Type Type

	Variadic bool
}

func (p *parser) paramType() ParamType {
	switch p.c {
	case gotoken.IDENT:
		tok := p.tok()
		switch p.n() {
		case gotoken.RPAREN, gotoken.COMMA:
			return &ParamTypeType{Type: newNamedType(Ident{tok, p.scope})}
		case gotoken.PERIOD:
			p.n()
			ident, _ := p.mustTok(gotoken.IDENT)
			if !isExported(ident.Val) {
				//TODO p.err unexported
			}
			return &ParamTypeType{Type: newQualifiedNamedType(Ident{tok, p.scope}, ident.Val)}
		case tokenLTLT:
			p.genericArgsOpt()
			return &ParamTypeType{Type: newNamedType(Ident{tok, p.scope})}
		case gotoken.ELLIPSIS:
			p.n()
			return &ParamTypeIdentType{Ident: tok, Type: newSliceType(p.tok().npos, p.typ()), Variadic: true}
		default:
			return &ParamTypeIdentType{Ident: tok, Type: p.typ()}
		}
	case gotoken.ELLIPSIS:
		p.n()
		return &ParamTypeType{Type: newSliceType(p.tok().npos, p.typ()), Variadic: true}
	default:
		return &ParamTypeType{Type: p.typ()}
	}
}

// ParamTypeList represents data reduced by productions
//
//	paramTypeList:
//		paramType
//	|	paramTypeList ',' paramType
type ParamTypeList interface {
	declare(*parser, *Scope, token.Pos)
	isParamTypeList()
	len() int
	tuple() *TupleType
	types() string
}

type paramTypeList struct{}

func (paramTypeList) isParamTypeList() {}

// ParamTypeListType is a ParamTypeList of ParamTypeType
type ParamTypeListType struct {
	paramTypeList
	List []*ParamTypeType
}

func (n *ParamTypeListType) len() int                           { return len(n.List) }
func (n *ParamTypeListType) declare(*parser, *Scope, token.Pos) {}

func (n *ParamTypeListType) types() string {
	if n == nil {
		return ""
	}

	var a []string
	for _, v := range n.List {
		switch {
		case v.Variadic:
			a = append(a, "..."+v.Type.(*SliceType).Element.String())
		default:
			a = append(a, v.Type.String())
		}
	}
	return strings.Join(a, ", ")
}

func (n *ParamTypeListType) tuple() *TupleType {
	if n == nil {
		return emptyTuple
	}

	var l []Type
	for _, v := range n.List {
		//TODO if t := v.Type; t.state() == 0 { //TODO-
		//TODO 	panic(fmt.Errorf("%v: TODO3088 %T %v", t.Position(), t, t))
		//TODO }

		l = append(l, v.Type)
	}
	r := newTupleType(npos{}, l)
	r.checked()
	return r
}

// ParamTypeListIdentType is a ParamTypeList of ParamTypeIdentType
type ParamTypeListIdentType struct {
	paramTypeList
	List []*ParamTypeIdentType
}

func (n *ParamTypeListIdentType) len() int { return len(n.List) }

func (n *ParamTypeListIdentType) declare(p *parser, s *Scope, visibility token.Pos) {
	for _, v := range n.List {
		d := newVarDecl(v.Ident, visibility, true)
		d.setType(v.Type)
		s.declare(p.sourceFile, d)
	}
}

func (n *ParamTypeListIdentType) types() string {
	if n == nil {
		return ""
	}

	var a []string
	for _, v := range n.List {
		switch {
		case v.Variadic:
			a = append(a, "..."+v.Type.(*SliceType).Element.String())
		default:
			a = append(a, v.Type.String())
		}
	}
	return strings.Join(a, ", ")
}

func (n *ParamTypeListIdentType) tuple() *TupleType {
	if n == nil {
		return emptyTuple
	}

	var l []Type
	for _, v := range n.List {
		//TODO if v.Type.state() == 0 {
		//TODO 	panic(fmt.Errorf("%v: TODO3139", v.Type.Position()))
		//TODO }

		l = append(l, v.Type)
	}
	r := newTupleType(npos{}, l)
	r.checked()
	return r
}

func (p *parser) paramTypeList() ParamTypeList {
	list := []ParamType{p.paramType()}
	for p.opt(gotoken.COMMA) && p.c != gotoken.RPAREN {
		list = append(list, p.paramType())
	}
	var names []*ParamTypeIdentType
	var types []*ParamTypeType
	for _, v := range list {
		switch x := v.(type) {
		case *ParamTypeIdentType:
			for _, w := range types {
				switch y := w.Type.(type) {
				case *NamedType:
					z := *x
					z.Ident = Token{Val: y.Name, npos: y.npos}
					names = append(names, &z)
					continue
				}

				p.err("syntax error: mixed named and unnamed function parameters")
			}
			names = append(names, x)
			types = types[:0]
		case *ParamTypeType:
			types = append(types, x)
		default:
			panic(fmt.Errorf("internal error 022: %T", v))
		}
	}
	if len(names) != 0 {
		if len(types) != 0 {
			p.err("syntax error: mixed named and unnamed function parameters")
		}
		return &ParamTypeListIdentType{List: names}
	}

	if len(types) != 0 {
		return &ParamTypeListType{List: types}
	}

	panic("internal error 023")
}

// paramTypeListCommaOptOpt:
//	|	paramTypeList commaOpt
func (p *parser) paramTypeListCommaOptOpt() ParamTypeList {
	if p.c != gotoken.RPAREN {
		return p.paramTypeList()
	}

	return emptyParamTypeList
}

// result:
//		%prec _NotParen
//	|	'(' paramTypeListCommaOptOpt ')'
//	|	qualifiedIdent genericArgsOpt
//	|	fnType
//	|	otherType
//	|	ptrType
//	|	rxChanType
func (p *parser) result() (r ParamTypeList) {
	r = emptyParamTypeList
	switch ch := p.c; ch {
	case gotoken.LBRACE, gotoken.RPAREN, gotoken.SEMICOLON, gotoken.COMMA, tokenBODY,
		gotoken.RBRACE, gotoken.COLON, gotoken.STRING, gotoken.ASSIGN, gotoken.RBRACK:
	case gotoken.LPAREN:
		p.n()
		r = p.paramTypeListCommaOptOpt()
		p.must(gotoken.RPAREN)
	case gotoken.IDENT:
		t := p.qualifiedIdent()
		p.genericArgsOpt()
		r = &ParamTypeListType{List: []*ParamTypeType{{Type: t}}}
	case gotoken.FUNC:
		t := p.fnType()
		p.pop()
		r = &ParamTypeListType{List: []*ParamTypeType{{Type: t}}}
	case gotoken.CHAN, gotoken.INTERFACE, gotoken.MAP, gotoken.STRUCT, gotoken.LBRACK:
		r = &ParamTypeListType{List: []*ParamTypeType{{Type: p.otherType(ch)}}}
	case gotoken.MUL:
		r = &ParamTypeListType{List: []*ParamTypeType{{Type: p.ptrType()}}}
	case gotoken.ARROW:
		r = &ParamTypeListType{List: []*ParamTypeType{{Type: p.rxChanType()}}}
	default:
		p.syntaxError(p)
		//TODO p.err()
	}
	return r
}

func (p *parser) shortVarDecl(visibility token.Pos, lhs, rhs []Expr) {
	new := false
	switch len(rhs) {
	case 0:
		p.err("TODO3167")
	case 1:
		for i, v := range lhs {
			x, ok := v.(*PrimaryExprIdent)
			if !ok {
				if !p.sourceFile.Package.ctx.tweaks.ignoreRedeclarations {
					p.err0(v.Position(), "invalid variable name")
				}
				continue
			}

			tok := x.Ident.Token
			if _, ok := p.scope.Bindings[tok.Val]; ok {
				continue
			}

			new = true
			d := newVarDecl(tok, visibility, false)
			d.Initializer = rhs[0]
			switch {
			case len(lhs) == 1:
				d.tupleIndex = -1
			default:
				d.tupleIndex = i
			}
			p.declare(d)
		}
	default:
		if len(lhs) != len(rhs) {
			p.err("TODO3190")
		}

		for i, v := range lhs {
			x, ok := v.(*PrimaryExprIdent)
			if !ok {
				if !p.sourceFile.Package.ctx.tweaks.ignoreRedeclarations {
					p.err0(v.Position(), "invalid variable name")
				}
				continue
			}

			tok := x.Ident.Token
			if _, ok := p.scope.Bindings[tok.Val]; ok {
				continue
			}

			new = true
			d := newVarDecl(tok, visibility, false)
			d.tupleIndex = -1
			if i < len(rhs) {
				d.Initializer = rhs[i]
			}
			p.declare(d)
		}
	}
	if !new {
		//TODO p.err("TODO3217") // no new vars left of :=
	}
}

// SimpleStmt represents data reduced by productions
//
//	simpleStmt:
//		expr
//	|	expr "%=" expr
//	|	expr "&=" expr
//	|	expr "&^=" expr
//	|	expr "*=" expr
//	|	expr "++"
//	|	expr "+=" expr
//	|	expr "--"
//	|	expr "-=" expr
//	|	expr "/=" expr
//	|	expr "<<=" expr
//	|	expr ">>=" expr
//	|	expr "^=" expr
//	|	expr "|=" expr
//	|	exprList ":=" exprList
//	|	exprList '=' exprList
type SimpleStmt Stmt

// SimpleStmtDecl describes a list of declarations.
type SimpleStmtDecl struct {
	stmt
	List []Declaration
}

func (n *SimpleStmtDecl) check(s *cstack) {
	for _, v := range n.List {
		s.check(v, n)
	}
}

// SimpleStmtExpr describes an expresssion statement.
type SimpleStmtExpr struct {
	stmt
	Expr Expr
}

func (n *SimpleStmtExpr) check(s *cstack) {
	s.check(n.Expr, n)
}

// SimpleStmtRemAssign describes expr %= expr.
type SimpleStmtRemAssign struct {
	stmt
	LHS Expr
	RHS Expr
}

func (n *SimpleStmtRemAssign) check(s *cstack) {
	//TODO
}

// SimpleStmtAndAssign describes expr &= expr.
type SimpleStmtAndAssign struct {
	stmt
	LHS Expr
	RHS Expr
}

func (n *SimpleStmtAndAssign) check(s *cstack) {
	//TODO
}

// SimpleStmtAndNotAssign describes expr &^= expr.
type SimpleStmtAndNotAssign struct {
	stmt
	LHS Expr
	RHS Expr
}

func (n *SimpleStmtAndNotAssign) check(s *cstack) {
	//TODO
}

// SimpleStmtMulAssign describes expr *= expr.
type SimpleStmtMulAssign struct {
	stmt
	LHS Expr
	RHS Expr
}

func (n *SimpleStmtMulAssign) check(s *cstack) {
	//TODO
}

// SimpleStmtInc describes expr++.
type SimpleStmtInc struct {
	stmt
	Expr Expr
}

func (n *SimpleStmtInc) check(s *cstack) {
	//TODO
}

// SimpleStmtAddAssign describes expr += expr.
type SimpleStmtAddAssign struct {
	stmt
	LHS Expr
	RHS Expr
}

func (n *SimpleStmtAddAssign) check(s *cstack) {
	//TODO
}

// SimpleStmtDec describes expr--.
type SimpleStmtDec struct {
	stmt
	Expr Expr
}

func (n *SimpleStmtDec) check(s *cstack) {
	//TODO
}

// SimpleStmtSubAssign describes expr -= expr.
type SimpleStmtSubAssign struct {
	stmt
	LHS Expr
	RHS Expr
}

func (n *SimpleStmtSubAssign) check(s *cstack) {
	//TODO
}

// SimpleStmtQuoAssign describes expr /= expr.
type SimpleStmtQuoAssign struct {
	stmt
	LHS Expr
	RHS Expr
}

func (n *SimpleStmtQuoAssign) check(s *cstack) {
	//TODO
}

// SimpleStmtShlAssign describes expr <<= expr.
type SimpleStmtShlAssign struct {
	stmt
	LHS Expr
	RHS Expr
}

func (n *SimpleStmtShlAssign) check(s *cstack) {
	//TODO
}

// SimpleStmtShrAssign describes expr >>= expr.
type SimpleStmtShrAssign struct {
	stmt
	LHS Expr
	RHS Expr
}

func (n *SimpleStmtShrAssign) check(s *cstack) {
	//TODO
}

// SimpleStmtXorAssign describes expr ^= expr.
type SimpleStmtXorAssign struct {
	stmt
	LHS Expr
	RHS Expr
}

func (n *SimpleStmtXorAssign) check(s *cstack) {
	//TODO
}

// SimpleStmtOrAssign describes expr |= expr.
type SimpleStmtOrAssign struct {
	stmt
	LHS Expr
	RHS Expr
}

func (n *SimpleStmtOrAssign) check(s *cstack) {
	//TODO
}

// SimpleStmtDefine describes exprList := exprList.
type SimpleStmtDefine struct {
	stmt
	LHS []Expr
	RHS []Expr
}

func (n *SimpleStmtDefine) check(s *cstack) {
	//TODO m := SimpleStmtAssignment{stmt: n.stmt}
	//TODO m.check2(s, n.LHS, n.RHS)
}

// SimpleStmtAssignment describes exprList = exprList.
type SimpleStmtAssignment struct {
	stmt
	LHS []Expr
	RHS []Expr
}

func (n *SimpleStmtAssignment) check(s *cstack) {
	n.check2(s, n.LHS, n.RHS)
}

func (n *SimpleStmtAssignment) check2(s *cstack, lhs, rhs []Expr) {
	for _, v := range lhs {
		if !skipBlank(v) {
			s.check(v, n)
		}
	}
	for _, v := range rhs {
		s.check(v, n)
	}
	switch len(rhs) {
	case 0:
		// nop
	case 1:
		r := rhs[0]
		switch t := r.Type().(type) {
		case nil:
			// nop
		case *TupleType:
			if g, e := len(lhs), len(t.List); g != e {
				n.err("assignment mismatch: %d variables but %d values", g, e)
				break
			}

			for i, l := range lhs {
				if t := t.List[i]; t != nil && t != Invalid && t.UnderlyingType() != Invalid {
					if l != nil && !skipBlank(l) {
						if v := l.Type(); v != nil && v != Invalid && t != Invalid && !t.AssignableTo(v) {
							n.err0(r.Position(), "cannot assign %s to %s (type %s) in multiple assignment", t, l.str(), v)
							break
						}
					}
				}
			}
		default:
			if t == Invalid {
				break
			}

			if len(lhs) != 1 {
				n.err("assignment mismatch: more than 1 variable but only 1 value")
			}

			if l := lhs[0]; l != nil && !skipBlank(l) {
				if v := l.Type(); v != nil && v != Invalid && !t.AssignableTo(v) {
					// dbg("%v:", n.Position())
					// dbg("%v: %T(%s)", t.Position(), t, t)
					// dbg("%v: %T(%s)", v.Position(), v, v)
					n.err0(r.Position(), "cannot use %s (type %s) as type %s in assignment", r.str(), t, v)
				}
			}
		}
	default:
		if len(lhs) != len(rhs) {
			n.err("TODO3517")
			break
		}

		for i, l := range lhs {
			if r := rhs[i]; r != nil {
				if l != nil && !skipBlank(l) {
					if t, v := r.Type(), l.Type(); t != nil && v != nil && v != Invalid && t != Invalid && !t.AssignableTo(v) {
						panic("TODO3529")
					}
				}
			}
		}
	}
}

// SimpleStmtSend describes expr <- expr.
type SimpleStmtSend struct {
	stmt
	LHS Expr
	RHS Expr
}

func (n *SimpleStmtSend) check(s *cstack) {
	//TODO
}

func (p *parser) simpleStmt(acceptRange bool) (label *PrimaryExprIdent, isRange bool, r SimpleStmt) {
	first := true
	e := p.expr()
	switch x := e.(type) {
	case *PrimaryExprIdent:
		if p.c == gotoken.COLON {
			return x, false, nil
		}
	case *ExprARROW:
		return nil, false, &SimpleStmtSend{LHS: x.LHS, RHS: x.RHS}
	}

	lhs := []Expr{e}
more:
	switch pc := p.c; pc {
	case gotoken.REM_ASSIGN:
		p.n()
		return nil, false, &SimpleStmtRemAssign{LHS: e, RHS: p.expr()}
	case gotoken.AND_ASSIGN:
		p.n()
		return nil, false, &SimpleStmtAndAssign{LHS: e, RHS: p.expr()}
	case gotoken.AND_NOT_ASSIGN:
		p.n()
		return nil, false, &SimpleStmtAndNotAssign{LHS: e, RHS: p.expr()}
	case gotoken.MUL_ASSIGN:
		p.n()
		return nil, false, &SimpleStmtMulAssign{LHS: e, RHS: p.expr()}
	case gotoken.ADD_ASSIGN:
		p.n()
		return nil, false, &SimpleStmtAddAssign{LHS: e, RHS: p.expr()}
	case gotoken.SUB_ASSIGN:
		p.n()
		return nil, false, &SimpleStmtSubAssign{LHS: e, RHS: p.expr()}
	case gotoken.QUO_ASSIGN:
		p.n()
		return nil, false, &SimpleStmtQuoAssign{LHS: e, RHS: p.expr()}
	case gotoken.SHL_ASSIGN:
		p.n()
		return nil, false, &SimpleStmtShlAssign{LHS: e, RHS: p.expr()}
	case gotoken.SHR_ASSIGN:
		p.n()
		return nil, false, &SimpleStmtShrAssign{LHS: e, RHS: p.expr()}
	case gotoken.XOR_ASSIGN:
		p.n()
		return nil, false, &SimpleStmtXorAssign{LHS: e, RHS: p.expr()}
	case gotoken.OR_ASSIGN:
		p.n()
		return nil, false, &SimpleStmtOrAssign{LHS: e, RHS: p.expr()}
	case gotoken.INC:
		p.n()
		return nil, false, &SimpleStmtInc{Expr: e}
	case gotoken.DEC:
		p.n()
		return nil, false, &SimpleStmtDec{Expr: e}
	case gotoken.COMMA:
		if !first {
			p.syntaxError(p)
			//TODO p.err()
			break
		}

		first = false
		p.n()
		lhs = append(lhs, p.exprList()...)
		goto more
	case gotoken.DEFINE, gotoken.ASSIGN:
		tok := p.tok()
		p.n()
		if acceptRange && p.opt(gotoken.RANGE) {
			isRange = true
		}
		rhs := p.exprList()
		switch {
		case pc == gotoken.DEFINE:
			p.shortVarDecl(p.pos(), lhs, rhs)
			return nil, isRange, &SimpleStmtDefine{LHS: lhs, RHS: rhs}
		default:
			return nil, isRange, &SimpleStmtAssignment{stmt: stmt{npos: newPos(tok)}, LHS: lhs, RHS: rhs}
		}
	}
	return nil, false, &SimpleStmtExpr{Expr: e}
}

// simpleStmtOpt:
//	|	simpleStmt
func (p *parser) simpleStmtOpt(acceptRange bool) (isRange bool, r SimpleStmt) {
	if p.c == gotoken.SEMICOLON || p.c == tokenBODY {
		return false, nil
	}

	label, isRange, r := p.simpleStmt(acceptRange)
	if label != nil {
		//TODO p.err()
	}
	return isRange, r
}

// IfHeader represents data reduced by productions
//
//	ifHeader:
//		simpleStmtOpt
//	|	simpleStmtOpt ';' simpleStmtOpt
type IfHeader interface {
	checker
	isIfHeader()
}

type ifHeader struct {
	sentinel
}

func (*ifHeader) isIfHeader() {}

// IfHeader1 describes if foo() ...
type IfHeader1 struct {
	ifHeader
	Stmt SimpleStmt
}

func (n *IfHeader1) check(s *cstack) {
	//TODO
}

// IfHeader2 describes if x := foo(); x > 3 ...
type IfHeader2 struct {
	ifHeader
	Stmt  SimpleStmt
	Stmt2 SimpleStmt
}

func (n *IfHeader2) check(s *cstack) {
	//TODO
}

func (p *parser) ifHeader(s string) (r IfHeader) {
	_, s1 := p.simpleStmtOpt(false)
	switch {
	case p.opt(gotoken.SEMICOLON):
		_, s2 := p.simpleStmtOpt(false)
		r = &IfHeader2{Stmt: s1, Stmt2: s2}
	default:
		r = &IfHeader1{Stmt: s1}
	}
	if p.c == gotoken.SEMICOLON {
		p.syntaxError(p)
		switch {
		case s != "":
			p.err(s)
		default:
			p.err("syntax error: unexpected %v, expecting { after if clause", p.unexpected())
		}
		p.skip(gotoken.LBRACE)
	}
	return r
}

// loopBody:
// 	BODY stmtList '}'
func (p *parser) loopBody() *StmtBlock {
	p.must(tokenBODY)
	s := p.push()
	l := p.stmtList()
	p.must(gotoken.RBRACE)
	p.pop()
	return &StmtBlock{s: s, List: l}
}

// ElseIf represents data reduced by productions
//
//	elseIfList:
//	|	elseIfList "else" "if" ifHeader loopBody
type ElseIf struct {
	IfHeader
	Body *StmtBlock
}

func (n *ElseIf) check(s *cstack) {
	//TODO s.check(n.IfHeader, n)
	if n.Body != nil {
		s.check(n.Body, n)
	}
}

func (p *parser) elseIfList() (isElse bool, r []ElseIf) {
	for p.opt(gotoken.ELSE) {
		if p.opt(gotoken.IF) {
			h := p.ifHeader("")
			b := p.loopBody()
			r = append(r, ElseIf{h, b})
			continue
		}

		return true, r // Consumed "else", "if" does not follow.
	}
	return false, r
}

//	compoundStmt:
//		'{' stmtList '}'
func (p *parser) compoundStmt(ch gotoken.Token) *StmtBlock {
	r := &StmtBlock{}
	switch p.c {
	case gotoken.LBRACE:
		p.n()
		r.s = p.push()
		r.List = p.stmtList()
		p.must(gotoken.RBRACE)
		p.pop()
	case gotoken.SEMICOLON:
		p.syntaxError(p)
		switch ch {
		case gotoken.ELSE:
			p.err("syntax error: else must be followed by if or statement block")
		default:
			//TODO p.err()
		}
	default:
		p.syntaxError(p)
		//TODO p.err()
		p.n()
	}
	return r
}

// CaseBlock represents data reduced by productions
//
//	caseBlockList:
//	|	caseBlockList "case" exprOrTypeList ":=" expr ':' stmtList
//	|	caseBlockList "case" exprOrTypeList ':' stmtList
//	|	caseBlockList "case" exprOrTypeList '=' expr ':' stmtList
//	|	caseBlockList "default" ':' stmtList
type CaseBlock struct {
	ExprOrTypeList []ExprOrType
	Expr           Expr
	Body           []Stmt //TODO StmtBlock

	Default bool // "default" clause
	Define  bool // true: case x := y; false: case y or case x = y.

	sentinel
}

func (n *CaseBlock) check(s *cstack) {
	//TODO for _, v := range n.ExprOrTypeList {
	//TODO 	s.check(v, n)
	//TODO }
	s.check(n.Expr, n)
	//TODO if n.Body != nil {
	//TODO 	s.check(n.Body, n)
	//TODO }
}

func (p *parser) caseBlockList() (r []CaseBlock) {
	for {
		p.push()
		switch p.c {
		case gotoken.CASE:
			p.n()
			n := CaseBlock{ExprOrTypeList: p.exprOrTypeList()}
			n.Define = p.c == gotoken.DEFINE
			if n.Define || p.c == gotoken.ASSIGN {
				p.n()
				n.Expr = p.expr()
			}
			p.must(gotoken.COLON)
			if n.Define {
				pos := p.pos()
				var lhs []Expr
				for _, v := range n.ExprOrTypeList {
					switch x := v.(type) {
					case *ExprOrTypeExpr:
						lhs = append(lhs, x.Expr)
					case
						*ExprOrTypeType,
						*exprOrType:

						p.err0(v.Position(), "TODO3714")
						lhs = append(lhs, nil)
					default:
						panic(fmt.Errorf("unexpected type %T", x))
					}
				}
				p.shortVarDecl(pos, lhs, []Expr{n.Expr})
			}
			r = append(r, n)
		case gotoken.DEFAULT:
			p.n()
			p.must(gotoken.COLON)
			r = append(r, CaseBlock{Default: true})
		case gotoken.IF:
			p.syntaxError(p)
			p.err("syntax error: unexpected if, expecting case or default or }")
			p.skip(gotoken.COLON)
		default:
			p.pop()
			if p.c != gotoken.RBRACE {
				p.syntaxError(p)
				//TODO p.err()
				p.skip(gotoken.RBRACE)
			}
			return
		}

		if n := len(r); n > 0 {
			r[n-1].Body = p.stmtList()
		}
		p.pop()
	}
}

// Stmt represents data reduced by productions
//
//	stmt:
//	|	"break" identOpt
//	|	"continue" identOpt
//	|	"defer" primaryExpr
//	|	"fallthrough"
//	|	"for" "range" expr loopBody
//	|	"for" exprList ":=" "range" expr loopBody
//	|	"for" exprList '=' "range" expr loopBody
//	|	"for" simpleStmtOpt ';' simpleStmtOpt ';' simpleStmtOpt loopBody
//	|	"for" simpleStmtOpt loopBody
//	|	"go" primaryExpr
//	|	"goto" IDENT
//	|	"if" ifHeader loopBody elseIfList
//	|	"if" ifHeader loopBody elseIfList "else" compoundStmt
//	|	"return"
//	|	"return" exprList
//	|	"select" BODY caseBlockList '}'
//	|	"switch" ifHeader BODY caseBlockList '}'
//	|	IDENT ':' stmt
//	|	commonDecl
//	|	compoundStmt
//	|	simpleStmt
type Stmt interface {
	checker
	isStmt()
}

type stmt struct {
	sentinel
	npos
}

func (n *stmt) isStmt() {}

// StmtBreak describes break label.
type StmtBreak struct {
	stmt
	Label Token
}

func (n *StmtBreak) check(s *cstack) {
	//TODO
}

// StmtContinue describes continue label.
type StmtContinue struct {
	stmt
	Label Token
}

func (n *StmtContinue) check(s *cstack) {
	//TODO
}

// StmtDefer describes defer expr.
type StmtDefer struct {
	stmt
	Expr Expr
}

func (n *StmtDefer) check(s *cstack) {
	s.check(n.Expr, n)
}

// StmtFallthrough describes falltrough.
type StmtFallthrough struct {
	stmt
}

func (n *StmtFallthrough) check(s *cstack) {
	//TODO
}

// StmtForRange describes for i := range expr { body } or the = variant.
type StmtForRange struct {
	ExprList []Expr
	Expr     Expr
	Body     *StmtBlock

	stmt
	Define bool // true: for exprlist := range ..., false: for exprlist = range ...
}

func (n *StmtForRange) check(s *cstack) {
	//TODO for _, v := range n.ExprList {
	//TODO 	s.check(v, n)
	//TODO }
	//TODO s.check(n.Expr, n)
	//TODO if n.Body != nil {
	//TODO 	s.check(n.Body, n)
	//TODO }
}

// StmtFor describes for stmt; stmt2; stmt3 { body }.
type StmtFor struct {
	stmt
	Stmt  SimpleStmt
	Stmt2 SimpleStmt
	Stmt3 SimpleStmt
	Body  *StmtBlock
}

func (n *StmtFor) check(s *cstack) {
	s.check(n.Stmt, n)
	s.check(n.Stmt2, n)
	s.check(n.Stmt3, n)
	if n.Body != nil {
		s.check(n.Body, n)
	}
}

// StmtGo describes go expr.
type StmtGo struct {
	stmt
	Expr Expr
}

func (n *StmtGo) check(s *cstack) {
	s.check(n.Expr, n)
}

// StmtGoto describes goto ident.
type StmtGoto struct {
	stmt
	Ident Token
}

func (n *StmtGoto) check(s *cstack) {
	//TODO
}

// StmtIf descibes if ifHeader body elseIfList else body2
type StmtIf struct {
	stmt
	IfHeader IfHeader
	Body     *StmtBlock
	ElseIf   []ElseIf
	Body2    *StmtBlock
}

func (n *StmtIf) check(s *cstack) {
	//TODO s.check(n.IfHeader, n)
	if n.Body != nil {
		s.check(n.Body, n)
	}
	for i := range n.ElseIf {
		s.check(&n.ElseIf[i], n)
	}
	if n.Body2 != nil {
		s.check(n.Body2, n)
	}
}

// StmtReturn describes return expr.
type StmtReturn struct {
	stmt
	ExprList []Expr
}

func (n *StmtReturn) check(s *cstack) {
	for _, v := range n.ExprList {
		s.check(v, n)
	}
}

// StmtSelect describes select {}.
type StmtSelect struct {
	stmt
	List []CaseBlock
}

func (n *StmtSelect) check(s *cstack) {
	for i := range n.List {
		s.check(&n.List[i], n)
	}
}

// StmtSwitch describes switch {}.
type StmtSwitch struct {
	stmt
	IfHeader IfHeader
	List     []CaseBlock
}

func (n *StmtSwitch) check(s *cstack) {
	//TODO s.check(n.IfHeader, n)
	for i := range n.List {
		s.check(&n.List[i], n)
	}
}

// StmtLabeled describes label: stmt.
type StmtLabeled struct {
	stmt
	Label Token
	Stmt  Stmt
}

func (n *StmtLabeled) check(s *cstack) {
	s.check(n.Stmt, n)
}

// StmtBlock describes { stmt; stmt2; ... }.
type StmtBlock struct {
	stmt
	List []Stmt
	s    *Scope
}

// Scope retruns the scope of the statement block.
func (n *StmtBlock) Scope() *Scope { return n.s }

func (n *StmtBlock) check(s *cstack) {
	var a []string
	if n.s != nil {
		for k := range n.s.Bindings {
			a = append(a, k)
		}
		sort.Strings(a)
		for _, k := range a {
			s.check(n.s.Bindings[k], nil)
		}
		for _, v := range n.s.Unbound {
			s.check(v, nil)
		}
	}
	for _, v := range n.List {
		s.check(v, nil)
	}
}

func (p *parser) stmt() (r Stmt) {
	switch ch := p.c; ch {
	case gotoken.SEMICOLON, gotoken.RBRACE, gotoken.CASE, gotoken.DEFAULT:
		// nop
	case gotoken.BREAK:
		p.n()
		tok := p.tok()
		switch {
		case p.opt(gotoken.IDENT):
			r = &StmtBreak{Label: tok}
		default:
			r = &StmtBreak{}
		}
	case gotoken.CONTINUE:
		p.n()
		tok := p.tok()
		switch {
		case p.opt(gotoken.IDENT):
			r = &StmtContinue{Label: tok}
		default:
			r = &StmtContinue{}
		}
	case gotoken.DEFER:
		p.n()
		e := p.expr()
		r = &StmtDefer{Expr: e}
		if _, ok := e.(*PrimaryExprCall); !ok {
			if _, ok := e.(*PrimaryExprParenExpr); ok {
				p.err("expression in defer must not be parenthesized")
				break
			}

			p.err("expression in defer must be function call")
		}
	case gotoken.GO:
		p.n()
		e := p.expr()
		r = &StmtGo{Expr: e}
		if _, ok := e.(*PrimaryExprCall); !ok {
			if _, ok := e.(*PrimaryExprParenExpr); ok {
				p.err("expression in go must not be parenthesized")
				break
			}
			p.err("expression in go must be function call")
		}
	case gotoken.FALLTHROUGH:
		r = &StmtFallthrough{}
		p.n()
	case gotoken.FOR:
		p.push()
		switch p.n() {
		case gotoken.RANGE:
			p.n()
			r = &StmtForRange{Expr: p.expr()}
		case gotoken.VAR:
			p.err("syntax error: var declaration not allowed in for initializer")
			p.n()
			fallthrough
		default:
			isRange, s := p.simpleStmtOpt(true)
			if isRange {
				switch x := s.(type) {
				case *SimpleStmtDefine:
					if len(x.RHS) != 1 {
						//TODO p.err()
						break
					}
					r = &StmtForRange{ExprList: x.LHS, Expr: x.RHS[0], Define: true}
				case *SimpleStmtAssignment:
					if len(x.RHS) != 1 {
						//TODO p.err()
						break
					}
					r = &StmtForRange{ExprList: x.LHS, Expr: x.RHS[0]}
				default:
					panic("internal error 024")
				}
				break
			}

			var s2, s3 SimpleStmt
			if p.opt(gotoken.SEMICOLON) {
				if p.c == tokenBODY {
					p.err("unexpected {, expecting for loop condition")
					break
				}

				_, s2 = p.simpleStmtOpt(false)
				p.must(gotoken.SEMICOLON)
				_, s3 = p.simpleStmtOpt(false)
			}
			r = &StmtFor{Stmt: s, Stmt2: s2, Stmt3: s3}
		}
		if p.c == gotoken.SEMICOLON {
			p.syntaxError(p)
			p.skip(tokenBODY)
		}
		l := p.loopBody()
		switch x := r.(type) {
		case *StmtForRange:
			x.Body = l
		case *StmtFor:
			x.Body = l
		}
		p.pop()
	case gotoken.GOTO:
		p.n()
		r = &StmtGoto{Ident: p.tok()}
		p.must(gotoken.IDENT)
	case gotoken.IF:
		p.n()
		p.push()
		if p.c == gotoken.VAR {
			p.err("syntax error: var declaration not allowed in if initializer")
			p.n()
		}
		s := &StmtIf{IfHeader: p.ifHeader("")}
		s.Body = p.loopBody()
		switch ok, l := p.elseIfList(); {
		case ok:
			s.ElseIf = l
			s.Body2 = p.compoundStmt(gotoken.ELSE)
		default:
			s.ElseIf = l
		}
		r = s
		p.pop()
	case gotoken.RETURN:
		p.n()
		s := &StmtReturn{}
		if p.not2(gotoken.SEMICOLON, gotoken.RBRACE) {
			s.ExprList = p.exprList()
			if p.not2(gotoken.SEMICOLON, gotoken.RBRACE) {
				p.syntaxError(p)
				p.skip(gotoken.SEMICOLON, gotoken.RBRACE)
			}
		}
		r = s
	case gotoken.SELECT:
		p.n()
		p.must(tokenBODY)
		p.push()
		r = &StmtSelect{List: p.caseBlockList()}
		p.must(gotoken.RBRACE)
		p.pop()
	case gotoken.SWITCH:
		p.n()
		p.push()
		if p.c == gotoken.VAR {
			p.err("syntax error: var declaration not allowed in switch initializer")
			p.n()
		}
		s := &StmtSwitch{IfHeader: p.ifHeader("syntax error: missing { after switch clause")}
		p.must(tokenBODY)
		p.push()
		s.List = p.caseBlockList()
		r = s
		p.must(gotoken.RBRACE)
		p.pop()
		p.pop()
	case gotoken.CONST, gotoken.TYPE, gotoken.VAR:
		r = &SimpleStmtDecl{List: p.commonDecl()}
	case gotoken.LBRACE:
		r = p.compoundStmt(gotoken.Token(-1))
	default:
		label, _, s := p.simpleStmt(false)
		switch {
		case label != nil && p.opt(gotoken.COLON):
			r = &StmtLabeled{Label: label.Ident.Token, Stmt: p.stmt()}
		default:
			r = s
		}
	}
	if p.c == gotoken.COLON {
		p.syntaxError(p)
		p.n()
	}
	return r
}

// stmtList:
// 	stmt
// |	stmtList ';' stmt
func (p *parser) stmtList() (r []Stmt) {
	if s := p.stmt(); s != nil {
		r = []Stmt{s}
	}
	for p.opt(gotoken.SEMICOLON) && p.not2(gotoken.RBRACE, gotoken.CASE, gotoken.DEFAULT) {
		if p.c == gotoken.ELSE {
			p.err("syntax error: unexpected else, expecting }")
			if ok, _ := p.elseIfList(); ok {
				p.compoundStmt(gotoken.ELSE)
			}
			continue
		}

		if s := p.stmt(); s != nil {
			r = append(r, s)
		}
	}
	if p.c == gotoken.LBRACE {
		p.err("syntax error: unexpected { at end of statement")
		p.skip(gotoken.RBRACE)
	}
	return r
}

// fnBody:
//	|	'{' stmtList '}'
func (p *parser) fnBody() (r []Stmt, ok bool) {
	if p.opt(gotoken.LBRACE) {
		ok = true
		r = p.stmtList()
		p.must(gotoken.RBRACE)
	}
	p.pop()
	return r, ok
}

// topLevelDeclList:
//	|	topLevelDeclList "func" '(' paramTypeListCommaOptOpt ')' IDENT genericParamsOpt '(' paramTypeListCommaOptOpt ')' result fnBody ';'
//	|	topLevelDeclList "func" IDENT genericParamsOpt '(' paramTypeListCommaOptOpt ')' result fnBody ';'
//	|	topLevelDeclList commonDecl ';'
func (p *parser) topLevelDeclList() {
	ps := p.sourceFile.Package.Scope
	for p.c != gotoken.EOF {
		p.scope = ps
		switch p.c {
		case gotoken.FUNC:
			tok := p.tok()
			var d Declaration
			rcvr := emptyParamTypeList
			switch p.n() {
			case gotoken.IDENT:
				tok := p.tok()
				d = newFuncDecl(p.tok(), token.NoPos)
				switch {
				case tok.Val == "init":
					p.sourceFile.InitFunctions = append(p.sourceFile.InitFunctions, d)
				default:
					ps.declare(p.sourceFile, d)
				}
				p.n()
				p.genericParamsOpt()
				switch p.c {
				case gotoken.LPAREN:
					p.n()
				default:
					p.syntaxError(p)
				}
			case gotoken.LPAREN:
				p.n()
				rcvr = p.paramTypeListCommaOptOpt()
				p.must(gotoken.RPAREN)
				var rx Type
				switch rcvr.len() {
				case 0:
					p.err("method has no receiver")
				case 1:
					switch x := rcvr.(type) {
					case *ParamTypeListIdentType:
						rx = x.List[0].Type
					case *ParamTypeListType:
						rx = x.List[0].Type
					default:
						panic("internal error 034")
					}
				default:
					p.err("method has multiple receivers")
				}
				if tok, ok := p.mustTok(gotoken.IDENT); ok {
					md := newMethodDecl(tok, token.NoPos)
					md.Receiver = rx
					d = md
				}
				p.genericParamsOpt()
				p.must2(gotoken.LPAREN)
			default:
				p.syntaxError(p)
			}

			params := p.paramTypeListCommaOptOpt()
			if p.c == gotoken.SEMICOLON {
				p.syntaxError(p)
				p.skip(gotoken.RPAREN)
			}
			p.must(gotoken.RPAREN)
			result := p.result()
			s := p.push()
			pos := p.pos()
			rcvr.declare(p, s, pos)
			params.declare(p, s, pos)
			result.declare(p, s, pos)
			var block *StmtBlock
			if body, ok := p.fnBody(); ok {
				block = &StmtBlock{s: s, List: body}
			}
			switch x := d.(type) {
			case nil:
				// nop
			case *FuncDecl:
				x.declaration.typ = newFunctionType(tok.npos, params, result)
				x.Body = block
			case *MethodDecl:
				x.declaration.typ = newFunctionType(tok.npos, params, result)
				x.Body = block
				p.sourceFile.TopLevelDecls = append(p.sourceFile.TopLevelDecls, x)
			default:
				panic(fmt.Errorf("internal error 025: %T", x))
			}
			switch p.c {
			case gotoken.SEMICOLON:
				if p.n() != gotoken.LBRACE {
					continue
				}

				p.err("syntax error: unexpected semicolon or newline before {")
				p.skip(gotoken.SEMICOLON)
			}
		case gotoken.CONST, gotoken.TYPE, gotoken.VAR:
			p.commonDecl()
		default:
			p.syntaxError(p)
			if s := p.stmt(); s != nil {
				p.err("syntax error: non-declaration statement outside function body")
			}
		}
		if p.c == gotoken.LBRACE {
			p.err("syntax error: unexpected { after top level declaration")
			p.skip(gotoken.RBRACE)
			p.n()
		}
		p.must(gotoken.SEMICOLON)
	}
}

// file:
// 	"package" IDENT ';' imports topLevelDeclList
func (p *parser) file() {
	if p.syntaxError == nil {
		p.syntaxError = func(*parser) {}
		p.noSyntaxErrorFunc = true
	}
	if p.n() != gotoken.PACKAGE {
		p.syntaxError(p)
		switch {
		case p.sourceFile.Package.ctx.tweaks.errchk:
			p.err0(token.Position{}, "syntax error: package statement must be first")
		default:
			p.err("syntax error: package statement must be first")
		}
		return
	}

	p.n()
	tok, ok := p.mustTok(gotoken.IDENT)
	if !ok {
		return
	}

	if tok.Val == "_" {
		p.err0(tok.Position(), "invalid package name _")
	}
	if p.must(gotoken.SEMICOLON) && p.sourceFile.build {
		pkg := p.sourceFile.Package
		switch nm := pkg.Name; {
		case nm == "":
			pkg.Name = tok.Val
			pkg.nameOrigin = newPos(tok)
		case nm != tok.Val:
			panic(fmt.Errorf("TODO4240 %v %v %v", tok.Position(), pkg.Name, pkg.nameOrigin.Position()))
		}
		p.imports()
		p.topLevelDeclList()
		if p.scope != nil && p.scope.Kind != PackageScope {
			panic("internal error 026")
		}
	}
}
