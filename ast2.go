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

	"github.com/cznic/mathutil"
	"github.com/cznic/xc"
)

// Node represents the interface all AST nodes, xc.Token and lex.Char implement.
type Node interface {
	Pos() token.Pos
}

// ------------------------------------------------------------------- Argument

func (n *Argument) check(ctx *context) (stop bool) {
	if n == nil {
		return
	}

	switch n.Case {
	case 0: // Expression
		if n.Expression.check(ctx) {
			return true
		}

		n.Value = n.Expression.Value
	case 1: // TypeLiteral
		if n.TypeLiteral.check(ctx) {
			return true
		}

		if t := n.TypeLiteral.Type; t != nil {
			n.Value = newTypeValue(t)
		}
	default:
		panic("internal error")
	}
	return false
}

func (n *Argument) ident() (t xc.Token) {
	if n.Case != 0 { // Expression
		return t
	}

	return n.Expression.ident()
}

// ------------------------------------------------------------------ ArrayType

func (n *ArrayType) check(ctx *context) (stop bool) {
	if n == nil {
		return false
	}

	done, stop := n.guard.check(ctx, nil)
	if done || stop {
		return stop
	}

	defer n.guard.done(nil)

	switch n.Case {
	case 0: // '[' "..." ']' Typ
		stop = n.Typ.check(ctx)
		if t := n.Typ.Type; t != nil {
			n.Type = newArrayType(ctx, t, n.items, !n.itemsSet)
		}
	case 1: // '[' Expression ']' Typ
		stop = n.Expression.check(ctx) || n.Typ.check(ctx)
		v := n.Expression.Value
		if v == nil {
			break
		}

		t := n.Typ.Type
		if t == nil {
			break
		}

		switch v.Kind() {
		case ConstValue:
			switch c := v.Const(); c.Kind() {
			case IntConst:
				if d := c.Convert(ctx.intType); d != nil {
					val := d.Const().(*intConst).val
					if val >= 0 {
						n.Type = newArrayType(ctx, t, val, false)
						break
					}

					todo(n, true)
					break
				}

				ctx.err(n.Expression, "array bound is too large")
			default:
				todo(n)
			}
		case NilValue:
			ctx.err(n.Expression, "invalid array bound nil")
		default:
			todo(n, true)
		}
	default:
		panic("internal error")
	}
	return stop
}

// ----------------------------------------------------------------------- Call

func (n *Call) check(ctx *context) (stop bool) {
	if n == nil {
		return false
	}

	for l := n.ArgumentList; l != nil; l = l.ArgumentList {
		if l.Argument.check(ctx) {
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

func (n *ChanType) check(ctx *context) (stop bool) {
	if n == nil {
		return false
	}

	done, stop := n.guard.check(ctx, nil)
	if done || stop {
		return stop
	}

	defer n.guard.done(nil)

	if n.Typ.check(ctx) {
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

func (n *CompLitItemList) check(ctx *context, t Type) (stop bool) {
	if t == nil || t.Kind() == Invalid {
		return false
	}

	var index, len int64
	var numField int
	var et Type
	tk := t.Kind()
	switch tk {
	case Array:
		len = t.Len()
		et = t.Elem()
	case Map:
		//kt = t.Key()
		et = t.Elem()
	case Slice:
		et = t.Elem()
	case Struct:
		numField = t.NumField()
	default:
		ctx.err(n, "invalid type for composite literal: %s", t)
		return false
	}

	m := map[int64]struct{}{}
	_ = et
	_ = len
	var keyed, unkeyed bool
	for ; n != nil; n = n.CompLitItemList {
		switch i := n.CompLitItem; i.Case {
		case 0: // CompLitValue
			switch tk {
			case Array:
				todo(i)
			case Map:
				todo(i, true)
			case Slice:
				todo(i)
			case Struct:
				unkeyed = true
				if keyed {
					todo(i, true) // mix
				}
				todo(i)
			}
		case 1: // CompLitValue ':' CompLitValue
			switch tk {
			case Array:
				todo(i, true) // invalid
			case Map:
				todo(i)
			case Slice:
				todo(i, true) // invalid
			case Struct:
				todo(i, true) // invalid
			}
		case 2: // CompLitValue ':' Expression
			switch tk {
			case Array:
				todo(i, true) // invalid
			case Map:
				todo(i)
			case Slice:
				todo(i, true) // invalid
			case Struct:
				todo(i, true) // invalid
			}
		case 3: // Expression
			switch tk {
			case Array:
				todo(i)
			case Map:
				todo(i, true)
			case Slice:
				todo(i)
			case Struct:
				unkeyed = true
				if keyed {
					todo(i, true) // mix
				}
				if int(index) >= numField {
					if ctx.err(i, "too many values in struct initializer") {
						return true
					}

					break
				}

				f := t.Field(int(index))
				if i.Expression.check(ctx) {
					return true
				}

				if v := i.Expression.Value; v != nil && f.Type != nil && !v.AssignableTo(f.Type) {
					todo(i, true)
				}
				index++
			default:
				panic("internal error")
			}
		case 4: // Expression ':' CompLitValue
			switch tk {
			case Array:
				if i.Expression.check(ctx) {
					return true
				}

				if v := i.Expression.Value; v != nil && !v.nonNegativeInteger() {
					todo(i.Expression, true) // index must be non-negative integer constant
				}

				todo(i)
			case Map:
				todo(i)
			case Slice:
				if i.Expression.check(ctx) {
					return true
				}

				if v := i.Expression.Value; v != nil && !v.nonNegativeInteger() {
					todo(i.Expression, true) // index must be non-negative integer constant
				}

				todo(i)
			case Struct:
				keyed = true
				if unkeyed {
					todo(i, true) // mix
				}
				todo(i)
			}
		case 5: // Expression ':' Expression
			switch tk {
			case Array:
				if i.Expression.check(ctx) {
					return true
				}

				if v := i.Expression.Value; v != nil && !v.nonNegativeInteger() {
					todo(i.Expression, true) // index must be non-negative integer constant
				}

				todo(i)
			case Map:
				todo(i)
			case Slice:
				if i.Expression.check(ctx) {
					return true
				}

				if v := i.Expression.Value; v != nil && !v.nonNegativeInteger() {
					todo(i.Expression, true) // index must be non-negative integer constant
				}

				todo(i)
			case Struct:
				keyed = true
				if unkeyed {
					if ctx.err(i.Expression, "mixture of field:value and value initializers") {
						return true
					}
				}
				fn, ok := i.Expression.isIdentifier()
				if !ok {
					if ctx.err(i.Expression, "invalid field name in struct initializer") {
						return true
					}

					break
				}

				if _, ok := m[int64(fn.Val)]; ok {
					if ctx.err(i.Expression, "duplicate field name in struct literal: %s", fn.S()) {
						return true
					}

					break
				}

				f := t.FieldByName(fn.Val)
				if f == nil {
					todo(i.Expression, true) // unknown field
					break
				}

				m[int64(fn.Val)] = struct{}{}
				if i.Expression2.check(ctx) {
					return true
				}

				if v := i.Expression2.Value; v != nil && f.Type != nil && !v.AssignableTo(f.Type) {
					todo(i.Expression2) // assignment fail
				}
			default:
				panic("internal error")
			}
		default:
			panic("internal error")
		}
	}
	return false
}

// ---------------------------------------------------------------- CompLitType

func (n *CompLitType) check(ctx *context) (stop bool) {
	if n == nil {
		return false
	}

	done, stop := n.guard.check(ctx, nil)
	if done || stop {
		return stop
	}

	defer n.guard.done(nil)

	switch n.Case {
	case 0: // ArrayType
		stop = n.ArrayType.check(ctx)
		n.Type = n.ArrayType.Type
	case 1: // MapType
		stop = n.MapType.check(ctx)
		n.Type = n.MapType.Type
	case 2: // SliceType
		stop = n.SliceType.check(ctx)
		n.Type = n.SliceType.Type
	case 3: // StructType
		stop = n.StructType.check(ctx)
		n.Type = n.StructType.Type
	default:
		panic("internal error")
	}
	return stop
}

// --------------------------------------------------------------- CompLitValue

func (n *CompLitValue) check(ctx *context, t Type) (stop bool) {
	if n == nil || t == nil || t.Kind() == Invalid {
		return false
	}

	n.Type = t
	switch n.Case {
	case 0: // '{' '}'
		// nop
	case 1: // '{' CompLitItemList CommaOpt '}'
		stop = n.CompLitItemList.check(ctx, t)
	default:
		panic("internal error")
	}
	return stop
}

// ------------------------------------------------------------------ ConstSpec

func (n *ConstSpec) decl(lx *lexer, t *Typ, el *ExpressionList) {
	defer func() {
		lx.firstConstSpec = false
		lx.iota++
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
		d := newConstDeclaration(l.ident(), t, e, lx.iota, scopeStart)
		lx.scope.declare(lx, d)
	}
	if el != nil {
		lx.err(el, "extra expression(s) in list")
	}
}

// ----------------------------------------------------------------- Expression

func (n *Expression) isIdentifier() (xc.Token, bool) {
	if n.Case != 0 {
		return xc.Token{}, false
	}

	return n.UnaryExpression.isIdentifier()
}

func (n *Expression) check(ctx *context) (stop bool) {
	if n == nil {
		return false
	}

	if n.Expression.check(ctx) ||
		n.Expression2.check(ctx) ||
		n.UnaryExpression.check(ctx) {
		return true
	}

	if n.Case != 0 && (n.Expression.Value == nil || n.Expression2.Value == nil) {
		return false
	}

	switch n.Case {
	case 0: // UnaryExpression
		n.Value = n.UnaryExpression.Value
	case 1: // Expression '%' Expression
		n.Value = n.Expression.Value.Mod(n.Token, n.Expression2.Value)
	case 2: // Expression '&' Expression
		n.Value = n.Expression.Value.And(n.Token, n.Expression2.Value)
	case 3: // Expression '*' Expression
		n.Value = n.Expression.Value.Mul(n.Token, n.Expression2.Value)
	case 4: // Expression '+' Expression
		n.Value = n.Expression.Value.Add(n.Token, n.Expression2.Value)
	case 5: // Expression '-' Expression
		n.Value = n.Expression.Value.Sub(n.Token, n.Expression2.Value)
	case 6: // Expression '/' Expression
		n.Value = n.Expression.Value.Div(n.Token, n.Expression2.Value)
	case 7: // Expression '<' Expression
		todo(n)
	case 8: // Expression '>' Expression
		todo(n)
	case 9: // Expression '^' Expression
		todo(n)
	case 10: // Expression '|' Expression
		n.Value = n.Expression.Value.Or(n.Token, n.Expression2.Value)
	case 11: // Expression "&&" Expression
		todo(n)
	case 12: // Expression "&^" Expression
		todo(n)
	case 13: // Expression "==" Expression
		todo(n)
	case 14: // Expression ">=" Expression
		todo(n)
	case 15: // Expression "<=" Expression
		todo(n)
	case 16: // Expression "<<" Expression
		n.Value = n.Expression.Value.Lsh(n.Token, n.Expression2.Value)
	case 17: // Expression "!=" Expression
		todo(n)
	case 18: // Expression "||" Expression
		todo(n)
	case 19: // Expression ">>" Expression
		n.Value = n.Expression.Value.Rsh(n.Token, n.Expression2.Value)
	case 20: // Expression "<-" Expression
		todo(n)
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
		todo(n)
	}
	return o.Token
}

//--------------------------------------------------------------- ExpressionOpt

func (n *ExpressionOpt) check(ctx *context) (stop bool) {
	if n == nil {
		return false
	}

	n.Expression.check(ctx)
	todo(n)
	return false
}

// ------------------------------------------------------------------- FuncType

func (n *FuncType) check(ctx *context) (stop bool) {
	if n == nil {
		return false
	}

	done, stop := n.guard.check(ctx, nil)
	if done || stop {
		return stop
	}

	defer n.guard.done(nil)

	if n.Signature.check(ctx) {
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

func (n *InterfaceType) check(ctx *context) (stop bool) {
	if n == nil {
		return false
	}

	done, stop := n.guard.check(ctx, nil)
	if done || stop {
		return stop
	}

	defer n.guard.done(nil)

	switch n.Case {
	case 0: // "interface" LBrace '}'
		n.Type = newInterfaceType(ctx, nil)
	case 1: // "interface" LBrace InterfaceMethodDeclList SemicolonOpt '}'
		s := n.methods
		a := make(declarations, 0, len(s.Bindings)+len(s.Unbound))
		for _, d := range s.Bindings {
			a = append(a, d)
		}
		for _, d := range s.Unbound {
			a = append(a, d)
		}
		var mta []Method
		var index int
		for l := n.InterfaceMethodDeclList; l != nil; l = l.InterfaceMethodDeclList {
			switch id := l.InterfaceMethodDecl; id.Case {
			case 1: //  QualifiedIdent
				if d := id.resolutionScope.mustLookupQI(ctx, id.QualifiedIdent, id.fileScope); d != nil {
					ctx.errf = func() bool {
						ctx.err(id.QualifiedIdent, "interface type loop involving %s", id.QualifiedIdent.str())
						return true
					}
					if d.check(ctx) {
						continue
					}

					switch x := d.(type) {
					case *TypeDeclaration:
						t := Type(x)
						switch t.Kind() {
						case Interface:
							for i := 0; i < t.NumMethod(); i++ {
								m := t.Method(i)
								switch {
								case s.Bindings[m.Name] != nil:
									todo(n, true)
								default:
									var pth int
									if !isExported(m.Name) {
										pth = n.pkgPath
									}
									mt := Method{m.Name, pth, m.Type, index, false}
									index++
									mta = append(mta, mt)
								}
							}
						default:
							if ctx.err(id, "interface contains embedded non-interface %s", id.QualifiedIdent.str()) {
								return true
							}
						}
					default:
						todo(n, true)
					}
				}
			}
		}
		sort.Sort(a)
		for _, m := range a {
			if m.check(ctx) {
				return true
			}

			if m.Name() != idUnderscore {
				var pth int
				fd := m.(*FuncDeclaration)
				if !fd.isExported {
					pth = n.pkgPath
				}
				mt := Method{m.Name(), pth, fd.Type, index, false}
				index++
				mta = append(mta, mt)
			}
		}
		n.Type = newInterfaceType(ctx, mta)
	default:
		panic("internal error")
	}

	return stop
}

// ------------------------------------------------------ LBraceCompLitItemList

func (n *LBraceCompLitItemList) check(ctx *context, t Type) (stop bool) {
	if n == nil || t == nil || t.Kind() == Invalid {
		return
	}

	var index, len int64
	var numField int
	var et Type
	tk := t.Kind()
	switch tk {
	case Array:
		len = t.Len()
		et = t.Elem()
	case Map:
		//kt = t.Key()
		et = t.Elem()
	case Slice:
		et = t.Elem()
	case Struct:
		numField = t.NumField()
	default:
		ctx.err(n, "invalid type for composite literal: %s", t)
		return false
	}

	m := map[int64]struct{}{}
	var keyed, unkeyed bool
	for ; n != nil; n = n.LBraceCompLitItemList {
		switch i := n.LBraceCompLitItem; i.Case {
		case 0: // Expression
			switch tk {
			case Array:
				if index >= len {
					if ctx.err(i, "array index %d out of bounds [0:%d]", index, len) {
						return true
					}

					break
				}

				if _, ok := m[index]; ok {
					todo(i, true)
				}

				m[index] = struct{}{}
				if i.Expression.check(ctx) {
					return true
				}

				if v := i.Expression.Value; v != nil && et != nil && !v.AssignableTo(et) {
					todo(i, true)
				}
				index++
			case Map:
				if ctx.err(i, "missing key in map literal") {
					return true
				}
			case Slice:
				if _, ok := m[index]; ok {
					todo(i, true)
				}

				m[index] = struct{}{}
				if i.Expression.check(ctx) {
					return true
				}

				if v := i.Expression.Value; v != nil && et != nil && !v.AssignableTo(et) {
					if ctx.compositeLiteralValueFail(i, v, et) {
						return true
					}
				}
				index++
			case Struct:
				unkeyed = true
				if keyed {
					todo(i, true) // mix
				}
				if int(index) >= numField {
					todo(i, true)
					break
				}

				f := t.Field(int(index))
				if i.Expression.check(ctx) {
					return true
				}

				if v := i.Expression.Value; v != nil && f.Type != nil && !v.AssignableTo(f.Type) {
					todo(i, true)
				}
				index++
			default:
				panic("internal error")
			}
		case 1: // Expression ':' Expression
			switch tk {
			case Array:
				if i.Expression.check(ctx) {
					return true
				}

				if v := i.Expression.Value; v != nil && !v.nonNegativeInteger() {
					if ctx.err(i.Expression, "index must be non-negative integer constant") {
						return true
					}
				}

				todo(i)
			case Map:
				todo(i)
			case Slice:
				if i.Expression.check(ctx) {
					return true
				}

				if v := i.Expression.Value; v != nil && !v.nonNegativeInteger() {
					if ctx.err(i.Expression, "index must be non-negative integer constant") {
						return true
					}
				}

				todo(i)
			case Struct:
				keyed = true
				if unkeyed {
					todo(i, true) // mix
				}
			default:
				panic("internal error")
			}
		case 2: // Expression ':' LBraceCompLitValue
			switch tk {
			case Array:
				if i.Expression.check(ctx) {
					return true
				}

				if v := i.Expression.Value; v != nil && !v.nonNegativeInteger() {
					todo(i.Expression, true) // index must be non-negative integer constant
				}

				if v := i.Expression.Value; v != nil {
					switch {
					case v.Kind() == ConstValue:
						c := v.Const().Convert(ctx.intType)
						if c == nil {
							todo(i, true)
							break
						}

						j := c.Const().(*intConst).val //TODO .IntValue()
						if j < 0 {
							todo(i, true)
							break
						}

						if j >= len {
							todo(i, true)
							break
						}

						index = j
					default:
						todo(i, true)
					}
				}
				if et == nil {
					break
				}

				et2 := et
				if et2.Kind() == Ptr && et2.Elem().Kind() == Struct {
					et2 = et2.Elem()
				}
				if i.LBraceCompLitValue.check(ctx, et2) {
					return true
				}

				if t := i.LBraceCompLitValue.Type; t != nil && !t.AssignableTo(et2) {
					todo(i, true)
				}
				index++
			case Map:
				todo(i)
			case Slice:
				if i.Expression.check(ctx) {
					return true
				}

				if v := i.Expression.Value; v != nil && !v.nonNegativeInteger() {
					todo(i.Expression, true) // index must be non-negative integer constant
				}

				todo(i)
			case Struct:
				keyed = true
				if unkeyed {
					todo(i, true) // mix
				}
				todo(i)
			default:
				panic("internal error")
			}
		case 3: // LBraceCompLitValue
			switch tk {
			case Array:
				if index >= len {
					todo(i, true)
				}

				if _, ok := m[index]; ok {
					todo(i, true)
				}

				m[index] = struct{}{}
				if et == nil {
					break
				}

				et2 := et
				if et2.Kind() == Ptr && et2.Elem().Kind() == Struct {
					et2 = et2.Elem()
				}
				if i.LBraceCompLitValue.check(ctx, et2) {
					return true
				}

				index++
			case Map:
				todo(i, true)
			case Slice:
				if _, ok := m[index]; ok {
					todo(i, true)
				}

				m[index] = struct{}{}
				if et == nil {
					break
				}

				et2 := et
				if et2.Kind() == Ptr && et2.Elem().Kind() == Struct {
					et2 = et2.Elem()
				}
				if i.LBraceCompLitValue.check(ctx, et2) {
					return true
				}

				index++
			case Struct:
				todo(i, true) // invalid
			default:
				panic("internal error")
			}
		default:
			panic("internal error")
		}
	}
	return false
}

// --------------------------------------------------------- LBraceCompLitValue

func (n *LBraceCompLitValue) check(ctx *context, t Type) (stop bool) {
	if n == nil || t == nil || t.Kind() == Invalid {
		return false
	}

	n.Type = t
	switch n.Case {
	case 0: // LBrace '}'
		// nop
	case 1: // LBrace LBraceCompLitItemList CommaOpt '}'
		stop = n.LBraceCompLitItemList.check(ctx, t)
	default:
		panic("internal error")
	}
	return stop
}

// -------------------------------------------------------------------- MapType

func (n *MapType) check(ctx *context) (stop bool) {
	if n == nil {
		return false
	}

	stack := ctx.stack
	ctx.stack = nil

	defer func() { ctx.stack = stack }()

	done, stop := n.guard.check(ctx, nil)
	if done || stop {
		return stop
	}

	defer n.guard.done(nil)

	// "map" '[' Typ ']' Typ
	stop = n.Typ.check(ctx) || n.Typ2.check(ctx)
	n.Type = newMapType(ctx, n.Typ.Type, n.Typ2.Type)
	return stop
}

// -------------------------------------------------------------------- Operand

func (n *Operand) isIdentifier() (xc.Token, bool) {
	if n.GenericArgumentsOpt != nil {
		todo(n)
	}
	return n.Token, n.Case == 4 // IDENTIFIER GenericArgumentsOpt
}

func (n *Operand) check(ctx *context) (stop bool) {
	if n == nil {
		return false
	}

	if n.Expression.check(ctx) ||
		n.FuncType.check(ctx) ||
		n.StatementList.check(ctx) {
		return true
	}

	switch n.Case {
	case 0: // '(' Expression ')'
		n.Value = n.Expression.Value
	case 1: // '(' TypeLiteral ')'
		todo(n)
	case 2: // BasicLiteral
		n.Value = n.BasicLiteral.Value
	case 3: // FuncType LBrace StatementList '}'
		todo(n)
	case 4: // IDENTIFIER GenericArgumentsOpt
		if n.GenericArgumentsOpt != nil {
			todo(n)
		}
		if n.Token.Val == idUnderscore {
			if ctx.err(n, "cannot use _ as value") {
				return true
			}

			break
		}

		d := n.resolutionScope.mustLookup(ctx, n.Token, n.fileScope)
		switch isPredeclared := ctx.isPredeclared(d); x := d.(type) {
		case nil:
			// nop
		case *ConstDeclaration:
			if x.check(ctx) {
				return true
			}

			if isPredeclared {
				switch d.Name() {
				case idFalse:
					n.Value = ctx.falseValue
				case idIota:
					n.Value = ctx.iota
					if n.Value == nil {
						ctx.err(n.Token, "undefined: iota")
					}
				case idTrue:
					n.Value = ctx.trueValue
				default:
					panic("internal error")
				}
				break
			}

			n.Value = x.Value
		case *FuncDeclaration:
			if isPredeclared || x.unsafe {
				n.Value = newPredeclaredFunctionValue(d)
				break
			}

			if x.check(ctx) {
				return true
			}

			if t := x.Type; t != nil {
				n.Value = newRuntimeValue(t)
			}
		case *ImportDeclaration:
			n.Value = newPackageValue(x.Package)
		case *TypeDeclaration:
			if x.check(ctx) {
				return true
			}

			n.Value = newTypeValue(x)
		case *VarDeclaration:
			if x.check(ctx) {
				return true
			}

			if isPredeclared {
				switch d.Name() {
				case idNil:
					n.Value = ctx.nilValue
				default:
					panic("internal error")
				}
				break
			}

			if t := x.Type; t != nil {
				n.Value = newAddressableValue(t)
			}
		default:
			//dbg("%s: %T", position(n.Pos()), d)
			todo(n)
		}
	default:
		panic("internal error")
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
					todo(n)
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

func (n *Parameters) check(ctx *context) (stop bool) {
	if n == nil {
		return false
	}

	for l := n.ParameterDeclList; l != nil; l = l.ParameterDeclList {
		switch i := l.ParameterDecl; i.Case {
		case
			0, // "..." Typ
			1, // IDENTIFIER "..." Typ
			2: // IDENTIFIER Typ
			stop = stop || i.Typ.check(ctx)
		case 3: // Typ
			if i.isParamName {
				stop = stop || i.typ.check(ctx)
				break
			}

			stop = stop || i.Typ.check(ctx)
		default:
			panic("internal error")
		}
	}
	return stop
}

// ---------------------------------------------------------- PrimaryExpression

func (n *PrimaryExpression) checkConversion(node Node, t Type, arg Value) {
	if t == nil || arg == nil {
		return
	}

	ctx := t.context()
	if t := arg.Type(); t != nil && !t.ConvertibleTo(t) {
		ctx.err(node, "cannot convert type %s to %s", arg.Type(), t)
		return
	}

	switch arg.Kind() {
	case ConstValue:
		c := arg.Const()
		if n.Value = c.Convert(t); n.Value == nil {
			ctx.constConversionFail(node, t, c)
		}
	case NilValue:
		if n.Value = arg.Convert(t); n.Value == nil {
			todo(n, true)
		}
	case RuntimeValue:
		if n.Value = arg.Convert(t); n.Value == nil {
			todo(n, true)
		}
	case TypeValue:
		todo(n)
	default:
		panic("internal error")
	}
}

func (n *PrimaryExpression) checkCall(ft Type, skip int) (stop bool) {
	if ft == nil {
		return false
	}

	args, ddd := n.Call.args()
	switch {
	case ft.IsVariadic():
		todo(n)
	default:
		if len(args) != ft.NumIn()-skip {
			todo(n) //TODO special case single tuple arg.
		}
		if ddd {
			todo(n)
		}
		for i, arg := range args {
			if arg != nil && !arg.AssignableTo(ft.In(i+skip)) {
				todo(n, true)
			}
		}
	}

	switch ft.NumOut() {
	case 0:
		todo(n)
	case 1:
		n.Value = newRuntimeValue(ft.Out(0))
	default:
		todo(n)
	}
	return false
}

func (n *PrimaryExpression) checkSelector(ctx *context, t Type, nm xc.Token) []Selector {
	t0 := t
	ignoreMethods := false
	if t.Kind() == Ptr {
		ignoreMethods = t.isNamed()
		t = t.Elem()
		switch t.Kind() {
		case Ptr, Interface:
			todo(n, true) // fail
			return nil
		}
	}

	var paths [][]Selector
	best := mathutil.MaxInt
	var w func(Type, []Selector)
	w = func(t Type, path []Selector) {
		if t == nil {
			return
		}

		if m := t.MethodByName(nm.Val); !ignoreMethods && m != nil && !m.merged {
			s := append(path[:len(path):len(path)], m)
			paths = append(paths, s)
			best = mathutil.Min(best, len(s))
			return
		}

		if t.Kind() != Struct {
			return
		}

		if f := t.FieldByName(nm.Val); f != nil {
			s := append(path[:len(path):len(path)], f)
			paths = append(paths, s)
			best = mathutil.Min(best, len(s))
			return
		}

		if len(path)+1 > best {
			return
		}

		for i := 0; i < t.NumField(); i++ {
			f := t.Field(i)
			t := f.Type
			if t.Kind() == Ptr {
				t = t.Elem()
			}
			w(t, append(path, f))
		}
	}

	w(t, nil)
	switch len(paths) {
	case 0:
		ctx.err(nm, "%s undefined (type %s has no field or method %s)", nm.S(), t0, nm.S())
		return nil
	case 1:
		return paths[0]
	}

	for i, v := range paths {
		if len(v) == best {
			for _, v := range paths[i+1:] {
				if len(v) == best {
					todo(n, true) // fail
					return nil
				}
			}

			return v
		}
	}
	panic("internal error")
}

func (n *PrimaryExpression) check(ctx *context) (stop bool) {
	if n == nil {
		return false
	}

	if n.Call.check(ctx) ||
		n.Expression.check(ctx) ||
		n.ExpressionOpt.check(ctx) ||
		n.ExpressionOpt2.check(ctx) ||
		n.ExpressionOpt3.check(ctx) ||
		n.Operand.check(ctx) ||
		n.PrimaryExpression.check(ctx) ||
		n.TypeLiteral.check(ctx) {
		return true
	}

	switch n.Case {
	case 0: // Operand
		n.Value = n.Operand.Value
	case 1: // CompLitType LBraceCompLitValue
		if ct := n.CompLitType; ct.Case == 0 { // ArrayType
			ct.ArrayType.items = n.LBraceCompLitValue.items
			ct.ArrayType.itemsSet = true
		}
		if n.CompLitType.check(ctx) {
			return true
		}

		if t := n.CompLitType.Type; t != nil {
			if n.LBraceCompLitValue.check(ctx, t) {
				return true
			}

			n.Value = newAddressableValue(t)
		}
	case 2: // PrimaryExpression '.' '(' "type" ')'
		todo(n)
	case 3: // PrimaryExpression '.' '(' Typ ')'
		todo(n)
	case 4: // PrimaryExpression '.' IDENTIFIER
		v := n.PrimaryExpression.Value
		if v == nil {
			break
		}

		if n.Token2.Val == idUnderscore {
			todo(n, true)
			break
		}

		switch v.Kind() {
		case PackageValue:
			t := n.Token2
			if !isExported(t.Val) {
				todo(n, true)
				break
			}

			if d := v.Package().Scope.Bindings[t.Val]; d != nil {
				switch x := d.(type) {
				case *ConstDeclaration:
					n.Value = x.Value
				case *FuncDeclaration:
					if x.unsafe {
						n.Value = newPredeclaredFunctionValue(d)
						break
					}

					if t := x.Type; t != nil {
						n.Value = newRuntimeValue(t)
					}
				default:
					//dbg("%s: %T", position(n.Pos()), d)
					todo(n)
				}
				break
			}

			ctx.err(t, "undefined: %s.%s", v.Package().Name, t.S())
		case RuntimeValue:
			t := v.Type()
			if t == nil {
				break
			}

			path := n.checkSelector(ctx, t, n.Token2)
			if path == nil {
				break
			}

			rv := v.(*runtimeValue)
			root := rv.root
			if root == nil {
				root = v
			}
			n.Value = newSelectorValue(path[len(path)-1].typ(), root, append(rv.path0, path[len(path)-1]), append(rv.path, path...))
		default:
			todo(n)
		}
	case 5: // PrimaryExpression '[' Expression ']'
		a := n.PrimaryExpression.Value
		if a == nil {
			break
		}

		// For a of pointer to array type:
		//
		// · a[x] is shorthand for (*a)[x]
		if a.Type().Kind() == Ptr && a.Type().Elem().Kind() == Array {
			a = newRuntimeValue(a.Type().Elem())
		}

		x := n.Expression.Value
		if x == nil {
			break
		}

		// If a is not a map:
		//
		// · the index x must be of integer type or untyped; it is in
		//   range if 0 <= x < len(a), otherwise it is out of range
		//
		// · a constant index must be non-negative and representable by
		//   a value of type int
		if a.Type().Kind() != Map {
			if x.Kind() == RuntimeValue && !x.Type().IntegerType() {
				if ctx.err(n.Expression, "non-integer array index") {
					return true
				}

				break
			}

			if x.Kind() == ConstValue {
				x2 := ctx.mustConvertConst(n.Expression, ctx.intType, x.Const())
				if x2 == nil {
					break
				}

				if x2 != nil && !x2.nonNegativeInteger() {
					todo(n.Expression, true)
					break
				}

				todo(n) // Check const array index is in range.
			}
		}

		todo(n)
	case 6: // PrimaryExpression '[' ExpressionOpt ':' ExpressionOpt ':' ExpressionOpt ']'
		todo(n)
	case 7: // PrimaryExpression '[' ExpressionOpt ':' ExpressionOpt ']'
		todo(n)
	case 8: // PrimaryExpression Call
		v := n.PrimaryExpression.Value
		if v == nil {
			break
		}

		switch v.Kind() {
		case RuntimeValue:
			t := v.Type()
			if t == nil {
				break
			}

			if t.Kind() != Func {
				todo(n, true)
				break
			}

			if d := v.PredeclaredFunction(); d != nil {
				switch {
				case ctx.isPredeclared(d):
					todo(n)
				default:
					if !d.(*FuncDeclaration).unsafe {
						panic("internal error")
					}

					switch d.Name() {
					case idAlignof:
						todo(n)
					case idOffsetof:
						n.Value = offsetof(ctx, n.Call)
					case idSizeof:
						n.Value = sizeof(ctx, n.Call)
					default:
						panic("internal error")
					}
				}
				break
			}

			return n.checkCall(t, 0)
		case TypeValue: // Conversion.
			t := v.Type()
			args, ddd := n.Call.args()
			if ddd {
				todo(n, true)
			}
			switch len(args) {
			case 0:
				todo(n, true)
			case 1:
				n.checkConversion(n.Call.ArgumentList.Argument, t, args[0])
			default:
				todo(n, true)
			}
		default:
			//dbg("", v.Kind())
			todo(n, true)
		}
	case 9: // PrimaryExpression CompLitValue
		v := n.PrimaryExpression.Value
		if v == nil {
			break
		}

		switch v.Kind() {
		case TypeValue:
			if n.CompLitValue.check(ctx, v.Type()) {
				return true
			}

			n.Value = newRuntimeValue(v.Type())
		default:
			todo(n, true)
		}
	case 10: // TypeLiteral '(' Expression CommaOpt ')'
		n.checkConversion(n.Expression, n.TypeLiteral.Type, n.Expression.Value)
	default:
		panic("internal error")
	}
	return false
}

func (n *PrimaryExpression) isIdentifier() (xc.Token, bool) {
	if n.Case != 0 { // Operand
		return xc.Token{}, false
	}

	return n.Operand.isIdentifier()
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
				todo(n)
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

func (n *ReceiverOpt) check(ctx *context) (stop bool) {
	if n == nil {
		return false
	}

	if nm := n.nm; nm.IsValid() {
		n.Type = n.resolutionScope.mustLookupLocalTLDType(ctx, nm)
	}
	return false
}

// ------------------------------------------------------------------ ResultOpt

func (n *ResultOpt) check(ctx *context) (stop bool) {
	if n == nil {
		return false
	}

	stack := ctx.stack
	ctx.stack = nil
	defer func() { ctx.stack = stack }()
	if n.Parameters.check(ctx) || n.Typ.check(ctx) {
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
			n.Type = ctx.voidType
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

func (n *Signature) check(ctx *context) (stop bool) {
	if n == nil {
		return false
	}

	if n.Parameters.check(ctx) || n.ResultOpt.check(ctx) {
		return true
	}

	if o := n.ResultOpt; o != nil {
		n.Type = o.Type
	}
	return false
}

// ------------------------------------------------------------------ SliceType

func (n *SliceType) check(ctx *context) (stop bool) {
	if n == nil {
		return false
	}

	stack := ctx.stack
	ctx.stack = nil

	defer func() { ctx.stack = stack }()

	done, stop := n.guard.check(ctx, nil)
	if done || stop {
		return stop
	}

	defer n.guard.done(nil)

	// '[' ']' Typ
	if n.Typ.check(ctx) {
		return true
	}

	n.Type = newSliceType(ctx, n.Typ.Type)
	return false
}

// ------------------------------------------------------------------ Statement

func (n *Statement) check(ctx *context) (stop bool) {
	todo(n)
	return false
}

// -------------------------------------------------------------- StatementList

func (n *StatementList) check(ctx *context) (stop bool) {
	for ; n != nil; n = n.StatementList {
		i := n.Statement
		if i.check(ctx) {
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
			lx.fileScope,
			lx.resolutionScope,
		))
	case 1: // IdentifierList Typ TagOpt
		for l := n.IdentifierList; l != nil; l = l.IdentifierList {
			lx.scope.declare(lx, newFieldDeclaration(
				l.ident(),
				n.Typ.normalize(),
				false,
				nil,
				tag,
				lx.fileScope,
				lx.resolutionScope,
			))
		}
	case 2: // QualifiedIdent TagOpt
		lx.scope.declare(lx, newFieldDeclaration(
			n.QualifiedIdent.i(),
			nil,
			false,
			n.QualifiedIdent,
			tag,
			lx.fileScope,
			lx.resolutionScope,
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

func (n *StructType) check(ctx *context) (stop bool) {
	if n == nil {
		return false
	}

	done, stop := n.guard.check(ctx, nil)
	if done || stop {
		return stop
	}

	defer func() {
		n.guard.done(nil)
		if n.Type != nil {
			t := n.Type.(*structType)
			t.align, t.fieldAlign, t.size = measure(t)
		}
	}()

	switch n.Case {
	case 0: // "struct" LBrace '}'
		n.Type = newStructType(ctx, nil, nil)
	case 1: // "struct" LBrace StructFieldDeclList SemicolonOpt '}'
		s := n.fields
		a := make(fields, 0, len(s.Bindings)+len(s.Unbound))
		for _, d := range s.Bindings {
			a = append(a, d.(*FieldDeclaration))
		}
		for _, d := range s.Unbound {
			a = append(a, d.(*FieldDeclaration))
		}
		sort.Sort(a)
		sf := make([]StructField, 0, len(a))
		var mta []Method
		for _, f := range a {
			if f.check(ctx) {
				return true
			}

			if ft := f.Type; ft != nil && f.isAnonymous && ft.NumMethod() != 0 {
				m := map[int]struct{}{}
				for i := 0; i < ft.NumMethod(); i++ {
					mt := ft.Method(i)
					if _, ok := m[mt.Name]; ok {
						continue
					}

					m[mt.Name] = struct{}{}
					var mt2 Method
					switch {
					case ft.Kind() == Interface:
						todo(zeroNode) // Merge interface method, must synthesize new method w/ receiver.
					default:
						mt2 = *mt
						mt2.merged = true
					}
					mt2.Index = len(mta)
					mta = append(mta, mt2)
				}
			}

			pkgPath := 0
			if !f.isExported {
				pkgPath = n.pkgPath
			}
			tag := 0
			if f.tag != nil {
				tag = int(f.tag.(stringID))
			}
			t := f.Type
			if f.isAnonymousPtr {
				t = newPtrType(ctx, t)
			}
			sf = append(sf, StructField{
				f.name,
				pkgPath,
				t,
				tag,
				0, // Offset is set in newStructType.
				nil,
				f.isAnonymous,
			})
		}
		n.Type = newStructType(ctx, sf, mta)
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

func (n *Typ) check(ctx *context) (stop bool) {
	if n == nil {
		return false
	}

	done, stop := n.guard.check(ctx, nil)
	if done || stop {
		return stop
	}

	defer n.guard.done(nil)

	t := n.normalize()
	switch t.Case {
	case 0: // '(' Typ ')'
		panic("internal error")
	case 1: // '*' Typ
		t2 := t.Typ.normalize()
		stack := ctx.stack
		ctx.stack = nil
		defer func() { ctx.stack = stack }()
		stop = t2.check(ctx)
		n.Type = newPtrType(ctx, t2.Type)
	case 2: // ArrayType
		stop = t.ArrayType.check(ctx)
		n.Type = t.ArrayType.Type
	case 3: // ChanType
		stop = t.ChanType.check(ctx)
		n.Type = t.ChanType.Type
	case 4: // FuncType
		stop = t.FuncType.check(ctx)
		n.Type = t.FuncType.Type
	case 5: // InterfaceType
		stop = t.InterfaceType.check(ctx)
		n.Type = t.InterfaceType.Type
	case 6: // MapType
		stop = t.MapType.check(ctx)
		n.Type = t.MapType.Type
	case 7: // QualifiedIdent GenericArgumentsOpt
		if t.GenericArgumentsOpt != nil {
			todo(n)
		}
		if t := n.resolutionScope.mustLookupType(ctx, n.QualifiedIdent, n.fileScope); t != nil {
			n.Type = t
			stop = t.check(ctx)
		}
	case 8: // SliceType
		stop = t.SliceType.check(ctx)
		n.Type = t.SliceType.Type
	case 9: // StructType
		stop = t.StructType.check(ctx)
		n.Type = t.StructType.Type
	default:
		panic("internal error")
	}
	return stop
}

// ---------------------------------------------------------------- TypeLiteral

func (n *TypeLiteral) check(ctx *context) (stop bool) {
	if n == nil {
		return false
	}

	done, stop := n.guard.check(ctx, nil)
	if done || stop {
		return stop
	}

	defer n.guard.done(nil)

	switch n.Case {
	case 0: // '*' TypeLiteral
		todo(n)
	case 1: // ArrayType
		todo(n)
	case 2: // ChanType
		todo(n)
	case 3: // FuncType
		todo(n)
	case 4: // InterfaceType
		todo(n)
	case 5: // MapType
		todo(n)
	case 6: // SliceType
		stop = n.SliceType.check(ctx)
		n.Type = n.SliceType.Type
	case 7: // StructType
		todo(n)
	default:
		panic("internal error")
	}
	return stop
}

// ------------------------------------------------------------ UnaryExpression

func (n *UnaryExpression) isIdentifier() (xc.Token, bool) {
	if n.Case != 7 { // PrimaryExpression
		return xc.Token{}, false
	}

	return n.PrimaryExpression.isIdentifier()
}

func (n *UnaryExpression) check(ctx *context) (stop bool) {
	if n == nil {
		return false
	}

	if n.UnaryExpression.check(ctx) || n.PrimaryExpression.check(ctx) {
		return true
	}

	var v Value
	if n.Case != 7 { // PrimaryExpression
		v = n.UnaryExpression.Value
		if v == nil {
			return false
		}
	}

	switch n.Case {
	case 0: // '!' UnaryExpression
		todo(n)
	case 1: // '&' UnaryExpression
		switch v.Kind() {
		default:
			//dbg("", v.Kind())
			todo(n)
		}
	case 2: // '*' UnaryExpression
		switch v.Kind() {
		case RuntimeValue:
			if t := v.Type(); t != nil {
				if t.Kind() != Ptr {
					todo(n, true)
					break
				}

				n.Value = newRuntimeValue(t.Elem())
			}
		case TypeValue:
			n.Value = newTypeValue(newPtrType(ctx, v.Type()))
		default:
			//dbg("", v.Kind())
			todo(n)
		}
	case 3: // '+' UnaryExpression
		todo(n)
	case 4: // '-' UnaryExpression
		switch v.Kind() {
		case ConstValue:
			n.Value = v.Neg(ctx.Context, n.UnaryExpression)
		default:
			//dbg("", v.Kind())
			todo(n)
		}
	case 5: // '^' UnaryExpression
		switch v.Kind() {
		case ConstValue:
			n.Value = v.Cpl(ctx.Context, n.UnaryExpression)
		default:
			//dbg("", v.Kind())
			todo(n)
		}
	case 6: // "<-" UnaryExpression
		todo(n)
	case 7: // PrimaryExpression
		n.Value = n.PrimaryExpression.Value
	default:
		panic("internal error")
	}
	return false
}
