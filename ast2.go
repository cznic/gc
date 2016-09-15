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
		//dbg("", position(n.Pos()))
		if n.Expression.check(ctx) {
			return true
		}

		n.Value = n.Expression.Value
		//dbg("%p", n.Value)
		n.flags = n.Expression.flags
	case 1: // TypeLiteral
		if n.TypeLiteral.check(ctx) {
			return true
		}

		if t := n.TypeLiteral.Type; t != nil {
			n.Value = newTypeValue(t)
			n.flags = n.TypeLiteral.Type.flags()
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

func (n *ArgumentList) check(ctx *context) (stop bool) {
	for ; n != nil; n = n.ArgumentList {
		if n.Argument.check(ctx) {
			return true
		}

		n.flags = n.flags | n.Argument.flags
	}
	return false
}

// ------------------------------------------------------------------ ArrayType

func (n *ArrayType) check(ctx *context) (stop bool) {
	if n == nil {
		return false
	}

	done, stop := n.guard.check(ctx, nil, n.Type)
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

		n.Type = newArrayType(t, -1, 0)
		if cv := n.compLitValue; cv != nil {
			return (*CompLitValue)(nil).check(ctx, cv, n.Type)
		}
	case 1: // '[' Expression ']' Typ
		ctx = ctx.setErrf(func(g *gate) bool {
			return ctx.err(n.Expression, "invalid array bound")
		})
		if n.Expression.check(ctx) || n.Typ.check(ctx) {
			return true
		}

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
			case FloatingPointConst:
				if !c.nonNegativeInteger(ctx) {
					return ctx.constAssignmentFail(n.Expression, ctx.intType, c)
				}

				if d := c.Convert(ctx.Context, ctx.intType); d != nil {
					val := d.Const().(*intConst).val
					if val >= 0 {
						n.Type = newArrayType(t, val, n.Expression.flags)
						break
					}

					todo(n, true)
					break
				}

				return ctx.err(n.Expression, "array bound is too large")
			case IntConst:
				if d := c.Convert(ctx.Context, ctx.intType); d != nil {
					val := d.Const().(*intConst).val
					if val >= 0 {
						n.Type = newArrayType(t, val, n.Expression.flags)
						break
					}

					return ctx.err(n.Expression, "array bound must be non-negative")
				}

				return ctx.err(n.Expression, "array bound is too large")
			default:
				return ctx.err(n.Expression, "invalid array bound %s", c)
			}
		case NilValue:
			return ctx.err(n.Expression, "invalid array bound nil")
		default:
			todo(n, true)
		}
	default:
		panic("internal error")
	}
	return false
}

// ----------------------------------------------------------------- Assignment

func (n *Assignment) check0(ctx *context, f func([]*Expression, []*Expression) bool, checkChanRx, checkIndex bool) (stop bool) {
	lhs := n.ExpressionList.list
	rhs := n.ExpressionList2.list
	switch len(rhs) {
	case 1:
		e := rhs[0]
		if e == nil {
			return
		}
		rv := e.Value
		if rv == nil {
			return
		}

		rt := rv.Type()
		if rt == nil {
			return
		}

		if len(lhs) == 2 {
			switch {
			case checkChanRx && e.isChanReceive():
				rt = newTupleType([]Type{ctx.untypedBoolType, rt})
				rv = newRuntimeValue(rt)
			case checkIndex && e.isIndex():
				t := e.UnaryExpression.PrimaryExpression.PrimaryExpression.Value.Type()
				if t == nil {
					return false
				}

				switch t.Kind() {
				case Map:
					rt = newTupleType([]Type{rt, ctx.untypedBoolType})
					rv = newRuntimeValue(rt)
				default:
					todo(n)
				}
			}
		}

		switch rt.Kind() {
		case Tuple:
			todo(n)
		default:
			if len(lhs) != 1 {
				return ctx.err(e, "assignment count mismatch: %d = 1", len(lhs))
			}

			e := lhs[0]
			if e == nil {
				return
			}

			lv := e.Value
			if lv == nil {
				return
			}

			if !lv.assignableTo() {
				return ctx.err(e, "invalid assignment destination")
			}

			lt := lv.Type()
			if lt == nil {
				return
			}

			if !rv.AssignableTo(ctx.Context, lt) {
				if ctx.err(n, "cannot use type %s as type %s in assignment", rv.Type(), lt) {
					return true
				}
			}
		}
	default:
		todo(n)
	}
	if f != nil {
		return f(lhs, rhs)
	}

	return false
}

func (n *Assignment) check(ctx *context) (stop bool) {
	if n == nil {
		return false
	}

	if n.ExpressionList.check(ctx, true) ||
		n.ExpressionList2.check(ctx, false) {
		return true
	}

	switch n.Case {
	case 0: // ExpressionList '=' ExpressionList
		n.check0(ctx, nil, true, true)
	case 1: // ExpressionList "+=" ExpressionList
		todo(n)
	case 2: // ExpressionList "&^=" ExpressionList
		todo(n)
	case 3: // ExpressionList "&=" ExpressionList
		todo(n)
	case 4: // ExpressionList "/=" ExpressionList
		todo(n)
	case 5: // ExpressionList "<<=" ExpressionList
		todo(n)
	case 6: // ExpressionList "%=" ExpressionList
		todo(n)
	case 7: // ExpressionList "*=" ExpressionList
		todo(n)
	case 8: // ExpressionList "|=" ExpressionList
		n.check0(ctx, func(lhs, rhs []*Expression) bool {
			for _, item := range rhs {
				if item == nil {
					continue
				}

				v := item.Value
				if v == nil {
					continue
				}

				if !v.Type().IntegerType() {
					if ctx.err(item, "invalid operation: |= (operator | not defined on %s)", v.Type()) {
						return true
					}
				}
			}
			return false
		}, false, false)
	case 9: // ExpressionList ">>=" ExpressionList
		todo(n)
	case 10: // ExpressionList "-=" ExpressionList
		todo(n)
	case 11: // ExpressionList "^=" ExpressionList
		todo(n)
	default:
		panic("internal error")
	}
	return false
}

// ---------------------------------------------------------------------- Block

func (n *Block) check(ctx *context) (stop bool) {
	if n == nil {
		return false
	}

	return n.StatementList.check(ctx)
}

// ----------------------------------------------------------------------- Body

func (n *Body) check(ctx *context) (stop bool) {
	if n == nil {
		return false
	}

	return n.StatementList.check(ctx)
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

		n.flags = n.flags | l.Argument.flags
	}
	return false
}

func (n *Call) args() (_ []Value, flags flags, ddd bool) {
	var a []Value
	for l := n.ArgumentList; l != nil; l = l.ArgumentList {
		if l.Argument.Expression != nil {
			flags = flags | l.Argument.Expression.flags
		}
		a = append(a, l.Argument.Value)
	}
	return a, flags, n.Case == 2 // '(' ArgumentList "..." CommaOpt ')'
}

// ------------------------------------------------------------------- ChanType

func (n *ChanType) check(ctx *context) (stop bool) {
	if n == nil {
		return false
	}

	done, stop := n.guard.check(ctx, nil, n.Type)
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

	done, stop := n.guard.check(ctx, nil, n.Type)
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
	done, stop := n.gate().check(ctx, nil, t)
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
	tt := t
	if t.Kind() == Ptr {
		tt = t.Elem()
	}
	if tt == nil {
		return
	}

	tk := tt.Kind()

	switch tk {
	case Array:
		len = tt.Len()
		et = tt.Elem()
		if len < 0 {
			defer func() {
				tt.(*arrayType).len = maxIndex + 1
			}()
		}
	case Map:
		kt = tt.Key()
		et = tt.Elem()
	case Slice:
		et = tt.Elem()
		len = -1
	case Struct:
		numField = tt.NumField()
	default:
		ctx.err(n, "invalid type for composite literal: %s", tt)
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

					if v.Kind() != ConstValue || !v.nonNegativeInteger(ctx) {
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

					if !v.AssignableTo(ctx.Context, kt) {
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
						case IntConst, RuneConst:
							if d := ck.Convert(ctx.Context, ctx.intType); d != nil {
								k.i = d.Const().(*intConst).val
								break
							}

							todo(n, true) // too large
						case StringConst:
							k.i = int64(ck.(*stringConst).val.id())
						default:
							//dbg("", ck.Kind())
							todo(e)
						}
						switch ck.Kind() {
						case IntConst, RuneConst, StringConst: //TODO more types
							if _, ok := m[k]; ok {
								if ctx.err(e, "duplicate key %s in map literal", ck) {
									return true
								}

								break
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

					field = tt.FieldByName(fn.Val)
					if field == nil {
						if ctx.err(e, "unknown %s field '%s' in struct literal", tt, fn.S()) {
							return true
						}

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
				if v != nil && !v.AssignableTo(ctx.Context, et) {
					//dbg("", position(val.Pos()))
					if ctx.compositeLiteralValueFail(val, v, et) {
						return true
					}
				}
				maxIndex = mathutil.MaxInt64(maxIndex, index)
				index++
			case Map:
				if v != nil && !v.AssignableTo(ctx.Context, et) {
					todo(val, true)
				}
			case Struct:
				switch {
				case hasKey:
					keyed = true
					if v != nil && field.Type != nil && !v.AssignableTo(ctx.Context, field.Type) {
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

					f := tt.Field(int(index))
					if v != nil && f.Type != nil && !v.AssignableTo(ctx.Context, f.Type) {
						if ctx.err(val, "cannot use type %s as type %s in field value", v.Type(), f.Type) {
							return true
						}
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

	if el == nil && t != nil {
		todo(n, true) // invalid
		return
	}

	if el != nil {
		lx.constExpr = el
		lx.constType = t
	}
	el0 := el
	shared := false
	if el0 == nil {
		shared = true
		el0 = lx.constExpr
		t = lx.constType
	}
	if el0 == nil && lx.firstConstSpec {
		lx.err(n, "constant declaration must have expression")
		return
	}

	scopeStart := lx.lookahead.Pos()
	el = el0
loop:
	for l := n.IdentifierList; l != nil; l = l.IdentifierList {
		var e *Expression
		switch {
		case el0 != nil:
			el = el0
			e = el0.Expression
			el0 = el0.ExpressionList
		default:
			lx.err(el, "not enough expression(s) in list")
			break loop
		}
		d := newConstDeclaration(l.ident(), t, e, lx.iota, scopeStart, shared)
		lx.declarationScope.declare(lx, d)
	}
	if el0 != nil {
		lx.err(el0, "extra expression(s) in list")
	}
}

// ----------------------------------------------------------------------- Elif

func (n *Elif) check(ctx *context) (stop bool) {
	if n == nil {
		return false
	}

	return n.IfHeader.check(ctx) || n.Body.check(ctx)
}

// ------------------------------------------------------------------- ElifList

func (n *ElifList) check(ctx *context) (stop bool) {
	for ; n != nil; n = n.ElifList {
		if n.Elif.check(ctx) {
			return true
		}
	}

	return false
}

// -------------------------------------------------------------------- ElseOpt

func (n *ElseOpt) check(ctx *context) (stop bool) {
	if n == nil {
		return false
	}

	return n.Block.check(ctx)
}

// ----------------------------------------------------------------- Expression

func (n *Expression) isCall() bool {
	if n.Case != 0 {
		return false
	}

	return n.UnaryExpression.isCall()
}

func (n *Expression) isIndex() bool { return n.Case == 0 && n.UnaryExpression.isIndex() }

func (n *Expression) isChanReceive() bool {
	if n.Case != 0 {
		return false
	}

	return n.UnaryExpression.isChanReceive()
}

func (n *Expression) isIdentifier() (xc.Token, bool) { //TODO remove, use .ident()
	if n.Case != 0 {
		return xc.Token{}, false
	}

	return n.UnaryExpression.isIdentifier()
}

func (n *Expression) check(ctx *context) (stop bool) {
	if n == nil {
		return false
	}

	//dbg("", position(n.Pos()))
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
		n.flags = n.UnaryExpression.flags
	default:
		n.flags = n.Expression.flags | n.Expression2.flags
	}

	switch n.Case {
	case 0: // UnaryExpression
		n.Value = n.UnaryExpression.Value
		//dbg("%p", n.Value)
	case 1: // Expression '%' Expression
		n.Value = n.Expression.Value.mod(ctx, n.Token, n.Expression2.Value)
	case 2: // Expression '&' Expression
		n.Value = n.Expression.Value.and(ctx, n.Token, n.Expression2.Value)
	case 3: // Expression '*' Expression
		n.Value = n.Expression.Value.mul(ctx, n.Token, n.Expression2.Value)
	case 4: // Expression '+' Expression
		n.Value = n.Expression.Value.add(ctx, n.Token, n.Expression2.Value)
	case 5: // Expression '-' Expression
		n.Value = n.Expression.Value.sub(ctx, n.Token, n.Expression2.Value)
	case 6: // Expression '/' Expression
		n.Value = n.Expression.Value.div(ctx, n.Token, n.Expression2.Value)
	case 7: // Expression '<' Expression
		n.Value = n.Expression.Value.lt(ctx, n.Token, n.Expression2.Value)
	case 8: // Expression '>' Expression
		n.Value = n.Expression.Value.gt(ctx, n.Token, n.Expression2.Value)
	case 9: // Expression '^' Expression
		n.Value = n.Expression.Value.xor(ctx, n.Token, n.Expression2.Value)
	case 10: // Expression '|' Expression
		n.Value = n.Expression.Value.or(ctx, n.Token, n.Expression2.Value)
	case 11: // Expression "&&" Expression
		n.Value = n.Expression.Value.boolAnd(ctx, n.Token, n.Expression2.Value)
	case 12: // Expression "&^" Expression
		n.Value = n.Expression.Value.andNot(ctx, n.Token, n.Expression2.Value)
	case 13: // Expression "==" Expression
		n.Value = n.Expression.Value.eq(ctx, n.Token, n.Expression2.Value)
	case 14: // Expression ">=" Expression
		n.Value = n.Expression.Value.ge(ctx, n.Token, n.Expression2.Value)
	case 15: // Expression "<=" Expression
		n.Value = n.Expression.Value.le(ctx, n.Token, n.Expression2.Value)
	case 16: // Expression "<<" Expression
		n.Value = n.Expression.Value.lsh(ctx, n.Token, n.Expression2.Value)
	case 17: // Expression "!=" Expression
		n.Value = n.Expression.Value.neq(ctx, n.Token, n.Expression2.Value)
	case 18: // Expression "||" Expression
		n.Value = n.Expression.Value.boolOr(ctx, n.Token, n.Expression2.Value)
	case 19: // Expression ">>" Expression
		n.Value = n.Expression.Value.rsh(ctx, n.Token, n.Expression2.Value)
	case 20: // Expression "<-" Expression
		n.Value = newRuntimeValue(ctx.voidType)
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

func (n *ExpressionList) check(ctx *context, ignoreBlank bool) (stop bool) {
	if n == nil {
		return false
	}

	if len(n.list) != 0 {
		todo(n)
		return false
		//TODO panic("internal error")
	}

	n0 := n
	for ; n != nil; n = n.ExpressionList {
		t := n.Expression.ident()
		if !ignoreBlank || t.Val != idUnderscore {
			if n.Expression.check(ctx) {
				return true
			}
		}

		n0.list = append(n0.list, n.Expression)
	}
	return false
}

// ---------------------------------------------------------- ExpressionListOpt

func (n *ExpressionListOpt) check(ctx *context) (stop bool) {
	if n == nil {
		return false
	}

	return n.ExpressionList.check(ctx, false)
}

//--------------------------------------------------------------- ExpressionOpt

func (n *ExpressionOpt) check(ctx *context) (stop bool) {
	if n == nil {
		return false
	}

	if n.Expression.check(ctx) {
		return true
	}

	n.flags = n.Expression.flags
	n.Value = n.Expression.Value
	return false
}

// ------------------------------------------------------------------ ForHeader

func (n *ForHeader) check(ctx *context) (stop bool) {
	if n == nil {
		return false
	}

	switch n.Case {
	case 0: // Range
		if n.Range.check(ctx) {
			return true
		}
	case 1: // SimpleStatementOpt ';' SimpleStatementOpt ';' SimpleStatementOpt
		if n.SimpleStatementOpt.check(ctx, false) ||
			n.SimpleStatementOpt2.check(ctx, true) ||
			n.SimpleStatementOpt3.check(ctx, false) {
			return true
		}

		if o := n.SimpleStatementOpt2; !o.boolExprOrNil() {
			e := o.SimpleStatement.Expression
			if e == nil {
				todo(o, true)
				break
			}

			v := e.Value
			if v == nil {
				todo(o, true)
				break
			}

			if ctx.err(e, "non-bool (type %s) used as for condition", v.Type()) {
				return true
			}

			break
		}
	case 2: // SimpleStatementOpt
		if n.SimpleStatementOpt.check(ctx, true) {
			return true
		}

		if !n.SimpleStatementOpt.boolExprOrNil() {
			todo(n, true) // need bool
			break
		}
	default:
		panic("internal error")
	}
	return false
}

// --------------------------------------------------------------- ForStatement

func (n *ForStatement) check(ctx *context) (stop bool) {
	if n == nil {
		return false
	}

	return n.ForHeader.check(ctx) || n.Body.check(ctx)
}

// ------------------------------------------------------------------- FuncType

func (n *FuncType) check(ctx *context) (stop bool) {
	if n == nil {
		return false
	}

	done, stop := n.guard.check(ctx, nil, n.Type)
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

// ------------------------------------------------------------------- IfHeader

func (n *IfHeader) check(ctx *context) (stop bool) {
	if n == nil {
		return false
	}

	switch n.Case {
	case 0: // SimpleStatementOpt
		return n.SimpleStatementOpt.check(ctx, true)
	case 1: // SimpleStatementOpt ';' SimpleStatementOpt
		return n.SimpleStatementOpt.check(ctx, false) || n.SimpleStatementOpt2.check(ctx, true)
	default:
		panic("internal error")
	}
}

// ---------------------------------------------------------------- IfStatement

func (n *IfStatement) check(ctx *context) (stop bool) {
	if n == nil {
		return false
	}

	if n.IfHeader.check(ctx) ||
		n.Body.check(ctx) ||
		n.ElifList.check(ctx) ||
		n.ElseOpt.check(ctx) {
		return true
	}
	return false
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

	done, stop := n.guard.check(ctx, nil, n.Type)
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
					ctx = ctx.setErrf(func(g *gate) bool {
						*g = gateCycle
						return ctx.err(qi, "interface type loop involving %s", qi.str())
					})

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

	done, stop := n.guard.check(ctx, nil, n.Type)
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

	//dbg("", position(n.Pos()), n.Case)
	if n.Expression.check(ctx) ||
		n.FuncType.check(ctx) ||
		n.StatementList.check(ctx) {
		return true
	}

	switch n.Case {
	case 0: // '(' Expression ')'
		n.Value = n.Expression.Value
		n.flags = n.Expression.flags
	case 1: // '(' TypeLiteral ')'
		if n.TypeLiteral.check(ctx) {
			return true
		}

		if t := n.TypeLiteral.Type; t != nil {
			n.Value = newTypeValue(t)
			n.flags = t.flags()
		}
	case 2: // BasicLiteral
		n.Value = n.BasicLiteral.Value
	case 3: // FuncType LBrace StatementList '}'
		t := n.FuncType.Type
		n.Value = newRuntimeValue(t)
		if t != nil {
			n.flags = t.flags()
		}
		if n.StatementList != nil {
			n.flags = n.flags | n.StatementList.flags
		}
	case 4: // IDENTIFIER GenericArgumentsOpt
		if n.GenericArgumentsOpt != nil {
			//TODO generics todo(n)
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

		if d.check(ctx) {
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
					n.flags = flagHasIota
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
			//dbg("%p", n.Value)
		case *FuncDeclaration:
			n.Value = newDeclarationValue(d, x.Type)
			if x.Type != nil {
				n.flags = x.Type.flags()
			}
		case *ImportDeclaration:
			n.Value = newPackageValue(x)
		case *ParameterDeclaration:
			n.Value = newAddressableValue(x.Type)
			if x.Type != nil {
				n.flags = x.Type.flags()
			}
		case *TypeDeclaration:
			n.Value = newTypeValue(x)
			n.flags = x.flags()
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

			n.Value = newAddressableValue(x.Type)
			if x.Type != nil {
				n.flags = x.Type.flags()
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
					//TODO generics todo(n)
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

func (n *PrimaryExpression) checkConversion(ctx *context, node Node, t Type, arg Value) {
	if t == nil || arg == nil {
		return
	}

	switch arg.Kind() {
	case ConstValue:
		c := arg.Const()
		if n.Value = c.Convert(ctx.Context, t); n.Value == nil {
			ctx.constConversionFail(node, t, c)
		}
	case NilValue:
		if n.Value = arg.Convert(ctx.Context, t); n.Value == nil {
			ctx.mustConvertNil(node, t)
		}
	//TODO- case RuntimeValue:
	//TODO- 	if n.Value = arg.Convert(ctx.Context, t); n.Value == nil {
	//TODO- 		switch {
	//TODO- 		case t.Kind() == Interface:
	//TODO- 			arg.Type().implementsFailed(ctx, node, "cannot convert %s to type %s:", t)
	//TODO- 		default:
	//TODO- 			ctx.err(node, "1888: cannot convert %s to type %s", arg.Type(), t)
	//TODO- 		}
	//TODO- 	}
	case RuntimeValue:
		if arg.Type() != nil && !arg.Type().ConvertibleTo(t) {
			ctx.err(node, "cannot convert type %s to %s", arg.Type(), t)
			return
		}

		n.Value = newRuntimeValue(t)
	case TypeValue:
		todo(n)
	default:
		panic("internal error")
	}
}

func (n *PrimaryExpression) variadicParam(ft Type, i int) Type {
	if n := ft.NumIn(); ft.IsVariadic() && i >= n-1 {
		return ft.In(n - 1).Elem()
	}

	return ft.In(i)
}

func (n *PrimaryExpression) checkCall(ctx *context, ft Type, skip int) (stop bool) {
	if ft == nil {
		return false
	}

	const (
		variadic = 1 << iota
		dots
		tuple
	)

	args, _, ddd := n.Call.args()
	var flags int
	if ft.IsVariadic() {
		flags |= variadic
	}
	if ddd {
		flags |= dots
	}
	if len(args) == 1 {
		arg := args[0]
		if arg != nil && arg.Kind() != NilValue {
			if t := arg.Type(); t != nil && t.Kind() == Tuple {
				flags |= tuple
			}
		}
	}

	in := ft.NumIn()
	switch flags {
	case variadic:
		if len(args) < in-1-skip {
			todo(n, true)
			break
		}

		for i, arg := range args {
			if arg != nil && !arg.AssignableTo(ctx.Context, n.variadicParam(ft, i+skip)) {
				ctx.valueAssignmentFail(n.Call.ArgumentList.node(i), n.variadicParam(ft, i+skip), arg)
			}
		}
	case dots:
		todo(n)
	case dots | variadic:
		todo(n)
	case tuple:
		args := args[0].Type().Elements()
		if len(args) < in-1-skip {
			todo(n, true)
			break
		}

		for i, arg := range args {
			if arg != nil && !arg.AssignableTo(ft.In(i+skip)) {
				todo(n, true)
			}
		}
	case tuple | variadic:
		args := args[0].Type().Elements()
		if len(args) < in-1-skip {
			todo(n, true)
			break
		}

		for i, arg := range args {
			if arg != nil && !arg.AssignableTo(n.variadicParam(ft, i+skip)) {
				todo(n, true)
			}
		}
	case tuple | dots:
		fallthrough
	case tuple | dots | variadic:
		ctx.err(n.Call.ArgumentList.node(0), "multiple-value in single-value context")
	default:
	}
	//TODO- switch {
	//TODO- case ft.IsVariadic():
	//TODO- 	todo(n)
	//TODO- default:
	//TODO- 	if len(args) != ft.NumIn()-skip {
	//TODO- 		todo(n) //TODO special case single tuple arg.
	//TODO- 	}
	//TODO- 	if ddd {
	//TODO- 		todo(n)
	//TODO- 	}
	//TODO- 	for i, arg := range args {
	//TODO- 		if arg != nil && !arg.AssignableTo(ft.In(i+skip)) {
	//TODO- 			todo(n)
	//TODO- 		}
	//TODO- 	}
	//TODO- }
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
			ctx.err(nm, "selecting %s.%s requires explicit dereference", t0, nm.S())
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

		if len(path) >= 100 {
			panic("internal error") //TODO need inf. recursion protection.
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
			if !f.Anonymous {
				continue
			}

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

func (n *PrimaryExpression) checkSliceIndex(ctx *context, o *ExpressionOpt) Const {
	if o == nil {
		return nil
	}

	if o.check(ctx) {
		return nil
	}

	v := o.Value
	if v == nil {
		return nil
	}

	if !v.Type().Numeric() || !v.nonNegativeInteger(ctx) {
		todo(n, true) // non-integer index
		return nil
	}

	switch v.Kind() {
	case ConstValue:
	default:
		//dbg("", v.Kind())
		todo(n)
		return nil
	}

	if c := v.Const().Convert(ctx.Context, ctx.intType); c != nil {
		return c.Const()
	}

	todo(n, true) // const index overflow
	return nil
}

func (n *PrimaryExpression) check(ctx *context) (stop bool) {
	if n == nil {
		return false
	}

	//dbg("", position(n.Pos()), n.Case)
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
		n.flags = n.Expression.flags
	}
	if n.ExpressionOpt != nil {
		n.flags = n.flags | n.ExpressionOpt.flags
	}
	if n.ExpressionOpt2 != nil {
		n.flags = n.flags | n.ExpressionOpt2.flags
	}
	if n.ExpressionOpt3 != nil {
		n.flags = n.flags | n.ExpressionOpt3.flags
	}
	switch n.Case {
	case 0: // Operand
		n.Value = n.Operand.Value
		n.flags = n.Operand.flags
		//dbg("%p", n.Value)
	case 1: // CompLitType LBraceCompLitValue
		if n.CompLitType.check(ctx) {
			return true
		}

		if t := n.CompLitType.Type; t != nil {
			if (*CompLitValue)(nil).check(ctx, n.LBraceCompLitValue, t) {
				return true
			}

			n.Value = newRuntimeValue(t)
			n.flags = t.flags()
		}
	case 2: // PrimaryExpression '.' '(' "type" ')'
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

			n.flags = t.flags()
			if t.Kind() != Interface {
				todo(n, true) // not an interface
			}
		default:
			//dbg("", v.Kind())
			todo(n)
		}
		n.flags = n.flags | n.PrimaryExpression.flags
	case 3: // PrimaryExpression '.' '(' Typ ')'
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

			n.flags = t.flags()
			if t.Kind() != Interface {
				todo(n, true) // not an interface
				break
			}

			t = n.Typ.Type
			n.Value = newRuntimeValue(t)
			if t != nil {
				n.flags = n.flags | t.flags()
			}
		default:
			//dbg("", v.Kind())
			todo(n)
		}
	case 4: // PrimaryExpression '.' IDENTIFIER
		v := n.PrimaryExpression.Value
		if v == nil {
			break
		}

		n.flags = n.PrimaryExpression.flags
		nm := n.Token2
		if nm.Val == idUnderscore && ctx.err(nm, "cannot refer to blank field or method") {
			return true
		}

		switch v.Kind() {
		case PackageValue:
			p := v.Declaration().(*ImportDeclaration).Package
			if !isExported(nm.Val) { //TODO use d.IsExported()
				if ctx.err(nm, "cannot refer to unexported name %s.%s", p.Name, nm.S()) {
					return true
				}

				break
			}

			if d := p.Scope.Bindings[nm.Val]; d != nil {
				switch x := d.(type) {
				case *ConstDeclaration:
					n.Value = x.Value
					if n.Value != nil && n.Value.Type() != nil {
						n.flags = n.flags | n.Value.Type().flags()
					}
				case *FuncDeclaration:
					n.Value = newDeclarationValue(d, x.Type)
					if x.Type != nil {
						n.flags = n.flags | x.Type.flags()
					}
				case *TypeDeclaration:
					n.Value = newTypeValue(x)
					n.flags = n.flags | x.flags()
				case *VarDeclaration:
					n.Value = newRuntimeValue(x.Type)
					if x.Type != nil {
						n.flags = n.flags | x.Type.flags()
					}
				default:
					//dbg("%T", d)
					todo(n)
				}
				break
			}

			ctx.err(nm, "undefined: %s.%s", p.Name, nm.S())
		case ConstValue:
			v = newRuntimeValue(v.Type())
			fallthrough
		case RuntimeValue:
			t := v.Type()
			if t == nil {
				break
			}

			n.flags = n.flags | t.flags()
			path := n.checkSelector(ctx, t, nm)
			if path == nil {
				break
			}

			rv := v.(*runtimeValue)
			root := rv.root
			if root == nil {
				root = v
			}
			val := path[len(path)-1]
			vt := val.typ()
			if m, ok := val.(*Method); ok && t.Kind() != Interface {
				t2 := t
				if t.Kind() == Ptr {
					t2 = t2.Elem()
				}
				m = t2.MethodByName(m.Name)
				if m != nil { // Method value.
					f0 := m.Type.(*funcType)
					f := *f0
					f.in = f0.in[1:]
					vt = &f
					f.typeBase.typ = vt
				}
			}
			n.Value = newSelectorValue(vt, root, append(rv.path0, path[len(path)-1]), append(rv.path, path...))
			if n.Value != nil && n.Value.Type() != nil {
				n.flags = n.flags | n.Value.Type().flags()
			}
		case TypeValue:
			t := v.Type()
			n.flags = n.flags | t.flags()
			if t.Kind() == Interface {
				todo(n, true)
				break
			}

			needPtrRx := false
			m := t.MethodByName(nm.Val)
			if m == nil && t.Kind() == Ptr {
				if e := t.Elem(); e != nil && e.Kind() != Interface {
					needPtrRx = true
					m = e.MethodByName(nm.Val)
				}
			}
			if m == nil {
				if needPtrRx {
					t = newPtrType(ctx.model.PtrBytes, t)
				}
				switch {
				case t.Kind() == Ptr:
					if ctx.err(nm, "(%s).%s undefined (type %s has no method %s)", t, nm.S(), t, nm.S()) {
						return true
					}
				default:
					if ctx.err(nm, "%s.%s undefined (type %s has no method %s)", t, nm.S(), t, nm.S()) {
						return true
					}
				}

				break
			}

			if t.isNamed() && t.PkgPath() != ctx.pkg.importPath && !isExported(nm.Val) {
				todo(n, true) // cannot use unexported method
				break
			}

			// Method expression.
			mt := m.Type
			if mt != nil {
				n.flags = n.flags | mt.flags()
				if needPtrRx {
					rxt := mt.In(0)
					if rxt != nil && rxt.Kind() != Ptr {
						mt = mt.(*funcType).ptrMethod(ctx)
					}
				}
			}
			n.Value = newRuntimeValue(mt)
		default:
			//dbg("", v.Kind())
			todo(n)
		}
	case 5: // PrimaryExpression '[' Expression ']'
		pv := n.PrimaryExpression.Value
		if pv == nil {
			break
		}

		pt := pv.Type()
		if pt == nil {
			break
		}

		n.flags = n.flags | pt.flags()
		// For a of pointer to array type:
		//
		// · a[x] is shorthand for (*a)[x]
		if pt.Kind() == Ptr && pt.Elem().Kind() == Array {
			pt = pt.Elem()
			pv = newRuntimeValue(pt)
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

				return false
			}

			if x.Kind() == ConstValue {
				constX = ctx.mustConvertConst(n.Expression, ctx.intType, x.Const())
				if constX == nil {
					return false
				}

				if !constX.nonNegativeInteger(ctx) {
					todo(n.Expression, true) // < 0
					return false
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
			if constX != nil {
				if i, j := constX.(*intConst).val, pt.Len(); i >= j {
					if ctx.err(n.Expression, "invalid array index %d (out of bounds for %d-element array)", i, j) {
						return true
					}

					break
				}
			}

			fallthrough
		case Slice:
			// For a of slice type S:
			//
			// · if x is out of range at run time, a run-time panic
			//   occurs
			//
			// · a[x] is the slice element at index x and the type
			//   of a[x] is the element type of S
			n.Value = newAddressableValue(pt.Elem())
			n.flags = n.flags | pt.Elem().flags()
		case String:
			// For a of string type:
			//
			// · a constant index must be in range if the string a
			//   is also constant
			//
			// · if x is out of range at run time, a run-time panic
			//   occurs
			//
			// · a[x] is the non-constant byte value at index x and
			//   the type of a[x] is byte
			//
			// · a[x] may not be assigned to
			switch pv.Kind() {
			case RuntimeValue:
				n.Value = newRuntimeValue(ctx.uint8Type)
			default:
				//dbg("", pv.Kind())
				todo(n)
			}
		case Map:
			// For a of map type M:
			//
			// · x's type must be assignable to the key type of M
			//
			// · if the map contains an entry with key x, a[x] is
			//   the map value with key x and the type of a[x] is the
			//   value type of M
			//
			// · if the map is nil or does not contain such an
			//   entry, a[x] is the zero value for the value type of
			//   M
			kt := pt.Key()
			if kt == nil {
				break
			}

			n.flags = n.flags | kt.flags()
			if !x.AssignableTo(ctx.Context, kt) {
				if ctx.err(n.Expression, "cannot use type %s as type %s in map index", x.Type(), kt) {
					return true
				}

				break
			}

			n.Value = newMapIndexValue(pt.Elem())
			n.flags = n.flags | pt.Elem().flags()
		default:
			// Otherwise a[x] is illegal.
			if ctx.err(n, "invalid operation: (type %s does not support indexing)", pt) {
				return true
			}
		}
	case 6: // PrimaryExpression '[' ExpressionOpt ':' ExpressionOpt ':' ExpressionOpt ']'
		todo(n)
		//TODO n.flags =
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

			n.flags = n.flags | t.flags()
			pa := false
			if t.Kind() == Ptr {
				if e := t.Elem(); e != nil && e.Kind() == Array {
					pa = true
					t = e
				}
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
				n.flags = n.flags | t.Elem().flags()
			default:
				//dbg("", t.Kind())
				todo(n, true) // cannot slice T
			}
		default:
			//dbg("", v.Kind())
			todo(n)
		}
		l := n.checkSliceIndex(ctx, n.ExpressionOpt)
		h := n.checkSliceIndex(ctx, n.ExpressionOpt2)
		if l != nil && h != nil {
			if l.(*intConst).val > h.(*intConst).val {
				todo(n, true) // l >= h
			}
		}
	case 8: // PrimaryExpression Call
		defer func() {
			n.flags = n.flags | n.Call.flags
			if n.Value != nil && n.Value.Kind() != ConstValue {
				n.flags = n.flags | flagLenPoisoned
			}
		}()

		v := n.PrimaryExpression.Value
		if v == nil {
			break
		}

		switch v.Kind() {
		case ConstValue:
			if ctx.err(n, "cannot call non-function (type %s)", v.Type()) {
				return true
			}

			break
		case PackageValue:
			p := v.Declaration().(*ImportDeclaration).Package
			if ctx.err(n, "use of package %s without selector", p.Name) {
				return true
			}

			break
		case RuntimeValue:
			t := v.Type()
			if t == nil {
				break
			}

			n.flags = n.flags | t.flags()
			if t.Kind() != Func {
				todo(n, true)
				break
			}

			if d := v.Declaration(); d != nil {
				switch {
				case ctx.isPredeclared(d):
					switch d.Name() {
					case idAppend:
						n.Value = builtinAppend(ctx, n.Call)
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
						n.Value = builtinImag(ctx, n.Call)
					case idLen:
						n.Value = builtinLen(ctx, n.Call)
					case idMake:
						n.Value = builtinMake(ctx, n.Call)
					case idNew:
						n.Value = builtinNew(ctx, n.Call)
					case idPanic:
						n.Value = builtinPanic(ctx, n.Call)
					case idPrint:
						n.Value = builtinPrint(ctx, n.Call, false)
					case idPrintln:
						n.Value = builtinPrint(ctx, n.Call, true)
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
							n.Value = alignof(ctx, n.Call)
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
			n.flags = n.flags | t.flags()
			args, _, ddd := n.Call.args()
			if ddd {
				if ctx.err(n.Call.ArgumentList.Argument, "invalid use of ... in type conversion to %s", t) {
					return true
				}
			}
			switch len(args) {
			case 0:
				todo(n, true)
			case 1:
				n.checkConversion(ctx, n.Call.ArgumentList.Argument, t, args[0])
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

		if v.Type() != nil {
			n.flags = n.flags | v.Type().flags()
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
		t := n.TypeLiteral.Type
		n.checkConversion(ctx, n.Expression, t, n.Expression.Value)
		if t != nil {
			n.flags = n.flags | t.flags()
		}
	default:
		panic("internal error")
	}
	return false
}

func (n *PrimaryExpression) isCall() bool {
	return n.Case == 8 // PrimaryExpression Call
}

func (n *PrimaryExpression) isIndex() bool { return n.Case == 5 } // PrimaryExpression '[' Expression ']'

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

// ---------------------------------------------------------------------- Range

func (n *Range) check(ctx *context) (stop bool) {
	if n == nil {
		return false
	}

	todo(n)
	return false
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
				//TODO generics todo(n)
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
			n.Type = newTupleType(a)
		}
	case 2: // Typ
		n.Type = n.Typ.Type
	default:
		panic("internal error")
	}
	return false
}

// ------------------------------------------------------------ SelectStatement

func (n *SelectStatement) check(ctx *context) (stop bool) {
	if n == nil {
		return false
	}

	return n.SwitchBody.check(ctx)
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

// ------------------------------------------------------------ SimpleStatement

func (n *SimpleStatement) check(ctx *context, used bool) (stop bool) {
	if n == nil {
		return false
	}

	if n.Assignment.check(ctx) ||
		n.Expression.check(ctx) ||
		n.Case != 4 && (n.ExpressionList.check(ctx, false) ||
			n.ExpressionList2.check(ctx, false)) {
		return true
	}

	//dbg("", position(n.Pos()))
	op := "++"
	switch n.Case {
	case 0: // Assignment
		// nop
	case 1: // Expression
		v := n.Expression.Value
		if v == nil {
			break
		}

		if used || isVoid(v.Type()) || n.Expression.isCall() || n.Expression.isChanReceive() {
			break
		}

		switch v.Kind() {
		case ConstValue:
			if ctx.err(n.Expression, "%s evaluated but not used", v.Const()) {
				return true
			}
		case TypeValue:
			if ctx.err(n.Expression, "type %s is not an expression", v.(*typeValue).Type()) {
				return true
			}
		default:
			todo(n, true) // unused value
		}
	case 2: // Expression "--"
		op = "--"
		fallthrough
	case 3: // Expression "++"
		v := n.Expression.Value
		if v == nil {
			break
		}

		if !v.Type().Numeric() {
			if ctx.err(n.Expression, "invalid operation: %s (non-numeric type %s)", op, v.Type()) {
				return true
			}

			break
		}

		if !v.assignableTo() {
			todo(n, true) // must be addressable or map index
			break
		}
	case 4: // ExpressionList ":=" ExpressionList
		// nop
	default:
		panic("internal error")
	}
	return false
}

func (n *SimpleStatement) boolExprOrNil() bool {
	if n == nil {
		return true
	}

	switch n.Case {
	case
		0, // Assignment
		2, // Expression "--"
		3, // Expression "++"
		4: // ExpressionList ":=" ExpressionList
		return false
	case 1: // Expression
		v := n.Expression.Value
		return v == nil || v.Type().Kind() == Bool || v.Type().Kind() == UntypedBool
	default:
		panic("internal error")
	}
}

// --------------------------------------------------------- SimpleStatementOpt

func (n *SimpleStatementOpt) check(ctx *context, used bool) (stop bool) {
	if n == nil {
		return false
	}

	return n.SimpleStatement.check(ctx, used)
}

func (n *SimpleStatementOpt) boolExprOrNil() bool {
	if n == nil {
		return true
	}

	return n.SimpleStatement.boolExprOrNil()
}

// ------------------------------------------------------------------ SliceType

func (n *SliceType) check(ctx *context) (stop bool) {
	if n == nil {
		return false
	}

	stack := ctx.stack
	ctx.stack = nil

	defer func() { ctx.stack = stack }()

	done, stop := n.guard.check(ctx, nil, n.Type)
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

	//TODO set n.flags
}

// -------------------------------------------------------------- StatementList

func (n *StatementList) check(ctx *context) (stop bool) {
	for ; n != nil; n = n.StatementList {
		i := n.Statement
		if i.check(ctx) {
			return true
		}

		if n.Statement != nil {
			n.flags = n.flags | n.Statement.flags
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
		//TODO verify valid context
		if o := n.IdentifierOpt; o != nil {
			if n.resolutionScope.lookupLabel(ctx, o.Token) == nil && ctx.err(o.Token, "break label not defined: %s", o.Token.S()) {
				return true
			}

			//TODO verify valid target
		}
	case 1: // "continue" IdentifierOpt
		todo(n)
	case 2: // "defer" Expression
		todo(n)
	case 3: // "fallthrough"
		//TODO check valid context
	case 4: // ForStatement
		return n.ForStatement.check(ctx)
	case 5: // "go" Expression
		todo(n)
	case 6: // "goto" IDENTIFIER
		n.resolutionScope.mustLookupLabel(ctx, n.Token2)
		//TODO verify valid target
	case 7: // IDENTIFIER ':' Statement
		return n.Statement.check(ctx)
	case 8: // IfStatement
		return n.IfStatement.check(ctx)
	case 9: // "return" ExpressionListOpt
		o := n.ExpressionListOpt
		if o == nil {
			break
		}

		if o.check(ctx) {
			return true
		}

		s := n.resolutionScope
		for !s.isFnScope {
			s = s.Parent
		}

		if s.fnType == nil {
			break
		}

		ft := *s.fnType
		if ft == nil {
			break
		}

		rt := ft.Result()
		if rt == nil {
			break
		}

		if isVoid(rt) {
			todo(n, true)
			break
		}

		list := o.ExpressionList.list
		switch rt.Kind() {
		case Tuple:
			rte := rt.Elements()
			switch {
			case len(list) == 1: // must be tuple as well
				v := list[0].Value
				if v == nil {
					break
				}

				if v.Type().Kind() != Tuple {
					todo(n, true) // mismatch
					break
				}

				le := v.Type().Elements()
				if len(le) > len(rte) {
					if ctx.err(list[0], "too many arguments to return") {
						return true
					}

					break
				}

				if len(le) < len(rte) {
					todo(n, true) // not enough
					break
				}

				for i, v := range le {
					if !v.AssignableTo(rte[i]) {
						todo(n, true) // mismatch
					}
				}
			default:
				if len(list) < len(rte) {
					todo(n, true) // not enough values to return
				}
				if len(list) > len(rte) {
					todo(n, true) // too many values to return
				}
				for i, e := range list[:mathutil.Min(len(list), len(rte))] {
					if e == nil {
						continue
					}

					v := e.Value
					if v == nil {
						continue
					}

					t := rte[i]
					if t == nil {
						continue
					}

					if !v.AssignableTo(ctx.Context, rte[i]) {
						todo(n, true) // type mismatch
					}
				}
			}
		default: // Single valued function
			if len(list) > 1 {
				todo(n, true) // too many values to return
			}
			v := list[0].Value
			if v == nil {
				break
			}

			if !v.AssignableTo(ctx.Context, rt) {
				switch v.Kind() {
				case NilValue:
					todo(n, true) // type mismatch
				default:
					if ctx.err(n, "cannot use type %s as type %s in return argument", v.Type(), rt) {
						return true
					}
				}
			}
		}
	case 10: // SelectStatement
		return n.SelectStatement.check(ctx)
	case 11: // SimpleStatement
		return n.SimpleStatement.check(ctx, false)
	case 12: // SwitchStatement
		return n.SwitchStatement.check(ctx)
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

	done, stop := n.guard.check(ctx, nil, n.Type)
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
		n.Type = newStructType(n, nil, nil)
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
		var fixes []*Type
		m := map[int]struct{}{}
		for _, f := range a {
			if f.check(ctx) {
				return true
			}

			if ft := f.Type; ft != nil && f.isAnonymous {
				ptr := ft.Kind() == Ptr
				if ptr {
					ft = ft.Elem()
				}
				if ft != nil && ft.NumMethod() != 0 {
					for i := 0; i < ft.NumMethod(); i++ {
						mt := ft.Method(i)
						if _, ok := m[mt.Name]; ok {
							continue
						}

						m[mt.Name] = struct{}{}
						var mt2 Method
						switch {
						case ft.Kind() == Interface:
							if ptr {
								break
							}

							// Merge interface method, must synthesize new method w/ receiver.
							mt2 = *mt
							in := []Type{nil}
							for i := 0; i < mt.Type.NumIn(); i++ {
								in = append(in, mt.Type.In(i))
							}
							fixes = append(fixes, &in[0])
							mt2.Type = newFuncType(ctx, mt.Name, in, mt.Type.Result(), isExported(mt.Name), mt.Type.IsVariadic())
						default:
							mt2 = *mt
						}
						mt2.merged = true
						mt2.Index = len(mta)
						mta = append(mta, mt2)
					}
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
				t = newPtrType(ctx.model.PtrBytes, t)
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
		n.Type = newStructType(n, sf, mta)
		for i := range fixes {
			*fixes[i] = n.Type
		}
	default:
		panic("internal error")
	}

	return stop
}

// ----------------------------------------------------------------- SwitchCase

func (n *SwitchCase) check(ctx *context) (stop bool) {
	if n == nil {
		return false
	}

	switch n.Case {
	case 0: // "case" ArgumentList ':'
		return n.ArgumentList.check(ctx)
	case 1: // "case" ArgumentList '=' Expression ':'
		todo(n)
	case 2: // "case" ArgumentList ":=" Expression ':'
		todo(n)
	case 3: // "default" ':'
		//TODO check unique
	case 4: // "case" error
		todo(n)
	case 5: // "default" error
		todo(n)
	default:
		panic("internal error")
	}
	return false
}

// ------------------------------------------------------------ SwitchCaseBlock

func (n *SwitchCaseBlock) check(ctx *context) (stop bool) {
	if n == nil {
		return false
	}

	return n.SwitchCase.check(ctx) || n.StatementList.check(ctx)
}

// ----------------------------------------------------------------- SwitchBody

func (n *SwitchBody) check(ctx *context) (stop bool) {
	for l := n.SwitchCaseList; l != nil; l = l.SwitchCaseList {
		if l.SwitchCaseBlock.check(ctx) {
			return true
		}
	}

	return false
}

// --------------------------------------------------------------- SwitchHeader

func (n *SwitchHeader) check(ctx *context) (stop bool) {
	if n == nil {
		return false
	}

	switch n.Case {
	case 0: // SimpleStatementOpt
		return n.SimpleStatementOpt.check(ctx, true)
	case 1: // SimpleStatementOpt ';'
		todo(n)
	case 2: // SimpleStatementOpt ';' Expression
		todo(n)
	case 3: // SimpleStatementOpt ';' IDENTIFIER ":=" PrimaryExpression '.' '(' "type" ')'
		todo(n)
	default:
		panic("internal error")
	}
	return false
}

// ------------------------------------------------------------ SwitchStatement

func (n *SwitchStatement) check(ctx *context) (stop bool) {
	if n == nil {
		return false
	}

	return n.SwitchHeader.check(ctx) || n.SwitchBody.check(ctx)
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

	done, stop := n.guard.check(ctx, nil, n.Type)
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
		n.Type = newPtrType(ctx.model.PtrBytes, t2.Type)
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
			//TODO generics todo(n)
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

	done, stop := n.guard.check(ctx, nil, n.Type)
	if done || stop {
		return stop
	}

	defer n.guard.done(nil)

	switch n.Case {
	case 0: // '*' TypeLiteral
		if n.TypeLiteral.check(ctx) {
			return true
		}

		if t := n.TypeLiteral.Type; t != nil {
			n.Type = newPtrType(ctx.model.PtrBytes, t)
		}
	case 1: // ArrayType
		if n.ArrayType.check(ctx) {
			return true
		}

		n.Type = n.ArrayType.Type
	case 2: // ChanType
		if n.ChanType.check(ctx) {
			return true
		}

		n.Type = n.ChanType.Type
	case 3: // FuncType
		todo(n)
	case 4: // InterfaceType
		if n.InterfaceType.check(ctx) {
			return true
		}

		n.Type = n.InterfaceType.Type
	case 5: // MapType
		if n.MapType.check(ctx) {
			return true
		}

		n.Type = n.MapType.Type
	case 6: // SliceType
		if n.SliceType.check(ctx) {
			return true
		}

		n.Type = n.SliceType.Type
	case 7: // StructType
		if n.StructType.check(ctx) {
			return true
		}

		n.Type = n.StructType.Type
	default:
		panic("internal error")
	}
	return false
}

// ------------------------------------------------------------ UnaryExpression

func (n *UnaryExpression) isCall() bool {
	if n.Case != 7 { // PrimaryExpression
		return false
	}

	return n.PrimaryExpression.isCall()
}

func (n *UnaryExpression) isChanReceive() bool {
	return n.Case == 6 // "<-" UnaryExpression
}

func (n *UnaryExpression) isIndex() bool { return n.Case == 7 && n.PrimaryExpression.isIndex() } // PrimaryExpression

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

	//dbg("", position(n.Pos()), n.Case)
	if n.UnaryExpression.check(ctx) || n.PrimaryExpression.check(ctx) {
		return true
	}

	var v Value
	switch n.Case {
	case 7: // PrimaryExpression
		n.flags = n.PrimaryExpression.flags
	default:
		n.flags = n.UnaryExpression.flags
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
			n.Value = newRuntimeValue(newPtrType(ctx.model.PtrBytes, v.Type()))
		default:
			//dbg("", v.Kind())
			todo(n)
		}
	case 2: // '*' UnaryExpression
		switch v.Kind() {
		case RuntimeValue:
			t := v.Type()
			if t.Kind() != Ptr {
				if ctx.err(n.UnaryExpression, "invalid indirect of type %s", t) {
					return true
				}

				break
			}

			n.Value = newAddressableValue(t.Elem())
		case TypeValue:
			n.Value = newTypeValue(newPtrType(ctx.model.PtrBytes, v.Type()))
		default:
			//dbg("", v.Kind())
			todo(n)
		}
	case 3: // '+' UnaryExpression
		switch v.Kind() {
		case ConstValue:
			if !v.Type().Numeric() {
				s := ""
				if v.Const().Untyped() {
					s = "untyped "
				}
				if ctx.err(n.UnaryExpression, "invalid operation: + %s%s", s, v.Type()) {
					return true
				}

				break
			}

			n.Value = v
		default:
			//dbg("", v.Kind())
			todo(n)
		}
	case 4: // '-' UnaryExpression
		switch v.Kind() {
		case ConstValue:
			n.Value = v.neg(ctx, n.UnaryExpression)
		default:
			//dbg("", v.Kind())
			todo(n)
		}
	case 5: // '^' UnaryExpression
		switch v.Kind() {
		case ConstValue:
			n.Value = v.cpl(ctx, n.UnaryExpression)
		case RuntimeValue:
			if !v.Type().IntegerType() {
				todo(n, true) // need int
				break
			}

			n.Value = newRuntimeValue(v.Type())
		default:
			//dbg("", v.Kind())
			todo(n)
		}
	case 6: // "<-" UnaryExpression
		n.flags = n.flags | flagLenPoisoned
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
		//dbg("%p", n.Value)
	default:
		panic("internal error")
	}
	return false
}

// -------------------------------------------------------------------- VarDecl

func (n *VarDecl) check(ctx *context) (stop bool) {
	// All is done in decl.go/varDecl()
	return false
}
