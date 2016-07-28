// Copyright 2016 The GC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gc

import (
	"bytes"
	"fmt"
	"go/token"
	"path/filepath"
	"sort"
	"strings"
	"unicode"

	"github.com/cznic/internal/buffer"
	"github.com/cznic/mathutil"
	"github.com/cznic/xc"
)

var (
	_ CompositeLiteralItemList = (*CompLitItemList)(nil)
	_ CompositeLiteralItemList = (*LBraceCompLitItemList)(nil)
	_ CompositeLiteralNode     = (*CompLitValue)(nil)
	_ CompositeLiteralNode     = (*LBraceCompLitValue)(nil)
	_ CompositeLiteralNode     = compositeLiteralNode{}
	_ CompositeLiteralValue    = (*CompLitValue)(nil)
	_ CompositeLiteralValue    = (*LBraceCompLitValue)(nil)
)

type compositeLiteralNode struct {
	e *Expression
}

func (n compositeLiteralNode) CompositeLiteralValue() CompositeLiteralValue { return nil }
func (n compositeLiteralNode) Expression() *Expression                      { return n.e }
func (n compositeLiteralNode) Pos() token.Pos                               { return n.e.Pos() }

// Node represents the interface all AST nodes, xc.Token and lex.Char implement.
type Node interface {
	Pos() token.Pos
}

// CompositeLiteralValue represents a composite literal value. Example:
//
//	var v = struct{ i int } /* composite literal value follows */ { 42 }
type CompositeLiteralValue interface {
	gate() *gate
	setType(Type)

	// Items returns the composite literal value item list. It will return
	// nil if the list is empty.
	Items() CompositeLiteralItemList

	// Value returns the Value of the composite literal value.
	Value() Value

	Node
}

// CompositeLiteralItemList represents a list of composite literal items.
// Example:
//
//	var v = []struct{ i int } /* composite literal item list follows */ {{42}, {314}}
type CompositeLiteralItemList interface {
	// Key returns the key of the current composite literal item. It
	// returns nil if the key is not present.
	//
	//	var v = []int{
	//		42,  // Item without a key.
	//		314: 278, // Key 314.
	//	}
	Key() CompositeLiteralNode

	// Value returns the value of the current composite literal item.
	//
	//	var v = []int{
	//		42,  // Value 42.
	//		314: 278, // Value 278.
	//	}
	Val() CompositeLiteralNode

	// Next moves to the next item in the list. It returns nil when there
	// are no more items in the list.
	Next() CompositeLiteralItemList

	Node
}

// CompositeLiteralNode represents the key or value node of a composite literal
// item.
type CompositeLiteralNode interface {
	// CompositeLiteralValue returns the node when it is a composite
	// literal value, or nil otherwise.  Example:
	//
	//	type t struct {
	//		i int
	//	}
	//
	//	var v = map[t]t{
	//		{42}: {314}, // Both key and value are composite literal values.
	//	}
	CompositeLiteralValue() CompositeLiteralValue

	// Expression returns the node when it is an expression, or nil
	// otherwise.  Example:
	//
	//	var v = map[int]int{
	//		42: 314, // Both key and value are expressions.
	//	}
	Expression() *Expression

	Node
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
		n.lenPoisoned = n.Expression.lenPoisoned
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

// --------------------------------------------------------------- ArgumentList

func (n *ArgumentList) node(i int) Node {
	for i != 0 {
		n = n.ArgumentList
		i--
	}
	return n
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
		if n.Typ.check(ctx) {
			return true
		}

		t := n.Typ.Type
		if t == nil {
			break
		}

		n.Type = newArrayType(ctx, t, -1)
		if cv := n.compLitValue; cv != nil {
			stop = (*CompLitValue)(nil).check(ctx, cv, n.Type)
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
						n.Type = newArrayType(ctx, t, val)
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

// ---------------------------------------------------------------------- Block

func (n *Block) check(ctx *context) (stop bool) {
	if n == nil {
		return false
	}

	todo(n)
	return false
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
		n.lenPoisoned = n.lenPoisoned || l.Argument.lenPoisoned
	}
	return false
}

func (n *Call) args() (_ []Value, lenPoisoned, ddd bool) {
	var a []Value
	first := true
	for l := n.ArgumentList; l != nil; l = l.ArgumentList {
		if first && l.Argument.Expression != nil {
			lenPoisoned = l.Argument.Expression.lenPoisoned
			first = false
		}
		a = append(a, l.Argument.Value)
	}
	return a, lenPoisoned, n.Case == 2 // '(' ArgumentList "..." CommaOpt ')'
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

// Key implements CompositeLiteralItemList.
func (n *CompLitItemList) Key() CompositeLiteralNode {
	switch i := n.CompLitItem; i.Case {
	case
		0, // CompLitValue
		3: // Expression
		return nil
	case
		1, // CompLitValue ':' CompLitValue
		2: // CompLitValue ':' Expression
		return i.CompLitValue
	case
		4, // Expression ':' CompLitValue
		5: // Expression ':' Expression
		return compositeLiteralNode{i.Expression}
	default:
		panic("internal error")
	}
}

// Next implements CompositeLiteralItemList.
func (n *CompLitItemList) Next() CompositeLiteralItemList {
	if l := n.CompLitItemList; l != nil {
		return l
	}

	return nil
}

// Val implements CompositeLiteralItemList.
func (n *CompLitItemList) Val() CompositeLiteralNode {
	switch i := n.CompLitItem; i.Case {
	case
		0, // CompLitValue
		4: // Expression ':' CompLitValue
		return i.CompLitValue
	case 1: // CompLitValue ':' CompLitValue
		return i.CompLitValue2
	case
		2, // CompLitValue ':' Expression
		3: // Expression
		return compositeLiteralNode{i.Expression}
	case 5: // Expression ':' Expression
		return compositeLiteralNode{i.Expression2}
	default:
		panic("internal error")
	}
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

// CompositeLiteralValue implements CompositeLiteralNode.
func (n *CompLitValue) CompositeLiteralValue() CompositeLiteralValue { return n }

// Expression implements CompositeLiteralNode.
func (n *CompLitValue) Expression() *Expression { return nil }

// Items implements CompositeLiteralValue.
func (n *CompLitValue) Items() CompositeLiteralItemList {
	if l := n.CompLitItemList; l != nil {
		return l
	}

	return nil
}

// Value implements CompositeLiteralValue.
func (n *CompLitValue) Value() Value {
	if n.Type != nil {
		return newRuntimeValue(n.Type)
	}

	return nil
}

func (n *CompLitValue) gate() *gate    { return &n.guard }
func (n *CompLitValue) setType(t Type) { n.Type = t }

func (*CompLitValue) check(ctx *context, n CompositeLiteralValue, t Type) (stop bool) {
	done, stop := n.gate().check(ctx, nil)
	if done || stop {
		return stop
	}

	defer n.gate().done(nil)

	if n == nil || t == nil || t.Kind() == Invalid {
		return false
	}

	var len, index, maxIndex int64
	var kt, et Type
	var numField int
	var keyed, unkeyed bool
	m := map[mapKey]struct{}{}
	tk := t.Kind()

	switch tk {
	case Array:
		len = t.Len()
		et = t.Elem()
		if len < 0 {
			defer func() {
				t.(*arrayType).len = maxIndex + 1
			}()
		}
	case Map:
		kt = t.Key()
		et = t.Elem()
	case Slice:
		et = t.Elem()
		len = -1
	case Struct:
		numField = t.NumField()
	default:
		ctx.err(n, "invalid type for composite literal: %s", t)
		return false
	}

	n.setType(t)
	for l := n.Items(); l != nil; l = l.Next() {
		var hasKey, cvkValid bool
		var cvk Value
		var field *StructField
		if key := l.Key(); key != nil {
			hasKey = true
			switch cv := key.CompositeLiteralValue(); {
			case cv != nil:
				switch tk {
				case Array:
					todo(cv)
				case Map:
					if kt == nil {
						break
					}

					if (*CompLitValue)(nil).check(ctx, cv, kt) {
						return true
					}

					cvk = cv.Value()
					cvkValid = true
				case Slice:
					todo(cv)
				case Struct:
					todo(cv, true)
				}
			default:
				e := key.Expression()
				switch tk {
				case Array, Slice:
					if e == nil {
						todo(key, true)
						break
					}

					if e.check(ctx) {
						return true
					}

					v := e.Value
					if v == nil {
						break
					}

					if !v.nonNegativeInteger() {
						if ctx.err(e, "index must be non-negative integer constant") {
							return true
						}

						break
					}

					i := v.Const().int()
					if len >= 0 && i >= len {
						todo(e, true)
						//if ctx.err(e, "array index %d out of bounds [0:%d]", i, len) {
						//	return true
						//}

						break
					}

					index = i
					maxIndex = mathutil.MaxInt64(maxIndex, index)
				case Map:
					var v Value
					switch {
					case cvkValid:
						v = cvk
					default:
						if e == nil {
							todo(key, true)
							break
						}

						if e.check(ctx) {
							return true
						}

						v = e.Value
					}
					if v == nil {
						break
					}

					if !v.AssignableTo(kt) {
						switch v.Kind() {
						case ConstValue:
							if ctx.err(e, "cannot use %s (type %s) as type %s in map key", v.Const(), v.Type(), kt) {
								return true
							}
						default:
							todo(e, true)
							// if ctx.err(n, "cannot use value of type %s as type %s in array or slice literal", v.Type(), t) {
							// 	return true
							// }
						}
						break
					}

					if v.Kind() != ConstValue {
						break
					}

					ck := v.Const()
					switch {
					case kt.Kind() == Interface:
						p := buffer.CGet(64)
						buf := bytes.NewBuffer((*p)[:0])
						switch {
						case ck.Untyped():
							fmt.Fprintf(buf, "%s", ck)
						default:
							fmt.Fprintf(buf, "%s(%s)", ck.Type(), ck)
						}
						k := mapKey{i: int64(dict.ID(buf.Bytes()))}
						buffer.Put(p)
						if _, ok := m[k]; ok {
							if ctx.err(e, "duplicate key %s in map literal", dict.S(int(k.i))) {
								return true
							}
						}

						m[k] = struct{}{}
					default:
						var k mapKey
						switch ck.Kind() {
						case StringConst:
							k.i = int64(ck.(*stringConst).val.id())
						default:
							//dbg("", ck.Kind())
							todo(e)
						}
						switch ck.Kind() {
						case StringConst: //TODO more types
							if _, ok := m[k]; ok {
								if ctx.err(e, "duplicate key %s in map literal", ck) {
									return true
								}
							}

							m[k] = struct{}{}
						}
					}
				case Struct:
					fn, ok := e.isIdentifier()
					if !ok {
						if ctx.err(e, "invalid field name in struct initializer") {
							return true
						}

						continue
					}

					field = t.FieldByName(fn.Val)
					if field == nil {
						todo(e, true) // unknown field
						continue
					}

					k := mapKey{i: int64(fn.Val)}
					if _, ok := m[k]; ok {
						if ctx.err(e, "duplicate field name in struct literal: %s", fn.S()) {
							return true
						}

						break
					}

					m[k] = struct{}{}
				}
			}
		}
		if val := l.Val(); val != nil {
			var v Value
			switch cv := val.CompositeLiteralValue(); {
			case cv != nil:
				switch tk {
				case Array, Slice, Map:
					if et == nil {
						break
					}

					if (*CompLitValue)(nil).check(ctx, cv, et) {
						return true
					}

					v = cv.Value()
				case Struct:
					if ctx.err(cv, "missing type in composite literal") {
						return true
					}

					continue
				}
			default:
				e := val.Expression()
				if e.check(ctx) {
					return true
				}

				v = e.Value
			}
			switch tk {
			case Array:
				if len >= 0 && index >= len {
					if ctx.err(val, "array index %d out of bounds [0:%d]", index, len) {
						return true
					}
				}
				fallthrough
			case Slice:
				k := mapKey{i: index}
				if _, ok := m[k]; ok {
					todo(val, true) // dup index
				}
				m[k] = struct{}{}
				if v != nil && !v.AssignableTo(et) {
					if ctx.compositeLiteralValueFail(val, v, et) {
						return true
					}
				}
				maxIndex = mathutil.MaxInt64(maxIndex, index)
				index++
			case Map:
				if v != nil && !v.AssignableTo(et) {
					todo(val, true)
				}
			case Struct:
				switch {
				case hasKey:
					keyed = true
					if v != nil && field.Type != nil && !v.AssignableTo(field.Type) {
						todo(val, true)
					}
				default:
					unkeyed = true
					if int(index) >= numField {
						if ctx.err(val, "too many values in struct initializer") {
							return true
						}

						break
					}

					f := t.Field(int(index))
					if v != nil && f.Type != nil && !v.AssignableTo(f.Type) {
						todo(val, true) // assign fail
					}
					index++
				}
				if keyed && unkeyed {
					if ctx.err(val, "mixture of field:value and value initializers") {
						return true
					}
				}

			}
		}
	}
	return false
}

// ------------------------------------------------------------------ ConstDecl

func (n *ConstDecl) check(ctx *context) (stop bool) {
	if n == nil {
		return false
	}

	todo(n)
	return false
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
		lx.declarationScope.declare(lx, d)
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

	switch {
	case n.Case == 0: // UnaryExpression
		n.lenPoisoned = n.UnaryExpression.lenPoisoned
	default:
		n.lenPoisoned = n.Expression.lenPoisoned || n.Expression2.lenPoisoned
	}

	switch n.Case {
	case 0: // UnaryExpression
		n.Value = n.UnaryExpression.Value
	case 1: // Expression '%' Expression
		n.Value = n.Expression.Value.mod(n.Token, n.Expression2.Value)
	case 2: // Expression '&' Expression
		n.Value = n.Expression.Value.and(n.Token, n.Expression2.Value)
	case 3: // Expression '*' Expression
		n.Value = n.Expression.Value.mul(n.Token, n.Expression2.Value)
	case 4: // Expression '+' Expression
		n.Value = n.Expression.Value.add(n.Token, n.Expression2.Value)
	case 5: // Expression '-' Expression
		n.Value = n.Expression.Value.sub(n.Token, n.Expression2.Value)
	case 6: // Expression '/' Expression
		n.Value = n.Expression.Value.div(n.Token, n.Expression2.Value)
	case 7: // Expression '<' Expression
		n.Value = n.Expression.Value.lt(n.Token, n.Expression2.Value)
	case 8: // Expression '>' Expression
		n.Value = n.Expression.Value.gt(n.Token, n.Expression2.Value)
	case 9: // Expression '^' Expression
		n.Value = n.Expression.Value.xor(n.Token, n.Expression2.Value)
	case 10: // Expression '|' Expression
		n.Value = n.Expression.Value.or(n.Token, n.Expression2.Value)
	case 11: // Expression "&&" Expression
		n.Value = n.Expression.Value.boolAnd(n.Token, n.Expression2.Value)
	case 12: // Expression "&^" Expression
		n.Value = n.Expression.Value.andNot(n.Token, n.Expression2.Value)
	case 13: // Expression "==" Expression
		n.Value = n.Expression.Value.eq(n.Token, n.Expression2.Value)
	case 14: // Expression ">=" Expression
		n.Value = n.Expression.Value.ge(n.Token, n.Expression2.Value)
	case 15: // Expression "<=" Expression
		todo(n)
	case 16: // Expression "<<" Expression
		n.Value = n.Expression.Value.lsh(n.Token, n.Expression2.Value)
	case 17: // Expression "!=" Expression
		n.Value = n.Expression.Value.neq(n.Token, n.Expression2.Value)
	case 18: // Expression "||" Expression
		n.Value = n.Expression.Value.boolOr(n.Token, n.Expression2.Value)
	case 19: // Expression ">>" Expression
		n.Value = n.Expression.Value.rsh(n.Token, n.Expression2.Value)
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

// ------------------------------------------------------------- ExpressionList

func (n *ExpressionList) check(ctx *context) (stop bool) {
	if len(n.list) != 0 {
		panic("internal error")
	}

	for ; n != nil; n = n.ExpressionList {
		if n.Expression.check(ctx) {
			return true
		}

		n.list = append(n.list, n.Expression)
	}
	return false
}

// ---------------------------------------------------------- ExpressionListOpt

func (n *ExpressionListOpt) check(ctx *context) (stop bool) {
	if n == nil {
		return false
	}

	return n.ExpressionList.check(ctx)
}

//--------------------------------------------------------------- ExpressionOpt

func (n *ExpressionOpt) check(ctx *context) (stop bool) {
	if n == nil {
		return false
	}

	n.Expression.check(ctx)
	n.lenPoisoned = n.Expression.lenPoisoned
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

	var in []Type
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
	n.Type = newFuncType(ctx, 0, in, n.Signature.Type, false, isVariadic)
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
	lx.imports = append(lx.imports, newImportDeclaration(nm, pos, bl, dict.SID(ip)))
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
				qi := id.QualifiedIdent
				if d := qi.resolutionScope.mustLookupQI(ctx, qi, qi.fileScope); d != nil {
					ctx.errf = func() bool {
						ctx.err(qi, "interface type loop involving %s", qi.str())
						return true
					}

					switch x := d.(type) {
					case *TypeDeclaration:
						if qi.checkTypeDeclaration(ctx, x) {
							continue
						}

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
							if ctx.err(id, "interface contains embedded non-interface %s", qi.str()) {
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

// Key implements CompositeLiteralItemList.
func (n *LBraceCompLitItemList) Key() CompositeLiteralNode {
	switch i := n.LBraceCompLitItem; i.Case {
	case
		0, // Expression
		3: // LBraceCompLitValue
		return nil
	case
		1, // Expression ':' Expression
		2: // Expression ':' LBraceCompLitValue
		return compositeLiteralNode{i.Expression}
	case
		4, // LBraceCompLitValue ':' Expression
		5: // LBraceCompLitValue ':' LBraceCompLitValue
		return i.LBraceCompLitValue
	default:
		panic("internal error")
	}
}

// Next implements CompositeLiteralItemList.
func (n *LBraceCompLitItemList) Next() CompositeLiteralItemList {
	if l := n.LBraceCompLitItemList; l != nil {
		return l
	}

	return nil
}

// Val implements CompositeLiteralItemList.
func (n *LBraceCompLitItemList) Val() CompositeLiteralNode {
	switch i := n.LBraceCompLitItem; i.Case {
	case
		0, // Expression
		4: // LBraceCompLitValue ':' Expression
		return compositeLiteralNode{i.Expression}
	case 1: // Expression ':' Expression
		return compositeLiteralNode{i.Expression2}
	case
		2, // Expression ':' LBraceCompLitValue
		3: // LBraceCompLitValue
		return i.LBraceCompLitValue
	case 5: // LBraceCompLitValue ':' LBraceCompLitValue
		return i.LBraceCompLitValue2
	default:
		panic("internal error")
	}
}

// --------------------------------------------------------- LBraceCompLitValue

// CompositeLiteralValue implements CompositeLiteralNode.
func (n *LBraceCompLitValue) CompositeLiteralValue() CompositeLiteralValue { return n }

// Expression implements CompositeLiteralNode.
func (n *LBraceCompLitValue) Expression() *Expression { return nil }

// Items implements CompositeLiteralValue.
func (n *LBraceCompLitValue) Items() CompositeLiteralItemList {
	if l := n.LBraceCompLitItemList; l != nil {
		return l
	}

	return nil
}

// Value implements CompositeLiteralValue.
func (n *LBraceCompLitValue) Value() Value {
	if n.Type != nil {
		return newRuntimeValue(n.Type)
	}

	return nil
}

func (n *LBraceCompLitValue) gate() *gate    { return &n.guard }
func (n *LBraceCompLitValue) setType(t Type) { n.Type = t }

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

func (n *Operand) checkDeclaration(ctx *context, d Declaration) (stop bool) {
	if ctx.isPredeclared(d) && ctx.pkg.ImportPath != "" {
		return false
	}

	return d.check(ctx)
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
		n.lenPoisoned = n.Expression.lenPoisoned
	case 1: // '(' TypeLiteral ')'
		if n.TypeLiteral.check(ctx) {
			return true
		}

		if t := n.TypeLiteral.Type; t != nil {
			n.Value = newTypeValue(t)
		}
	case 2: // BasicLiteral
		n.Value = n.BasicLiteral.Value
	case 3: // FuncType LBrace StatementList '}'
		n.Value = newRuntimeValue(n.FuncType.Type)
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
		if d == nil {
			break
		}

		if n.checkDeclaration(ctx, d) {
			return true
		}

		switch isPredeclared := ctx.isPredeclared(d); x := d.(type) {
		case *ConstDeclaration:
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
			n.Value = newDeclarationValue(d, x.Type)
		case *ImportDeclaration:
			n.Value = newPackageValue(x)
		case *TypeDeclaration:
			n.Value = newTypeValue(x)
		case *VarDeclaration:
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
			pd.typ = pd.Typ
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
			pd.typ = pd.Typ
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
			if lx.declarationScope.Bindings[nm.Val] != nil {
				lx.err(v, "duplicate argument %s", nm.S())
				continue
			}

			lx.declarationScope.declare(lx, newParamaterDeclaration(nm, v.typ, v.isVariadic, lx.lookahead.Pos()))
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
			ctx.mustConvertNil(node, t)
		}
	case RuntimeValue:
		if n.Value = arg.Convert(t); n.Value == nil {
			switch {
			case t.Kind() == Interface:
				arg.Type().implementsFailed(ctx, node, "cannot convert %s to type %s:", t)
			default:
				//dbg("", arg.Type(), t)
				todo(n, true)
			}
		}
	case TypeValue:
		todo(n)
	default:
		panic("internal error")
	}
}

func (n *PrimaryExpression) checkCall(ctx *context, ft Type, skip int) (stop bool) {
	if ft == nil {
		return false
	}

	args, _, ddd := n.Call.args()
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
				todo(n)
			}
		}
	}

	n.Value = newRuntimeValue(ft.Result())
	return false
}

func (n *PrimaryExpression) checkSelector(ctx *context, t Type, nm xc.Token) []Selector {
	t0 := t
	ignoreMethods := false
	if t.Kind() == Ptr {
		ignoreMethods = t.isNamed()
		t = t.Elem()
		switch t.Kind() {
		case Interface:
			ctx.err(nm, "%s undefined (type *%s is pointer to interface, not interface)", nm.S(), t)
			return nil
		case Ptr:
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

func (n *PrimaryExpression) checkIndexExpr(ctx *context, o *ExpressionOpt) Value {
	if o == nil {
		return nil
	}

	todo(o)
	return nil
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
		n.Typ.check(ctx) ||
		n.TypeLiteral.check(ctx) {
		return true
	}

	if n.Expression != nil {
		n.lenPoisoned = n.Expression.lenPoisoned
	}
	if n.ExpressionOpt != nil {
		n.lenPoisoned = n.lenPoisoned || n.ExpressionOpt.lenPoisoned
	}
	if n.ExpressionOpt2 != nil {
		n.lenPoisoned = n.lenPoisoned || n.ExpressionOpt2.lenPoisoned
	}
	if n.ExpressionOpt3 != nil {
		n.lenPoisoned = n.lenPoisoned || n.ExpressionOpt3.lenPoisoned
	}
	switch n.Case {
	case 0: // Operand
		n.Value = n.Operand.Value
		n.lenPoisoned = n.Operand.lenPoisoned
	case 1: // CompLitType LBraceCompLitValue
		if n.CompLitType.check(ctx) {
			return true
		}

		if t := n.CompLitType.Type; t != nil {
			if (*CompLitValue)(nil).check(ctx, n.LBraceCompLitValue, t) {
				return true
			}

			n.Value = newRuntimeValue(t)
		}
	case 2: // PrimaryExpression '.' '(' "type" ')'
		todo(n)
	case 3: // PrimaryExpression '.' '(' Typ ')'
		v := n.PrimaryExpression.Value
		if v == nil {
			break
		}

		nm := n.Token2
		if nm.Val == idUnderscore {
			todo(n, true)
			break
		}

		switch v.Kind() {
		case RuntimeValue:
			t := v.Type()
			if t == nil {
				break
			}

			if t.Kind() != Interface {
				todo(n, true) // not an interface
				break
			}

			n.Value = newRuntimeValue(n.Typ.Type)
		default:
			//dbg("", v.Kind())
			todo(n)
		}
	case 4: // PrimaryExpression '.' IDENTIFIER
		v := n.PrimaryExpression.Value
		if v == nil {
			break
		}

		nm := n.Token2
		if nm.Val == idUnderscore {
			todo(n, true)
			break
		}

		switch v.Kind() {
		case PackageValue:
			if !isExported(nm.Val) { //TODO use d.IsExported()
				todo(n, true)
				break
			}

			p := v.Declaration().(*ImportDeclaration).Package
			if d := p.Scope.Bindings[nm.Val]; d != nil {
				switch x := d.(type) {
				case *ConstDeclaration:
					n.Value = x.Value
				case *FuncDeclaration:
					n.Value = newDeclarationValue(d, x.Type)
				case *TypeDeclaration:
					n.Value = newTypeValue(x)
				case *VarDeclaration:
					n.Value = newRuntimeValue(x.Type)
				default:
					//dbg("%T", d)
					todo(n)
				}
				break
			}

			ctx.err(nm, "undefined: %s.%s", p.Name, nm.S())
		case RuntimeValue:
			t := v.Type()
			if t == nil {
				break
			}

			path := n.checkSelector(ctx, t, nm)
			if path == nil {
				break
			}

			rv := v.(*runtimeValue)
			root := rv.root
			if root == nil {
				root = v
			}
			n.Value = newSelectorValue(path[len(path)-1].typ(), root, append(rv.path0, path[len(path)-1]), append(rv.path, path...))
		case TypeValue:
			t := v.Type()
			if t.Kind() == Interface {
				todo(n, true)
				break
			}

			m := t.MethodByName(nm.Val)
			if m == nil {
				todo(n)
				break
			}

			if t.isNamed() && t.PkgPath() != ctx.pkg.importPath && !isExported(nm.Val) {
				todo(n, true) // cannot use unexported method
				break
			}

			n.Value = newRuntimeValue(m.Type)
		default:
			//dbg("", v.Kind())
			todo(n)
		}
	case 5: // PrimaryExpression '[' Expression ']'
		a := n.PrimaryExpression.Value
		if a == nil {
			break
		}

		pt := a.Type()
		if pt == nil {
			break
		}

		// For a of pointer to array type:
		//
		// · a[x] is shorthand for (*a)[x]
		if pt.Kind() == Ptr && pt.Elem().Kind() == Array {
			a = newRuntimeValue(pt.Elem())
		}

		x := n.Expression.Value
		if x == nil {
			break
		}

		var constX Const
		// If a is not a map:
		//
		// · the index x must be of integer type or untyped; it is in
		//   range if 0 <= x < len(a), otherwise it is out of range
		//
		// · a constant index must be non-negative and representable by
		//   a value of type int
		if pt.Kind() != Map {
			if x.Kind() == RuntimeValue && !x.Type().IntegerType() {
				switch pt.Kind() {
				case Array:
					if ctx.err(n.Expression, "non-integer array index") {
						return true
					}
				case Slice:
					if ctx.err(n.Expression, "non-integer slice index") {
						return true
					}
				case String:
					if ctx.err(n.Expression, "non-integer string index") {
						return true
					}
				}

				break
			}

			if x.Kind() == ConstValue {
				constX = ctx.mustConvertConst(n.Expression, ctx.intType, x.Const())
				if constX == nil {
					break
				}

				if !constX.nonNegativeInteger() {
					todo(n.Expression, true) // < 0
					break
				}
			}
		}

		switch pt.Kind() {
		case Array:
			// For a of array type A:
			//
			// · a constant index must be in range
			//
			// · if x is out of range at run time, a run-time panic
			//   occurs
			//
			// · a[x] is the array element at index x and the type
			//   of a[x] is the element type of A
			if i, j := constX.(*intConst).val, pt.Len(); i >= j {
				if ctx.err(n.Expression, "invalid array index %d (out of bounds for %d-element array)", i, j) {
					return true
				}

				break
			}

			fallthrough
		case Slice:
			n.Value = newRuntimeValue(pt.Elem())
		case String:
			todo(n)
		case Map:
			kt := pt.Key()
			if kt == nil {
				break
			}

			if !x.AssignableTo(kt) {
				todo(n, true) // invalid index type
				break
			}

			n.Value = newRuntimeValue(pt.Elem())
		default:
			todo(n)
		}
	case 6: // PrimaryExpression '[' ExpressionOpt ':' ExpressionOpt ':' ExpressionOpt ']'
		todo(n)
	case 7: // PrimaryExpression '[' ExpressionOpt ':' ExpressionOpt ']'
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

			pa := false
			if e := t.Elem(); e != nil && e.Kind() == Array {
				pa = true
				t = e
			}

			switch t.Kind() {
			case Array:
				if !pa && !v.Addressable() {
					if ctx.err(n.PrimaryExpression, "invalid operation on %s value (slice of unaddressable value)", t) {
						return true
					}
					break
				}

				n.Value = newRuntimeValue(newSliceType(ctx, t.Elem()))
			default:
				//dbg("", t.Kind())
				todo(n, true) // cannot slice T
			}
		default:
			//dbg("", v.Kind())
			todo(n)
		}
		l := n.checkIndexExpr(ctx, n.ExpressionOpt)
		if l != nil {
			todo(n) // check L
		}
		h := n.checkIndexExpr(ctx, n.ExpressionOpt2)
		if h != nil {
			todo(n) // check H
		}
		if l != nil && h != nil {
			todo(n) // check l <= H
		}
	case 8: // PrimaryExpression Call
		defer func() {
			n.lenPoisoned = n.Call.lenPoisoned
			if n.Value != nil && n.Value.Kind() != ConstValue {
				n.lenPoisoned = true
			}
		}()

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

			if d := v.Declaration(); d != nil {
				switch {
				case ctx.isPredeclared(d):
					switch d.Name() {
					case idAppend:
						todo(n)
					case idCap:
						n.Value = builtinCap(ctx, n.Call)
					case idClose:
						todo(n)
					case idComplex:
						n.Value = builtinComplex(ctx, n.Call)
					case idCopy:
						todo(n)
					case idDelete:
						todo(n)
					case idImag:
						todo(n)
					case idLen:
						n.Value = builtinLen(ctx, n.Call)
					case idMake:
						n.Value = builtinMake(ctx, n.Call)
					case idNew:
						n.Value = builtinNew(ctx, n.Call)
					case idPanic:
						todo(n)
					case idPrint:
						todo(n)
					case idPrintln:
						todo(n)
					case idReal:
						n.Value = builtinReal(ctx, n.Call)
					case idRecover:
						todo(n)
					default:
						panic("internal error")
					}
					return false
				default:
					if x, ok := d.(*FuncDeclaration); ok && x.unsafe {
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
						return false
					}
				}
			}

			return n.checkCall(ctx, t, 0)
		case TypeValue: // Conversion.
			t := v.Type()
			args, _, ddd := n.Call.args()
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
			if (*CompLitValue)(nil).check(ctx, n.CompLitValue, v.Type()) {
				return true
			}

			n.Value = n.CompLitValue.Value()
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

func (n *Prologue) post(lx *lexer) (stop bool) {
	for _, v := range lx.imports {
		v.once = lx.oncePackage(v.bl, string(dict.S(v.ip)))
	}
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

		if v.name == 0 {
			return true
		}

		lx.fileScope.declare(lx, v)
	}
	return false
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

func (n *QualifiedIdent) checkTypeDeclaration(ctx *context, t *TypeDeclaration) (stop bool) {
	if n.Case == 0 /* IDENTIFIER */ && ctx.isPredeclared(t) && ctx.pkg.ImportPath != "" {
		if t.Kind() == Invalid {
			//dbg("", position(n.Pos()), ctx.pkg.ImportPath)
			panic("internal error")
		}

		return false

	}

	return t.check(ctx)
}

func (n *QualifiedIdent) checkTypeName(ctx *context) (_ Type, stop bool) {
	if t := n.resolutionScope.mustLookupType(ctx, n, n.fileScope); t != nil {
		return t, n.checkTypeDeclaration(ctx, t)
	}

	return nil, false
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
	if n == nil {
		return false
	}

	return n.Block.check(ctx) ||
		n.ConstDecl.check(ctx) ||
		n.StatementNonDecl.check(ctx) ||
		n.TypeDecl.check(ctx) ||
		n.VarDecl.check(ctx)
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

// ----------------------------------------------------------- StatementNonDecl

func (n *StatementNonDecl) check(ctx *context) (stop bool) {
	if n == nil {
		return false
	}

	switch n.Case {
	case 0: // "break" IdentifierOpt
		todo(n)
	case 1: // "continue" IdentifierOpt
		todo(n)
	case 2: // "defer" Expression
		todo(n)
	case 3: // "fallthrough"
		todo(n)
	case 4: // ForStatement
		todo(n)
	case 5: // "go" Expression
		todo(n)
	case 6: // "goto" IDENTIFIER
		todo(n)
	case 7: // IDENTIFIER ':' Statement
		todo(n)
	case 8: // IfStatement
		todo(n)
	case 9: // "return" ExpressionListOpt
		if o := n.ExpressionListOpt; o != nil {
			if o.check(ctx) {
				return true
			}

			s := n.resolutionScope
			for !s.isFnScope {
				s = s.Parent
			}

			ft := *s.fnType
			if ft == nil {
				break
			}

			rt := ft.Result()
			if rt == nil {
				break
			}

			list := o.ExpressionList.list
			switch rt.Kind() {
			case Tuple:
				todo(n)
			default:
				if len(list) > 1 {
					todo(n, true) // too many values to return
				}
				v := list[0].Value
				if v == nil {
					break
				}

				if !v.AssignableTo(rt) {
					todo(n, true) // type mismatch
				}
			}
		}
	case 10: // SelectStatement
		todo(n)
	case 11: // SimpleStatement
		todo(n)
	case 12: // SwitchStatement
		todo(n)
	default:
		panic("internal error")
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
		lx.declarationScope.declare(lx, newFieldDeclaration(
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
			lx.declarationScope.declare(lx, newFieldDeclaration(
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
		lx.declarationScope.declare(lx, newFieldDeclaration(
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
		n.Type, stop = n.QualifiedIdent.checkTypeName(ctx)
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

// ------------------------------------------------------------------- TypeDecl

func (n *TypeDecl) check(ctx *context) (stop bool) {
	if n == nil {
		return false
	}

	todo(n)
	return false
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
		if n.ArrayType.check(ctx) {
			return true
		}

		n.Type = n.ArrayType.Type
	case 2: // ChanType
		todo(n)
	case 3: // FuncType
		todo(n)
	case 4: // InterfaceType
		todo(n)
	case 5: // MapType
		stop = n.MapType.check(ctx)
		n.Type = n.MapType.Type
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
	switch n.Case {
	case 7: // PrimaryExpression
		n.lenPoisoned = n.PrimaryExpression.lenPoisoned
	default:
		n.lenPoisoned = n.UnaryExpression.lenPoisoned
		v = n.UnaryExpression.Value
		if v == nil {
			return false
		}
	}

	switch n.Case {
	case 0: // '!' UnaryExpression
		switch v.Kind() {
		case ConstValue:
			switch c := v.Const(); c.Type().Kind() {
			case Bool:
				n.Value = newConstValue(newBoolConst(!c.(*boolConst).val, c.Type(), false))
			case UntypedBool:
				n.Value = newConstValue(newBoolConst(!c.(*boolConst).val, c.Type(), true))
			default:
				todo(n, true) // inv op
			}
		default:
			//dbg("", v.Kind())
			todo(n)
		}
	case 1: // '&' UnaryExpression
		switch v.Kind() {
		case RuntimeValue:
			//TODO disallow taking address of a function (but not of a variable of a function type).
			n.Value = newRuntimeValue(newPtrType(ctx, v.Type()))
		default:
			//dbg("", v.Kind())
			todo(n)
		}
	case 2: // '*' UnaryExpression
		switch v.Kind() {
		case RuntimeValue:
			t := v.Type()
			if t.Kind() != Ptr {
				todo(n, true)
				break
			}

			n.Value = newRuntimeValue(t.Elem())
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
			n.Value = v.neg(ctx.Context, n.UnaryExpression)
		default:
			//dbg("", v.Kind())
			todo(n)
		}
	case 5: // '^' UnaryExpression
		switch v.Kind() {
		case ConstValue:
			n.Value = v.cpl(ctx.Context, n.UnaryExpression)
		default:
			//dbg("", v.Kind())
			todo(n)
		}
	case 6: // "<-" UnaryExpression
		n.lenPoisoned = true
		switch v.Kind() {
		case RuntimeValue:
			t := v.Type()
			if t == nil {
				break
			}

			n.Value = newRuntimeValue(t.Elem())
		default:
			//dbg("", v.Kind())
			todo(n)
		}
	case 7: // PrimaryExpression
		n.Value = n.PrimaryExpression.Value
	default:
		panic("internal error")
	}
	return false
}

// -------------------------------------------------------------------- VarDecl

func (n *VarDecl) check(ctx *context) (stop bool) {
	if n == nil {
		return false
	}

	todo(n)
	return false
}
