// Copyright 2016 The GC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gc

import (
	"fmt"
	"math"
	"math/big"
	"unicode"

	"github.com/cznic/xc"
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
	PackageValue           // Value is an imported package qualifier.
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
	// Add implements the binary addition operation and returns the new
	// Value, or nil if the operation is invalid. Invalid operations are
	// reported as errors at node position.
	Add(node Node, op Value) Value

	// Addressable reports whether a value is addressable. It panics if
	// value's Kind is not RuntimeValue.
	Addressable() bool

	// And implements the binary bitwise and operation and returns the new
	// Value, or nil if the operation is invalid. Invalid operations are
	// reported as errors at node position.
	And(node Node, op Value) Value

	// AssignableTo reports whether this value is assignable to type t.
	// It panics if value's Kind is not NilValue, RuntimeValue, TypeValue
	// or ConstValue.
	//
	// See https://golang.org/ref/spec#Assignability
	AssignableTo(t Type) bool

	// Const returns the value's constant value. It panics if value Kind is
	// not ConstValue.
	Const() Const

	// Convert converts this value to type t and returns the new Value. It
	// returns nil if the conversion is not possible.  It panics if value's
	// Kind is not NilValue, TypeValue or ConstValue.
	//
	// See https://golang.org/ref/spec#Conversions
	Convert(t Type) Value

	// Cpl implements the unary bitwise complement operation and returns
	// the new Value, or nil if the operation is invalid. Invalid
	// operations are reported as errors at node position.
	Cpl(ctx *Context, node Node) Value

	// Div implements the binary divide operation and returns the new
	// Value, or nil if the operation is invalid. Invalid operations are
	// reported as errors at node position.
	Div(node Node, op Value) Value

	// Kind returns the specific kind of this value.
	Kind() ValueKind

	// Lsh implements the binary left shift operation and returns the new
	// Value, or nil if the operation is invalid. Invalid operations are
	// reported as errors at node position.
	Lsh(node Node, op Value) Value

	// Mod implements the binary modulo operation and returns the new
	// Value, or nil if the operation is invalid. Invalid operations are
	// reported as errors at node position.
	Mod(node Node, op Value) Value

	// Mul implements the binary multiply operation and returns the new
	// Value, or nil if the operation is invalid. Invalid operations are
	// reported as errors at node position.
	Mul(node Node, op Value) Value

	// Neg implements the unary negate operation and returns the new Value,
	// or nil if the operation is invalid. Invalid operations are reported
	// as errors at node position.
	Neg(ctx *Context, node Node) Value

	// Package returns the package a PackageValue refers to. It panic if
	// the value Kind is not PackageValue.
	Package() *Package

	// PredeclaredFunction returns a declaration of a predeclared function
	// or nil if the value does not represent a predeclared function. It
	// panics if value's Kind is not RuntimeValue or if the value's type is
	// not a function.
	PredeclaredFunction() Declaration

	// Rsh implements the binary right shift operation and returns the new
	// Value, or nil if the operation is invalid. Invalid operations are
	// reported as errors at node position.
	Rsh(node Node, op Value) Value

	// StructField returns a *StructField and the fields cumulative offset
	// if the value represents a struct field or (nil, 0) otherwise.  It
	// panics if the value's Kind is not RuntimeValue. The result pointee,
	// if any, is read only.
	StructField() (*StructField, uint64)

	// Sub implements the binary subtraction operation and returns the new
	// Value, or nil if the operation is invalid. Invalid operations are
	// reported as errors at node position.
	Sub(node Node, op Value) Value

	// Type returns the type of this value or nil if the value's type
	// cannot be determined due to errors or if the value is a NilValue.
	Type() Type
}

type valueBase struct{ kind ValueKind }

func (v *valueBase) Addressable() bool                   { panic("Addressable of inappropriate value") }
func (v *valueBase) AssignableTo(Type) bool              { panic("AssignableTo of inappropriate value") }
func (v *valueBase) Const() Const                        { panic("Const of inappropriate value") }
func (v *valueBase) Convert(Type) Value                  { panic("internal error") }
func (v *valueBase) Kind() ValueKind                     { return v.kind }
func (v *valueBase) Package() *Package                   { panic("Package of inappropriate value") }
func (v *valueBase) StructField() (*StructField, uint64) { panic("StructField of inappropriate value") }
func (v *valueBase) Type() Type                          { return nil }

func (v *valueBase) Add(n Node, op Value) Value {
	op.Type().context().err(n, "invalid operand for binary +")
	return nil
}

func (v *valueBase) And(n Node, op Value) Value {
	op.Type().context().err(n, "invalid operand for binary &")
	return nil
}

func (v *valueBase) Cpl(ctx *Context, n Node) Value {
	ctx.err(n, "invalid operand for unary ^")
	return nil
}

func (v *valueBase) Div(n Node, op Value) Value {
	op.Type().context().err(n, "invalid operand for binary /")
	return nil
}

func (v *valueBase) Lsh(n Node, op Value) Value {
	op.Type().context().err(n, "invalid operand for binary <<")
	return nil
}

func (v *valueBase) Mod(n Node, op Value) Value {
	op.Type().context().err(n, "invalid operand for binary %")
	return nil
}

func (v *valueBase) Mul(n Node, op Value) Value {
	op.Type().context().err(n, "invalid operand for binary *")
	return nil
}

func (v *valueBase) Neg(ctx *Context, n Node) Value {
	ctx.err(n, "invalid operand for unary -")
	return nil
}

func (v *valueBase) PredeclaredFunction() Declaration {
	panic("PredeclaredFunction of inappropriate value")
}

func (v *valueBase) Rsh(n Node, op Value) Value {
	op.Type().context().err(n, "invalid operand for binary >>")
	return nil
}

func (v *valueBase) Sub(n Node, op Value) Value {
	op.Type().context().err(n, "invalid operand for binary -")
	return nil
}

// ----------------------------------------------------------------- constValue

type constValue struct {
	c Const
	valueBase
}

func newConstValue(c Const) *constValue { return &constValue{c, valueBase{ConstValue}} }

func (v *constValue) Add(n Node, op Value) Value     { return v.c.Add(n, op) }
func (v *constValue) And(n Node, op Value) Value     { return v.c.And(n, op) }
func (v *constValue) AssignableTo(t Type) bool       { return v.c.AssignableTo(t) }
func (v *constValue) Const() Const                   { return v.c }
func (v *constValue) Convert(u Type) Value           { return v.c.Convert(u) }
func (v *constValue) Cpl(ctx *Context, n Node) Value { return v.c.Cpl(ctx, n) }
func (v *constValue) Div(n Node, op Value) Value     { return v.c.Div(n, op) }
func (v *constValue) Lsh(n Node, op Value) Value     { return v.c.Lsh(n, op) }
func (v *constValue) Mod(n Node, op Value) Value     { return v.c.Mod(n, op) }
func (v *constValue) Mul(n Node, op Value) Value     { return v.c.Mul(n, op) }
func (v *constValue) Neg(ctx *Context, n Node) Value { return v.c.Neg(ctx, n) }
func (v *constValue) Rsh(n Node, op Value) Value     { return v.c.Rsh(n, op) }
func (v *constValue) String() string                 { return v.c.String() }
func (v *constValue) Sub(n Node, op Value) Value     { return v.c.Sub(n, op) }
func (v *constValue) Type() Type                     { return v.c.Type() }

// ------------------------------------------------------------------- nilValue

type nilValue struct{ valueBase }

func newNilValue() *nilValue { return &nilValue{valueBase{NilValue}} }

func (v *nilValue) AssignableTo(t Type) bool {
	switch t.Kind() {
	case Ptr, Func, Slice, Map, Chan, Interface:
		return true
	}

	return false
}

func (v *nilValue) Convert(u Type) Value {
	switch u.Kind() {
	case Ptr, Func, Slice, Map, Chan, Interface:
		return newRuntimeValue(u)
	}

	return nil
}

// --------------------------------------------------------------- packageValue

type packageValue struct {
	pkg *Package
	valueBase
}

func newPackageValue(pkg *Package) *packageValue { return &packageValue{pkg, valueBase{PackageValue}} }

func (v *packageValue) Package() *Package { return v.pkg }

// --------------------------------------------------------------- runtimeValue

type runtimeValue struct {
	valueBase
	d           Declaration
	field       *StructField
	offset      uint64
	typ         Type
	addressable bool
}

func newRuntimeValue(typ Type) *runtimeValue {
	return &runtimeValue{
		typ:       typ,
		valueBase: valueBase{RuntimeValue},
	}
}

func newAddressableValue(typ Type) *runtimeValue {
	return &runtimeValue{
		addressable: true,
		typ:         typ,
		valueBase:   valueBase{RuntimeValue},
	}
}

func newPredeclaredFunctionValue(d Declaration) *runtimeValue {
	return &runtimeValue{
		d:         d,
		typ:       d.(*FuncDeclaration).Type,
		valueBase: valueBase{RuntimeValue},
	}
}

func newStructFieldalue(f *StructField, offset uint64) *runtimeValue {
	return &runtimeValue{
		addressable: true,
		field:       f,
		offset:      offset,
		typ:         f.Type,
		valueBase:   valueBase{RuntimeValue},
	}
}

func (v *runtimeValue) Addressable() bool                   { return v.addressable }
func (v *runtimeValue) AssignableTo(t Type) bool            { return v.Type().AssignableTo(t) }
func (v *runtimeValue) PredeclaredFunction() Declaration    { return v.d }
func (v *runtimeValue) StructField() (*StructField, uint64) { return v.field, v.offset }
func (v *runtimeValue) Type() Type                          { return v.typ }

func (v *runtimeValue) Add(n Node, op Value) Value {
	todo(n)
	return nil
}

func (v *runtimeValue) And(n Node, op Value) Value {
	todo(n)
	return nil
}

func (v *runtimeValue) Convert(u Type) Value {
	if t := v.Type(); t != nil && t.ConvertibleTo(u) {
		return newRuntimeValue(u)
	}

	return nil
}

func (v *runtimeValue) Cpl(ctx *Context, n Node) Value {
	switch t := v.Type(); {
	case t.IntegerType():
		return v
	default:
		return nil
	}
}

func (v *runtimeValue) Div(n Node, op Value) Value {
	todo(n)
	return nil
}

func (v *runtimeValue) Lsh(n Node, op Value) Value {
	todo(n)
	return nil
}

func (v *runtimeValue) Mod(n Node, op Value) Value {
	todo(n)
	return nil
}

func (v *runtimeValue) Mul(n Node, op Value) Value {
	todo(n)
	return nil
}

func (v *runtimeValue) Neg(ctx *Context, n Node) Value {
	switch t := v.Type(); {
	case t.IntegerType(), t.FloatingPointType(), t.ComplexType():
		return v
	default:
		return v.valueBase.Neg(ctx, n)
	}
}

func (v *runtimeValue) Rsh(n Node, op Value) Value {
	todo(n)
	return nil
}

func (v *runtimeValue) Sub(n Node, op Value) Value {
	todo(n)
	return nil
}

// ------------------------------------------------------------------ typeValue

type typeValue struct {
	t Type
	valueBase
}

func newTypeValue(t Type) *typeValue { return &typeValue{t, valueBase{TypeValue}} }

func (v *typeValue) AssignableTo(t Type) bool { return v.Type().AssignableTo(t) }
func (v *typeValue) String() string           { return v.t.String() }
func (v *typeValue) Type() Type               { return v.t }

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
	add(n Node, t Type, untyped bool, op Const) Const
	and(n Node, t Type, untyped bool, op Const) Const
	convert(Type) Const // Result is untyped.
	div(n Node, t Type, untyped bool, op Const) Const
	mod(n Node, t Type, untyped bool, op Const) Const
	mul(n Node, t Type, untyped bool, op Const) Const
	mustConvert(Node, Type) Const // Result is untyped.
	normalize() Const             // Keeps exact value or returns nil on overflow.
	representable(Type) Const     // Can be inexact for floats.
	sub(n Node, t Type, untyped bool, op Const) Const

	// Add implements the binary addition operation and returns the new
	// Value, or nil if the operation is invalid. Invalid operations are
	// reported as errors at node position.
	Add(node Node, op Value) Value

	// And implements the binary bitwise and operation and returns the new
	// Value, or nil if the operation is invalid. Invalid operations are
	// reported as errors at node position.
	And(node Node, op Value) Value

	// AssignableTo reports whether this constant is assignable to type t.
	//
	// See https://golang.org/ref/spec#Assignability
	AssignableTo(t Type) bool

	// Convert converts this constant to type t and returns the new Value.
	// It returns nil if the conversion is not possible.
	//
	// See https://golang.org/ref/spec#Conversions
	Convert(t Type) Value

	// Cpl implements the unary bitwise complement operation and returns
	// the new Value, or nil if the operation is invalid. Invalid
	// operations are reported as errors at node position.
	Cpl(ctx *Context, node Node) Value

	// Div implements the binary divide operation and returns the new
	// Value, or nil if the operation is invalid. Invalid operations are
	// reported as errors at node position.
	Div(n Node, op Value) Value

	// Integral reports whether the constant's value is ∈ Z. It panics if
	// the constant is not numeric.
	Integral() bool

	// Kind returns the specific kind of the constant.
	Kind() ConstKind

	// Lsh implements the binary left shift operation and returns the new
	// Value, or nil if the operation is invalid. Invalid operations are
	// reported as errors at node position.
	Lsh(node Node, op Value) Value

	// Mod implements the binary modulo operation and returns the new
	// Value, or nil if the operation is invalid. Invalid operations are
	// reported as errors at node position.
	Mod(node Node, op Value) Value

	// Mul implements the binary multiply operation and returns the new
	// Value, or nil if the operation is invalid. Invalid operations are
	// reported as errors at node position.
	Mul(n Node, op Value) Value

	// Neg implements the unary negate operation and returns the new Const,
	// or nil if the operation is invalid. Invalid operations are reported
	// as errors at node position if ctx is not nil.
	Neg(ctx *Context, node Node) Value

	// Numeric reports whether the constant's Kind is one of RuneConst,
	// IntConst, FloatingPointConst or ComplexConst.
	Numeric() bool

	// Rsh implements the binary right shift operation and returns the new
	// Value, or nil if the operation is invalid. Invalid operations are
	// reported as errors at node position.
	Rsh(node Node, op Value) Value

	// String returns a string representation of the constant's value.
	String() string

	// Sub implements the binary subtraction operation and returns the new
	// Value, or nil if the operation is invalid. Invalid operations are
	// reported as errors at node position.
	Sub(node Node, op Value) Value

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

func (c *constBase) add(Node, Type, bool, Const) Const { panic("internal error") }
func (c *constBase) and(Node, Type, bool, Const) Const { panic("internal error") }
func (c *constBase) AssignableTo(Type) bool            { panic("internal error") }
func (c *constBase) convert(Type) Const                { panic("internal error") }
func (c *constBase) Convert(Type) Value                { panic("internal error") }
func (c *constBase) div(Node, Type, bool, Const) Const { panic("internal error") }
func (c *constBase) Integral() bool                    { panic("internal error") }
func (c *constBase) Kind() ConstKind                   { return c.kind }
func (c *constBase) mod(Node, Type, bool, Const) Const { panic("internal error") }
func (c *constBase) mul(Node, Type, bool, Const) Const { panic("internal error") }
func (c *constBase) mustConvert(Node, Type) Const      { panic("internal error") }
func (c *constBase) normalize() Const                  { panic("internal error") }
func (c *constBase) representable(Type) Const          { panic("internal error") }
func (c *constBase) String() string                    { panic("internal error") }
func (c *constBase) sub(Node, Type, bool, Const) Const { panic("internal error") }
func (c *constBase) Type() Type                        { return c.typ }
func (c *constBase) Untyped() bool                     { return c.untyped }

func (c *constBase) mustConvertConst(n Node, d Const, t Type) Const {
	e := d.convert(t)
	if e.representable(t) != nil {
		return e
	}

	t.context().constConversionFail(n, t, d)
	return nil
}

func (c *constBase) assignableTo(cc Const, t Type) bool {
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
		return c.representableIntFromBigFloat(big.NewFloat(0).SetPrec(t.context().floatConstPrec).SetFloat64(v), t)
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

func (c *constBase) Add(n Node, op Value) Value {
	op.Type().context().err(n, "invalid operand for binary +")
	return nil
}

func (c *constBase) And(n Node, op Value) Value {
	op.Type().context().err(n, "invalid operand for binary &")
	return nil
}

func (c *constBase) Cpl(ctx *Context, n Node) Value {
	ctx.err(n, "invalid operand for unary ^")
	return nil
}

func (c *constBase) Div(n Node, op Value) Value {
	op.Type().context().err(n, "invalid operand for binary /")
	return nil
}

func (c *constBase) Lsh(n Node, op Value) Value {
	op.Type().context().err(n, "invalid operand for binary *")
	return nil
}

func (c *constBase) Mod(n Node, op Value) Value {
	op.Type().context().err(n, "invalid operand for binary %")
	return nil
}

func (c *constBase) Mul(n Node, op Value) Value {
	op.Type().context().err(n, "invalid operand for binary *")
	return nil
}

func (c *constBase) Neg(ctx *Context, n Node) Value {
	ctx.err(n, "invalid operand for unary -")
	return nil
}

func (c *constBase) Rsh(n Node, op Value) Value {
	op.Type().context().err(n, "invalid operand for binary >>")
	return nil
}

func (c *constBase) Sub(n Node, op Value) Value {
	op.Type().context().err(n, "invalid operand for binary -")
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

func (c *boolConst) AssignableTo(t Type) bool { return c.assignableTo(c, t) }
func (c *boolConst) normalize() Const         { return c }

func (c *boolConst) representable(t Type) Const {
	if t.Kind() == Bool {
		return c
	}

	return nil
}

func (c *boolConst) Convert(t Type) Value {
	if t.Kind() == Bool {
		return newConstValue(newBoolConst(c.val, t, false))
	}

	return nil
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

func (c *complexConst) AssignableTo(t Type) bool         { return c.assignableTo(c, t) }
func (c *complexConst) mustConvert(n Node, t Type) Const { return c.mustConvertConst(n, c, t) }

func (c *complexConst) Add(n Node, op Value) Value {
	todo(n)
	return nil
}

func (c *complexConst) Convert(t Type) Value {
	if c.Type().ConvertibleTo(t) {
		if d := c.representable(t); d != nil {
			return newConstValue(d)
		}
	}

	return nil
}

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

func (c *complexConst) convert(t Type) Const {
	todo(xc.Token{}) //TODO
	return nil
}

func (c *complexConst) div(n Node, t Type, untyped bool, op Const) Const {
	todo(n)
	return nil
}

func (c *complexConst) Div(n Node, op Value) Value {
	todo(n)
	return nil
}

func (c *complexConst) mul(n Node, t Type, untyped bool, op Const) Const {
	todo(n)
	return nil
}

func (c *complexConst) Mul(n Node, op Value) Value {
	todo(n)
	return nil
}

func (c *complexConst) Neg(ctx *Context, n Node) Value {
	todo(n)
	return nil
}

func (c *complexConst) Sub(n Node, op Value) Value {
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

func (c *floatConst) AssignableTo(t Type) bool         { return c.assignableTo(c, t) }
func (c *floatConst) mustConvert(n Node, t Type) Const { return c.mustConvertConst(n, c, t) }

func (c *floatConst) Add(n Node, op Value) Value {
	todo(n)
	return nil
}

func (c *floatConst) Convert(t Type) Value {
	if c.Type().ConvertibleTo(t) {
		if d := c.representable(t); d != nil {
			return newConstValue(d)
		}
	}

	return nil
}

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

func (c *floatConst) convert(t Type) Const {
	ctx := t.context()
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
		todo(xc.Token{})
	default:
		panic("internal error")
	}
	return nil
}

func (c *floatConst) div(n Node, t Type, untyped bool, op Const) Const {
	ctx := t.context()
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

func (c *floatConst) Div(n Node, op Value) Value {
	ctx := op.Type().context()
	switch op.Kind() {
	case ConstValue:
		t, untyped, a, b := ctx.arithmeticBinOpShape(c, op.Const(), n)
		if t != nil {
			if d := a.div(n, t, untyped, b); d != nil {
				return newConstValue(d)
			}
		}
	default:
		//dbg("", op.Kind())
		todo(n)
	}
	return nil
}

func (c *floatConst) mod(n Node, t Type, untyped bool, op Const) Const {
	t.context().err(n, "illegal constant expression: floating-point %% operation")
	return nil
}

func (c *floatConst) mul(n Node, t Type, untyped bool, op Const) Const {
	ctx := t.context()
	var d big.Float
	d.Mul(c.bigVal, op.(*floatConst).bigVal)
	if untyped {
		t = nil
	}
	return ctx.mustConvertConst(n, t, newFloatConst(0, &d, ctx.float64Type, true))
}

func (c *floatConst) Lsh(n Node, op Value) Value {
	ctx := c.Type().context()
	if i := c.convert(ctx.intType); i != nil {
		return i.Lsh(n, op)
	}

	ctx.err(n, "invalid operand for binary <<")
	return nil
}

func (c *floatConst) Mul(n Node, op Value) Value {
	ctx := op.Type().context()
	switch op.Kind() {
	case ConstValue:
		t, untyped, a, b := ctx.arithmeticBinOpShape(c, op.Const(), n)
		if t != nil {
			if d := a.mul(n, t, untyped, b); d != nil {
				return newConstValue(d)
			}
		}
	default:
		//dbg("", op.Kind())
		todo(n)
	}
	return nil
}

func (c *floatConst) Neg(ctx *Context, n Node) Value {
	if c.bigVal != nil {
		var d big.Float
		return newConstValue(newFloatConst(0, d.Neg(c.bigVal), c.Type(), c.Untyped()))
	}

	return newConstValue(newFloatConst(-c.val, nil, c.Type(), c.Untyped()))
}

func (c *floatConst) Rsh(n Node, op Value) Value {
	todo(n)
	return nil
}

func (c *floatConst) Sub(n Node, op Value) Value {
	todo(n)
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

func (c *intConst) AssignableTo(t Type) bool         { return c.assignableTo(c, t) }
func (c *intConst) Integral() bool                   { return true }
func (c *intConst) mustConvert(n Node, t Type) Const { return c.mustConvertConst(n, c, t) }

func (c *intConst) add(n Node, t Type, untyped bool, op Const) Const {
	ctx := t.context()
	var d big.Int
	d.Add(c.bigVal, op.(*intConst).bigVal)
	if untyped {
		t = nil
	}
	return ctx.mustConvertConst(n, t, newIntConst(0, &d, ctx.intType, true))
}

func (c *intConst) Add(n Node, op Value) Value {
	ctx := op.Type().context()
	switch op.Kind() {
	case ConstValue:
		t, untyped, a, b := ctx.arithmeticBinOpShape(c, op.Const(), n)
		if t != nil {
			if d := a.add(n, t, untyped, b); d != nil {
				return newConstValue(d)
			}
		}
	default:
		//dbg("", op.Kind())
		todo(n)
	}
	return nil
}

func (c *intConst) and(n Node, t Type, untyped bool, op Const) Const {
	ctx := t.context()
	var d big.Int
	d.And(c.bigVal, op.(*intConst).bigVal)
	if untyped {
		t = nil
	}
	return ctx.mustConvertConst(n, t, newIntConst(0, &d, ctx.intType, true))
}

func (c *intConst) And(n Node, op Value) Value {
	ctx := op.Type().context()
	switch op.Kind() {
	case ConstValue:
		t, untyped, a, b := ctx.arithmeticBinOpShape(c, op.Const(), n)
		if t != nil {
			if d := a.and(n, t, untyped, b); d != nil {
				return newConstValue(d)
			}
		}
	default:
		//dbg("", op.Kind())
		todo(n)
	}
	return nil
}

func (c *intConst) Convert(t Type) Value {
	if c.Type().ConvertibleTo(t) {
		switch t.Kind() {
		case String:
			return newRuntimeValue(t)
		default:
			if d := c.representable(t); d != nil {
				return newConstValue(d)
			}
		}
	}

	return nil
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
	case t.ComplexType():
		if v := c.bigVal; v != nil {
			return c.representableComplexFromBigInt(v, t)
		}

		return c.representableComplexFromInt(c.val, t)
	case t.FloatingPointType():
		if v := c.bigVal; v != nil {
			return c.representableFloatFromBigInt(v, t)
		}

		return c.representableFloatFromInt(c.val, t)
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

func (c *intConst) convert(t Type) Const {
	ctx := t.context()
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
		todo(xc.Token{})
	default:
		panic("internal error")
	}
	return nil
}

func (c *intConst) Cpl(ctx *Context, n Node) Value {
	if c.bigVal != nil {
		var v big.Int
		if d := newIntConst(0, v.Sub(v.Neg(c.bigVal), bigInt1), c.Type(), c.Untyped()).normalize(); d != nil {
			return newConstValue(d)
		}

		return nil
	}

	if d := newIntConst(^c.val, nil, c.Type(), c.Untyped()).normalize(); d != nil {
		return newConstValue(d)
	}

	return nil
}

func (c *intConst) div(n Node, t Type, untyped bool, op Const) Const {
	ctx := t.context()
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

func (c *intConst) Div(n Node, op Value) Value {
	ctx := op.Type().context()
	switch op.Kind() {
	case ConstValue:
		t, untyped, a, b := ctx.arithmeticBinOpShape(c, op.Const(), n)
		if t != nil {
			if d := a.div(n, t, untyped, b); d != nil {
				return newConstValue(d)
			}
		}
	default:
		//dbg("", op.Kind())
		todo(n)
	}
	return nil
}

func (c *intConst) Lsh(n Node, op Value) Value {
	ctx := c.Type().context()
	switch op.Kind() {
	case ConstValue:
		d := op.Const()
		if !d.Integral() {
			todo(n)
			return nil
		}

		// The right operand in a shift expression must have unsigned
		// integer type or be an untyped constant that can be converted
		// to unsigned integer type.
		var e *intConst
		switch {
		case d.Untyped():
			if e = d.convert(ctx.intType).(*intConst); e.bigVal.Sign() < 0 {
				todo(n, true)
				return nil
			}
		default:
			if !d.Type().UnsignedIntegerType() {
				todo(n, true)
				return nil
			}

			e = d.convert(ctx.intType).(*intConst)
		}

		if uint(e.bigVal.BitLen()) > ctx.intConstBits {
			todo(n, true)
			return nil
		}

		f := c.convert(ctx.intType).(*intConst)
		var g big.Int
		h := newIntConst(0, g.Lsh(f.bigVal, uint(e.bigVal.Uint64())), ctx.intType, true).normalize()
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

func (c *intConst) mod(n Node, t Type, untyped bool, op Const) Const {
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
	//	e.Div(c.bigVal, d)
	//	if untyped {
	//		t = nil
	//	}
	//	return ctx.mustConvertConst(n, t, newIntConst(0, &e, ctx.intType, true))
}

func (c *intConst) Mod(n Node, op Value) Value {
	ctx := op.Type().context()
	switch op.Kind() {
	case ConstValue:
		t, untyped, a, b := ctx.arithmeticBinOpShape(c, op.Const(), n)
		if t != nil {
			if d := a.mod(n, t, untyped, b); d != nil {
				return newConstValue(d)
			}
		}
	default:
		//dbg("", op.Kind())
		todo(n)
	}
	return nil
}

func (c *intConst) mul(n Node, t Type, untyped bool, op Const) Const {
	ctx := t.context()
	var d big.Int
	d.Mul(c.bigVal, op.(*intConst).bigVal)
	if untyped {
		t = nil
	}
	e := newIntConst(0, &d, ctx.intType, true)
	if e.normalize() == nil {
		ctx.err(n, "constant multiplication overflow")
		return nil
	}

	return ctx.mustConvertConst(n, t, e)
}

func (c *intConst) Mul(n Node, op Value) Value {
	ctx := op.Type().context()
	switch op.Kind() {
	case ConstValue:
		t, untyped, a, b := ctx.arithmeticBinOpShape(c, op.Const(), n)
		if t != nil {
			if d := a.mul(n, t, untyped, b); d != nil {
				return newConstValue(d)
			}
		}
	default:
		//dbg("", op.Kind())
		todo(n)
	}
	return nil
}

func (c *intConst) Neg(ctx *Context, n Node) Value {
	if c.bigVal != nil {
		var v big.Int
		if d := newIntConst(0, v.Neg(v.Set(c.bigVal)), c.Type(), c.Untyped()).normalize(); d != nil {
			return newConstValue(d)
		}

		return nil
	}

	var d Const
	switch {
	case c.val == math.MinInt64:
		var v big.Int
		d = newIntConst(0, v.Neg(v.SetInt64(c.val)), c.Type(), c.Untyped()).normalize()
	default:
		d = newIntConst(-c.val, nil, c.Type(), c.Untyped()).normalize()
	}
	if d != nil {
		return newConstValue(d)
	}

	return nil
}

func (c *intConst) Rsh(n Node, op Value) Value {
	ctx := c.Type().context()
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
			if e = d.convert(ctx.intType).(*intConst); e.bigVal.Sign() < 0 {
				todo(n, true)
				return nil
			}
		default:
			if !d.Type().UnsignedIntegerType() {
				todo(n, true)
				return nil
			}

			e = d.convert(ctx.intType).(*intConst)
		}

		if uint(e.bigVal.BitLen()) > ctx.intConstBits {
			todo(n, true)
			return nil
		}

		f := c.convert(ctx.intType).(*intConst)
		var g big.Int
		h := newIntConst(0, g.Rsh(f.bigVal, uint(e.bigVal.Uint64())), ctx.intType, true).normalize()
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

func (c *intConst) sub(n Node, t Type, untyped bool, op Const) Const {
	ctx := t.context()
	var d big.Int
	d.Sub(c.bigVal, op.(*intConst).bigVal)
	if untyped {
		t = nil
	}
	return ctx.mustConvertConst(n, t, newIntConst(0, &d, ctx.intType, true))
}

func (c *intConst) Sub(n Node, op Value) Value {
	ctx := op.Type().context()
	switch op.Kind() {
	case ConstValue:
		t, untyped, a, b := ctx.arithmeticBinOpShape(c, op.Const(), n)
		if t != nil {
			if d := a.sub(n, t, untyped, b); d != nil {
				return newConstValue(d)
			}
		}
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

func (c *stringConst) AssignableTo(t Type) bool { return c.assignableTo(c, t) }
func (c *stringConst) normalize() Const         { return c }

func (c *stringConst) Add(n Node, op Value) Value {
	todo(n)
	return nil
}

func (c *stringConst) Convert(t Type) Value {
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

func (c *stringConst) representable(t Type) Const {
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
