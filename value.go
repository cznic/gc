// Copyright 2016 The GC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gc

import (
	"fmt"
	"math"
	"math/big"
	"unicode"
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

	bigIntMinInt32 = big.NewInt(math.MinInt32)
	bigIntMaxInt32 = big.NewInt(math.MaxInt32)
	bigIntMinInt64 = big.NewInt(math.MinInt64)
	bigIntMaxInt64 = big.NewInt(math.MaxInt64)

	bigIntMaxUint32 = big.NewInt(math.MaxUint32)
	bigIntMaxUint64 = big.NewInt(0).SetUint64(math.MaxUint64)

	bigFloatMaxUint64 = big.NewFloat(0).SetPrec(DefaultFloatConstPrec).SetUint64(math.MaxUint64)
	bigFloatMinInt64  = big.NewFloat(0).SetPrec(DefaultFloatConstPrec).SetInt64(math.MinInt64)

	floatZero    = 0.
	floatNegZero = -1 / (1 / floatZero)
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
	ConstValue             // Value known at compile time.
	NilValue               // Typeless zero value.
	RuntimeValue           // Value known at run time.
	TypeValue              // Value represents a type.
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
	// AssignableTo reports whether this value is assignable to value u.
	// It panics if value's Kind is not NilValue, RuntimeValue, TypeValue
	// or ConstValue.
	//
	// See https://golang.org/ref/spec#Assignability
	AssignableTo(v Value) bool

	// BinaryMultiply implements the binary '*' operation and returns the
	// new Value, or nil if the operation is invalid. Invalid operations
	// are reported as errors at node position.
	BinaryMultiply(ctx *Context, node Node, op Value) Value

	// Const returns the value's constant value. It panics if value Kind is
	// not ConstValue.
	Const() Const

	// Convert converts this value to type t and returns the new Value. It
	// returns nil if the conversion is not possible.  It panics if value's
	// Kind is not NilValue, TypeValue or ConstValue.
	//
	// See https://golang.org/ref/spec#Conversions
	Convert(t Type) Value

	// Kind returns the specific kind of this value.
	Kind() ValueKind

	// Type returns the type of this value or nil if the value's type
	// cannot be determined due to errors or if the value is a NilValue.
	Type() Type

	// UnaryMinus implements the unary '-' operation and returns the new
	// Value, or nil if the operation is invalid. Invalid operations are
	// reported as errors at node position.
	UnaryMinus(ctx *Context, node Node) Value
}

type valueBase struct{ kind ValueKind }

func (v *valueBase) AssignableTo(Value) bool { panic("AssignableTo of inappropriate value") }
func (v *valueBase) Const() Const            { panic("Const of inappropriate value") }
func (v *valueBase) Convert(Type) Value      { panic("internal error") }
func (v *valueBase) Kind() ValueKind         { return v.kind }
func (v *valueBase) Type() Type              { return nil }

func (v *valueBase) BinaryMultiply(ctx *Context, n Node, op Value) Value {
	ctx.err(n, "invalid operand for binary *")
	return nil
}

func (v *valueBase) UnaryMinus(ctx *Context, n Node) Value {
	ctx.err(n, "invalid operand for unary -")
	return nil
}

// ----------------------------------------------------------------- constValue

type constValue struct {
	c Const
	valueBase
}

func newConstValue(c Const) *constValue { return &constValue{c, valueBase{ConstValue}} }

func (v *constValue) AssignableTo(val Value) bool { return v.c.AssignableTo(val) }
func (v *constValue) Const() Const                { return v.c }
func (v *constValue) String() string              { return v.c.String() }
func (v *constValue) Type() Type                  { return v.c.Type() }

func (v *constValue) Convert(u Type) Value {
	c := v.c.Convert(u)
	if c != nil {
		return newConstValue(c)
	}

	return nil
}

func (v *constValue) BinaryMultiply(ctx *Context, n Node, op Value) Value {
	return v.c.BinaryMultiply(ctx, n, op)
}

func (v *constValue) UnaryMinus(ctx *Context, n Node) Value {
	if c := v.c.UnaryMinus(ctx, n); c != nil {
		return newConstValue(c)
	}

	return nil
}

// ------------------------------------------------------------------- nilValue

type nilValue struct{ valueBase }

func newNilValue() *nilValue { return &nilValue{valueBase{NilValue}} }

func (v *nilValue) AssignableTo(val Value) bool {
	switch val.Type().Kind() {
	case Ptr, Func, Slice, Map, Chan, Interface:
		return true
	}
	return false
}

func (v *nilValue) Convert(u Type) Value {
	if u.Kind() == Ptr {
		return newRuntimeValue(u)
	}

	return nil
}

// --------------------------------------------------------------- runtimeValue

type runtimeValue struct {
	typ Type
	valueBase
}

func newRuntimeValue(typ Type) *runtimeValue { return &runtimeValue{typ, valueBase{RuntimeValue}} }

func (v *runtimeValue) AssignableTo(u Value) bool { return v.Type().AssignableTo(u.Type()) }
func (v *runtimeValue) Type() Type                { return v.typ }

func (v *runtimeValue) Convert(u Type) Value {
	t := v.Type()
	if t != nil && t.ConvertibleTo(u) {
		return newRuntimeValue(u)
	}

	return nil
}

func (v *runtimeValue) BinaryMultiply(ctx *Context, n Node, op Value) Value {
	todo(n) //TODO
	return nil
}

func (v *runtimeValue) UnaryMinus(ctx *Context, n Node) Value {
	switch t := v.Type(); {
	case t.IntegerType(), t.FloatingPointType(), t.ComplexType():
		return v
	default:
		return v.valueBase.UnaryMinus(ctx, n)
	}
}

// ------------------------------------------------------------------ typeValue

type typeValue struct {
	t Type
	valueBase
}

func newTypeValue(t Type) *typeValue { return &typeValue{t, valueBase{TypeValue}} }

func (v *typeValue) AssignableTo(val Value) bool { return v.Type().AssignableTo(val.Type()) }
func (v *typeValue) String() string              { return v.t.String() }
func (v *typeValue) Type() Type                  { return v.t }

func (v *typeValue) Convert(u Type) Value {
	if v.Type().ConvertibleTo(u) {
		return newTypeValue(u)
	}

	return nil
}

// Const is the representatopn of a constant Go expression.
//
// Not all methods apply to all kinds of constants.  Restrictions, if any, are
// noted in the documentation for each method.  Use the Kind method to find out
// the kind of constant before calling kind-specific methods.  Calling a method
// inappropriate to the kind of constant causes a run-time panic.
type Const interface {
	normalize() Const         // Keeps exact value, can return nil on overflow.
	representable(Type) Const // Can be inexact for floats.

	// AssignableTo reports whether this constant is assignable to value v.
	// It panics if v is a ConstValue.
	//
	// See https://golang.org/ref/spec#Assignability
	AssignableTo(v Value) bool

	// Convert converts this constant to type t and returns the new
	// Constant. It returns nil if the conversion is not possible.
	//
	// See https://golang.org/ref/spec#Conversions
	Convert(t Type) Const

	// Kind returns the specific kind of the constant.
	Kind() ConstKind

	// Numeric reports whether the constant's Kind is one of RuneConst,
	// IntConst, FloatingPointConst or ComplexConst.
	Numeric() bool

	// String returns a string representation of the constant's value.
	String() string

	// Type returns the constant's type.
	Type() Type

	// BinaryMultiply implements the binary '*' operation and returns the
	// new Value, or nil if the operation is invalid. Invalid operations
	// are reported as errors at node position.
	BinaryMultiply(ctx *Context, node Node, op Value) Value

	// UnaryMinus implements the unary '-' operation and returns the new
	// Const, or nil if the operation is invalid. Invalid operations are
	// reported as errors at node position if ctx is not nil.
	UnaryMinus(ctx *Context, node Node) Const

	// Untyped reports whether the constant is untyped.
	Untyped() bool
}

type constBase struct {
	kind    ConstKind
	typ     Type
	untyped bool
}

func (c *constBase) AssignableTo(Value) bool  { panic("internal error") }
func (c *constBase) Convert(Type) Const       { panic("internal error") }
func (c *constBase) Kind() ConstKind          { return c.kind }
func (c *constBase) String() string           { panic("internal error") }
func (c *constBase) Type() Type               { return c.typ }
func (c *constBase) Untyped() bool            { return c.untyped }
func (c *constBase) normalize() Const         { panic("internal error") }
func (c *constBase) representable(Type) Const { panic("internal error") }

func (c *constBase) assignableTo(cc Const, v Value) bool {
	if v.Kind() == ConstValue {
		panic("internal error")
	}

	t := v.Type()

	if c.Untyped() && cc.representable(t) != nil {
		return true
	}

	return c.Type().AssignableTo(t)
}

func (c *constBase) Numeric() bool {
	return c.kind == ComplexConst || c.kind == FloatingPointConst ||
		c.kind == IntConst || c.kind == RuneConst
}

func (c *constBase) representableComplexFromBigComplex(v *bigComplex, t Type) Const {
	switch t.Kind() {
	case Complex64, Complex128:
		ctx := t.context()
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

func (c *constBase) representableComplexFromBigFloat(v *big.Float, t Type) Const {
	return c.representableComplexFromBigComplex(&bigComplex{v, big.NewFloat(0)}, t)
}

func (c *constBase) representableComplexFromBigInt(v *big.Int, t Type) Const {
	return c.representableComplexFromBigFloat(big.NewFloat(0).SetPrec(t.context().floatConstPrec).SetInt(v), t)
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
	case Float32, Float64:
		f, _ := v.Float64()
		return c.representableFloatFromFloat(f, t)
	}
	panic("internal error")
}

func (c *constBase) representableFloatFromBigInt(v *big.Int, t Type) Const {
	return c.representableFloatFromBigFloat(big.NewFloat(0).SetPrec(t.context().floatConstPrec).SetInt(v), t)
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
		if math.Abs(v) <= math.MaxFloat32 {
			return newFloatConst(v, nil, t, false)
		}
	case Float64:
		return newFloatConst(v, nil, t, false)
	default:
		panic("internal error")
	}
	return nil
}

func (c *constBase) representableFloatFromInt(v int64, t Type) Const {
	return c.representableFloatFromFloat(float64(v), t)
}

func (c *constBase) representableIntFromBigComplex(v *bigComplex, t Type) Const {
	if v.im.Sign() != 0 {
		return nil
	}

	return c.representableIntFromBigFloat(v.re, t)
}

func (c *constBase) representableIntFromBigFloat(v *big.Float, t Type) Const {
	if !v.IsInt() {
		return nil
	}

	if v.Cmp(bigFloatMinInt64) < 0 || v.Cmp(bigFloatMaxUint64) > 0 {
		return nil
	}

	n, _ := v.Int(nil)
	return c.representableIntFromBigInt(n, t)
}

func (c *constBase) representableIntFromBigInt(v *big.Int, t Type) Const {
	if v.Cmp(bigIntMinInt64) < 0 || v.Cmp(bigIntMaxUint64) > 0 {
		return nil
	}

	if v.Cmp(bigIntMaxInt64) <= 0 {
		return c.representableIntFromInt(v.Int64(), t)
	}

	// v in (math.MaxInt64, math.MaxUint64].
	switch t.Kind() {
	case Int8, Int16, Int32, Int64, Int, Uint8, Uint16, Uint32:
		return nil
	case Uint64:
		return newIntConst(0, v, t, false)
	case Uint:
		if t.context().model.IntBytes == 8 {
			return newIntConst(0, v, t, false)
		}
	case Uintptr:
		if t.context().model.PtrBytes == 8 {
			return newIntConst(0, v, t, false)
		}
	default:
		panic("internal error")
	}
	return nil
}

func (c *constBase) representableIntFromComplex(v complex128, t Type) Const {
	if imag(v) != 0 {
		return nil
	}

	return c.representableIntFromFloat(real(v), t)
}

func (c *constBase) representableIntFromFloat(v float64, t Type) Const {
	if math.Floor(v) != v {
		return nil
	}

	if v < math.MinInt64 || v > math.MaxInt64 {
		panic("internal error")
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
		ctx := t.context()
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
		ctx := t.context()
		if n >= 0 && uint64(n) <= ctx.maxUint {
			return newIntConst(n, nil, t, false)
		}
	case Uintptr:
		ctx := t.context()
		if n >= 0 && uint64(n) <= ctx.maxUintptr {
			return newIntConst(n, nil, t, false)
		}
	default:
		panic("internal error")
	}
	return nil
}

func (c *constBase) representableIntFromInt(v int64, t Type) Const {
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
		ctx := t.context()
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
		ctx := t.context()
		if v >= 0 && uint64(v) <= ctx.maxUint {
			return newIntConst(v, nil, t, false)
		}
	case Uintptr:
		ctx := t.context()
		if v >= 0 && uint64(v) <= ctx.maxUintptr {
			return newIntConst(v, nil, t, false)
		}
	default:
		panic("internal error")
	}
	return nil
}

func (c *constBase) BinaryMultiply(ctx *Context, n Node, op Value) Value {
	ctx.err(n, "invalid operand for binary *")
	return nil
}

func (c *constBase) UnaryMinus(ctx *Context, n Node) Const {
	ctx.err(n, "invalid operand for unary -")
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

func (c *boolConst) AssignableTo(v Value) bool { return c.assignableTo(c, v) }
func (c *boolConst) normalize() Const          { return c }

func (c *boolConst) representable(t Type) Const {
	if t.Kind() == Bool {
		return c
	}

	return nil
}

func (c *boolConst) Convert(t Type) Const {
	if t.Kind() != Bool {
		return nil
	}

	return newBoolConst(c.val, t, false)
}

func (c *boolConst) String() string { return fmt.Sprint(c.val) }

// --------------------------------------------------------------- complexConst

type complexConst struct {
	constBase
	val    complex128
	bigVal *bigComplex
}

func newComplexConst(val complex128, bigVal *bigComplex, typ Type, untyped bool) *complexConst {
	return &complexConst{constBase{ComplexConst, typ, untyped}, val, bigVal}
}

func (c *complexConst) AssignableTo(v Value) bool { return c.assignableTo(c, v) }
func (c *complexConst) Convert(t Type) (r Const)  { return c.representable(t) }

func (c *complexConst) normalize() Const {
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

	return c.representable(c.Type())
}

func (c *complexConst) representable(t Type) Const {
	switch {
	case t.IntegerType():
		if v := c.bigVal; v != nil {
			return c.representableIntFromBigComplex(v, t)
		}

		return c.representableIntFromComplex(c.val, t)
	case t.FloatingPointType():
		if v := c.bigVal; v != nil {
			return c.representableFloatFromBigComplex(v, t)
		}

		return c.representableFloatFromComplex(c.val, t)
	case t.ComplexType():
		if v := c.bigVal; v != nil {
			return c.representableComplexFromBigComplex(v, t)
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

func (c *complexConst) BinaryMultiply(ctx *Context, n Node, op Value) Value {
	todo(n) //TODO
	return nil
}

func (c *complexConst) UnaryMinus(ctx *Context, n Node) Const {
	todo(n)
	return nil
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

func (c *floatConst) AssignableTo(v Value) bool { return c.assignableTo(c, v) }
func (c *floatConst) Convert(t Type) Const      { return c.representable(t) }

func (c *floatConst) normalize() Const {
	if v := c.bigVal; v != nil {
		if f, ok := v.Float64(); ok == big.Exact {
			c = newFloatConst(f, nil, c.Type(), c.Untyped())
		}
	}

	if c.Untyped() {
		return c
	}

	return c.representable(c.Type())
}

func (c *floatConst) representable(t Type) Const {
	switch {
	case t.IntegerType():
		if v := c.bigVal; v != nil {
			return c.representableIntFromBigFloat(v, t)
		}

		return c.representableIntFromFloat(c.val, t)
	case t.FloatingPointType():
		if v := c.bigVal; v != nil {
			return c.representableFloatFromBigFloat(v, t)
		}

		return c.representableFloatFromFloat(c.val, t)
	case t.ComplexType():
		if v := c.bigVal; v != nil {
			return c.representableComplexFromBigFloat(v, t)
		}

		return c.representableComplexFromFloat(c.val, t)
	}
	return nil
}

func (c *floatConst) String() string {
	if c.bigVal != nil {
		return bigFloatString(c.bigVal)
	}

	return fmt.Sprint(c.val)
}

func (c *floatConst) BinaryMultiply(ctx *Context, n Node, op Value) Value {
	todo(n) //TODO
	return nil
}

func (c *floatConst) UnaryMinus(ctx *Context, n Node) Const {
	todo(n)
	return nil
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

func (c *intConst) AssignableTo(v Value) bool { return c.assignableTo(c, v) }
func (c *intConst) Convert(t Type) Const      { return c.representable(t) }

func (c *intConst) big() *big.Int {
	if c.bigVal != nil {
		return c.bigVal
	}

	return big.NewInt(c.val)
}

func (c *intConst) normalize() Const {
	if v := c.bigVal; v != nil {
		if uint64(v.BitLen()) > uint64(c.Type().context().intConstBits) {
			return nil
		}

		if v.Cmp(bigIntMinInt64) >= 0 && v.Cmp(bigIntMaxInt64) <= 0 {
			c = newIntConst(v.Int64(), nil, c.Type(), c.Untyped())
		}
	}

	if c.Untyped() {
		return c
	}

	return c.representable(c.Type())
}

func (c *intConst) representable(t Type) Const {
	switch {
	case t.IntegerType():
		if v := c.bigVal; v != nil {
			return c.representableIntFromBigInt(v, t)
		}

		return c.representableIntFromInt(c.val, t)
	case t.FloatingPointType():
		if v := c.bigVal; v != nil {
			return c.representableFloatFromBigInt(v, t)
		}

		return c.representableFloatFromInt(c.val, t)
	case t.ComplexType():
		if v := c.bigVal; v != nil {
			return c.representableComplexFromBigInt(v, t)
		}

		return c.representableComplexFromInt(c.val, t)
	case t.Kind() == String:
		if d := c.representable(t.context().int32Type); d != nil {
			s := string(d.(*intConst).val)
			return newStringConst(stringID(dict.SID(s)), t, false)
		}
	}
	return nil
}

func (c *intConst) String() string {
	if c.bigVal != nil {
		return fmt.Sprint(c.bigVal)
	}

	return fmt.Sprint(c.val)
}

func (c *intConst) BinaryMultiply(ctx *Context, n Node, op Value) Value {
	switch op.Kind() {
	case ConstValue:
		a, b := ctx.arithmeticBinOpShape(c, op.Const(), n)
		if a != nil {
			var c big.Int
			c.Mul(a.(*intConst).big(), b.(*intConst).big())
			d := newIntConst(0, &c, a.Type(), a.Untyped())
			if e := d.normalize(); e != nil {
				return newConstValue(e)
			}

			ctx.err(n, "constant %s overflows %s", d, a.Type())
		}
	default:
		//dbg("", op.Kind())
		todo(n) //TODO
	}
	return nil
}

func (c *intConst) UnaryMinus(ctx *Context, n Node) Const {
	if c.bigVal != nil {
		var v big.Int
		return newIntConst(0, v.Neg(v.Set(c.bigVal)), c.Type(), c.Untyped()).normalize()
	}

	switch {
	case c.val == math.MinInt64:
		var v big.Int
		return newIntConst(0, v.Neg(v.SetInt64(c.val)), c.Type(), c.Untyped()).normalize()
	default:
		return newIntConst(-c.val, nil, c.Type(), c.Untyped()).normalize()
	}
}

// ------------------------------------------------------------------ runeConst

type runeConst struct {
	intConst
}

func newRuneConst(val rune, bigVal *big.Int, typ Type, untyped bool) *runeConst {
	return &runeConst{intConst{constBase{RuneConst, typ, untyped}, int64(val), bigVal}}
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

func (c *runeConst) BinaryMultiply(ctx *Context, n Node, op Value) Value {
	todo(n) //TODO
	return nil
}

func (c *runeConst) UnaryMinus(ctx *Context, n Node) Const {
	if d := c.intConst.UnaryMinus(ctx, n); d != nil {
		return &runeConst{*d.(*intConst)}
	}
	return nil
}

// ---------------------------------------------------------------- stringConst

type stringConst struct {
	constBase
	val stringValue
}

func newStringConst(val stringValue, typ Type, untyped bool) *stringConst {
	return &stringConst{constBase{StringConst, typ, untyped}, val}
}

func (c *stringConst) AssignableTo(v Value) bool { return c.assignableTo(c, v) }
func (c *stringConst) Convert(t Type) Const      { return c.representable(t) }
func (c *stringConst) normalize() Const          { return c }

func (c *stringConst) representable(t Type) Const {
	if t.Kind() != String {
		return nil
	}

	return newStringConst(c.val, t, false)
}

func (c *stringConst) String() string { return fmt.Sprintf("%q", c.val.s()) }

// ----------------------------------------------------------------- bigComplex

// bigComplex is the value of a untyped complex constant.
type bigComplex struct {
	re, im *big.Float
}

func (c *bigComplex) String() string {
	if c.im.Sign() >= 0 {
		return fmt.Sprintf("%s+%si", bigFloatString(c.re), bigFloatString(c.im))
	}

	return fmt.Sprintf("%s%si", bigFloatString(c.re), bigFloatString(c.im))
}

// ---------------------------------------------------------------- stringValue

type stringValue interface {
	cat(v stringValue) stringValue
	len() int
	s() []byte
}

type stringID int

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

type stringIDs []stringID

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
	const digits = 6
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
