// Copyright 2016 The GC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Portions of this source file is a derived work. The original code is at
//
//	https://github.com/golang/go/blob/e805bf39458915365924228dc53969ce04e32813/src/reflect/type.go
//
// and contains the following copyright notice.
//
// ---------------------------------------------------
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// ---------------------------------------------------
//
// The LICENSE file mentioned just above is available as GO-LICENSE in this
// repository.

package gc

import (
	"bytes"
	"fmt"

	"github.com/cznic/mathutil"
)

var (
	_ Selector = (*Method)(nil)
	_ Selector = (*StructField)(nil)

	_ Type = (*TypeDeclaration)(nil)
	_ Type = (*arrayType)(nil)
	_ Type = (*chanType)(nil)
	_ Type = (*typeBase)(nil)
	_ Type = (*funcType)(nil)
	_ Type = (*interfaceType)(nil)
	_ Type = (*mapType)(nil)
	_ Type = (*ptrType)(nil)
	_ Type = (*sliceType)(nil)
	_ Type = (*structType)(nil)
	_ Type = (*tupleType)(nil)
)

// Selector is either a *StructField or a *Method. The pointee is read only.
type Selector interface {
	typ() Type
}

// Values of type ChanDir.
const (
	RecvDir ChanDir             = 1 << iota // <-chan
	SendDir                                 // chan<-
	BothDir = RecvDir | SendDir             // chan
)

// Values of type Kind.
const (
	Invalid Kind = iota
	Bool
	Int
	Int8
	Int16
	Int32
	Int64
	Uint
	Uint8
	Uint16
	Uint32
	Uint64
	Uintptr
	Float32
	Float64
	Complex64
	Complex128
	Array
	Chan
	Func
	Interface
	Map
	Ptr
	Slice
	String
	Struct
	UnsafePointer
	Tuple
	UntypedBool
	maxKind
)

// ChanDir represents a channel type's direction.
type ChanDir int

// A Kind represents the specific kind of type that a Type represents. The zero
// Kind is not a valid kind.
type Kind uint

// Type is the representation of a Go type.
//
// Not all methods apply to all kinds of types.  Restrictions, if any, are
// noted in the documentation for each method.  Use the Kind method to find out
// the kind of type before calling kind-specific methods.  Calling a method
// inappropriate to the kind of type causes a run-time panic.
type Type interface {
	// --------------------------------------------------------------------
	// Methods applicable to all types.

	base() *typeBase
	field(int) Type
	flags() flags
	implementsFailed(*context, Node, string, Type) bool
	isNamed() bool
	numField() int
	setOffset(int, uint64)
	signatureString(int) string
	str(*bytes.Buffer)

	// Align returns the alignment in bytes of a value of this type when
	// allocated in memory.
	Align() int

	// FieldAlign returns the alignment in bytes of a value of this type
	// when used as a field in a struct.
	FieldAlign() int

	// Method returns the i'th method in the type's method set.  It panics
	// if i is not in the range [0, NumMethod()).
	//
	// For a non-interface type T or *T, the returned Method's Type
	// describes a function whose first argument is the receiver.
	//
	// For an interface type, the returned Method's Type field gives the
	// method signature, without a receiver. The result pointee is read
	// only.
	Method(int) *Method

	// MethodByName returns the method with that name ID in the type's
	// method set or nil if the method was not found.
	//
	// For a non-interface type T or *T, the returned Method's Type
	// describes a function whose first argument is the receiver.
	//
	// For an interface type, the returned Method's Type field gives the
	// method signature, without a receiver. The result pointee is read
	// only.
	MethodByName(int) *Method

	// NumMethod returns the number of methods in the type's method set.
	NumMethod() int

	// Name returns the type's name ID within its package.  It returns an
	// empty string for unnamed types.
	Name() int

	// PkgPath returns a named type's package path ID, that is, the import
	// path that uniquely identifies the package, such as
	// "encoding/base64".  If the type was predeclared (string, error) or
	// unnamed (*T, struct{}, []int), the package path ID will be zero.
	PkgPath() int

	// Size returns the number of bytes needed to store a value of the
	// given type; it is analogous to unsafe.Sizeof.
	Size() uint64

	// String returns a string representation of the type.  The string
	// representation may use shortened package names (e.g., base64 instead
	// of "encoding/base64") and is not guaranteed to be unique among
	// types.  To test for equality, compare the Types directly.
	String() string

	// Kind returns the specific kind of this type.
	Kind() Kind

	// Identical reports whether this type is identical to u.
	//
	// See https://golang.org/ref/spec#Type_identity
	Identical(u Type) bool

	// Implements reports whether the type implements the interface type u.
	Implements(u Type) bool

	// UnderlyingType returns the underlying type of this type.
	//
	// See https://golang.org/ref/spec#Types
	UnderlyingType() Type

	// AssignableTo reports whether a value of the type is assignable to
	// type u.
	//
	// See https://golang.org/ref/spec#Assignability
	AssignableTo(u Type) bool

	// IntegerType reports whether this type's Kind is one of Int*, Uint*
	// including Uintptr.
	IntegerType() bool

	// FloatingPointType reports whether this type's Kind is Float32 or
	// Float64.
	FloatingPointType() bool

	// ComplexType reports whether this type's Kind is Complex64 or
	// Complex128.
	ComplexType() bool

	// ConvertibleTo reports whether a value of the type is convertible to
	// type u.
	//
	// See https://golang.org/ref/spec#Conversions
	ConvertibleTo(u Type) bool

	// Comparable reports whether values of this type are comparable.
	Comparable() bool

	// Ordered reports whether values of this type are ordered.
	Ordered() bool

	// UnsignedIntegerType reports whether this type's Kind is one of Uint*
	// including Uintptr.
	UnsignedIntegerType() bool

	// Numeric reports whether the this type is a number.
	Numeric() bool

	// --------------------------------------------------------------------
	// Methods applicable only to some types, depending on Kind.
	// The methods allowed for each kind are:
	//
	//	Int*, Uint*, Float*, Complex*: Bits
	//	Array: Elem, Len
	//	Chan: ChanDir, Elem
	//	Func: In, NumIn, Out, NumOut, IsVariadic, Result.
	//	Map: Key, Elem
	//	Ptr: Elem
	//	Slice: Elem
	//	Struct: Field, FieldByIndex, FieldByName, FieldByNameFunc, NumField
	//	Tuple: Elements

	// Bits returns the size of the type in bits.  It panics if the type's
	// Kind is not one of the sized or unsized Int, Uint, Float, or Complex
	// kinds.
	Bits(*Context) int

	// ChanDir returns a channel type's direction.  It panics if the type's
	// Kind is not Chan.
	ChanDir() ChanDir

	// IsVariadic reports whether a function type's final input parameter
	// is a "..." parameter.  If so, t.In(t.NumIn() - 1) returns the
	// parameter's implicit actual type []T.
	//
	// For concreteness, if t represents func(x int, y ... float64), then
	//
	//	t.NumIn() == 2
	//	t.In(0) is the reflect.Type for "int"
	//	t.In(1) is the reflect.Type for "[]float64"
	//	t.IsVariadic() == true
	//
	// IsVariadic panics if the type's Kind is not Func.
	IsVariadic() bool

	// Elem returns a type's element type.
	// It panics if the type's Kind is not Array, Chan, Map, Ptr, or Slice.
	Elem() Type

	// Elements return the types this type consist of. It panics if the
	// type's Kind is not Tuple. The result is read only.
	Elements() []Type

	// Field returns a struct type's i'th field.  It panics if the type's
	// Kind is not Struct.  It panics if i is not in the range [0,
	// NumField()). The result pointee is read only.
	Field(i int) *StructField

	// FieldByIndex returns the nested field corresponding to the index
	// sequence.  It is equivalent to calling Field successively for each
	// index i.  It panics if the type's Kind is not Struct.
	FieldByIndex(index []int) StructField

	// FieldByName returns the struct field with the given name ID or nil
	// if the field was not found. The result pointee is read only.
	FieldByName(name int) *StructField

	//TODO
	//
	// BUG(rsc): FieldByName and related functions consider struct field
	// names to be equal if the names are equal, even if they are
	// unexported names originating in different packages. The practical
	// effect of this is that the result of t.FieldByName("x") is not well
	// defined if the struct type t contains multiple fields named x
	// (embedded from different packages).  FieldByName may return one of
	// the fields named x or may report that there are none.  See
	// golang.org/issue/4876 for more details.

	// FieldByNameFunc returns the first struct field with a name ID that
	// satisfies the match function or nil if the field was not found. The
	// result pointee is read only.
	FieldByNameFunc(match func(int) bool) *StructField

	// In returns the type of a function type's i'th input parameter.  It
	// panics if the type's Kind is not Func.  It panics if i is not in the
	// range [0, NumIn()).
	In(i int) Type

	// Key returns a map type's key type.  It panics if the type's Kind is
	// not Map.
	Key() Type

	// Len returns an array type's length.  It panics if the type's Kind is
	// not Array.
	Len() int64

	// It panics if the type's Kind is not Struct.
	// NumField returns a struct type's field count.
	NumField() int

	// NumIn returns a function type's input parameter count.  It panics if
	// the type's Kind is not Func.
	NumIn() int

	// NumOut returns a function type's output parameter count.  It panics
	// if the type's Kind is not Func.
	NumOut() int

	// Out returns the type of a function type's i'th output parameter.  It
	// panics if the type's Kind is not Func.  It panics if i is not in the
	// range [0, NumOut()).
	Out(i int) Type

	// Result returns the function result type. It panics if the type's
	// Kind is not Func.
	Result() Type
}

// Method represents a single method.
type Method struct {
	// Name is the method name ID.  PkgPath is the package path ID that
	// qualifies a lower case (unexported) method name.  It is zero for
	// upper case (exported) method names.  The combination of PkgPath and
	// Name uniquely identifies a method in a method set.
	//
	// See https://golang.org/ref/spec#Uniqueness_of_identifiers
	Name    int
	PkgPath int

	Type   Type // Method type.
	Index  int  // Index for Type.Method.
	merged bool
}

func (m *Method) typ() Type { return m.Type }

// A StructField describes a single field in a struct.
type StructField struct {
	// Name is the field name ID.
	Name int
	// PkgPath is the package path ID that qualifies a lower case
	// (unexported) field name.  It is zero for upper case (exported) field
	// names.
	//
	// See https://golang.org/ref/spec#Uniqueness_of_identifiers
	PkgPath int

	Type      Type   // Field type.
	Tag       int    // Field tag string ID.
	Offset    uint64 // Offset within struct, in bytes.
	Index     []int  // Index sequence for Type.FieldByIndex.
	Anonymous bool   // Is an embedded field.
}

func (s *StructField) typ() Type { return s.Type }

type typeBase struct {
	align      int
	fieldAlign int
	kind       Kind
	methods    []Method
	pkgPath    int
	size       uint64
	typ        Type
}

func (t *typeBase) Align() int                        { return t.align }
func (t *typeBase) base() *typeBase                   { return t }
func (t *typeBase) ChanDir() ChanDir                  { panic("ChanDir of non-chan type") }
func (t *typeBase) Elements() []Type                  { panic("Elements of non-tuple type") }
func (t *typeBase) Field(int) *StructField            { panic("Field of non-struct type") }
func (t *typeBase) field(int) Type                    { panic("internal error") }
func (t *typeBase) FieldAlign() int                   { return t.fieldAlign }
func (t *typeBase) FieldByIndex([]int) StructField    { panic("FieldByIndex of non-struct type") }
func (t *typeBase) FieldByName(name int) *StructField { panic("FieldByName of non-struct type") }
func (t *typeBase) flags() flags                      { panic("internal error") }
func (t *typeBase) Identical(Type) bool               { panic("internal error") }
func (t *typeBase) In(i int) Type                     { panic("In of a non-func type") }
func (t *typeBase) isNamed() bool                     { return t.pkgPath != 0 }
func (t *typeBase) IsVariadic() bool                  { panic("IsVariadic of non-func type") }
func (t *typeBase) Key() Type                         { panic("Key of non-map type") }
func (t *typeBase) Kind() Kind                        { return t.kind }
func (t *typeBase) Len() int64                        { panic("Len of non-array type") }
func (t *typeBase) Name() int                         { return 0 }
func (t *typeBase) numField() int                     { panic("internal error") }
func (t *typeBase) NumField() int                     { panic("NumField of non-struct type") }
func (t *typeBase) NumIn() int                        { panic("NumIn of non-func type") }
func (t *typeBase) NumMethod() int                    { return len(t.methods) }
func (t *typeBase) NumOut() int                       { panic("NumOut of non-func type") }
func (t *typeBase) Out(i int) Type                    { panic("Out of a non-func type") }
func (t *typeBase) PkgPath() int                      { return t.pkgPath }
func (t *typeBase) Result() Type                      { panic("Result of a non-func type") }
func (t *typeBase) setOffset(int, uint64)             { panic("internal error") }
func (t *typeBase) signatureString(int) string        { panic("internal error") }
func (t *typeBase) Size() uint64                      { return t.size }
func (t *typeBase) str(*bytes.Buffer)                 { panic("internal error") }

func (t *typeBase) AssignableTo(u Type) bool {
	v := t.typ

	// A value x is assignable to a variable of type U ("x is assignable to
	// U") in any of these cases:

	// · x's type is identical to U.
	if v.Identical(u) {
		return true
	}

	// · Untyped bool is assignable to any bool type.
	if v.Kind() == UntypedBool && u.Kind() == Bool {
		return true
	}

	// · x's type V and U have identical underlying types and at least one
	// of V or U is not a named type.
	if (v.Name() == 0 || u.Name() == 0) &&
		v.UnderlyingType().Identical(u.UnderlyingType()) {
		return true
	}

	// · U is an interface type and x implements U.
	if u.Kind() == Interface && v.Implements(u) {
		return true
	}

	// · x is a bidirectional channel value, U is a channel type, x's type
	// V and U have identical element types, and at least one of V or U is
	// not a named type.
	if (v.Name() == 0 || u.Name() == 0) && v.Kind() == Chan &&
		u.Kind() == Chan && v.ChanDir() == BothDir &&
		v.Elem().Identical(u.Elem()) {
		return true
	}

	// · x is the predeclared identifier nil and U is a pointer, function,
	// slice, map, channel, or interface type.
	//
	// N/A here

	// · x is a untyped constant representable by a value of type U.
	//
	// N/A here

	return false
}

func (t *typeBase) Bits(ctx *Context) int {
	switch t.Kind() {
	case Int8, Uint8:
		return 8
	case Int16, Uint16:
		return 16
	case Int32, Uint32, Float32:
		return 32
	case Int64, Uint64, Float64, Complex64:
		return 64
	case Int, Uint:
		return ctx.model.IntBytes * 8
	case Complex128:
		return 128
	}
	panic("Bits of inappropriate type")
}

func (t *typeBase) Comparable() bool {
	k := t.Kind()
	tt := t.typ

	// · Boolean values are comparable. Two boolean values are equal if
	// they are either both true or both false.
	if k == Bool || k == UntypedBool {
		return true
	}

	// · Integer values are comparable and ordered, in the usual way.
	if t.IntegerType() {
		return true
	}

	// · Floating point values are comparable and ordered, as defined by
	// the IEEE-754 standard.
	if t.FloatingPointType() {
		return true
	}

	// · Complex values are comparable. Two complex values u and v are
	// equal if both real(u) == real(v) and imag(u) == imag(v).
	if t.ComplexType() {
		return true
	}

	// · String values are comparable and ordered, lexically byte-wise.
	if k == String {
		return true
	}

	// · Pointer values are comparable. Two pointer values are equal if
	// they point to the same variable or if both have value nil. Pointers
	// to distinct zero-size variables may or may not be equal.
	if k == Ptr {
		return true
	}

	// · Channel values are comparable. Two channel values are equal if
	// they were created by the same call to make or if both have value
	// nil.
	if k == Chan {
		return true
	}

	// · Interface values are comparable. Two interface values are equal if
	// they have identical dynamic types and equal dynamic values or if
	// both have value nil.
	if k == Interface {
		return true
	}

	// · A value x of non-interface type X and a value t of interface type
	// T are comparable when values of type X are comparable and X
	// implements T. They are equal if t's dynamic type is identical to X
	// and t's dynamic value is equal to x.
	//
	// N/A here.

	// · Struct values are comparable if all their fields are comparable.
	// Two struct values are equal if their corresponding non-blank fields
	// are equal.
	if k == Struct {
		n := tt.NumField()
		for i := 0; i < n; i++ {
			if !tt.Field(i).Type.Comparable() {
				return false
			}
		}
		return true
	}

	// · Array values are comparable if values of the array element type
	// are comparable. Two array values are equal if their corresponding
	// elements are equal.
	if k == Array {
		return true
	}

	return false
}

func (t *typeBase) ComplexType() bool {
	k := t.Kind()
	return k == Complex64 || k == Complex128
}

func (t *typeBase) ConvertibleTo(u Type) bool {
	xt := t.typ
	if xt == nil {
		return false
	}

	// A non-constant value x can be converted to type U in any of these
	// cases:

	// · x is assignable to U.
	if xt.AssignableTo(u) {
		return true
	}

	ptr := 0
	ptrBytes := int(xt.Size())
	for xt.Kind() == Ptr {
		ptr++
		xt = xt.Elem()
		if xt == nil {
			return false
		}
	}

	if ptr != 0 {
		xt = xt.UnderlyingType()
		if xt == nil {
			return false
		}

		for ; ptr != 0; ptr-- {
			xt = newPtrType(ptrBytes, xt)
		}
	}

	// · x's type and U have identical underlying types.
	if xt.UnderlyingType().Identical(u.UnderlyingType()) {
		return true
	}

	// · x's type and U are unnamed pointer types and their pointer base
	// types have identical underlying types.
	if xt.Name() == 0 && u.Name() == 0 && xt.Kind() == Ptr &&
		u.Kind() == Ptr && xt.Elem().Identical(u.Elem()) {
		return true
	}

	// · x's type and U are both integer or floating point types.
	if (xt.IntegerType() || xt.FloatingPointType()) && (u.IntegerType() || u.FloatingPointType()) {
		return true
	}

	// · x's type and U are both complex types.
	if xt.ComplexType() && u.ComplexType() {
		return true
	}

	// · x is an integer or a slice of bytes or runes and U is a string
	// type.
	if (xt.IntegerType() ||
		xt.Kind() == Slice && (xt.Elem().Kind() == Uint8 || xt.Elem().Kind() == Int32)) &&
		u.Kind() == String {
		return true
	}

	// · x is a string and U is a slice of bytes or runes.
	if xt.Kind() == String &&
		(u.Kind() == Slice && (u.Elem().Kind() == Uint8 || u.Elem().Kind() == Int32)) {
		return true
	}

	if u.Kind() == UnsafePointer && (xt.Kind() == Ptr || xt.Kind() == Uintptr || xt.Kind() == UnsafePointer) {
		return true
	}

	if xt.Kind() == UnsafePointer && (u.Kind() == Ptr || u.Kind() == Uintptr || u.Kind() == UnsafePointer) {
		return true
	}

	return false
}

func (t *typeBase) Elem() Type {
	switch n := t.typ.UnderlyingType(); n.Kind() {
	case Array, Slice, Ptr:
		return n.Elem()
	default:
		panic("Elem of inappropriate type")
	}
}

func (t *typeBase) FieldByNameFunc(match func(int) bool) *StructField {
	panic("FieldByNameFunc of a non-struct type")
}

func (t *typeBase) FloatingPointType() bool {
	k := t.Kind()
	return k == Float32 || k == Float64
}

func (t *typeBase) implementsFailed(ctx *context, n Node, msg string, u Type) (stop bool) {
	if u.Kind() != Interface {
		panic("internal error")
	}

	s := fmt.Sprintf(msg, t, u)

	if t.Kind() == Ptr {
		t = t.Elem().base()
	}

	nu := u.NumMethod()
	if nu == 0 {
		panic("internal error")
	}

	var off int
	if t.Kind() != Interface {
		off = 1
	}
	for i := 0; i < nu; i++ {
		mu := u.Method(i)
		mt := t.MethodByName(mu.Name)
		if mt == nil || mt.PkgPath != mu.PkgPath {
			return ctx.err(n, "%s\n\t%s does not implement %s (missing %s method)", s, t, u, dict.S(mu.Name))
		}

		tt := mt.Type
		it := tt.NumIn()
		ot := tt.NumOut()
		tu := mu.Type
		iu := tu.NumIn()
		ou := tu.NumOut()
		if it-off != iu || ot != ou || tt.IsVariadic() != tu.IsVariadic() {
			nm := dict.S(mu.Name)
			return ctx.err(
				n,
				`%s
	%s does not implement %s (wrong type for %s method)
		have %s%s
		want %s%s`,
				s, t, u, nm,
				nm, tt.signatureString(1),
				nm, tu.signatureString(0),
			)
		}

		for j := 0; j < iu; j++ {
			if !tt.In(j + off).Identical(tu.In(j)) {
				todo(n, true)
				return false
			}
		}
		for j := 0; j < ou; j++ {
			if !tt.Out(j).Identical(tu.Out(j)) {
				todo(n, true)
				return false
			}
		}
	}
	panic("internal error")
}

func (t *typeBase) Implements(u Type) bool {
	if u.Kind() != Interface {
		panic("non-interface argument passed to Implements")
	}

	if t.Kind() == Ptr {
		t = t.Elem().base()
	}

	nu := u.NumMethod()
	if nu == 0 {
		return true // Fast path: Every type implements interface{}.
	}

	nt := t.NumMethod()
	if nt < nu { // Fast path: t has not enough methods to implement u.
		return false
	}

	var off int
	if t.Kind() != Interface {
		off = 1
	}
	for i := 0; i < nu; i++ {
		mu := u.Method(i)
		mt := t.MethodByName(mu.Name)
		if mt == nil || mt.PkgPath != mu.PkgPath {
			return false
		}

		tt := mt.Type
		it := tt.NumIn()
		ot := tt.NumOut()
		tu := mu.Type
		iu := tu.NumIn()
		ou := tu.NumOut()
		if it-off != iu || ot != ou || tt.IsVariadic() != tu.IsVariadic() {
			return false
		}

		for j := 0; j < iu; j++ {
			if !tt.In(j + off).Identical(tu.In(j)) {
				return false
			}
		}
		for j := 0; j < ou; j++ {
			if !tt.Out(j).Identical(tu.Out(j)) {
				return false
			}
		}
	}
	return true
}

func (t *typeBase) IntegerType() bool {
	switch t.Kind() {
	case
		Int8, Int16, Int32, Int64, Int,
		Uint8, Uint16, Uint32, Uint64, Uint, Uintptr:
		return true
	}
	return false
}

func (t *typeBase) Method(i int) *Method {
	if i < 0 || i >= len(t.methods) {
		panic("Method: index out of range")
	}

	return &t.methods[i]
}

func (t *typeBase) MethodByName(nm int) *Method {
	for i, v := range t.methods {
		if v.Name == nm {
			return &t.methods[i]
		}
	}
	return nil
}

func (t *typeBase) Numeric() bool {
	u := t.typ
	return u.IntegerType() || u.FloatingPointType() || u.ComplexType()
}

func (t *typeBase) Ordered() bool {
	switch t.Kind() {
	case
		Int8, Int16, Int32, Int64, Int,
		Uint8, Uint16, Uint32, Uint64, Uint, Uintptr,
		Float32, Float64,
		Complex64, Complex128,
		String:
		return true
	}
	return false
}

func (t *typeBase) UnderlyingType() Type {
	switch t := t.typ; x := t.(type) {
	case *TypeDeclaration:
		if x.typ != nil {
			return x.typ.UnderlyingType()
		}

		return x
	default:
		return t
	}
}

func (t *typeBase) UnsignedIntegerType() bool {
	switch t.Kind() {
	case
		Uint8, Uint16, Uint32, Uint64, Uint, Uintptr:
		return true
	}
	return false
}

func (t *typeBase) String() string {
	var buf bytes.Buffer
	safeTypeStr(t.typ, &buf)
	return buf.String()
}

// ------------------------------------------------------------------ arrayType

type arrayType struct {
	typeBase
	elem Type
	len  int64
	flgs flags
}

func newArrayType(elem Type, len int64, flags flags) *arrayType {
	t := &arrayType{elem: elem, len: len}
	t.flgs = flags
	t.align = elem.Align()
	t.fieldAlign = elem.FieldAlign()
	t.kind = Array
	t.size = elem.Size() * uint64(len)
	if elem != nil {
		t.flgs = t.flgs | elem.flags()
	}
	t.typ = t
	return t
}

func (t *arrayType) Elem() Type   { return t.elem }
func (t *arrayType) Len() int64   { return t.len }
func (t *arrayType) flags() flags { return t.flgs }

func (t *arrayType) Identical(u Type) bool {
	return u.Kind() == Array && t.Len() == u.Len() && t.Elem().Identical(u.Elem())
}

func (t *arrayType) str(w *bytes.Buffer) {
	w.WriteByte('[')
	switch {
	case t.len < 0:
		w.WriteString("...")
	default:
		fmt.Fprintf(w, "%v", t.Len())
	}
	w.WriteByte(']')
	safeTypeStr(t.Elem(), w)
}

// ------------------------------------------------------------------- chanType

type chanType struct {
	typeBase
	dir  ChanDir
	elem Type
	flgs flags
}

func newChanType(ctx *context, dir ChanDir, elem Type) *chanType {
	t := &chanType{dir: dir, elem: elem}
	t.align = ctx.model.PtrBytes
	t.fieldAlign = ctx.model.PtrBytes
	t.kind = Chan
	t.size = uint64(ctx.model.PtrBytes)
	if elem != nil {
		t.flgs = elem.flags()
	}
	t.typ = t
	return t
}

func (t *chanType) ChanDir() ChanDir { return t.dir }
func (t *chanType) Elem() Type       { return t.elem }
func (t *chanType) flags() flags     { return t.flgs }

func (t *chanType) Identical(u Type) bool {
	return u.Kind() == Chan && t.ChanDir() == u.ChanDir() && t.Elem().Identical(u.Elem())
}

func (t *chanType) str(w *bytes.Buffer) {
	switch t.dir {
	case RecvDir:
		w.WriteString("<-chan ")
	case SendDir:
		w.WriteString("chan<- ")
	case BothDir:
		w.WriteString("chan ")
	default:
		panic("internal error")
	}
	safeTypeStr(t.Elem(), w)
}

// ------------------------------------------------------------------- funcType

type funcType struct {
	in         []Type
	isExported bool
	isVariadic bool
	name       int // For methods only.
	result     Type
	typeBase
	flgs flags
}

func newFuncType(ctx *context, name int, in []Type, result Type, isExported, isVariadic bool) *funcType {
	t := &funcType{name: name, in: in, result: result, isExported: isExported, isVariadic: isVariadic}
	t.align = ctx.model.PtrBytes
	t.fieldAlign = ctx.model.PtrBytes
	t.kind = Func
	t.size = uint64(ctx.model.PtrBytes)
	if result != nil {
		t.flgs = result.flags()
	}
	for _, v := range in {
		if v != nil {
			t.flgs = t.flgs | v.flags()
		}
	}
	t.typ = t
	return t
}

func (t *funcType) IsVariadic() bool { return t.isVariadic }
func (t *funcType) NumIn() int       { return len(t.in) }
func (t *funcType) Result() Type     { return t.result }
func (t *funcType) flags() flags     { return t.flgs }

func (t *funcType) ptrMethod(ctx *context) *funcType {
	u := *t
	u.in = append([]Type(nil), t.in...)
	u.in[0] = newPtrType(ctx.model.PtrBytes, t.in[0])
	u.typ = &u
	return &u
}

func (t *funcType) Identical(u Type) bool {
	if u.Kind() != Func {
		return false
	}

	it := t.NumIn()
	ot := t.NumOut()
	iu := u.NumIn()
	ou := u.NumOut()
	if it != iu || ot != ou || t.IsVariadic() != u.IsVariadic() {
		return false
	}

	for i := 0; i < it; i++ {
		if !t.In(i).Identical(u.In(i)) {
			return false
		}
	}

	for i := 0; i < ot; i++ {
		if !t.Out(i).Identical(u.Out(i)) {
			return false
		}
	}

	return true
}

func (t *funcType) In(i int) Type {
	if i < 0 || i >= len(t.in) {
		panic("In: index out of range")
	}

	return t.in[i]
}

func (t *funcType) NumOut() int {
	switch {
	case t.result == nil:
		return 0
	case t.result.Kind() == Tuple:
		return len(t.Result().Elements())
	default:
		return 1
	}
}

func (t *funcType) Out(i int) Type {
	if i < 0 || i >= t.NumOut() {
		panic("Out: index out of range")
	}

	switch {
	case t.result.Kind() == Tuple:
		return t.Result().Elements()[i]
	default:
		return t.result
	}
}

func (t *funcType) str(w *bytes.Buffer) {
	w.WriteString("func")
	t.signatureStr(0, w)
}

func (t *funcType) signatureStr(skip int, w *bytes.Buffer) {
	w.WriteByte('(')
	in := t.in[skip:]
	for i, v := range in {
		safeTypeStr(v, w)
		if i != len(in)-1 {
			w.WriteString(", ")
		}
	}
	w.WriteByte(')')
	if t.NumOut() != 0 {
		w.WriteByte(' ')
		safeTypeStr(t.result, w)
	}
}

func (t *funcType) signatureString(skip int) string {
	var buf bytes.Buffer
	t.signatureStr(skip, &buf)
	return buf.String()
}

// -------------------------------------------------------------- interfaceType

type interfaceType struct {
	typeBase
	flgs flags
}

func newInterfaceType(ctx *context, methods []Method) *interfaceType {
	t := &interfaceType{}
	t.align = ctx.model.PtrBytes
	t.fieldAlign = ctx.model.PtrBytes
	t.kind = Interface
	t.methods = methods
	t.size = uint64(2 * ctx.model.PtrBytes)
	for _, m := range methods {
		if m.Type != nil {
			t.flgs = t.flgs | m.Type.flags()
		}
	}
	t.typ = t
	return t
}

func (t *interfaceType) flags() flags { return t.flgs }

func (t *interfaceType) Identical(u Type) bool {
	if u.Kind() != Interface {
		return false
	}

	nt := t.NumMethod()
	nu := u.NumMethod()
	if nt != nu {
		return false
	}

	for i := 0; i < nt; i++ {
		mt := t.Method(i)
		mu := t.MethodByName(mt.Name)
		if mu == nil || mt.PkgPath != mu.PkgPath || !mt.Type.Identical(mu.Type) {
			return false
		}
	}
	return true
}

func (t *interfaceType) str(w *bytes.Buffer) {
	w.WriteString("interface{")
	for i, v := range t.methods {
		w.Write(dict.S(v.Name))
		v.Type.(*funcType).signatureStr(0, w)
		if i != len(t.methods)-1 {
			w.WriteString("; ")
		}
	}
	w.WriteByte('}')
}

// -------------------------------------------------------------------- mapType

type mapType struct {
	typeBase
	elem Type
	key  Type
	flgs flags
}

func newMapType(ctx *context, key, elem Type) *mapType {
	t := &mapType{elem: elem, key: key}
	if key != nil {
		t.flgs = key.flags()
	}
	if elem != nil {
		t.flgs = t.flgs | elem.flags()
	}
	t.align = ctx.model.PtrBytes
	t.fieldAlign = ctx.model.PtrBytes
	t.kind = Map
	t.size = uint64(ctx.model.PtrBytes)
	t.typ = t
	return t
}

func (t *mapType) Elem() Type   { return t.elem }
func (t *mapType) Key() Type    { return t.key }
func (t *mapType) flags() flags { return t.flgs }

func (t *mapType) Identical(u Type) bool {
	return u.Kind() == Map && t.Key().Identical(u.Key()) && t.Elem().Identical(u.Elem())
}

func (t *mapType) str(w *bytes.Buffer) {
	w.WriteString("map[")
	safeTypeStr(t.Key(), w)
	w.WriteByte(']')
	safeTypeStr(t.Elem(), w)
}

// -------------------------------------------------------------------- ptrType

type ptrType struct {
	typeBase
	elem Type
	flgs flags
}

func newPtrType(ptrBytes int, elem Type) *ptrType {
	t := &ptrType{elem: elem}
	if elem != nil {
		t.flgs = elem.flags()
	}
	t.align = ptrBytes
	t.fieldAlign = ptrBytes
	t.kind = Ptr
	t.size = uint64(ptrBytes)
	t.typ = t
	return t
}

func (t *ptrType) Elem() Type            { return t.elem }
func (t *ptrType) Identical(u Type) bool { return u.Kind() == Ptr && t.Elem().Identical(u.Elem()) }
func (t *ptrType) flags() flags          { return t.flgs }

func (t *ptrType) str(w *bytes.Buffer) {
	w.WriteByte('*')
	safeTypeStr(t.Elem(), w)
}

// ------------------------------------------------------------------ sliceType

type sliceType struct {
	typeBase
	elem Type
	flgs flags
}

func newSliceType(ctx *context, elem Type) *sliceType {
	t := &sliceType{elem: elem}
	if elem != nil {
		t.flgs = elem.flags()
	}
	t.align = ctx.model.PtrBytes
	t.fieldAlign = ctx.model.PtrBytes
	t.kind = Slice
	t.size = uint64(3 * ctx.model.PtrBytes)
	t.typ = t
	return t
}

func (t *sliceType) Elem() Type            { return t.elem }
func (t *sliceType) Identical(u Type) bool { return u.Kind() == Slice && t.Elem().Identical(u.Elem()) }
func (t *sliceType) flags() flags          { return t.flgs }

func (t *sliceType) str(w *bytes.Buffer) {
	w.WriteString("[]")
	safeTypeStr(t.Elem(), w)
}

// ----------------------------------------------------------------- structType

type structType struct {
	fields []StructField
	typeBase
	flgs flags
	node Node //TODO-
}

func newStructType(n Node, fields []StructField, methods []Method) *structType {
	t := &structType{fields: fields, node: n}
	for _, v := range fields {
		if v.Type != nil {
			t.flgs = t.flgs | v.Type.flags()
		}
	}
	t.kind = Struct
	t.methods = methods
	t.typ = t
	return t
}

func (t *structType) field(i int) Type          { return t.fields[i].Type }
func (t *structType) NumField() int             { return len(t.fields) }
func (t *structType) numField() int             { return len(t.fields) }
func (t *structType) setOffset(i int, o uint64) { t.Field(i).Offset = o }
func (t *structType) flags() flags              { return t.flgs }

func (t *structType) Field(i int) *StructField {
	if i < 0 || i > len(t.fields) {
		panic("Field: index out of range")
	}

	return &t.fields[i]
}

func (t *structType) FieldByIndex(index []int) StructField {
	var u Type = t
	var f *StructField
	for _, v := range index {
		f := u.Field(v)
		u = f.Type
	}
	r := *f
	r.Index = index
	return r
}

func (t *structType) FieldByName(name int) *StructField {
	for i, v := range t.fields {
		if name == v.Name {
			return &t.fields[i]
		}
	}
	return nil
}

func (t *structType) FieldByNameFunc(match func(int) bool) *StructField {
	for i, v := range t.fields {
		if match(v.Name) {
			return &t.fields[i]
		}
	}
	return nil
}

func (t *structType) Identical(u Type) bool {
	if u.Kind() != Struct {
		return false
	}

	n := t.NumField()
	if n != u.NumField() {
		return false
	}

	for i := 0; i < n; i++ {
		ft := t.Field(i)
		fu := u.Field(i)
		if ft.Anonymous != fu.Anonymous || ft.Name != fu.Name ||
			ft.PkgPath != fu.PkgPath || ft.Tag != fu.Tag ||
			!ft.Type.Identical(fu.Type) {
			return false
		}
	}
	return true
}

func (t *structType) str(w *bytes.Buffer) {
	w.WriteString("struct{")
	for i, v := range t.fields {
		if !v.Anonymous {
			w.Write(dict.S(v.Name))
			w.WriteByte(' ')
		}
		safeTypeStr(v.Type, w)
		if i != len(t.fields)-1 {
			w.WriteString("; ")
		}
	}
	w.WriteByte('}')
}

// ------------------------------------------------------------------ tupleType

type tupleType struct {
	typeBase
	elements []Type
	flgs     flags
}

func newTupleType(elements []Type) *tupleType {
	t := &tupleType{elements: elements}
	for _, v := range elements {
		if v != nil {
			t.flgs = t.flgs | v.flags()
		}
	}
	t.kind = Tuple
	t.typ = t
	t.align, t.fieldAlign, t.size = measure(t)
	return t
}

func (t *tupleType) Elements() []Type      { return t.elements }
func (t *tupleType) field(i int) Type      { return t.elements[i] }
func (t *tupleType) numField() int         { return len(t.elements) }
func (t *tupleType) setOffset(int, uint64) {}
func (t *tupleType) flags() flags          { return t.flgs }

func (t *tupleType) Identical(u Type) bool {
	if u.Kind() != Tuple {
		return false
	}

	et := t.Elements()
	eu := u.Elements()
	if len(et) != len(eu) {
		return false
	}

	for i, v := range et {
		if !v.Identical(eu[i]) {
			return false
		}
	}
	return true
}

func (t *tupleType) str(w *bytes.Buffer) {
	if len(t.elements) == 0 {
		return
	}

	w.WriteByte('(')
	for i, v := range t.elements {
		safeTypeStr(v, w)
		if i != len(t.elements)-1 {
			w.WriteString(", ")
		}
	}
	w.WriteByte(')')
}

func safeTypeStr(t Type, w *bytes.Buffer) {
	if t != nil {
		t.str(w)
		return
	}

	w.WriteString("<nil>")
}

// ----------------------------------------------------------------------------

func align(n, to uint64) uint64 {
	n += to - 1
	return n - n%to
}

func measure(t Type) (algn, fieldAlign int, size uint64) {
	algn = 1
	var ofs uint64
	for i := 0; i < t.numField(); i++ {
		ft := t.field(i)
		if ft == nil || ft.Kind() == Invalid {
			continue
		}

		falign := mathutil.Max(ft.FieldAlign(), 1)
		sz := ft.Size()
		ofs = align(ofs, uint64(falign))
		t.setOffset(i, ofs)
		algn = mathutil.Max(algn, falign)
		ofs += sz
	}
	return algn, algn, align(ofs, uint64(algn))
}
