// Copyright 2016 The GC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gc

import (
	"fmt"
	"go/token"
	"path/filepath"
	"sort"
	"strings"
	"unicode"

	"github.com/cznic/xc"
)

// Node represents the interface all AST nodes, xc.Token and lex.Char implement.
type Node interface {
	Pos() token.Pos
}

// ------------------------------------------------------------------- Argument

func (n *Argument) check(ctx *Context, stack []Declaration, node Node) (stop bool) {
	if n == nil {
		return
	}

	return n.Expression.check(ctx, stack, node) || n.TypeLiteral.check(ctx, stack, node)
}

func (n *Argument) ident() (t xc.Token) {
	if n.Case != 0 { // Expression
		return t
	}

	return n.Expression.ident()
}

// ------------------------------------------------------------------ ArrayType

func (n *ArrayType) check(ctx *Context, stack []Declaration, node Node) (stop bool) {
	if n == nil {
		return false
	}

	done, stop := n.guard.check(ctx, stack, node)
	if done || stop {
		return stop
	}

	defer n.guard.done()

	switch n.Case {
	case 0: // '[' "..." ']' Typ
		stop = n.Typ.check(ctx, stack, node)
		todo(n) //TODO
	case 1: // '[' Expression ']' Typ
		stop = n.Expression.check(ctx, stack, node) || n.Typ.check(ctx, stack, node)
		todo(n) //TODO
	default:
		panic("internal error")
	}
	return stop
}

// ----------------------------------------------------------------------- Call

func (n *Call) check(ctx *Context, stack []Declaration, node Node) (stop bool) {
	if n == nil {
		return false
	}

	for l := n.ArgumentList; l != nil; l = l.ArgumentList {
		if l.Argument.check(ctx, stack, node) {
			return true
		}
	}
	return false
}

func (n *Call) args() (_ []Value, ddd bool) {
	var a []Value
	for l := n.ArgumentList; l != nil; l = l.ArgumentList {
		a = append(a, l.Argument.Value)
	}
	return a, n.Case == 2 // '(' ArgumentList "..." CommaOpt ')'
}

// ------------------------------------------------------------------- ChanType

func (n *ChanType) check(ctx *Context, stack []Declaration, node Node) (stop bool) {
	if n == nil {
		return false
	}

	done, stop := n.guard.check(ctx, stack, node)
	if done || stop {
		return stop
	}

	defer n.guard.done()

	if n.Typ.check(ctx, stack, node) {
		return true
	}

	switch t := n.Typ.Type; n.Case {
	case 0: // "chan" Typ
		n.Type = newChanType(ctx, BothDir, t)
	case 1: // "chan" TXCHAN Typ
		n.Type = newChanType(ctx, SendDir, t)
	case 2: // RXCHAN "chan" Typ
		n.Type = newChanType(ctx, RecvDir, t)
	default:
		panic("internal error")
	}

	return stop
}

// ------------------------------------------------------------ CompLitItemList

func (n *CompLitItemList) check(ctx *Context, stack []Declaration, node Node) (stop bool) {
	for ; n != nil; n = n.CompLitItemList {
		i := n.CompLitItem
		if i.CompLitValue.check(ctx, stack, node) ||
			i.CompLitValue2.check(ctx, stack, node) ||
			i.Expression.check(ctx, stack, node) ||
			i.Expression2.check(ctx, stack, node) {
			return true
		}

		switch i.Case {
		case 0: // CompLitValue
			todo(n) //TODO
		case 1: // CompLitValue ':' CompLitValue
			todo(n) //TODO
		case 2: // CompLitValue ':' Expression
			todo(n) //TODO
		case 3: // Expression
			todo(n) //TODO
		case 4: // Expression ':' CompLitValue
			todo(n) //TODO
		case 5: // Expression ':' Expression
			todo(n) //TODO
		default:
			panic("internal error")
		}
	}
	return false
}

// ---------------------------------------------------------------- CompLitType

func (n *CompLitType) check(ctx *Context, stack []Declaration, node Node) (stop bool) {
	if n == nil {
		return false
	}

	done, stop := n.guard.check(ctx, stack, node)
	if done || stop {
		return stop
	}

	defer n.guard.done()

	switch n.Case {
	case 0: // ArrayType
		stop = n.ArrayType.check(ctx, stack, node)
		todo(n) //TODO
	case 1: // MapType
		stop = n.MapType.check(ctx, stack, node)
		todo(n) //TODO
	case 2: // SliceType
		stop = n.SliceType.check(ctx, stack, node)
		todo(n) //TODO
	case 3: // StructType
		stop = n.StructType.check(ctx, stack, node)
		todo(n) //TODO
	default:
		panic("internal error")
	}
	return stop
}

// --------------------------------------------------------------- CompLitValue

func (n *CompLitValue) check(ctx *Context, stack []Declaration, node Node) (stop bool) {
	if n == nil {
		return false
	}

	switch n.Case {
	case 0: // '{' '}'
		todo(n) //TODO
	case 1: // '{' CompLitItemList CommaOpt '}'
		stop = n.CompLitItemList.check(ctx, stack, node)
		todo(n) //TODO
	default:
		panic("internal error")
	}
	return stop
}

// ------------------------------------------------------------------ ConstSpec

func (n *ConstSpec) decl(lx *lexer, t *Typ, el *ExpressionList) {
	defer func() {
		lx.firstConstSpec = false
	}()

	el0 := el
	if el0 == nil {
		el0 = lx.constExpr
	}
	if el0 == nil && lx.firstConstSpec {
		lx.err(n, "constant declaration must have expression")
		return
	}
	scopeStart := lx.lookahead.Pos()
loop:
	for l := n.IdentifierList; l != nil; l = l.IdentifierList {
		var e *Expression
		switch {
		case el != nil:
			e = el.Expression
			el = el.ExpressionList
		case el0 != nil:
			lx.err(el0, "not enough expression(s) in list")
			break loop
		}
		d := newConstDeclaration(l.ident(), t, e, scopeStart)
		lx.scope.declare(lx, d)
	}
	if el != nil {
		lx.err(el, "extra expression(s) in list")
	}
}

// ----------------------------------------------------------------- Expression

func (n *Expression) check(ctx *Context, stack []Declaration, node Node) (stop bool) {
	if n == nil {
		return false
	}

	if n.Expression.check(ctx, stack, node) ||
		n.Expression2.check(ctx, stack, node) ||
		n.UnaryExpression.check(ctx, stack, node) {
		return true
	}

	if n.Case != 0 && (n.Expression.Value == nil || n.Expression2.Value == nil) {
		return false
	}

	switch n.Case {
	case 0: // UnaryExpression
		n.Value = n.UnaryExpression.Value
	case 1: // Expression '%' Expression
		todo(n) //TODO
	case 2: // Expression '&' Expression
		todo(n) //TODO
	case 3: // Expression '*' Expression
		n.Value = n.Expression.Value.BinaryMultiply(ctx, n.Token, n.Expression2.Value)
	case 4: // Expression '+' Expression
		todo(n) //TODO
	case 5: // Expression '-' Expression
		todo(n) //TODO
	case 6: // Expression '/' Expression
		todo(n) //TODO
	case 7: // Expression '<' Expression
		todo(n) //TODO
	case 8: // Expression '>' Expression
		todo(n) //TODO
	case 9: // Expression '^' Expression
		todo(n) //TODO
	case 10: // Expression '|' Expression
		todo(n) //TODO
	case 11: // Expression "&&" Expression
		todo(n) //TODO
	case 12: // Expression "&^" Expression
		todo(n) //TODO
	case 13: // Expression "==" Expression
		todo(n) //TODO
	case 14: // Expression ">=" Expression
		todo(n) //TODO
	case 15: // Expression "<=" Expression
		todo(n) //TODO
	case 16: // Expression "<<" Expression
		todo(n) //TODO
	case 17: // Expression "!=" Expression
		todo(n) //TODO
	case 18: // Expression "||" Expression
		todo(n) //TODO
	case 19: // Expression ">>" Expression
		todo(n) //TODO
	case 20: // Expression "<-" Expression
		todo(n) //TODO
	default:
		panic("internal error")
	}
	return false
}

func (n *Expression) ident() (t xc.Token) {
	if n.Case != 0 { // UnaryExpression
		return t
	}

	u := n.UnaryExpression
	if u.Case != 7 { // PrimaryExpression
		return t
	}

	p := u.PrimaryExpression
	if p.Case != 0 { // Operand
		return t
	}

	o := p.Operand
	if o.Case != 4 { // IDENTIFIER GenericArgumentsOpt
		return t
	}

	if o.GenericArgumentsOpt != nil {
		todo(n) //TODO
	}
	return o.Token
}

//--------------------------------------------------------------- ExpressionOpt

func (n *ExpressionOpt) check(ctx *Context, stack []Declaration, node Node) (stop bool) {
	if n == nil {
		return false
	}

	n.Expression.check(ctx, stack, node)
	todo(n) //TODO
	return false
}

// ------------------------------------------------------------------- FuncType

func (n *FuncType) check(ctx *Context, stack []Declaration, node Node) (stop bool) {
	if n == nil {
		return false
	}

	done, stop := n.guard.check(ctx, stack, node)
	if done || stop {
		return stop
	}

	defer n.guard.done()

	if n.Signature.check(ctx, stack, node) {
		return true
	}

	var in, out []Type
	isVariadic := false
	for l := n.Signature.Parameters.ParameterDeclList; l != nil; l = l.ParameterDeclList {
		switch i := l.ParameterDecl; i.Case {
		case 0: // "..." Typ
			in = append(in, newSliceType(ctx, i.Typ.Type))
			isVariadic = true
		case 1: // IDENTIFIER "..." Typ
			in = append(in, newSliceType(ctx, i.Typ.Type))
			isVariadic = true
		case 2: // IDENTIFIER Typ
			in = append(in, i.Typ.Type)
		case 3: // Typ
			switch {
			case i.isParamName:
				in = append(in, i.typ.Type)
			default:
				in = append(in, i.Typ.Type)
			}
		default:
			panic("internal error")
		}
	}
	switch t := n.Signature.Type; {
	case t != nil && t.Kind() == Tuple:
		out = t.Elements()
	case t == nil:
		// nop
	default:
		out = []Type{t}
	}
	n.Type = newFuncType(ctx, 0, in, out, false, isVariadic)
	return false
}

// ------------------------------------------------------------- IdentifierList

func (n *IdentifierList) ident() xc.Token {
	if n.Case == 0 { // IDENTIFIER
		return n.Token
	}

	// case 1: // IdentifierList ',' IDENTIFIER
	return n.Token2

}

// ----------------------------------------------------------------- ImportSpec

func (n *ImportSpec) post(lx *lexer) {
	var nm int
	var pos token.Pos
	switch n.Case {
	case
		0, // '.' BasicLiteral
		2: // '.' BasicLiteral error
		nm = idDot
		pos = n.Token.Pos()
	case
		1, // IdentifierOpt BasicLiteral
		3: // IdentifierOpt BasicLiteral error
		if o := n.IdentifierOpt; o != nil {
			nm = o.Token.Val
			pos = o.Token.Pos()
		}
	default:
		panic("internal error")
	}
	var ip string
	var bl *BasicLiteral
	switch bl = n.BasicLiteral; bl.Case {
	case
		0, //CHAR_LIT
		1, //FLOAT_LIT
		2, //IMAG_LIT
		3: //INT_LIT
		lx.err(bl, "import statement not a string")
		return
	case 4: //STRING_LIT
		if !pos.IsValid() {
			pos = bl.Pos()
		}
		val := bl.stringValue
		if val == nil {
			return
		}

		ip = string(val.s())
		if ip == "" {
			lx.err(bl, "import path is empty")
			return
		}

		if ip[0] == '/' {
			lx.err(bl, "import path cannot be absolute path")
			return
		}

		for _, r := range ip {
			if strings.ContainsAny(ip, "!\"#$%&'()*,:;<=>?[\\]^`{|}") || r == '\ufffd' {
				lx.err(bl, "import path contains invalid character")
				return
			}

			if !unicode.In(r, unicode.L, unicode.M, unicode.N, unicode.P, unicode.S) {
				lx.err(bl, "invalid import path")
				return
			}
		}

		if isRelativeImportPath(ip) {
			pi := lx.pkg.ImportPath
			if t := lx.test; t != nil {
				if _, ok := t.pkgMap[pi]; ok {
					pi = filepath.Dir(pi)
				}
			}
			ip = filepath.Join(pi, ip)
		}
	default:
		panic("internal error")
	}
	lx.imports = append(lx.imports, newImportDeclaration(lx.pkg, nm, pos, lx.oncePackage(bl, ip))) // Async load.
}

// -------------------------------------------------------------- InterfaceType

func (n *InterfaceType) check(ctx *Context, stack []Declaration, node Node) (stop bool) {
	if n == nil {
		return false
	}

	done, stop := n.guard.check(ctx, stack, node)
	if done || stop {
		return stop
	}

	defer n.guard.done()

	switch n.Case {
	case 0: // "interface" LBrace '}'
		n.Type = newInterfaceType(ctx, nil)
	case 1: // "interface" LBrace InterfaceMethodDeclList SemicolonOpt '}'
		//TODO process embedded interfaces
		s := n.methods
		a := make(declarations, 0, len(s.Bindings)+len(s.Unbound))
		for _, d := range s.Bindings {
			a = append(a, d)
		}
		for _, d := range s.Unbound {
			a = append(a, d)
		}
		sort.Sort(a)
		for _, m := range a {
			if m.check(ctx, stack, node) {
				return true
			}
			todo(n) //TODO
		}
	default:
		panic("internal error")
	}

	return stop
}

// ------------------------------------------------------ LBraceCompLitItemList

func (n *LBraceCompLitItemList) check(ctx *Context, stack []Declaration, node Node) (stop bool) {
	for ; n != nil; n = n.LBraceCompLitItemList {
		i := n.LBraceCompLitItem
		if i.Expression.check(ctx, stack, node) ||
			i.Expression2.check(ctx, stack, node) ||
			i.LBraceCompLitValue.check(ctx, stack, node) {
			return true
		}

		switch i.Case {
		case 0: // Expression
			todo(n) //TODO
		case 1: // Expression ':' Expression
			todo(n) //TODO
		case 2: // Expression ':' LBraceCompLitValue
			todo(n) //TODO
		case 3: // LBraceCompLitValue
			todo(n) //TODO
		default:
			panic("internal error")
		}

	}
	return false
}

// --------------------------------------------------------- LBraceCompLitValue

func (n *LBraceCompLitValue) check(ctx *Context, stack []Declaration, node Node) (stop bool) {
	if n == nil {
		return false
	}

	switch n.Case {
	case 0: // LBrace '}'
		todo(n) //TODO
	case 1: // LBrace LBraceCompLitItemList CommaOpt '}'
		stop = n.LBraceCompLitItemList.check(ctx, stack, node)
		todo(n) //TODO
	default:
		panic("internal error")
	}
	return stop
}

// -------------------------------------------------------------------- MapType

func (n *MapType) check(ctx *Context, stack []Declaration, node Node) (stop bool) {
	if n == nil {
		return false
	}

	stack = nil
	done, stop := n.guard.check(ctx, stack, node)
	if done || stop {
		return stop
	}

	defer n.guard.done()

	// "map" '[' Typ ']' Typ
	stop = n.Typ.check(ctx, stack, node) || n.Typ2.check(ctx, stack, node)
	n.Type = newMapType(ctx, n.Typ.Type, n.Typ2.Type)
	return stop
}

// -------------------------------------------------------------------- Operand

func (n *Operand) check(ctx *Context, stack []Declaration, node Node) (stop bool) {
	if n == nil {
		return false
	}

	if n.Expression.check(ctx, stack, node) ||
		n.FuncType.check(ctx, stack, node) ||
		n.StatementList.check(ctx, stack, node) {
		return true
	}

	switch n.Case {
	case 0: // '(' Expression ')'
		todo(n) //TODO
	case 1: // '(' TypeLiteral ')'
		todo(n) //TODO
	case 2: // BasicLiteral
		n.Value = n.BasicLiteral.Value
	case 3: // FuncType LBrace StatementList '}'
		todo(n) //TODO
	case 4: // IDENTIFIER GenericArgumentsOpt
		if n.GenericArgumentsOpt != nil {
			todo(n) //TODO
		}
		if n.Token.Val == idUnderscore {
			if ctx.err(n, "cannot use _ as value") {
				return true
			}

			break
		}

		d := n.resolutionScope.mustLookup(ctx, n.Token, n.fileScope)
		switch isBuiltin := ctx.isBuiltin(d); x := d.(type) {
		case nil:
			// nop
		case *ConstDeclaration:
			if isBuiltin {
				switch d.Name() {
				case idIota:
					todo(n) //TODO
				}
				break
			}

			n.Value = x.Value
		case *TypeDeclaration:
			n.Value = newTypeValue(x)
		default:
			//dbg("%T", d)
			todo(n) //TODO
		}
	default:
		panic("TODO")
	}
	return false
}

// -------------------------------------------------------------- PackageClause

func (n *PackageClause) post(lx *lexer) {
	p := lx.pkg
	t := n.Token2
	nm := t.Val
	switch {
	case p.Name != "":
		if p.name != nm {
			lx.close(t,
				"package %s: found packages %s (%s) and %s (%s) in %s",
				p.ImportPath,
				p.Name, filepath.Base(p.namedBy),
				t.S(), filepath.Base(lx.name),
				p.Directory,
			)
			return
		}
	default:
		p.Name = string(t.S())
		p.name = nm
		p.namedBy = lx.name
	}
}

// ----------------------------------------------------------------- Parameters

func (n *Parameters) post(lx *lexer) {
	if n == nil {
		return
	}

	hasNamedParams := false
	i := 0
	for l := n.ParameterDeclList; l != nil; l = l.ParameterDeclList {
		pd := l.ParameterDecl
		n.list = append(n.list, pd)
		switch pd.Case {
		case 0: // "..." Typ
			if l.ParameterDeclList != nil {
				lx.err(pd, "can only use ... as final argument in list")
			}

			if hasNamedParams {
				lx.err(pd, "mixed named and unnamed parameters")
			}
		case 1: // IDENTIFIER "..." Typ
			hasNamedParams = true
			pd.nm = pd.Token
			pd.isVariadic = true
			typ := pd.Typ
			if l.ParameterDeclList != nil {
				lx.err(pd, "can only use ... as final argument in list")
			}

			for j := i - 1; j >= 0 && n.list[j].Case == 3; j-- {
				pd := n.list[j]
				pd.isParamName = true
				pd.typ = typ
				lx.err(pd, "can only use ... as final argument in list")
			}
		case 2: // IDENTIFIER Typ
			hasNamedParams = true
			pd.nm = pd.Token
			typ := pd.Typ
			for j := i - 1; j >= 0 && n.list[j].Case == 3; j-- {
				pd := n.list[j]
				pd.isParamName = true
				pd.typ = typ
			}
		case 3: // Typ
			switch t := pd.Typ; t.Case {
			case 7: // QualifiedIdent GenericArgumentsOpt
				if t.GenericArgumentsOpt != nil {
					todo(n) //TODO
				}
				switch qi := t.QualifiedIdent; qi.Case {
				case 0: // IDENTIFIER
					pd.nm = qi.Token
				}
			}
		}
		i++
	}
	if !hasNamedParams {
		return
	}

	for _, v := range n.list {
		nm := v.nm
		if nm.IsValid() {
			if lx.scope.Bindings[nm.Val] != nil {
				lx.err(v, "duplicate argument %s", nm.S())
				continue
			}

			lx.scope.declare(lx, newParamaterDeclaration(nm, v.typ, v.isVariadic, lx.lookahead.Pos()))
			continue
		}

		if v.isParamName && !nm.IsValid() {
			lx.err(v, "mixed named and unnamed parameters")
		}
	}
}

func (n *Parameters) check(ctx *Context, stack []Declaration, node Node) (stop bool) {
	if n == nil {
		return false
	}

	for l := n.ParameterDeclList; l != nil; l = l.ParameterDeclList {
		switch i := l.ParameterDecl; i.Case {
		case
			0, // "..." Typ
			1, // IDENTIFIER "..." Typ
			2: // IDENTIFIER Typ
			if i.Typ.check(ctx, stack, node) {
				return true
			}
		case 3: // Typ
			if i.isParamName {
				if i.typ.check(ctx, stack, node) {
					return true
				}

				todo(n) //TODO
				break
			}

			if i.Typ.check(ctx, stack, node) {
				return true
			}

			todo(n) //TODO
		default:
			panic("internal error")
		}
	}
	return false
}

// ---------------------------------------------------------- PrimaryExpression

func (n *PrimaryExpression) check(ctx *Context, stack []Declaration, node Node) (stop bool) {
	if n == nil {
		return false
	}

	if n.Call.check(ctx, stack, node) ||
		n.CompLitType.check(ctx, stack, node) ||
		//TODO- n.CompLitValue.check(ctx, stack, node) ||
		n.Expression.check(ctx, stack, node) ||
		n.ExpressionOpt.check(ctx, stack, node) ||
		n.ExpressionOpt2.check(ctx, stack, node) ||
		n.ExpressionOpt3.check(ctx, stack, node) ||
		//TODO- n.LBraceCompLitValue.check(ctx, stack, node) ||
		n.Operand.check(ctx, stack, node) ||
		n.PrimaryExpression.check(ctx, stack, node) ||
		n.TypeLiteral.check(ctx, stack, node) {
		return true
	}

	switch n.Case {
	case 0: // Operand
		n.Value = n.Operand.Value
	case 1: // CompLitType LBraceCompLitValue
		//TODO if n.LBraceCompLitValue.check(ctx, stack, node) {
		//TODO 	return true
		//TODO }
		todo(n) //TODO
	case 2: // PrimaryExpression '.' '(' "type" ')'
		todo(n) //TODO
	case 3: // PrimaryExpression '.' '(' Typ ')'
		todo(n) //TODO
	case 4: // PrimaryExpression '.' IDENTIFIER
		todo(n) //TODO
	case 5: // PrimaryExpression '[' Expression ']'
		todo(n) //TODO
	case 6: // PrimaryExpression '[' ExpressionOpt ':' ExpressionOpt ':' ExpressionOpt ']'
		todo(n) //TODO
	case 7: // PrimaryExpression '[' ExpressionOpt ':' ExpressionOpt ']'
		todo(n) //TODO
	case 8: // PrimaryExpression Call
		v := n.PrimaryExpression.Value
		if v == nil {
			break
		}

		switch v.Kind() {
		case TypeValue: // Conversion.
			args, ddd := n.Call.args()
			if ddd {
				todo(n) //TODO
			}
			switch len(args) {
			case 0:
				todo(n) //TODO
			case 1:
				todo(n) //TODO
			default:
				todo(n) //TODO
			}
		default:
			//dbg("", v.Kind())
			todo(n) //TODO
		}
	case 9: // PrimaryExpression CompLitValue
		//TODO if n.CompLitValue.check(ctx, stack, node) {
		//TODO 	return true
		//TODO }
		todo(n) //TODO
	case 10: // TypeLiteral '(' Expression CommaOpt ')'
		todo(n) //TODO
	default:
		panic("internal error")
	}
	return false
}

// ------------------------------------------------------------------- Prologue

func (n *Prologue) post(lx *lexer) {
	for _, v := range lx.imports {
		p := v.once.Value().(*Package)
		v.Package = p
		switch v.Name() {
		case idDot:
			lx.dotImports = append(lx.dotImports, v)
			for _, v := range p.Scope.Bindings {
				if isExported(v.Name()) {
					lx.fileScope.declare(lx, v)
				}
			}
			continue
		case idUnderscore:
			lx.unboundImports = append(lx.unboundImports, v)
			continue
		case 0:
			switch {
			case p.ImportPath == "C":
				v.name = idC
			default:
				v.name = p.name
			}
		}

		lx.fileScope.declare(lx, v)
	}
}

// ------------------------------------------------------------- QualifiedIdent

func (n *QualifiedIdent) q() xc.Token {
	switch n.Case {
	case 0: // IDENTIFIER
		return xc.Token{}
	case 1: // IDENTIFIER '.' IDENTIFIER
		return n.Token
	default:
		panic("internal error")
	}
}

func (n *QualifiedIdent) i() xc.Token {
	switch n.Case {
	case 0: // IDENTIFIER
		return n.Token
	case 1: // IDENTIFIER '.' IDENTIFIER
		return n.Token3
	default:
		panic("internal error")
	}
}

func (n *QualifiedIdent) str() string {
	switch q, i := n.q(), n.i(); {
	case q.IsValid():
		return fmt.Sprintf("%s.%s", q.S(), i.S())
	default:
		return string(i.S())
	}
}

// ---------------------------------------------------------------- ReceiverOpt

func (n *ReceiverOpt) post(lx *lexer) {
	cnt := 0
	l := n.Parameters.ParameterDeclList
	for ; l != nil && cnt < 1; l = l.ParameterDeclList {
		cnt++
	}
	if cnt == 0 {
		lx.err(n.Parameters, "method has no receiver")
		return
	}

	if l != nil {
		lx.err(l.ParameterDecl, "method has multiple receivers")
		return
	}

	switch pd := n.Parameters.ParameterDeclList.ParameterDecl; pd.Case {
	case
		0, // "..." Typ
		1: // IDENTIFIER "..." Typ
		lx.err(pd, "invalid receiver type")
	case
		2, // IDENTIFIER Typ
		3: // Typ
		t := pd.Typ
		cnt := 0
	again:
		switch t.Case {
		case 0: // '(' Typ ')'
			t = t.Typ
			goto again
		case 1: // '*' Typ
			n.isPtr = true
			t = t.Typ
			cnt++
			if cnt < 2 {
				goto again
			}
		case 7: // QualifiedIdent GenericArgumentsOpt
			if t.GenericArgumentsOpt != nil {
				todo(n) //TODO
			}
			switch qi := t.QualifiedIdent; qi.Case {
			case 0: // IDENTIFIER
				//ok
				n.nm = qi.Token
				return
			case 1: // IDENTIFIER '.' IDENTIFIER
				lx.err(qi, "cannot define new methods on non-local type %s.%s", qi.Token.S(), qi.Token3.S())
			default:
				panic("internal error")
			}
		}
		lx.err(t, "invalid receiver type")
	default:
		panic("internal error")
	}
}

func (n *ReceiverOpt) check(ctx *Context, stack []Declaration, node Node) (stop bool) {
	if n == nil {
		return false
	}

	if nm := n.nm; nm.IsValid() {
		n.Type = n.resolutionScope.mustLookupLocalTLDType(ctx, nm)
	}
	return false
}

// ------------------------------------------------------------------ ResultOpt

func (n *ResultOpt) check(ctx *Context, stack []Declaration, node Node) (stop bool) {
	if n == nil {
		return false
	}

	stack = nil
	if n.Parameters.check(ctx, stack, node) || n.Typ.check(ctx, stack, node) {
		return true
	}

	switch n.Case {
	case 0: // /* empty */
		// nop
	case 1: // Parameters
		var a []Type
		for l := n.Parameters.ParameterDeclList; l != nil; l = l.ParameterDeclList {
			switch i := l.ParameterDecl; i.Case {
			case
				0, // "..." Typ
				1: // IDENTIFIER "..." Typ
				// Invalid
			case 2: // IDENTIFIER Typ
				a = append(a, i.Typ.Type)
			case 3: // Typ
				switch {
				case i.isParamName:
					a = append(a, i.typ.Type)
				default:
					a = append(a, i.Typ.Type)
				}
			default:
				panic("internal error")
			}
		}
		switch len(a) {
		case 0:
			// nop
		case 1:
			n.Type = a[0]
		default:
			n.Type = newTupleType(ctx, a)
		}
	case 2: // Typ
		n.Type = n.Typ.Type
	default:
		panic("internal error")
	}
	return false
}

// ------------------------------------------------------------------ Signature

func (n *Signature) post(lx *lexer) {
	n.Parameters.post(lx)
	if o := n.ResultOpt; o != nil {
		o.Parameters.post(lx)
	}
}

func (n *Signature) check(ctx *Context, stack []Declaration, node Node) (stop bool) {
	if n == nil {
		return false
	}

	if n.Parameters.check(ctx, stack, node) ||
		n.ResultOpt.check(ctx, stack, node) {
		return true
	}

	if o := n.ResultOpt; o != nil {
		n.Type = o.Type
	}
	return false
}

// ------------------------------------------------------------------ SliceType

func (n *SliceType) check(ctx *Context, stack []Declaration, node Node) (stop bool) {
	if n == nil {
		return false
	}

	stack = nil
	done, stop := n.guard.check(ctx, stack, node)
	if done || stop {
		return stop
	}

	defer n.guard.done()

	// '[' ']' Typ
	stop = n.Typ.check(ctx, stack, node)
	todo(n) //TODO

	return stop
}

// ------------------------------------------------------------------ Statement

func (n *Statement) check(ctx *Context, stack []Declaration, node Node) (stop bool) {
	todo(n) //TODO
	return false
}

// -------------------------------------------------------------- StatementList

func (n *StatementList) check(ctx *Context, stack []Declaration, node Node) (stop bool) {
	for ; n != nil; n = n.StatementList {
		i := n.Statement
		if i.check(ctx, stack, node) {
			return true
		}
	}
	return false
}

// ------------------------------------------------------------ StructFieldDecl

func (n *StructFieldDecl) decl(lx *lexer) {
	var tag stringValue
	if o := n.TagOpt; o != nil {
		tag = o.stringValue
	}
	switch n.Case {
	case 0: // '*' QualifiedIdent TagOpt
		lx.scope.declare(lx, newFieldDeclaration(
			n.QualifiedIdent.i(),
			nil,
			true,
			n.QualifiedIdent,
			tag,
		))
	case 1: // IdentifierList Typ TagOpt
		for l := n.IdentifierList; l != nil; l = l.IdentifierList {
			lx.scope.declare(lx, newFieldDeclaration(
				l.ident(),
				nil,
				false,
				nil,
				tag,
			))
		}
	case 2: // QualifiedIdent TagOpt
		lx.scope.declare(lx, newFieldDeclaration(
			n.QualifiedIdent.i(),
			nil,
			false,
			n.QualifiedIdent,
			tag,
		))
	case 3: // '(' QualifiedIdent ')' TagOpt
		lx.err(n.Token, "cannot parenthesize embedded type")
	case 4: // '(' '*' QualifiedIdent ')' TagOpt
		lx.err(n.Token, "cannot parenthesize embedded type")
	case 5: // '*' '(' QualifiedIdent ')' TagOpt
		lx.err(n.Token2, "cannot parenthesize embedded type")
	default:
		panic("internal error")
	}
}

// ----------------------------------------------------------------- StructType

func (n *StructType) check(ctx *Context, stack []Declaration, node Node) (stop bool) {
	if n == nil {
		return false
	}

	done, stop := n.guard.check(ctx, stack, node)
	if done || stop {
		return stop
	}

	defer n.guard.done()

	switch n.Case {
	case 0: // "struct" LBrace '}'
		n.Type = newStructType(ctx, nil)
	case 1: // "struct" LBrace StructFieldDeclList SemicolonOpt '}'
		//TODO process embedded fields
		s := n.fields
		a := make(fields, 0, len(s.Bindings)+len(s.Unbound))
		for _, d := range s.Bindings {
			a = append(a, d.(*FieldDeclaration))
		}
		for _, d := range s.Unbound {
			a = append(a, d.(*FieldDeclaration))
		}
		sort.Sort(a)
		for _, f := range a {
			if f.check(ctx, stack, node) {
				return true
			}
			todo(n) //TODO
		}
	default:
		panic("internal error")
	}

	return stop
}

// ------------------------------------------------------------------------ Typ

func (n *Typ) normalize() *Typ {
	for {
		switch n.Case {
		case 0: // '(' Typ ')'
			n = n.Typ
		default:
			return n
		}
	}
}

func (n *Typ) check(ctx *Context, stack []Declaration, node Node) (stop bool) {
	if n == nil {
		return false
	}

	done, stop := n.guard.check(ctx, stack, node)
	if done || stop {
		return stop
	}

	defer n.guard.done()

	t := n.normalize()
	switch t.Case {
	case 0: // '(' Typ ')'
		panic("internal error")
	case 1: // '*' Typ
		t2 := t.Typ.normalize()
		if t2.check(ctx, stack, node) {
			return true
		}
		n.Type = newPtrType(ctx, t2.Type)
	case 2: // ArrayType
		if t.ArrayType.check(ctx, stack, node) {
			return true
		}

		todo(n) //TODO
	case 3: // ChanType
		if t.ChanType.check(ctx, stack, node) {
			return true
		}

		n.Type = t.ChanType.Type
	case 4: // FuncType
		if t.FuncType.check(ctx, stack, node) {
			return true
		}

		n.Type = t.FuncType.Type
	case 5: // InterfaceType
		if t.InterfaceType.check(ctx, stack, node) {
			return true
		}

		n.Type = t.InterfaceType.Type
	case 6: // MapType
		if t.MapType.check(ctx, stack, node) {
			return true
		}

		n.Type = t.MapType.Type
	case 7: // QualifiedIdent GenericArgumentsOpt
		if t.GenericArgumentsOpt != nil {
			todo(n) //TODO
		}
		if t := n.resolutionScope.mustLookupType(ctx, n.QualifiedIdent, n.fileScope); t != nil {
			n.Type = t
			t.check(ctx, stack, node)
		}
	case 8: // SliceType
		if t.SliceType.check(ctx, stack, node) {
			return true
		}

		n.Type = t.SliceType.Type
	case 9: // StructType
		if t.StructType.check(ctx, stack, node) {
			return true
		}

		n.Type = t.StructType.Type
	default:
		panic("internal error")
	}
	return false
}

// ---------------------------------------------------------------- TypeLiteral

func (n *TypeLiteral) check(ctx *Context, stack []Declaration, node Node) (stop bool) {
	if n == nil {
		return false
	}

	done, stop := n.guard.check(ctx, stack, node)
	if done || stop {
		return stop
	}

	defer n.guard.done()

	todo(n) //TODO
	return false
}

// ------------------------------------------------------------ UnaryExpression

func (n *UnaryExpression) check(ctx *Context, stack []Declaration, node Node) (stop bool) {
	if n == nil {
		return false
	}

	if n.UnaryExpression.check(ctx, stack, node) ||
		n.PrimaryExpression.check(ctx, stack, node) {
		return true
	}

	switch n.Case {
	case 0: // '!' UnaryExpression
		todo(n) //TODO
	case 1: // '&' UnaryExpression
		todo(n) //TODO
	case 2: // '*' UnaryExpression
		todo(n) //TODO
	case 3: // '+' UnaryExpression
		todo(n) //TODO
	case 4: // '-' UnaryExpression
		if v := n.UnaryExpression.Value; v != nil {
			n.Value = v.UnaryMinus(ctx, n.UnaryExpression)
		}
	case 5: // '^' UnaryExpression
		todo(n) //TODO
	case 6: // "<-" UnaryExpression
		todo(n) //TODO
	case 7: // PrimaryExpression
		n.Value = n.PrimaryExpression.Value
	default:
		panic("internal error")
	}
	return false
}
