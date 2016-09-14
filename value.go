// Copyright 2016 The GC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gc

import (
	"fmt"
	"math"
	"math/big"
	"strings"
	"unicode"
)

const (
	digits = 6
)

var (
	_ Const = (*boolConst)(nil)
	_ Const = (*complexConst)(nil)
	_ Const = (*constBase)(nil)
	_ Const = (*floatConst)(nil)
	_ Const = (*intConst)(nil)
	_ Const = (*runeConst)(nil)
	_ Const = (*stringConst)(nil)

	_ stringValue = stringID(0)
	_ stringValue = stringIDs(nil)

	_ Value = (*constValue)(nil)
	_ Value = (*runtimeValue)(nil)
	_ Value = (*valueBase)(nil)

	bigInt1        = big.NewInt(1)
	bigIntMaxInt32 = big.NewInt(math.MaxInt32)
	bigIntMaxInt64 = big.NewInt(math.MaxInt64)
	bigIntMinInt32 = big.NewInt(math.MinInt32)
	bigIntMinInt64 = big.NewInt(math.MinInt64)

	bigIntMaxUint32 = big.NewInt(math.MaxUint32)
	bigIntMaxUint64 = big.NewInt(0).SetUint64(math.MaxUint64)

	bigFloatMaxUint64 = big.NewFloat(0).SetPrec(DefaultFloatConstPrec).SetUint64(math.MaxUint64)
	bigFloatMinInt64  = big.NewFloat(0).SetPrec(DefaultFloatConstPrec).SetInt64(math.MinInt64)

	floatZero    = 0.
	floatNegZero = -1 / (1 / floatZero)

	// see https://github.com/golang/go/blob/master/test/float_lit2.go#L21
	bigFloatHalfwayPointFloat32, _ = big.NewFloat(0).SetPrec(DefaultFloatConstPrec).SetString("340282356779733661637539395458142568448")
)

// Values of type ConstKind.
const (
	_ ConstKind = iota
	BoolConst
	ComplexConst
	FloatingPointConst
	IntConst
	RuneConst
	StringConst
)

// ConstKind is the specific kinf of a Const. The zero ConstKind is not a valid
// kind.
type ConstKind int

// Values of type ValueKind.
const (
	_            ValueKind = iota
	ConstValue             // Value known at compile time. Constant values have a type.
	NilValue               // Typeless zero value.
	PackageValue           // Value represents an imported package qualifier.
	RuntimeValue           // Value known only at run time. Runtime values have a type.
	TypeValue              // Value represents a type. Type values are a type.
)

// ValueKind is a the specific kind of a Value. The zero ValueKind is not a
// valid kind.
type ValueKind int

// Value is the representation of a Go expression.
//
// Not all methods apply to all kinds of values.  Restrictions, if any, are
// noted in the documentation for each method.  Use the Kind method to find out
// the kind of value before calling kind-specific methods.  Calling a method
// inappropriate to the kind of value causes a run-time panic.
type Value interface {
	add(ctx *context, node Node, op Value) Value
	and(ctx *context, node Node, op Value) Value
	andNot(ctx *context, node Node, op Value) Value
	assignableTo() bool
	boolAnd(ctx *context, node Node, op Value) Value
	boolOr(ctx *context, node Node, op Value) Value
	cpl(ctx *context, node Node) Value
	div(ctx *context, node Node, op Value) Value
	eq(ctx *context, node Node, op Value) Value
	ge(ctx *context, node Node, op Value) Value
	gt(ctx *context, node Node, op Value) Value
	lsh(ctx *context, node Node, op Value) Value
	le(ctx *context, node Node, op Value) Value
	lt(ctx *context, node Node, op Value) Value
	mod(ctx *context, node Node, op Value) Value
	mul(ctx *context, node Node, op Value) Value
	neg(ctx *context, node Node) Value
	neq(ctx *context, node Node, op Value) Value
	nonNegativeInteger(ctx *context) bool
	or(ctx *context, node Node, op Value) Value
	rsh(ctx *context, node Node, op Value) Value
	sub(ctx *context, node Node, op Value) Value
	xor(ctx *context, node Node, op Value) Value

	// Addressable reports whether a value is addressable. It panics if
	// value's Kind is not RuntimeValue.
	Addressable() bool

	// AssignableTo reports whether this value is assignable to type t.
	// It panics if value's Kind is not NilValue, RuntimeValue, TypeValue
	// or ConstValue.
	//
	// See https://golang.org/ref/spec#Assignability
	AssignableTo(ctx *Context, t Type) bool

	// Const returns the value's constant value. It panics if value Kind is
	// not ConstValue.
	Const() Const

	// Convert converts this value to type t and returns the new Value. It
	// returns nil if the conversion is not possible.  It panics if value's
	// Kind is not NilValue, TypeValue or ConstValue.
	//
	// See https://golang.org/ref/spec#Conversions
	Convert(ctx *Context, t Type) Value

	// Declaration returns the declaration a value refers to, if any.
	Declaration() Declaration

	// Integral reports whether the value is ∈ Z. It panics if the value is
	// not numeric.
	Integral() bool

	// Kind returns the specific kind of this value.
	Kind() ValueKind

	// Nil reports whether the value is nil. Nil panics if the value's Kind not NilValue
	// and the value's Type is not Ptr.
	Nil() bool

	// Selector returns the root value of a selector and its paths or (nil,
	// nil, nil) of the value does not represent a selector. It panics if
	// value's Kind is not RuntimeValue.
	//
	// path0 represents the original selector path, path represents the
	// full selector path. For example, in
	//
	//	type T struct {
	//		U
	//	}
	//
	//	type U struct {
	//		f int
	//	}
	//
	//	var v U
	//	var w = v.f
	//
	//	path0 is [f] and path is [U, f].
	Selector() (root Value, path0, path []Selector)

	// Type returns the type of this value or nil if the value's type
	// cannot be determined due to errors or if the value is a NilValue.
	Type() Type
}

type valueBase struct{ kind ValueKind }

func (v *valueBase) Addressable() bool                    { panic("Addressable of inappropriate value") }
func (v *valueBase) assignableTo() bool                   { panic("assignableTo of inappropriate value") }
func (v *valueBase) AssignableTo(*Context, Type) bool     { panic("AssignableTo of inappropriate value") }
func (v *valueBase) Const() Const                         { panic("Const of inappropriate value") }
func (v *valueBase) Convert(*Context, Type) Value         { panic("internal error") }
func (v *valueBase) Declaration() Declaration             { return nil }
func (v *valueBase) Integral() bool                       { panic("Integral of non-numeric value") }
func (v *valueBase) Kind() ValueKind                      { return v.kind }
func (v *valueBase) Nil() bool                            { panic("Nil of inappropriate value") }
func (v *valueBase) nonNegativeInteger(ctx *context) bool { return false }
func (v *valueBase) Type() Type                           { return nil }

func (v *valueBase) add(ctx *context, n Node, op Value) Value {
	ctx.err(n, "invalid operand for binary +")
	return nil
}

func (v *valueBase) and(ctx *context, n Node, op Value) Value {
	ctx.err(n, "invalid operand for binary &")
	return nil
}

func (v *valueBase) andNot(ctx *context, n Node, op Value) Value {
	ctx.err(n, "invalid operand for binary &^")
	return nil
}

func (v *valueBase) boolAnd(ctx *context, n Node, op Value) Value {
	ctx.err(n, "invalid operand for binary &&")
	return nil
}

func (v *valueBase) boolOr(ctx *context, n Node, op Value) Value {
	ctx.err(n, "invalid operand for binary ||")
	return nil
}

func (v *valueBase) cpl(ctx *context, n Node) Value {
	ctx.err(n, "invalid operand for unary ^")
	return nil
}

func (v *valueBase) div(ctx *context, n Node, op Value) Value {
	ctx.err(n, "invalid operand for binary /")
	return nil
}

func (v *valueBase) eq(ctx *context, n Node, op Value) Value {
	ctx.err(n, "invalid operand for binary ==")
	return nil
}

func (v *valueBase) le(ctx *context, n Node, op Value) Value {
	ctx.err(n, "invalid operand for binary <=")
	return nil
}

func (v *valueBase) lt(ctx *context, n Node, op Value) Value {
	ctx.err(n, "invalid operand for binary <")
	return nil
}

func (v *valueBase) ge(ctx *context, n Node, op Value) Value {
	ctx.err(n, "invalid operand for binary >=")
	return nil
}

func (v *valueBase) gt(ctx *context, n Node, op Value) Value {
	ctx.err(n, "invalid operand for binary >")
	return nil
}

func (v *valueBase) lsh(ctx *context, n Node, op Value) Value {
	ctx.err(n, "invalid operand for binary <<")
	return nil
}

func (v *valueBase) mod(ctx *context, n Node, op Value) Value {
	ctx.err(n, "invalid operand for binary %")
	return nil
}

func (v *valueBase) mul(ctx *context, n Node, op Value) Value {
	ctx.err(n, "invalid operand for binary *")
	return nil
}

func (v *valueBase) neg(ctx *context, n Node) Value {
	ctx.err(n, "invalid operand for unary -")
	return nil
}

func (v *valueBase) neq(ctx *context, n Node, op Value) Value {
	ctx.err(n, "invalid operand for binary !=")
	return nil
}

func (v *valueBase) or(ctx *context, n Node, op Value) Value {
	ctx.err(n, "invalid operand for binary |")
	return nil
}

func (v *valueBase) rsh(ctx *context, n Node, op Value) Value {
	ctx.err(n, "invalid operand for binary >>")
	return nil
}

func (v *valueBase) Selector() (Value, []Selector, []Selector) {
	panic("Selector of inappropriate value")
}

func (v *valueBase) sub(ctx *context, n Node, op Value) Value {
	ctx.err(n, "invalid operand for binary -")
	return nil
}

func (v *valueBase) xor(ctx *context, n Node, op Value) Value {
	ctx.err(n, "invalid operand for binary ^")
	return nil
}

// ----------------------------------------------------------------- constValue

type constValue struct {
	c Const
	valueBase
}

func newConstValue(c Const) *constValue { return &constValue{c, valueBase{ConstValue}} }

func (v *constValue) add(ctx *context, n Node, op Value) Value     { return v.c.add(ctx, n, op) }
func (v *constValue) and(ctx *context, n Node, op Value) Value     { return v.c.and(ctx, n, op) }
func (v *constValue) andNot(ctx *context, n Node, op Value) Value  { return v.c.andNot(ctx, n, op) }
func (v *constValue) assignableTo() bool                           { return false }
func (v *constValue) AssignableTo(ctx *Context, t Type) bool       { return v.c.AssignableTo(ctx, t) }
func (v *constValue) boolAnd(ctx *context, n Node, op Value) Value { return v.c.boolAnd(ctx, n, op) }
func (v *constValue) boolOr(ctx *context, n Node, op Value) Value  { return v.c.boolOr(ctx, n, op) }
func (v *constValue) Const() Const                                 { return v.c }
func (v *constValue) Convert(ctx *Context, u Type) Value           { return v.c.Convert(ctx, u) }
func (v *constValue) cpl(ctx *context, n Node) Value               { return v.c.cpl(ctx, n) }
func (v *constValue) div(ctx *context, n Node, op Value) Value     { return v.c.div(ctx, n, op) }
func (v *constValue) eq(ctx *context, n Node, op Value) Value      { return v.c.eq(ctx, n, op) }
func (v *constValue) ge(ctx *context, n Node, op Value) Value      { return v.c.ge(ctx, n, op) }
func (v *constValue) gt(ctx *context, n Node, op Value) Value      { return v.c.gt(ctx, n, op) }
func (v *constValue) Integral() bool                               { return v.c.Integral() }
func (v *constValue) lsh(ctx *context, n Node, op Value) Value     { return v.c.lsh(ctx, n, op) }
func (v *constValue) le(ctx *context, n Node, op Value) Value      { return v.c.le(ctx, n, op) }
func (v *constValue) lt(ctx *context, n Node, op Value) Value      { return v.c.lt(ctx, n, op) }
func (v *constValue) mod(ctx *context, n Node, op Value) Value     { return v.c.mod(ctx, n, op) }
func (v *constValue) mul(ctx *context, n Node, op Value) Value     { return v.c.mul(ctx, n, op) }
func (v *constValue) neg(ctx *context, n Node) Value               { return v.c.neg(ctx, n) }
func (v *constValue) neq(ctx *context, n Node, op Value) Value     { return v.c.neq(ctx, n, op) }
func (v *constValue) nonNegativeInteger(ctx *context) bool         { return v.c.nonNegativeInteger(ctx) }
func (v *constValue) or(ctx *context, n Node, op Value) Value      { return v.c.or(ctx, n, op) }
func (v *constValue) rsh(ctx *context, n Node, op Value) Value     { return v.c.rsh(ctx, n, op) }
func (v *constValue) String() string                               { return v.c.String() }
func (v *constValue) sub(ctx *context, n Node, op Value) Value     { return v.c.sub(ctx, n, op) }
func (v *constValue) Type() Type                                   { return v.c.Type() }
func (v *constValue) xor(ctx *context, n Node, op Value) Value     { return v.c.xor(ctx, n, op) }

// ------------------------------------------------------------------- nilValue

type nilValue struct{ valueBase }

func newNilValue() *nilValue { return &nilValue{valueBase{NilValue}} }

func (v *nilValue) Nil() bool { return true }

func (v *nilValue) assignableTo() bool { return false }

func (v *nilValue) AssignableTo(ctx *Context, t Type) bool {
	switch t.Kind() {
	case Ptr, Func, Slice, Map, Chan, Interface, UnsafePointer:
		return true
	}

	return false
}

func (v *nilValue) Convert(ctx *Context, u Type) Value {
	switch u.Kind() {
	case Ptr:
		return newNilPtrValue(u)
	case Func, Slice, Map, Chan, Interface:
		return newRuntimeValue(u)
	}

	return nil
}

// --------------------------------------------------------------- packageValue

type packageValue struct {
	d Declaration
	valueBase
}

func newPackageValue(d Declaration) *packageValue { return &packageValue{d, valueBase{PackageValue}} }

func (v *packageValue) assignableTo() bool { return false }

func (v *packageValue) Declaration() Declaration { return v.d }

// --------------------------------------------------------------- runtimeValue

type runtimeValue struct {
	addressable bool
	d           Declaration
	mapIndex    bool
	nil         bool
	path        []Selector
	path0       []Selector
	root        Value
	typ         Type
	valueBase

	//TODO shrink runtimeValue size by introducing a payload field.
}

func newRuntimeValue(typ Type) Value {
	if typ == nil {
		return nil
	}

	return &runtimeValue{
		typ:       typ,
		valueBase: valueBase{RuntimeValue},
	}
}

func newAddressableValue(typ Type) Value {
	if typ == nil {
		return nil
	}

	return &runtimeValue{
		addressable: true,
		typ:         typ,
		valueBase:   valueBase{RuntimeValue},
	}
}

func newMapIndexValue(typ Type) Value {
	if typ == nil {
		return nil
	}

	return &runtimeValue{
		mapIndex:  true,
		typ:       typ,
		valueBase: valueBase{RuntimeValue},
	}
}

func newDeclarationValue(d Declaration, t Type) Value {
	if d == nil || t == nil {
		return nil
	}

	return &runtimeValue{
		d:         d,
		typ:       t,
		valueBase: valueBase{RuntimeValue},
	}
}

func newSelectorValue(t Type, root Value, path0, path []Selector) Value {
	if t == nil {
		return nil
	}

	return &runtimeValue{
		addressable: true,
		path0:       path0,
		path:        path,
		root:        root,
		typ:         t,
		valueBase:   valueBase{RuntimeValue},
	}
}

func newNilPtrValue(t Type) *runtimeValue {
	if t.Kind() != Ptr && t.Kind() != UnsafePointer {
		panic("internal error")
	}

	return &runtimeValue{
		nil:       true,
		typ:       t,
		valueBase: valueBase{RuntimeValue},
	}
}

func (v *runtimeValue) Addressable() bool                      { return v.addressable }
func (v *runtimeValue) assignableTo() bool                     { return v.addressable || v.mapIndex }
func (v *runtimeValue) AssignableTo(ctx *Context, t Type) bool { return v.Type().AssignableTo(t) }
func (v *runtimeValue) Declaration() Declaration               { return v.d }

func (v *runtimeValue) nonNegativeInteger(ctx *context) bool {
	return v.Type().Numeric() && v.Integral()
}
func (v *runtimeValue) Selector() (Value, []Selector, []Selector) { return v.root, v.path0, v.path }
func (v *runtimeValue) Type() Type                                { return v.typ }

func (v *runtimeValue) add(ctx *context, n Node, op Value) Value {
	switch op.Kind() {
	case ConstValue:
		if !op.Const().AssignableTo(ctx.Context, v.Type()) {
			ctx.err(n, "invalid operation: mismatched types %s and %s", v.Type(), op.Type())
			return nil
		}

		return v
	default:
		//dbg("", op.Kind())
		todo(n)
	}
	return nil
}

func (v *runtimeValue) and(ctx *context, n Node, op Value) Value {
	switch op.Kind() {
	case ConstValue:
		if !op.Const().AssignableTo(ctx.Context, v.Type()) {
			todo(n, true) // ctx.err(n, "invalid operation: mismatched types %s and %s", v.Type(), op.Type())
			return nil
		}

		return v
	case RuntimeValue:
		if !op.Type().Identical(v.Type()) {
			todo(n, true) // type mismatch
			break
		}

		if !v.Type().IntegerType() {
			todo(n, true) //  need int
			break
		}

		return v
	default:
		todo(n)
	}
	return nil
}

func (v *runtimeValue) andNot(ctx *context, n Node, op Value) Value {
	todo(n)
	return nil
}

func (v *runtimeValue) boolAnd(ctx *context, n Node, op Value) Value {
	if t := ctx.booleanBinOpShape(v.Type(), op.Type(), n); t != nil {
		return newRuntimeValue(t)
	}

	ctx.err(n, "invalid operation: && (mismatched types %s and %s)", v.Type(), op.Type())
	return nil
}

func (v *runtimeValue) boolOr(ctx *context, n Node, op Value) Value {
	if t := ctx.booleanBinOpShape(v.Type(), op.Type(), n); t != nil {
		return newRuntimeValue(t)
	}

	ctx.err(n, "invalid operation: || (mismatched types %s and %s)", v.Type(), op.Type())
	return nil
}

func (v *runtimeValue) Convert(ctx *Context, u Type) Value {
	t := v.Type()
	switch {
	case t.Kind() == UnsafePointer:
		if k := u.Kind(); k == Ptr || k == Uintptr {
			return newRuntimeValue(u)
		}
	case u.Kind() == UnsafePointer:
		if t.Kind() == Ptr && v.Nil() {
			return newNilPtrValue(u)
		}

		if t.Kind() == Ptr || t.Kind() == Uintptr {
			return newRuntimeValue(u)
		}
	default:
		if t != nil && t.ConvertibleTo(u) {
			return newRuntimeValue(u)
		}
	}

	return nil
}

func (v *runtimeValue) cpl(ctx *context, n Node) Value {
	switch t := v.Type(); {
	case t.IntegerType():
		return v
	default:
		return nil
	}
}

func (v *runtimeValue) div(ctx *context, n Node, op Value) Value {
	todo(n)
	return nil
}

func (v *runtimeValue) eq(ctx *context, n Node, op Value) Value {
	ot := op.Type()
	switch op.Kind() {
	case ConstValue, RuntimeValue:
		if !v.AssignableTo(ctx.Context, ot) && !op.AssignableTo(ctx.Context, v.Type()) {
			ctx.err(n, "invalid operation: == (mismatched types %s and %s)", v.Type(), ot)
			break
		}

		if !v.Type().Comparable() {
			ctx.err(n, "invalid operation: == (operator == not defined on %s)", v.Type())
			break
		}

		if !ot.Comparable() {
			ctx.err(n, "invalid operation: == (operator == not defined on %s)", ot)
			break
		}

		return newRuntimeValue(ctx.untypedBoolType)
	case NilValue:
		if !op.AssignableTo(ctx.Context, v.Type()) {
			todo(n, true) // invalid operand
			break
		}

		return newRuntimeValue(ctx.untypedBoolType)
	default:
		//dbg("", op.Kind())
		todo(n)
	}
	return nil
}

func (v *runtimeValue) Integral() bool {
	if !v.Type().Numeric() {
		panic("Integral of non-numeric value")
	}

	return v.Type().IntegerType() || v.Kind() == ConstValue && v.Const().Integral()
}

func (v *runtimeValue) le(ctx *context, n Node, op Value) Value {
	ot := op.Type()
	switch op.Kind() {
	case ConstValue, RuntimeValue:
		if !v.AssignableTo(ctx.Context, ot) && !op.AssignableTo(ctx.Context, v.Type()) {
			//dbg("", v.Type(), ot)
			todo(n, true) // type mismatch
			break
		}

		if !v.Type().Ordered() {
			todo(n, true) // not ordered
			break
		}

		if !ot.Ordered() {
			todo(n, true) // not ordered
			break
		}

		return newRuntimeValue(ctx.untypedBoolType)
	default:
		//dbg("", op.Kind())
		todo(n)
	}
	return nil
}

func (v *runtimeValue) lt(ctx *context, n Node, op Value) Value {
	ot := op.Type()
	switch op.Kind() {
	case ConstValue, RuntimeValue:
		if !v.AssignableTo(ctx.Context, ot) && !op.AssignableTo(ctx.Context, v.Type()) {
			//dbg("", v.Type(), ot)
			todo(n, true) // type mismatch
			break
		}

		if !v.Type().Ordered() {
			todo(n, true) // not ordered
			break
		}

		if !ot.Ordered() {
			todo(n, true) // not ordered
			break
		}

		return newRuntimeValue(ctx.untypedBoolType)
	default:
		//dbg("", op.Kind())
		todo(n)
	}
	return nil
}

func (v *runtimeValue) ge(ctx *context, n Node, op Value) Value {
	ot := op.Type()
	switch op.Kind() {
	case ConstValue, RuntimeValue:
		if !v.AssignableTo(ctx.Context, ot) && !op.AssignableTo(ctx.Context, v.Type()) {
			ctx.err(n, "invalid operation: >= (mismatched types %s and %s)", v.Type(), ot)
			break
		}

		if !v.Type().Ordered() {
			ctx.err(n, "invalid operation: >= (operator >= not defined on %s)", v.Type())
			break
		}

		if !ot.Ordered() {
			ctx.err(n, "invalid operation: >= (operator >= not defined on %s)", ot)
			break
		}

		return newRuntimeValue(ctx.untypedBoolType)
	default:
		//dbg("", op.Kind())
		todo(n)
	}
	return nil
}

func (v *runtimeValue) gt(ctx *context, n Node, op Value) Value {
	ot := op.Type()
	switch op.Kind() {
	case ConstValue, RuntimeValue:
		if !v.AssignableTo(ctx.Context, ot) && !op.AssignableTo(ctx.Context, v.Type()) {
			ctx.err(n, "invalid operation: > (mismatched types %s and %s)", v.Type(), ot)
			break
		}

		if !v.Type().Ordered() {
			ctx.err(n, "invalid operation: > (operator > not defined on %s)", v.Type())
			break
		}

		if !ot.Ordered() {
			ctx.err(n, "invalid operation: > (operator > not defined on %s)", ot)
			break
		}

		return newRuntimeValue(ctx.untypedBoolType)
	default:
		//dbg("", op.Kind())
		todo(n)
	}
	return nil
}

func (v *runtimeValue) lsh(ctx *context, n Node, op Value) Value {
	//TODO non const shift rules
	return nil
}

func (v *runtimeValue) mod(ctx *context, n Node, op Value) Value {
	todo(n)
	return nil
}

func (v *runtimeValue) mul(ctx *context, n Node, op Value) Value {
	todo(n)
	return nil
}

func (v *runtimeValue) neg(ctx *context, n Node) Value {
	switch t := v.Type(); {
	case t.IntegerType(), t.FloatingPointType(), t.ComplexType():
		return v
	default:
		return v.valueBase.neg(ctx, n)
	}
}

func (v *runtimeValue) neq(ctx *context, n Node, op Value) Value {
	ot := op.Type()
	switch op.Kind() {
	case ConstValue, RuntimeValue:
		if !v.AssignableTo(ctx.Context, ot) && !op.AssignableTo(ctx.Context, v.Type()) {
			ctx.err(n, "invalid operation: != (mismatched types %s and %s)", v.Type(), ot)
			break
		}

		if !v.Type().Comparable() {
			ctx.err(n, "invalid operation: != (operator != not defined on %s)", v.Type())
			break
		}

		if !ot.Comparable() {
			ctx.err(n, "invalid operation: != (operator != not defined on %s)", ot)
			break
		}

		return newRuntimeValue(ctx.untypedBoolType)
	case NilValue:
		if !op.AssignableTo(ctx.Context, v.Type()) {
			todo(n, true) // invalid operand
			break
		}

		return newRuntimeValue(ctx.untypedBoolType)
	default:
		//dbg("", op.Kind())
		todo(n)
	}
	return nil
}

func (v *runtimeValue) Nil() bool {
	if v.Type().Kind() != Ptr {
		panic("internal error")
	}

	return v.nil
}

func (v *runtimeValue) or(ctx *context, n Node, op Value) Value {
	ot := op.Type()
	switch op.Kind() {
	case ConstValue:
		if !ot.IntegerType() {
			todo(n, true) // need int
			break
		}

		if !op.Const().Untyped() && !ot.AssignableTo(v.Type()) {
			todo(n, true) // type mismatch
			break
		}

		if !v.Type().IntegerType() {
			todo(n, true) // invalid operation
			break
		}

		return newRuntimeValue(v.Type())
	case RuntimeValue:
		if !v.Type().Identical(ot) {
			todo(n, true) //ctx.err(n, "invalid operation: | (mismatched types %s and %s)", v.Type(), ot)
			break
		}

		return newRuntimeValue(v.Type())
	default:
		//dbg("", op.Kind())
		todo(n)
	}
	return nil
}

func (v *runtimeValue) rsh(ctx *context, n Node, op Value) Value {
	todo(n)
	return nil
}

func (v *runtimeValue) sub(ctx *context, n Node, op Value) Value {
	if !v.Type().Numeric() {
		todo(n, true) // invalid operand
		return nil
	}

	ot := op.Type()
	if ot == nil {
		return nil
	}

	if !ot.Numeric() {
		todo(n, true) // invalid operand
		return nil
	}

	switch op.Kind() {
	case ConstValue:
		if op.Const().Untyped() && op.AssignableTo(ctx.Context, v.Type()) || v.Type().Identical(ot) {
			return newRuntimeValue(v.Type())
		}
	default:
		//dbg("", op.Kind())
		todo(n)
	}
	return nil
}

func (v *runtimeValue) xor(ctx *context, n Node, op Value) Value {
	ot := op.Type()
	switch op.Kind() {
	case ConstValue:
		if !ot.IntegerType() {
			todo(n, true) // need int
			break
		}

		if !op.Const().Untyped() && !ot.AssignableTo(v.Type()) { //TODO see .sub
			todo(n, true) // type mismatch
			break
		}

		if !v.Type().IntegerType() {
			ctx.err(n, "invalid operation: ^ (mismatched types %s and %s)", v.Type(), ot)
			break
		}

		return newRuntimeValue(v.Type())
	case RuntimeValue:
		if !v.Type().Identical(ot) {
			ctx.err(n, "invalid operation: ^ (mismatched types %s and %s)", v.Type(), ot)
			break
		}

		return newRuntimeValue(v.Type())
	default:
		//dbg("", op.Kind())
		todo(n)
	}
	return nil
}

// ------------------------------------------------------------------ typeValue

type typeValue struct {
	t Type
	valueBase
}

func newTypeValue(t Type) Value {
	if t == nil {
		return nil
	}

	return &typeValue{t, valueBase{TypeValue}}
}

func (v *typeValue) assignableTo() bool                     { return false }
func (v *typeValue) AssignableTo(ctx *Context, t Type) bool { return v.Type().AssignableTo(t) }
func (v *typeValue) String() string                         { return v.t.String() }
func (v *typeValue) Type() Type                             { return v.t }

func (v *typeValue) Convert(ctx *Context, u Type) Value {
	if v.Type().ConvertibleTo(u) {
		return newTypeValue(u)
	}

	return nil
}

type mapKey struct {
	b bool
	i int64
	c complex128
}

// Const is the representatopn of a constant Go expression.
//
// Not all methods apply to all kinds of constants.  Restrictions, if any, are
// noted in the documentation for each method.  Use the Kind method to find out
// the kind of constant before calling kind-specific methods.  Calling a method
// inappropriate to the kind of constant causes a run-time panic.
type Const interface {
	add(ctx *context, node Node, op Value) Value
	add0(ctx *context, n Node, t Type, untyped bool, op Const) Const
	and(ctx *context, node Node, op Value) Value
	and0(ctx *context, n Node, t Type, untyped bool, op Const) Const
	andNot(ctx *context, node Node, op Value) Value
	andNot0(ctx *context, n Node, t Type, untyped bool, op Const) Const
	boolAnd(ctx *context, node Node, op Value) Value
	boolOr(ctx *context, node Node, op Value) Value
	convert(*context, Type) Const // Result is untyped.
	cpl(ctx *context, node Node) Value
	div(ctx *context, n Node, op Value) Value
	div0(ctx *context, n Node, t Type, untyped bool, op Const) Const
	eq(ctx *context, node Node, op Value) Value
	eq0(ctx *context, n Node, t Type, untyped bool, op Const) Const
	ge(ctx *context, node Node, op Value) Value
	ge0(ctx *context, n Node, t Type, untyped bool, op Const) Const
	gt(ctx *context, node Node, op Value) Value
	gt0(ctx *context, n Node, t Type, untyped bool, op Const) Const
	int() int64
	lsh(ctx *context, node Node, op Value) Value
	le(ctx *context, node Node, op Value) Value
	le0(ctx *context, n Node, t Type, untyped bool, op Const) Const
	lt(ctx *context, node Node, op Value) Value
	lt0(ctx *context, n Node, t Type, untyped bool, op Const) Const
	mapKey() mapKey
	mod(ctx *context, node Node, op Value) Value
	mod0(ctx *context, n Node, t Type, untyped bool, op Const) Const
	mul(ctx *context, n Node, op Value) Value
	mul0(ctx *context, n Node, t Type, untyped bool, op Const) Const
	mustConvert(*context, Node, Type) Const // Result is untyped.
	neg(ctx *context, node Node) Value
	neq(ctx *context, node Node, op Value) Value
	neq0(ctx *context, n Node, t Type, untyped bool, op Const) Const
	nonNegativeInteger(*context) bool
	normalize(ctx *context) Const // Keeps exact value or returns nil on overflow.
	or(ctx *context, node Node, op Value) Value
	or0(ctx *context, n Node, t Type, untyped bool, op Const) Const
	representable(*context, Type) Const // Can be inexact for floats.
	rsh(ctx *context, node Node, op Value) Value
	sub(ctx *context, node Node, op Value) Value
	sub0(ctx *context, n Node, t Type, untyped bool, op Const) Const
	xor(ctx *context, node Node, op Value) Value
	xor0(ctx *context, n Node, t Type, untyped bool, op Const) Const

	// AssignableTo reports whether this constant is assignable to type t.
	//
	// See https://golang.org/ref/spec#Assignability
	AssignableTo(ctx *Context, t Type) bool

	// Convert converts this constant to type t and returns the new Value.
	// It returns nil if the conversion is not possible.
	//
	// See https://golang.org/ref/spec#Conversions
	Convert(ctx *Context, t Type) Value

	// ConvertibleTo reports whether this constant is convertible to type
	// u.
	//
	// See https://golang.org/ref/spec#Conversions
	ConvertibleTo(ctx *Context, u Type) bool

	// Integral reports whether the constant's value is ∈ Z. It panics if
	// the constant is not numeric.
	Integral() bool

	// Kind returns the specific kind of the constant.
	Kind() ConstKind

	// Numeric reports whether the constant's Kind is one of RuneConst,
	// IntConst, FloatingPointConst or ComplexConst.
	Numeric() bool

	// String returns a string representation of the constant's value.
	String() string

	// Type returns the constant's type.
	Type() Type

	// Untyped reports whether the constant is untyped.
	Untyped() bool
}

type constBase struct {
	kind    ConstKind
	typ     Type
	untyped bool
}

func (c *constBase) add0(*context, Node, Type, bool, Const) Const    { panic("internal error") }
func (c *constBase) and0(*context, Node, Type, bool, Const) Const    { panic("internal error") }
func (c *constBase) andNot0(*context, Node, Type, bool, Const) Const { panic("internal error") }
func (c *constBase) AssignableTo(*Context, Type) bool                { panic("internal error") }
func (c *constBase) convert(*context, Type) Const                    { panic("internal error") }
func (c *constBase) Convert(*Context, Type) Value                    { panic("internal error") }
func (c *constBase) ConvertibleTo(ctx *Context, u Type) bool         { panic("internal error") }
func (c *constBase) div0(*context, Node, Type, bool, Const) Const    { panic("internal error") }
func (c *constBase) eq0(*context, Node, Type, bool, Const) Const     { panic("internal error") }
func (c *constBase) ge0(*context, Node, Type, bool, Const) Const     { panic("internal error") }
func (c *constBase) gt0(*context, Node, Type, bool, Const) Const     { panic("internal error") }
func (c *constBase) int() int64                                      { panic("internal error") }
func (c *constBase) Integral() bool                                  { panic("internal error") }
func (c *constBase) Kind() ConstKind                                 { return c.kind }
func (c *constBase) le0(*context, Node, Type, bool, Const) Const     { panic("internal error") }
func (c *constBase) lt0(*context, Node, Type, bool, Const) Const     { panic("internal error") }
func (c *constBase) mapKey() mapKey                                  { panic("internal error") }
func (c *constBase) mod0(*context, Node, Type, bool, Const) Const    { panic("internal error") }
func (c *constBase) mul0(*context, Node, Type, bool, Const) Const    { panic("internal error") }
func (c *constBase) mustConvert(*context, Node, Type) Const          { panic("internal error") }
func (c *constBase) neq0(*context, Node, Type, bool, Const) Const    { panic("internal error") }
func (c *constBase) nonNegativeInteger(*context) bool                { return false }
func (c *constBase) normalize(ctx *context) Const                    { panic("internal error") }
func (c *constBase) or0(*context, Node, Type, bool, Const) Const     { panic("internal error") }
func (c *constBase) representable(*context, Type) Const              { panic("internal error") }
func (c *constBase) String() string                                  { panic("internal error") }
func (c *constBase) sub0(*context, Node, Type, bool, Const) Const    { panic("internal error") }
func (c *constBase) Type() Type                                      { return c.typ }
func (c *constBase) Untyped() bool                                   { return c.untyped }
func (c *constBase) xor0(*context, Node, Type, bool, Const) Const    { panic("internal error") }

func (c *constBase) mustConvertConst(ctx *context, n Node, d Const, t Type) Const {
	e := d.convert(ctx, t)
	if e == nil || e.representable(ctx, t) != nil {
		return e
	}

	ctx.constConversionFail(n, t, d)
	return nil
}

func (c *constBase) assignableTo(ctx *context, cc Const, t Type) bool {
	if c.Untyped() && cc.representable(ctx, t) != nil {
		return true
	}

	return c.Type().AssignableTo(t)
}

func (c *constBase) Numeric() bool {
	return c.kind == ComplexConst || c.kind == FloatingPointConst ||
		c.kind == IntConst || c.kind == RuneConst
}

func (c *constBase) representableComplexFromBigComplex(ctx *context, v *bigComplex, t Type) Const {
	switch t.Kind() {
	case Complex64, Complex128:
		re := c.representableFloatFromBigFloat(v.re, ctx.float64Type)
		if re == nil {
			return nil
		}

		im := c.representableFloatFromBigFloat(v.im, ctx.float64Type)
		if im == nil {
			return nil
		}

		return c.representableComplexFromComplex(complex(re.(*floatConst).val, im.(*floatConst).val), t)
	}
	panic("internal error")
}

func (c *constBase) representableComplexFromBigFloat(ctx *context, v *big.Float, t Type) Const {
	return c.representableComplexFromBigComplex(ctx, &bigComplex{v, big.NewFloat(0)}, t)
}

func (c *constBase) representableComplexFromBigInt(ctx *context, v *big.Int, t Type) Const {
	return c.representableComplexFromBigFloat(ctx, big.NewFloat(0).SetPrec(ctx.floatConstPrec).SetInt(v), t)
}

func (c *constBase) representableComplexFromComplex(v complex128, t Type) Const {
	switch t.Kind() {
	case Complex64:
		if math.Abs(real(v)) <= math.MaxFloat32 && math.Abs(imag(v)) <= math.MaxFloat32 {
			return newComplexConst(v, nil, t, false)
		}
	case Complex128:
		return newComplexConst(v, nil, t, false)
	default:
		panic("internal error")
	}
	return nil
}

func (c *constBase) representableComplexFromFloat(v float64, t Type) Const {
	return c.representableComplexFromComplex(complex(v, 0), t)
}

func (c *constBase) representableComplexFromInt(v int64, t Type) Const {
	return c.representableComplexFromFloat(float64(v), t)
}

func (c *constBase) representableFloatFromBigComplex(v *bigComplex, t Type) Const {
	if v.im.Sign() != 0 {
		return nil
	}

	return c.representableFloatFromBigFloat(v.re, t)
}

func (c *constBase) representableFloatFromBigFloat(v *big.Float, t Type) Const {
	switch t.Kind() {
	case Float32:
		var w big.Float
		if w.Abs(v).Cmp(bigFloatHalfwayPointFloat32) < 0 {
			f, _ := v.Float64()
			return newFloatConst(f, nil, t, false)
		}
	case Float64:
		f, _ := v.Float64()
		return c.representableFloatFromFloat(f, t)
	default:
		panic("internal error")
	}
	return nil
}

func (c *constBase) representableFloatFromBigInt(ctx *context, v *big.Int, t Type) Const {
	return c.representableFloatFromBigFloat(big.NewFloat(0).SetPrec(ctx.floatConstPrec).SetInt(v), t)
}

func (c *constBase) representableFloatFromComplex(v complex128, t Type) Const {
	if imag(v) != 0 {
		return nil
	}

	return c.representableFloatFromFloat(real(v), t)
}

func (c *constBase) representableFloatFromFloat(v float64, t Type) Const {
	if math.IsInf(v, 0) {
		return nil
	}

	if v == floatNegZero {
		v = 0
	}
	switch t.Kind() {
	case Float32:
		if f := float32(v); !math.IsInf(float64(f), 0) {
			return newFloatConst(v, nil, t, false)
		}
	case Float64:
		return newFloatConst(v, nil, t, false)
	default:
		//dbg("", t.Kind())
		panic("internal error")
	}
	return nil
}

func (c *constBase) representableFloatFromInt(v int64, t Type) Const {
	return c.representableFloatFromFloat(float64(v), t)
}

func (c *constBase) representableIntFromBigComplex(ctx *context, v *bigComplex, t Type) Const {
	if v.im.Sign() != 0 {
		return nil
	}

	return c.representableIntFromBigFloat(ctx, v.re, t)
}

func (c *constBase) representableIntFromBigFloat(ctx *context, v *big.Float, t Type) Const {
	if !v.IsInt() {
		return nil
	}

	if v.Cmp(bigFloatMinInt64) < 0 || v.Cmp(bigFloatMaxUint64) > 0 {
		return nil
	}

	n, _ := v.Int(nil)
	return c.representableIntFromBigInt(ctx, n, t)
}

func (c *constBase) representableIntFromBigInt(ctx *context, v *big.Int, t Type) Const {
	if v.Cmp(bigIntMinInt64) < 0 || v.Cmp(bigIntMaxUint64) > 0 {
		return nil
	}

	if v.Cmp(bigIntMaxInt64) <= 0 {
		return c.representableIntFromInt(ctx, v.Int64(), t)
	}

	// v in (math.MaxInt64, math.MaxUint64].
	switch t.Kind() {
	case Int8, Int16, Int32, Int64, Int, Uint8, Uint16, Uint32:
		return nil
	case Uint64:
		return newIntConst(0, v, t, false)
	case Uint:
		if ctx.model.IntBytes == 8 {
			return newIntConst(0, v, t, false)
		}
	case Uintptr:
		if ctx.model.PtrBytes == 8 {
			return newIntConst(0, v, t, false)
		}
	default:
		panic("internal error")
	}
	return nil
}

func (c *constBase) representableIntFromComplex(ctx *context, v complex128, t Type) Const {
	if imag(v) != 0 {
		return nil
	}

	return c.representableIntFromFloat(ctx, real(v), t)
}

func (c *constBase) representableIntFromFloat(ctx *context, v float64, t Type) Const {
	if math.Floor(v) != v {
		return nil
	}

	if v < math.MinInt64 || v > math.MaxInt64 {
		return c.representableIntFromBigFloat(ctx, big.NewFloat(0).SetPrec(ctx.floatConstPrec).SetFloat64(v), t)
	}

	n := int64(v)
	switch t.Kind() {
	case Int8:
		if n >= math.MinInt8 && n <= math.MaxInt8 {
			return newIntConst(n, nil, t, false)
		}
	case Int16:
		if n >= math.MinInt16 && n <= math.MaxInt16 {
			return newIntConst(n, nil, t, false)
		}
	case Int32:
		if n >= math.MinInt32 && n <= math.MaxInt32 {
			return newIntConst(n, nil, t, false)
		}
	case Int64:
		return newIntConst(n, nil, t, false)
	case Int:
		if n >= ctx.minInt && n <= ctx.maxInt {
			return newIntConst(n, nil, t, false)
		}
	case Uint8:
		if n >= 0 && n <= math.MaxUint8 {
			return newIntConst(n, nil, t, false)
		}
	case Uint16:
		if n >= 0 && n <= math.MaxUint16 {
			return newIntConst(n, nil, t, false)
		}
	case Uint32:
		if n >= 0 && n <= math.MaxUint32 {
			return newIntConst(n, nil, t, false)
		}
	case Uint64:
		if n >= 0 {
			return newIntConst(n, nil, t, false)
		}
	case Uint:
		if n >= 0 && uint64(n) <= ctx.maxUint {
			return newIntConst(n, nil, t, false)
		}
	case Uintptr:
		if n >= 0 && uint64(n) <= ctx.maxUintptr {
			return newIntConst(n, nil, t, false)
		}
	default:
		panic("internal error")
	}
	return nil
}

func (c *constBase) representableIntFromInt(ctx *context, v int64, t Type) Const {
	switch t.Kind() {
	case Int8:
		if v >= math.MinInt8 && v <= math.MaxInt8 {
			return newIntConst(v, nil, t, false)
		}
	case Int16:
		if v >= math.MinInt16 && v <= math.MaxInt16 {
			return newIntConst(v, nil, t, false)
		}
	case Int32:
		if v >= math.MinInt32 && v <= math.MaxInt32 {
			return newIntConst(v, nil, t, false)
		}
	case Int64:
		return newIntConst(v, nil, t, false)
	case Int:
		if v >= ctx.minInt && v <= ctx.maxInt {
			return newIntConst(v, nil, t, false)
		}
	case Uint8:
		if v >= 0 && v <= math.MaxUint8 {
			return newIntConst(v, nil, t, false)
		}
	case Uint16:
		if v >= 0 && v <= math.MaxUint16 {
			return newIntConst(v, nil, t, false)
		}
	case Uint32:
		if v >= 0 && v <= math.MaxUint32 {
			return newIntConst(v, nil, t, false)
		}
	case Uint64:
		if v >= 0 {
			return newIntConst(v, nil, t, false)
		}
	case Uint:
		if v >= 0 && uint64(v) <= ctx.maxUint {
			return newIntConst(v, nil, t, false)
		}
	case Uintptr:
		if v >= 0 && uint64(v) <= ctx.maxUintptr {
			return newIntConst(v, nil, t, false)
		}
	default:
		panic("internal error")
	}
	return nil
}

func (c *constBase) add(ctx *context, n Node, op Value) Value {
	ctx.err(n, "invalid operand for binary +")
	return nil
}

func (c *constBase) and(ctx *context, n Node, op Value) Value {
	ctx.err(n, "invalid operand for binary &")
	return nil
}

func (c *constBase) andNot(ctx *context, n Node, op Value) Value {
	ctx.err(n, "invalid operand for binary &^")
	return nil
}

func (c *constBase) boolAnd(ctx *context, n Node, op Value) Value {
	ctx.err(n, "invalid operand for binary &&")
	return nil
}

func (c *constBase) boolOr(ctx *context, n Node, op Value) Value {
	ctx.err(n, "invalid operand for binary ||")
	return nil
}

func (c *constBase) cpl(ctx *context, n Node) Value {
	ctx.err(n, "invalid operand for unary ^")
	return nil
}

func (c *constBase) div(ctx *context, n Node, op Value) Value {
	ctx.err(n, "invalid operand for binary /")
	return nil
}

func (c *constBase) eq(ctx *context, n Node, op Value) Value {
	ctx.err(n, "invalid operand for binary ==")
	return nil
}

func (c *constBase) le(ctx *context, n Node, op Value) Value {
	ctx.err(n, "invalid operand for binary <=")
	return nil
}

func (c *constBase) lt(ctx *context, n Node, op Value) Value {
	ctx.err(n, "invalid operand for binary <")
	return nil
}

func (c *constBase) ge(ctx *context, n Node, op Value) Value {
	ctx.err(n, "invalid operand for binary >=")
	return nil
}

func (c *constBase) gt(ctx *context, n Node, op Value) Value {
	ctx.err(n, "invalid operand for binary >")
	return nil
}

func (c *constBase) lsh(ctx *context, n Node, op Value) Value {
	ctx.err(n, "invalid operand for binary *")
	return nil
}

func (c *constBase) mod(ctx *context, n Node, op Value) Value {
	ctx.err(n, "invalid operand for binary %")
	return nil
}

func (c *constBase) mul(ctx *context, n Node, op Value) Value {
	ctx.err(n, "invalid operand for binary *")
	return nil
}

func (c *constBase) neg(ctx *context, n Node) Value {
	ctx.err(n, "invalid operand for unary -")
	return nil
}

func (c *constBase) neq(ctx *context, n Node, op Value) Value {
	ctx.err(n, "invalid operand for binary !=")
	return nil
}

func (c *constBase) or(ctx *context, n Node, op Value) Value {
	ctx.err(n, "invalid operand for binary |")
	return nil
}

func (c *constBase) rsh(ctx *context, n Node, op Value) Value {
	ctx.err(n, "invalid operand for binary >>")
	return nil
}

func (c *constBase) sub(ctx *context, n Node, op Value) Value {
	ctx.err(n, "invalid operand for binary -")
	return nil
}

func (c *constBase) xor(ctx *context, n Node, op Value) Value {
	ctx.err(n, "invalid operand for binary ^")
	return nil
}

// ------------------------------------------------------------------ boolConst

type boolConst struct {
	constBase
	val bool
}

func newBoolConst(val bool, typ Type, untyped bool) *boolConst {
	return &boolConst{constBase{BoolConst, typ, untyped}, val}
}

func (c *boolConst) AssignableTo(ctx *Context, t Type) bool {
	return c.assignableTo(&context{Context: ctx}, c, t)
}

func (c *boolConst) mapKey() mapKey               { return mapKey{b: c.val} }
func (c *boolConst) normalize(ctx *context) Const { return c }

func (c *boolConst) ConvertibleTo(ctx *Context, u Type) bool {
	// A constant value x can be converted to type T
	//
	// · x is representable by a value of type T.

	// · x is a floating-point constant, T is a floating-point type, and x is
	// representable by a value of type T after rounding using IEEE 754
	// round-to-even rules, but with an IEEE -0.0 further rounded to an unsigned
	// 0.0. The constant T(x) is the rounded value.

	// · x is an integer constant and T is a string type. The same rule as for
	// non-constant x applies in this case.

	return false
}

func (c *boolConst) boolAnd(ctx *context, n Node, op Value) Value {
	switch op.Kind() {
	case ConstValue:
		t, untyped, a, b := ctx.constBooleanBinOpShape(c, op.Const(), n)
		if t == nil {
			break
		}

		return newConstValue(newBoolConst(a.(*boolConst).val && b.(*boolConst).val, t, untyped))
	case RuntimeValue:
		if t := ctx.booleanBinOpShape(c.Type(), op.Type(), n); t != nil {
			return newRuntimeValue(t)
		}

		ctx.err(n, "invalid operation: && (mismatched types %s and %s)", c.Type(), op.Type())
	default:
		//dbg("", op.Kind())
		todo(n)
	}
	return nil
}

func (c *boolConst) boolOr(ctx *context, n Node, op Value) Value {
	switch op.Kind() {
	case ConstValue:
		t, untyped, a, b := ctx.constBooleanBinOpShape(c, op.Const(), n)
		if t == nil {
			break
		}

		return newConstValue(newBoolConst(a.(*boolConst).val || b.(*boolConst).val, t, untyped))
	case RuntimeValue:
		if t := ctx.booleanBinOpShape(c.Type(), op.Type(), n); t != nil {
			return newRuntimeValue(t)
		}

		ctx.err(n, "invalid operation: || (mismatched types %s and %s)", c.Type(), op.Type())
	default:
		//dbg("", op.Kind())
		todo(n)
	}
	return nil
}

func (c *boolConst) eq(ctx *context, n Node, op Value) Value {
	todo(n)
	return nil
}

func (c *boolConst) neq(ctx *context, n Node, op Value) Value {
	todo(n)
	return nil
}

func (c *boolConst) representable(ctx *context, t Type) Const {
	if t.Kind() == Bool {
		return c
	}

	return nil
}

func (c *boolConst) Convert(ctx *Context, t Type) Value {
	if t.Kind() == Bool {
		return newConstValue(newBoolConst(c.val, t, false))
	}

	return nil
}

func (c *boolConst) String() string { return fmt.Sprint(c.val) }

func (c *boolConst) Type() Type {
	if c.untyped {
		return untypedBoolType
	}

	return c.typ
}

// --------------------------------------------------------------- complexConst

type complexConst struct {
	constBase
	val    complex128
	bigVal *bigComplex
}

func newComplexConst(val complex128, bigVal *bigComplex, typ Type, untyped bool) *complexConst {
	return &complexConst{constBase{ComplexConst, typ, untyped}, val, bigVal}
}

func (c *complexConst) AssignableTo(ctx *Context, t Type) bool {
	return c.assignableTo(&context{Context: ctx}, c, t)
}

func (c *complexConst) mustConvert(ctx *context, n Node, t Type) Const {
	return c.mustConvertConst(ctx, n, c, t)
}

func (c *complexConst) ConvertibleTo(ctx *Context, u Type) bool {
	// A constant value x can be converted to type T
	//
	// · x is representable by a value of type T.

	return c.representable(&context{Context: ctx}, u) != nil

	// · x is a floating-point constant, T is a floating-point type, and x is
	// representable by a value of type T after rounding using IEEE 754
	// round-to-even rules, but with an IEEE -0.0 further rounded to an unsigned
	// 0.0. The constant T(x) is the rounded value.

	// · x is an integer constant and T is a string type. The same rule as for
	// non-constant x applies in this case.
}

func (c *complexConst) add(ctx *context, n Node, op Value) Value {
	todo(n)
	return nil
}

func (c *complexConst) eq(ctx *context, n Node, op Value) Value {
	todo(n)
	return nil
}

func (c *complexConst) Convert(ctx *Context, t Type) Value {
	if d := c.representable(&context{Context: ctx}, t); d != nil {
		return newConstValue(d)
	}

	return nil
}

func (c *complexConst) mapKey() mapKey {
	if c.bigVal != nil {
		re, _ := c.bigVal.re.Float64()
		im, _ := c.bigVal.im.Float64()
		return mapKey{c: complex(re, im)}
	}

	return mapKey{c: c.val}
}

func (c *complexConst) neq(ctx *context, n Node, op Value) Value {
	todo(n)
	return nil
}

func (c *complexConst) nonNegativeInteger(ctx *context) bool {
	return c.representable(ctx, ctx.intType) != nil
}

func (c *complexConst) normalize(ctx *context) Const {
	if v := c.bigVal; v != nil {
		if re, ok := v.re.Float64(); ok == big.Exact {
			if im, ok := v.im.Float64(); ok == big.Exact {
				c = newComplexConst(complex(re, im), nil, c.Type(), c.Untyped())
			}
		}
	}

	if c.Untyped() {
		return c
	}

	return c.representable(ctx, c.Type())
}

func (c *complexConst) representable(ctx *context, t Type) Const {
	switch {
	case t.IntegerType():
		if v := c.bigVal; v != nil {
			return c.representableIntFromBigComplex(ctx, v, t)
		}

		return c.representableIntFromComplex(ctx, c.val, t)
	case t.FloatingPointType():
		if v := c.bigVal; v != nil {
			return c.representableFloatFromBigComplex(v, t)
		}

		return c.representableFloatFromComplex(c.val, t)
	case t.ComplexType():
		if v := c.bigVal; v != nil {
			return c.representableComplexFromBigComplex(ctx, v, t)
		}

		return c.representableComplexFromComplex(c.val, t)
	}
	return nil
}

func (c *complexConst) String() string {
	if c.bigVal != nil {
		return c.bigVal.String()
	}

	return fmt.Sprint(c.val)
}

func (c *complexConst) convert(ctx *context, t Type) Const {
	switch {
	case t.IntegerType():
		todo(zeroNode)
	case t.FloatingPointType():
		todo(zeroNode)
	case t.ComplexType():
		todo(zeroNode)
	default:
		panic("internal error")
	}
	return nil
}

func (c *complexConst) div0(ctx *context, n Node, t Type, untyped bool, op Const) Const {
	todo(n)
	return nil
}

func (c *complexConst) div(ctx *context, n Node, op Value) Value {
	todo(n)
	return nil
}

func (c *complexConst) mul0(ctx *context, n Node, t Type, untyped bool, op Const) Const {
	var d bigComplex
	d.mul(ctx, c.bigVal, op.(*complexConst).bigVal)
	if untyped {
		t = nil
	}
	e := newComplexConst(0, &d, ctx.complex128Type, true)
	if e.normalize(ctx) == nil {
		todo(n, true) // ctx.err(n, "constant multiplication overflow")
		return nil
	}

	return ctx.mustConvertConst(n, t, e)
}

func (c *complexConst) mul(ctx *context, n Node, op Value) Value {
	switch op.Kind() {
	case ConstValue:
		t, untyped, a, b := ctx.arithmeticBinOpShape(c, op.Const(), n)
		if t != nil {
			if d := a.mul0(ctx, n, t, untyped, b); d != nil {
				return newConstValue(d)
			}
		}
	default:
		todo(n)
	}
	return nil
}

func (c *complexConst) neg(ctx *context, n Node) Value {
	todo(n)
	return nil
}

func (c *complexConst) sub(ctx *context, n Node, op Value) Value {
	todo(n)
	return nil
}

func (c *complexConst) Integral() bool {
	if c.bigVal != nil {
		if c.bigVal.im.Sign() != 0 {
			return false
		}

		return c.bigVal.re.IsInt()
	}

	return imag(c.val) == 0 && math.Floor(real(c.val)) == real(c.val)
}

// ----------------------------------------------------------------- floatConst

type floatConst struct {
	constBase
	val    float64
	bigVal *big.Float
}

func newFloatConst(val float64, bigVal *big.Float, typ Type, untyped bool) *floatConst {
	return &floatConst{constBase{FloatingPointConst, typ, untyped}, val, bigVal}
}

func (c *floatConst) AssignableTo(ctx *Context, t Type) bool {
	return c.assignableTo(&context{Context: ctx}, c, t)
}

func (c *floatConst) mustConvert(ctx *context, n Node, t Type) Const {
	return c.mustConvertConst(ctx, n, c, t)
}

func (c *floatConst) ConvertibleTo(ctx *Context, u Type) bool {
	// A constant value x can be converted to type T
	//
	// · x is representable by a value of type T.

	if c.representable(&context{Context: ctx}, u) != nil {
		return true
	}

	// · x is a floating-point constant, T is a floating-point type, and x is
	// representable by a value of type T after rounding using IEEE 754
	// round-to-even rules, but with an IEEE -0.0 further rounded to an unsigned
	// 0.0. The constant T(x) is the rounded value.

	todo(zeroNode)
	return false

	// · x is an integer constant and T is a string type. The same rule as for
	// non-constant x applies in this case.
}

func (c *floatConst) add0(ctx *context, n Node, t Type, untyped bool, op Const) Const {
	var d big.Float
	d.Add(c.bigVal, op.(*floatConst).bigVal)
	if untyped {
		t = nil
	}
	e := newFloatConst(0, &d, ctx.float64Type, true)
	if e.normalize(ctx) == nil {
		todo(n, true) // {over,under}flow
		return nil
	}

	return ctx.mustConvertConst(n, t, e)
}

func (c *floatConst) add(ctx *context, n Node, op Value) Value {
	switch op.Kind() {
	case ConstValue:
		t, untyped, a, b := ctx.arithmeticBinOpShape(c, op.Const(), n)
		if t != nil {
			if d := a.add0(ctx, n, t, untyped, b); d != nil {
				return newConstValue(d)
			}
		}
	default:
		//dbg("", op.Kind())
		todo(n)
	}
	return nil
}

func (c *floatConst) Convert(ctx *Context, t Type) Value {
	if d := c.representable(&context{Context: ctx}, t); d != nil {
		return newConstValue(d)
	}

	return nil
}

func (c *floatConst) mapKey() mapKey {
	if c.bigVal != nil {
		f, _ := c.bigVal.Float64()
		return mapKey{c: complex(f, 0)}
	}

	return mapKey{c: complex(c.val, 0)}
}

func (c *floatConst) nonNegativeInteger(ctx *context) bool {
	return c.representable(ctx, ctx.intType) != nil
}

func (c *floatConst) normalize(ctx *context) Const {
	if v := c.bigVal; v != nil {
		if f, ok := v.Float64(); ok == big.Exact {
			c = newFloatConst(f, nil, c.Type(), c.Untyped())
		}
	}

	if c.Untyped() {
		return c
	}

	return c.representable(ctx, c.Type())
}

func (c *floatConst) representable(ctx *context, t Type) Const {
	switch {
	case t.IntegerType():
		if v := c.bigVal; v != nil {
			return c.representableIntFromBigFloat(ctx, v, t)
		}

		return c.representableIntFromFloat(ctx, c.val, t)
	case t.FloatingPointType():
		if v := c.bigVal; v != nil {
			return c.representableFloatFromBigFloat(v, t)
		}

		return c.representableFloatFromFloat(c.val, t)
	case t.ComplexType():
		if v := c.bigVal; v != nil {
			return c.representableComplexFromBigFloat(ctx, v, t)
		}

		return c.representableComplexFromFloat(c.val, t)
	}
	return nil
}

func (c *floatConst) String() string {
	if c.bigVal != nil {
		return bigFloatString(c.bigVal)
	}

	s := fmt.Sprint(c.val)
	d := strings.IndexByte(s, '.')
	if d < 0 {
		return s
	}

	e := strings.IndexByte(s, 'e')
	if e < 0 {
		e = len(s)
	}
	f := e - d - 1
	if f > 5 {
		f = 5
	}
	return s[:d+1] + s[d+1:d+1+f] + s[e:]
}

func (c *floatConst) convert(ctx *context, t Type) Const {
	switch {
	case t.IntegerType():
		if !c.Integral() {
			return nil
		}

		switch {
		case c.bigVal != nil:
			i, _ := c.bigVal.Int(nil)
			return newIntConst(0, i, ctx.intType, true)
		default:
			i, _ := big.NewFloat(0).SetPrec(ctx.floatConstPrec).SetFloat64(c.val).Int(nil)
			return newIntConst(0, i, ctx.intType, true)
		}
	case t.FloatingPointType():
		if c.bigVal != nil {
			return newFloatConst(0, c.bigVal, ctx.float64Type, true)
		}

		return newFloatConst(0, big.NewFloat(0).SetPrec(ctx.floatConstPrec).SetFloat64(c.val), ctx.float64Type, true)
	case t.ComplexType():
		if c.bigVal != nil {
			return newComplexConst(
				0,
				&bigComplex{
					c.bigVal,
					big.NewFloat(0).SetPrec(ctx.floatConstPrec).SetFloat64(0),
				},
				ctx.complex128Type,
				true,
			)
		}

		return newComplexConst(
			0,
			&bigComplex{
				big.NewFloat(0).SetPrec(ctx.floatConstPrec).SetFloat64(c.val),
				big.NewFloat(0).SetPrec(ctx.floatConstPrec).SetFloat64(0),
			},
			ctx.complex128Type,
			true,
		)
	default:
		panic("internal error")
	}
}

func (c *floatConst) div0(ctx *context, n Node, t Type, untyped bool, op Const) Const {
	d := op.(*floatConst).bigVal
	if d.Sign() == 0 {
		ctx.err(n, "division by zero")
		return nil
	}

	var e big.Float
	e.Quo(c.bigVal, d)
	if untyped {
		t = nil
	}
	return ctx.mustConvertConst(n, t, newFloatConst(0, &e, ctx.float64Type, true))
}

func (c *floatConst) div(ctx *context, n Node, op Value) Value {
	switch op.Kind() {
	case ConstValue:
		t, untyped, a, b := ctx.arithmeticBinOpShape(c, op.Const(), n)
		if t != nil {
			if d := a.div0(ctx, n, t, untyped, b); d != nil {
				return newConstValue(d)
			}
		}
	default:
		//dbg("", op.Kind())
		todo(n)
	}
	return nil
}

func (c *floatConst) eq(ctx *context, n Node, op Value) Value {
	todo(n)
	return nil
}

func (c *floatConst) le(ctx *context, n Node, op Value) Value {
	todo(n)
	return nil
}

func (c *floatConst) lt(ctx *context, n Node, op Value) Value {
	todo(n)
	return nil
}

func (c *floatConst) lsh(ctx *context, n Node, op Value) Value {
	switch op.Kind() {
	case ConstValue:
		if i := c.convert(ctx, ctx.intType); i != nil {
			return i.lsh(ctx, n, op)
		}

		ctx.err(n, "invalid operand for binary <<")
	case RuntimeValue:
		//TODO non const shift rules
	default:
		//dbg("", op.Kind())
		todo(n)
	}
	return nil
}

func (c *floatConst) ge(ctx *context, n Node, op Value) Value {
	todo(n)
	return nil
}

func (c *floatConst) gt(ctx *context, n Node, op Value) Value {
	todo(n)
	return nil
}

func (c *floatConst) mod0(ctx *context, n Node, t Type, untyped bool, op Const) Const {
	ctx.err(n, "illegal constant expression: floating-point %% operation")
	return nil
}

func (c *floatConst) mul0(ctx *context, n Node, t Type, untyped bool, op Const) Const {
	var d big.Float
	d.Mul(c.bigVal, op.(*floatConst).bigVal)
	if untyped {
		t = nil
	}
	e := newFloatConst(0, &d, ctx.float64Type, true)
	if e.normalize(ctx) == nil {
		todo(n, true) // ctx.err(n, "constant multiplication overflow")
		return nil
	}

	return ctx.mustConvertConst(n, t, e)
}

func (c *floatConst) mul(ctx *context, n Node, op Value) Value {
	switch op.Kind() {
	case ConstValue:
		t, untyped, a, b := ctx.arithmeticBinOpShape(c, op.Const(), n)
		if t != nil {
			if d := a.mul0(ctx, n, t, untyped, b); d != nil {
				return newConstValue(d)
			}
		}
	default:
		//dbg("", op.Kind())
		todo(n)
	}
	return nil
}

func (c *floatConst) neg(ctx *context, n Node) Value {
	if c.bigVal != nil {
		var d big.Float
		return newConstValue(newFloatConst(0, d.Neg(c.bigVal), c.Type(), c.Untyped()))
	}

	return newConstValue(newFloatConst(-c.val, nil, c.Type(), c.Untyped()))
}

func (c *floatConst) neq(ctx *context, n Node, op Value) Value {
	todo(n)
	return nil
}

func (c *floatConst) rsh(ctx *context, n Node, op Value) Value {
	todo(n)
	return nil
}

func (c *floatConst) sub0(ctx *context, n Node, t Type, untyped bool, op Const) Const {
	var d big.Float
	d.Sub(c.bigVal, op.(*floatConst).bigVal)
	if untyped {
		t = nil
	}
	e := newFloatConst(0, &d, ctx.float64Type, true)
	if e.normalize(ctx) == nil {
		todo(n, true) // {over,under}flow
		return nil
	}

	return ctx.mustConvertConst(n, t, e)
}

func (c *floatConst) sub(ctx *context, n Node, op Value) Value {
	switch op.Kind() {
	case ConstValue:
		t, untyped, a, b := ctx.arithmeticBinOpShape(c, op.Const(), n)
		if t != nil {
			if d := a.sub0(ctx, n, t, untyped, b); d != nil {
				return newConstValue(d)
			}
		}
	default:
		//dbg("", op.Kind())
		todo(n)
	}
	return nil
}

func (c *floatConst) Integral() bool {
	if c.bigVal != nil {
		return c.bigVal.IsInt()
	}

	return math.Floor(c.val) == c.val
}

// ------------------------------------------------------------------- intConst

type intConst struct {
	constBase
	val    int64
	bigVal *big.Int
}

func newIntConst(val int64, bigVal *big.Int, typ Type, untyped bool) *intConst {
	return &intConst{constBase{IntConst, typ, untyped}, val, bigVal}
}

func (c *intConst) AssignableTo(ctx *Context, t Type) bool {
	return c.assignableTo(&context{Context: ctx}, c, t)
}

func (c *intConst) Integral() bool { return true }
func (c *intConst) int() int64     { return c.val }

func (c *intConst) mustConvert(ctx *context, n Node, t Type) Const {
	return c.mustConvertConst(ctx, n, c, t)
}

func (c *intConst) ConvertibleTo(ctx *Context, u Type) bool {
	// A constant value x can be converted to type T
	//
	// · x is representable by a value of type T.

	// · x is a floating-point constant, T is a floating-point type, and x is
	// representable by a value of type T after rounding using IEEE 754
	// round-to-even rules, but with an IEEE -0.0 further rounded to an unsigned
	// 0.0. The constant T(x) is the rounded value.

	// · x is an integer constant and T is a string type. The same rule as for
	// non-constant x applies in this case.

	return c.Convert(ctx, u) != nil
}

func (c *intConst) add0(ctx *context, n Node, t Type, untyped bool, op Const) Const {
	var d big.Int
	d.Add(c.bigVal, op.(*intConst).bigVal)
	if untyped {
		t = nil
	}
	return ctx.mustConvertConst(n, t, newIntConst(0, &d, ctx.intType, true))
}

func (c *intConst) add(ctx *context, n Node, op Value) Value {
	switch op.Kind() {
	case ConstValue:
		t, untyped, a, b := ctx.arithmeticBinOpShape(c, op.Const(), n)
		if t != nil {
			if d := a.add0(ctx, n, t, untyped, b); d != nil {
				return newConstValue(d)
			}
		}
	default:
		//dbg("", op.Kind())
		todo(n)
	}
	return nil
}

func (c *intConst) and0(ctx *context, n Node, t Type, untyped bool, op Const) Const {
	var d big.Int
	d.And(c.bigVal, op.(*intConst).bigVal)
	if untyped {
		t = nil
	}
	return ctx.mustConvertConst(n, t, newIntConst(0, &d, ctx.intType, true))
}

func (c *intConst) and(ctx *context, n Node, op Value) Value {
	switch op.Kind() {
	case ConstValue:
		t, untyped, a, b := ctx.arithmeticBinOpShape(c, op.Const(), n)
		if t != nil {
			if d := a.and0(ctx, n, t, untyped, b); d != nil {
				return newConstValue(d)
			}
		}
	default:
		//dbg("", op.Kind())
		todo(n)
	}
	return nil
}

func (c *intConst) andNot0(ctx *context, n Node, t Type, untyped bool, op Const) Const {
	var d big.Int
	d.AndNot(c.bigVal, op.(*intConst).bigVal)
	if untyped {
		t = nil
	}
	return ctx.mustConvertConst(n, t, newIntConst(0, &d, ctx.intType, true))
}

func (c *intConst) andNot(ctx *context, n Node, op Value) Value {
	switch op.Kind() {
	case ConstValue:
		t, untyped, a, b := ctx.arithmeticBinOpShape(c, op.Const(), n)
		if t != nil {
			if d := a.andNot0(ctx, n, t, untyped, b); d != nil {
				return newConstValue(d)
			}
		}
	default:
		//dbg("", op.Kind())
		todo(n)
	}
	return nil
}

func (c *intConst) Convert(ctx *Context, t Type) Value {
	switch t.Kind() {
	case String:
		return newRuntimeValue(t)
	default:
		if d := c.representable(&context{Context: ctx}, t); d != nil {
			return newConstValue(d)
		}

		if c.Type().Kind() == Uintptr && t.Kind() == UnsafePointer {
			return newRuntimeValue(t)
		}
	}

	return nil
}

func (c *intConst) mapKey() mapKey {
	if c.bigVal != nil {
		return mapKey{i: c.bigVal.Int64()}
	}

	return mapKey{i: c.val}
}

func (c *intConst) eq0(ctx *context, n Node, t Type, untyped bool, op Const) Const {
	return newBoolConst(c.bigVal.Cmp(op.(*intConst).bigVal) == 0, ctx.boolType, true)
}

func (c *intConst) eq(ctx *context, n Node, op Value) Value {
	switch op.Kind() {
	case ConstValue:
		t, untyped, a, b := ctx.arithmeticBinOpShape(c, op.Const(), n)
		if t != nil {
			if d := a.eq0(ctx, n, t, untyped, b); d != nil {
				return newConstValue(d)
			}
		}
	case RuntimeValue:
		ot := op.Type()
		if !c.AssignableTo(ctx.Context, ot) && !ot.AssignableTo(c.Type()) {
			ctx.err(n, "invalid operation: == (mismatched types %s and %s)", c.Type(), ot)
			break
		}

		if !ot.Comparable() {
			ctx.err(n, "invalid operation: == (operator == not defined on %s)", ot)
			break
		}

		return newRuntimeValue(untypedBoolType)
	default:
		//dbg("", op.Kind())
		todo(n)
	}
	return nil
}

func (c *intConst) le0(ctx *context, n Node, t Type, untyped bool, op Const) Const {
	return newBoolConst(c.bigVal.Cmp(op.(*intConst).bigVal) <= 0, ctx.boolType, true)
}

func (c *intConst) le(ctx *context, n Node, op Value) Value {
	switch op.Kind() {
	case ConstValue:
		t, untyped, a, b := ctx.arithmeticBinOpShape(c, op.Const(), n)
		if t != nil {
			if d := a.le0(ctx, n, t, untyped, b); d != nil {
				return newConstValue(d)
			}
		}
	default:
		//dbg("", op.Kind())
		todo(n)
	}
	return nil
}

func (c *intConst) lt0(ctx *context, n Node, t Type, untyped bool, op Const) Const {
	return newBoolConst(c.bigVal.Cmp(op.(*intConst).bigVal) < 0, ctx.boolType, true)
}

func (c *intConst) lt(ctx *context, n Node, op Value) Value {
	switch op.Kind() {
	case ConstValue:
		t, untyped, a, b := ctx.arithmeticBinOpShape(c, op.Const(), n)
		if t != nil {
			if d := a.lt0(ctx, n, t, untyped, b); d != nil {
				return newConstValue(d)
			}
		}
	default:
		//dbg("", op.Kind())
		todo(n)
	}
	return nil
}

func (c *intConst) ge0(ctx *context, n Node, t Type, untyped bool, op Const) Const {
	return newBoolConst(c.bigVal.Cmp(op.(*intConst).bigVal) >= 0, ctx.boolType, true)
}

func (c *intConst) ge(ctx *context, n Node, op Value) Value {
	switch op.Kind() {
	case ConstValue:
		t, untyped, a, b := ctx.arithmeticBinOpShape(c, op.Const(), n)
		if t != nil {
			if d := a.gt0(ctx, n, t, untyped, b); d != nil {
				return newConstValue(d)
			}
		}
	case RuntimeValue:
		ot := op.Type()
		if !c.AssignableTo(ctx.Context, ot) && !ot.AssignableTo(c.Type()) {
			ctx.err(n, "invalid operation: >= (mismatched types %s and %s)", c.Type(), ot)
			break
		}

		if !ot.Ordered() {
			ctx.err(n, "invalid operation: >= (operator >= not defined on %s)", ot)
			break
		}

		return newRuntimeValue(untypedBoolType)
	default:
		//dbg("", op.Kind())
		todo(n)
	}
	return nil
}

func (c *intConst) gt0(ctx *context, n Node, t Type, untyped bool, op Const) Const {
	return newBoolConst(c.bigVal.Cmp(op.(*intConst).bigVal) > 0, ctx.boolType, true)
}

func (c *intConst) gt(ctx *context, n Node, op Value) Value {
	switch op.Kind() {
	case ConstValue:
		t, untyped, a, b := ctx.arithmeticBinOpShape(c, op.Const(), n)
		if t != nil {
			if d := a.gt0(ctx, n, t, untyped, b); d != nil {
				return newConstValue(d)
			}
		}
	default:
		//dbg("", op.Kind())
		todo(n)
	}
	return nil
}

func (c *intConst) neq0(ctx *context, n Node, t Type, untyped bool, op Const) Const {
	return newBoolConst(c.bigVal.Cmp(op.(*intConst).bigVal) != 0, ctx.boolType, true)
}

func (c *intConst) neq(ctx *context, n Node, op Value) Value {
	switch op.Kind() {
	case ConstValue:
		t, untyped, a, b := ctx.arithmeticBinOpShape(c, op.Const(), n)
		if t != nil {
			if d := a.eq0(ctx, n, t, untyped, b); d != nil {
				return newConstValue(d)
			}
		}
	case RuntimeValue:
		ot := op.Type()
		if !c.AssignableTo(ctx.Context, ot) && !op.Type().AssignableTo(c.Type()) {
			ctx.err(n, "invalid operation: != (mismatched types %s and %s)", c.Type(), ot)
			break
		}

		if !ot.Comparable() {
			ctx.err(n, "invalid operation: != (operator != not defined on %s)", ot)
			break
		}

		return newRuntimeValue(untypedBoolType)
	default:
		//dbg("", op.Kind())
		todo(n)
	}
	return nil
}

func (c *intConst) nonNegativeInteger(ctx *context) bool {
	return c.representable(ctx, ctx.intType) != nil
}

func (c *intConst) normalize(ctx *context) Const {
	if v := c.bigVal; v != nil {
		if uint64(v.BitLen()) > uint64(ctx.intConstBits) {
			return nil
		}

		if v.Cmp(bigIntMinInt64) >= 0 && v.Cmp(bigIntMaxInt64) <= 0 {
			c = newIntConst(v.Int64(), nil, c.Type(), c.Untyped())
		}
	}

	if c.Untyped() {
		return c
	}

	return c.representable(ctx, c.Type())
}

func (c *intConst) or0(ctx *context, n Node, t Type, untyped bool, op Const) Const {
	var d big.Int
	d.Or(c.bigVal, op.(*intConst).bigVal)
	if untyped {
		t = nil
	}
	return ctx.mustConvertConst(n, t, newIntConst(0, &d, ctx.intType, true))
}

func (c *intConst) or(ctx *context, n Node, op Value) Value {
	switch op.Kind() {
	case ConstValue:
		t, untyped, a, b := ctx.arithmeticBinOpShape(c, op.Const(), n)
		if t != nil {
			if d := a.or0(ctx, n, t, untyped, b); d != nil {
				return newConstValue(d)
			}
		}
	case RuntimeValue:
		if !op.Type().IntegerType() {
			todo(n, true) // need integer
			break
		}

		if d := c.Convert(ctx.Context, op.Type()); d == nil {
			todo(n, true) // type/value mismatch
			break
		}

		return newRuntimeValue(op.Type())
	default:
		//dbg("", op.Kind())
		todo(n)
	}
	return nil
}

func (c *intConst) representable(ctx *context, t Type) Const {
	switch {
	case t.IntegerType():
		if v := c.bigVal; v != nil {
			return c.representableIntFromBigInt(ctx, v, t)
		}

		return c.representableIntFromInt(ctx, c.val, t)
	case t.ComplexType():
		if v := c.bigVal; v != nil {
			return c.representableComplexFromBigInt(ctx, v, t)
		}

		return c.representableComplexFromInt(c.val, t)
	case t.FloatingPointType():
		if v := c.bigVal; v != nil {
			return c.representableFloatFromBigInt(ctx, v, t)
		}

		return c.representableFloatFromInt(c.val, t)
	}
	return nil
}

func (c *intConst) String() string {
	if c.bigVal != nil {
		return fmt.Sprint(c.bigVal)
	}

	return fmt.Sprint(c.val)
}

func (c *intConst) convert(ctx *context, t Type) Const {
	switch {
	case t.IntegerType():
		if c.bigVal != nil {
			return newIntConst(0, c.bigVal, ctx.intType, true)
		}

		return newIntConst(0, big.NewInt(c.val), ctx.intType, true)
	case t.FloatingPointType():
		if c.bigVal != nil {
			return newFloatConst(0, big.NewFloat(0).SetPrec(ctx.floatConstPrec).SetInt(c.bigVal), ctx.intType, true)
		}

		return newFloatConst(0, big.NewFloat(0).SetPrec(ctx.floatConstPrec).SetInt64(c.val), ctx.intType, true)
	case t.ComplexType():
		todo(zeroNode)
	default:
		panic("internal error")
	}
	return nil
}

func (c *intConst) cpl(ctx *context, n Node) Value {
	if c.bigVal != nil {
		var v big.Int
		if d := newIntConst(0, v.Sub(v.Neg(c.bigVal), bigInt1), c.Type(), c.Untyped()).normalize(ctx); d != nil {
			return newConstValue(d)
		}

		return nil
	}

	if d := newIntConst(^c.val, nil, c.Type(), c.Untyped()).normalize(ctx); d != nil {
		return newConstValue(d)
	}

	return nil
}

func (c *intConst) div0(ctx *context, n Node, t Type, untyped bool, op Const) Const {
	d := op.(*intConst).bigVal
	if d.Sign() == 0 {
		ctx.err(n, "division by zero")
		return nil
	}

	var e big.Int
	e.Div(c.bigVal, d)
	if untyped {
		t = nil
	}
	return ctx.mustConvertConst(n, t, newIntConst(0, &e, ctx.intType, true))
}

func (c *intConst) div(ctx *context, n Node, op Value) Value {
	switch op.Kind() {
	case ConstValue:
		t, untyped, a, b := ctx.arithmeticBinOpShape(c, op.Const(), n)
		if t != nil {
			if d := a.div0(ctx, n, t, untyped, b); d != nil {
				return newConstValue(d)
			}
		}
	case RuntimeValue:
		if !op.Type().Numeric() {
			todo(n, true) // invalid operand
			break
		}

		if !c.ConvertibleTo(ctx.Context, op.Type()) {
			todo(n, true) // type mismatch
			break
		}

		return newRuntimeValue(op.Type())
	default:
		//dbg("", op.Kind())
		todo(n)
	}
	return nil
}

func (c *intConst) lsh(ctx *context, n Node, op Value) Value {
	switch op.Kind() {
	case ConstValue:
		d := op.Const()
		if !d.Integral() {
			todo(n, true) // invalid shift count
			return nil
		}

		// The right operand in a shift expression must have unsigned
		// integer type or be an untyped constant that can be converted
		// to unsigned integer type.
		var e *intConst
		switch {
		case d.Untyped():
			if e = d.convert(ctx, ctx.intType).(*intConst); e.bigVal.Sign() < 0 {
				todo(n, true)
				return nil
			}
		default:
			if !d.Type().UnsignedIntegerType() {
				todo(n, true)
				return nil
			}

			e = d.convert(ctx, ctx.intType).(*intConst)
		}

		if uint(e.bigVal.BitLen()) > ctx.intConstBits {
			todo(n, true)
			return nil
		}

		f := c.convert(ctx, ctx.intType).(*intConst)
		var g big.Int
		h := newIntConst(0, g.Lsh(f.bigVal, uint(e.bigVal.Uint64())), ctx.intType, true).normalize(ctx)
		if h == nil {
			ctx.err(n, "constant shift overflow")
			return nil
		}

		// If the left operand of a constant shift expression is an
		// untyped constant, the result is an integer constant;
		// otherwise it is a constant of the same type as the left
		// operand, which must be of integer type.
		switch {
		case c.Untyped():
			return newConstValue(h)
		default:
			if i := ctx.mustConvertConst(n, c.Type(), h); i != nil {
				return newConstValue(i)
			}
		}
	case RuntimeValue:
		//TODO non const shift rules
	default:
		//dbg("", op.Kind())
		todo(n)
	}
	return nil
}

func (c *intConst) mod0(ctx *context, n Node, t Type, untyped bool, op Const) Const {
	todo(n)
	return nil
	//	ctx := t.context()
	//	d := op.(*intConst).bigVal
	//	if d.Sign() == 0 {
	//		ctx.err(n, "division by zero")
	//		return nil
	//	}
	//
	//	var e big.Int
	//	e.div(c.bigVal, d)
	//	if untyped {
	//		t = nil
	//	}
	//	return ctx.mustConvertConst(n, t, newIntConst(0, &e, ctx.intType, true))
}

func (c *intConst) mod(ctx *context, n Node, op Value) Value {
	switch op.Kind() {
	case ConstValue:
		t, untyped, a, b := ctx.arithmeticBinOpShape(c, op.Const(), n)
		if t != nil {
			if d := a.mod0(ctx, n, t, untyped, b); d != nil {
				return newConstValue(d)
			}
		}
	default:
		//dbg("", op.Kind())
		todo(n)
	}
	return nil
}

func (c *intConst) mul0(ctx *context, n Node, t Type, untyped bool, op Const) Const {
	var d big.Int
	d.Mul(c.bigVal, op.(*intConst).bigVal)
	if untyped {
		t = nil
	}
	e := newIntConst(0, &d, ctx.intType, true)
	if e.normalize(ctx) == nil {
		ctx.err(n, "constant multiplication overflow")
		return nil
	}

	return ctx.mustConvertConst(n, t, e)
}

func (c *intConst) mul(ctx *context, n Node, op Value) Value {
	switch op.Kind() {
	case ConstValue:
		t, untyped, a, b := ctx.arithmeticBinOpShape(c, op.Const(), n)
		if t != nil {
			if d := a.mul0(ctx, n, t, untyped, b); d != nil {
				return newConstValue(d)
			}
		}
	default:
		//dbg("", op.Kind())
		todo(n)
	}
	return nil
}

func (c *intConst) neg(ctx *context, n Node) Value {
	if c.bigVal != nil {
		var v big.Int
		if d := newIntConst(0, v.Neg(v.Set(c.bigVal)), c.Type(), c.Untyped()).normalize(&context{}); d != nil {
			return newConstValue(d)
		}

		return nil
	}

	var d Const
	switch {
	case c.val == math.MinInt64:
		var v big.Int
		d = newIntConst(0, v.Neg(v.SetInt64(c.val)), c.Type(), c.Untyped()).normalize(ctx)
	default:
		d = newIntConst(-c.val, nil, c.Type(), c.Untyped()).normalize(ctx)
	}
	if d != nil {
		return newConstValue(d)
	}

	return nil
}

func (c *intConst) rsh(ctx *context, n Node, op Value) Value {
	switch op.Kind() {
	case ConstValue:
		d := op.Const()
		if !d.Integral() {
			todo(n, true)
			return nil
		}

		// The right operand in a shift expression must have unsigned
		// integer type or be an untyped constant that can be converted
		// to unsigned integer type.
		var e *intConst
		switch {
		case d.Untyped():
			if e = d.convert(ctx, ctx.intType).(*intConst); e.bigVal.Sign() < 0 {
				todo(n, true)
				return nil
			}
		default:
			if !d.Type().UnsignedIntegerType() {
				todo(n, true)
				return nil
			}

			e = d.convert(ctx, ctx.intType).(*intConst)
		}

		if uint(e.bigVal.BitLen()) > ctx.intConstBits {
			todo(n, true)
			return nil
		}

		f := c.convert(ctx, ctx.intType).(*intConst)
		var g big.Int
		h := newIntConst(0, g.Rsh(f.bigVal, uint(e.bigVal.Uint64())), ctx.intType, true).normalize(ctx)
		if h == nil {
			ctx.err(n, "constant shift overflow")
			return nil
		}

		// If the left operand of a constant shift expression is an
		// untyped constant, the result is an integer constant;
		// otherwise it is a constant of the same type as the left
		// operand, which must be of integer type.
		switch {
		case c.Untyped():
			return newConstValue(h)
		default:
			if i := ctx.mustConvertConst(n, c.Type(), h); i != nil {
				return newConstValue(i)
			}
		}
	default:
		//dbg("", op.Kind())
		todo(n)
	}
	return nil
}

func (c *intConst) sub0(ctx *context, n Node, t Type, untyped bool, op Const) Const {
	var d big.Int
	d.Sub(c.bigVal, op.(*intConst).bigVal)
	if untyped {
		t = nil
	}
	return ctx.mustConvertConst(n, t, newIntConst(0, &d, ctx.intType, true))
}

func (c *intConst) sub(ctx *context, n Node, op Value) Value {
	switch op.Kind() {
	case ConstValue:
		t, untyped, a, b := ctx.arithmeticBinOpShape(c, op.Const(), n)
		if t != nil {
			if d := a.sub0(ctx, n, t, untyped, b); d != nil {
				return newConstValue(d)
			}
		}
	default:
		//dbg("", op.Kind())
		todo(n)
	}
	return nil
}

func (c *intConst) xor0(ctx *context, n Node, t Type, untyped bool, op Const) Const {
	var d big.Int
	d.Xor(c.bigVal, op.(*intConst).bigVal)
	if untyped {
		t = nil
	}
	return ctx.mustConvertConst(n, t, newIntConst(0, &d, ctx.intType, true))
}

func (c *intConst) xor(ctx *context, n Node, op Value) Value {
	ot := op.Type()
	switch op.Kind() {
	case ConstValue:
		t, untyped, a, b := ctx.arithmeticBinOpShape(c, op.Const(), n)
		if t != nil {
			if d := a.xor0(ctx, n, t, untyped, b); d != nil {
				return newConstValue(d)
			}
		}
	case RuntimeValue:
		if !ot.IntegerType() {
			ctx.err(n, "invalid operation: ^ (mismatched types %s and %s)", c.Type(), ot)
			break
		}

		if !c.untyped && !c.Type().AssignableTo(ot) {
			todo(n, true) // type mismatch
			break
		}

		return newRuntimeValue(ot)
	default:
		//dbg("", op.Kind())
		todo(n)
	}
	return nil
}

// ------------------------------------------------------------------ runeConst

type runeConst struct {
	intConst
}

func newRuneConst(val rune, bigVal *big.Int, typ Type, untyped bool) *runeConst {
	return &runeConst{intConst{constBase{RuneConst, typ, untyped}, int64(val), bigVal}}
}

func (c *runeConst) ConvertibleTo(ctx *Context, u Type) bool {
	// A constant value x can be converted to type T
	//
	// · x is representable by a value of type T.

	// · x is a floating-point constant, T is a floating-point type, and x is
	// representable by a value of type T after rounding using IEEE 754
	// round-to-even rules, but with an IEEE -0.0 further rounded to an unsigned
	// 0.0. The constant T(x) is the rounded value.

	// · x is an integer constant and T is a string type. The same rule as for
	// non-constant x applies in this case.

	return c.Convert(ctx, u) != nil
}

func (c *runeConst) String() string {
	if c.bigVal != nil {
		return fmt.Sprint(c.bigVal)
	}

	if v := c.val; v >= 0 && v <= unicode.MaxRune {
		return fmt.Sprintf("'%c'", v)
	}

	return fmt.Sprint(c.val)
}

// ---------------------------------------------------------------- stringConst

type stringConst struct {
	constBase
	val stringValue
}

func newStringConst(val stringValue, typ Type, untyped bool) *stringConst {
	return &stringConst{constBase{StringConst, typ, untyped}, val}
}

func (c *stringConst) AssignableTo(ctx *Context, t Type) bool {
	return c.assignableTo(&context{Context: ctx}, c, t)
}

func (c *stringConst) mapKey() mapKey               { return mapKey{i: int64(dict.ID(c.val.s()))} }
func (c *stringConst) normalize(ctx *context) Const { return c }

func (c *stringConst) mustConvert(ctx *context, n Node, t Type) Const {
	return c.mustConvertConst(ctx, n, c, t)
}

func (c *stringConst) ConvertibleTo(ctx *Context, u Type) bool {
	// A constant value x can be converted to type T
	//
	// · x is representable by a value of type T.

	return u.Kind() == String

	// · x is a floating-point constant, T is a floating-point type, and x is
	// representable by a value of type T after rounding using IEEE 754
	// round-to-even rules, but with an IEEE -0.0 further rounded to an unsigned
	// 0.0. The constant T(x) is the rounded value.

	// · x is an integer constant and T is a string type. The same rule as for
	// non-constant x applies in this case.
}

func (c *stringConst) add(ctx *context, n Node, op Value) Value {
	switch op.Kind() {
	case ConstValue:
		t, untyped, a, b := ctx.constStringBinOpShape(c, op.Const(), n)
		if t != nil {
			return newConstValue(newStringConst(a.(*stringConst).val.cat(b.(*stringConst).val), t, untyped))
		}
	case RuntimeValue:
		if t := ctx.stringBinOpShape(newConstValue(c), op, n); t != nil {
			return newRuntimeValue(t)
		}
	default:
		//dbg("", op.Kind())
		todo(n)
	}
	ctx.err(n, "invalid operation: + (mismatched types %s and %s)", c.Type(), op.Type())
	return nil
}

func (c *stringConst) convert(ctx *context, t Type) Const {
	switch t.Kind() {
	case String:
		return newStringConst(c.val, ctx.stringType, true)
	default:
		panic("internal error")
	}
}

func (c *stringConst) eq0(ctx *context, n Node, t Type, untyped bool, op Const) Const {
	return newBoolConst(c.val.eq(op.(*stringConst).val), ctx.boolType, true)
}

func (c *stringConst) eq(ctx *context, n Node, op Value) Value {
	switch op.Kind() {
	case ConstValue:
		t, untyped, a, b := ctx.constStringBinOpShape(c, op.Const(), n)
		if t != nil {
			if d := a.eq0(ctx, n, t, untyped, b); d != nil {
				return newConstValue(d)
			}
		}
	default:
		//dbg("", op.Kind())
		todo(n)
	}
	return nil
}

func (c *stringConst) le(ctx *context, n Node, op Value) Value {
	todo(n)
	return nil
}

func (c *stringConst) lt(ctx *context, n Node, op Value) Value {
	todo(n)
	return nil
}

func (c *stringConst) ge(ctx *context, n Node, op Value) Value {
	todo(n)
	return nil
}

func (c *stringConst) gt(ctx *context, n Node, op Value) Value {
	todo(n)
	return nil
}

func (c *stringConst) Convert(ctx *Context, t Type) Value {
	switch t.Kind() {
	case Slice:
		if e := t.Elem(); e.Kind() == Uint8 || e.Kind() == Int32 {
			return newRuntimeValue(t)
		}
	case String:
		return newConstValue(newStringConst(c.val, t, false))
	}
	return nil
}

func (c *stringConst) representable(ctx *context, t Type) Const {
	if t.Kind() == String {
		return newStringConst(c.val, t, false)
	}

	return nil
}

func (c *stringConst) String() string { return fmt.Sprintf("%q", c.val.s()) }

// ----------------------------------------------------------------- bigComplex

// bigComplex is the value of a untyped complex constant.
type bigComplex struct {
	re, im *big.Float
}

func (z *bigComplex) String() string {
	if z.im.Sign() >= 0 {
		return fmt.Sprintf("%s+%si", bigFloatString(z.re), bigFloatString(z.im))
	}

	return fmt.Sprintf("%s%si", bigFloatString(z.re), bigFloatString(z.im))
}

func (z *bigComplex) mul(ctx *context, x, y *bigComplex) *bigComplex {
	// (a+bi)(c+di) = (ac-bd)+(bc+ad)i
	var r bigComplex
	var s big.Float
	r.re.Sub(s.Mul(x.re, y.re), s.Mul(x.im, y.im))
	r.re.Add(s.Mul(x.im, y.re), s.Mul(x.re, y.im))
	*z = r
	return z
}

// ---------------------------------------------------------------- stringValue

type stringValue interface {
	cat(v stringValue) stringValue
	eq(stringValue) bool
	id() int
	len() int
	s() []byte
}

type stringID int

func (s stringID) id() int   { return int(s) }
func (s stringID) len() int  { return len(dict.S(int(s))) }
func (s stringID) s() []byte { return dict.S(int(s)) }

func (s stringID) cat(v stringValue) stringValue {
	switch x := v.(type) {
	case stringID:
		if s == 0 {
			return v
		}

		if x == 0 {
			return s
		}

		return stringIDs{s, x}
	case stringIDs:
		if s == 0 {
			return v
		}

		return append(stringIDs{s}, x...)
	default:
		panic("internal error")
	}
}

func (s stringID) eq(v stringValue) bool {
	switch x := v.(type) {
	case stringID:
		return s == x
	case stringIDs:
		panic("TODO")
	default:
		panic("internal error")
	}
}

type stringIDs []stringID

func (s stringIDs) eq(v stringValue) bool { panic("TODO") }
func (s stringIDs) id() int               { return dict.ID(s.s()) }

func (s stringIDs) cat(v stringValue) stringValue {
	switch x := v.(type) {
	case stringID:
		if x == 0 {
			return s
		}

		return append(s, x)
	case stringIDs:
		return append(s, x...)
	default:
		panic("internal error")
	}
}

func (s stringIDs) len() int {
	n := 0
	for _, v := range s {
		n += v.len()
	}
	return n
}

func (s stringIDs) s() []byte {
	b := make([]byte, 0, s.len())
	for _, v := range s {
		b = append(b, v.s()...)
	}
	return b
}

func bigFloatString(a *big.Float) string {
	var m2 big.Float
	e2 := a.MantExp(&m2)
	neg := m2.Sign() < 0
	m2.Abs(&m2)
	switch e10 := math.Floor(float64(e2) * math.Ln2 / math.Ln10); {
	case e10 >= 0 && e10 <= 100:
		if a.IsInt() {
			return fmt.Sprintf("%.0f", a)
		}

		fallthrough
	case e10 < 0 && e10 >= -1:
		return fmt.Sprintf("%.*g", digits, a)
	default:
		m2b, _ := m2.Float64()
		m10 := math.Exp(math.Log(m2b) + math.Ln2*float64(e2) - math.Ln10*e10)
		if neg {
			m10 = -m10
		}
		return fmt.Sprintf("%.*ge%+.0f", digits, m10, e10)
	}
}
